package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"romanapi/roman"

	"fmt"

	"strconv"
)

const (
	MIN_DECIMAL int = 1
	MAX_DECIMAL int = 3999
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
	fmt.Println( romans.IntToRoman(10))
	router.Run() // will by default use port 8080
}

func homePage(c *gin.Context) {
	lower_bound, err := strconv.Atoi(c.DefaultQuery("min", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter format. Must be an integer"})
		return
	}
	upper_bound, err := strconv.Atoi(c.DefaultQuery("max", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter format. Must be an integer"})
		return
	}

	if lower_bound > upper_bound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter format. Min must be higher than or equal to max"})
		return
	}

	if lower_bound < MIN_DECIMAL || upper_bound > MAX_DECIMAL {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid parameter format. Values must be between %d and %d", MIN_DECIMAL, MAX_DECIMAL)})
		return
	}

	c.String(http.StatusOK, "This is my home page min %d, max %d", lower_bound, upper_bound)
}
