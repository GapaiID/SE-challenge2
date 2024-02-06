package routes_test

import (
	"errors"
	"github.com/GapaiID/SE-challenge2/api/dto"
	"github.com/GapaiID/SE-challenge2/api/models"
	"time"
)

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
