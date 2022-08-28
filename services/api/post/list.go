package post

import (
	"github.com/gin-gonic/gin"
	"github.com/lily-lee/blog-server/config"
	"github.com/lily-lee/blog-server/models"
)

type (
	ListParam struct {
		Keyword  string `form:"keyword"`
		UserID   uint64 `form:"user_id"`
		VolumeID uint64 `form:"volume_id"`
		Page     int    `form:"page"`
		PerPage  int    `form:"per_page"`
	}

	ListResp struct {
		Page    int    `json:"page"`
		PerPage int    `json:"per_page"`
		Count   int64  `json:"count"`
		Data    []Item `json:"data"`
	}

	Item struct {
		models.Post
		LikeCount    int64 `json:"like_count"`
		CommentCount int64 `json:"comment_count"`
	}
)

func (param *ListParam) Check() error {
	if param.Page <= 0 {
		param.Page = 1
	}

	if param.PerPage <= 0 {
		param.PerPage = 10
	}

	return nil
}

func (param *ListParam) Do(c *gin.Context) (interface{}, error) {
	db := config.DB
	where := db.Model(&models.Post{}).Select(models.PostSelectFields)
	if param.Keyword != "" {
		where = where.Where("match(title, digest, content) against(?)", param.Keyword)
	}

	if param.UserID > 0 {
		where = where.Where("user_id = ?", param.UserID)
	}

	if param.VolumeID > 0 {
		where = where.Where("volume_id = ?", param.VolumeID)
	}

	var count int64
	_ = where.Count(&count).Error

	posts := make([]models.Post, 0)
	err := where.Limit(param.PerPage).Offset((param.Page - 1) * param.PerPage).Order("id DESC").Find(&posts).Error

	resp := ListResp{
		Page:    param.Page,
		PerPage: param.PerPage,
		Count:   count,
		Data:    make([]Item, len(posts)),
	}

	for i := range posts {
		resp.Data[i].Post = posts[i]
		resp.Data[i].CommentCount = GetCommentNum(resp.Data[i].ID)
		resp.Data[i].LikeCount = GetLikeNum(resp.Data[i].ID)
	}

	return resp, err
}
