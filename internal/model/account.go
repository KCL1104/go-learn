package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// User 使用者
type User struct {
	gorm.Model        // GORM 內建包含 ID, CreatedAt, UpdatedAt, DeletedAt
	Email      string `gorm:"uniqueIndex;type:varchar(100);not null"`
	Password   string `gorm:"not null"` // 實際存儲時必須 Hash 過，不能存明文！
}

// Account 資產帳戶
// 複合唯一索引：一個 User 對同一種 Currency 只能有一條記錄
type Account struct {
	gorm.Model
	UserID   uint   `gorm:"uniqueIndex:idx_user_currency"`                           // 外鍵
	Currency string `gorm:"uniqueIndex:idx_user_currency;type:varchar(10);not null"` // e.g., "BTC", "USDT"

	// 重點：使用 decimal 類型，並指定資料庫精度 (20位整數，8位小數)
	Balance decimal.Decimal `gorm:"type:decimal(20,8);not null;default:0"`

	// 凍結金額：下單時，錢不是消失，而是從 Balance 移動到 Locked
	Locked decimal.Decimal `gorm:"type:decimal(20,8);not null;default:0"`
}
