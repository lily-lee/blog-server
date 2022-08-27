package models

import "time"

type PostLike struct {
	PostID    uint64    `json:"post_id"`
	UserID    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
