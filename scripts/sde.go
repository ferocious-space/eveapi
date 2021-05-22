package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

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
	if err := filepath.Walk(
		"./testing", func(path string, info fs.FileInfo, err error) error {
			runtime.GC()
			if info.IsDir() {
				return nil
			}
			fmt.Println("Processing", filepath.Base(path))
			if filepath.Base(path) == "region.staticdata" {
				file, err := yamlparser.ParseFile(path)
				if err != nil {
					return err
				}
				regionYaml = append(regionYaml, file)
			}
			if filepath.Base(path) == "constellation.staticdata" {
				file, err := yamlparser.ParseFile(path)
				if err != nil {
					return err
				}
				constellationYaml = append(constellationYaml, file)
			}
			if filepath.Base(path) == "solarsystem.staticdata" {
				file, err := yamlparser.ParseFile(path)
				if err != nil {
					return err
				}
				solarsystemYaml = append(solarsystemYaml, file)
			}
			if filepath.Ext(path) != ".yaml" {
				return nil
			}
			fmt.Println("generating", path)
			typeInterface, err := yamlparser.ParseFile(path)
			if err != nil {
				return err
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
				return err
			}
			if err := os.WriteFile(
				filepath.Join("./sde", typeName+".go"), file,
				os.ModePerm,
			); err != nil {
				return err
			}
			return nil
		},
	); err != nil {
		fmt.Println("walking", err.Error())
		return
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
