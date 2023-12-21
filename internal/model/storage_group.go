package model

import "github.com/uptrace/bun"

type StorageGroup struct {
	bun.BaseModel `bun:"table:Group,alias:g"`

	Id        int    `bun:"id,pk,autoincrement"`
	SpaceId   int    `bun:"space_id"`
	Name      string `json:"name"`
	Size      int    `json:"size"`
	NumOfFree int    `json:"numOfFree" bun:"numOfFree"`
	Price     int    `json:"price" bun:"price"`
	PricePer  int    `json:"pricePer" bun:"pricePer"`
}

type UpdateGroupInput struct {
	Name     *string `json:"name"`
	Price    *int    `json:"price"`
	PricePer *int    `json:"pricePer"`
}
