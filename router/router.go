package router

import (
	"github.com/gin-gonic/gin"

	_ "github.com/lily-lee/blog-server/services/validators"
)

type IRouter interface {
	Register(*gin.Engine)
}

func New() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())

	new(userRouter).Register(r)
	new(draftRouter).Register(r)
	new(postRouter).Register(r)

	return r
}
