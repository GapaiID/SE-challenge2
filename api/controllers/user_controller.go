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

type UserController struct {
	authService services.AuthService
	userService services.UserService
	userPolicy  policies.UserPolicy
}

func NewUserController(authService services.AuthService, userService services.UserService, userPolicy policies.UserPolicy) UserController {
	return UserController{
		authService: authService,
		userService: userService,
		userPolicy:  userPolicy,
	}
}

// Me godoc
//
//	@Summary		Get data Logged-in user
//	@Description	Get data Logged-in user
//	@Tags			user
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/me [get]
//	@Security 		BearerAuth
//	@Success		200  {object}  response.Response{data=dto.MeResponse}  "ok"
func (c UserController) Me(ctx echo.Context) error {
	user := ctx.Get(constants.CurrentUser).(*models.User)
	if user.IsAnonymous() {
		return response.Response{
			Code: http.StatusUnauthorized,
		}.JSON(ctx)
	}

	return ctx.JSON(http.StatusOK, dto.MeResponse{
		Name:  user.Name,
		Email: user.Email,
	})
}

// MePassword godoc
//
//	@Summary		Get data Logged-in user
//	@Description	Get data Logged-in user
//	@Tags			user
//	@Accept			application/json
//	@Produce		application/json
//	@Param 			data body dto.MePasswordRequest true "Post"
//	@Router			/me/password [post]
//	@Security 		BearerAuth
//	@Success		200  {object}  response.Response{}  "ok"
func (c UserController) MePassword(ctx echo.Context) error {
	user := ctx.Get(constants.CurrentUser).(*models.User)
	if user.IsAnonymous() {
		return response.Response{
			Code: http.StatusUnauthorized,
		}.JSON(ctx)
	}

	var passwordReq dto.MePasswordRequest
	if err := ctx.Bind(&passwordReq); err != nil {
		return response.Response{
			Error: err,
		}.JSONValidationError(ctx)
	}

	err := c.userService.MeUpdatePassword(user, &passwordReq)
	if err != nil {
		return response.Response{
			Code:    http.StatusBadRequest,
			Message: err,
		}.JSON(ctx)
	}

	return response.Response{
		Code: http.StatusAccepted,
	}.JSON(ctx)
}

// MeUpdate godoc
//
//	@Summary		Update data Logged-in user
//	@Description	Update data Logged-in user
//	@Tags			user
//	@Accept			application/json
//	@Produce		application/json
//	@Param 			data body dto.MeUpdateRequest true "Post"
//	@Router			/me [patch]
//	@Security 		BearerAuth
//	@Success		200  {object}  response.Response{data=dto.MeResponse}  "ok"
func (c UserController) MeUpdate(ctx echo.Context) error {
	user := ctx.Get(constants.CurrentUser).(*models.User)
	if user.IsAnonymous() {
		return response.Response{
			Code: http.StatusUnauthorized,
		}.JSON(ctx)
	}

	var meRequest dto.MeUpdateRequest
	if err := ctx.Bind(&meRequest); err != nil {
		return response.Response{
			Error: err,
		}.JSONValidationError(ctx)
	}

	err := c.userService.MeUpdate(user, &meRequest)
	if err != nil {
		return response.Response{
			Code:    http.StatusBadRequest,
			Message: err,
		}.JSON(ctx)
	}

	return ctx.JSON(http.StatusOK, dto.MeResponse{
		Name:  user.Name,
		Email: user.Email,
	})
}

// List godoc
//
//	@Summary		Get Pagination and Several Users
//	@Description	Get Pagination and Several Users
//	@Tags			blog
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/users [get]
//	@Success		200  {object}  response.Response{data=dto.UserPagination}  "ok"
func (c UserController) List(ctx echo.Context) error {
	queryParams := new(dto.UserQueryParams)
	if err := ctx.Bind(queryParams); err != nil {
		return response.Response{
			Error: err,
		}.JSONValidationError(ctx)
	}

	paginationResp, err := c.userService.Query(queryParams)
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
//	@Summary		Get Detail User
//	@Description	Get Detail User
//	@Tags			user
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/users/{id} [get]
//	@Success		200  {object}  response.Response{data=dto.User}  "ok"
func (c UserController) Detail(ctx echo.Context) error {
	userID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return response.Response{
			Code:    http.StatusInternalServerError,
			Message: err,
		}.JSON(ctx)
	}

	userResp, err := c.userService.Get(uint(userID))
	if err != nil {
		return response.Response{
			Code:    http.StatusNotFound,
			Message: err,
		}.JSON(ctx)
	}

	return response.Response{
		Data: userResp,
	}.JSON(ctx)
}

// Follow godoc
//
//	@Summary		Following User
//	@Description	Following User
//	@Tags			user
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/users/{id}/follow [post]
//	@Security 		BearerAuth
//	@Success		202  {object}  response.Response{}  "accepted"
func (c UserController) Follow(ctx echo.Context) error {
	userID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return response.Response{
			Code:    http.StatusInternalServerError,
			Message: err,
		}.JSON(ctx)
	}

	err = c.userPolicy.CanFollow(ctx)
	if err != nil {
		return response.Response{
			Error: err,
		}.JSONPolicyError(ctx)
	}

	followReq := &dto.FollowerCreateRequest{UserID: uint(userID)}
	user := ctx.Get(constants.CurrentUser).(*models.User)
	if err := c.userService.Follow(user, followReq); err != nil {
		return response.Response{
			Code:    http.StatusBadRequest,
			Message: err,
		}.JSON(ctx)
	}

	return response.Response{
		Code: http.StatusAccepted,
	}.JSON(ctx)
}

// UnFollow godoc
//
//	@Summary		UnFollowing User
//	@Description	UnFollowing User
//	@Tags			user
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/users/{id}/unfollow [post]
//	@Security 		BearerAuth
//	@Success		202  {object}  response.Response{}  "accepted"
func (c UserController) UnFollow(ctx echo.Context) error {
	userID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return response.Response{
			Code:    http.StatusInternalServerError,
			Message: err,
		}.JSON(ctx)
	}

	err = c.userPolicy.CanUnFollow(ctx)
	if err != nil {
		return response.Response{
			Error: err,
		}.JSONPolicyError(ctx)
	}

	unFollowReq := &dto.UnFollowerCreateRequest{UserID: uint(userID)}
	user := ctx.Get(constants.CurrentUser).(*models.User)
	if err := c.userService.UnFollow(user, unFollowReq); err != nil {
		return response.Response{
			Code:    http.StatusBadRequest,
			Message: err,
		}.JSON(ctx)
	}

	return response.Response{
		Code: http.StatusAccepted,
	}.JSON(ctx)
}
