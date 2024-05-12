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
	router.PATCH("/order/:orderId/status", api.updateOrderStatus)
	router.GET("/order/:orderId/history", api.getOrgerHistory)
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

func (api *OrderAPI) updateOrderStatus(c *gin.Context) {
	orderId := c.Param("orderId")

	c.Status(http.StatusCreated)

	// userid := c.Param("userid")
	// message := "userid is " + userid
	// c.String(http.StatusOK, message)
	// fmt.Println(message)
}

func (api *OrderAPI) getOrgerHistory(c *gin.Context) {
	orderId := c.Param("orderId")

	c.Status(http.StatusCreated)
}
