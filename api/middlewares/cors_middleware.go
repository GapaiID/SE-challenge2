package middlewares

import (
	"github.com/agung96tm/miblog/lib"
	"github.com/labstack/echo/v4/middleware"
)

type CorsMiddleware struct {
	handler    lib.HttpHandler
	corsConfig *lib.CorsConfig
}

func NewCorsMiddleware(handler lib.HttpHandler, config lib.Config) CorsMiddleware {
	return CorsMiddleware{
		handler:    handler,
		corsConfig: config.Cors,
	}
}

func (m CorsMiddleware) Setup() {
	m.handler.Engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: m.corsConfig.AllowOrigins,
		AllowMethods: m.corsConfig.AllowMethods,
	}))
}
