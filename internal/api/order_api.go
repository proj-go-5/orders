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
}

func (api *OrderAPI) listOrders(c *gin.Context) {
	c.JSON(http.StatusOK, []models.Order{
		{ID: 1, Status: status.Actived, CustomerID: 1, TotalPrice: 100},
	})
}

func (api *OrderAPI) createOrder(c *gin.Context) {
	c.Status(http.StatusCreated)
}
