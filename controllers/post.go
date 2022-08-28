package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lily-lee/blog-server/services/api/post"
	"github.com/lily-lee/blog-server/services/request"
)

// ListPost list posts
// @Summary     list posts
// @Description list posts
// @Tags        posts
// @Accept      json
// @Produce     json
// @Param       param           query    post.ListParam true  "draft id"
// @Param       Authorization   header   string         false "login token"
// @Success     200             {object} post.ListResp
// @Failure     401,404,422,500 {object} request.BizErr
// @Router      /api/posts [get]
func ListPost(c *gin.Context) {
	param := new(post.ListParam)
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetPost get one post
// @Summary     get a specific post
// @Description return one post
// @Tags        posts
// @Accept      json
// @Produce     json
// @Param       id              path     int64  true "post id"
// @Param       Authorization   header   string false "login token"
// @Success     200             {object} post.Post
// @Failure     401,404,422,500 {object} request.BizErr
// @Router      /api/posts/{id} [get]
func GetPost(c *gin.Context) {
	param := new(post.GetParam)
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}

// CommentPost comment a post
// @Summary     comment a post
// @Description comment a post
// @Tags        posts
// @Accept      json
// @Produce     json
// @Param       id              path     int64  				 true "post id"
// @Param       param           body     post.CreateCommentParam true "create comment param"
// @Param       Authorization   header   string                  true "login token"
// @Success     200             {object} models.Comment
// @Failure     401,404,422,500 {object} request.BizErr
// @Router      /api/posts/{id}/comments [post]
func CommentPost(c *gin.Context) {
	param := new(post.CreateCommentParam)
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}

// LikePost like or dislike post.
// @Summary     like or dislike post.
// @Description When you request this api if you had like a post, then you'll unlike the post.
// @Tags        posts
// @Accept      json
// @Produce     json
// @Param       id              path     int64  true "post id"
// @Param       Authorization   header   string true "login token"
// @Success     200             {object} post.LikeResp
// @Failure     401,404,422,500 {object} request.BizErr
// @Router      /api/posts/{id}/like [post]
func LikePost(c *gin.Context) {
	param := new(post.LikePostParam)
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}
