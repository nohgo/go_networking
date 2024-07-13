package repo

import (
	"testing"

	db "github.com/nohgo/go_networking/api/database"
	"github.com/nohgo/go_networking/api/models"
)

func TestUserRepository(t *testing.T) {
	db.CreatePool()
	defer db.ClosePool()
	ur := NewUserRepository()
	user := models.User{
		Username: "bob",
		Password: "hello",
	}

	if err := ur.Add(user); err != nil {
		t.Fatalf("Add is broken: %v", err)
	}

	if exists, err := ur.AreValidCredentials(user); err != nil || !exists {
		t.Fatalf("AreValidCredentials is broken: %v", err)
	}

	if err := ur.Delete(user); err != nil {
		t.Fatalf("Delete is broken: %v", err)
	}
}
