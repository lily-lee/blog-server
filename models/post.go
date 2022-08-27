package models

import (
	"time"

	"github.com/lily-lee/blog-server/services/types"
)

var PostSelectFields = []string{"id", "draft_id", "volume_id", "user_id", "title", "digest", "cover_url", "tag", "created_at", "updated_at"}

type Post struct {
	ID        uint64    `json:"id"`
	DraftID   uint64    `json:"draft_id"`
	VolumeID  uint64    `json:"volume_id"`
	UserID    uint64    `json:"user_id"`
	Title     string    `json:"title"`
	Digest    string    `json:"digest"`
	CoverURL  string    `json:"cover_url"`
	Tag       types.Tag `json:"tag"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
