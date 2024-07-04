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
