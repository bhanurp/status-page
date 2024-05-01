package incident

import (
	"encoding/json"
	"status-page/logger"
)

// createIncidentMetaData creates incident meta data with all pipelines the component
// is failing. Should be updated with all pipelines
func createIncidentMetaData(incidentMetadata Metadata) []byte {
	j, err := json.Marshal(incidentMetadata)
	if err != nil {
		logger.Error("Failed marshalling pipelines into json data")
		return []byte{}
	}
	return j
}
