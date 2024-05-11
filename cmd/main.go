package main

import (
	"orders/internal/api"
	"orders/internal/config"
	"orders/internal/db"
	"orders/internal/repositories"
	"orders/internal/server"
	"orders/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(config.Env("GIN_MODE"))

	connection := db.GetConnection()

	orderRepository := repositories.NewOrderRepository(connection)
	orderProductRepository := repositories.NewOrderProductRepository(connection)
	orderManager := services.NewOrderManager(orderRepository, orderProductRepository)

	var apis = []server.Routable{
		api.NewOrderAPI(orderManager),
	}

	if err := runServer(apis); err != nil {
		panic(err)
	}
}

func runServer(apis []server.Routable) error {
	router := gin.Default()
	s := server.NewServer(router)
	s.RegisterRoutes(apis)

	return s.Start()
}
