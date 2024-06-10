package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	v1 "github.com/kalanihaubrick/todo-api/api/v1"
	"github.com/kalanihaubrick/todo-api/configs"
	"github.com/kalanihaubrick/todo-api/internal/database"
	"github.com/kalanihaubrick/todo-api/internal/models"
	"github.com/kalanihaubrick/todo-api/pkg/logger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file: ", err)
	}

	dsn := configs.GetDSN()
	err = database.Connect(dsn)

	if err != nil {
		logger.Fatal("failed to connect database: ", err)
	}

	err = models.Migrate(database.DB)

	if err != nil {
		logger.Fatal("failed to migrate database: ", err)
	}

	r := gin.Default()
	v1.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logger.Info("Starting server", "port", port)
	r.Run(":" + port)
}
