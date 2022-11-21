// Code generated by YAML2GO. DO NOT EDIT.

package sde

import (
	yamlv3 "gopkg.in/yaml.v3"
	"os"
)

type AncestryMap map[int32]Ancestry

func (x *AncestryMap) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return yamlv3.NewDecoder(f).Decode(x)
}
func (x AncestryMap) Get(ID int32) *Ancestry {
	if a, ok := x[ID]; ok {
		return &a
	}
	return nil
}

type Ancestry struct {
	BloodlineID      *int32                 `bson:"bloodlineID,omitempty" db:"bloodlineID,omitempty" json:"bloodlineID,omitempty" yaml:"bloodlineID,omitempty"`
	Charisma         *int32                 `bson:"charisma,omitempty" db:"charisma,omitempty" json:"charisma,omitempty" yaml:"charisma,omitempty"`
	DescriptionID    *AncestryDescriptionID `bson:"descriptionID,omitempty" db:"descriptionID,omitempty" json:"descriptionID,omitempty" yaml:"descriptionID,omitempty"`
	IconID           *int32                 `bson:"iconID,omitempty" db:"iconID,omitempty" json:"iconID,omitempty" yaml:"iconID,omitempty"`
	Intelligence     *int32                 `bson:"intelligence,omitempty" db:"intelligence,omitempty" json:"intelligence,omitempty" yaml:"intelligence,omitempty"`
	Memory           *int32                 `bson:"memory,omitempty" db:"memory,omitempty" json:"memory,omitempty" yaml:"memory,omitempty"`
	NameID           *AncestryNameID        `bson:"nameID,omitempty" db:"nameID,omitempty" json:"nameID,omitempty" yaml:"nameID,omitempty"`
	Perception       *int32                 `bson:"perception,omitempty" db:"perception,omitempty" json:"perception,omitempty" yaml:"perception,omitempty"`
	ShortDescription *string                `bson:"shortDescription,omitempty" db:"shortDescription,omitempty" json:"shortDescription,omitempty" yaml:"shortDescription,omitempty"`
	Willpower        *int32                 `bson:"willpower,omitempty" db:"willpower,omitempty" json:"willpower,omitempty" yaml:"willpower,omitempty"`
}
type AncestryDescriptionID struct {
	De *string `bson:"de,omitempty" db:"de,omitempty" json:"de,omitempty" yaml:"de,omitempty"`
	En *string `bson:"en,omitempty" db:"en,omitempty" json:"en,omitempty" yaml:"en,omitempty"`
	Es *string `bson:"es,omitempty" db:"es,omitempty" json:"es,omitempty" yaml:"es,omitempty"`
	Fr *string `bson:"fr,omitempty" db:"fr,omitempty" json:"fr,omitempty" yaml:"fr,omitempty"`
	Ja *string `bson:"ja,omitempty" db:"ja,omitempty" json:"ja,omitempty" yaml:"ja,omitempty"`
	Ko *string `bson:"ko,omitempty" db:"ko,omitempty" json:"ko,omitempty" yaml:"ko,omitempty"`
	Ru *string `bson:"ru,omitempty" db:"ru,omitempty" json:"ru,omitempty" yaml:"ru,omitempty"`
	Zh *string `bson:"zh,omitempty" db:"zh,omitempty" json:"zh,omitempty" yaml:"zh,omitempty"`
}
type AncestryNameID struct {
	De *string `bson:"de,omitempty" db:"de,omitempty" json:"de,omitempty" yaml:"de,omitempty"`
	En *string `bson:"en,omitempty" db:"en,omitempty" json:"en,omitempty" yaml:"en,omitempty"`
	Es *string `bson:"es,omitempty" db:"es,omitempty" json:"es,omitempty" yaml:"es,omitempty"`
	Fr *string `bson:"fr,omitempty" db:"fr,omitempty" json:"fr,omitempty" yaml:"fr,omitempty"`
	Ja *string `bson:"ja,omitempty" db:"ja,omitempty" json:"ja,omitempty" yaml:"ja,omitempty"`
	Ko *string `bson:"ko,omitempty" db:"ko,omitempty" json:"ko,omitempty" yaml:"ko,omitempty"`
	Ru *string `bson:"ru,omitempty" db:"ru,omitempty" json:"ru,omitempty" yaml:"ru,omitempty"`
	Zh *string `bson:"zh,omitempty" db:"zh,omitempty" json:"zh,omitempty" yaml:"zh,omitempty"`
}
