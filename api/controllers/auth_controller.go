package controllers

import (
	"github.com/agung96tm/miblog/api/dto"
	"github.com/agung96tm/miblog/api/services"
	"github.com/agung96tm/miblog/pkg/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return AuthController{
		authService: authService,
	}
}

// Register godoc
//
//	@Summary		Register a new User
//	@Description	register a new user
//	@Tags			auth
//	@Accept			application/json
//	@Produce		application/json
//	@Param 			data body dto.RegisterRequest true "Post"
//	@Router			/register [post]
//	@Success		201  {object}  response.Response{data=dto.RegisterResponse}  "created"
//	@Failure		400  {object}  response.Response{data=[]response.ValidationErrors}  "bad request"
func (c AuthController) Register(ctx echo.Context) error {
	registerReq := new(dto.RegisterRequest)
	if err := ctx.Bind(registerReq); err != nil {
		return response.Response{Error: err}.JSONValidationError(ctx)
	}

	registerResp, err := c.authService.Register(registerReq)
	if err != nil {
		return response.Response{
			Code:    http.StatusBadRequest,
			Message: err,
		}.JSON(ctx)
	}

	return response.Response{
		Code: http.StatusCreated,
		Data: registerResp,
	}.JSON(ctx)
}

// Login godoc
//
//	@Summary		Login User
//	@Description	Login user and get response JWT Token
//	@Tags			auth
//	@Accept			application/json
//	@Produce		application/json
//	@Param 			data body dto.LoginRequest true "Post"
//	@Router			/login [post]
//	@Success		200  {object}  response.Response{data=dto.LoginResponse}  "ok"
//	@Failure		400  {object}  response.Response{data=[]response.ValidationErrors}  "bad request"
func (c AuthController) Login(ctx echo.Context) error {
	loginReq := new(dto.LoginRequest)
	if err := ctx.Bind(loginReq); err != nil {
		return response.Response{Error: err}.JSONValidationError(ctx)
	}

	registerResp, err := c.authService.Login(loginReq)
	if err != nil {
		return response.Response{
			Code:    http.StatusBadRequest,
			Message: err,
		}.JSON(ctx)
	}

	return response.Response{
		Code: http.StatusOK,
		Data: registerResp,
	}.JSON(ctx)
}
