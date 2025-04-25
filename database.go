package config

import (
	"crud-app/models"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	// Konfigurasi khusus untuk MAMP
	config := map[string]string{
		"DB_USER":     "root",      // Username default MAMP
		"DB_PASSWORD": "root",      // Password default MAMP
		"DB_HOST":     "localhost", // Host
		"DB_PORT":     "8889",      // Port MySQL MAMP
		"DB_NAME":     "crud_db",   // Nama database
	}

	// Set environment variables jika belum ada
	for key, value := range config {
		if os.Getenv(key) == "" {
			os.Setenv(key, value)
		}
	}

	// Format DSN untuk MAMP
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	var db *gorm.DB
	var err error

	// Retry connection
	for i := 0; i < 5; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Printf("Connection attempt %d failed. Retrying in 5 seconds...\n", i+1)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Set connection pool
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Auto migrate
	if err := db.AutoMigrate(&models.Category{}, &models.Product{}); err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	DB = db
	fmt.Println("🎉 Connected to database successfully!")
	return nil
}
