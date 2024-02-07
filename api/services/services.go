package services

import "go.uber.org/fx"

var Module = fx.Module(
	"services",
	fx.Provide(fx.Annotate(NewAuthService, fx.As(new(IAuthService)))),
	fx.Provide(NewUserService),
	fx.Provide(fx.Annotate(NewBlogService, fx.As(new(IBlogService)))),
	fx.Provide(NewCommentService),
)
