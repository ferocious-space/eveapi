// Code generated by YAML2GO. DO NOT EDIT.

package sde

import (
	yamlv3 "gopkg.in/yaml.v3"
	"os"
)

type GroupIDMap map[int32]GroupID

func (x *GroupIDMap) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return yamlv3.NewDecoder(f).Decode(x)
}
func (x GroupIDMap) Get(ID int32) *GroupID {
	if a, ok := x[ID]; ok {
		return &a
	}
	return nil
}

type GroupID struct {
	Anchorable           *bool        `bson:"anchorable,omitempty" db:"anchorable,omitempty" json:"anchorable,omitempty" yaml:"anchorable,omitempty"`
	Anchored             *bool        `bson:"anchored,omitempty" db:"anchored,omitempty" json:"anchored,omitempty" yaml:"anchored,omitempty"`
	CategoryID           *int32       `bson:"categoryID,omitempty" db:"categoryID,omitempty" json:"categoryID,omitempty" yaml:"categoryID,omitempty"`
	FittableNonSingleton *bool        `bson:"fittableNonSingleton,omitempty" db:"fittableNonSingleton,omitempty" json:"fittableNonSingleton,omitempty" yaml:"fittableNonSingleton,omitempty"`
	Name                 *GroupIDName `bson:"name,omitempty" db:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Published            *bool        `bson:"published,omitempty" db:"published,omitempty" json:"published,omitempty" yaml:"published,omitempty"`
	UseBasePrice         *bool        `bson:"useBasePrice,omitempty" db:"useBasePrice,omitempty" json:"useBasePrice,omitempty" yaml:"useBasePrice,omitempty"`
	IconID               *int32       `bson:"iconID,omitempty" db:"iconID,omitempty" json:"iconID,omitempty" yaml:"iconID,omitempty"`
}
type GroupIDName struct {
	De *string `bson:"de,omitempty" db:"de,omitempty" json:"de,omitempty" yaml:"de,omitempty"`
	En *string `bson:"en,omitempty" db:"en,omitempty" json:"en,omitempty" yaml:"en,omitempty"`
	Fr *string `bson:"fr,omitempty" db:"fr,omitempty" json:"fr,omitempty" yaml:"fr,omitempty"`
	Ja *string `bson:"ja,omitempty" db:"ja,omitempty" json:"ja,omitempty" yaml:"ja,omitempty"`
	Ru *string `bson:"ru,omitempty" db:"ru,omitempty" json:"ru,omitempty" yaml:"ru,omitempty"`
	Zh *string `bson:"zh,omitempty" db:"zh,omitempty" json:"zh,omitempty" yaml:"zh,omitempty"`
	Es *string `bson:"es,omitempty" db:"es,omitempty" json:"es,omitempty" yaml:"es,omitempty"`
	It *string `bson:"it,omitempty" db:"it,omitempty" json:"it,omitempty" yaml:"it,omitempty"`
}
