package sde

import (
	"bytes"
	"fmt"
	"github.com/goccy/go-json"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/valyala/bytebufferpool"
	"go.etcd.io/bbolt"
	yamlv3 "gopkg.in/yaml.v3"
)

var bufferPool = bytebufferpool.Pool{}

func getType(myvar interface{}) []byte {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return []byte(t.Elem().Name())
	} else {
		return []byte(t.Name())
	}
}

func SDEBolt(boltPath string) (*bbolt.DB, error) {
	opts := bbolt.DefaultOptions
	opts.FreelistType = bbolt.FreelistMapType
	opts.InitialMmapSize = 256 << 20
	open, err := bbolt.Open(boltPath, os.ModePerm, opts)
	if err != nil {
		return nil, err
	}
	open.MaxBatchSize = 1000
	open.MaxBatchDelay = 10 * time.Millisecond
	open.AllocSize = 1 << 20

	return open, nil
}

func ConvertDogmaAttributes(materialsPath string, db *bbolt.DB) error {
	materials := make(DogmaAttributeMap)
	err := materials.Load(materialsPath)
	if err != nil {
		return err
	}

	wait := sync.WaitGroup{}
	errorChannel := make(chan error, 1)
	parallel := make(chan struct{}, 1000)
	wait.Add(len(materials))
	for k, v := range materials {
		parallel <- struct{}{}
		go func(k int32, inv DogmaAttribute) {
			defer wait.Done()
			err := db.Batch(
				func(tx *bbolt.Tx) error {
					btx, err := tx.CreateBucketIfNotExists(getType(inv))
					if err != nil {
						return err
					}
					buf := bufferPool.Get()
					err = json.NewEncoder(buf).Encode(inv)
					if err != nil {
						return err
					}
					err = btx.Put([]byte(fmt.Sprintf("%d", k)), buf.Bytes())
					if err != nil {
						return err
					}

					bufferPool.Put(buf)
					return nil
				},
			)
			if err != nil {
				errorChannel <- err
			}
			<-parallel
		}(k, v)

	}
	go func() {
		wait.Wait()
		close(errorChannel)
	}()
	select {
	case err, closed := <-errorChannel:
		if !closed {
			return err
		} else {
			return nil
		}
	}
}

func ConvertDogmaEffects(materialsPath string, db *bbolt.DB) error {
	materials := make(DogmaEffectMap)
	err := materials.Load(materialsPath)
	if err != nil {
		return err
	}

	wait := sync.WaitGroup{}
	errorChannel := make(chan error, 1)
	parallel := make(chan struct{}, 1000)
	wait.Add(len(materials))
	for k, v := range materials {
		parallel <- struct{}{}
		bktType := getType(materials[0])
		go func(k int32, inv DogmaEffect) {
			defer wait.Done()
			err := db.Batch(
				func(tx *bbolt.Tx) error {
					btx, err := tx.CreateBucketIfNotExists(bktType)
					if err != nil {
						return err
					}
					buf := bufferPool.Get()
					err = json.NewEncoder(buf).Encode(inv)
					if err != nil {
						return err
					}
					err = btx.Put([]byte(fmt.Sprintf("%d", k)), buf.Bytes())
					if err != nil {
						return err
					}

					bufferPool.Put(buf)
					return nil
				},
			)
			if err != nil {
				errorChannel <- err
			}
			<-parallel
		}(k, v)

	}
	go func() {
		wait.Wait()
		close(errorChannel)
	}()
	select {
	case err, closed := <-errorChannel:
		if !closed {
			return err
		} else {
			return nil
		}
	}
}

func QueryBolt(db *bbolt.DB, key interface{}, value interface{}) error {
	return db.View(
		func(tx *bbolt.Tx) error {
			bucket := tx.Bucket(getType(value))
			if bucket == nil {
				return errors.Errorf("bucket %s not found", string(getType(value)))
			}

			get := bucket.Get([]byte(fmt.Sprintf("%d", key)))
			if get != nil {
				err := json.NewDecoder(bytes.NewBuffer(get)).Decode(value)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func ConvertTypeDogma(materialsPath string, db *bbolt.DB) error {
	materials := make(TypeDogmaMap)
	err := materials.Load(materialsPath)
	if err != nil {
		return err
	}

	wait := sync.WaitGroup{}
	errorChannel := make(chan error, 1)
	parallel := make(chan struct{}, 1000)
	wait.Add(len(materials))
	bktType := getType(materials[0])
	for k, v := range materials {
		parallel <- struct{}{}
		go func(k int32, inv TypeDogma) {
			defer wait.Done()
			err := db.Batch(
				func(tx *bbolt.Tx) error {
					btx, err := tx.CreateBucketIfNotExists(bktType)
					if err != nil {
						return err
					}
					buf := bufferPool.Get()
					err = json.NewEncoder(buf).Encode(inv)
					if err != nil {
						return err
					}
					err = btx.Put([]byte(fmt.Sprintf("%d", k)), buf.Bytes())
					if err != nil {
						return err
					}

					bufferPool.Put(buf)
					return nil
				},
			)
			if err != nil {
				errorChannel <- err
			}
			<-parallel
		}(k, v)

	}
	go func() {
		wait.Wait()
		close(errorChannel)
	}()
	select {
	case err, closed := <-errorChannel:
		if !closed {
			return err
		} else {
			return nil
		}
	}
}

func ConvertTypes(materialsPath string, db *bbolt.DB) error {
	materials := make(TypeIDMap)
	err := materials.Load(materialsPath)
	if err != nil {
		return err
	}

	wait := sync.WaitGroup{}
	errorChannel := make(chan error, 1)
	parallel := make(chan struct{}, 1000)
	wait.Add(len(materials))
	bktType := getType(materials[0])
	for k, v := range materials {
		parallel <- struct{}{}
		go func(k int32, inv TypeID) {
			defer wait.Done()
			err := db.Batch(
				func(tx *bbolt.Tx) error {
					btx, err := tx.CreateBucketIfNotExists(bktType)
					if err != nil {
						return err
					}
					buf := bufferPool.Get()
					err = json.NewEncoder(buf).Encode(inv)
					if err != nil {
						return err
					}
					err = btx.Put([]byte(fmt.Sprintf("%d", k)), buf.Bytes())
					if err != nil {
						return err
					}

					bufferPool.Put(buf)
					return nil
				},
			)
			if err != nil {
				errorChannel <- err
			}
			<-parallel
		}(k, v)

	}
	go func() {
		wait.Wait()
		close(errorChannel)
	}()
	select {
	case err, closed := <-errorChannel:
		if !closed {
			return err
		} else {
			return nil
		}
	}
}

func ConvertMaterials(materialsPath string, db *bbolt.DB) error {
	materials := make(TypeMaterialMap)
	err := materials.Load(materialsPath)
	if err != nil {
		return err
	}

	wait := sync.WaitGroup{}
	errorChannel := make(chan error, 1)
	parallel := make(chan struct{}, 1000)
	wait.Add(len(materials))
	bktType := getType(materials[0])
	for k, v := range materials {
		parallel <- struct{}{}
		go func(k int32, inv TypeMaterial) {
			defer wait.Done()
			err := db.Batch(
				func(tx *bbolt.Tx) error {
					btx, err := tx.CreateBucketIfNotExists(bktType)
					if err != nil {
						return err
					}
					buf := bufferPool.Get()
					err = json.NewEncoder(buf).Encode(inv)
					if err != nil {
						return err
					}
					err = btx.Put([]byte(fmt.Sprintf("%d", k)), buf.Bytes())
					if err != nil {
						return err
					}

					bufferPool.Put(buf)
					return nil
				},
			)
			if err != nil {
				errorChannel <- err
			}
			<-parallel
		}(k, v)

	}
	go func() {
		wait.Wait()
		close(errorChannel)
	}()
	select {
	case err, closed := <-errorChannel:
		if !closed {
			return err
		} else {
			return nil
		}
	}
}

func ConvertNames(namesPath string, db *bbolt.DB) error {
	names := make(InvNameList, 0)
	err := names.Load(namesPath)
	if err != nil {
		return err
	}
	wait := sync.WaitGroup{}
	errorChannel := make(chan error, 1)
	parallel := make(chan struct{}, 1000)
	wait.Add(len(names))
	bktType := getType(names[0])
	for _, v := range names {
		parallel <- struct{}{}
		go func(inv InvName) {
			defer wait.Done()
			err := db.Batch(
				func(tx *bbolt.Tx) error {
					btx, err := tx.CreateBucketIfNotExists(bktType)
					if err != nil {
						return err
					}
					buf := bufferPool.Get()
					err = json.NewEncoder(buf).Encode(inv)
					if err != nil {
						return err
					}
					err = btx.Put([]byte(fmt.Sprintf("%d", inv.ItemID)), buf.Bytes())
					if err != nil {
						return err
					}

					bufferPool.Put(buf)
					return nil
				},
			)
			if err != nil {
				errorChannel <- err
			}
			<-parallel
		}(v)

	}
	go func() {
		wait.Wait()
		close(errorChannel)
	}()
	select {
	case err, closed := <-errorChannel:
		if !closed {
			return err
		} else {
			return nil
		}
	}
}

func ConvertUniverse(fromPath string, db *bbolt.DB) error {
	var wait sync.WaitGroup
	errorChannel := make(chan error, 1)
	parallel := make(chan struct{}, 1000)
	wait.Add(1)
	go func() {
		defer wait.Done()

		err := filepath.WalkDir(
			fromPath, func(path string, info fs.DirEntry, err error) error {
				if info.IsDir() {
					return nil
				}
				wait.Add(1)
				parallel <- struct{}{}
				go func(path string) {
					defer wait.Done()
					var data interface{}
					var key interface{}
					file := new(os.File)
					switch filepath.Base(path) {
					case "solarsystem.staticdata":
						data = new(Solarsystem)
					case "constellation.staticdata":
						data = new(Constellation)
					case "region.staticdata":
						data = new(Region)
					default:
						panic("wtf")
					}
					file, err := os.Open(path)
					if err != nil {
						errorChannel <- err
						return
					}
					defer file.Close()
					err = yamlv3.NewDecoder(file).Decode(data)
					if err != nil {
						errorChannel <- err
						return
					}
					buf := bufferPool.Get()
					err = json.NewEncoder(buf).Encode(data)
					if err != nil {
						errorChannel <- err
						return
					}
					switch filepath.Base(path) {
					case "solarsystem.staticdata":
						key = *data.(*Solarsystem).SolarSystemID
					case "constellation.staticdata":
						key = *data.(*Constellation).ConstellationID
					case "region.staticdata":
						key = *data.(*Region).RegionID
					default:
						panic("wtf")
					}
					err = db.Batch(
						func(tx *bbolt.Tx) error {
							btx, err := tx.CreateBucketIfNotExists(getType(data))
							if err != nil {
								return err
							}
							err = btx.Put([]byte(fmt.Sprintf("%d", key)), buf.Bytes())
							if err != nil {
								return err
							}
							return nil
						},
					)
					if err != nil {
						errorChannel <- err
						return
					}

					bufferPool.Put(buf)
					<-parallel
				}(path)
				return nil
			},
		)
		if err != nil {
			errorChannel <- err
		}
	}()
	go func() {
		wait.Wait()
		close(errorChannel)
	}()
	select {
	case err, closed := <-errorChannel:
		if !closed {
			return err
		} else {
			return nil
		}
	}
}
