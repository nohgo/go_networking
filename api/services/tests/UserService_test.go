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
	user = models.User{
		Username: "bob",
		Password: "hello",
	}
	code := m.Run()
	os.Exit(code)
}

func TestRegister(t *testing.T) {
	us := svc.NewUserService(repo_mock.NewMockUserRepository())
	if err := us.Register(user); err != nil {
		t.Fatalf("Register failed: %v", err)
	}
}

func TestLogin(t *testing.T) {
	us := svc.NewUserService(repo_mock.NewMockUserRepository())
	token, err := us.Login(user)

	if err != nil {
		t.Fatalf("Login returned an error: %v", err)
	}
	if token == "" {
		t.Fatalf("Login returned a blank token: %v", token)
	}
}
