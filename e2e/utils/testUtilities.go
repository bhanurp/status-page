package utils

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/bhanurp/status-page/incident"
	"github.com/bhanurp/status-page/logger"
)

func Inite2e() {
	logger.Info("Starting e2e tests")
	envPath := "utils/statuspage.env"
	if _, err := os.Stat(envPath); err == nil {
		if err := godotenv.Load(envPath); err != nil {
			logger.Warn("Error loading statuspage.env file: %v", zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
		}
	} else {
		logger.Info("File statuspage.env not found, proceeding without it")
	}
}

func CreateIncident(apiKey, statusPageComponentID, statusPageID, incidentName, incidentBody string, t *testing.T) {
	client := incident.NewDefaultIncident(apiKey, statusPageComponentID, statusPageID, incidentName, incidentBody)
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
			return i
		}
	}
	logger.Error("failed to find incident with name: ", zap.String("incidentName", incidentName))
	return incidentIdentified
}

func CreateStatusPageURL(pageID, incidentID string) string {
	return "https://api.statuspage.io/v1/pages/" + pageID + "/incidents/" + incidentID
}

func FetchAPIKey() string {
	return os.Getenv("STATUS_PAGE_BEARER_TOKEN")
}

func FetchStatusPageID() string {
	return os.Getenv("STATUS_PAGE_ID")
}

func FetchStatusPageComponentID() string {
	return os.Getenv("STATUS_PAGE_COMPONENT_ID")
}

func FetchHostName() string {
	return "api.statuspage.io"
}

func GetIncidentName() string {
	return "TestFailureIncident"
}
