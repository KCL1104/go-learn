package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type OrderType string
type OrderSide string
type OrderStatus string

const (
	SideBuy  OrderSide = "buy"
	SideSell OrderSide = "sell"

	StatusPending   OrderStatus = "pending"
	StatusFilled    OrderStatus = "filled"
	StatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	gorm.Model

	UserID uint `gorm:"not null"`

	Symbol string    `gorm:"type:varchar(20);not null"` // e.g., "BTC/USDT"
	Side   OrderSide `gorm:"type:varchar(10);not null"` // "buy" or "sell"

	Price    decimal.Decimal `gorm:"type:decimal(20,8);not null"`
	Quantity decimal.Decimal `gorm:"type:decimal(20,8);not null"`

	Status OrderStatus `gorm:"type:varchar(20);not null;default:'pending'"`
}
