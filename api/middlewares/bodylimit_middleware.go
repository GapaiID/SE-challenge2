package middlewares

import (
	"github.com/agung96tm/miblog/lib"
	"github.com/labstack/echo/v4/middleware"
)

type BodyLimitMiddleware struct {
	handler lib.HttpHandler
}

func NewBodyLimitMiddleware(handler lib.HttpHandler) BodyLimitMiddleware {
	return BodyLimitMiddleware{
		handler: handler,
	}
}

func (m BodyLimitMiddleware) Setup() {
	m.handler.Engine.Use(middleware.BodyLimit("2M"))
}
