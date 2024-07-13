package repo

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"github.com/nohgo/go_networking/api/database"
	"github.com/nohgo/go_networking/api/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Add(models.User) error
	AreValidCredentials(models.User) (bool, error)
	Delete(user models.User) error
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
	_, err = ur.pool.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, string(hashedPassword))
	return err
}

func (ur *postgresUserRepository) AreValidCredentials(user models.User) (bool, error) {
	row := ur.pool.QueryRow("SELECT password FROM users WHERE username = $1;", user.Username)

	var foundPassword string
	err := row.Scan(&foundPassword)
	if err != nil {
		return false, errors.New("Invalid credentials")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(foundPassword), []byte(user.Password)); err != nil {
		return false, nil
	}

	return true, nil
}

func (ur *postgresUserRepository) Delete(user models.User) error {
	_, err := ur.pool.Exec("DELETE FROM users WHERE username=$1", user.Username)
	return err
}
