// Code generated by YAML2GO. DO NOT EDIT.

package sde

import (
	yamlv3 "gopkg.in/yaml.v3"
	"os"
)

type InvItemList []InvItem

func (x *InvItemList) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return yamlv3.NewDecoder(f).Decode(x)
}

type InvItem struct {
	FlagID     *int32 `bson:"flagID,omitempty" json:"flagID,omitempty" storm:"index" yaml:"flagID,omitempty"`
	ItemID     *int64 `bson:"itemID,omitempty" json:"itemID,omitempty" storm:"index" yaml:"itemID,omitempty"`
	LocationID *int64 `bson:"locationID,omitempty" json:"locationID,omitempty" storm:"index" yaml:"locationID,omitempty"`
	OwnerID    *int32 `bson:"ownerID,omitempty" json:"ownerID,omitempty" storm:"index" yaml:"ownerID,omitempty"`
	Quantity   *int32 `bson:"quantity,omitempty" json:"quantity,omitempty" yaml:"quantity,omitempty"`
	TypeID     *int32 `bson:"typeID,omitempty" json:"typeID,omitempty" storm:"index" yaml:"typeID,omitempty"`
}
