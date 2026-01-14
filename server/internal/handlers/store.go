package handlers

import (
	"net/http"

	"github.com/adk-saugat/stash/server/internal/models"
	"github.com/gin-gonic/gin"
)

type ShareRequest struct {
	ProjectId   string `json:"project_id"`
	ProjectName string `json:"project_name"`
	Store       models.Store `json:"store"`
}

func ShareStore(ctx *gin.Context) {
	var req ShareRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Get user info from JWT claims
	userId, _ := ctx.Get("user_id")
	userEmail, _ := ctx.Get("email")

	// Check if project exists, create if not
	project := models.Project{}
	if !project.Exists(req.ProjectId) {
		project.ProjectId = req.ProjectId
		project.Name = req.ProjectName
		project.OwnerId = userId.(string)

		if err := project.Create(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create project"})
			return
		}
	}

	// Set author from JWT and create store
	req.Store.Author = userEmail.(string)
	req.Store.ProjectId = req.ProjectId

	if err := req.Store.Create(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create store"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":    "Store shared successfully",
		"store_id":   req.Store.StoreId,
		"project_id": req.ProjectId,
	})
}
