package internal

import (
	"first_practice_lesson/internal/app"
	"first_practice_lesson/internal/app/db"
	"first_practice_lesson/internal/app/handlers"
	"first_practice_lesson/internal/app/services/car"
	"first_practice_lesson/internal/app/services/car/repository"
	"first_practice_lesson/internal/app/usecases"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Run() {
	dbPool, err := db.Connect("postgres://postgres:Aldiyar2004@localhost:5432/practices?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
		return
	}

	carsRepo := repository.NewRepository(dbPool)
	carsServices := car.NewService(carsRepo)
	carsUseCase := usecases.NewCarsUseCase(carsServices)

	router := gin.Default()
	di := &app.DI{CarsUseCase: carsUseCase}

	handlers.AppRouter(router, di)

	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
