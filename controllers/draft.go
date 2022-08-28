package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lily-lee/blog-server/services/api/draft"
	"github.com/lily-lee/blog-server/services/request"
)

// CreateDraft create a draft article.
// @Summary     create a draft article
// @Description create a draft article
// @Tags        drafts
// @Accept      json
// @Produce     json
// @Param       draft           body     draft.CreateParam true "create param"
// @Param       Authorization   header   string            true "login token"
// @Success     200             {object} models.Draft
// @Failure     401,404,422,500 {object} request.BizErr
// @Router      /api/drafts [post]
func CreateDraft(c *gin.Context) {
	param := new(draft.CreateParam)
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}

// EditDraft edit an article.
// @Summary     edit an article
// @Description edit an article
// @Tags        drafts
// @Accept      json
// @Produce     json
// @Param       id              path     int64           true "draft id"
// @Param       draft           body     draft.EditParam true "create param"
// @Param       Authorization   header   string          true "login token"
// @Success     200             {object} models.Draft
// @Failure     401,404,422,500 {object} request.BizErr
// @Router      /api/drafts/{id} [put]
func EditDraft(c *gin.Context) {
	param := new(draft.EditParam)
	_, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, nil)
}

// PostDraft post an article.
// @Summary     edit an article
// @Description edit an article
// @Tags        drafts
// @Accept      json
// @Produce     json
// @Param       id              path     int64  true "draft id"
// @Param       Authorization   header   string true "login token"
// @Success     200             {object} models.Draft
// @Failure     401,404,422,500 {object} request.BizErr
// @Router      /api/drafts/{id}/posts [post]
func PostDraft(c *gin.Context) {
	param := new(draft.PostParam)
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}

// ListDraft return draft list
// @Summary     list drafts
// @Description list drafts
// @Tags        drafts
// @Accept      json
// @Produce     json
// @Param       param           query    draft.ListParam true "draft id"
// @Param       Authorization   header   string          true "login token"
// @Success     200             {object} draft.ListResp
// @Failure     401,404,422,500 {object} request.BizErr
// @Router      /api/drafts [get]
func ListDraft(c *gin.Context) {
	param := new(draft.ListParam)
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetDraft return one draft
// @Summary     get a specific article
// @Description get a specific article
// @Tags        drafts
// @Accept      json
// @Produce     json
// @Param       id              path     int64  true "draft id"
// @Param       Authorization   header   string true "login token"
// @Success     200             {object} models.Draft
// @Failure     401,404,422,500 {object} request.BizErr
// @Router      /api/drafts/{id} [get]
func GetDraft(c *gin.Context) {
	param := new(draft.GetParam)
	data, err := request.Handle(c, param)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}
