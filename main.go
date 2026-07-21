package main

import (
	"log"
	"os"

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

	r.Run(os.Getenv("APP_HOST"))
}
