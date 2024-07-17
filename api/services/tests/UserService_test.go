package svc_test

import (
	"os"
	"testing"

	repo_mock "github.com/nohgo/go_networking/api/database/repositories/mocks"
	"github.com/nohgo/go_networking/api/models"
	svc "github.com/nohgo/go_networking/api/services"
)

var user models.User

func TestMain(m *testing.M) {
	mur := repo_mock.NewMockUserRepository()
	us := svc.NewUserService(mur)
	user = models.User{
		Username: "bob",
		Password: "hello",
	}
	code := m.Run()
	os.Exit(code)
}
