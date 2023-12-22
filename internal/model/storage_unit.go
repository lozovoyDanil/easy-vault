package model

import (
	"time"

	"github.com/uptrace/bun"
)

type StorageUnit struct {
	bun.BaseModel `bun:"table:Unit,alias:u"`

	Id         int    `json:"id" bun:"id,pk,autoincrement"`
	GroupId    int    `json:"group_id" bun:"group_id"`
	UserId     int    `json:"user_id" bun:"user_id"`
	Name       string `json:"name"`
	IsOccupied bool   `json:"isOccupied" bun:"isOccupied"`
	LastUsed   string `json:"lastUsed" bun:"lastUsed"`
	BusyUntil  string `json:"busyUntil" bun:"busyUntil"`
}

type UnitInput struct {
	UserId     *int       `json:"userId"`
	Name       *string    `json:"name"`
	IsOccupied *bool      `json:"isOccupied"`
	LastUsed   *time.Time `json:"lastUsed"`
	BusyUntil  *time.Time `json:"busyUntil"`
}

type UnitDetails struct {
	StorageUnit
	History []UnitHistory `json:"history"`
}
