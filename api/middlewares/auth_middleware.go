package middlewares

import (
	"github.com/GapaiID/SE-challenge2/api/models"
	"github.com/GapaiID/SE-challenge2/api/services"
	"github.com/GapaiID/SE-challenge2/constants"
	"github.com/GapaiID/SE-challenge2/lib"
	"github.com/labstack/echo/v4"
	"strings"
)

type AuthMiddleware struct {
	handler     lib.HttpHandler
	authService services.AuthService
}

func NewAuthMiddleware(handler lib.HttpHandler, authService services.AuthService) AuthMiddleware {
	return AuthMiddleware{
		handler:     handler,
		authService: authService,
	}
}

func (m AuthMiddleware) Setup() {
	m.handler.Engine.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			request := ctx.Request()

			var (
				auth   = request.Header.Get("Authorization")
				prefix = "Bearer "
				token  string
			)

			if auth != "" && strings.HasPrefix(auth, prefix) {
				token = auth[len(prefix):]
			}

			user, err := m.authService.AuthorizeJWTToken(token)
			if err != nil {
				user = models.AnonymousUser
			}

			ctx.Set(constants.CurrentUser, user)
			return next(ctx)
		}
	})
}
