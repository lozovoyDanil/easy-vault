package model

import "time"

type StorageUnit struct {
	Id         int
	Name       string    `json:"name" binding:"required"`
	IsOccupied bool      `json:"isOccupied"`
	LastUsed   time.Time `json:"lastUsed"`
	BusyUntil  time.Time `json:"busyUntil"`
}

type UpdateUnitInput struct {
	Name       *string    `json:"name"`
	IsOccupied *bool      `json:"isOccupied"`
	LastUsed   *time.Time `json:"lastUsed"`
	BusyUntil  *time.Time `json:"busyUntil"`
}
