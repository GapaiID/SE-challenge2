package services

import "github.com/agung96tm/miblog/api/repositories"

type BlogService struct {
	blogPostRepository repositories.BlogPostRepository
}

func NewBlogService(blogPostRepository repositories.BlogPostRepository) BlogService {
	return BlogService{
		blogPostRepository: blogPostRepository,
	}
}
