package incident

import (
	"encoding/json"
	"net/http"

	"github.com/bhanurp/rest"
	"github.com/bhanurp/status-page/common"
	"github.com/bhanurp/status-page/logger"
	"github.com/bhanurp/status-page/statuspageurl"
	"go.uber.org/zap"
)

// fetchUnresolvedIncidents returns all unresolved incidents for the status page
func fetchUnresolvedIncidents() ([]Incident, error) {
	apiKey, pageID, _, _ := common.FetchStatusPageDetails()
	get := rest.GetRequest{}
	resp, err := get.Do(statuspageurl.BaseURL+"pages/"+pageID+"/incidents/unresolved", nil, map[string]string{"Authorization": "OAuth " + apiKey}, 10)
	if err != nil {
		return []Incident{}, err
	}
	return handleResponse(resp)
}

func handleResponse(resp *rest.Response) ([]Incident, error) {
	incidents := make([]Incident, 0)
	if resp.StatusCode == http.StatusOK {
		err := json.Unmarshal(resp.Body, &incidents)
		if err != nil {
			return incidents, err
		}
	}
	logger.Debug(string(resp.Body))
	return incidents, nil
}

func fetchIncidentByIncidentID(apiKey, pageID, incidentID string) (*Incident, error) {
	incident := new(Incident)
	get := rest.GetRequest{}
	resp, err := get.Do(statuspageurl.BaseURL+"pages/"+pageID+"/incidents/"+incidentID, nil, map[string]string{"Authorization": "OAuth " + apiKey}, 10)
	if err != nil {
		return incident, err
	}
	response := string(resp.Body)
	logger.Debug("Response from fetch incident by ID ", zap.String("Response", response))
	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(resp.Body, incident)
		if err != nil {
			return incident, err
		}
	}
	logger.Debug(string(resp.Body))
	return incident, nil
}

func fetchUpcomingIncidents() ([]Incident, error) {
	apiKey, pageID, _, _ := common.FetchStatusPageDetails()
	get := rest.GetRequest{}
	resp, err := get.Do(statuspageurl.BaseURL+"pages/"+pageID+"/incidents/upcoming", nil, map[string]string{"Authorization": "OAuth " + apiKey}, 10)
	if err != nil {
		return []Incident{}, err
	}
	return handleResponse(resp)
}

func fetchUpcomingActiveMaintainaances() ([]Incident, error) {
	apiKey, pageID, _, _ := common.FetchStatusPageDetails()
	get := rest.GetRequest{}
	resp, err := get.Do(statuspageurl.BaseURL+"pages/"+pageID+"/incidents/upcoming_active_maintenances", nil, map[string]string{"Authorization": "OAuth " + apiKey}, 10)
	if err != nil {
		return []Incident{}, err
	}
	return handleResponse(resp)
}

func fetchUpcomingScheduledIncidents() ([]Incident, error) {
	apiKey, pageID, _, _ := common.FetchStatusPageDetails()
	get := rest.GetRequest{}
	resp, err := get.Do(statuspageurl.BaseURL+"pages/"+pageID+"/incidents/scheduled", nil, map[string]string{"Authorization": "OAuth " + apiKey}, 10)
	if err != nil {
		return []Incident{}, err
	}
	return handleResponse(resp)
}
