package draft

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lily-lee/blog-server/config"
	"github.com/lily-lee/blog-server/models"
	"github.com/lily-lee/blog-server/services/jwttoken"
	"github.com/lily-lee/blog-server/services/request"
)

type GetParam struct {
	ID uint64 `uri:"id" json:"-"`
}

func (param *GetParam) Do(c *gin.Context) (interface{}, error) {
	if err := c.BindUri(param); err != nil {
		return nil, &request.BizErr{HttpCode: http.StatusUnprocessableEntity, ErrMsg: err.Error()}
	}

	user, err := jwttoken.GetUser(c)
	if err != nil {
		return nil, err
	}

	draft := &models.Draft{}
	db := config.DB
	err = db.Where("id = ?", param.ID).Where("user_id = ?", user.ID).Take(draft).Error

	return draft, err
}
