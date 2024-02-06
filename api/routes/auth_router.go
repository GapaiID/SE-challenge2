package routes

import (
	"github.com/GapaiID/SE-challenge2/api/controllers"
	"github.com/GapaiID/SE-challenge2/lib"
)

type AuthRouter struct {
	handler        lib.HttpHandler
	authController controllers.AuthController
}

func NewAuthRouter(handler lib.HttpHandler, authController controllers.AuthController) AuthRouter {
	return AuthRouter{
		handler:        handler,
		authController: authController,
	}
}

func (r AuthRouter) Setup() {
	r.handler.Engine.POST("/auth/register", r.authController.Register)
	r.handler.Engine.POST("/auth/login", r.authController.Login)
}
