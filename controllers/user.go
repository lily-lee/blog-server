package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lily-lee/blog-server/services/api/user"
	"github.com/lily-lee/blog-server/services/request"
)

// Register provides user registration http api.
func Register(c *gin.Context) {
	param := &user.RegisterParam{}
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}

// Login returns user information and jwt token.
func Login(c *gin.Context) {
	param := &user.LoginParam{}
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}

func Follow(c *gin.Context) {

}
