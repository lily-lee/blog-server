package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lily-lee/blog-server/config"
	"github.com/lily-lee/blog-server/models"
	"github.com/lily-lee/blog-server/services/jwttoken"
	"github.com/lily-lee/blog-server/services/request"
)

type (
	GetParam struct {
		ID uint64 `uri:"id"`
	}

	Post struct {
		models.Post
		LikeCount    int64 `json:"like_count"`
		CommentCount int64 `json:"comment_count"`
		Liked        bool  `json:"liked"`
	}
)

func (param *GetParam) Do(c *gin.Context) (interface{}, error) {
	if err := c.BindUri(param); err != nil {
		return nil, &request.BizErr{HttpCode: http.StatusUnprocessableEntity, ErrMsg: err.Error()}
	}

	user, _ := jwttoken.GetUser(c)
	db := config.DB

	post := models.Post{}
	err := db.Model(&models.Post{}).Where("id = ?", param.ID).Take(&post).Error
	if err != nil {
		return nil, err
	}

	resp := Post{
		Post:         post,
		LikeCount:    GetLikeNum(param.ID),
		CommentCount: GetCommentNum(param.ID),
	}

	if user != nil {
		err = db.Model(&models.PostLike{}).
			Where("post_id = ?", param.ID).Where("user_id = ?", user.ID).Take(&models.PostLike{}).Error
		resp.Liked = err == nil
	}

	return resp, nil
}
