package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"orders/configs"
	"orders/internal/api"
)

func main() {
	gin.SetMode(configs.Env("GIN_MODE"))

	router := gin.Default()
	api.RegisterRoutes(router)

	port := configs.Env("PORT")
	fmt.Println(port)
	err := router.Run(":" + port)
	if err != nil {
		fmt.Println(err)
	}
}
