package policies

import (
	"github.com/agung96tm/miblog/api/models"
	"github.com/agung96tm/miblog/api/services"
	"github.com/agung96tm/miblog/constants"
	appErrors "github.com/agung96tm/miblog/errors"
	"github.com/labstack/echo/v4"
)

type BlogPolicy struct {
	blogService services.BlogService
}

func NewBlogPolicy(blogService services.BlogService) BlogPolicy {
	return BlogPolicy{
		blogService: blogService,
	}
}

func (u BlogPolicy) CanCreate(ctx echo.Context) (bool, error) {
	user, ok := ctx.Get(constants.CurrentUser).(*models.User)
	if !ok || user.IsAnonymous() {
		return false, appErrors.ErrPolicyUnauthorized
	}
	return false, nil
}

func (u BlogPolicy) CanUpdate(ctx echo.Context, postID uint) (bool, error) {
	user, ok := ctx.Get(constants.CurrentUser).(*models.User)
	if !ok || user.IsAnonymous() {
		return false, appErrors.ErrPolicyUnauthorized
	}
	return false, nil
}

func (u BlogPolicy) CanDelete(ctx echo.Context, postID uint) (bool, error) {
	user, ok := ctx.Get(constants.CurrentUser).(*models.User)
	if !ok || user.IsAnonymous() {
		return false, appErrors.ErrPolicyUnauthorized
	}
	return false, nil
}
