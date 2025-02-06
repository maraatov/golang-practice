package handlers

import (
	"first_practice_lesson/internal/app"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	DI *app.DI
}

func AppRouter(router *gin.Engine, di *app.DI) {
	h := Handler{DI: di}

	api := router.Group("/api/v1")
	cars := api.Group("/cars") // http://localhost:8080/api/v1/cars

	cars.GET("/", h.getCars)
}
