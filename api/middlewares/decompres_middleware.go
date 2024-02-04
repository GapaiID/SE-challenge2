package middlewares

import (
	"github.com/agung96tm/miblog/lib"
	"github.com/labstack/echo/v4/middleware"
)

type DecompressMiddleware struct {
	handler lib.HttpHandler
}

func NewDecompressMiddleware(handler lib.HttpHandler) DecompressMiddleware {
	return DecompressMiddleware{
		handler: handler,
	}
}

func (m DecompressMiddleware) Setup() {
	m.handler.Engine.Use(middleware.Decompress())
}
