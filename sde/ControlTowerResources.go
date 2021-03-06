// Code generated by YAML2GO. DO NOT EDIT.

package sde

import (
	yamlv3 "gopkg.in/yaml.v3"
	"os"
)

type ControlTowerResourceMap map[int32]ControlTowerResource

func (x *ControlTowerResourceMap) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return yamlv3.NewDecoder(f).Decode(x)
}
func (x ControlTowerResourceMap) Get(ID int32) *ControlTowerResource {
	if a, ok := x[ID]; ok {
		return &a
	}
	return nil
}

type ControlTowerResource struct {
	Resources []ControlTowerResourceResources `bson:"resources,omitempty" db:"resources,omitempty" json:"resources,omitempty" yaml:"resources,omitempty"`
}
type ControlTowerResourceResources struct {
	Purpose          *int32   `bson:"purpose,omitempty" db:"purpose,omitempty" json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Quantity         *int32   `bson:"quantity,omitempty" db:"quantity,omitempty" json:"quantity,omitempty" yaml:"quantity,omitempty"`
	ResourceTypeID   *int32   `bson:"resourceTypeID,omitempty" db:"resourceTypeID,omitempty" json:"resourceTypeID,omitempty" yaml:"resourceTypeID,omitempty"`
	FactionID        *int32   `bson:"factionID,omitempty" db:"factionID,omitempty" json:"factionID,omitempty" yaml:"factionID,omitempty"`
	MinSecurityLevel *float64 `bson:"minSecurityLevel,omitempty" db:"minSecurityLevel,omitempty" json:"minSecurityLevel,omitempty" yaml:"minSecurityLevel,omitempty"`
}
