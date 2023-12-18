package model

import (
	"time"

	"github.com/uptrace/bun"
)

type StorageUnit struct {
	bun.BaseModel `bun:"table:Unit,alias:u"`

	Id         int       `json:"id" bun:"id,pk,autoincrement"`
	GroupId    int       `bun:"group_id"`
	UserId     int       `bun:"user_id"`
	Name       string    `json:"name" binding:"required"`
	IsOccupied bool      `json:"isOccupied" bun:"isOccupied"`
	LastUsed   time.Time `json:"lastUsed" bun:"lastUsed"`
	BusyUntil  time.Time `json:"busyUntil" bun:"busyUntil"`
}

type UpdateUnitInput struct {
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
