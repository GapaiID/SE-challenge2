package services

import (
	"errors"
	"github.com/agung96tm/miblog/api/dto"
	"github.com/agung96tm/miblog/api/models"
	"github.com/agung96tm/miblog/api/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return UserService{
		userRepository: userRepository,
	}
}

func (s UserService) MeUpdatePassword(user *models.User, passwordReq *dto.MePasswordRequest) error {
	if !passwordReq.PasswordMatches() {
		return errors.New("new and confirm password not matches")
	}
	valid, err := user.CheckPassword(passwordReq.OldPassword)
	if err != nil || !valid {
		return errors.New("invalid old password")
	}

	user.Password = passwordReq.NewPassword

	err = s.userRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (s UserService) MeUpdate(user *models.User, meReq *dto.MeUpdateRequest) error {
	if meReq.Name != "" {
		user.Name = meReq.Name
	}

	err := s.userRepository.Update(user)
	if err != nil {
		return err
	}

	return nil
}
