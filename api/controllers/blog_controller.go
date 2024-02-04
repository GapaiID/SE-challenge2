package controllers

import (
	"github.com/agung96tm/miblog/api/dto"
	"github.com/agung96tm/miblog/api/models"
	"github.com/agung96tm/miblog/api/policies"
	"github.com/agung96tm/miblog/api/services"
	"github.com/agung96tm/miblog/constants"
	"github.com/agung96tm/miblog/pkg/response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type BlogController struct {
	blogService services.BlogService
	blogPolicy  policies.BlogPolicy
}

func NewBlogController(blogService services.BlogService, blogPolicy policies.BlogPolicy) BlogController {
	return BlogController{
		blogService: blogService,
		blogPolicy:  blogPolicy,
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
	queryParams := new(dto.BlogPostQueryParams)
	if err := ctx.Bind(queryParams); err != nil {
		return response.Response{
			Error: err,
		}.JSONValidationError(ctx)
	}

	paginationResp, err := c.blogService.Query(queryParams)
	if err != nil {
		return response.Response{
			Code:    http.StatusBadRequest,
			Message: err,
		}.JSON(ctx)
	}

	return response.Response{
		Data: paginationResp,
	}.JSON(ctx)
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
	postID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return response.Response{
			Code:    http.StatusInternalServerError,
			Message: err,
		}.JSON(ctx)
	}

	postResp, err := c.blogService.Get(uint(postID))
	if err != nil {
		return response.Response{
			Code:    http.StatusNotFound,
			Message: err,
		}.JSON(ctx)
	}

	return response.Response{
		Code: http.StatusOK,
		Data: postResp,
	}.JSON(ctx)
}

// Create godoc
//
//	@Summary		Create a Post
//	@Description	Create a Post
//	@Tags			blog
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/blog_posts [post]
//	@Success		201  {object}  response.Response{data=dto.BlogPostCreateRequest}  "created"
func (c BlogController) Create(ctx echo.Context) error {
	err := c.blogPolicy.CanCreate(ctx)
	if err != nil {
		return response.Response{
			Error: err,
		}.JSONPolicyError(ctx)
	}

	blogPostReq := new(dto.BlogPostCreateRequest)
	if err := ctx.Bind(blogPostReq); err != nil {
		return response.Response{
			Error: err,
		}.JSONValidationError(ctx)
	}

	user := ctx.Get(constants.CurrentUser).(*models.User)
	postResp, err := c.blogService.Create(user, blogPostReq)
	if err != nil {
		return response.Response{
			Code:    http.StatusBadRequest,
			Message: err,
		}.JSON(ctx)
	}

	return response.Response{
		Code: http.StatusCreated,
		Data: postResp,
	}.JSON(ctx)
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
	postID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return response.Response{
			Code:    http.StatusInternalServerError,
			Message: err,
		}.JSON(ctx)
	}

	err = c.blogPolicy.CanDelete(ctx, uint(postID))
	if err != nil {
		return response.Response{
			Error: err,
		}.JSONPolicyError(ctx)
	}

	err = c.blogService.Delete(uint(postID))
	if err != nil {
		return response.Response{
			Code:    http.StatusBadRequest,
			Message: err,
		}.JSON(ctx)
	}

	return response.Response{
		Code: http.StatusNoContent,
	}.JSON(ctx)
}
