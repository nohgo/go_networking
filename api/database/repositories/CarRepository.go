package repo

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/nohgo/go_networking/api/database"
	"github.com/nohgo/go_networking/api/models"
)

// The CarRepository interface provides a way to interact with the car database
type CarRepository interface {
	GetAll(string) ([]models.Car, error)
	Add(models.Car, string) error
	Delete(int, string) error
}

// postgresCarRepository is a concrete implementation of [repo.CarRepository]
type postgresCarRepository struct {
	pool *sql.DB
}

// Returns a new [repo.postgresCarRepository]
func NewCarRepository() *postgresCarRepository {
	return &postgresCarRepository{pool: db.Pool}
}

// Gets all cars associated with a given username
func (cr *postgresCarRepository) GetAll(username string) ([]models.Car, error) {
	rows, err := cr.pool.Query("SELECT * FROM cars WHERE username=$1", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cars := make([]models.Car, 0)
	for rows.Next() {
		var car models.Car
		var _username string
		if err := rows.Scan(&car.Id, &car.Make, &car.Model, &car.Year, &_username); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cars, nil
}

// Adds a car to the database
func (cr *postgresCarRepository) Add(car models.Car, username string) error {
	_, err := cr.pool.Exec("INSERT INTO cars (make, model, year, username) VALUES($1, $2, $3, $4)", car.Make, car.Model, car.Year, username)
	return err
}

// Deletes a car from the database
func (cr *postgresCarRepository) Delete(id int, username string) error {
	_, err := cr.pool.Exec("DELETE FROM cars WHERE id=$1", id)
	return err
}
