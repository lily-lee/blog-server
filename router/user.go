package router

import (
	"github.com/gin-gonic/gin"

	"github.com/lily-lee/blog-server/controllers"
	"github.com/lily-lee/blog-server/middlewares"
)

type userRouter struct{}

func (*userRouter) Register(r *gin.Engine) {
	r.POST("/api/registration", controllers.Register)
	r.POST("/api/login", controllers.Login)

	user := r.Group("/api/users", middlewares.JwtAuth())
	{
		user.GET("/:id", controllers.Register)
		user.GET("/:id/followers")
		user.POST("/:id/followers")
		user.DELETE("/:id/followers")
	}
}
