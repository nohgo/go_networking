package svc

import (
	repo "github.com/nohgo/go_networking/api/database/repositories"
	"github.com/nohgo/go_networking/api/models"
)

type UserService struct {
	ur repo.UserRepository
}

func NewUserService(ur repo.UserRepository) *UserService {
	return &UserService{ur: ur}
}
func (us *UserService) Register(name string, pass string) (err error) {
	err = us.ur.Add(models.User{Username: name, Password: pass})
	return
}
func (us *UserService) GetAll() ([]models.User, error) {
	return us.ur.GetAll()
}
