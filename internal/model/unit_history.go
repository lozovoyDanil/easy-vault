package model

import (
	"time"

	"github.com/uptrace/bun"
)

type UnitHistory struct {
	bun.BaseModel `bun:"table:Unit_History,alias:uh"`

	Id         int       `json:"id" bun:"id,pk,autoincrement"`
	UnitId     int       `bun:"unit_id"`
	UserId     int       `bun:"user_id"`
	Status     int       `json:"status" bun:"status"`
	Action     string    `json:"action" bun:"action"`
	ActionDate time.Time `json:"actionDate" bun:"actionDate"`
}
