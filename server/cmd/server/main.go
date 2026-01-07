package main

import (
	"log"
	"os"

	"github.com/adk-saugat/stash/server/internal/routes"
	"github.com/adk-saugat/stash/server/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// loading environment
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatal("Could not load environment.")
	}

	// initialize database connection pool
	database.InitializeDatabase()
	defer database.Close()

	// run migrations
	database.RunMigrations("../../migrations")

	// initializing server
	server := gin.Default()
	PORT := os.Getenv("PORT")

	routes.RegisterRoutes(server)

	err = server.Run(":" + PORT)
	if err != nil {
		log.Fatal("Could not run server on port " + PORT)
	}
}
