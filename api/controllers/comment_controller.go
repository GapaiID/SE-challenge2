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

type CommentController struct {
	commentService services.CommentService
	commentPolicy  policies.CommentPolicy
}

func NewCommentController(commentService services.CommentService, commentPolicy policies.CommentPolicy) CommentController {
	return CommentController{
		commentService: commentService,
		commentPolicy:  commentPolicy,
	}
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
	queryParams := new(dto.CommentQueryParams)
	if err := ctx.Bind(queryParams); err != nil {
		return response.Response{
			Error: err,
		}.JSONValidationError(ctx)
	}

	paginationResp, err := c.commentService.Query(queryParams)
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
//	@Summary		Get detail a Comment
//	@Description	Get detail a Comment
//	@Tags			comment
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/comments/{id} [get]
//	@Success		200  {object}  response.Response{data=dto.Comment}  "ok"
func (c CommentController) Detail(ctx echo.Context) error {
	commentID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return response.Response{
			Code:    http.StatusInternalServerError,
			Message: err,
		}.JSON(ctx)
	}

	postResp, err := c.commentService.Get(uint(commentID))
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
//	@Summary		Create a Comment
//	@Description	Create a Comment
//	@Tags			comment
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/comments [post]
//	@Security 		BearerAuth
//	@Success		201  {object}  response.Response{data=dto.CommentCreateResponse}  "created"
func (c CommentController) Create(ctx echo.Context) error {
	err := c.commentPolicy.CanCreate(ctx)
	if err != nil {
		return response.Response{
			Error: err,
		}.JSONPolicyError(ctx)
	}

	commentReq := new(dto.CommentCreateRequest)
	if err := ctx.Bind(commentReq); err != nil {
		return response.Response{
			Error: err,
		}.JSONValidationError(ctx)
	}

	user := ctx.Get(constants.CurrentUser).(*models.User)
	postResp, err := c.commentService.Create(user, commentReq)
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
//	@Summary		Update a Comment
//	@Description	Update a Comment
//	@Tags			comment
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/comments/{id} [patch]
//	@Security 		BearerAuth
//	@Success		200  {object}  response.Response{data=dto.CommentUpdateResponse}  "ok"
func (c CommentController) Update(ctx echo.Context) error {
	commentID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return response.Response{
			Code:    http.StatusInternalServerError,
			Message: err,
		}.JSON(ctx)
	}

	err = c.commentPolicy.CanUpdate(ctx, uint(commentID))
	if err != nil {
		return response.Response{
			Error: err,
		}.JSONPolicyError(ctx)
	}

	commentReq := new(dto.CommentUpdateRequest)
	if err := ctx.Bind(commentReq); err != nil {
		return response.Response{
			Error: err,
		}.JSONValidationError(ctx)
	}

	user := ctx.Get(constants.CurrentUser).(*models.User)
	postResp, err := c.commentService.Update(user, uint(commentID), commentReq)
	if err != nil {
		return response.Response{
			Code:    http.StatusBadRequest,
			Message: err,
		}.JSON(ctx)
	}

	return response.Response{
		Code: http.StatusOK,
		Data: postResp,
	}.JSON(ctx)
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
	commentID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return response.Response{
			Code:    http.StatusInternalServerError,
			Message: err,
		}.JSON(ctx)
	}

	err = c.commentPolicy.CanDelete(ctx, uint(commentID))
	if err != nil {
		return response.Response{
			Error: err,
		}.JSONPolicyError(ctx)
	}

	err = c.commentService.Delete(uint(commentID))
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
