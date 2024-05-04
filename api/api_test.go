package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePageHandler(t *testing.T) {
	tests := []struct {
		input string
		wantCode int
		wantMessage string
	}{
		{"", http.StatusOK, ""}, 			 		// get home page
		{"something", http.StatusNotFound, ""}, 	// get not existing endpoint

	}
	r := gin.Default()	
	r.GET("/", HomePageHandler)

	for _, tt := range tests {
		req, _ := http.NewRequest("GET", "/" + tt.input, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, tt.wantCode, w.Code)
	}
}

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
		{"?mini=1&max=a", http.StatusBadRequest, ""}, 		// wrong parameter name
	}

	r := gin.Default()	
	r.GET("/api/v1/romans", GetRomansHandler)

	for _, tt := range tests {
		req, _ := http.NewRequest("GET", "/api/v1/romans" + tt.input, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, tt.wantCode, w.Code)
		
	}
}
