package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	// Get API key from environment variable or use the provided one
	apiKey := os.Getenv("FOOTBALL_DATA_KEY")
	if apiKey == "" {
		apiKey = "3e51e2d6acab4d3a89ea86bc3f112fdd" // Use the provided key if not set in environment
	}

	fmt.Printf("Testing API key: %s\n", apiKey)

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Test API connection with a simple request
	baseURL := "https://api.football-data.org/v4"

	// 1. Test competitions endpoint
	fmt.Println("\n1. Testing competitions endpoint...")
	testEndpoint(client, baseURL+"/competitions", apiKey)

	// 2. Test matches endpoint for yesterday
	fmt.Println("\n2. Testing matches endpoint for yesterday...")
	yesterday := time.Now().AddDate(0, 0, -1)
	yesterdayStr := yesterday.Format("2006-01-02")
	testEndpoint(client, baseURL+"/matches?dateFrom="+yesterdayStr+"&dateTo="+yesterdayStr, apiKey)

	// 3. Test Premier League matches
	fmt.Println("\n3. Testing Premier League matches...")
	testEndpoint(client, baseURL+"/competitions/PL/matches", apiKey)
}

func testEndpoint(client *http.Client, url string, apiKey string) {
	// Create request
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Add API key header
	req.Header.Set("X-Auth-Token", apiKey)

	// Execute request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error executing request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Check response status
	fmt.Printf("Status code: %d\n", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API request failed with status code: %d\n", resp.StatusCode)
		return
	}

	// Parse and print a summary of the response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("Error decoding response: %v\n", err)
		return
	}

	// Print a summary of the response
	fmt.Println("Response summary:")
	for key, value := range result {
		switch v := value.(type) {
		case []interface{}:
			fmt.Printf("  %s: array with %d items\n", key, len(v))
		default:
			fmt.Printf("  %s: %v\n", key, value)
		}
	}
}
