// Code generated by YAML2GO. DO NOT EDIT.

package sde

type GraphicIDMap map[int32]GraphicID
type GraphicID struct {
	Description    *string            `bson:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	GraphicFile    *string            `bson:"graphicFile,omitempty" json:"graphicFile,omitempty" yaml:"graphicFile,omitempty"`
	IconInfo       *GraphicIDIconInfo `bson:"iconInfo,omitempty" json:"iconInfo,omitempty" yaml:"iconInfo,omitempty"`
	SofFactionName *string            `bson:"sofFactionName,omitempty" json:"sofFactionName,omitempty" yaml:"sofFactionName,omitempty"`
	SofHullName    *string            `bson:"sofHullName,omitempty" json:"sofHullName,omitempty" yaml:"sofHullName,omitempty"`
	SofRaceName    *string            `bson:"sofRaceName,omitempty" json:"sofRaceName,omitempty" yaml:"sofRaceName,omitempty"`
}
type GraphicIDIconInfo struct {
	Folder *string `bson:"folder,omitempty" json:"folder,omitempty" yaml:"folder,omitempty"`
}
