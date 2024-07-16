package svc_test

import (
	"os"
	"testing"

	db "github.com/nohgo/go_networking/api/database"
	"github.com/nohgo/go_networking/api/models"
)

var user models.User

func TestMain(m *testing.M) {
	db.CreatePool()
	defer db.ClosePool()
	user = models.User{
		Username: "bob",
		Password: "hello",
	}
	code := m.Run()
	os.Exit(code)
}
