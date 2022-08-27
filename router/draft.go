package router

import (
	"github.com/gin-gonic/gin"

	"github.com/lily-lee/blog-server/controllers"
	"github.com/lily-lee/blog-server/middlewares"
)

type draftRouter struct{}

func (*draftRouter) Register(r *gin.Engine) {
	drafts := r.Group("/api/drafts", middlewares.JwtAuth())
	{
		drafts.POST("", controllers.CreateDraft)
		drafts.PUT("/:id", controllers.EditDraft)
		drafts.POST("/:id/posts", controllers.PostDraft)
		drafts.GET("/:id", controllers.GetDraft)
		drafts.GET("", controllers.ListDraft)
	}
}
