package main

import (
	"my-exchange/api"
	"my-exchange/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// --- 依賴注入 (Wiring) 開始 ---

	// 1. 先初始化 Service (只做一次)
	orderSvc := service.NewOrderService()

	// 2. 初始化 Handler，並把 Service 注入進去
	orderHandler := api.NewHandler(orderSvc)

	// --- Wiring 結束 ---

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// 3. 註冊路由
	// 注意這裡要用 orderHandler 實例的方法
	r.POST("/order", orderHandler.CreateOrder)

	r.Run(":8080")
}
