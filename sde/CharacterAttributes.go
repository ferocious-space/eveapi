// Code generated by YAML2GO. DO NOT EDIT.

package sde

import (
	yamlv3 "gopkg.in/yaml.v3"
	"os"
)

type CharacterAttributeMap map[int32]CharacterAttribute

func (x *CharacterAttributeMap) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return yamlv3.NewDecoder(f).Decode(x)
}
func (x CharacterAttributeMap) Get(ID int32) *CharacterAttribute {
	if a, ok := x[ID]; ok {
		return &a
	}
	return nil
}

type CharacterAttribute struct {
	Description      *string                   `bson:"description,omitempty" db:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	IconID           *int32                    `bson:"iconID,omitempty" db:"iconID,omitempty" json:"iconID,omitempty" yaml:"iconID,omitempty"`
	NameID           *CharacterAttributeNameID `bson:"nameID,omitempty" db:"nameID,omitempty" json:"nameID,omitempty" yaml:"nameID,omitempty"`
	Notes            *string                   `bson:"notes,omitempty" db:"notes,omitempty" json:"notes,omitempty" yaml:"notes,omitempty"`
	ShortDescription *string                   `bson:"shortDescription,omitempty" db:"shortDescription,omitempty" json:"shortDescription,omitempty" yaml:"shortDescription,omitempty"`
}
type CharacterAttributeNameID struct {
	De *string `bson:"de,omitempty" db:"de,omitempty" json:"de,omitempty" yaml:"de,omitempty"`
	En *string `bson:"en,omitempty" db:"en,omitempty" json:"en,omitempty" yaml:"en,omitempty"`
	Fr *string `bson:"fr,omitempty" db:"fr,omitempty" json:"fr,omitempty" yaml:"fr,omitempty"`
	Ja *string `bson:"ja,omitempty" db:"ja,omitempty" json:"ja,omitempty" yaml:"ja,omitempty"`
	Ko *string `bson:"ko,omitempty" db:"ko,omitempty" json:"ko,omitempty" yaml:"ko,omitempty"`
	Ru *string `bson:"ru,omitempty" db:"ru,omitempty" json:"ru,omitempty" yaml:"ru,omitempty"`
	Zh *string `bson:"zh,omitempty" db:"zh,omitempty" json:"zh,omitempty" yaml:"zh,omitempty"`
}
