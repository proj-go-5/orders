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
	"github.com/proj-go-5/accounts/pkg/authorization"
)

type OrderManager interface {
	List(ctx context.Context) ([]models.Order, error)
	Create(ctx context.Context, orderDTO *models.Order) error
	UpdateStatus(ctx context.Context, orderID int, newStatus status.Status) error
}

type OrderHistoryManager interface {
	List(ctx context.Context, orderID int) ([]models.OrderHistory, error)
}

func NewOrderAPI(orderManager OrderManager, orderHistoryManager OrderHistoryManager) *OrderAPI {
	return &OrderAPI{orderManager, orderHistoryManager}
}

type OrderAPI struct {
	orderManager        OrderManager
	orderHistoryManager OrderHistoryManager
}

func (api *OrderAPI) RegisterRoutes(router *gin.Engine) {
	jwtService := authorization.NewJwtService("test", 100)
	authService := authorization.NewAuthServie(jwtService)

	router.POST("/orders", api.createOrder)
	router.GET("/orders", gin.WrapF(authService.AdminMiddleware(func(w http.ResponseWriter, r *http.Request) {
		ctx := w.(*responseRecorder).Context
		api.listOrders(ctx)
	})))

	router.GET("/order/:orderID/history", api.getOrderHistory)
	router.PATCH("/order/:orderID/status", gin.WrapF(authService.AdminMiddleware(func(w http.ResponseWriter, r *http.Request) {
		ctx := w.(*responseRecorder).Context
		api.updateOrderStatus(ctx)
	})))

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

	err = api.orderManager.UpdateStatus(ctx, orderID, statusDTO.Status)
	if err != nil {
		err := ctx.AbortWithError(http.StatusBadRequest, err)
		if err != nil {
			log.Println("Error while aborting request:", err)
		}
		return
	}
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

type responseRecorder struct {
	gin.ResponseWriter
	Context    *gin.Context
	statusCode int
}

func (r *responseRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}
