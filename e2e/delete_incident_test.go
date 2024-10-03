package e2e

import (
	"log"
	"testing"

	"github.com/bhanurp/status-page/e2e/utils"
	"github.com/bhanurp/status-page/incident"
)

// TestDeleteIncident tests the deletion of an incident
func TestDeleteIncident(t *testing.T) {
	apiKey := utils.FetchAPIKey()
	statusPageID := utils.FetchStatusPageID()
	statusPageComponentID := utils.FetchStatusPageComponentID()

	deleteIncident := incident.BuildDeleteIncident()
	deleteIncident.SetAPIKey(apiKey).
		SetPageID(statusPageID).
		SetTitle("Test Incident for Deletion").
		SetMessage("Test Incident for Deletion")

	createdIncident := incident.NewDefaultIncident(apiKey, statusPageComponentID, statusPageID, "Test Incident for Deletion", "")
	_, err := createdIncident.SendCreateIncidentRequest()
	if err != nil {
		t.Fatalf("Failed to create incident: %v", err)
	}
	unresolvedIncidentsBeforeDelete := utils.FetchUnresolvedIncidentsCount()
	if unresolvedIncidentsBeforeDelete == 0 {
		t.Fatalf("Failed to create incident")
	}
	incidentToBeDeleted := utils.FetchIncidentByNameFromUnresolvedIncidents(incident.IncidentNamePrefix + " Test Incident for Deletion")
	deleteIncident.ID = incidentToBeDeleted.ID
	err = deleteIncident.DeleteIncidentHandler()
	if err != nil {
		t.Fatalf("Failed to delete incident: %v", err)
	}
	unresolvedIncidentsAfterDelete := utils.FetchUnresolvedIncidentsCount()
	if unresolvedIncidentsAfterDelete != unresolvedIncidentsBeforeDelete-1 {
		t.Fatalf("Failed to delete incident")
	}
	log.Println("Deleted incident successfully")
}
