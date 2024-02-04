package services

import (
	"github.com/agung96tm/miblog/api/dto"
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

func (s BlogService) Get(id uint) (*dto.BlogPost, error) {
	post, err := s.blogPostRepository.Get(id)
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
