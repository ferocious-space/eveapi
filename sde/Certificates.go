// Code generated by YAML2GO. DO NOT EDIT.

package sde

import (
	yamlv3 "gopkg.in/yaml.v3"
	"os"
)

type CertificateMap map[int32]Certificate

func (x *CertificateMap) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return yamlv3.NewDecoder(f).Decode(x)
}
func (x CertificateMap) Get(ID int32) *Certificate {
	if a, ok := x[ID]; ok {
		return &a
	}
	return nil
}

type Certificate struct {
	Description    *string                         `bson:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	GroupID        *int32                          `bson:"groupID,omitempty" json:"groupID,omitempty" storm:"index" yaml:"groupID,omitempty"`
	Name           *string                         `bson:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	RecommendedFor []int32                         `bson:"recommendedFor,omitempty" json:"recommendedFor,omitempty" yaml:"recommendedFor,omitempty"`
	SkillTypes     map[int32]CertificateSkillTypes `bson:"skillTypes,omitempty" json:"skillTypes,omitempty" yaml:"skillTypes,omitempty"`
}
type CertificateSkillTypes struct {
	Advanced *int32 `bson:"advanced,omitempty" json:"advanced,omitempty" yaml:"advanced,omitempty"`
	Basic    *int32 `bson:"basic,omitempty" json:"basic,omitempty" yaml:"basic,omitempty"`
	Elite    *int32 `bson:"elite,omitempty" json:"elite,omitempty" yaml:"elite,omitempty"`
	Improved *int32 `bson:"improved,omitempty" json:"improved,omitempty" yaml:"improved,omitempty"`
	Standard *int32 `bson:"standard,omitempty" json:"standard,omitempty" yaml:"standard,omitempty"`
}
