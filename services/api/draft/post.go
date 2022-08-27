package draft

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"

	"github.com/lily-lee/blog-server/config"
	"github.com/lily-lee/blog-server/models"
	"github.com/lily-lee/blog-server/services/jwttoken"
	"github.com/lily-lee/blog-server/services/request"
)

type PostParam struct {
	ID uint64 `uri:"id"`
}

func (param *PostParam) Do(c *gin.Context) (interface{}, error) {
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

	post := &models.Post{
		ID:       draft.PostID,
		DraftID:  draft.ID,
		VolumeID: draft.VolumeID,
		UserID:   user.ID,
		Title:    draft.Title,
		Content:  draft.Content,
		Digest:   draft.Digest,
		CoverURL: draft.CoverURL,
		Tag:      draft.Tag,
	}

	postId := draft.PostID

	if postId == 0 {
		postId, err = config.SnowFlake.NextID()
		post.ID = postId
		if err != nil {
			logrus.Error("get snow id failed.", err)
			return nil, err
		}
	}

	err = db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(post).Error

	if err != nil {
		return nil, err
	}

	err = db.Model(&models.Draft{}).Where("id = ?", draft.ID).Updates(&models.Draft{
		PostID: postId,
		Posted: true,
	}).Error

	post.ID = postId

	return post, db.Commit().Error
}
