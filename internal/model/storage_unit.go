package model

import "time"

type StorageUnit struct {
	Id         int
	Name       string
	IsOccupied bool
	LastUsed   time.Time
	BusyUntil  time.Time
}
