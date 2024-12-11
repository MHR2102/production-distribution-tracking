package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/MHR2102/production-distribution-tracking/models"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=prod_dist_tracking port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal menghubungkan ke database: ", err)
	}
	fmt.Println("Koneksi database berhasil!")
}

func MigrateDB() {
	err := DB.AutoMigrate(&models.User{}, &models.ProductionBatch{}, &models.Distribution{})
	if err != nil {
		log.Fatal("Gagal melakukan migrasi database: ", err)
	}
	fmt.Println("Migrasi database berhasil!")
}
