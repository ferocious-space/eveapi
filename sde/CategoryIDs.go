// Code generated by YAML2GO. DO NOT EDIT.

package sde

import (
	yamlv3 "gopkg.in/yaml.v3"
	"os"
)

type CategoryIDMap map[int32]CategoryID

func (x *CategoryIDMap) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return yamlv3.NewDecoder(f).Decode(x)
}
func (x CategoryIDMap) Get(ID int32) *CategoryID {
	if a, ok := x[ID]; ok {
		return &a
	}
	return nil
}

type CategoryID struct {
	Name      *CategoryIDName `bson:"name,omitempty" db:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Published *bool           `bson:"published,omitempty" db:"published,omitempty" json:"published,omitempty" yaml:"published,omitempty"`
	IconID    *int32          `bson:"iconID,omitempty" db:"iconID,omitempty" json:"iconID,omitempty" yaml:"iconID,omitempty"`
}
type CategoryIDName struct {
	De *string `bson:"de,omitempty" db:"de,omitempty" json:"de,omitempty" yaml:"de,omitempty"`
	En *string `bson:"en,omitempty" db:"en,omitempty" json:"en,omitempty" yaml:"en,omitempty"`
	Es *string `bson:"es,omitempty" db:"es,omitempty" json:"es,omitempty" yaml:"es,omitempty"`
	Fr *string `bson:"fr,omitempty" db:"fr,omitempty" json:"fr,omitempty" yaml:"fr,omitempty"`
	Ja *string `bson:"ja,omitempty" db:"ja,omitempty" json:"ja,omitempty" yaml:"ja,omitempty"`
	Ru *string `bson:"ru,omitempty" db:"ru,omitempty" json:"ru,omitempty" yaml:"ru,omitempty"`
	Zh *string `bson:"zh,omitempty" db:"zh,omitempty" json:"zh,omitempty" yaml:"zh,omitempty"`
}
