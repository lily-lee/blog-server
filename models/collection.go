package models

import "time"

type Collection struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	UserID      uint64    `json:"user_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
