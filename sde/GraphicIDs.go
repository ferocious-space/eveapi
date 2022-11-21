// Code generated by YAML2GO. DO NOT EDIT.

package sde

import (
	yamlv3 "gopkg.in/yaml.v3"
	"os"
)

type GraphicIDMap map[int32]GraphicID

func (x *GraphicIDMap) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return yamlv3.NewDecoder(f).Decode(x)
}
func (x GraphicIDMap) Get(ID int32) *GraphicID {
	if a, ok := x[ID]; ok {
		return &a
	}
	return nil
}

type GraphicID struct {
	Description    *string            `bson:"description,omitempty" db:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	GraphicFile    *string            `bson:"graphicFile,omitempty" db:"graphicFile,omitempty" json:"graphicFile,omitempty" yaml:"graphicFile,omitempty"`
	IconInfo       *GraphicIDIconInfo `bson:"iconInfo,omitempty" db:"iconInfo,omitempty" json:"iconInfo,omitempty" yaml:"iconInfo,omitempty"`
	SofFactionName *string            `bson:"sofFactionName,omitempty" db:"sofFactionName,omitempty" json:"sofFactionName,omitempty" yaml:"sofFactionName,omitempty"`
	SofHullName    *string            `bson:"sofHullName,omitempty" db:"sofHullName,omitempty" json:"sofHullName,omitempty" yaml:"sofHullName,omitempty"`
	SofRaceName    *string            `bson:"sofRaceName,omitempty" db:"sofRaceName,omitempty" json:"sofRaceName,omitempty" yaml:"sofRaceName,omitempty"`
	SofLayout      []string           `bson:"sofLayout,omitempty" db:"sofLayout,omitempty" json:"sofLayout,omitempty" yaml:"sofLayout,omitempty"`
}
type GraphicIDIconInfo struct {
	Folder *string `bson:"folder,omitempty" db:"folder,omitempty" json:"folder,omitempty" yaml:"folder,omitempty"`
}
