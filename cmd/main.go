package main

import (
	"context"
	"fmt"
	"github.com/proj-go-5/accounts/pkg/authorization"
	"net/http"
	"orders/internal/api"
	"orders/internal/api/middleware"
	"orders/internal/config"
	"orders/internal/db"
	"orders/internal/repositories"
	"orders/internal/server"
	"orders/internal/services/history"
	"orders/internal/services/order"
	"orders/internal/services/product"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
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
	orderHistoryRepository := repositories.NewOrderHistoryRepository(conn)

	client := product.NewClient(http.DefaultClient)
	productFetcher := product.NewFetcher(client)

	orderManager := order.NewOrderManager(orderRepository, orderHistoryRepository, productFetcher)
	historyManager := history.NewOrderHistoryManager(orderHistoryRepository)

	JwtExpiration, err := strconv.Atoi(config.Env("JWT_EXPIRATION"))
	if err != nil {
		panic(err)
	}
	jwtService := authorization.NewJwtService(config.Env("JWT_SECRET"), JwtExpiration)
	authService := authorization.NewAuthServie(jwtService)
	adminMiddlewareAdapter := middleware.NewAdminMiddlewareAdapter(authService)

	var apis = []server.Routable{
		api.NewOrderAPI(orderManager, historyManager, adminMiddlewareAdapter),
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
