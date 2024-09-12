package incident

import (
	"encoding/json"
	"status-page/logger"
)

// createIncidentMetaData creates incident meta data
func createIncidentMetaData(incidentMetadata Metadata) []byte {
	j, err := json.Marshal(incidentMetadata)
	if err != nil {
		logger.Error("Failed marshalling metadata into json data")
		return []byte{}
	}
	return j
}
