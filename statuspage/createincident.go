package incident

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"reflect"
	"status-page/logger"
	"strings"
	"sync"
)

// NewCreateIncident creates incident struct and returns the pointer to it
func NewCreateIncident(apiKey, hostName, componentID, pageID string) *CreateIncident {
	c := new(CreateIncident)
	c.APIKey = apiKey
	c.HostName = hostName
	c.ComponentID = componentID
	c.PageID = pageID
	c.Metadata = Metadata{}
	return c
}

// PostIncident Checks if an Incident is already present for given component on the status page
// and invokes SendCreateIncidentRequest
func (c *CreateIncident) PostIncident(wg *sync.WaitGroup) error {
	defer wg.Done()
	c.IncidentName = c.PipelineName
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
				err := c.UpdateIncidentIfPipelineFailuresOnComponentChanges(updateIncidents, incident)
				if err != nil {
					return err
				}
				SetOutputVariables(incident)
				// Returning from here since there is no need to create one more incident
				return nil
			}
		}
	}
	// Create http request with incident name, status, component ID and status
	err = c.SendCreateIncidentRequest(c.APIKey)
	if err != nil {
		return err
	}
	return nil
}

// UpdateIncidentIfPipelineFailuresOnComponentChanges updates incidents
// inCase if the pipelines mapped to a component failing has changed
// from earlier runs. For example if ServiceA, ServiceB failed initially
// and an incident is created and now only ServiceB is failing updates
// incident with message ServiceB is failing
func (c *CreateIncident) UpdateIncidentIfPipelineFailuresOnComponentChanges(updateIncidents []string, incident Incident) error {
	logger.Info("Pipelines failing has changed from earlier identified")
	// If pipelines failing from earlier monitor is same as current monitor run no need to update incident for the component
	if reflect.DeepEqual(incident.Metadata.Data, c.Metadata.Data) {
		return nil
	}
	// If pipelines from earlier monitor is different to current monitor run update incidents
	u := NewUpdateIncident(c.APIKey, c.HostName, c.ComponentID, c.PageID, c.IncidentBody, c.PipelineName, IncidentStatusIdentified)
	u.IncidentName = c.IncidentName
	u.Metadata = c.Metadata

	err := u.UpdateIncidentMatchingWithComponent(incident.ID, ComponentStatusMajorOutage, string(createIncidentMetaData(c.Metadata)))
	if err != nil {
		return err
	}
	return nil
}

// SetOutputVariables creates logger output variables which can be used after executing task
func SetOutputVariables(createdIncident Incident) {
	logger.Debug("Updating output variables, ",
		zap.String("IncidentID", createdIncident.ID),
		zap.String("CreatedIncidentURL", createdIncident.Shortlink))
}

// SendCreateIncidentRequest calls to create incident on status page, component CreateIncident configured.
func (c *CreateIncident) SendCreateIncidentRequest(apiKey string) error {
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
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", c.HostName+"/v1/pages/"+c.PageID+"/incidents", body)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "OAuth "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error("Unable to close response body")
		}
	}(resp.Body)
	logger.Debug("Created Incident on ",
		zap.String("ComponentID", c.ComponentID))
	if resp.StatusCode/100 == 2 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		createdIncident := Incident{}
		err = json.Unmarshal(bodyBytes, &createdIncident)
		if err != nil {
			return err
		}
		// Set output variables
		SetOutputVariables(createdIncident)
		logger.Info("Created task incident ID: %s and incident URL is %s",
			zap.String("CreatedIncidentID", createdIncident.ID),
			zap.String("CreatedIncidentURL", createdIncident.Shortlink))
	} else {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		logger.Debug(string(bodyBytes))
	}
	return nil
}

// FetchUnresolvedIncidents returns all unresolved incidents for the status page
func FetchUnresolvedIncidents(apiKey, hostName, pageID string) ([]Incident, error) {
	incidents := make([]Incident, 0)
	req, err := http.NewRequest("GET", hostName+"/v1/pages/"+pageID+"/incidents/unresolved", nil)
	if err != nil {
		return incidents, err
	}
	req.Header.Set("Authorization", "OAuth "+apiKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return incidents, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return incidents, err
	}
	response := fmt.Sprintf("%s", b)
	logger.Debug("Response from fetch all unresolved ", zap.String("Response", response))
	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(b, &incidents)
		if err != nil {
			return incidents, err
		}
	} else {
		logger.Debug(string(b))
	}
	return incidents, nil
}
