package routes

import (
	"github.com/agung96tm/miblog/api/controllers"
	"github.com/agung96tm/miblog/lib"
)

type BlogRouter struct {
	handler        lib.HttpHandler
	blogController controllers.BlogController
}

func NewBlogRouter(handler lib.HttpHandler, blogController controllers.BlogController) BlogRouter {
	return BlogRouter{
		handler:        handler,
		blogController: blogController,
	}
}

func (r BlogRouter) Setup() {
	r.handler.Engine.GET("/blog_posts", r.blogController.List)
	r.handler.Engine.POST("/blog_posts", r.blogController.Create)
	r.handler.Engine.GET("/blog_posts/:id", r.blogController.Detail)
	r.handler.Engine.PATCH("/blog_posts/:id", r.blogController.Update)
	r.handler.Engine.DELETE("/blog_posts/:id", r.blogController.Delete)
}
