package incident

import (
	"encoding/json"
	"errors"
	"github.com/bhanurp/status-page/logger"
	"io"
	"net/http"

	"go.uber.org/zap"
)

type StatusPageHTTPClient struct {
	req *http.Request
}

func (s *StatusPageHTTPClient) SendHTTPRequest(incident string) error {
	resp, err := http.DefaultClient.Do(s.req)
	logger.Debug("incident id " + incident + " " + resp.Status)
	if err != nil {
		logger.Warn("Failed to resolve incident: " + incident)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	if resp.StatusCode >= http.StatusOK && resp.StatusCode < 300 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		updatedIncident := Incident{}
		err = json.Unmarshal(bodyBytes, &updatedIncident)
		if err != nil {
			return err
		}
		// Set output variables
		SetOutputVariables(updatedIncident)
		logger.Info("Updated task incident",
			zap.String("incidentID", updatedIncident.ID),
			zap.String("incidentURL", updatedIncident.Shortlink),
		)
	} else {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		logger.Debug("incident id " + incident + " " + string(bodyBytes))
		logger.Error("Received failure status code",
			zap.Int("statusCode", resp.StatusCode),
			zap.String("message", "unable to update incident"),
		)
		return errors.New("failed to update incident")
	}
	return nil
}
