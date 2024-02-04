package services

import (
	"github.com/agung96tm/miblog/api/dto"
	"github.com/agung96tm/miblog/api/models"
	"github.com/agung96tm/miblog/api/repositories"
)

type CommentService struct {
	commentRepository repositories.CommentRepository
}

func NewCommentService(commentRepository repositories.CommentRepository) CommentService {
	return CommentService{
		commentRepository: commentRepository,
	}
}

func (s CommentService) Query(params *dto.CommentQueryParams) (any, error) {
	list, pagination, err := s.commentRepository.Query(params)
	if err != nil {
		return nil, err
	}

	var comments []*dto.Comment
	for _, comment := range *list {
		comments = append(comments, &dto.Comment{
			ID:   comment.ID,
			Body: comment.Body,
			User: &dto.UserInComment{
				ID:   comment.User.ID,
				Name: comment.User.Name,
			},
			Post: &dto.PostInComment{
				ID:    comment.Post.ID,
				Title: comment.Post.Title,
			},
		})
	}

	return &dto.CommentPagination{
		List:       comments,
		Pagination: pagination,
	}, nil
}

func (s CommentService) Get(commentID uint) (*dto.Comment, error) {
	comment, err := s.commentRepository.Get(commentID)
	if err != nil {
		return nil, err
	}

	return &dto.Comment{
		ID:   comment.ID,
		Body: comment.Body,
		User: &dto.UserInComment{
			ID:   comment.User.ID,
			Name: comment.User.Name,
		},
		Post: &dto.PostInComment{
			ID:    comment.Post.ID,
			Title: comment.Post.Title,
		},
	}, nil
}

func (s CommentService) Create(user *models.User, postReq *dto.BlogPostCreateRequest) (*dto.BlogPostCreateResponse, error) {
	return nil, nil
}

func (s CommentService) Delete(commentID uint) error {
	var post models.Comment
	post.ID = commentID

	err := s.commentRepository.Delete(&post)
	if err != nil {
		return err
	}

	return nil
}
