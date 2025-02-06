package repository

import (
	"context"
	car "first_practice_lesson/internal/app/models"
)

type Repository interface {
	GetCars(context.Context) ([]*car.Car, error)
}
