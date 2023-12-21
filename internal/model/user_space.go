package model

import "github.com/uptrace/bun"

type UserSpace struct {
	bun.BaseModel `bun:"table:User_Spaces,alias:us"`

	Id      int `bun:"id,pk,autoincrement"`
	UserId  int `bun:"user_id"`
	SpaceId int `bun:"space_id"`
}
