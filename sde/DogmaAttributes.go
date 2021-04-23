// Code generated by YAML2GO. DO NOT EDIT.

package sde

type DogmaAttributeMap map[int32]DogmaAttribute
type DogmaAttribute struct {
	AttributeID          *int32                              `bson:"attributeID,omitempty" json:"attributeID,omitempty" yaml:"attributeID,omitempty"`
	CategoryID           *int32                              `bson:"categoryID,omitempty" json:"categoryID,omitempty" yaml:"categoryID,omitempty"`
	DataType             *int32                              `bson:"dataType,omitempty" json:"dataType,omitempty" yaml:"dataType,omitempty"`
	DefaultValue         *float64                            `bson:"defaultValue,omitempty" json:"defaultValue,omitempty" yaml:"defaultValue,omitempty"`
	Description          *string                             `bson:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	HighIsGood           *bool                               `bson:"highIsGood,omitempty" json:"highIsGood,omitempty" yaml:"highIsGood,omitempty"`
	Name                 *string                             `bson:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Published            *bool                               `bson:"published,omitempty" json:"published,omitempty" yaml:"published,omitempty"`
	Stackable            *bool                               `bson:"stackable,omitempty" json:"stackable,omitempty" yaml:"stackable,omitempty"`
	DisplayNameID        *DogmaAttributeDisplayNameID        `bson:"displayNameID,omitempty" json:"displayNameID,omitempty" yaml:"displayNameID,omitempty"`
	IconID               *int32                              `bson:"iconID,omitempty" json:"iconID,omitempty" yaml:"iconID,omitempty"`
	TooltipDescriptionID *DogmaAttributeTooltipDescriptionID `bson:"tooltipDescriptionID,omitempty" json:"tooltipDescriptionID,omitempty" yaml:"tooltipDescriptionID,omitempty"`
	TooltipTitleID       *DogmaAttributeTooltipTitleID       `bson:"tooltipTitleID,omitempty" json:"tooltipTitleID,omitempty" yaml:"tooltipTitleID,omitempty"`
	UnitID               *int32                              `bson:"unitID,omitempty" json:"unitID,omitempty" yaml:"unitID,omitempty"`
	ChargeRechargeTimeID *int64                              `bson:"chargeRechargeTimeID,omitempty" json:"chargeRechargeTimeID,omitempty" yaml:"chargeRechargeTimeID,omitempty"`
	MaxAttributeID       *int32                              `bson:"maxAttributeID,omitempty" json:"maxAttributeID,omitempty" yaml:"maxAttributeID,omitempty"`
}
type DogmaAttributeDisplayNameID struct {
	De *string `bson:"de,omitempty" json:"de,omitempty" yaml:"de,omitempty"`
	En *string `bson:"en,omitempty" json:"en,omitempty" yaml:"en,omitempty"`
	Fr *string `bson:"fr,omitempty" json:"fr,omitempty" yaml:"fr,omitempty"`
	Ja *string `bson:"ja,omitempty" json:"ja,omitempty" yaml:"ja,omitempty"`
	Ko *string `bson:"ko,omitempty" json:"ko,omitempty" yaml:"ko,omitempty"`
	Ru *string `bson:"ru,omitempty" json:"ru,omitempty" yaml:"ru,omitempty"`
	Zh *string `bson:"zh,omitempty" json:"zh,omitempty" yaml:"zh,omitempty"`
}
type DogmaAttributeTooltipDescriptionID struct {
	De *string `bson:"de,omitempty" json:"de,omitempty" yaml:"de,omitempty"`
	En *string `bson:"en,omitempty" json:"en,omitempty" yaml:"en,omitempty"`
	Fr *string `bson:"fr,omitempty" json:"fr,omitempty" yaml:"fr,omitempty"`
	Ja *string `bson:"ja,omitempty" json:"ja,omitempty" yaml:"ja,omitempty"`
	Ko *string `bson:"ko,omitempty" json:"ko,omitempty" yaml:"ko,omitempty"`
	Ru *string `bson:"ru,omitempty" json:"ru,omitempty" yaml:"ru,omitempty"`
	Zh *string `bson:"zh,omitempty" json:"zh,omitempty" yaml:"zh,omitempty"`
}
type DogmaAttributeTooltipTitleID struct {
	De *string `bson:"de,omitempty" json:"de,omitempty" yaml:"de,omitempty"`
	En *string `bson:"en,omitempty" json:"en,omitempty" yaml:"en,omitempty"`
	Fr *string `bson:"fr,omitempty" json:"fr,omitempty" yaml:"fr,omitempty"`
	Ja *string `bson:"ja,omitempty" json:"ja,omitempty" yaml:"ja,omitempty"`
	Ko *string `bson:"ko,omitempty" json:"ko,omitempty" yaml:"ko,omitempty"`
	Ru *string `bson:"ru,omitempty" json:"ru,omitempty" yaml:"ru,omitempty"`
	Zh *string `bson:"zh,omitempty" json:"zh,omitempty" yaml:"zh,omitempty"`
}
