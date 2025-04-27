package swagger

import (
	"fmt"
	"os"

	docs "management-be/docs/swagger"
)

// Init initializes the Swagger documentation with the correct host and schemes
// This function should be called at application startup
func Init() {
	// Get API host and port from environment variables
	apiHost := os.Getenv("API_HOST")
	apiPort := os.Getenv("API_PORT")

	// If environment variables are not set, use defaults
	if apiHost == "" {
		apiHost = "localhost"
	}
	if apiPort == "" {
		apiPort = "8080"
	}

	// Update the host with the environment variables
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", apiHost, apiPort)

	// Add schemes
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Log the Swagger configuration
	fmt.Printf("Swagger UI available at http://%s:%s/swagger/index.html\n", apiHost, apiPort)
}
