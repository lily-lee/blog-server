package models

import "time"

type Follower struct {
	UserID     uint64    `json:"user_id"`
	FollowerID uint64    `json:"follower_id"`
	CreatedAt  time.Time `json:"created_at"`
}
