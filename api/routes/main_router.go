package routes

import (
	"github.com/agung96tm/miblog/api/controllers"
	"github.com/agung96tm/miblog/lib"
)

type MainRouter struct {
	Handler        lib.HttpHandler
	mainController controllers.MainController
}

func NewMainRouter(Handler lib.HttpHandler, mainController controllers.MainController) MainRouter {
	return MainRouter{
		Handler:        Handler,
		mainController: mainController,
	}
}

func (r MainRouter) Setup() {
	r.Handler.Engine.GET("/", r.mainController.Index)
}
