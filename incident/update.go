package incident

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/bhanurp/rest"
	"github.com/bhanurp/status-page/logger"
	"github.com/bhanurp/status-page/statuspageurl"

	"go.uber.org/zap"
)

func NewDefaultUpdateIncident() *UpdateIncident {
	return new(UpdateIncident)
}

func NewUpdateIncident(apiKey, componentID, pageID, incidentBody, incidentHeader, incidentStatus string) *UpdateIncident {
	u := new(UpdateIncident)
	u.APIKey = apiKey
	u.ComponentID = componentID
	u.PageID = pageID
	u.IncidentStatus = incidentStatus
	u.IncidentHeader = incidentHeader
	u.IncidentBody = incidentBody
	u.Metadata = Metadata{}
	return u
}

func (u *UpdateIncident) SetIncidentName(incidentName string) {
	u.IncidentName = incidentName
}

func (u *UpdateIncident) SetAPIKey(apiKey string) {
	u.APIKey = apiKey
}

func (u *UpdateIncident) SetHostName(hostName string) {
	u.HostName = hostName
}

func (u *UpdateIncident) SetComponentID(componentID string) {
	u.ComponentID = componentID
}

func (u *UpdateIncident) SetPageID(pageID string) {
	u.PageID = pageID
}

func (u *UpdateIncident) SetIncidentStatus(incidentStatus string) {
	u.IncidentStatus = incidentStatus
}

func (u *UpdateIncident) SetIncidentHeader(incidentHeader string) {
	u.IncidentHeader = incidentHeader
}

func (u *UpdateIncident) SetIncidentBody(incidentBody string) {
	u.IncidentBody = incidentBody
}

// ResolveIncidents fetches all unresolved incidents on the status page and filters
// components interested in to be updated
func (u *UpdateIncident) ResolveIncidents() error {
	var unresolvedIncidents []string
	// Create http request with incident name, status, component ID and status
	incidents, err := u.FetchUnresolvedIncidents()
	if err != nil {
		return err
	}
	logger.Debug("unresolved incidents for page: ", zap.Any("incidents", incidents))
	for _, incident := range incidents {
		for _, component := range incident.Components {
			// Consider updating incidents only when component ID matches and has prefix in component Name as IncidentNamePrefix
			if strings.Trim(u.ComponentID, " ") == component.ID && (strings.HasPrefix(incident.Name, IncidentNamePrefix)) {
				unresolvedIncidents = append(unresolvedIncidents, incident.ID)
				logger.Debug("unresolved incidents for component id : ", zap.Strings("unresolvedIncidents", unresolvedIncidents))
				u.IncidentName = incident.Name
				err = u.UpdateIncidentMatchingWithComponent(incident.ID, ComponentStatusOperational)
			}
		}
	}
	return err
}

// FetchUnresolvedIncidents fetches all unresolved incidents on component ID given as input
func (u *UpdateIncident) FetchUnresolvedIncidents() ([]Incident, error) {
	result := make([]Incident, 0)
	req, err := http.NewRequest("GET", u.HostName+"/v1/pages/"+u.PageID+"/incidents/unresolved", nil)
	if err != nil {
		return result, err
	}
	req.Header.Set("Authorization", "OAuth "+u.APIKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return result, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error("failed to close response body", zap.Error(err))
		}
	}(resp.Body)
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// UpdateIncidentMatchingWithComponent Sends which incidents has to be resolved on a given component ID
func (u *UpdateIncident) UpdateIncidentMatchingWithComponent(unresolvedIncident string, componentStatus string) error {
	payloadBytes := u.prepareIncidentBodyRequest(componentStatus)
	m := make(map[string]string)
	m["Authorization"] = "OAuth " + u.APIKey
	m["Content-Type"] = "application/json"
	put := &rest.PutRequest{}
	resp, err := put.Do(statuspageurl.BaseURL+"pages/"+u.PageID+"/incidents/"+unresolvedIncident, payloadBytes, m, 10)
	if err != nil {
		return err
	}
	logger.Info("Response", zap.String("response", string(resp.Body)))
	return nil
}

// prepareIncidentBodyRequest prepares url values with incidentName, incidentStatus, incidentBody
// componentID and componentStatus
func (u *UpdateIncident) prepareIncidentBodyRequest(componentStatus string) []byte {
	//u.IncidentName = fmt.Sprintf("%s %s", utils.IncidentNamePrefix, u.IncidentName)
	data := Payload{}
	data.Incident.Body = ""
	data.Incident.Name = u.IncidentName
	data.Incident.Status = u.IncidentStatus
	data.Incident.ComponentIds = append(data.Incident.ComponentIds, u.ComponentID)
	m := make(map[string]string)
	m[u.ComponentID] = componentStatus
	data.Incident.Components = m
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		logger.Error("failed to marshal incident data")
		return []byte{}
	}
	return payloadBytes
}
