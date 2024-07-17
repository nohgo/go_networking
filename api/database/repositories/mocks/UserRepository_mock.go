package repo_mock

import (
	"errors"

	"github.com/nohgo/go_networking/api/models"
)

type mockUserRepository struct {
}

func NewMockUserRepository() *mockUserRepository {
	return &mockUserRepository{}
}

func (*mockUserRepository) Add(user models.User) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("Invalid fields")
	}
	return nil
}

func (*mockUserRepository) AreValidCredentials(user models.User) (bool, error) {
	if user.Username == "" || user.Password == "" {
		return false, errors.New("Invalid fields")
	}
	return true, nil
}

func (*mockUserRepository) Delete(user models.User) error {
	if user.Username == "" {
		return errors.New("invalid fields")
	}
	return nil
}
