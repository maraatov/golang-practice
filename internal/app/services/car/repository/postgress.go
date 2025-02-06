package repository

import (
	"context"
	car "first_practice_lesson/internal/app/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

// TODO Реализовать селект из базы
func (r *repository) GetCars(ctx context.Context) ([]*car.Car, error) {
	rows, err := r.db.Query(ctx, "SELECT brand, year, model FROM cars")
	if err != nil {
		return nil, err
	}

	var cars []*car.Car

	for rows.Next() {
		car := new(car.Car)
		err := rows.Scan(&car.Brand, &car.Year, &car.Model)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	return cars, nil
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}
