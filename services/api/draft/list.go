package draft

import (
	"github.com/gin-gonic/gin"

	"github.com/lily-lee/blog-server/config"
	"github.com/lily-lee/blog-server/models"
	"github.com/lily-lee/blog-server/services/jwttoken"
)

type (
	ListParam struct {
		VolumeID uint64 `form:"volume_id"`
		Page     int    `form:"page"`
		PerPage  int    `form:"per_page"`
	}

	ListResp struct {
		Page    int            `json:"page"`
		PerPage int            `json:"per_page"`
		Count   int64          `json:"count"`
		Data    []models.Draft `json:"data"`
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
	user, err := jwttoken.GetUser(c)
	if err != nil {
		return nil, err
	}

	drafts := make([]models.Draft, 0)
	db := config.DB

	where := db.Model(&models.Draft{}).Select(models.DraftSelectFields).
		Where("user_id = ?", user.ID).
		Where("post_id = ?", 0).
		Where("posted = ?", 0)
	if param.VolumeID > 0 {
		where = where.Where("volume_id = ?", param.VolumeID)
	}

	var count int64
	_ = where.Count(&count).Error

	err = where.Limit(param.PerPage).Offset((param.Page - 1) * param.PerPage).Find(&drafts).Error
	if err != nil {
		return nil, err
	}

	result := ListResp{
		Page:    param.Page,
		PerPage: param.PerPage,
		Count:   count,
		Data:    drafts,
	}

	return result, nil
}
