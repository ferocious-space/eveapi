// Code generated by YAML2GO. DO NOT EDIT.

package sde

import (
	yamlv3 "gopkg.in/yaml.v3"
	"os"
)

type DogmaEffectMap map[int32]DogmaEffect

func (x *DogmaEffectMap) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return yamlv3.NewDecoder(f).Decode(x)
}
func (x DogmaEffectMap) Get(ID int32) *DogmaEffect {
	if a, ok := x[ID]; ok {
		return &a
	}
	return nil
}

type DogmaEffect struct {
	DisallowAutoRepeat             *bool                     `bson:"disallowAutoRepeat,omitempty" json:"disallowAutoRepeat,omitempty" yaml:"disallowAutoRepeat,omitempty"`
	DischargeAttributeID           *int32                    `bson:"dischargeAttributeID,omitempty" json:"dischargeAttributeID,omitempty" storm:"index" yaml:"dischargeAttributeID,omitempty"`
	DurationAttributeID            *int32                    `bson:"durationAttributeID,omitempty" json:"durationAttributeID,omitempty" storm:"index" yaml:"durationAttributeID,omitempty"`
	EffectCategory                 *int32                    `bson:"effectCategory,omitempty" json:"effectCategory,omitempty" yaml:"effectCategory,omitempty"`
	EffectID                       *int32                    `bson:"effectID,omitempty" json:"effectID,omitempty" storm:"index" yaml:"effectID,omitempty"`
	EffectName                     *string                   `bson:"effectName,omitempty" json:"effectName,omitempty" yaml:"effectName,omitempty"`
	ElectronicChance               *bool                     `bson:"electronicChance,omitempty" json:"electronicChance,omitempty" yaml:"electronicChance,omitempty"`
	Guid                           *string                   `bson:"guid,omitempty" json:"guid,omitempty" yaml:"guid,omitempty"`
	IsAssistance                   *bool                     `bson:"isAssistance,omitempty" json:"isAssistance,omitempty" yaml:"isAssistance,omitempty"`
	IsOffensive                    *bool                     `bson:"isOffensive,omitempty" json:"isOffensive,omitempty" yaml:"isOffensive,omitempty"`
	IsWarpSafe                     *bool                     `bson:"isWarpSafe,omitempty" json:"isWarpSafe,omitempty" yaml:"isWarpSafe,omitempty"`
	PropulsionChance               *bool                     `bson:"propulsionChance,omitempty" json:"propulsionChance,omitempty" yaml:"propulsionChance,omitempty"`
	Published                      *bool                     `bson:"published,omitempty" json:"published,omitempty" yaml:"published,omitempty"`
	RangeChance                    *bool                     `bson:"rangeChance,omitempty" json:"rangeChance,omitempty" yaml:"rangeChance,omitempty"`
	Distribution                   *int32                    `bson:"distribution,omitempty" json:"distribution,omitempty" yaml:"distribution,omitempty"`
	FalloffAttributeID             *int32                    `bson:"falloffAttributeID,omitempty" json:"falloffAttributeID,omitempty" storm:"index" yaml:"falloffAttributeID,omitempty"`
	RangeAttributeID               *int32                    `bson:"rangeAttributeID,omitempty" json:"rangeAttributeID,omitempty" storm:"index" yaml:"rangeAttributeID,omitempty"`
	TrackingSpeedAttributeID       *int32                    `bson:"trackingSpeedAttributeID,omitempty" json:"trackingSpeedAttributeID,omitempty" storm:"index" yaml:"trackingSpeedAttributeID,omitempty"`
	DescriptionID                  *DogmaEffectDescriptionID `bson:"descriptionID,omitempty" json:"descriptionID,omitempty" storm:"index" yaml:"descriptionID,omitempty"`
	DisplayNameID                  *DogmaEffectDisplayNameID `bson:"displayNameID,omitempty" json:"displayNameID,omitempty" storm:"index" yaml:"displayNameID,omitempty"`
	IconID                         *int32                    `bson:"iconID,omitempty" json:"iconID,omitempty" storm:"index" yaml:"iconID,omitempty"`
	ModifierInfo                   []DogmaEffectModifierInfo `bson:"modifierInfo,omitempty" json:"modifierInfo,omitempty" yaml:"modifierInfo,omitempty"`
	SfxName                        *string                   `bson:"sfxName,omitempty" json:"sfxName,omitempty" yaml:"sfxName,omitempty"`
	NpcUsageChanceAttributeID      *int32                    `bson:"npcUsageChanceAttributeID,omitempty" json:"npcUsageChanceAttributeID,omitempty" storm:"index" yaml:"npcUsageChanceAttributeID,omitempty"`
	NpcActivationChanceAttributeID *int32                    `bson:"npcActivationChanceAttributeID,omitempty" json:"npcActivationChanceAttributeID,omitempty" storm:"index" yaml:"npcActivationChanceAttributeID,omitempty"`
	FittingUsageChanceAttributeID  *int32                    `bson:"fittingUsageChanceAttributeID,omitempty" json:"fittingUsageChanceAttributeID,omitempty" storm:"index" yaml:"fittingUsageChanceAttributeID,omitempty"`
	ResistanceAttributeID          *int32                    `bson:"resistanceAttributeID,omitempty" json:"resistanceAttributeID,omitempty" storm:"index" yaml:"resistanceAttributeID,omitempty"`
}
type DogmaEffectDescriptionID struct {
	De *string `bson:"de,omitempty" json:"de,omitempty" yaml:"de,omitempty"`
	En *string `bson:"en,omitempty" json:"en,omitempty" yaml:"en,omitempty"`
	Fr *string `bson:"fr,omitempty" json:"fr,omitempty" yaml:"fr,omitempty"`
	Ja *string `bson:"ja,omitempty" json:"ja,omitempty" yaml:"ja,omitempty"`
	Ko *string `bson:"ko,omitempty" json:"ko,omitempty" yaml:"ko,omitempty"`
	Ru *string `bson:"ru,omitempty" json:"ru,omitempty" yaml:"ru,omitempty"`
	Zh *string `bson:"zh,omitempty" json:"zh,omitempty" yaml:"zh,omitempty"`
}
type DogmaEffectDisplayNameID struct {
	De *string `bson:"de,omitempty" json:"de,omitempty" yaml:"de,omitempty"`
	En *string `bson:"en,omitempty" json:"en,omitempty" yaml:"en,omitempty"`
	Fr *string `bson:"fr,omitempty" json:"fr,omitempty" yaml:"fr,omitempty"`
	Ja *string `bson:"ja,omitempty" json:"ja,omitempty" yaml:"ja,omitempty"`
	Ko *string `bson:"ko,omitempty" json:"ko,omitempty" yaml:"ko,omitempty"`
	Ru *string `bson:"ru,omitempty" json:"ru,omitempty" yaml:"ru,omitempty"`
	Zh *string `bson:"zh,omitempty" json:"zh,omitempty" yaml:"zh,omitempty"`
}
type DogmaEffectModifierInfo struct {
	Domain               *string `bson:"domain,omitempty" json:"domain,omitempty" yaml:"domain,omitempty"`
	Func                 *string `bson:"func,omitempty" json:"func,omitempty" yaml:"func,omitempty"`
	ModifiedAttributeID  *int32  `bson:"modifiedAttributeID,omitempty" json:"modifiedAttributeID,omitempty" storm:"index" yaml:"modifiedAttributeID,omitempty"`
	ModifyingAttributeID *int32  `bson:"modifyingAttributeID,omitempty" json:"modifyingAttributeID,omitempty" storm:"index" yaml:"modifyingAttributeID,omitempty"`
	Operation            *int32  `bson:"operation,omitempty" json:"operation,omitempty" yaml:"operation,omitempty"`
	GroupID              *int32  `bson:"groupID,omitempty" json:"groupID,omitempty" storm:"index" yaml:"groupID,omitempty"`
	SkillTypeID          *int32  `bson:"skillTypeID,omitempty" json:"skillTypeID,omitempty" storm:"index" yaml:"skillTypeID,omitempty"`
	EffectID             *int32  `bson:"effectID,omitempty" json:"effectID,omitempty" storm:"index" yaml:"effectID,omitempty"`
}
