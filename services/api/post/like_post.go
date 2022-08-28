package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lily-lee/blog-server/config"
	"github.com/lily-lee/blog-server/models"
	"github.com/lily-lee/blog-server/services/cache"
	"github.com/lily-lee/blog-server/services/jwttoken"
	"github.com/lily-lee/blog-server/services/request"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type (
	LikePostParam struct {
		PostID uint64 `uri:"id" json:"-"`
	}

	LikeResp struct {
		PostID uint64 `json:"post_id"`
		UserID uint64 `json:"user_id"`
		Like   bool   `json:"like"`
	}
)

func (param *LikePostParam) Do(c *gin.Context) (interface{}, error) {
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

	like := new(models.PostLike)
	err = db.Model(&models.PostLike{}).
		Where("post_id = ?", param.PostID).
		Where("user_id = ?", user.ID).Take(like).Error

	resp := LikeResp{PostID: param.PostID, UserID: user.ID}
	switch err {
	case nil:
		err = db.Model(like).
			Where("post_id = ?", param.PostID).
			Where("user_id = ?", user.ID).Delete(like).Error
	case gorm.ErrRecordNotFound:
		resp.Like = true
		like.PostID = param.PostID
		like.UserID = user.ID
		err = db.Create(like).Error
	default:
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	SaveNum(LikeKey(param.PostID), resp.Like, param.PostID, GetLikeNum)

	return resp, nil
}

func GetLikeNum(postId uint64) int64 {
	key := LikeKey(postId)
	num, err := cache.GetInt64(key, func() (int64, error) {
		var count int64
		err := config.DB.Model(&models.PostLike{}).Where("post_id = ?", postId).Count(&count).Error
		return count, err
	})

	if err != nil {
		logrus.Errorf("get like num failed. post_id:%d, err:%v", postId, err)
	}

	return num
}
