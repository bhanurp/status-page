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
