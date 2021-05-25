package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/go-openapi/inflect"

	"github.com/ferocious-space/eveapi"
	"github.com/ferocious-space/eveapi/internal/pkg/yamlparser"
)

func sdegen() {
	err := eveapi.DownloadSDE(&http.Client{}, "./testing", true)
	if err != nil {
		return
	}
	regionYaml := make([]reflect.Type, 0)
	constellationYaml := make([]reflect.Type, 0)
	solarsystemYaml := make([]reflect.Type, 0)
	var wait sync.WaitGroup
	errChannel := make(chan error, 1)
	defer close(errChannel)
	waitChannel := make(chan struct{}, 1)
	parallel := make(chan struct{}, 10)
	if err := filepath.Walk(
		"./testing", func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			parallel <- struct{}{}
			wait.Add(1)
			go func(path string) {
				defer func() {
					_ = <-parallel
				}()
				defer wait.Done()
				fmt.Println("Processing", filepath.Base(path))
				if filepath.Base(path) == "region.staticdata" {
					file, err := yamlparser.ParseFile(path)
					if err != nil {
						errChannel <- err
					}
					regionYaml = append(regionYaml, file)
				}
				if filepath.Base(path) == "constellation.staticdata" {
					file, err := yamlparser.ParseFile(path)
					if err != nil {
						errChannel <- err
					}
					constellationYaml = append(constellationYaml, file)
				}
				if filepath.Base(path) == "solarsystem.staticdata" {
					file, err := yamlparser.ParseFile(path)
					if err != nil {
						errChannel <- err
					}
					solarsystemYaml = append(solarsystemYaml, file)
				}
				if filepath.Ext(path) != ".yaml" {
					return
				}
				fmt.Println("generating", path)
				typeInterface, err := yamlparser.ParseFile(path)
				if err != nil {
					errChannel <- err
				}
				typeName := inflect.Capitalize(
					inflect.Camelize(
						strings.TrimSuffix(
							filepath.Base(path),
							filepath.Ext(path),
						),
					),
				)
				file, err := yamlparser.GenerateType(
					typeInterface,
					typeName,
					"github.com/ferocious-space/eveapi/sde",
					"sde",
				)
				if err != nil {
					errChannel <- err
				}
				if err := os.WriteFile(
					filepath.Join("./sde", typeName+".go"), file,
					os.ModePerm,
				); err != nil {
					errChannel <- err
				}
			}(path)
			return nil
		},
	); err != nil {
		fmt.Println("walking", err.Error())
		return
	}
	go func() {
		defer close(waitChannel)
		wait.Wait()
	}()
	select {
	case err := <-errChannel:
		fmt.Println(err.Error())
		return
	case <-waitChannel:
	}
	regionType, err := yamlparser.GenerateType(yamlparser.DeepMergeList(regionYaml), "Regions", "github.com/fericous-space/eveapi/sde", "sde")
	if err != nil {
		return
	}
	if err := os.WriteFile(
		filepath.Join("./sde", "Regions.go"), regionType,
		os.ModePerm,
	); err != nil {
		return
	}
	constellationType, err := yamlparser.GenerateType(yamlparser.DeepMergeList(constellationYaml), "Constellations", "github.com/fericous-space/eveapi/sde", "sde")
	if err != nil {
		return
	}
	if err := os.WriteFile(
		filepath.Join("./sde", "Constellations.go"), constellationType,
		os.ModePerm,
	); err != nil {
		return
	}
	solarsystemType, err := yamlparser.GenerateType(yamlparser.DeepMergeList(solarsystemYaml), "Solarsystems", "github.com/fericous-space/eveapi/sde", "sde")
	if err != nil {
		return
	}
	if err := os.WriteFile(
		filepath.Join("./sde", "Solarsystems.go"), solarsystemType,
		os.ModePerm,
	); err != nil {
		return
	}
}
