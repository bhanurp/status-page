package incident

import (
	"encoding/json"
	"net/http"

	"github.com/bhanurp/status-page/logger"
	"go.uber.org/zap"
)

// Incident represents an incident in the status page
type DeleteIncident struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

// DeleteIncidentHandler handles the deletion of an incident
func DeleteIncidentHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var incident Incident
	err := json.NewDecoder(r.Body).Decode(&incident)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Delete the incident from the status page
	err = deleteIncident(incident.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
}

// deleteIncident deletes an incident from the status page
func deleteIncident(incidentID string) error {
	logger.Info("Deleting incident", zap.String("incidentID", incidentID))
	// Implement the logic to delete the incident from the status page
	// For example, you can delete the incident from a database or a file

	// Delete the incident with the given ID
	// ...

	return nil
}
