package router

import (
	"github.com/gin-gonic/gin"

	"github.com/lily-lee/blog-server/controllers"
	"github.com/lily-lee/blog-server/middlewares"
)

type postRouter struct{}

func (*postRouter) Register(r *gin.Engine) {
	r.GET("/api/posts", middlewares.JwtAuth(true), controllers.ListPost)
	r.GET("/api/posts/:id", middlewares.JwtAuth(true), controllers.GetPost)

	posts := r.Group("/api/posts", middlewares.JwtAuth())
	{
		posts.POST("/:id/comments", controllers.CommentPost)
		posts.POST("/:id/like", controllers.LikePost)
	}
}
