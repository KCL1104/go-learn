package service

import (
	"my-exchange/internal/model"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// OrderService handles order-related business logic
type OrderService struct {
	db *gorm.DB
}

// NewOrderService creates a new OrderService with DB dependency
func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{
		db: db,
	}
}

// PlaceOrderParams is the DTO for order placement
type PlaceOrderParams struct {
	Symbol   string
	Price    decimal.Decimal
	Quantity decimal.Decimal
	Side     string
}

func (s *OrderService) PlaceOrder(params PlaceOrderParams) (decimal.Decimal, decimal.Decimal, error) {
	grossValue := params.Price.Mul(params.Quantity)
	fee := grossValue.Mul(decimal.NewFromFloat(0.001))

	var total decimal.Decimal
	if params.Side == "buy" {
		total = grossValue.Add(fee)
	} else {
		total = grossValue.Sub(fee)
	}

	order := model.Order{
		UserID:   1, // TODO: hardcoded, implement auth later
		Symbol:   params.Symbol,
		Side:     model.OrderSide(params.Side),
		Price:    params.Price,
		Quantity: params.Quantity,
		Status:   model.StatusPending,
	}

	if err := s.db.Create(&order).Error; err != nil {
		return decimal.Zero, decimal.Zero, err
	}

	return total, fee, nil
}
