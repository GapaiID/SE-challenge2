package routes

import (
	"github.com/agung96tm/miblog/api/controllers"
	"github.com/agung96tm/miblog/lib"
)

type CommentRouter struct {
	handler       lib.HttpHandler
	commentRouter controllers.CommentController
}

func NewCommentRouter(handler lib.HttpHandler, commentRouter controllers.CommentController) CommentRouter {
	return CommentRouter{
		handler:       handler,
		commentRouter: commentRouter,
	}
}

func (r CommentRouter) Setup() {
	r.handler.Engine.GET("/comments", r.commentRouter.List)
	r.handler.Engine.POST("/comments", r.commentRouter.Create)
	r.handler.Engine.GET("/comments/:id", r.commentRouter.Detail)
	r.handler.Engine.PATCH("/comments/:id", r.commentRouter.Update)
	r.handler.Engine.DELETE("/comments/:id", r.commentRouter.Delete)
}
