package e2e

import (
	"testing"

	"github.com/bhanurp/status-page/e2e/utils"
	"github.com/bhanurp/status-page/incident"
)

func TestCreateIncident(t *testing.T) {
	utils.Inite2e()
	apiKey := utils.FetchAPIKey()
	statusPageID := utils.FetchStatusPageID()
	statusPageComponentID := utils.FetchStatusPageComponentID()
	incidentName := utils.GetIncidentName()
	numberOfIncidents := utils.FetchUnresolvedIncidentsCount()
	createIncidentClient := incident.NewDefaultIncident(apiKey, statusPageComponentID, statusPageID, incidentName, "")
	_, err := incident.CreateNewIncident(createIncidentClient)
	if err != nil {
		t.Fatalf("Failed to create incident: %v", err)
	}
	// Verify the created incident
	numberOfIncidentsAfterCreate := utils.FetchUnresolvedIncidentsCount()
	if numberOfIncidentsAfterCreate != numberOfIncidents+1 {
		t.Fatalf("Failed to create incident")
	}
}
