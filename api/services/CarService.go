// The svc package provides all the services of the backend
package svc

import (
	repo "github.com/nohgo/go_networking/api/database/repositories"
	"github.com/nohgo/go_networking/api/models"
)

// The car service provides a way to execute commands that effect the car repository
// It contains one car repository that it calls
type CarService struct {
	cr repo.CarRepository
}

// NewCarService returns a new [svc.CarService]
// Calling it requires a [repo.CarRepository]
func NewCarService(cr repo.CarRepository) *CarService {
	return &CarService{cr: cr}
}

// Gets all cars associated with a username
func (cs *CarService) GetAll(username string) ([]models.Car, error) {
	return cs.cr.GetAll(username)
}

// Adds the provided car to the database
func (cs *CarService) Add(car models.Car, username string) error {
	return cs.cr.Add(car, username)
}

// Deletes the provided car from the database
func (cs *CarService) Delete(id int, username string) error {
	return cs.cr.Delete(id, username)
}
