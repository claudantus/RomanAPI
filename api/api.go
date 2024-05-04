package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/go-playground/validator/v10"
	"errors"
	"romanapi/roman"
)

// struct for validation of input for the range of romans
// TODO: make limits configurable
type rangeParams struct {
	Min int `form:"min" binding:"required,gte=1,lte=3999,ltefield=Max"`
	Max int `form:"max" binding:"required,gte=1,lte=3999,gtefield=Min"`
}

// struct for basic error messages
type ErrorMsg struct {
    Field string `json:"field"`
    Message   string `json:"message"`
}

// utility function to set up the router
func SetUpRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/api/v1/romans", GetRomansHandler)
	router.GET("/", HomePageHandler)
	return router
}

// create error message from field validation
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
	var listOfRomans []string

	// create list of romans
	for decimal := params.Min; decimal < params.Max + 1; decimal++ {
		// get roman number from decimal
		rom, err := roman.DecimalToRoman(decimal)

		// if successful, append the roman to the array
		if err == nil {
			listOfRomans = append(listOfRomans, rom)
		}
	}

	// check if list if empty
	if len(listOfRomans) == 0 {
		out := make([]ErrorMsg, 1)
		out[0] = ErrorMsg{"any", "bad input"}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
	}

	c.JSON(http.StatusOK, listOfRomans)
}

func HomePageHandler(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the Roman Numeral API.\n" +
		"Get a range of roman numerals via /romans with the query parameters \n" +
		`"min" for the lower and` + "\n" + `"max" for the upper bound`)
}