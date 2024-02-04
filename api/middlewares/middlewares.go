package middlewares

import "go.uber.org/fx"

var Module = fx.Module(
	"middlewares",
	fx.Provide(NewBodyLimitMiddleware),
	fx.Provide(NewCorsMiddleware),
	fx.Provide(NewDecompressMiddleware),
	fx.Provide(NewGZipMiddleware),
	fx.Provide(NewAuthMiddleware),

	fx.Provide(NewMiddlewares),
)

type Middlewares []Middleware

type Middleware interface {
	Setup()
}

func NewMiddlewares(
	bodyLimit BodyLimitMiddleware,
	cors CorsMiddleware,
	decompress DecompressMiddleware,
	gzip GZipMiddleware,
	auth AuthMiddleware,
) Middlewares {
	return Middlewares{
		bodyLimit,
		cors,
		decompress,
		gzip,
		auth,
	}
}

func (middlewares Middlewares) Setup() {
	for _, middleware := range middlewares {
		middleware.Setup()
	}
}
