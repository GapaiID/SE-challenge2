package controllers

import (
	"github.com/agung96tm/miblog/api/dto"
	"github.com/agung96tm/miblog/api/models"
	"github.com/agung96tm/miblog/api/services"
	"github.com/agung96tm/miblog/constants"
	"github.com/agung96tm/miblog/pkg/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	authService services.AuthService
}

func NewUserController(authService services.AuthService) UserController {
	return UserController{
		authService: authService,
	}
}

func (c UserController) Me(ctx echo.Context) error {
	user, ok := ctx.Get(constants.CurrentUser).(*models.User)
	if !ok || user.IsAnonymous() {
		return response.Response{
			Code: http.StatusUnauthorized,
		}.JSON(ctx)
	}

	return ctx.JSON(http.StatusOK, dto.MeResponse{
		Name:  user.Name,
		Email: user.Email,
	})
}
