package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"orders/internal/api"
)

func main() {
	gin.SetMode(gin.DebugMode) // gin.ReleaseMode for production

	router := gin.Default()
	api.RegisterRoutes(router)

	err := router.Run(":8082")
	if err != nil {
		fmt.Println(err)
	}
}
