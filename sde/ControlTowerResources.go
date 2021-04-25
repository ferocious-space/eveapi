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
	Resources []ControlTowerResourceResources `bson:"resources,omitempty" json:"resources,omitempty" yaml:"resources,omitempty"`
}
type ControlTowerResourceResources struct {
	Purpose          *int32   `bson:"purpose,omitempty" json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Quantity         *int32   `bson:"quantity,omitempty" json:"quantity,omitempty" yaml:"quantity,omitempty"`
	ResourceTypeID   *int32   `bson:"resourceTypeID,omitempty" json:"resourceTypeID,omitempty" storm:"index" yaml:"resourceTypeID,omitempty"`
	FactionID        *int32   `bson:"factionID,omitempty" json:"factionID,omitempty" storm:"index" yaml:"factionID,omitempty"`
	MinSecurityLevel *float64 `bson:"minSecurityLevel,omitempty" json:"minSecurityLevel,omitempty" yaml:"minSecurityLevel,omitempty"`
}
