package main

import (
	"my-exchange/api"
	"my-exchange/internal/infrastructure" // 引入
	"my-exchange/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 初始化資料庫
	infrastructure.InitDB()

	r := gin.Default()

	orderSvc := service.NewOrderService()
	orderHandler := api.NewHandler(orderSvc)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.POST("/order", orderHandler.CreateOrder)

	r.Run(":8080")
}
