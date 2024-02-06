package lib

import "go.uber.org/fx"

var Module = fx.Module(
	"lib",
	fx.Provide(NewHttpHandler),
	fx.Provide(NewDatabase),
	fx.Provide(NewConfig),
	fx.Provide(NewMigration),
	fx.Provide(NewJWT),
	fx.Provide(NewMail),
	fx.Provide(fx.Annotate(NewRedis, fx.As(new(IRedis)))),
)
