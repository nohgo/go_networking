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
	GetAll(string) ([]models.Car, error)
	AreValidCredentials(models.User) (bool, error)
}

type postgresUserRepository struct {
	pool *sql.DB
}

func NewUserRepository() *postgresUserRepository {
	return &postgresUserRepository{db.Pool}
}

func (ur *postgresUserRepository) Add(user models.User) error {
	_, err := ur.pool.Exec(fmt.Sprintf("INSERT INTO users (username, password) VALUES ('%v', '%v')", user.Username, user.Password))
	return err
}

func (ur *postgresUserRepository) GetAll(username string) ([]models.Car, error) {
	rows, err := ur.pool.Query(fmt.Sprintf("SELECT * FROM cars WHERE username='%v'", username))
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	cars := make([]models.Car, 0)
	for rows.Next() {
		var carMake string
		var model string
		var year int
		if err := rows.Scan(&carMake, &model, &year); err != nil {
			return nil, err
		}
		cars = append(cars, models.Car{Make: carMake, Model: model, Year: year})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cars, nil
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
