package component

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bhanurp/rest"
	"github.com/bhanurp/status-page/common"
	"github.com/bhanurp/status-page/logger"
	"github.com/bhanurp/status-page/statuspageurl"
)

// createComponent creates a new component in the StatusPage API
func createComponent(statusPageID string, newComponent Component) (*Component, error) {
	var err error
	url := statuspageurl.ConstructURL("pages/%s/components", statusPageID)
	jsonData, err := fetchComponentData(newComponent)
	if err != nil {
		return nil, err
	}
	postRequest := &rest.PostRequest{}
	resp, createErr := postRequest.Do(url, jsonData, common.CreateHeaders(), 10)
	if createErr != nil {
		return nil, createErr
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create component: %d", resp.StatusCode)
	}
	var createdComponent Component
	err = json.Unmarshal(resp.Body, &createdComponent)
	if err != nil {
		return nil, err
	}

	logger.Info("Created component successfully")
	return &createdComponent, nil
}
