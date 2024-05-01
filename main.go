package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create Gin router
	router := gin.Default()

	// Register Routes
	router.GET("/", homePage)

	// Start the server
	// Don't use localhost server address within a Docker container
	// If you do this only 'localhost'
	// (i.e. any service within the Docker container's network) can reach it.
	router.Run() // will by default use port 8080
}

func homePage(c *gin.Context) {
	c.String(http.StatusOK, "This is my home page")
}
