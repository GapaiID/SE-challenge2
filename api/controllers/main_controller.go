package controllers

import (
	"github.com/agung96tm/miblog/constants"
	"github.com/agung96tm/miblog/lib"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type MainController struct {
	redis lib.Redis
}

func NewMainController(redis lib.Redis) MainController {
	return MainController{
		redis: redis,
	}
}

func (c MainController) Index(ctx echo.Context) error {
	var message map[string]any

	_ = c.redis.Get(constants.CacheBaseUrl, &message)
	if message == nil {
		message = map[string]any{"message": "OK"}
		_ = c.redis.Set(constants.CacheBaseUrl, message, time.Second*3600)
	}

	return ctx.JSON(http.StatusOK, message)
}
