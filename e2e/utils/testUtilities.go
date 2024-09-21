package utils

import (
	"log"
	"os"
	"testing"

	"github.com/bhanurp/status-page/incident"
	"github.com/bhanurp/status-page/logger"
)

func Inite2e() {
	logger.Info("Starting e2e tests")
}

func FetchStatusPageDetails() (string, string, string, string) {
	apiKey := os.Getenv("STATUS_PAGE_BEARER_TOKEN")
	statusPageID := os.Getenv("STATUS_PAGE_ID")
	statusPageComponentID := os.Getenv("STATUS_PAGE_COMPONENT_ID")
	hostName := os.Getenv("STATUS_PAGE_HOSTNAME")
	return apiKey, statusPageID, statusPageComponentID, hostName
}

func CreateIncident(apiKey, hostName, statusPageComponentID, statusPageID, incidentName, incidentBody string, t *testing.T) {
	client := incident.NewDefaultIncident(apiKey, hostName, statusPageComponentID, statusPageID, incidentName, incidentBody)
	_, err := client.SendCreateIncidentRequest(apiKey)
	if err != nil {
		t.Fatalf("Failed to post incident: %v", err)
	}
}

func FetchUnresolvedIncidentsCount() int {
	incidents, err := FetchUnresolvedIncidents()
	if err != nil {
		return 0
	}
	return len(incidents)
}

func FetchUnresolvedIncidents() ([]incident.Incident, error) {
	Inite2e()
	apiKey, statusPageID, _, hostName := FetchStatusPageDetails()
	incidents, err := incident.FetchUnresolvedIncidents(apiKey, hostName, statusPageID)
	if err != nil {
		return nil, err
	}
	return incidents, nil
}

func FetchIncidentByNameFromUnresolvedIncidents(incidentName string) incident.Incident {
	var incidentIdentified incident.Incident
	unresolvedIncidents, err := FetchUnresolvedIncidents()
	if err != nil {
		log.Fatalf("Failed to fetch unresolved incidents: %v", err)
	}
	for _, i := range unresolvedIncidents {
		log.Println(i.Name)
		if i.Name == incidentName {
			log.Println("Incident to be updated: ", i.Name)
			incidentIdentified = i
		}
	}
	return incidentIdentified
}

func CreateStatusPageURL(pageID, incidentID string) string {
	return "https://api.statuspage.io/v1/pages/" + pageID + "/incidents/" + incidentID
}
