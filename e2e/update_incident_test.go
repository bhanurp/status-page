package e2e

import (
	"log"
	"testing"

	"github.com/bhanurp/status-page/incident"
)

func TestUpdateIncident(t *testing.T) {
	inite2e()
	apiKey, statusPageID, statusPageComponentID, hostName := fetchStatusPageDetails()

	createIncident(apiKey, hostName, statusPageComponentID, statusPageID, "TestUpdateIncident", "", t)
	incidents, err := fetchUnresolvedIncidents()
	if err != nil {
		t.Fatalf("Failed to fetch unresolved incidents: %v", err)
	}
	var incidentToBeUpdated incident.Incident
	for _, i := range incidents {
		log.Println(i.Name)
		if i.Name == incident.IncidentNamePrefix+" TestUpdateIncident" {
			log.Println("Incident to be updated: ", i.Name)
			incidentToBeUpdated = i
		}
	}
	if incidentToBeUpdated.ID == "" {
		t.Fatalf("Failed to find incident with name 'TestUpdateIncident' in unresolved incidents")
	}
	countOfIncidentsBeforeUpdate := len(incidents)
	updateIncident := incident.NewUpdateIncident(apiKey, hostName, statusPageComponentID, statusPageID, "UpdatedIncidentBody", "", incident.IncidentStatusResolved)
	updateIncident.SetIncidentName(incidentToBeUpdated.Name)
	// Create an incident to update
	err = updateIncident.UpdateIncidentMatchingWithComponent(incidentToBeUpdated.ID, incident.ComponentStatusOperational)
	if err != nil {
		t.Fatalf("Failed to update incident: %v", err)
	}
	afterUpdateIncidents := fetchUnresolvedIncidentsCount()
	if countOfIncidentsBeforeUpdate-1 != afterUpdateIncidents {
		t.Fatalf("countOfIncidentsBeforeUpdate: %d is not greater than afterUpdateIncidents: %d", countOfIncidentsBeforeUpdate, afterUpdateIncidents)
	}
}
