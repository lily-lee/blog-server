package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lily-lee/blog-server/services/api/post"
	"github.com/lily-lee/blog-server/services/request"
)

func ListPost(c *gin.Context) {
	param := new(post.ListParam)
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetPost(c *gin.Context) {
	param := new(post.GetParam)
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}

func CommentPost(c *gin.Context) {
	param := new(post.CreateCommentParam)
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}

// LikePost like or dislike post.
func LikePost(c *gin.Context) {
	param := new(post.LikePostParam)
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}
