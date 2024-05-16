package api

import (
	"context"
	"log"
	"net/http"
	"orders/internal/dto"
	"orders/internal/mapper"
	"orders/internal/models"

	"github.com/gin-gonic/gin"
)

type OrderManager interface {
	List(ctx context.Context) ([]models.Order, error)
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

	if err := api.manager.Create(ctx, order); err != nil {
		err := ctx.AbortWithError(http.StatusNotFound, err)
		if err != nil {
			log.Println("Error while aborting request:", err)
		}
		return
	}

	ctx.JSON(http.StatusCreated, &order)
}
