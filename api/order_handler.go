package api

import (
	"my-exchange/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type OrderRequest struct {
	Symbol   string          `json:"symbol" binding:"required"`
	Price    decimal.Decimal `json:"price" binding:"required"`
	Quantity decimal.Decimal `json:"quantity" binding:"required"`
	Side     string          `json:"side" binding:"required,oneof=buy sell"`
}

// Handler holds the service dependencies
type Handler struct {
	orderService *service.OrderService
}

// NewHandler creates a new Handler with injected service
func NewHandler(s *service.OrderService) *Handler {
	return &Handler{
		orderService: s,
	}
}

// CreateOrder handles order creation requests
func (h *Handler) CreateOrder(c *gin.Context) {
	var req OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Price.LessThanOrEqual(decimal.Zero) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "price must be greater than 0"})
		return
	}

	if req.Quantity.LessThanOrEqual(decimal.Zero) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "quantity must be greater than 0"})
		return
	}

	if req.Side != "buy" && req.Side != "sell" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "side must be 'buy' or 'sell'"})
		return
	}

	total, fee, err := h.orderService.PlaceOrder(service.PlaceOrderParams{
		Symbol:   req.Symbol,
		Price:    req.Price,
		Quantity: req.Quantity,
		Side:     req.Side,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "order_created",
		"order":       req,
		"total_value": total,
		"fee":         fee,
	})
}
