package main

import (
	"github.com/gin-gonic/gin"
	"romanapi/api"
)

// @title Roman Numeral API
// @version 0.1
// @description Returns roman numerals in a range specified by query parameters
// @host localhost:8080
// @BasePath /
// @produce json
// @servers




func main() {
	// Create Gin router
	router := gin.Default()

	// Register Routes
	router.GET("/", api.HomePageHandler)
	router.GET("/api/v1/romans", api.GetRomansHandler)

	// Start the server
	// Don't use localhost server address within a Docker container
	// If you do this only 'localhost'
	// (i.e. any service within the Docker container's network) can reach it.
	err := router.Run() // will by default use port 8080
	if err != nil {
		return
	}
}


