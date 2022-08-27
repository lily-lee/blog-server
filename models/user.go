package models

import (
	"time"

	"github.com/lily-lee/blog-server/services/types"
)

type User struct {
	ID        uint64         `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Avatar    string         `json:"avatar"`
	Gender    int            `json:"gender"`
	Birthday  types.Birthday `json:"birthday"`
	Salt      string         `json:"-"`
	Password  string         `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
