package incident

const (
	// IncidentNamePrefix Prefix for Incidents Created via JFrog pipelines tasks
	IncidentNamePrefix = "[JFPIP]"
	// IncidentNameOldPrefix for Incidents created via JFrog pipelines tasks to be deleted
	IncidentNameOldPrefix = "[JFROG PIPELINES]"
	// HostName Atlassian status page host name
	HostName = "https://api.statuspage.io"
	// IncidentName name of the incident which is created
	// IncidentName = "Pipelines Builds are Failing"
	// IncidentStatusIdentified default incident status for create incident
	IncidentStatusIdentified = "identified"
	// IncidentStatus resolved
	IncidentStatusResolved = "resolved"
	// Component Operational
	ComponentStatusOperational = "operational"
	// Component major outage
	ComponentStatusMajorOutage = "major_outage"
)
