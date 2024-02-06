package policies

import (
	"github.com/agung96tm/miblog/api/models"
	"github.com/agung96tm/miblog/api/services"
	"github.com/agung96tm/miblog/constants"
	appErrors "github.com/agung96tm/miblog/errors"
	"github.com/labstack/echo/v4"
)

type UserPolicy struct {
	authService services.AuthService
}

func NewUserPolicy(authService services.AuthService) UserPolicy {
	return UserPolicy{
		authService: authService,
	}
}

func (u UserPolicy) CanUpdate(ctx echo.Context, userID uint) (bool, error) {
	user, ok := ctx.Get(constants.CurrentUser).(*models.User)
	if !ok || user.IsAnonymous() {
		return false, appErrors.ErrPolicyUnauthorized
	}
	if user.ID == userID {
		return true, appErrors.ErrPolicyForbidden
	}
	return false, nil
}

func (u UserPolicy) CanFollow(ctx echo.Context) error {
	user, ok := ctx.Get(constants.CurrentUser).(*models.User)
	if !ok || user.IsAnonymous() {
		return appErrors.ErrPolicyUnauthorized
	}
	return nil
}

func (u UserPolicy) CanUnFollow(ctx echo.Context) error {
	user, ok := ctx.Get(constants.CurrentUser).(*models.User)
	if !ok || user.IsAnonymous() {
		return appErrors.ErrPolicyUnauthorized
	}
	return nil
}
