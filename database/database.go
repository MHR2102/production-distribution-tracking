package database

import (
	"fmt"
	"log"

	"github.com/yourusername/production-distribution-tracking/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=prod_dist_tracking port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	fmt.Println("Database connected!")
}

func MigrateDB() {
	err := DB.AutoMigrate(&models.User{}, &models.ProductionBatch{}, &models.Distribution{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	fmt.Println("Database migrated!")
}
