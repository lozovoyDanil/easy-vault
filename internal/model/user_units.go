package model

import "time"

type UserUnits struct {
	Id          int
	UserId      int
	UnitId      int
	TimeCreated time.Time
}
