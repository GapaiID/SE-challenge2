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

func (u BlogPolicy) CanCreate(ctx echo.Context) error {
	user, ok := ctx.Get(constants.CurrentUser).(*models.User)
	if !ok || user.IsAnonymous() {
		return appErrors.ErrPolicyUnauthorized
	}
	return nil
}

func (u BlogPolicy) CanUpdate(ctx echo.Context, postID uint) error {
	user, ok := ctx.Get(constants.CurrentUser).(*models.User)
	if !ok || user.IsAnonymous() {
		return appErrors.ErrPolicyUnauthorized
	}

	postRes, err := u.blogService.Get(postID)
	if err != nil {
		return err
	}

	if user.ID != postRes.User.ID {
		return appErrors.ErrPolicyForbidden
	}
	return nil
}

func (u BlogPolicy) CanDelete(ctx echo.Context, postID uint) error {
	user, ok := ctx.Get(constants.CurrentUser).(*models.User)
	if !ok || user.IsAnonymous() {
		return appErrors.ErrPolicyUnauthorized
	}

	postRes, err := u.blogService.Get(postID)
	if err != nil {
		return err
	}

	if user.ID != postRes.User.ID {
		return appErrors.ErrPolicyForbidden
	}
	return nil
}
