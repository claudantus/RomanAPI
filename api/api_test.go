package api

import (
	// "github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO: decide whether to combine all handlers in one test function or leave them separated for readability
// TODO: test the content such as roman numerals or error messages

func TestHomePageHandler(t *testing.T) {
	tests := []struct {
		input string
		wantCode int
		wantMessage string
	}{
		{"", http.StatusOK, ""}, 			 		// get home page
		{"something", http.StatusNotFound, ""}, 	// get not existing endpoint
	}
	
	r := SetUpRouter()

	for _, tt := range tests {
		req, _ := http.NewRequest("GET", "/" + tt.input, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, tt.wantCode, w.Code)
	}
}

// TODO: Test response message in detail
func TestGetRomansHandler(t *testing.T) {
	tests := []struct {
		input string
		wantCode int
		wantMessage string
	}{
		{"", http.StatusBadRequest, ""}, 			 		// no query parameters
		{"?min=1&max=2", http.StatusOK, ""}, 		 		// good parameters
		{"?min=2&max=1", http.StatusBadRequest, ""}, 		// min larger than max
		{"?min=0&max=1", http.StatusBadRequest, ""}, 		// min out of bounds
		{"?min=3999&max=4000", http.StatusBadRequest, ""}, 	// max out of bounds
		{"?min=a&max=1", http.StatusBadRequest, ""}, 		// min wrong type
		{"?min=1&max=a", http.StatusBadRequest, ""}, 		// max wrong type
		{"?max=2", http.StatusBadRequest, ""}, 				// min field missing
		{"?min=1", http.StatusBadRequest, ""}, 				// max field missing
		{"?mini=1&max=a", http.StatusBadRequest, ""}, 		// wrong parameter name
	}

	r := SetUpRouter()

	for _, tt := range tests {
		req, _ := http.NewRequest("GET", "/api/v1/romans" + tt.input, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, tt.wantCode, w.Code)
	}
}
