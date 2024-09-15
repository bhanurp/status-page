package e2e

import (
	"os"
	"testing"

	"github.com/bhanurp/status-page/statuspage" // Replace "your-package-path" with the actual package path
)

func TestCreateIncident(t *testing.T) {
	client := statuspage.NewClient(os.Getenv("STATUS_PAGE_BEARER_TOKEN"))

	// Create Incident
	createIncident := client.Incident{
		Name:   "Test Incident",
		Status: "investigating",
	}
	createdIncident, err := client.CreateIncident(createIncident)
	if err != nil {
		t.Fatalf("Failed to create incident: %v", err)
	}

	// Verify the created incident
	if createdIncident.Name != createIncident.Name {
		t.Fatalf("Expected incident name %s, got %s", createIncident.Name, createdIncident.Name)
	}
}
