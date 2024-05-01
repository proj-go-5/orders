package api

import (
	"github.com/gin-gonic/gin"
	"orders/internal/db"
	"orders/internal/repositories"
	"orders/internal/services"
)

type Routable interface {
	RegisterRoutes(router *gin.Engine)
}

func RegisterRoutes(router *gin.Engine) {
	connection := db.GetConnection()

	orderRepository := repositories.NewOrderRepository(connection)
	orderManager := services.NewOrderManager(orderRepository)

	var apis = []Routable{
		NewOrderAPI(orderManager),
	}

	for _, a := range apis {
		a.RegisterRoutes(router)
	}
}
