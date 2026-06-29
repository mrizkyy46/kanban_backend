package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrizkyy46/kanban-backend/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Grouping rute API
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}
	}

	return r
}
