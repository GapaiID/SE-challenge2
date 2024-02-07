package policies

import (
	"github.com/GapaiID/SE-challenge2/api/models"
	"github.com/GapaiID/SE-challenge2/api/services"
	"github.com/GapaiID/SE-challenge2/constants"
	appErrors "github.com/GapaiID/SE-challenge2/errors"
	"github.com/labstack/echo/v4"
)

type CommentPolicy struct {
	commentService services.CommentService
}

func NewCommentPolicy(commentService services.CommentService) CommentPolicy {
	return CommentPolicy{
		commentService: commentService,
	}
}

func (u CommentPolicy) CanCreate(ctx echo.Context) error {
	user, ok := ctx.Get(constants.CurrentUser).(*models.User)
	if !ok || user.IsAnonymous() {
		return appErrors.ErrPolicyUnauthorized
	}
	return nil
}

func (u CommentPolicy) CanUpdate(ctx echo.Context, commentID uint) error {
	user, ok := ctx.Get(constants.CurrentUser).(*models.User)
	if !ok || user.IsAnonymous() {
		return appErrors.ErrPolicyUnauthorized
	}

	commentRes, err := u.commentService.Get(commentID)
	if err != nil {
		return err
	}

	if user.ID != commentRes.User.ID {
		return appErrors.ErrPolicyForbidden
	}
	return nil
}

func (u CommentPolicy) CanDelete(ctx echo.Context, commentID uint) error {
	user, ok := ctx.Get(constants.CurrentUser).(*models.User)
	if !ok || user.IsAnonymous() {
		return appErrors.ErrPolicyUnauthorized
	}

	commentRes, err := u.commentService.Get(commentID)
	if err != nil {
		return err
	}

	if user.ID != commentRes.User.ID {
		return appErrors.ErrPolicyForbidden
	}
	return nil
}
