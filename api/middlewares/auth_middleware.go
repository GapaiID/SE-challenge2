package middlewares

import (
	"github.com/agung96tm/miblog/api/models"
	"github.com/agung96tm/miblog/api/services"
	"github.com/agung96tm/miblog/constants"
	"github.com/agung96tm/miblog/lib"
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
