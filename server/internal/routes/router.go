package routes

import (
	"net/http"

	"github.com/adk-saugat/stash/server/internal/handlers"
	"github.com/adk-saugat/stash/server/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Server running okay!",
		})
	})

	// Auth routes
	server.POST("/api/login", handlers.LoginUser)
	server.POST("/api/register", handlers.RegisterUser)

	// Protected routes
	protected := server.Group("/api")
	protected.Use(middlewares.AuthMiddleware)
	{
		protected.POST("/share", handlers.ShareStore)
	}
}