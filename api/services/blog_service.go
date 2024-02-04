package services

import (
	"github.com/agung96tm/miblog/api/dto"
	"github.com/agung96tm/miblog/api/models"
	"github.com/agung96tm/miblog/api/repositories"
)

type BlogService struct {
	blogPostRepository repositories.BlogPostRepository
}

func NewBlogService(blogPostRepository repositories.BlogPostRepository) BlogService {
	return BlogService{
		blogPostRepository: blogPostRepository,
	}
}

func (s BlogService) Get(postID uint) (*dto.BlogPost, error) {
	post, err := s.blogPostRepository.Get(postID)
	if err != nil {
		return nil, err
	}

	return &dto.BlogPost{
		ID:    post.ID,
		Title: post.Title,
		Body:  post.Body,
		User: &dto.BlogUser{
			ID:   post.User.ID,
			Name: post.User.Name,
		},
	}, nil
}

func (s BlogService) Create(user *models.User, postReq *dto.BlogPostCreateRequest) (*dto.BlogPostCreateResponse, error) {
	var post models.BlogPost

	post.Title = postReq.Title
	post.Body = postReq.Body
	post.UserID = user.ID

	err := s.blogPostRepository.Create(&post)
	if err != nil {
		return nil, err
	}

	return &dto.BlogPostCreateResponse{
		ID:    post.ID,
		Title: post.Title,
		Body:  post.Body,
	}, nil
}

func (s BlogService) Delete(postID uint) error {
	var post models.BlogPost
	post.ID = postID

	err := s.blogPostRepository.Delete(&post)
	if err != nil {
		return err
	}

	return nil
}
