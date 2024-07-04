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
