package db

import (
	"BACKEND-GOLANG-MVVM/internal/app/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Initialize() (*gorm.DB, error) {
	// Dapatkan URL database dari environment variable
	dsn := os.Getenv("DATABASE_URL")
	log.Println("Database:", dsn)

	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set in environment")
	}

	// Membuka koneksi ke SQL Server menggunakan DSN langsung
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate the User model
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	return db, nil
}
