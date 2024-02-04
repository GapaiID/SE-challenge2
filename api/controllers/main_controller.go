package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type MainController struct{}

func NewMainController() MainController {
	return MainController{}
}

func (c MainController) Index(ctx echo.Context) error {
	message := map[string]any{
		"message": "Hello World",
	}
	return ctx.JSON(http.StatusOK, message)
}
