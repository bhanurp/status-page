package e2e

import (
	"os"
	"testing"

	"github.com/bhanurp/status-page/incident"
	"github.com/bhanurp/status-page/logger"
)

func inite2e() {
	logger.Info("Starting e2e tests")
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
