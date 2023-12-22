package model

import "github.com/uptrace/bun"

type StorageGroup struct {
	bun.BaseModel `bun:"table:StorageGroup,alias:g"`

	Id        int    `bun:"id,pk,autoincrement"`
	SpaceId   int    `bun:"space_id"`
	Name      string `json:"name"`
	Size      int    `json:"size"`
	NumOfFree int    `json:"numOfFree" bun:"numOfFree"`
	Price     int    `json:"price" bun:"price"`
	PricePer  string `json:"pricePer" bun:"pricePer"`
}

type GroupInput struct {
	Name     *string `json:"name"`
	Price    *int    `json:"price"`
	PricePer *int    `json:"pricePer"`
}
