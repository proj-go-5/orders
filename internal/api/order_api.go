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
	UpdateStatus(ctx context.Context, orderID int, newStatus status.Status, comment string) error
}

type OrderHistoryManager interface {
	List(ctx context.Context, orderID int) ([]models.OrderHistory, error)
}

type AdminMiddleware interface {
	Handler() gin.HandlerFunc
}

func NewOrderAPI(
	orderManager OrderManager,
	orderHistoryManager OrderHistoryManager,
	adminMiddleware AdminMiddleware,
) *OrderAPI {
	return &OrderAPI{orderManager, orderHistoryManager, adminMiddleware}
}

type OrderAPI struct {
	orderManager        OrderManager
	orderHistoryManager OrderHistoryManager
	adminMiddleware     AdminMiddleware
}

func (api *OrderAPI) RegisterRoutes(router *gin.Engine) {
	adminRoutes := router.Group("/")
	adminRoutes.Use(api.adminMiddleware.Handler())
	{
		adminRoutes.GET("/orders", api.listOrders)
		adminRoutes.PATCH("/orders/:orderID/status", api.updateOrderStatus)
	}

	adminRoutes.GET("/orders/:orderID/history", api.getOrderHistory)
	router.POST("/orders", api.createOrder)
}

func (api *OrderAPI) listOrders(ctx *gin.Context) {
	orders, err := api.orderManager.List(ctx)

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

	if validationError := orderDTO.Validate(); validationError != nil {
		err := ctx.AbortWithError(http.StatusBadRequest, validationError)
		if err != nil {
			log.Println("Error while aborting request:", err)
		}
		return
	}

	var order = mapper.ConvertOrderDTOToModel(&orderDTO)

	if err := api.orderManager.Create(ctx, order); err != nil {
		err := ctx.AbortWithError(http.StatusNotFound, err)
		if err != nil {
			log.Println("Error while aborting request:", err)
		}
		return
	}

	ctx.JSON(http.StatusCreated, &order)
}

func (api *OrderAPI) updateOrderStatus(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("orderID"))
	if err != nil {
		return
	}

	var statusDTO dto.OrderStatus
	if err := ctx.BindJSON(&statusDTO); err != nil {
		err := ctx.AbortWithError(http.StatusBadRequest, err)
		if err != nil {
			log.Println("Error while aborting request:", err)
		}
		return
	}

	err = api.orderManager.UpdateStatus(ctx, orderID, statusDTO.Status, statusDTO.Comment)
	if err != nil {
		err := ctx.AbortWithError(http.StatusBadRequest, err)
		if err != nil {
			log.Println("Error while aborting request:", err)
		}
		return
	}

	return
}

func (api *OrderAPI) getOrderHistory(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("orderID"))
	if err != nil {
		return
	}

	history, err := api.orderHistoryManager.List(ctx, orderID)
	if err != nil {
		err := ctx.AbortWithError(http.StatusNotFound, err)
		if err != nil {
			log.Println("Error while aborting request:", err)
		}
		return
	}

	ctx.JSON(http.StatusOK, history)
}
