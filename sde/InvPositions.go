// Code generated by YAML2GO. DO NOT EDIT.

package sde

import (
	yamlv3 "gopkg.in/yaml.v3"
	"os"
)

type InvPositionList []InvPosition

func (x *InvPositionList) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return yamlv3.NewDecoder(f).Decode(x)
}

type InvPosition struct {
	ItemID *int64   `bson:"itemID,omitempty" db:"itemID,omitempty" json:"itemID,omitempty" yaml:"itemID,omitempty"`
	Pitch  *float64 `bson:"pitch,omitempty" db:"pitch,omitempty" json:"pitch,omitempty" yaml:"pitch,omitempty"`
	Roll   *float64 `bson:"roll,omitempty" db:"roll,omitempty" json:"roll,omitempty" yaml:"roll,omitempty"`
	X      *float64 `bson:"x,omitempty" db:"x,omitempty" json:"x,omitempty" yaml:"x,omitempty"`
	Y      *float64 `bson:"y,omitempty" db:"y,omitempty" json:"y,omitempty" yaml:"y,omitempty"`
	Yaw    *float64 `bson:"yaw,omitempty" db:"yaw,omitempty" json:"yaw,omitempty" yaml:"yaw,omitempty"`
	Z      *float64 `bson:"z,omitempty" db:"z,omitempty" json:"z,omitempty" yaml:"z,omitempty"`
}
