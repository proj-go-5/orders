package main

import (
	"context"
	"fmt"
	"orders/internal/api"
	"orders/internal/config"
	"orders/internal/db"
	"orders/internal/repositories"
	"orders/internal/server"
	"orders/internal/services"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(config.Env("GIN_MODE"))

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Env("DB_HOST"),
		config.Env("DB_USER"),
		config.Env("DB_PASSWORD"),
		config.Env("DB_NAME"),
	)

	database := db.NewDatabase(dsn)
	conn, stop, err := database.GetConnection()
	defer stop()

	if err != nil {
		panic(err)
	}

	orderRepository := repositories.NewOrderRepository(conn)
	orderManager := services.NewOrderManager(orderRepository)

	var apis = []server.Routable{
		api.NewOrderAPI(orderManager),
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	router := gin.Default()
	srv := server.NewServer(router)
	srv.RegisterRoutes(apis)

	if err := srv.Start(ctx); err != nil {
		panic(err)
	}
}
