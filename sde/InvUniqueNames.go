// Code generated by YAML2GO. DO NOT EDIT.

package sde

import (
	yamlv3 "gopkg.in/yaml.v3"
	"os"
)

type InvUniqueNameList []InvUniqueName

func (x *InvUniqueNameList) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return yamlv3.NewDecoder(f).Decode(x)
}

type InvUniqueName struct {
	GroupID  *int32  `bson:"groupID,omitempty" db:"groupID,omitempty" json:"groupID,omitempty" yaml:"groupID,omitempty"`
	ItemID   *int64  `bson:"itemID,omitempty" db:"itemID,omitempty" json:"itemID,omitempty" yaml:"itemID,omitempty"`
	ItemName *string `bson:"itemName,omitempty" db:"itemName,omitempty" json:"itemName,omitempty" yaml:"itemName,omitempty"`
}
