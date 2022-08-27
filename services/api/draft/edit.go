package draft

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lily-lee/blog-server/config"
	"github.com/lily-lee/blog-server/models"
	"github.com/lily-lee/blog-server/services/jwttoken"
	"github.com/lily-lee/blog-server/services/request"
	"github.com/lily-lee/blog-server/services/types"
)

type EditParam struct {
	ID       uint64    `uri:"id"`
	VolumeID uint64    `json:"volume_id"`
	Title    string    `json:"title" binding:"required"`
	Content  string    `json:"content" binding:"required"`
	Digest   string    `json:"digest"`
	CoverURL string    `json:"cover_url"`
	Tag      types.Tag `json:"tag"`
}

func (param *EditParam) Do(c *gin.Context) (interface{}, error) {
	if err := c.BindUri(param); err != nil {
		return nil, &request.BizErr{HttpCode: http.StatusUnprocessableEntity, ErrMsg: err.Error()}
	}

	user, err := jwttoken.GetUser(c)
	if err != nil {
		return nil, err
	}

	db := config.DB
	db = db.Begin()
	defer db.Rollback()

	draft := &models.Draft{}
	err = db.Where("id = ? and user_id = ?", param.ID, user.ID).Take(draft).Error
	if err != nil {
		return nil, err
	}

	err = db.Model(&models.Draft{}).Where("id = ?", param.ID).Updates(&models.Draft{
		VolumeID: param.VolumeID,
		Title:    param.Title,
		Content:  param.Content,
	}).Error
	if err != nil {
		return nil, err
	}

	return nil, db.Commit().Error
}
