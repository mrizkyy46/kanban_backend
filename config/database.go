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
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	cfg := postgres.Config{
		DSN: dsn,
	}

	database, err := gorm.Open(postgres.New(cfg), &gorm.Config{})

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
