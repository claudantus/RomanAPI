package api

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"github.com/go-playground/validator/v10"
	"errors"
	"romanapi/roman"
	"romanapi/docs"
)

const (
	welcomeMessageStr string = "Welcome to the Roman Numeral API.\n" +
		"Get a range of roman numerals via /api/v1/romans with the query parameters \n" +
		"\"min\" for the lower and" + "\n" + "\"max\" for the upper bound"
)

// rangeParams defines the input validation structure for the range of Romans.
// TODO: make limits configurable
type rangeParams struct {
	Min int `form:"min" binding:"required,gte=1,lte=3999,ltefield=Max"`
	Max int `form:"max" binding:"required,gte=1,lte=3999,gtefield=Min"`
}

// decimalRoman defines the return type structure.
type decimalRoman struct {
	Decimal int `form:"decimal" example:"10"`
	Roman string `form:"roman" example:"X"`
}

// welcomeMessage define the welcome message structure.
type welcomeMessage struct {
	Message string `form:"message" example:"Welcome to the Roman Numeral API.\nGet a range of roman numerals via /api/v1/romans with the query parameters \n\"min\" for the lower and\n\"max\" for the upper bound"`
}

// ErrorMsg defines the structure for basic error messages.
type ErrorMsg struct {
    Field string `json:"field" example:"Min"`
    Message   string `json:"message" example:"This field is required"`
}

// @title Roman Numeral API
// @description This API allows users to convert decimal numbers to Roman numerals and retrieve a range of Roman numerals.
// @BasePath /api/v1
// TODO: add contact info
func SetUpRouter() *gin.Engine {
	docs.SwaggerInfo.BasePath = "/api/v1"
	router := gin.Default()
	router.GET("/api/v1/romans", GetRomansHandler)
	router.GET("/", HomePageHandler)
	// use ginSwagger middleware to serve the API docs
	router.GET("api/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

// getErrorMsg generates descriptive error messages based on field validation errors.
// It takes a validator.FieldError as input and returns a string describing the error.
func getErrorMsg(fe validator.FieldError) string {
    switch fe.Tag() {
        case "required":
            return "This field is required"
        case "lte":
            return "Should be less than or equal to " + fe.Param()
        case "gte":
            return "Should be greater than or equal to " + fe.Param()
		case "ltefield":
			return "Should be greater than or equal to " + fe.Param()
		case "gtefield":
			return "Should be greater than or equal to " + fe.Param()
    }
    return "Unknown error"
}

// @Summary Returns Roman numerals in a specified range
// @Description Uses a min and a max parameter to define the range
// @Accept json
// @Produce json
// @Param min query int true "Min" minimum:1 maximum:3999 example:1
// @Param max query int true "Max" minimum:1 maximum:3999 example:100
// @Success 200 {array} decimalRoman "Successful operation"
// @Failure 400 {array} ErrorMsg "Validation error"
// @Router /romans [get]
func GetRomansHandler(c *gin.Context) {
	// input validation using bindings
	var params rangeParams

	// return status bad request and collect errors for output
	// TODO: improve error messages
	if err := c.Bind(&params); err != nil {
		var ve validator.ValidationErrors
        if errors.As(err, &ve) {
            out := make([]ErrorMsg, len(ve))
            for i, fe := range ve {
                out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
            }
            c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	// initialize array for the list of romans to return
	var decimalRomans []decimalRoman
	// create list of romans
	for decimal := params.Min; decimal < params.Max + 1; decimal++ {
		// get roman number from decimal
		rom, err := roman.DecimalToRoman(decimal)

		// if successful, append the roman to the array
		if err == nil {
			decimalRomans = append(decimalRomans, decimalRoman{decimal, rom})
		}
	}

	// check if list if empty
	if len(decimalRomans) == 0 {
		out := make([]ErrorMsg, 1)
		out[0] = ErrorMsg{"any", "bad input"}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
	}

	c.JSON(http.StatusOK, decimalRomans)
}

// @Summary Shows welcome page
// @Description Displays a short description of how to use the API
// @Accept */*
// @Produce json
// @Success 200 {object} welcomeMessage "Welcome message" 
// @Router / [get]
func HomePageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, welcomeMessage{welcomeMessageStr})
}