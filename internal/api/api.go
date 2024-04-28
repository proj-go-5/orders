package api

import (
	"github.com/gin-gonic/gin"
	"orders/internal/observers"
	"orders/internal/repositories"
	"orders/internal/services"
)

type Routable interface {
	RegisterRoutes(router *gin.Engine)
}

func RegisterRoutes(router *gin.Engine) {
	orderObservers := []services.OrderObserver{
		observers.NewEmailObserver(),
	}
	orderRepository := repositories.NewOrderRepository()
	orderManager := services.NewOrderManager(orderRepository, orderObservers)

	var apis = []Routable{
		NewOrderAPI(orderManager),
	}

	for _, a := range apis {
		a.RegisterRoutes(router)
	}
}
