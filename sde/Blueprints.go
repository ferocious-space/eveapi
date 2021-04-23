// Code generated by YAML2GO. DO NOT EDIT.

package sde

type BlueprintMap map[int32]Blueprint
type BlueprintActivitiesReactionProducts struct {
	Quantity *int32 `bson:"quantity,omitempty" json:"quantity,omitempty" yaml:"quantity,omitempty"`
	TypeID   *int32 `bson:"typeID,omitempty" json:"typeID,omitempty" yaml:"typeID,omitempty"`
}
type BlueprintActivitiesResearchMaterialSkills struct {
	Level  *int32 `bson:"level,omitempty" json:"level,omitempty" yaml:"level,omitempty"`
	TypeID *int32 `bson:"typeID,omitempty" json:"typeID,omitempty" yaml:"typeID,omitempty"`
}
type BlueprintActivitiesResearchTimeProducts struct{}
type BlueprintActivitiesResearchTimeMaterials struct {
	Quantity *int32 `bson:"quantity,omitempty" json:"quantity,omitempty" yaml:"quantity,omitempty"`
	TypeID   *int32 `bson:"typeID,omitempty" json:"typeID,omitempty" yaml:"typeID,omitempty"`
}
type BlueprintActivitiesInventionMaterials struct {
	Quantity *int32 `bson:"quantity,omitempty" json:"quantity,omitempty" yaml:"quantity,omitempty"`
	TypeID   *int32 `bson:"typeID,omitempty" json:"typeID,omitempty" yaml:"typeID,omitempty"`
}
type BlueprintActivitiesResearchTime struct {
	Time      *int64                                     `bson:"time,omitempty" json:"time,omitempty" yaml:"time,omitempty"`
	Materials []BlueprintActivitiesResearchTimeMaterials `bson:"materials,omitempty" json:"materials,omitempty" yaml:"materials,omitempty"`
	Skills    []BlueprintActivitiesResearchTimeSkills    `bson:"skills,omitempty" json:"skills,omitempty" yaml:"skills,omitempty"`
	Products  []BlueprintActivitiesResearchTimeProducts  `bson:"products,omitempty" json:"products,omitempty" yaml:"products,omitempty"`
}
type BlueprintActivitiesCopyingProducts struct{}
type BlueprintActivitiesManufacturingProducts struct {
	Quantity *int32 `bson:"quantity,omitempty" json:"quantity,omitempty" yaml:"quantity,omitempty"`
	TypeID   *int32 `bson:"typeID,omitempty" json:"typeID,omitempty" yaml:"typeID,omitempty"`
}
type BlueprintActivitiesInventionSkills struct {
	Level  *int32 `bson:"level,omitempty" json:"level,omitempty" yaml:"level,omitempty"`
	TypeID *int32 `bson:"typeID,omitempty" json:"typeID,omitempty" yaml:"typeID,omitempty"`
}
type BlueprintActivitiesCopyingSkills struct {
	Level  *int32 `bson:"level,omitempty" json:"level,omitempty" yaml:"level,omitempty"`
	TypeID *int32 `bson:"typeID,omitempty" json:"typeID,omitempty" yaml:"typeID,omitempty"`
}
type BlueprintActivitiesResearchMaterial struct {
	Time      *int64                                         `bson:"time,omitempty" json:"time,omitempty" yaml:"time,omitempty"`
	Materials []BlueprintActivitiesResearchMaterialMaterials `bson:"materials,omitempty" json:"materials,omitempty" yaml:"materials,omitempty"`
	Skills    []BlueprintActivitiesResearchMaterialSkills    `bson:"skills,omitempty" json:"skills,omitempty" yaml:"skills,omitempty"`
	Products  []BlueprintActivitiesResearchMaterialProducts  `bson:"products,omitempty" json:"products,omitempty" yaml:"products,omitempty"`
}
type BlueprintActivitiesReactionMaterials struct {
	Quantity *int32 `bson:"quantity,omitempty" json:"quantity,omitempty" yaml:"quantity,omitempty"`
	TypeID   *int32 `bson:"typeID,omitempty" json:"typeID,omitempty" yaml:"typeID,omitempty"`
}
type BlueprintActivitiesReaction struct {
	Materials []BlueprintActivitiesReactionMaterials `bson:"materials,omitempty" json:"materials,omitempty" yaml:"materials,omitempty"`
	Products  []BlueprintActivitiesReactionProducts  `bson:"products,omitempty" json:"products,omitempty" yaml:"products,omitempty"`
	Skills    []BlueprintActivitiesReactionSkills    `bson:"skills,omitempty" json:"skills,omitempty" yaml:"skills,omitempty"`
	Time      *int64                                 `bson:"time,omitempty" json:"time,omitempty" yaml:"time,omitempty"`
}
type BlueprintActivitiesResearchTimeSkills struct {
	Level  *int32 `bson:"level,omitempty" json:"level,omitempty" yaml:"level,omitempty"`
	TypeID *int32 `bson:"typeID,omitempty" json:"typeID,omitempty" yaml:"typeID,omitempty"`
}
type Blueprint struct {
	Activities         *BlueprintActivities `bson:"activities,omitempty" json:"activities,omitempty" yaml:"activities,omitempty"`
	BlueprintTypeID    *int32               `bson:"blueprintTypeID,omitempty" json:"blueprintTypeID,omitempty" yaml:"blueprintTypeID,omitempty"`
	MaxProductionLimit *int32               `bson:"maxProductionLimit,omitempty" json:"maxProductionLimit,omitempty" yaml:"maxProductionLimit,omitempty"`
}
type BlueprintActivitiesResearchMaterialProducts struct{}
type BlueprintActivitiesInventionProducts struct {
	Probability *float64 `bson:"probability,omitempty" json:"probability,omitempty" yaml:"probability,omitempty"`
	Quantity    *int32   `bson:"quantity,omitempty" json:"quantity,omitempty" yaml:"quantity,omitempty"`
	TypeID      *int32   `bson:"typeID,omitempty" json:"typeID,omitempty" yaml:"typeID,omitempty"`
}
type BlueprintActivitiesCopying struct {
	Time      *int64                                `bson:"time,omitempty" json:"time,omitempty" yaml:"time,omitempty"`
	Materials []BlueprintActivitiesCopyingMaterials `bson:"materials,omitempty" json:"materials,omitempty" yaml:"materials,omitempty"`
	Skills    []BlueprintActivitiesCopyingSkills    `bson:"skills,omitempty" json:"skills,omitempty" yaml:"skills,omitempty"`
	Products  []BlueprintActivitiesCopyingProducts  `bson:"products,omitempty" json:"products,omitempty" yaml:"products,omitempty"`
}
type BlueprintActivitiesCopyingMaterials struct {
	Quantity *int32 `bson:"quantity,omitempty" json:"quantity,omitempty" yaml:"quantity,omitempty"`
	TypeID   *int32 `bson:"typeID,omitempty" json:"typeID,omitempty" yaml:"typeID,omitempty"`
}
type BlueprintActivitiesResearchMaterialMaterials struct {
	Quantity *int32 `bson:"quantity,omitempty" json:"quantity,omitempty" yaml:"quantity,omitempty"`
	TypeID   *int32 `bson:"typeID,omitempty" json:"typeID,omitempty" yaml:"typeID,omitempty"`
}
type BlueprintActivitiesManufacturing struct {
	Materials []BlueprintActivitiesManufacturingMaterials `bson:"materials,omitempty" json:"materials,omitempty" yaml:"materials,omitempty"`
	Products  []BlueprintActivitiesManufacturingProducts  `bson:"products,omitempty" json:"products,omitempty" yaml:"products,omitempty"`
	Time      *int64                                      `bson:"time,omitempty" json:"time,omitempty" yaml:"time,omitempty"`
	Skills    []BlueprintActivitiesManufacturingSkills    `bson:"skills,omitempty" json:"skills,omitempty" yaml:"skills,omitempty"`
}
type BlueprintActivitiesManufacturingMaterials struct {
	Quantity *int32 `bson:"quantity,omitempty" json:"quantity,omitempty" yaml:"quantity,omitempty"`
	TypeID   *int32 `bson:"typeID,omitempty" json:"typeID,omitempty" yaml:"typeID,omitempty"`
}
type BlueprintActivitiesInvention struct {
	Materials []BlueprintActivitiesInventionMaterials `bson:"materials,omitempty" json:"materials,omitempty" yaml:"materials,omitempty"`
	Products  []BlueprintActivitiesInventionProducts  `bson:"products,omitempty" json:"products,omitempty" yaml:"products,omitempty"`
	Skills    []BlueprintActivitiesInventionSkills    `bson:"skills,omitempty" json:"skills,omitempty" yaml:"skills,omitempty"`
	Time      *int64                                  `bson:"time,omitempty" json:"time,omitempty" yaml:"time,omitempty"`
}
type BlueprintActivitiesReactionSkills struct {
	Level  *int32 `bson:"level,omitempty" json:"level,omitempty" yaml:"level,omitempty"`
	TypeID *int32 `bson:"typeID,omitempty" json:"typeID,omitempty" yaml:"typeID,omitempty"`
}
type BlueprintActivities struct {
	Copying          *BlueprintActivitiesCopying          `bson:"copying,omitempty" json:"copying,omitempty" yaml:"copying,omitempty"`
	Manufacturing    *BlueprintActivitiesManufacturing    `bson:"manufacturing,omitempty" json:"manufacturing,omitempty" yaml:"manufacturing,omitempty"`
	ResearchMaterial *BlueprintActivitiesResearchMaterial `bson:"research_material,omitempty" json:"research_material,omitempty" yaml:"research_material,omitempty"`
	ResearchTime     *BlueprintActivitiesResearchTime     `bson:"research_time,omitempty" json:"research_time,omitempty" yaml:"research_time,omitempty"`
	Invention        *BlueprintActivitiesInvention        `bson:"invention,omitempty" json:"invention,omitempty" yaml:"invention,omitempty"`
	Reaction         *BlueprintActivitiesReaction         `bson:"reaction,omitempty" json:"reaction,omitempty" yaml:"reaction,omitempty"`
}
type BlueprintActivitiesManufacturingSkills struct {
	Level  *int32 `bson:"level,omitempty" json:"level,omitempty" yaml:"level,omitempty"`
	TypeID *int32 `bson:"typeID,omitempty" json:"typeID,omitempty" yaml:"typeID,omitempty"`
}
