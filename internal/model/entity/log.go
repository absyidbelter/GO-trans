package model

import "time"

type Log struct {
	ID        uint
	UserID    uint
	Log       string
	CreatedAt time.Time
}
