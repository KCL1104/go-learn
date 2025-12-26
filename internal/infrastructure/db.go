package infrastructure

import (
	"fmt"
	"log"
	"my-exchange/internal/model" // 引入剛剛定義的 model

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 全局 DB 變數 (在真實專案中，我們會把它放在 Config struct 裡傳遞，這裡先簡化)
var DB *gorm.DB

func InitDB() {
	// 這裡請填入你本地 PostgreSQL 的資訊
	// dsn = Data Source Name
	dsn := "host=localhost user=postgres password=aa124578 dbname=exchange port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("🚀 Database connected successfully!")

	// AutoMigrate 會自動根據 Struct 建立或更新資料庫表結構
	// 這是 GORM 最方便的功能之一
	err = DB.AutoMigrate(&model.User{}, &model.Account{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("✅ Database migration completed!")
}
