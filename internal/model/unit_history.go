package model

import "time"

type UnitHistory struct {
	Id         int
	UnitId     int
	UserId     int
	UserInfo   string
	Action     string
	ActionDate time.Time
}
