package model

import "time"

type Subscription struct {
	Id        int       `bun:"id,pk,autoincrement"`
	UserId    int       `bun:"user_id"`
	StartedAt time.Time `json:"startedAt" bun:"startTime"`
	ExpiresAt time.Time `json:"expiresAt" bun:"endTime"`
	IsActive  bool      `json:"isActive" bun:"isActive"`
	Tier      int       `json:"tier" bun:"tier"`
}
