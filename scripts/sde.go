package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	etcdIOutil "github.com/coreos/etcd/pkg/ioutil"
	"github.com/go-openapi/inflect"
	"github.com/hashicorp/go-cleanhttp"
	"go.uber.org/zap"

	"github.com/ferocious-space/eveapi"
	"github.com/ferocious-space/eveapi/internal/pkg/yamlparser"
)

func sdegen() {
	log := zap.NewExample()
	err := eveapi.DownloadSDE(cleanhttp.DefaultClient(), "./testing", false)
	if err != nil {
		log.Fatal("downloading SDE", zap.Error(err))
	}
	if err := filepath.Walk(
		"./testing", func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if filepath.Ext(path) != ".yaml" {
				return nil
			}
			log.Info("generating", zap.String("file", path))
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
			if err := etcdIOutil.WriteAndSyncFile(
				filepath.Join("./sde", typeName+".go"), file,
				os.ModePerm,
			); err != nil {
				return err
			}
			return nil
		},
	); err != nil {
		log.Fatal("walking", zap.Error(err))
		return
	}
	return
}
