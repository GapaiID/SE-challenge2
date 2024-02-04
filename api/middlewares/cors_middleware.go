package middlewares

import (
	"github.com/agung96tm/miblog/lib"
	"github.com/labstack/echo/v4/middleware"
)

type CorsMiddleware struct {
	handler lib.HttpHandler
}

func NewCorsMiddleware(handler lib.HttpHandler) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
	}
}

func (m CorsMiddleware) Setup() {
	m.handler.Engine.Use(middleware.CORS())
}
