package repo_mock

import (
	"errors"

	"github.com/nohgo/go_networking/api/models"
)

type MockUserRepository struct {
}

func NewMockUserRepository() MockUserRepository {
	return MockUserRepository{}
}

func (*MockUserRepository) Add(user models.User) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("Invalid fields")
	}
	return nil
}
