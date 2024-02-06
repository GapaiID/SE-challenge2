package routes

import (
	"github.com/GapaiID/SE-challenge2/api/controllers"
	"github.com/GapaiID/SE-challenge2/lib"
)

type BlogRouter struct {
	Handler        lib.HttpHandler
	blogController controllers.BlogController
}

func NewBlogRouter(Handler lib.HttpHandler, blogController controllers.BlogController) BlogRouter {
	return BlogRouter{
		Handler:        Handler,
		blogController: blogController,
	}
}

func (r BlogRouter) Setup() {
	r.Handler.Engine.GET("/following_blog_posts", r.blogController.FollowingBlogPostList)

	r.Handler.Engine.GET("/blog_posts", r.blogController.List)
	r.Handler.Engine.POST("/blog_posts", r.blogController.Create)
	r.Handler.Engine.GET("/blog_posts/:id", r.blogController.Detail)
	r.Handler.Engine.PATCH("/blog_posts/:id", r.blogController.Update)
	r.Handler.Engine.DELETE("/blog_posts/:id", r.blogController.Delete)
}
