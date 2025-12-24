package service

import (
	"github.com/shopspring/decimal"
)

// 1. 定義 OrderService 結構體
// 目前它是空的，未來這裡會放 Database 的連線實例 (Repository)
type OrderService struct{}

// 2. 為了方便，我們創建一個 New 函數 (Constructor pattern)
func NewOrderService() *OrderService {
	return &OrderService{}
}

// 3. 定義入參結構 (DTO)
// 這是 Service 層看得懂的參數，解耦 HTTP 層
type PlaceOrderParams struct {
	Symbol   string
	Price    decimal.Decimal
	Quantity decimal.Decimal
	Side     string
}

// 4. 實現業務邏輯方法
// 注意：這個方法掛載在 *OrderService 上
func (s *OrderService) PlaceOrder(params PlaceOrderParams) (decimal.Decimal, decimal.Decimal) {
	// 邏輯封裝在這裡
	grossValue := params.Price.Mul(params.Quantity)
	fee := grossValue.Mul(decimal.NewFromFloat(0.001))

	// 簡單區分買賣邏輯
	var total decimal.Decimal
	if params.Side == "buy" {
		total = grossValue.Add(fee)
	} else {
		total = grossValue.Sub(fee) // 賣出的話，拿到的是扣除手續費後的錢
	}

	return total, fee
}
