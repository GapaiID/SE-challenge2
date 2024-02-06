package routes

import (
	"github.com/GapaiID/SE-challenge2/api/controllers"
	"github.com/GapaiID/SE-challenge2/lib"
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
