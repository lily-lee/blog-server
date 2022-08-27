package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lily-lee/blog-server/services/api/draft"
	"github.com/lily-lee/blog-server/services/request"
)

// CreateDraft create a draft article.
func CreateDraft(c *gin.Context) {
	param := new(draft.CreateParam)
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}

// EditDraft edit an article.
func EditDraft(c *gin.Context) {
	param := new(draft.EditParam)
	_, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, nil)
}

// PostDraft post an article.
func PostDraft(c *gin.Context) {
	param := new(draft.PostParam)
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}

// ListDraft return draft list
func ListDraft(c *gin.Context) {
	param := new(draft.ListParam)
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetDraft return one draft
func GetDraft(c *gin.Context) {
	param := new(draft.GetParam)
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}
