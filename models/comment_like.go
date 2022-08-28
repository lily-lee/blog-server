package models

import "time"

type CommentLike struct {
	PostID    uint64    `json:"post_id"`
	CommentID uint64    `json:"comment_id"`
	UserID    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
