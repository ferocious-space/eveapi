// Code generated by YAML2GO. DO NOT EDIT.

package sde

type TypeDogmaMap map[int32]TypeDogma
type TypeDogma struct {
	DogmaAttributes []TypeDogmaDogmaAttributes `bson:"dogmaAttributes,omitempty" json:"dogmaAttributes,omitempty" yaml:"dogmaAttributes,omitempty"`
	DogmaEffects    []TypeDogmaDogmaEffects    `bson:"dogmaEffects,omitempty" json:"dogmaEffects,omitempty" yaml:"dogmaEffects,omitempty"`
}
type TypeDogmaDogmaAttributes struct {
	AttributeID *int32   `bson:"attributeID,omitempty" json:"attributeID,omitempty" storm:"index" yaml:"attributeID,omitempty"`
	Value       *float64 `bson:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}
type TypeDogmaDogmaEffects struct {
	EffectID  *int32 `bson:"effectID,omitempty" json:"effectID,omitempty" storm:"index" yaml:"effectID,omitempty"`
	IsDefault *bool  `bson:"isDefault,omitempty" json:"isDefault,omitempty" yaml:"isDefault,omitempty"`
}
