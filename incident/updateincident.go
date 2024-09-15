package incident

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/bhanurp/status-page/logger"

	"go.uber.org/zap"
)

func NewUpdateIncident(apiKey, hostName, componentID, pageID, incidentBody, incidentHeader, incidentStatus string) *UpdateIncident {
	u := new(UpdateIncident)
	u.APIKey = apiKey
	u.HostName = hostName
	u.ComponentID = componentID
	u.PageID = pageID
	u.IncidentStatus = incidentStatus
	u.IncidentHeader = incidentHeader
	u.IncidentBody = incidentBody
	u.Metadata = Metadata{}
	return u
}

// ResolveIncidents fetches all unresolved incidents on the status page and filters
// components interested in to be updated
func (u *UpdateIncident) ResolveIncidents(wg *sync.WaitGroup) error {
	var unresolvedIncidents []string
	defer wg.Done()
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
	statusPageHTTPClient := new(StatusPageHTTPClient)
	payloadBytes := u.prepareIncidentBodyRequest(componentStatus)
	logger.Debug(" Received for updating incident", zap.String("Payload", string(payloadBytes)))
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("PUT", u.HostName+"/v1/pages/"+u.PageID+"/incidents/"+unresolvedIncident, body)
	if err != nil {
		logger.Warn("Failed to update incident: " + unresolvedIncident)
	}
	req.Header.Set("Authorization", "OAuth "+u.APIKey)
	req.Header.Set("Content-Type", "application/json")
	statusPageHTTPClient.req = req
	err = statusPageHTTPClient.SendHTTPRequest(unresolvedIncident)
	if err != nil {
		return err
	}
	logger.Info("Completed performing updating of incidents")
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
