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

func (r UserRouter) Setup() {
	r.handler.Engine.GET("/me", r.userController.Me)
	r.handler.Engine.PATCH("/me", r.userController.MeUpdate)
	r.handler.Engine.POST("/me/password", r.userController.MePassword)

	r.handler.Engine.GET("/users/:id", r.userController.Detail)
	r.handler.Engine.POST("/users/:id/follow", r.userController.Follow)
}
