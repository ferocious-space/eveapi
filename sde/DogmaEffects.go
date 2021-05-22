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
	DisallowAutoRepeat             *bool                     `bson:"disallowAutoRepeat,omitempty" db:"disallowAutoRepeat,omitempty" json:"disallowAutoRepeat,omitempty" yaml:"disallowAutoRepeat,omitempty"`
	DischargeAttributeID           *int32                    `bson:"dischargeAttributeID,omitempty" db:"dischargeAttributeID,omitempty" json:"dischargeAttributeID,omitempty" yaml:"dischargeAttributeID,omitempty"`
	DurationAttributeID            *int32                    `bson:"durationAttributeID,omitempty" db:"durationAttributeID,omitempty" json:"durationAttributeID,omitempty" yaml:"durationAttributeID,omitempty"`
	EffectCategory                 *int32                    `bson:"effectCategory,omitempty" db:"effectCategory,omitempty" json:"effectCategory,omitempty" yaml:"effectCategory,omitempty"`
	EffectID                       *int32                    `bson:"effectID,omitempty" db:"effectID,omitempty" json:"effectID,omitempty" yaml:"effectID,omitempty"`
	EffectName                     *string                   `bson:"effectName,omitempty" db:"effectName,omitempty" json:"effectName,omitempty" yaml:"effectName,omitempty"`
	ElectronicChance               *bool                     `bson:"electronicChance,omitempty" db:"electronicChance,omitempty" json:"electronicChance,omitempty" yaml:"electronicChance,omitempty"`
	Guid                           *string                   `bson:"guid,omitempty" db:"guid,omitempty" json:"guid,omitempty" yaml:"guid,omitempty"`
	IsAssistance                   *bool                     `bson:"isAssistance,omitempty" db:"isAssistance,omitempty" json:"isAssistance,omitempty" yaml:"isAssistance,omitempty"`
	IsOffensive                    *bool                     `bson:"isOffensive,omitempty" db:"isOffensive,omitempty" json:"isOffensive,omitempty" yaml:"isOffensive,omitempty"`
	IsWarpSafe                     *bool                     `bson:"isWarpSafe,omitempty" db:"isWarpSafe,omitempty" json:"isWarpSafe,omitempty" yaml:"isWarpSafe,omitempty"`
	PropulsionChance               *bool                     `bson:"propulsionChance,omitempty" db:"propulsionChance,omitempty" json:"propulsionChance,omitempty" yaml:"propulsionChance,omitempty"`
	Published                      *bool                     `bson:"published,omitempty" db:"published,omitempty" json:"published,omitempty" yaml:"published,omitempty"`
	RangeChance                    *bool                     `bson:"rangeChance,omitempty" db:"rangeChance,omitempty" json:"rangeChance,omitempty" yaml:"rangeChance,omitempty"`
	Distribution                   *int32                    `bson:"distribution,omitempty" db:"distribution,omitempty" json:"distribution,omitempty" yaml:"distribution,omitempty"`
	FalloffAttributeID             *int32                    `bson:"falloffAttributeID,omitempty" db:"falloffAttributeID,omitempty" json:"falloffAttributeID,omitempty" yaml:"falloffAttributeID,omitempty"`
	RangeAttributeID               *int32                    `bson:"rangeAttributeID,omitempty" db:"rangeAttributeID,omitempty" json:"rangeAttributeID,omitempty" yaml:"rangeAttributeID,omitempty"`
	TrackingSpeedAttributeID       *int32                    `bson:"trackingSpeedAttributeID,omitempty" db:"trackingSpeedAttributeID,omitempty" json:"trackingSpeedAttributeID,omitempty" yaml:"trackingSpeedAttributeID,omitempty"`
	DescriptionID                  *DogmaEffectDescriptionID `bson:"descriptionID,omitempty" db:"descriptionID,omitempty" json:"descriptionID,omitempty" yaml:"descriptionID,omitempty"`
	DisplayNameID                  *DogmaEffectDisplayNameID `bson:"displayNameID,omitempty" db:"displayNameID,omitempty" json:"displayNameID,omitempty" yaml:"displayNameID,omitempty"`
	IconID                         *int32                    `bson:"iconID,omitempty" db:"iconID,omitempty" json:"iconID,omitempty" yaml:"iconID,omitempty"`
	ModifierInfo                   []DogmaEffectModifierInfo `bson:"modifierInfo,omitempty" db:"modifierInfo,omitempty" json:"modifierInfo,omitempty" yaml:"modifierInfo,omitempty"`
	SfxName                        *string                   `bson:"sfxName,omitempty" db:"sfxName,omitempty" json:"sfxName,omitempty" yaml:"sfxName,omitempty"`
	NpcUsageChanceAttributeID      *int32                    `bson:"npcUsageChanceAttributeID,omitempty" db:"npcUsageChanceAttributeID,omitempty" json:"npcUsageChanceAttributeID,omitempty" yaml:"npcUsageChanceAttributeID,omitempty"`
	NpcActivationChanceAttributeID *int32                    `bson:"npcActivationChanceAttributeID,omitempty" db:"npcActivationChanceAttributeID,omitempty" json:"npcActivationChanceAttributeID,omitempty" yaml:"npcActivationChanceAttributeID,omitempty"`
	FittingUsageChanceAttributeID  *int32                    `bson:"fittingUsageChanceAttributeID,omitempty" db:"fittingUsageChanceAttributeID,omitempty" json:"fittingUsageChanceAttributeID,omitempty" yaml:"fittingUsageChanceAttributeID,omitempty"`
	ResistanceAttributeID          *int32                    `bson:"resistanceAttributeID,omitempty" db:"resistanceAttributeID,omitempty" json:"resistanceAttributeID,omitempty" yaml:"resistanceAttributeID,omitempty"`
}
type DogmaEffectDescriptionID struct {
	De *string `bson:"de,omitempty" db:"de,omitempty" json:"de,omitempty" yaml:"de,omitempty"`
	En *string `bson:"en,omitempty" db:"en,omitempty" json:"en,omitempty" yaml:"en,omitempty"`
	Fr *string `bson:"fr,omitempty" db:"fr,omitempty" json:"fr,omitempty" yaml:"fr,omitempty"`
	Ja *string `bson:"ja,omitempty" db:"ja,omitempty" json:"ja,omitempty" yaml:"ja,omitempty"`
	Ko *string `bson:"ko,omitempty" db:"ko,omitempty" json:"ko,omitempty" yaml:"ko,omitempty"`
	Ru *string `bson:"ru,omitempty" db:"ru,omitempty" json:"ru,omitempty" yaml:"ru,omitempty"`
	Zh *string `bson:"zh,omitempty" db:"zh,omitempty" json:"zh,omitempty" yaml:"zh,omitempty"`
}
type DogmaEffectDisplayNameID struct {
	De *string `bson:"de,omitempty" db:"de,omitempty" json:"de,omitempty" yaml:"de,omitempty"`
	En *string `bson:"en,omitempty" db:"en,omitempty" json:"en,omitempty" yaml:"en,omitempty"`
	Fr *string `bson:"fr,omitempty" db:"fr,omitempty" json:"fr,omitempty" yaml:"fr,omitempty"`
	Ja *string `bson:"ja,omitempty" db:"ja,omitempty" json:"ja,omitempty" yaml:"ja,omitempty"`
	Ko *string `bson:"ko,omitempty" db:"ko,omitempty" json:"ko,omitempty" yaml:"ko,omitempty"`
	Ru *string `bson:"ru,omitempty" db:"ru,omitempty" json:"ru,omitempty" yaml:"ru,omitempty"`
	Zh *string `bson:"zh,omitempty" db:"zh,omitempty" json:"zh,omitempty" yaml:"zh,omitempty"`
}
type DogmaEffectModifierInfo struct {
	Domain               *string `bson:"domain,omitempty" db:"domain,omitempty" json:"domain,omitempty" yaml:"domain,omitempty"`
	Func                 *string `bson:"func,omitempty" db:"func,omitempty" json:"func,omitempty" yaml:"func,omitempty"`
	ModifiedAttributeID  *int32  `bson:"modifiedAttributeID,omitempty" db:"modifiedAttributeID,omitempty" json:"modifiedAttributeID,omitempty" yaml:"modifiedAttributeID,omitempty"`
	ModifyingAttributeID *int32  `bson:"modifyingAttributeID,omitempty" db:"modifyingAttributeID,omitempty" json:"modifyingAttributeID,omitempty" yaml:"modifyingAttributeID,omitempty"`
	Operation            *int32  `bson:"operation,omitempty" db:"operation,omitempty" json:"operation,omitempty" yaml:"operation,omitempty"`
	GroupID              *int32  `bson:"groupID,omitempty" db:"groupID,omitempty" json:"groupID,omitempty" yaml:"groupID,omitempty"`
	SkillTypeID          *int32  `bson:"skillTypeID,omitempty" db:"skillTypeID,omitempty" json:"skillTypeID,omitempty" yaml:"skillTypeID,omitempty"`
	EffectID             *int32  `bson:"effectID,omitempty" db:"effectID,omitempty" json:"effectID,omitempty" yaml:"effectID,omitempty"`
}
