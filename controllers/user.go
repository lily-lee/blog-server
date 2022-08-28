package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lily-lee/blog-server/services/api/user"
	"github.com/lily-lee/blog-server/services/request"
)

// Register provides user registration http api.
// @Summary     user register
// @Description register a user
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       param           body     user.RegisterParam true "user register param"
// @Success     200             {object} user.RegisterResp
// @Failure     401,404,422,500 {object} request.BizErr
// @Router      /api/registration [post]
func Register(c *gin.Context) {
	param := new(user.RegisterParam)
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}

// Login returns user information and jwt token.
// @Summary     user login
// @Description user login
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       param           body     user.LoginParam true "user login param"
// @Success     200             {object} user.RegisterResp
// @Failure     401,404,422,500 {object} request.BizErr
// @Router      /api/login [post]
func Login(c *gin.Context) {
	param := new(user.LoginParam)
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}

func Follow(c *gin.Context) {

}
