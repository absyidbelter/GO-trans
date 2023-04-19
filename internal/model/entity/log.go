package model

import "time"

type Log struct {
	ID        int       `json:"id"`
	UserID    uint      `json:"user_id" `
	Event     string    `json:"event" `
	CreatedAt time.Time `json:"created_at" `
}
