package incident

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/bhanurp/status-page/logger"

	"go.uber.org/zap"
)

// createIncidentMetaData creates incident meta data
func createIncidentMetaData(incidentMetadata Metadata) []byte {
	j, err := json.Marshal(incidentMetadata)
	if err != nil {
		logger.Error("Failed marshalling metadata into json data")
		return []byte{}
	}
	return j
}

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
	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		logger.Fatal("API key not found in config file or environment variables")
	}

	return apiKey
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
