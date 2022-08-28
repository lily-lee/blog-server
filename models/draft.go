package models

import (
	"time"

	"github.com/lily-lee/blog-server/services/types"
)

var DraftSelectFields = []string{"id", "post_id", "volume_id", "user_id", "title", "digest", "cover_url", "tag", "created_at", "updated_at"}

type Draft struct {
	ID        uint64    `json:"id"`
	PostID    uint64    `json:"post_id"`
	VolumeID  uint64    `json:"volume_id"`
	UserID    uint64    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content,omitempty"`
	Digest    string    `json:"digest"`
	CoverURL  string    `json:"cover_url"`
	Tag       types.Tag `json:"tag"`
	Posted    bool      `json:"posted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
