package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"orders/internal/api"
	"orders/internal/config"
)

func main() {
	gin.SetMode(config.Env("GIN_MODE"))

	router := gin.Default()
	api.RegisterRoutes(router)
	port := config.Env("PORT")
	err := router.Run(":" + port)
	if err != nil {
		fmt.Println(err)
	}
}
