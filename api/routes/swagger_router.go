package routes

import (
	"github.com/GapaiID/SE-challenge2/docs"
	"github.com/GapaiID/SE-challenge2/lib"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type SwaggerRouter struct {
	Handler lib.HttpHandler
	config  lib.Config
}

func NewSwaggerRouter(Handler lib.HttpHandler, config lib.Config) SwaggerRouter {
	return SwaggerRouter{
		Handler: Handler,
		config:  config,
	}
}

func (r SwaggerRouter) Setup() {
	docs.SwaggerInfo.Title = r.config.Swagger.Title
	docs.SwaggerInfo.Description = r.config.Swagger.Description
	docs.SwaggerInfo.Version = r.config.Swagger.Version

	r.Handler.Engine.GET("/docs/*", echoSwagger.WrapHandler)
}
