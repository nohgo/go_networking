package repo

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nohgo/go_networking/api/database"
	"github.com/nohgo/go_networking/api/models"
)

type CarRepository interface {
	GetAll(string) ([]models.Car, error)
	Add(models.Car, string) error
}

type postgresCarRepository struct {
	pool *sql.DB
}

func NewCarRepository() *postgresCarRepository {
	return &postgresCarRepository{pool: db.Pool}
}

func (cr *postgresCarRepository) GetAll(username string) ([]models.Car, error) {
	rows, err := cr.pool.Query(fmt.Sprintf("SELECT * FROM cars WHERE username='%v'", username))
	defer rows.Close()
	if err != nil {
		return nil, err
	}
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

func (cr *postgresCarRepository) Add(car models.Car, username string) error {
	_, err := cr.pool.Exec(fmt.Sprintf("INSERT INTO cars (make, model, year, username) VALUES('%v', '%v', '%v', '%v')", car.Make, car.Model, car.Year, username))
	return err
}
