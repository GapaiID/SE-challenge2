package routes

import (
	"github.com/agung96tm/miblog/api/controllers"
	"github.com/agung96tm/miblog/lib"
)

type UserRouter struct {
	handler        lib.HttpHandler
	userController controllers.UserController
}

func NewUserRouter(handler lib.HttpHandler, userController controllers.UserController) UserRouter {
	return UserRouter{
		handler:        handler,
		userController: userController,
	}
}

func (u UserRouter) Setup() {
	u.handler.Engine.GET("/me", u.userController.Me)
}
