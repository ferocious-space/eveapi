// Code generated by YAML2GO. DO NOT EDIT.

package sde

import (
	yamlv3 "gopkg.in/yaml.v3"
	"os"
)

type InvFlagList []InvFlag

func (x *InvFlagList) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return yamlv3.NewDecoder(f).Decode(x)
}

type InvFlag struct {
	FlagID   *int32  `bson:"flagID,omitempty" db:"flagID,omitempty" json:"flagID,omitempty" yaml:"flagID,omitempty"`
	FlagName *string `bson:"flagName,omitempty" db:"flagName,omitempty" json:"flagName,omitempty" yaml:"flagName,omitempty"`
	FlagText *string `bson:"flagText,omitempty" db:"flagText,omitempty" json:"flagText,omitempty" yaml:"flagText,omitempty"`
	OrderID  *int32  `bson:"orderID,omitempty" db:"orderID,omitempty" json:"orderID,omitempty" yaml:"orderID,omitempty"`
}
