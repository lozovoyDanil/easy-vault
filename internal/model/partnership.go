package model

import "time"

type Partnership struct {
	UserId    int       `json:"userId" bun:"user_id,pk"`
	StartedAt time.Time `json:"startedAt" bun:"startTime"`
	Tier      int       `json:"tier" bun:"tier"`
}
