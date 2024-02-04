package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type CommentController struct {
}

func NewCommentController() CommentController {
	return CommentController{}
}

// List godoc
//
//	@Summary		Get Pagination and Several Comments
//	@Description	Get Pagination and Several Comments
//	@Tags			comment
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/comments [get]
//	@Success		200  {object}  response.Response{data=dto.CommentPagination}  "ok"
func (c CommentController) List(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "list")
}

// Detail godoc
//
//	@Summary		Get detail a Comment
//	@Description	Get detail a Comment
//	@Tags			comment
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/comments/{id} [get]
//	@Success		200  {object}  response.Response{data=dto.Comment}  "ok"
func (c CommentController) Detail(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "detail")
}

// Create godoc
//
//	@Summary		Create a Comment
//	@Description	Create a Comment
//	@Tags			comment
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/comments [post]
//	@Security 		BearerAuth
//	@Success		201  {object}  response.Response{data=dto.CommentCreateResponse}  "created"
func (c CommentController) Create(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "create")
}

// Update godoc
//
//	@Summary		Update a Comment
//	@Description	Update a Comment
//	@Tags			blog
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/comments/{id} [patch]
//	@Security 		BearerAuth
//	@Success		201  {object}  response.Response{data=dto.CommentUpdateResponse}  "created"
func (c CommentController) Update(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "update")
}

// Delete godoc
//
//	@Summary		Delete a Comment
//	@Description	Delete a Comment
//	@Tags			comment
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/comments/{id} [delete]
//	@Security 		BearerAuth
//	@Success		204  {object}  response.Response{}  "no content"
func (c CommentController) Delete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "delete")
}
