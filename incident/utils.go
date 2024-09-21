package incident

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

func createHeaders(apiKey string) map[string]string {
	headers := make(map[string]string, 0)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "OAuth " + apiKey
	return headers
}

// CreateStatusPageURL constructs the URL for the status page incident.
func createStatusPageURL(pageID, incidentID string) string {
	return "https://api.statuspage.io/v1/pages/" + pageID + "/incidents/" + incidentID
}
