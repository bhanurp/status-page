package incident

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"sync"

	"github.com/bhanurp/rest"
	"github.com/bhanurp/status-page/logger"

	"go.uber.org/zap"
)

// NewDefaultIncident creates incident struct and returns the pointer to it
func NewDefaultIncident(apiKey, hostName, componentID, pageID, incidentName, incidentBody string) *CreateIncident {
	c := new(CreateIncident)
	c.APIKey = apiKey
	c.HostName = hostName
	c.ComponentID = componentID
	c.PageID = pageID
	c.IncidentName = incidentName
	c.IncidentBody = incidentBody
	c.IncidentStatus = IncidentStatusIdentified
	c.Metadata = Metadata{}
	return c
}

// PostIncident Checks if an Incident is already present for given component on the status page
// and invokes SendCreateIncidentRequest
func (c *CreateIncident) PostIncident(wg *sync.WaitGroup) error {
	defer wg.Done()
	c.IncidentStatus = IncidentStatusIdentified
	incidents, err := FetchUnresolvedIncidents(c.APIKey, c.HostName, c.PageID)
	if err != nil {
		return err
	}
	logger.Debug("unresolved incidents for page: ",
		zap.Any("UnresolvedIncidents", incidents))
	updateIncidents := make([]string, 0)
	// Iterate over existing incidents for evert incident
	for _, incident := range incidents {
		for _, component := range incident.Components {
			logger.Debug("component details",
				zap.String("componentID", c.ComponentID),
				zap.String("existing componentID", component.ID),
				zap.String("IncidentName", incident.Name),
				zap.String("IncidentPrefix", IncidentNamePrefix),
			)
			if strings.Trim(c.ComponentID, " ") == component.ID && strings.HasPrefix(incident.Name, IncidentNamePrefix) {
				logger.Warn("An incident on component exist ",
					zap.String("IncidentURL", incident.Shortlink))
				c.IncidentName = fmt.Sprintf("%s (more...)", incident.Name)
				// Update incidents without resolving
				err := c.UpdateIncidentonFailureReasonChange(updateIncidents, incident)
				if err != nil {
					return err
				}
				return nil
			}
		}
	}
	// Create http request with incident name, status, component ID and status
	_, err = c.SendCreateIncidentRequest(c.APIKey)
	if err != nil {
		return err
	}
	return nil
}

// UpdateIncidentonFailureReasonChange updates incidents
// inCase if the reason for failing of component has changed
// from earlier runs. For example if ServiceA, ServiceB failed initially
// and an incident is created and now only ServiceB is failing updates
// incident with message ServiceB is failing
func (c *CreateIncident) UpdateIncidentonFailureReasonChange(updateIncidents []string, incident Incident) error {
	logger.Info("Incident failure reason has changed from earlier identified")
	// If reasons for failing from earlier monitor is same as current monitor run no need to update incident for the component
	if reflect.DeepEqual(incident.Metadata.Data, c.Metadata.Data) {
		return nil
	}
	// If incident metadata is
	u := NewUpdateIncident(c.APIKey, c.HostName, c.ComponentID, c.PageID, c.IncidentBody, c.IncidentHeader, IncidentStatusIdentified)
	u.IncidentName = c.IncidentName
	u.Metadata = c.Metadata

	err := u.UpdateIncidentMatchingWithComponent(incident.ID, ComponentStatusMajorOutage)
	if err != nil {
		return err
	}
	return nil
}

// SendCreateIncidentRequest calls to create incident on status page, component CreateIncident configured.
func (c *CreateIncident) SendCreateIncidentRequest(apiKey string) (*Incident, error) {
	logger.Debug("Sending Create Incident request")
	c.IncidentName = fmt.Sprintf("%s %s", IncidentNamePrefix, c.IncidentName)
	data := Payload{}
	data.Incident.Body = c.IncidentBody
	data.Incident.Name = c.IncidentName
	data.Incident.Status = c.IncidentStatus
	data.Incident.ComponentIds = append(data.Incident.ComponentIds, c.ComponentID)
	m := make(map[string]string, 0)
	m[c.ComponentID] = ComponentStatusMajorOutage
	data.Incident.Components = m
	data.Incident.Metadata.Data = c.Metadata.Data
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		logger.Error("Failed to marshal incident data")
	}
	headers := make(map[string]string, 0)
	headers["Authorization"] = "OAuth " + apiKey
	headers["Content-Type"] = "application/json"
	p := rest.PostRequest{}
	resp, err := p.Do("https://"+c.HostName+"/v1/pages/"+c.PageID+"/incidents", payloadBytes, headers, 10)
	if err != nil {
		return nil, err
	}
	logger.Debug("Created Incident on ",
		zap.String("ComponentID", c.ComponentID))
	var createdIncident Incident
	if resp.StatusCode/100 == 2 {
		createdIncident = Incident{}
		err = json.Unmarshal(resp.Body, &createdIncident)
		if err != nil {
			return nil, err
		}
		logger.Info("Created task incident ID: %s and incident URL is %s",
			zap.String("CreatedIncidentID", createdIncident.ID),
			zap.String("CreatedIncidentURL", createdIncident.Shortlink))

	}
	logger.Debug(string(resp.Body))
	return &createdIncident, nil
}

// FetchUnresolvedIncidents returns all unresolved incidents for the status page
func FetchUnresolvedIncidents(apiKey, hostName, pageID string) ([]Incident, error) {
	incidents := make([]Incident, 0)
	get := rest.GetRequest{}
	resp, err := get.Do("https://"+hostName+"/v1/pages/"+pageID+"/incidents/unresolved", nil, map[string]string{"Authorization": "OAuth " + apiKey}, 10)
	if err != nil {
		return incidents, err
	}
	response := string(resp.Body)
	logger.Debug("Response from fetch all unresolved ", zap.String("Response", response))
	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(resp.Body, &incidents)
		if err != nil {
			return incidents, err
		}
	} else {
		logger.Debug(string(resp.Body))
	}
	return incidents, nil
}
