package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine){
	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message" : "Server running okay!",
		})
	})
}