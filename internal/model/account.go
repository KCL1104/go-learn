package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;type:varchar(100);not null"`
	Password string `gorm:"not null"` // Must be hashed, never store plaintext
}

// Account represents a user's asset account (unique per user per currency)
type Account struct {
	gorm.Model
	UserID   uint   `gorm:"uniqueIndex:idx_user_currency"`
	Currency string `gorm:"uniqueIndex:idx_user_currency;type:varchar(10);not null"` // e.g., "BTC", "USDT"

	Balance decimal.Decimal `gorm:"type:decimal(20,8);not null;default:0"`
	Locked  decimal.Decimal `gorm:"type:decimal(20,8);not null;default:0"` // Funds moved from Balance when placing orders
}
