package draft

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/lily-lee/blog-server/config"
	"github.com/lily-lee/blog-server/models"
	"github.com/lily-lee/blog-server/services/jwttoken"
	"github.com/lily-lee/blog-server/services/types"
	"github.com/sirupsen/logrus"
)

type CreateParam struct {
	VolumeID uint64    `json:"volume_id"`
	Title    string    `json:"title" binding:"required"`
	Content  string    `json:"content" binding:"required"`
	Digest   string    `json:"digest"`
	CoverURL string    `json:"cover_url"`
	Tag      types.Tag `json:"tag"`
}

func (param *CreateParam) Do(c *gin.Context) (interface{}, error) {
	user, err := jwttoken.GetUser(c)
	if err != nil {
		return nil, err
	}

	draftId, err := config.SnowFlake.NextID()
	if err != nil {
		logrus.Error("get snow id failed.", err)
		return nil, err
	}

	draft := new(models.Draft)
	_ = copier.Copy(draft, param)
	draft.ID = draftId
	draft.UserID = user.ID

	db := config.DB
	err = db.Create(draft).Error

	return draft, err
}
