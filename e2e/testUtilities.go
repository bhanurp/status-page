package e2e

import (
	"os"
	"testing"

	"github.com/bhanurp/status-page/incident"
)

func inite2e() {
	os.Setenv("STATUS_PAGE_BEARER_TOKEN", "a18f8b4b9a164b708e3bc3957258829b")
	os.Setenv("STATUS_PAGE_ID", "3mz82cxnz2x9")
	os.Setenv("STATUS_PAGE_COMPONENT_ID", "9dmxjrwqp3qc")
	os.Setenv("STATUS_PAGE_HOSTNAME", "api.statuspage.io")
}

func fetchStatusPageDetails() (string, string, string, string) {
	apiKey := os.Getenv("STATUS_PAGE_BEARER_TOKEN")
	statusPageID := os.Getenv("STATUS_PAGE_ID")
	statusPageComponentID := os.Getenv("STATUS_PAGE_COMPONENT_ID")
	hostName := os.Getenv("STATUS_PAGE_HOSTNAME")
	return apiKey, statusPageID, statusPageComponentID, hostName
}

func createIncident(apiKey, hostName, statusPageComponentID, statusPageID, incidentName, incidentBody string, t *testing.T) {
	client := incident.NewDefaultIncident(apiKey, hostName, statusPageComponentID, statusPageID, incidentName, incidentBody)
	err := client.SendCreateIncidentRequest(apiKey)
	if err != nil {
		t.Fatalf("Failed to post incident: %v", err)
	}
}

func fetchUnresolvedIncidentsCount() int {
	incidents, err := fetchUnresolvedIncidents()
	if err != nil {
		return 0
	}
	return len(incidents)
}

func fetchUnresolvedIncidents() ([]incident.Incident, error) {
	inite2e()
	apiKey, statusPageID, _, hostName := fetchStatusPageDetails()
	incidents, err := incident.FetchUnresolvedIncidents(apiKey, hostName, statusPageID)
	if err != nil {
		return nil, err
	}
	return incidents, nil
}
