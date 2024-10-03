package e2e

import (
	"testing"

	"github.com/bhanurp/status-page/e2e/utils"
	"github.com/bhanurp/status-page/incident"
)

func TestUpdateIncident(t *testing.T) {
	utils.Inite2e()
	apiKey := utils.FetchAPIKey()
	statusPageID := utils.FetchStatusPageID()
	statusPageComponentID := utils.FetchStatusPageComponentID()

	utils.CreateIncident(apiKey, statusPageComponentID, statusPageID, "TestUpdateIncident", "", t)
	incidents, err := utils.FetchUnresolvedIncidents()
	if err != nil {
		t.Fatalf("Failed to fetch unresolved incidents: %v", err)
	}
	incidentToBeUpdated := utils.FetchIncidentByNameFromUnresolvedIncidents(incident.IncidentNamePrefix + " TestUpdateIncident")
	if incidentToBeUpdated.ID == "" {
		t.Fatalf("Failed to find incident with name 'TestUpdateIncident' in unresolved incidents")
	}
	countOfIncidentsBeforeUpdate := len(incidents)
	updateIncident := incident.NewUpdateIncident(apiKey, statusPageComponentID, statusPageID, "UpdatedIncidentBody", "", incident.IncidentStatusResolved)
	updateIncident.SetIncidentName(incidentToBeUpdated.Name)
	// Create an incident to update
	err = updateIncident.UpdateIncidentMatchingWithComponent(incidentToBeUpdated.ID, incident.ComponentStatusOperational)
	if err != nil {
		t.Fatalf("Failed to update incident: %v", err)
	}
	afterUpdateIncidents := utils.FetchUnresolvedIncidentsCount()
	if countOfIncidentsBeforeUpdate-1 != afterUpdateIncidents {
		t.Fatalf("countOfIncidentsBeforeUpdate: %d is not greater than afterUpdateIncidents: %d", countOfIncidentsBeforeUpdate, afterUpdateIncidents)
	}
}
