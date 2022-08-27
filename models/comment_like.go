package models

import "time"

type CommentLike struct {
	BlogID    uint64    `json:"blog_id"`
	CommentID uint64    `json:"comment_id"`
	UserID    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
