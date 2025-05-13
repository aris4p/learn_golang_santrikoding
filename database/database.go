package database

import (
	"fmt"
	"log"

	"github.com/aris4p/config"
	"github.com/aris4p/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Load konfigurasi database dari .env
	dbHost := config.GetEnv("DB_HOST", "localhost")
	dbPort := config.GetEnv("DB_PORT", "3306")
	dbUser := config.GetEnv("DB_USER", "root")
	dbPass := config.GetEnv("DB_PASSWORD", "")
	dbName := config.GetEnv("DB_NAME", "db_golang")

	// Format DSN (Data Source Name) untuk MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	// Membuat koneksi ke database
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connected successfully!!")

	// **Auto Migrate Models**
	err = DB.AutoMigrate(&models.User{}) // Bisa tamabahkan model lain jika perlu
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("Database Migrated Succesfully!")
}
