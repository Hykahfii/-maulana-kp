package main

import (
	"crud-app/config"
	"crud-app/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.ConnectDatabase(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := gin.Default()

	// Categories endpoints
	r.GET("/categories", controllers.GetCategories)
	r.POST("/categories", controllers.CreateCategory)

	// Products endpoints
	r.GET("/products", controllers.GetProducts)
	r.POST("/products", controllers.CreateProduct)

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Start server
	if err := r.Run(":8082"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
