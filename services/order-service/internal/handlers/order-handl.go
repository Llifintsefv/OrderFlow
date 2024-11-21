package handlers

import (
	"net/http"
	"order-service/internal/models"
	"order-service/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)


type OrderHandlers struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *OrderHandlers{
	return &OrderHandlers{orderService: orderService}
}

func (h *OrderHandlers) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.BindJSON(&order);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	if err := h.orderService.CreateOrder(&order); err != nil {
		c.JSON(http.StatusBadGateway,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusCreated,order)
}

func (h *OrderHandlers) GetOrder(c *gin.Context) {
	idStr := c.Param("id")
	id,err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	order,err := h.orderService.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,order)
	
}