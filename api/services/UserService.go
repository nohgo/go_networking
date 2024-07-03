package svc

import (
	"errors"

	"github.com/nohgo/go_networking/api/auth"
	repo "github.com/nohgo/go_networking/api/database/repositories"
	"github.com/nohgo/go_networking/api/models"
)

type UserService struct {
	ur repo.UserRepository
}

func NewUserService(ur repo.UserRepository) *UserService {
	return &UserService{ur: ur}
}
func (us *UserService) Register(user models.User) (err error) {
	err = us.ur.Add(user)
	return
}
func (us *UserService) GetAll(username string) ([]models.Car, error) {
	return us.ur.GetAll(username)
}
func (us *UserService) Login(user models.User) (token string, err error) {
	exists, err := us.ur.AreValidCredentials(user)

	if err != nil {
		return "", err
	}
	if !exists {
		return "", errors.New("Invalid credentials")
	}

	token, err = auth.CreateToken(user.Username)

	if err != nil {
		return "", err
	}

	return token, nil
}
