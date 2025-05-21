package http

import (
	"ecommerce/order-service/internal/domain"
	"ecommerce/order-service/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	usecase usecase.OrderUsecase
}

func NewOrderHandler(router *gin.Engine, uc usecase.OrderUsecase) {
	handler := &OrderHandler{usecase: uc}

	router.POST("/orders", handler.CreateOrder)
	router.GET("/orders/:id", handler.GetOrder)
	router.PATCH("/orders/:id", handler.UpdateOrder)
	router.DELETE("/orders/:id", handler.DeleteOrder)
	router.GET("/orders", handler.ListOrders)
	router.GET("/users/:user_id/orders", handler.ListUserOrders)
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.usecase.Create(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := h.usecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var orderUpdates domain.Order
	if err := c.ShouldBindJSON(&orderUpdates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orderUpdates.ID = id
	if err := h.usecase.Update(&orderUpdates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderUpdates)
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	orders, err := h.usecase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) ListUserOrders(c *gin.Context) {
	userID := c.Param("user_id")
	orders, err := h.usecase.ListByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}
