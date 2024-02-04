package policies

import (
	"github.com/labstack/echo/v4"
)

type CommentPolicy struct {
}

func NewCommentPolicy() CommentPolicy {
	return CommentPolicy{}
}

func (u CommentPolicy) CanCreate(ctx echo.Context) error {
	return nil
}

func (u CommentPolicy) CanUpdate(ctx echo.Context, postID uint) (bool, error) {
	return false, nil
}

func (u CommentPolicy) CanDelete(ctx echo.Context, postID uint) error {
	return nil
}
