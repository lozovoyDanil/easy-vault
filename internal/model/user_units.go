package model

import "time"

type UserUnits struct {
	Id           int
	UserId       int
	UnitId       int
	OccupiedFrom time.Time
	BusyUntil    time.Time
}
