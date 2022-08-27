package user

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/lily-lee/blog-server/config"
	"github.com/lily-lee/blog-server/models"
	"github.com/lily-lee/blog-server/services/common"
	"github.com/lily-lee/blog-server/services/jwttoken"
	"github.com/lily-lee/blog-server/services/request"
	"github.com/lily-lee/blog-server/services/types"
	"github.com/sirupsen/logrus"
)

type (
	RegisterParam struct {
		Name     string         `json:"name" binding:"required"`
		Email    string         `json:"email" binding:"required,email"`
		Mobile   string         `json:"mobile"`
		Avatar   string         `json:"avatar"`
		Gender   int            `json:"gender"`
		Birthday types.Birthday `json:"birthday" binding:"birthday"`
		Password string         `json:"password" binding:"required"`
	}

	RegisterResp struct {
		User  *models.User    `json:"user"`
		Token *jwttoken.Token `json:"token"`
	}
)

func (param *RegisterParam) Do(c *gin.Context) (interface{}, error) {
	user := new(models.User)
	if err := copier.Copy(user, param); err != nil {
		return nil, err
	}

	salt := common.Random(6)
	password := common.EncodePassword(salt, param.Password)

	uid, err := config.SnowFlake.NextID()
	if err != nil {
		logrus.Error("user get snowflake id failed.")
		return nil, err
	}
	user.ID = uid
	user.Salt = salt
	user.Password = password

	db := config.DB
	err = db.Create(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, &request.BizErr{HttpCode: http.StatusUnprocessableEntity, ErrMsg: "user exists"}
		}
		return nil, err
	}

	token, err := jwttoken.GenToken(user)
	if err != nil {
		return nil, err
	}

	resp := RegisterResp{
		User:  user,
		Token: token,
	}

	return resp, nil
}
