package repo

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nohgo/go_networking/api/database"
	"github.com/nohgo/go_networking/api/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Add(models.User) error
	AreValidCredentials(models.User) (bool, error)
}

type postgresUserRepository struct {
	pool *sql.DB
}

func NewUserRepository() *postgresUserRepository {
	return &postgresUserRepository{db.Pool}
}

func (ur *postgresUserRepository) Add(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	_, err = ur.pool.Exec(fmt.Sprintf("INSERT INTO users (username, password) VALUES ('%v', '%v')", user.Username, string(hashedPassword)))
	return err
}

func (ur *postgresUserRepository) AreValidCredentials(user models.User) (bool, error) {
	row := ur.pool.QueryRow(fmt.Sprintf("SELECT password FROM users WHERE username = '%v';", user.Username))

	var foundPassword string
	err := row.Scan(&foundPassword)
	if err != nil {
		return false, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(foundPassword), []byte(user.Password)); err != nil {
		return false, nil
	}

	return true, nil
}
