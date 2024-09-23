package component

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bhanurp/rest"
	"github.com/bhanurp/status-page/common"
	"github.com/bhanurp/status-page/logger"
	"go.uber.org/zap"
)

func fetchComponentByComponentID(hostName, statusPageID, componentID string) (*Component, error) {
	components, err := fetchAllComponents(hostName, statusPageID)
	if err != nil {
		return nil, err
	}
	for i, component := range components {
		if component.ID == componentID {
			return &components[i], nil
		}
	}
	return nil, fmt.Errorf("component not found with ID: %s", componentID)
}

func fetchComponentByComponentName(hostName, statusPageID, componentName string) (*Component, error) {
	components, err := fetchAllComponents(hostName, statusPageID)
	if err != nil {
		return nil, err
	}
	for i, component := range components {
		if component.Name == componentName {
			return &components[i], nil
		}
	}
	return nil, fmt.Errorf("component not found with name: %s", componentName)
}

func fetchAllComponents(hostName, statusPageID string) ([]Component, error) {
	url := fmt.Sprintf("https://%s/v1/pages/%s/components", hostName, statusPageID)
	headers := common.CreateHeaders()
	getComponents := rest.GetRequest{}
	resp, err := getComponents.Do(url, []byte{}, headers, 10)
	if err != nil {
		logger.Debug("response", zap.String("body", string(resp.Body)))
		return nil, err
	}

	var components []Component
	err = json.Unmarshal(resp.Body, &components)
	if err != nil {
		return nil, err
	}

	return components, nil
}

// createComponent creates a new component in the StatusPage API
func createComponent(hostName, statusPageID string, newComponent Component) (*Component, error) {
	var err error
	url := fmt.Sprintf("https://%s/v1/pages/%s/components", hostName, statusPageID)
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

	return &createdComponent, nil
}

func fetchComponentData(newComponent Component) ([]byte, error) {
	componentData := map[string]map[string]string{
		"component": {
			"name":        newComponent.Name,
			"description": newComponent.Description,
			"status":      newComponent.Status,
		},
	}

	jsonData, err := json.Marshal(componentData)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
