package model

import "time"

type UnitHistory struct {
	Id         int       `json:"id" bun:"id,pk,autoincrement"`
	UnitId     int       `bun:"unit_id"`
	UserId     int       `bun:"user_id"`
	UserInfo   string    `json:"userInfo" bun:"userInfo"`
	Action     string    `json:"action" bun:"action"`
	ActionDate time.Time `json:"actionDate" bun:"actionDate"`
}
