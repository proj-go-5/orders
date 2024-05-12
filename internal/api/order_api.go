package api

import (
	"context"
	"log"
	"net/http"
	"orders/internal/dto"
	"orders/internal/enums/status"
	"orders/internal/mapper"
	"orders/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderManager interface {
	List(ctx context.Context) ([]models.Order, error)
	Create(ctx context.Context, orderDTO *models.Order) error
	UpdateStatusByOrderId(ctx context.Context, orderId int, newStatus status.Status) error
	GetHistoryByOrderId(ctx context.Context, orderID int) ([]models.OrderHistory, error)
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
	router.PATCH("/order/:orderID/status", api.updateOrderStatus)
	router.GET("/order/:orderID/history", api.getOrgerHistory)
}

func (api *OrderAPI) listOrders(ctx *gin.Context) {
	orders, err := api.manager.List(ctx)

	if err != nil {
		err := ctx.AbortWithError(http.StatusNotFound, err)
		if err != nil {
			log.Println("Error while aborting request:", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, &orders)
}

func (api *OrderAPI) createOrder(ctx *gin.Context) {
	var orderDTO dto.Order

	if err := ctx.BindJSON(&orderDTO); err != nil {
		err := ctx.AbortWithError(http.StatusBadRequest, err)
		if err != nil {
			log.Println("Error while aborting request:", err)
		}
		return
	}

	var order = mapper.ConvertOrderDTOToModel(&orderDTO)

	if err := api.manager.Create(ctx, order); err != nil {
		err := ctx.AbortWithError(http.StatusNotFound, err)
		if err != nil {
			log.Println("Error while aborting request:", err)
		}
		return
	}

	ctx.JSON(http.StatusCreated, &order)
}

func (api *OrderAPI) updateOrderStatus(ctx *gin.Context) {
	// if err := api.manager.UpdateStatusByOrderId()
}

func (api *OrderAPI) getOrgerHistory(ctx *gin.Context) {
	orderIDstr := ctx.Param("orderID")
	orderID, err := strconv.Atoi(orderIDstr)
	if err != nil {
		return
	}

	history, err := api.manager.GetHistoryByOrderId(ctx, orderID)
	if err != nil {
		err := ctx.AbortWithError(http.StatusNotFound, err)
		if err != nil {
			log.Println("Error while aborting request:", err)
		}
		return
	}

	ctx.JSON(http.StatusCreated, history)
}
