package services

import (
	"errors"
	"profoliohub-api/models"
	"profoliohub-api/repositories"
	"profoliohub-api/utils"
)

type AuthService struct {
	userRepository *repositories.UserRepositories
}

func NewAuthService(userRepo *repositories.UserRepositories) *AuthService {
	return &AuthService{
		userRepository: userRepo,
	}
}

func (s *AuthService) RegisterUser(user models.User) error {
	existingUser, _ := s.userRepository.GetUserByEmail(user.Email)
	if existingUser.Email != "" {
		return errors.New("user already exists")
	}

	user.Password = utils.HashPassword(user.Password)
	return s.userRepository.CreateUser(user)
}

func (s *AuthService) Login(email string, password string) (string, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID.Hex())
	if err != nil {
		return "", err
	}

	return token, nil
}
