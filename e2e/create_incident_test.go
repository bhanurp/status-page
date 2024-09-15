package e2e

import (
	"testing"
)

func TestCreateIncident(t *testing.T) {
	inite2e()
	numberOfIncidents := fetchUnresolvedIncidentsCount()
	apiKey, statusPageID, statusPageComponentID, hostName := fetchStatusPageDetails()
	createIncident(apiKey, hostName, statusPageComponentID, statusPageID, "TestFailure1", "", t)
	// Verify the created incident
	numberOfIncidentsAfterCreate := fetchUnresolvedIncidentsCount()
	if numberOfIncidentsAfterCreate != numberOfIncidents+1 {
		t.Fatalf("Expected to find incident with name 'TestFailure1' in unresolved incidents")
	}
}
