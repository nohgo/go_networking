// The repo package contains all the repositories for the api
package repo

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"github.com/nohgo/go_networking/api/database"
	"github.com/nohgo/go_networking/api/models"
	"golang.org/x/crypto/bcrypt"
)

// The user repository interface provides simple methods to interact with the user database
type UserRepository interface {
	Add(models.User) error
	AreValidCredentials(models.User) (bool, error)
	Delete(user models.User) error
}

// postgresUserRepository is a concrete implementation of the [repo.UserRepository] interface using postgres SQL
type postgresUserRepository struct {
	pool *sql.DB
}

// Returns a [repo.postgresUserRepository] using the pool variable from [db.Pool]
func NewPostgresUserRepository() *postgresUserRepository {
	return &postgresUserRepository{db.Pool}
}

// Adds a user to the database
func (ur *postgresUserRepository) Add(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	_, err = ur.pool.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, string(hashedPassword))
	return err
}

// Verifies a user's credentials are valid
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

// deletes a user from the database
func (ur *postgresUserRepository) Delete(user models.User) error {
	_, err := ur.pool.Exec("DELETE FROM users WHERE username=$1", user.Username)
	return err
}
