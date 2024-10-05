package incident

import "github.com/bhanurp/status-page/common"

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
	apiKey, statusPageID, _, _ := common.FetchStatusPageDetails()
	return fetchIncidentByIncidentID(apiKey, statusPageID, incidentID)
}

func FetchUpcomingIncidents() ([]Incident, error) {
	return fetchUpcomingIncidents()
}

func FetchUpcomingActiveMaintainaances() ([]Incident, error) {
	return fetchUpcomingActiveMaintainaances()
}

func FetchUpcomingScheduledIncidents() ([]Incident, error) {
	return fetchUpcomingScheduledIncidents()
}
