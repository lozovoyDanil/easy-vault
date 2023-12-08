package model

import "github.com/uptrace/bun"

type Space struct {
	bun.BaseModel `bun:"table:Space,alias:s"`

	Id        int    `json:"id" bun:"id,pk,autoincrement"`
	Name      string `json:"name"`
	Addr      string `json:"addr"`
	Size      int    `json:"size"`
	NumOfFree int    `json:"numOfFree" bun:"numOfFree"`
}

type UpdateSpaceInput struct {
	Name *string `json:"name"`
	Addr *string `json:"addr"`
}
