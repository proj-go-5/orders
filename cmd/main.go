package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"orders/internal/api"
	"orders/internal/config"
	"orders/internal/db"
	"orders/internal/repositories"
	"orders/internal/server"
	"orders/internal/services"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	gin.SetMode(config.Env("GIN_MODE"))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Env("DB_HOST"),
		config.Env("DB_USER"),
		config.Env("DB_PASSWORD"),
		config.Env("DB_NAME"),
	)

	database := db.NewDatabase(dsn)
	connection, err := database.GetConnection(ctx)
	if err != nil {
		panic(err)
	}

	orderRepository := repositories.NewOrderRepository(connection)
	orderManager := services.NewOrderManager(orderRepository)

	var apis = []server.Routable{
		api.NewOrderAPI(orderManager),
	}

	router := gin.Default()
	newServer := server.NewServer(router)
	newServer.RegisterRoutes(apis)
	if err := newServer.Start(ctx); err != nil {
		panic(err)
	}
}
