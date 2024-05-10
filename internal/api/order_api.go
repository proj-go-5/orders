package api

import (
	"context"
	"fmt"
	"net/http"
	"orders/internal/models"

	"github.com/gin-gonic/gin"
)

type OrderManager interface {
	List(ctx context.Context) ([]*models.Order, error)
	Create(ctx context.Context, orderDTO *models.Order) error
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

func (api *OrderAPI) listOrders(ctx *gin.Context) {
	orders, err := api.manager.List(ctx)

	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, &orders)
}

func (api *OrderAPI) createOrder(ctx *gin.Context) {
	var order models.Order

	if err := ctx.BindJSON(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := api.manager.Create(ctx, &order); err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.Header("Location", fmt.Sprintf("/orders/%d", order.ID))
	ctx.JSON(http.StatusCreated, &order)
}
