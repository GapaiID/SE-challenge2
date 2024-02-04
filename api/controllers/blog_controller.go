package controllers

import (
	"github.com/agung96tm/miblog/api/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BlogController struct {
	blogService services.BlogService
}

func NewBlogController(blogService services.BlogService) BlogController {
	return BlogController{
		blogService: blogService,
	}
}

// List godoc
//
//	@Summary		Get Pagination and Several Posts
//	@Description	Get Pagination and Several Posts
//	@Tags			blog
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/blog_posts [get]
//	@Success		200  {object}  response.Response{data=dto.BlogPostPagination}  "ok"
func (c BlogController) List(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "list")
}

// Detail godoc
//
//	@Summary		Get detail a post
//	@Description	Get detail a post
//	@Tags			blog
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/blog_posts/{id} [get]
//	@Success		200  {object}  response.Response{data=dto.BlogPost}  "ok"
func (c BlogController) Detail(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "detail")
}

// Create godoc
//
//	@Summary		Create a Post
//	@Description	Create a Post
//	@Tags			blog
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/blog_posts [post]
//	@Success		201  {object}  response.Response{}  "created"
func (c BlogController) Create(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "create")
}

// Update godoc
//
//	@Summary		Update a Post
//	@Description	Update a Post
//	@Tags			blog
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/blog_posts/{id} [patch]
//	@Success		200  {object}  response.Response{data=dto.BlogPost}  "ok"
func (c BlogController) Update(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "update")
}

// Delete godoc
//
//	@Summary		Delete a Post
//	@Description	Delete a Post
//	@Tags			blog
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/blog_posts/{id} [delete]
//	@Success		204  {object}  response.Response{}  "no content"
func (c BlogController) Delete(ctx echo.Context) error {
	return ctx.JSON(http.StatusNoContent, "delete")
}
