package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/lily-lee/blog-server/config"
	"github.com/lily-lee/blog-server/models"
	"github.com/lily-lee/blog-server/services/cache"
	"github.com/lily-lee/blog-server/services/jwttoken"
	"github.com/lily-lee/blog-server/services/request"
	"github.com/sirupsen/logrus"
)

type CreateCommentParam struct {
	PostID    uint64 `uri:"id" json:"-"`
	PID       uint64 `json:"pid"`
	Content   string `json:"content" binding:"required"`
	Anonymous bool   `json:"anonymous"`
}

func (param *CreateCommentParam) Do(c *gin.Context) (interface{}, error) {
	if err := c.BindUri(param); err != nil {
		return nil, &request.BizErr{HttpCode: http.StatusUnprocessableEntity, ErrMsg: err.Error()}
	}

	user, err := jwttoken.GetUser(c)
	if err != nil {
		return nil, err
	}

	db := config.DB
	err = db.Model(models.Post{}).Where("id = ?", param.PostID).Take(&models.Post{}).Error
	if err != nil {
		return nil, err
	}

	comment := new(models.Comment)
	if err = copier.Copy(comment, param); err != nil {
		return nil, err
	}

	cid, err := config.SnowFlake.NextID()
	if err != nil {
		logrus.Error("comment get snowflake failed.err:", err)
	}

	if cid > 0 {
		comment.ID = cid
	}

	comment.UserID = user.ID
	err = db.Create(comment).Error
	if err != nil {
		return nil, err
	}

	if param.Anonymous {
		comment.UserID = 0
	}

	SaveNum(CommentKey(param.PostID), true, param.PostID, GetCommentNum)

	return comment, nil
}

func GetCommentNum(postId uint64) int64 {
	key := CommentKey(postId)
	num, err := cache.GetInt64(key, func() (int64, error) {
		var count int64
		err := config.DB.Model(&models.Comment{}).Where("post_id = ?", postId).Count(&count).Error
		return count, err
	})

	if err != nil {
		logrus.Errorf("get comment num failed. post_id:%d, err:%v", postId, err)
	}

	return num
}
