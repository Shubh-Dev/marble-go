package database

import (
	"log"
	"github.com/Marbleture/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDB() {
	dsn := "host = localhost user = postgres password = motorola123 dbname = marble port = 5432 sslmode = disable TimeZone = Asia/Shanghai"
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connec DB \n", err)
	}

	log.Println("DB connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running Migration")
	db.AutoMigrate(&models.Marble{})

	DB = Dbinstance{Db: db}
}