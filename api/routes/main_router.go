package routes

import (
	"github.com/agung96tm/miblog/api/controllers"
	"github.com/agung96tm/miblog/docs"
	"github.com/agung96tm/miblog/lib"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type MainRouter struct {
	handler        lib.HttpHandler
	mainController controllers.MainController
	config         lib.Config
}

func NewMainRouter(handler lib.HttpHandler, mainController controllers.MainController, config lib.Config) MainRouter {
	return MainRouter{
		handler:        handler,
		mainController: mainController,
		config:         config,
	}
}

func (r MainRouter) Setup() {
	r.handler.Engine.GET("/", r.mainController.Index)

	docs.SwaggerInfo.Title = r.config.Swagger.Title
	docs.SwaggerInfo.Description = r.config.Swagger.Description
	docs.SwaggerInfo.Version = r.config.Swagger.Version

	r.handler.Engine.GET("/docs/*", echoSwagger.WrapHandler)
}
