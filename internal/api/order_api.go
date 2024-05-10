package api

import (
	"net/http"
	"orders/internal/enums/status"
	"orders/internal/models"

	"github.com/gin-gonic/gin"
)

type OrderManager interface {
	List() []*models.Order
	Create(orderDTO *models.Order) error
}

func NewOrderAPI(manager OrderManager) *OrderAPI {
	return &OrderAPI{manager}
}

type OrderAPI struct {
	manager OrderManager
}

func (api *OrderAPI) RegisterRoutes(router *gin.Engine) {
	router.GET("/orders", api.listOrders)
	router.POST("/orders", api.createOrder)
	router.PATCH("/order/{id}/status", api.updateOrderStatus)
	router.GET("/order/{id}/history", api.getOrgerHistory)
}

func (api *OrderAPI) listOrders(c *gin.Context) {
	c.JSON(http.StatusOK, []models.Order{
		{ID: 1, Status: status.Active, CustomerID: 1, TotalPrice: 100},
	})
}

func (api *OrderAPI) createOrder(c *gin.Context) {
	c.Status(http.StatusCreated)
}

func (api *OrderAPI) updateOrderStatus(c *gin.Context) {
	c.Status(http.StatusCreated)
}

func (api *OrderAPI) getOrgerHistory(c *gin.Context) {
	c.Status(http.StatusCreated)
}
