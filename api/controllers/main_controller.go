package controllers

import (
	"github.com/GapaiID/SE-challenge2/constants"
	"github.com/GapaiID/SE-challenge2/lib"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type MainController struct {
	redis lib.IRedis
}

func NewMainController(redis lib.IRedis) MainController {
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
