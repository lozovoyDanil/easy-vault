package model

import "time"

type Subscription struct {
	Id        int
	Tier      int
	ExpiresAt time.Time
}
