package infrastructure

import (
	"fmt"
	"log"
	"my-exchange/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global database instance
var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=password dbname=exchange port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("🚀 Database connected successfully!")

	// AutoMigrate automatically creates or updates table schemas based on structs
	err = DB.AutoMigrate(&model.User{}, &model.Account{}, &model.Order{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("✅ Database migration completed!")
}
