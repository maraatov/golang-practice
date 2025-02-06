package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getCars(ctx *gin.Context) {
	res, err := h.DI.CarsUseCase.GetCars(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}
