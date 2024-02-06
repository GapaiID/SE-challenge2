package routes_test

import (
	"errors"
	"github.com/GapaiID/SE-challenge2/api/dto"
	"github.com/GapaiID/SE-challenge2/api/models"
	"github.com/labstack/echo/v4"
	"time"
)

/* ===================================================
/*	Lib
/* =================================================== */

type LibRedisMock struct {
	IsCalled map[string]any
}

func NewLibRedisMock() *LibRedisMock {
	return &LibRedisMock{
		IsCalled: make(map[string]interface{}),
	}
}

func (r *LibRedisMock) Set(key string, value any, expiration time.Duration) error {
	r.IsCalled["Set"] = []any{value, expiration}
	return nil
}

func (r *LibRedisMock) Get(key string, value any) error {
	r.IsCalled["Get"] = []any{key, value}
	return nil
}

func (r *LibRedisMock) Delete(keys ...string) (bool, error) {
	return false, nil
}

func (r *LibRedisMock) Check(keys ...string) (bool, error) {
	return false, nil
}

type AuthServiceMock struct{}

/* ===================================================
/*	Services
/* =================================================== */

func NewAuthServiceMock() AuthServiceMock {
	return AuthServiceMock{}
}

func (auth AuthServiceMock) Register(registerReq *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	if registerReq.Email == "valid@example.com" {
		return &dto.RegisterResponse{
			Email: registerReq.Email,
			Name:  registerReq.Name,
		}, nil
	}
	return nil, errors.New("register invalid")
}
func (auth AuthServiceMock) Login(loginReq *dto.LoginRequest) (*dto.LoginResponse, error) {
	if loginReq.Email == "valid@example.com" {
		return &dto.LoginResponse{
			Token: "token-jwt",
		}, nil
	}
	return nil, errors.New("invalid email")
}
func (auth AuthServiceMock) AuthorizeJWTToken(token string) (*models.User, error) {
	return nil, nil
}

type BlogServiceMock struct{}

func NewBlogServiceMock() BlogServiceMock {
	return BlogServiceMock{}
}

func (s BlogServiceMock) Query(params *dto.BlogPostQueryParams) (*dto.BlogPostPagination, error) {
	return &dto.BlogPostPagination{
		List: []*dto.BlogPost{
			{
				ID:    1,
				Title: "Blog 1",
				Body:  "Description 1",
				User:  &dto.UserInBlogPost{ID: 1, Name: "Dragon"},
			},
			{
				ID:    2,
				Title: "Blog 2",
				Body:  "Description 2",
				User:  &dto.UserInBlogPost{ID: 1, Name: "Dragon"},
			},
		},
		Pagination: &dto.Pagination{
			PageSize: 1,
			Total:    2,
			LastPage: 1,
			Current:  1,
		},
	}, nil
}
func (s BlogServiceMock) Get(postID uint) (*dto.BlogPost, error) {
	if postID == 1 {
		return &dto.BlogPost{
			ID:    1,
			Title: "Blog 1",
			Body:  "Description 1",
			User:  &dto.UserInBlogPost{ID: 1, Name: "Dragon"},
		}, nil
	}
	return nil, errors.New("not found")
}
func (s BlogServiceMock) Create(user *models.User, postReq *dto.BlogPostCreateRequest) (*dto.BlogPostCreateResponse, error) {
	if user.ID != 1 {
		return nil, errors.New("unauthorized")
	}

	return &dto.BlogPostCreateResponse{
		ID:    1,
		Title: postReq.Title,
		Body:  postReq.Body,
	}, nil
}
func (s BlogServiceMock) Update(userID *models.User, postID uint, postReq *dto.BlogPostUpdateRequest) (*dto.BlogPostUpdateResponse, error) {
	return nil, nil
}
func (s BlogServiceMock) Delete(postID uint) error {
	return nil
}
func (s BlogServiceMock) QueryByFollowing(user *models.User, params *dto.BlogPostQueryParams) (*dto.BlogPostPagination, error) {
	return nil, nil
}

/* ===================================================
/*	Policies
/* =================================================== */

type BlogPolicyMock struct{}

func NewBlogPolicyMock() BlogPolicyMock {
	return BlogPolicyMock{}
}

func (BlogPolicyMock) CanCreate(ctx echo.Context) error {
	return nil
}
func (BlogPolicyMock) CanUpdate(ctx echo.Context, postID uint) error {
	return nil
}
func (BlogPolicyMock) CanDelete(ctx echo.Context, postID uint) error {
	return nil
}
func (BlogPolicyMock) CanSeeFollowingPosts(ctx echo.Context) error {
	return nil
}
