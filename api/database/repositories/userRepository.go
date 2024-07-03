package repo

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nohgo/go_networking/api/database"
	"github.com/nohgo/go_networking/api/models"
)

type UserRepository interface {
	Add(models.User) error
	GetAll() ([]models.User, error)
	AreValidCredentials(models.User) (bool, error)
}

type postgresUserRepository struct {
	pool *sql.DB
}

func NewUserRepository() *postgresUserRepository {
	return &postgresUserRepository{db.Pool}
}

func (ur *postgresUserRepository) Add(user models.User) error {
	_, err := ur.pool.Exec(fmt.Sprintf("INSERT INTO users (username, password, cars) VALUES ('%v', '%v')", user.Username, user.Password))
	return err
}

func (ur *postgresUserRepository) GetAll() ([]models.User, error) {
	rows, err := ur.pool.Query("SELECT * FROM users")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	users := make([]models.User, 0)
	for rows.Next() {
		var username string
		var password string
		if err := rows.Scan(&username, &password); err != nil {
			return nil, err
		}
		users = append(users, models.User{Username: username, Password: password})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *postgresUserRepository) AreValidCredentials(user models.User) (bool, error) {
	row := ur.pool.QueryRow(fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM users WHERE username = '%v' AND password = '%v');", user.Username, user.Password))

	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
