package routes

import (
	"github.com/GapaiID/SE-challenge2/api/controllers"
	"github.com/GapaiID/SE-challenge2/lib"
)

type AuthRouter struct {
	Handler        lib.HttpHandler
	authController controllers.AuthController
}

func NewAuthRouter(handler lib.HttpHandler, authController controllers.AuthController) AuthRouter {
	return AuthRouter{
		Handler:        handler,
		authController: authController,
	}
}

func (r AuthRouter) Setup() {
	r.Handler.Engine.POST("/auth/register", r.authController.Register)
	r.Handler.Engine.POST("/auth/login", r.authController.Login)
}
