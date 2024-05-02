package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"romanapi/roman"

	"fmt"

	"strconv"

	"strings"
)

const (
	// lower boundary of accepted input to get roman number
	MIN_DECIMAL int = 1
	// upper boundary of accepted input to get roman number
	MAX_DECIMAL int = 3999 
)

func main() {
	// Create Gin router
	router := gin.Default()

	// Register Routes
	router.GET("/", getHome)
	router.GET("/romans", getRomans)

	// Start the server
	// Don't use localhost server address within a Docker container
	// If you do this only 'localhost'
	// (i.e. any service within the Docker container's network) can reach it.
	router.Run() // will by default use port 8080
}

func getRomans(c *gin.Context) {
	// get range for the romans from query parameter range. Default to 1-10
	romansRange := c.Query("range")
	romansRangeList := strings.Split(romansRange, "-")
	if len(romansRangeList) != 2 && romansRange != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid range for romans. Must be in the format min-max"})
		return
	}

	// split
	// get lower boundary for range of roman numbers from query parameter. Default to 1
	// convert to integer
	var lowerBoundStr string
	var upperBoundStr string
	if romansRange == "" {
		lowerBoundStr = c.DefaultQuery("min", "1")
		upperBoundStr = c.DefaultQuery("max", "10")
	} else {
		lowerBoundStr = romansRangeList[0]
		upperBoundStr = romansRangeList[1]
	}

	lowerBound, err := strconv.Atoi(lowerBoundStr)

	// if lowerBound cannot be converted to an integer, return status bad request
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter format. Must be an integer"})
		return
	}

	// get upper boundary for range of roman numbers from query parameter. Default to 10
	// convert to integer
	upperBound, err := strconv.Atoi(upperBoundStr)
		
	// if upperBound cannot be converted to an integer, return status bad request
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter format. Must be an integer"})
		return
	}

	// if lowerBound is higher than upperBound, return status bad request
	if lowerBound > upperBound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter format. Min must be higher than or equal to max"})
		return
	}

	// if the lower and upper bound are not in the requested range of 1-3999, return bad request
	if lowerBound < MIN_DECIMAL || upperBound > MAX_DECIMAL {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid parameter format. Values must be between %d and %d", MIN_DECIMAL, MAX_DECIMAL)})
		return
	}

	// initialize array for the list of romans to return
	var listOfRomans []string

	// create list of romans
	for decimal := lowerBound; decimal < upperBound + 1; decimal++ {
		// get roman number from decimal
		rom, err := roman.IntToRoman(decimal)

		// if successful, append the roman to the array
		if err == nil {
			listOfRomans = append(listOfRomans, rom)
		}
	}

	// return status ok and the list of romans
	c.JSON(http.StatusOK, listOfRomans)
}

func getHome(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the Roman Numeral API.\n" +
		"Get a range of roman numerals via /romans with the query parameters " +
		"min for the lower and max for the upper bound \n" +
		"or with the parameter range in the forman min-max (e.g. 1-10)")
}
