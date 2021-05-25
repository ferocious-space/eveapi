/*
 *    Copyright 2021 FerociousBite and Contributors
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package eveapi

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	openapiRuntime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/ferocious-space/eveapi/esi"
)

func NewAPIClient(cachedClient *http.Client) *esi.EVESwaggerInterface {
	apiRuntime := httptransport.NewWithClient(esi.DefaultHost, esi.DefaultBasePath, esi.DefaultSchemes, cachedClient)
	apiRuntime.Consumers[openapiRuntime.JSONMime] = openapiRuntime.ConsumerFunc(
		func(reader io.Reader, data interface{}) error {
			dec := json.NewDecoder(reader)
			dec.UseNumber()
			return dec.Decode(data)
		},
	)
	apiRuntime.Producers[openapiRuntime.JSONMime] = openapiRuntime.ProducerFunc(
		func(writer io.Writer, data interface{}) error {
			enc := json.NewEncoder(writer)
			enc.SetEscapeHTML(false)
			return enc.Encode(data)
		},
	)
	return esi.New(apiRuntime, strfmt.Default)
}

type downloadInfo struct {
	Etag         string
	LastModified time.Time
}

const sdePath = "https://eve-static-data-export.s3-eu-west-1.amazonaws.com/tranquility/sde.zip"

func DownloadSDE(client *http.Client, targetDir string, extract bool) error {
	path := ""
	if filepath.IsAbs(targetDir) {
		path = targetDir
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		path = filepath.Clean(filepath.Join(cwd, targetDir))
	}
	if info, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {
		if !info.IsDir() {
			return errors.New("targetDir must be directory")
		}
	}
	info := &downloadInfo{}
	infoFile := new(os.File)
	infoPath := filepath.Join(path, "sdeInfo.yaml")
	if _, err := os.Stat(infoPath); os.IsNotExist(err) {
		infoFile, err = os.Create(infoPath)
		if err != nil {
			return err
		}
		defer infoFile.Close()
		err = yaml.NewEncoder(infoFile).Encode(&info)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {
		infoFile, err = os.Open(infoPath)
		if err != nil {
			return err
		}
		defer infoFile.Close()
	}
	err := yaml.NewDecoder(infoFile).Decode(&info)
	if err != nil {
		info = &downloadInfo{}
	}
	sdeFilePath := filepath.Join(path, "sde.zip")
	force := false
	if _, err := os.Stat(sdeFilePath); os.IsNotExist(err) {
		force = true
	} else if err != nil {
		return err
	}

	dlRequest := &http.Request{}
	dlRequest, err = http.NewRequest("HEAD", sdePath, nil)
	if err != nil {
		return err
	}
	dlRequest.Header.Set("cache-control", "no-store, no-cache")
	if !force {

		if info.Etag != "" {
			dlRequest.Header.Set("if-none-match", info.Etag)
		}
		if !info.LastModified.IsZero() {
			dlRequest.Header.Set("if-modified-since", info.LastModified.Format(time.RFC1123))
		}
	}
	if resp, err := client.Do(dlRequest); err != nil {
		return err
	} else {
		switch resp.StatusCode {
		case http.StatusOK:
			if dlRequest.Method == "HEAD" {
				dlRequest.Method = "GET"
				resp, err = client.Do(dlRequest)
				if err != nil {
					return err
				}
				defer func() {
					_, _ = io.Copy(io.Discard, resp.Body)
					_ = resp.Body.Close()
				}()
			}
			expires, err := time.Parse(time.RFC1123, resp.Header.Get("last-modified"))
			if err != nil {
				return err
			}
			sdeFile, err := os.Create(sdeFilePath)
			if err != nil {
				return err
			}
			w, err := io.Copy(sdeFile, resp.Body)
			if err != nil {
				return err
			}
			if w != resp.ContentLength {
				return errors.Errorf("dowload size differs %d != %d", w, resp.ContentLength)
			}
			err = sdeFile.Close()
			if err != nil {
				return err
			}
			info.Etag = resp.Header.Get("etag")
			info.LastModified = expires
			out, err := yaml.Marshal(&info)
			if err != nil {
				return err
			}
			err = os.WriteFile(infoPath, out, os.ModePerm)
			if err != nil {
				return err
			}
		case http.StatusNotModified:
		default:
			return errors.Errorf("unexpected status code: %d", resp.StatusCode)
		}
	}
	if extract {
		var zipSize int64
		if stat, err := os.Stat(sdeFilePath); err != nil {
			return err
		} else {
			zipSize = stat.Size()
		}
		zipFile, err := os.Open(sdeFilePath)
		if err != nil {
			return err
		}
		zReader, err := zip.NewReader(zipFile, zipSize)
		if err != nil {
			return err
		}
		defer zipFile.Close()
		guard := sync.Mutex{}

		parallelExecution := make(chan struct{}, runtime.NumCPU()*4+1)
		wg := sync.WaitGroup{}
		for i := range zReader.File {
			parallelExecution <- struct{}{}
			wg.Add(1)
			go func(zFile *zip.File) {
				defer func() {
					<-parallelExecution
					wg.Done()
				}()
				if err := extractFn(
					zFile, filepath.Dir(sdeFilePath), func() *sync.Mutex {
						return &guard
					},
				); err != nil {
					log.Fatal(err)
				}
			}(zReader.File[i])

		}
		wg.Wait()
	}
	return nil
}

func extractFn(f *zip.File, path string, lockKeeper func() *sync.Mutex) error {
	lockKeeper().Lock()
	defer lockKeeper().Unlock()

	targetPath := filepath.Join(path, f.Name)
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()
	if !strings.HasPrefix(targetPath, filepath.Clean(filepath.Join(path)+string(os.PathSeparator))) {
		return errors.Errorf("%s: illegal file path", targetPath)
	}
	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(targetPath, os.ModePerm); err != nil {
			return err
		}
	} else {
		if err := os.MkdirAll(filepath.Dir(targetPath), os.ModePerm); err != nil {
			return err
		}
		out, err := os.Create(targetPath)
		if err != nil {
			return err
		}
		defer out.Close()
		wSize, err := io.Copy(out, rc)
		if err != nil {
			return err
		}
		if uint64(wSize) != f.UncompressedSize64 {
			return errors.New("decompress: file sizes does not match")
		}
	}
	return nil
}
