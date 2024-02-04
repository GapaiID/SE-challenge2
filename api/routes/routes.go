package routes

import "go.uber.org/fx"

var Module = fx.Module(
	"routes",
	fx.Provide(NewMainRouter),
	fx.Provide(NewAuthRouter),
	fx.Provide(NewUserRouter),
	fx.Provide(NewRoutes),
	fx.Provide(NewBlogRouter),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(mainRouter MainRouter, authRouter AuthRouter, userRouter UserRouter, blogRouter BlogRouter) Routes {
	return Routes{
		mainRouter,
		authRouter,
		userRouter,
		blogRouter,
	}
}

func (routes Routes) Setup() {
	for _, router := range routes {
		router.Setup()
	}
}
