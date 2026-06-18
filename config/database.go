package config

import (
	"fmt"
	"log"
	"os"

	"github.com/mrizkyy46/kanban-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbName, port)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	// Sinkronisasi otomatis (Membuat tabel Users & Tasks jika belum ada di database)
	err = database.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		log.Fatal("Gagal melakukan Auto Migrate:", err)
	}

	fmt.Println("Koneksi Database & Auto Migrate Berhasil!")
	DB = database
}
