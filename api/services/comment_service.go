package services

import (
	"github.com/agung96tm/miblog/api/dto"
	"github.com/agung96tm/miblog/api/models"
)

type CommentService struct {
}

func NewCommentService() CommentService {
	return CommentService{}
}

func (s CommentService) Query(params *dto.BlogPostQueryParams) (any, error) {
	return nil, nil
}

func (s CommentService) Get(postID uint) (*dto.BlogPost, error) {
	return nil, nil
}

func (s CommentService) Create(user *models.User, postReq *dto.BlogPostCreateRequest) (*dto.BlogPostCreateResponse, error) {
	return nil, nil
}

func (s CommentService) Delete(postID uint) error {
	return nil
}
