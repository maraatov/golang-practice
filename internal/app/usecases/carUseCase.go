package usecases

import (
	"context"
	models "first_practice_lesson/internal/app/models"
	car "first_practice_lesson/internal/app/services/car"
)

type CarsUseCase interface {
	GetCars(ctx context.Context) ([]*models.Car, error)
}

type carUseCase struct {
	carsService car.Service
}

func (c *carUseCase) GetCars(ctx context.Context) ([]*models.Car, error) {
	cars, err := c.carsService.GetCars(ctx)
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func NewCarsUseCase(carsService car.Service) CarsUseCase {
	return &carUseCase{carsService}
}
