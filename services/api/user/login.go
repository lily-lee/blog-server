package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lily-lee/blog-server/config"
	"github.com/lily-lee/blog-server/models"
	"github.com/lily-lee/blog-server/services/common"
	"github.com/lily-lee/blog-server/services/jwttoken"
	"github.com/lily-lee/blog-server/services/request"
	"gorm.io/gorm"
)

type LoginParam struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (param *LoginParam) Do(c *gin.Context) (interface{}, error) {
	db := config.DB
	user := &models.User{}

	err := db.Where("email = ?", param.Email).Take(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &request.BizErr{HttpCode: http.StatusNotFound, ErrMsg: "user not found"}
		}
		return nil, err
	}

	if !common.CheckPassword(user.Salt, param.Password, user.Password) {
		return nil, &request.BizErr{HttpCode: http.StatusUnprocessableEntity, ErrMsg: "email or password error"}
	}

	token, err := jwttoken.GenToken(user)
	if err != nil {
		return nil, err
	}

	return RegisterResp{User: user, Token: token}, nil
}
