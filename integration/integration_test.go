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

// Define your test suite struct
type IntegrationTestSuite struct {
    suite.Suite
    container testcontainers.Container
}

// SetupSuite is called once before the tests in the suite are run
func (suite *IntegrationTestSuite) SetupSuite() {
    ctx := context.Background()

    // Define the container with your API image
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

// TestAPIIntegration tests the API integration
func (suite *IntegrationTestSuite) TestAPIIntegration() {
    ctx := context.Background()

    // Get the container's host and port
    endpoint, err := suite.container.Endpoint(ctx, "")
    assert.NoError(suite.T(), err)

    // Make HTTP request to the API
    resp, err := http.Get(fmt.Sprintf("http://%s/api/v1/romans?min=1&max=10", endpoint))
    assert.NoError(suite.T(), err)

    // Ensure status code is 200 OK
    assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)
}

// Run the test suite
func TestIntegrationTestSuite(t *testing.T) {
    suite.Run(t, new(IntegrationTestSuite))
}
