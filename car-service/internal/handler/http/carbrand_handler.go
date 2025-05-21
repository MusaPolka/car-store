package http

import (
	"ecommerce/car-service/internal/domain"
	"ecommerce/car-service/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CarBrandHandler struct {
	usecase usecase.CarBrandUsecase
}

func NewCarBrandHandler(router *gin.Engine, uc usecase.CarBrandUsecase) {
	handler := &CarBrandHandler{usecase: uc}

	router.POST("/carbrands", handler.CreateCarBrand)
	router.GET("/carbrands/:id", handler.GetCarBrand)
	router.PATCH("/carbrands/:id", handler.UpdateCarBrand)
	router.DELETE("/carbrands/:id", handler.DeleteCarBrand)
	router.GET("/carbrands", handler.ListCarBrands)
}

func (h *CarBrandHandler) CreateCarBrand(c *gin.Context) {
	var brand domain.CarBrand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.usecase.Create(&brand); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, brand)
}

func (h *CarBrandHandler) GetCarBrand(c *gin.Context) {
	id := c.Param("id")
	brand, err := h.usecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car brand not found"})
		return
	}
	c.JSON(http.StatusOK, brand)
}

func (h *CarBrandHandler) UpdateCarBrand(c *gin.Context) {
	id := c.Param("id")
	var updates domain.CarBrand
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updates.ID = id
	if err := h.usecase.Update(&updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updates)
}

func (h *CarBrandHandler) DeleteCarBrand(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *CarBrandHandler) ListCarBrands(c *gin.Context) {
	brands, err := h.usecase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brands)
}
