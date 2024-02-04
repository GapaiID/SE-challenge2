package repositories

import "go.uber.org/fx"

var Module = fx.Module(
	"repositories",
	fx.Provide(NewUserRepository),
	fx.Provide(NewBlogPostRepository),
	fx.Provide(NewCommentRepository),
)
