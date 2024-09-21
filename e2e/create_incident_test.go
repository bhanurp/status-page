package e2e

import (
	"testing"

	"github.com/bhanurp/status-page/e2e/utils"
)

func TestCreateIncident(t *testing.T) {
	utils.Inite2e()
	numberOfIncidents := utils.FetchUnresolvedIncidentsCount()
	apiKey, statusPageID, statusPageComponentID, hostName := utils.FetchStatusPageDetails()
	utils.CreateIncident(apiKey, hostName, statusPageComponentID, statusPageID, "TestFailure1", "", t)
	// Verify the created incident
	numberOfIncidentsAfterCreate := utils.FetchUnresolvedIncidentsCount()
	if numberOfIncidentsAfterCreate != numberOfIncidents+1 {
		t.Fatalf("Expected to find incident with name 'TestFailure1' in unresolved incidents")
	}
}
