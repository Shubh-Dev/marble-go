package database

import (
	"fmt"
	"log"

	"os"

	"github.com/Shubh-Dev/marble-go/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDB() {
	var runMigration bool = false
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	if databasePassword == "" {
		fmt.Println("Database password not set")
	} else {
		fmt.Println("Password retrived from env")
	}
	// dsn := "host = localhost user = postgres password = %v dbname = marble port = 5432 sslmode = disable TimeZone = Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		"localhost", "postgres", databasePassword, "marble", "5432", "disable", "Asia/Shanghai")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connec DB \n", err)
	}

	log.Println("DB connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	if runMigration {
		log.Println("running Migration")
		db.AutoMigrate(&models.Coupon{}, &models.CustomerAddress{}, &models.Customer{},
			&models.Discount{}, &models.OrderDiscount{}, &models.OrderItem{},
			&models.Order{}, &models.ProductDiscount{}, &models.ProductImage{}, &models.ProductStock{},
			&models.ProductVariation{}, &models.Product{}, &models.StockLocation{}, &models.Voucher{})
	}

	DB = Dbinstance{Db: db}
}
