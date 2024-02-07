package services

import (
	"errors"
	"github.com/GapaiID/SE-challenge2/api/dto"
	"github.com/GapaiID/SE-challenge2/api/mails"
	"github.com/GapaiID/SE-challenge2/api/models"
	"github.com/GapaiID/SE-challenge2/api/repositories"
	"github.com/GapaiID/SE-challenge2/lib"
)

type IAuthService interface {
	Register(registerReq *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(loginReq *dto.LoginRequest) (*dto.LoginResponse, error)
	AuthorizeJWTToken(token string) (*models.User, error)
}

type AuthService struct {
	userRepository repositories.UserRepository
	jwt            lib.JWT
	authMail       mails.AuthMail
}

func NewAuthService(userRepository repositories.UserRepository, jwt lib.JWT, authMail mails.AuthMail) AuthService {
	return AuthService{
		userRepository: userRepository,
		jwt:            jwt,
		authMail:       authMail,
	}
}

func (a AuthService) Register(registerReq *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	user := models.User{
		Name:     registerReq.Name,
		Password: registerReq.Password,
		Email:    registerReq.Email,
	}

	if err := a.userRepository.Create(&user); err != nil {
		return nil, err
	}

	go func() {
		a.authMail.Register(&user)
	}()

	return &dto.RegisterResponse{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (a AuthService) Login(loginReq *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := a.userRepository.GetByEmail(loginReq.Email)
	if err != nil {
		return nil, err
	}

	valid, err := user.CheckPassword(loginReq.Password)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, errors.New("email or Password not matched")
	}

	token, err := a.jwt.GenerateToken(int64(user.ID))
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token: token,
	}, nil
}

func (a AuthService) AuthorizeJWTToken(token string) (*models.User, error) {
	userID, err := a.jwt.GetSubjectFromToken(token)
	if err != nil {
		return nil, err
	}

	user, err := a.userRepository.Get(uint(userID))
	if err != nil {
		return nil, err
	}

	return user, nil
}
