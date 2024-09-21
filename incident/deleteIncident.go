package incident

import (
	"github.com/bhanurp/rest"
	"github.com/bhanurp/status-page/logger"
	"go.uber.org/zap"
)

// Incident represents an incident in the status page
type DeleteIncident struct {
	ID      string
	Title   string
	Message string
	APIKey  string
	PageID  string
}

func (d *DeleteIncident) SetTitle(title string) *DeleteIncident {
	d.Title = title
	return d
}

func (d *DeleteIncident) SetMessage(message string) *DeleteIncident {
	d.Message = message
	return d
}

func (d *DeleteIncident) SetAPIKey(apiKey string) *DeleteIncident {
	d.APIKey = apiKey
	return d
}

func (d *DeleteIncident) SetPageID(pageID string) *DeleteIncident {
	d.PageID = pageID
	return d
}

func BuildDeleteIncident() *DeleteIncident {
	return new(DeleteIncident)
}

func (d *DeleteIncident) SetID(id string) *DeleteIncident {
	d.ID = id
	return d
}

// DeleteIncidentHandler handles the deletion of an incident
func (d *DeleteIncident) DeleteIncidentHandler() error {
	headers := createHeaders(d.APIKey)

	// Delete the incident from the status page
	responseBody, err := deleteIncident(d.ID, d.PageID, headers)
	if err != nil {
		logger.Debug("response", zap.String("body", string(responseBody)))
		return err
	}
	logger.Info("Incident deleted successfully", zap.String("incidentID", d.ID))
	return nil
}

// deleteIncident deletes an incident from the status page
func deleteIncident(incidentID, pageID string, headers map[string]string) ([]byte, error) {
	logger.Info("Deleting incident", zap.String("incidentID", incidentID))
	delete := rest.DeleteRequest{}
	deleteResponse, err := delete.Do(createStatusPageURL(pageID, incidentID), []byte{}, headers, 10)
	return deleteResponse.Body, err
}
