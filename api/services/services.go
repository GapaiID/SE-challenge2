package services

import "go.uber.org/fx"

var Module = fx.Module(
	"services",
	fx.Provide(NewAuthService),
	fx.Provide(NewUserService),
	fx.Provide(NewBlogService),
	fx.Provide(NewCommentService),
)
