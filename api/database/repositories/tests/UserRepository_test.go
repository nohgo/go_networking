package repo_test

import (
	"os"
	"testing"

	db "github.com/nohgo/go_networking/api/database"
	"github.com/nohgo/go_networking/api/database/repositories"
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

func TestAdd(t *testing.T) {
	ur := repo.NewPostgresUserRepository()

	if err := ur.Add(user); err != nil {
		t.Fatalf("Add is broken: %v", err)
	}
}

func TestAreValidCredentials(t *testing.T) {
	ur := repo.NewPostgresUserRepository()

	exists, err := ur.AreValidCredentials(user)

	if err != nil {
		t.Fatalf("AreValidCredentials returned an error: %v", err)
	}

	if !exists {
		t.Fatalf("AreValidCredentials returned exists = %v", exists)
	}

}

func TestDelete(t *testing.T) {
	ur := repo.NewPostgresUserRepository()

	if err := ur.Delete(user); err != nil {
		t.Fatalf("Delete is broken: %v", err)
	}
}
