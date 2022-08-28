package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lily-lee/blog-server/config"
	_const "github.com/lily-lee/blog-server/const"
	_ "github.com/lily-lee/blog-server/docs"
	_ "github.com/lily-lee/blog-server/services/validators"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type IRouter interface {
	Register(*gin.Engine)
}

func New() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())

	// serve swagger api docs
	if config.Conf.Env == _const.EnvDev {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	new(userRouter).Register(r)
	new(draftRouter).Register(r)
	new(postRouter).Register(r)

	return r
}
