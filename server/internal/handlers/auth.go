package handlers

import (
	"net/http"

	"github.com/adk-saugat/stash/server/internal/models"
	"github.com/adk-saugat/stash/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginUser(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if user exists
	var user models.User
	err := user.FindByEmail(req.Email)

	if err != nil {
		// User doesn't exist
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// User exists - verify password
	if !utils.CheckPassword(req.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logged in successfully",
		"email":   user.Email,
	})
}

func RegisterUser(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if user already exists
	var existingUser models.User
	err := existingUser.FindByEmail(req.Email)
	if err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not process password"})
		return
	}

	// Create new user
	newUser := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := newUser.Create(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Account created successfully",
		"email":   req.Email,
	})
}