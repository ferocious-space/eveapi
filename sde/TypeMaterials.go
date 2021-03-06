// Code generated by YAML2GO. DO NOT EDIT.

package sde

import (
	yamlv3 "gopkg.in/yaml.v3"
	"os"
)

type TypeMaterialMap map[int32]TypeMaterial

func (x *TypeMaterialMap) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return yamlv3.NewDecoder(f).Decode(x)
}
func (x TypeMaterialMap) Get(ID int32) *TypeMaterial {
	if a, ok := x[ID]; ok {
		return &a
	}
	return nil
}

type TypeMaterial struct {
	Materials []TypeMaterialMaterials `bson:"materials,omitempty" db:"materials,omitempty" json:"materials,omitempty" yaml:"materials,omitempty"`
}
type TypeMaterialMaterials struct {
	MaterialTypeID *int32 `bson:"materialTypeID,omitempty" db:"materialTypeID,omitempty" json:"materialTypeID,omitempty" yaml:"materialTypeID,omitempty"`
	Quantity       *int32 `bson:"quantity,omitempty" db:"quantity,omitempty" json:"quantity,omitempty" yaml:"quantity,omitempty"`
}
