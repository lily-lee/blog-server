package models

import "time"

type ArticleCollection struct {
	CollectionID uint64    `json:"collection_id"`
	ArticleID    uint64    `json:"article_id"`
	UserID       uint64    `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
}
