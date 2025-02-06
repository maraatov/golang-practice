package car

import (
	"context"
	models "first_practice_lesson/internal/app/models"
	"first_practice_lesson/internal/app/services/car/repository"
)

type Service interface {
	GetCars(context.Context) ([]*models.Car, error)
}

type service struct {
	carsRepo repository.Repository
}

// TODO Создать слайс и пройтись по машинам и взять их имена
func (s *service) GetCars(ctx context.Context) ([]*models.Car, error) {
	cars, err := s.carsRepo.GetCars(ctx)
	if err != nil {
		return nil, err
	}

	return cars, nil
}

func NewService(repo repository.Repository) Service {
	return &service{repo}
}
