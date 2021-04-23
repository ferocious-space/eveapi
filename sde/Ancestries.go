// Code generated by YAML2GO. DO NOT EDIT.

package sde

type AncestryMap map[int32]Ancestry
type AncestryNameID struct {
	De *string `bson:"de,omitempty" json:"de,omitempty" yaml:"de,omitempty"`
	En *string `bson:"en,omitempty" json:"en,omitempty" yaml:"en,omitempty"`
	Fr *string `bson:"fr,omitempty" json:"fr,omitempty" yaml:"fr,omitempty"`
	Ja *string `bson:"ja,omitempty" json:"ja,omitempty" yaml:"ja,omitempty"`
	Ko *string `bson:"ko,omitempty" json:"ko,omitempty" yaml:"ko,omitempty"`
	Ru *string `bson:"ru,omitempty" json:"ru,omitempty" yaml:"ru,omitempty"`
	Zh *string `bson:"zh,omitempty" json:"zh,omitempty" yaml:"zh,omitempty"`
}
type AncestryDescriptionID struct {
	De *string `bson:"de,omitempty" json:"de,omitempty" yaml:"de,omitempty"`
	En *string `bson:"en,omitempty" json:"en,omitempty" yaml:"en,omitempty"`
	Fr *string `bson:"fr,omitempty" json:"fr,omitempty" yaml:"fr,omitempty"`
	Ja *string `bson:"ja,omitempty" json:"ja,omitempty" yaml:"ja,omitempty"`
	Ko *string `bson:"ko,omitempty" json:"ko,omitempty" yaml:"ko,omitempty"`
	Ru *string `bson:"ru,omitempty" json:"ru,omitempty" yaml:"ru,omitempty"`
	Zh *string `bson:"zh,omitempty" json:"zh,omitempty" yaml:"zh,omitempty"`
}
type Ancestry struct {
	BloodlineID      *int32                 `bson:"bloodlineID,omitempty" json:"bloodlineID,omitempty" yaml:"bloodlineID,omitempty"`
	Charisma         *int32                 `bson:"charisma,omitempty" json:"charisma,omitempty" yaml:"charisma,omitempty"`
	DescriptionID    *AncestryDescriptionID `bson:"descriptionID,omitempty" json:"descriptionID,omitempty" yaml:"descriptionID,omitempty"`
	IconID           *int32                 `bson:"iconID,omitempty" json:"iconID,omitempty" yaml:"iconID,omitempty"`
	Intelligence     *int32                 `bson:"intelligence,omitempty" json:"intelligence,omitempty" yaml:"intelligence,omitempty"`
	Memory           *int32                 `bson:"memory,omitempty" json:"memory,omitempty" yaml:"memory,omitempty"`
	NameID           *AncestryNameID        `bson:"nameID,omitempty" json:"nameID,omitempty" yaml:"nameID,omitempty"`
	Perception       *int32                 `bson:"perception,omitempty" json:"perception,omitempty" yaml:"perception,omitempty"`
	ShortDescription *string                `bson:"shortDescription,omitempty" json:"shortDescription,omitempty" yaml:"shortDescription,omitempty"`
	Willpower        *int32                 `bson:"willpower,omitempty" json:"willpower,omitempty" yaml:"willpower,omitempty"`
}
