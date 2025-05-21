package http

import (
	"ecommerce/car-service/internal/domain"
	"ecommerce/car-service/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CarHandler struct {
	usecase usecase.CarUsecase
}

func NewCarHandler(router *gin.Engine, uc usecase.CarUsecase) {
	handler := &CarHandler{usecase: uc}

	router.POST("/cars", handler.CreateCar)
	router.GET("/cars/:id", handler.GetCar)
	router.PATCH("/cars/:id", handler.UpdateCar)
	router.DELETE("/cars/:id", handler.DeleteCar)
	router.GET("/cars", handler.ListCars)
}

func (h *CarHandler) CreateCar(c *gin.Context) {
	var car domain.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.usecase.Create(&car); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, car)
}

func (h *CarHandler) GetCar(c *gin.Context) {
	id := c.Param("id")
	car, err := h.usecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}
	c.JSON(http.StatusOK, car)
}

func (h *CarHandler) UpdateCar(c *gin.Context) {
	id := c.Param("id")
	var carUpdates domain.Car
	if err := c.ShouldBindJSON(&carUpdates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	carUpdates.ID = id
	if err := h.usecase.Update(&carUpdates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, carUpdates)
}

func (h *CarHandler) DeleteCar(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *CarHandler) ListCars(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}
	cars, err := h.usecase.List(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"page":  page,
		"limit": limit,
		"cars":  cars,
	})
}
