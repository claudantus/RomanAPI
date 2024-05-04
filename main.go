package main

import (
	"romanapi/api"
)

// TODO: provide OpenAPI specs in api at http://localhost:8080/docs

// api specs
// @title Roman Numeral API
// @version 1.0
// @description Returns roman numerals in a range specified by query parameters
// @host localhost:8080
// @BasePath /
// @produce json
// @servers

// HomePageHandler godoc
// @Summary      Shows welcome page
// @Description  Displays a short descprition on how to use the API
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]

func main() {
	// Create Gin router
	router := api.SetUpRouter()

	// Start the server
	// Don't use localhost server address within a Docker container
	// If you do this only 'localhost'
	// (i.e. any service within the Docker container's network) can reach it.
	err := router.Run() // will by default use port 8080
	if err != nil {
		return
	}
}


