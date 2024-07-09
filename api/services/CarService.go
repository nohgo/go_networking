package svc

import (
	repo "github.com/nohgo/go_networking/api/database/repositories"
	"github.com/nohgo/go_networking/api/models"
)

type CarService struct {
	cr repo.CarRepository
}

func NewCarService(cr repo.CarRepository) *CarService {
	return &CarService{cr: cr}
}

func (cs *CarService) GetAll(username string) ([]models.Car, error) {
	return cs.cr.GetAll(username)
}

func (cs *CarService) Add(car models.Car, username string) error {
	return cs.cr.Add(car, username)
}

func (cs *CarService) Delete(id int, username string) error {
	return cs.cr.Delete(id, username)
}
