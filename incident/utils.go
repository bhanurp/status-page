package incident

import (
	"encoding/json"

	"github.com/bhanurp/status-page/logger"
)

const (
	// IncidentNamePrefix Prefix for Incidents Created
	IncidentNamePrefix = "[TEST]"
	// HostName Atlassian status page host name
	HostName = "https://api.statuspage.io"
	// IncidentName name of the incident which is created
	// IncidentStatusIdentified default incident status for create incident
	IncidentStatusIdentified = "identified"
	// IncidentStatus resolved
	IncidentStatusResolved      = "resolved"
	IncidentStatusInvestigating = "investigating"
	IncidentStatusMonitoring    = "monitoring"
	// Component Operational
	ComponentStatusOperational = "operational"
	// Component major outage
	ComponentStatusMajorOutage = "major_outage"
)

// createIncidentMetaData creates incident meta data
func CreateIncidentMetaData(incidentMetadata Metadata) []byte {
	j, err := json.Marshal(incidentMetadata)
	if err != nil {
		logger.Error("Failed marshalling metadata into json data")
		return []byte{}
	}
	return j
}

// CreateStatusPageURL constructs the URL for the status page incident.
func createStatusPageURL(pageID, incidentID string) string {
	return "https://api.statuspage.io/v1/pages/" + pageID + "/incidents/" + incidentID
}
