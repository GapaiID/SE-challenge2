package policies

import "go.uber.org/fx"

var Module = fx.Module(
	"policies",
	fx.Provide(NewUserPolicy),
	fx.Provide(fx.Annotate(NewBlogPolicy, fx.As(new(IBlogPolicy)))),
	fx.Provide(NewCommentPolicy),
)
