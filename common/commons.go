package common

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/bhanurp/status-page/logger"
	"go.uber.org/zap"
)

// GetAPIKey to retrieve the API key from config file or environment variable
func GetAPIKey() string {
	// Get the $HOME directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logger.Fatal("Unable to find home directory: %s", zap.String("error", err.Error()))
	}

	// Path to the config file
	configFilePath := filepath.Join(homeDir, ".secrets/statuspage", "apikey")

	// Attempt to read the API key from the config file
	apiKey, err := readAPIKeyFromFile(configFilePath)
	if err == nil && apiKey != "" {
		return apiKey
	}

	// If not found in file, fallback to environment variable
	apiKey = os.Getenv("STATUS_PAGE_BEARER_TOKEN")
	if apiKey == "" {
		logger.Fatal("API key not found in config file or environment variables")
	}

	return apiKey
}

func FetchStatusPageDetails() (string, string, string, string) {
	apiKey := os.Getenv("STATUS_PAGE_BEARER_TOKEN")
	statusPageID := os.Getenv("STATUS_PAGE_ID")
	statusPageComponentID := os.Getenv("STATUS_PAGE_COMPONENT_ID")
	hostName := "api.statuspage.io"
	return apiKey, statusPageID, statusPageComponentID, hostName
}

// Helper function to read the API key from a file
func readAPIKeyFromFile(filePath string) (string, error) {
	// Read the content of the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Trim any whitespace or newlines from the API key
	return strings.TrimSpace(string(data)), nil
}

func CreateHeaders() map[string]string {
	apiKey := GetAPIKey()
	headers := make(map[string]string, 0)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "OAuth " + apiKey
	return headers
}
