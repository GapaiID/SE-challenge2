package services

import (
	"github.com/GapaiID/SE-challenge2/api/dto"
	"github.com/GapaiID/SE-challenge2/api/models"
	"github.com/GapaiID/SE-challenge2/api/repositories"
)

type IBlogService interface {
	Query(params *dto.BlogPostQueryParams) (*dto.BlogPostPagination, error)
	Get(postID uint) (*dto.BlogPost, error)
	Create(user *models.User, postReq *dto.BlogPostCreateRequest) (*dto.BlogPostCreateResponse, error)
	Update(userID *models.User, postID uint, postReq *dto.BlogPostUpdateRequest) (*dto.BlogPostUpdateResponse, error)
	Delete(postID uint) error
	QueryByFollowing(user *models.User, params *dto.BlogPostQueryParams) (*dto.BlogPostPagination, error)
}

type BlogService struct {
	blogPostRepository repositories.BlogPostRepository
}

func NewBlogService(blogPostRepository repositories.BlogPostRepository) BlogService {
	return BlogService{
		blogPostRepository: blogPostRepository,
	}
}

func (s BlogService) Query(params *dto.BlogPostQueryParams) (*dto.BlogPostPagination, error) {
	list, pagination, err := s.blogPostRepository.Query(params)
	if err != nil {
		return nil, err
	}

	var posts []*dto.BlogPost
	for _, post := range *list {
		posts = append(posts, &dto.BlogPost{
			ID:    post.ID,
			Title: post.Title,
			Body:  post.Body,
			User: &dto.UserInBlogPost{
				ID:   post.User.ID,
				Name: post.User.Name,
			},
		})
	}

	return &dto.BlogPostPagination{
		List:       posts,
		Pagination: pagination,
	}, nil
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
		User: &dto.UserInBlogPost{
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

func (s BlogService) Update(userID *models.User, postID uint, postReq *dto.BlogPostUpdateRequest) (*dto.BlogPostUpdateResponse, error) {
	post, err := s.blogPostRepository.Get(postID)
	if err != nil {
		return nil, err
	}

	if postReq.Title != "" {
		post.Title = postReq.Title
	}
	if postReq.Body != "" {
		post.Body = postReq.Body
	}

	if err := s.blogPostRepository.Update(postID, post); err != nil {
		return nil, err
	}

	return &dto.BlogPostUpdateResponse{
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

func (s BlogService) QueryByFollowing(user *models.User, params *dto.BlogPostQueryParams) (*dto.BlogPostPagination, error) {
	list, pagination, err := s.blogPostRepository.QueryByFollowing(user, params)
	if err != nil {
		return nil, err
	}

	var posts []*dto.BlogPost
	for _, post := range *list {
		posts = append(posts, &dto.BlogPost{
			ID:    post.ID,
			Title: post.Title,
			Body:  post.Body,
			User: &dto.UserInBlogPost{
				ID:   post.User.ID,
				Name: post.User.Name,
			},
		})
	}

	return &dto.BlogPostPagination{
		List:       posts,
		Pagination: pagination,
	}, nil
}
