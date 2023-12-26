package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Partnership struct {
	bun.BaseModel `bun:"Partnership,alias:part"`

	UserId    int       `json:"userId" bun:"user_id,pk"`
	StartedAt time.Time `json:"startedAt" bun:"startTime"`
	Tier      int       `json:"tier" bun:"tier"`
}
