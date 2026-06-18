package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrizkyy46/kanban-backend/config"
)

func main() {
	config.ConnectDatabase()
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Backend terkoneksi dengan database dan siap digunakan!",
		})
	})

	r.Run("127.0.0.1:8080")
}