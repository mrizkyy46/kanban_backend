package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mrizkyy46/kanban-backend/config"
	"github.com/mrizkyy46/kanban-backend/routes"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Gagal memuat file .env")
	}

	config.ConnectDatabase()

	r := routes.SetupRouter()

	r.Run("127.0.0.1:8080")
}
