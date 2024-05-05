package integration

import (
    "context"
    "fmt"
    "net/http"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "github.com/testcontainers/testcontainers-go"
    "github.com/testcontainers/testcontainers-go/wait"
)

// Define the test suite struct
type IntegrationTestSuite struct {
    suite.Suite
    container testcontainers.Container
}

// SetupSuite is called once before the tests in the suite are run
func (suite *IntegrationTestSuite) SetupSuite() {
    ctx := context.Background()

    // Define the container with the API image
    req := testcontainers.ContainerRequest{
        Image:        "romanapi",
        ExposedPorts: []string{"8080/tcp"},
        WaitingFor:   wait.ForHTTP("http://127.0.0.1:8080/").WithStartupTimeout(30 * time.Second),
    }

    // Start the container
    container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
        ContainerRequest: req,
        Started:          true,
    })
    assert.NoError(suite.T(), err)
    suite.container = container
}

// TearDownSuite is called once after all the tests in the suite have been run
func (suite *IntegrationTestSuite) TearDownSuite() {
    // Stop and remove the container
    err := suite.container.Terminate(context.Background())
    assert.NoError(suite.T(), err)
}

// TODO: Test response messages in detail
// TODO: Copy paste from api_test.go... Import this.
// TestAPIIntegration tests the API integration
func (suite *IntegrationTestSuite) TestAPIIntegration() {
    tests := []struct {
		input string
		wantCode int
		wantMessage string
	}{
        {"", http.StatusBadRequest, ""}, 			 		// no query parameters
		{"?min=1&max=2", http.StatusOK, ""}, 		        // good parameters
		{"?min=2&max=1", http.StatusBadRequest, ""}, 		// min larger than max
		{"?min=0&max=1", http.StatusBadRequest, ""}, 		// min out of bounds
		{"?min=3999&max=4000", http.StatusBadRequest, ""}, 	// max out of bounds
		{"?min=a&max=1", http.StatusBadRequest, ""}, 		// min wrong type
		{"?min=1&max=a", http.StatusBadRequest, ""}, 		// max wrong type
		{"?max=2", http.StatusBadRequest, ""}, 				// min field missing
		{"?min=1", http.StatusBadRequest, ""}, 				// max field missing
		{"?mini=1&max=1", http.StatusBadRequest, ""}, 		// wrong parameter name
    }
    ctx := context.Background()

    // Get the container's host and port
    endpoint, err := suite.container.Endpoint(ctx, "")
    assert.NoError(suite.T(), err)

    // Make HTTP request to the API
    for _, tt := range tests {
        resp, err := http.Get(fmt.Sprintf("http://%s/api/v1/romans" + tt.input, endpoint))
        assert.NoError(suite.T(), err)

        // Ensure status code is OK
        assert.Equal(suite.T(), tt.wantCode, resp.StatusCode)
    }
}

// Run the test suite
func TestIntegrationTestSuite(t *testing.T) {
    suite.Run(t, new(IntegrationTestSuite))
}
