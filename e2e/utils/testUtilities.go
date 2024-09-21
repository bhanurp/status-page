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

func CreateIncident(apiKey, hostName, statusPageComponentID, statusPageID, incidentName, incidentBody string, t *testing.T) {
	client := incident.NewDefaultIncident(apiKey, hostName, statusPageComponentID, statusPageID, incidentName, incidentBody)
	_, err := client.SendCreateIncidentRequest()
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
	incidents, err := incident.FetchAllUnresolvedIncidents()
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

func FetchAPIKey() string {
	return os.Getenv("API_KEY")
}

func FetchStatusPageID() string {
	return os.Getenv("STATUS_PAGE_ID")
}

func FetchStatusPageComponentID() string {
	return os.Getenv("STATUS_PAGE_COMPONENT_ID")
}

func FetchHostName() string {
	return os.Getenv("HOST_NAME")
}

func GetIncidentName() string {
	return "TestFailureIncident"
}
