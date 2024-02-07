package policies

import (
	"github.com/GapaiID/SE-challenge2/api/models"
	"github.com/GapaiID/SE-challenge2/api/services"
	"github.com/GapaiID/SE-challenge2/constants"
	appErrors "github.com/GapaiID/SE-challenge2/errors"
	"github.com/labstack/echo/v4"
)

type UserPolicy struct {
	authService services.IAuthService
}

func NewUserPolicy(authService services.IAuthService) UserPolicy {
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
