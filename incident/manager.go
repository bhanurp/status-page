package incident

func ResolveAllIncident() error {
	incident := UpdateIncident{}
	incident.ResolveIncidents()
	return nil
}

func CreateNewIncident(incident *CreateIncident) (*Incident, error) {
	return incident.SendCreateIncidentRequest()
}

func SaveIncidentChanges(incident *UpdateIncident) error {
	return incident.UpdateIncidentMatchingWithComponent(incident.ID, ComponentStatusOperational)
}

func FetchAllUnresolvedIncidents() ([]Incident, error) {
	return fetchUnresolvedIncidents()
}

func DeleteGivenIncident(incident *DeleteIncident) error {
	return incident.DeleteIncidentHandler()
}

func FetchIncidentByIncidentID(incidentID string) (*Incident, error) {
	apiKey, statusPageID, _, hostName := FetchStatusPageDetails()
	return fetchIncidentByIncidentID(apiKey, hostName, statusPageID, incidentID)
}
