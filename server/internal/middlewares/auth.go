package middlewares

import (
	"net/http"
	"strings"

	"github.com/adk-saugat/stash/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT token from Authorization header
func AuthMiddleware(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		ctx.Abort()
		return
	}

	// Extract token from "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
		ctx.Abort()
		return
	}

	token := parts[1]

	// Validate token
	claims, err := utils.ValidateToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		ctx.Abort()
		return
	}

	// Set user info in context for handlers to use
	ctx.Set("user_id", claims.UserID)
	ctx.Set("email", claims.Email)

	ctx.Next()
}
