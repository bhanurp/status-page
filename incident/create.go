package incident

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/bhanurp/rest"
	"github.com/bhanurp/status-page/logger"
	"github.com/bhanurp/status-page/statuspageurl"

	"go.uber.org/zap"
)

// NewDefaultIncident creates incident struct and returns the pointer to it
func NewDefaultIncident(apiKey, componentID, pageID, incidentName, incidentBody string) *CreateIncident {
	c := new(CreateIncident)
	c.APIKey = apiKey
	c.ComponentID = componentID
	c.PageID = pageID
	c.IncidentName = incidentName
	c.IncidentBody = incidentBody
	c.IncidentStatus = IncidentStatusIdentified
	c.Metadata = Metadata{}
	return c
}

func BuildCreateIncident() *CreateIncident {
	return new(CreateIncident)
}

// SetAPIKey sets the API key and returns the CreateIncident pointer
func (c *CreateIncident) SetAPIKey(apiKey string) *CreateIncident {
	c.APIKey = apiKey
	return c
}

// SetComponentID sets the component ID and returns the CreateIncident pointer
func (c *CreateIncident) SetComponentID(componentID string) *CreateIncident {
	c.ComponentID = componentID
	return c
}

// SetPageID sets the page ID and returns the CreateIncident pointer
func (c *CreateIncident) SetPageID(pageID string) *CreateIncident {
	c.PageID = pageID
	return c
}

// SetIncidentName sets the incident name and returns the CreateIncident pointer
func (c *CreateIncident) SetIncidentName(incidentName string) *CreateIncident {
	c.IncidentName = incidentName
	return c
}

// SetIncidentBody sets the incident body and returns the CreateIncident pointer
func (c *CreateIncident) SetIncidentBody(incidentBody string) *CreateIncident {
	c.IncidentBody = incidentBody
	return c
}

// SetIncidentStatus sets the incident status and returns the CreateIncident pointer
func (c *CreateIncident) SetIncidentStatus(incidentStatus string) *CreateIncident {
	c.IncidentStatus = incidentStatus
	return c
}

// SetMetadata sets the metadata and returns the CreateIncident pointer
func (c *CreateIncident) SetMetadata(metadata Metadata) *CreateIncident {
	c.Metadata = metadata
	return c
}

// SetIncidentHeader sets the incident header and returns the CreateIncident pointer
func (c *CreateIncident) SetIncidentHeader(incidentHeader string) *CreateIncident {
	c.IncidentHeader = incidentHeader
	return c
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
	u := NewUpdateIncident(c.APIKey, c.ComponentID, c.PageID, c.IncidentBody, c.IncidentHeader, IncidentStatusIdentified)
	u.IncidentName = c.IncidentName
	u.Metadata = c.Metadata

	err := u.UpdateIncidentMatchingWithComponent(incident.ID, ComponentStatusMajorOutage)
	if err != nil {
		return err
	}
	return nil
}

// SendCreateIncidentRequest calls to create incident on status page, component CreateIncident configured.
func (c *CreateIncident) SendCreateIncidentRequest() (*Incident, error) {
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
	headers["Authorization"] = "OAuth " + c.APIKey
	headers["Content-Type"] = "application/json"
	p := rest.PostRequest{}
	resp, err := p.Do(statuspageurl.BaseURL+"pages/"+c.PageID+"/incidents", payloadBytes, headers, 10)
	if err != nil {
		return nil, err
	}
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
	return &createdIncident, err
}
