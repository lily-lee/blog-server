package models

import "time"

type Comment struct {
	ID        uint64    `json:"id"`
	PostID    uint64    `json:"post_id"`
	PID       uint64    `json:"pid"`
	UserID    uint64    `json:"user_id"`
	Anonymous bool      `json:"anonymous"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
