package model

import "github.com/uptrace/bun"

type Space struct {
	bun.BaseModel `bun:"table:Space,alias:s"`

	Id          int    `json:"id" bun:"id,pk,autoincrement"`
	Name        string `json:"name"`
	Addr        string `json:"addr"`
	Size        int    `json:"size"`
	NumOfGroups int    `json:"numOfGroups" bun:"numOfGroups"`
	NumOfFree   int    `json:"numOfFree" bun:"numOfFree"`
}

type SpaceInput struct {
	Name *string `json:"name"`
	Addr *string `json:"addr"`
}

type SpaceFilter struct {
	Name    *string `json:"name"`
	Addr    *string `json:"addr"`
	Order   *string `json:"order"`
	MinSize *int    `json:"minSize"`
	MaxSize *int    `json:"maxSize"`
}
