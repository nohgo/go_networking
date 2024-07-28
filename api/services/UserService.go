package svc

import (
	"errors"

	"github.com/nohgo/go_networking/api/auth"
	repo "github.com/nohgo/go_networking/api/database/repositories"
	"github.com/nohgo/go_networking/api/models"
)

// The UserService struct provides a way to access the user repository
// It contains one [repo.UserRepository]
type UserService struct {
	ur repo.UserRepository
}

// NewUserService returns a new [svc.UserService] and takes one [repo.UserRepository]
func NewUserService(ur repo.UserRepository) *UserService {
	return &UserService{ur: ur}
}

// Adds a user to the database
func (us *UserService) Register(user models.User) (err error) {
	if len(user.Username) == 0 || len(user.Password) == 0 {
		return &models.UserError{}
	}
	err = us.ur.Add(user)
	return
}

// Checks whether a user's credentials are in the datbase and returns a JWT on success
func (us *UserService) Login(user models.User) (token string, err error) {
	if len(user.Username) == 0 || len(user.Password) == 0 {
		return "", &models.UserError{}
	}
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

func (us *UserService) Delete(user models.User) error {
	if len(user.Username) == 0 {
		return &models.UserError{}
	}
	return us.ur.Delete(user)
}
