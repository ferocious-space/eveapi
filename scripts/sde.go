package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-openapi/inflect"

	"github.com/ferocious-space/eveapi"
	"github.com/ferocious-space/eveapi/internal/pkg/yamlparser"
)

func sdegen() {
	err := eveapi.DownloadSDE(&http.Client{}, "./testing", false)
	if err != nil {
		return
	}
	if err := filepath.Walk(
		"./testing", func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
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
		fmt.Println("waling", err.Error())
		return
	}
	return
}
