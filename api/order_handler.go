package api

import (
	"my-exchange/internal/service" // 引入 service 包
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

// Request Body 保持不變
type OrderRequest struct {
	Symbol   string          `json:"symbol" binding:"required"`
	Price    decimal.Decimal `json:"price" binding:"required,gt=0"`
	Quantity decimal.Decimal `json:"quantity" binding:"required,gt=0"`
	Side     string          `json:"side" binding:"required,oneof=buy sell"`
}

// 2. 定義 Handler 結構體
// 這個結構體持有 Service 的指針
type Handler struct {
	orderService *service.OrderService
}

// 3. 構造函數 (Constructor)
// 在外面把 Service 傳進來 (這就是注入！)
func NewHandler(s *service.OrderService) *Handler {
	return &Handler{
		orderService: s,
	}
}

// 4. 將 CreateOrder 變成 Handler 的方法 (Method)
// 注意這裡變成了 (h *Handler)
func (h *Handler) CreateOrder(c *gin.Context) {
	var req OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 5. 使用結構體中已經注入的 service，而不是自己 New
	total, fee := h.orderService.PlaceOrder(service.PlaceOrderParams{
		Symbol:   req.Symbol,
		Price:    req.Price,
		Quantity: req.Quantity,
		Side:     req.Side,
	})

	c.JSON(http.StatusOK, gin.H{
		"status":      "order_created",
		"order":       req,
		"total_value": total,
		"fee":         fee,
	})
}
