package policies

import (
	"github.com/agung96tm/miblog/api/models"
	"github.com/agung96tm/miblog/api/services"
	"github.com/agung96tm/miblog/constants"
	appErrors "github.com/agung96tm/miblog/errors"
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
	return nil
}

func (u CommentPolicy) CanUpdate(ctx echo.Context, postID uint) (bool, error) {
	return false, nil
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
