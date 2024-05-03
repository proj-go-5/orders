package main

import (
	"github.com/gin-gonic/gin"
	"orders/internal/api"
	"orders/internal/config"
	"orders/internal/db"
	"orders/internal/repositories"
	"orders/internal/server"
	"orders/internal/services"
)

func main() {
	gin.SetMode(config.Env("GIN_MODE"))

	connection := db.GetConnection()

	orderRepository := repositories.NewOrderRepository(connection)
	orderManager := services.NewOrderManager(orderRepository)

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
