package main

import (
	"my-exchange/api"
	"my-exchange/internal/infrastructure"
	"my-exchange/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	infrastructure.InitDB()

	db := infrastructure.DB

	r := gin.Default()

	orderSvc := service.NewOrderService(db)
	orderHandler := api.NewHandler(orderSvc)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.POST("/order", orderHandler.CreateOrder)

	r.Run(":8080")
}
