package e2e

import (
	"os"
	"testing"

	"github.com/bhanurp/status-page/incident"
)

func TestUpdateIncident(t *testing.T) {
	inite2e()
	apiKey := os.Getenv("STATUS_PAGE_BEARER_TOKEN")
	statusPageID := os.Getenv("STATUS_PAGE_ID")
	statusPageComponentID := os.Getenv("STATUS_PAGE_COMPONENT_ID")
	hostName := os.Getenv("STATUS_PAGE_HOSTNAME")

	createIncident(apiKey, hostName, statusPageComponentID, statusPageID, "TestUpdateIncident", "", t)
	incidents, err := fetchUnresolvedIncidents()
	if err != nil {
		t.Fatalf("Failed to fetch unresolved incidents: %v", err)
	}
	var incidentToBeUpdated incident.Incident
	for _, incident := range incidents {
		if incident.Name == "TestUpdateIncident" {
			incidentToBeUpdated = incident
		}
	}
	countOfIncidentsBeforeUpdate := len(incidents)
	updateIncident := incident.NewUpdateIncident(apiKey, hostName, statusPageComponentID, statusPageID, "UpdatedIncidentBody", "", "TestUpdateIncident")
	// Create an incident to update
	err = updateIncident.UpdateIncidentMatchingWithComponent(incidentToBeUpdated.ID, "operational")
	if err != nil {
		t.Fatalf("Failed to update incident: %v", err)
	}

	afterUpdateIncidents := fetchUnresolvedIncidentsCount()

	if countOfIncidentsBeforeUpdate-1 != afterUpdateIncidents {
		t.Fatalf("Expected to find incident with name 'TestUpdateIncident' in unresolved incidents")
	}
}
