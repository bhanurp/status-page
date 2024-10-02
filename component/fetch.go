package component

import (
	"encoding/json"
	"fmt"

	"github.com/bhanurp/rest"
	"github.com/bhanurp/status-page/common"
	"github.com/bhanurp/status-page/logger"
	"github.com/bhanurp/status-page/statuspageurl"
	"go.uber.org/zap"
)

func fetchComponentByComponentID(statusPageID, componentID string) (*Component, error) {
	components, err := fetchAllComponents(statusPageID)
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

func fetchComponentByComponentName(statusPageID, componentName string) (*Component, error) {
	components, err := fetchAllComponents(statusPageID)
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

func fetchAllComponents(statusPageID string) ([]Component, error) {
	url := statuspageurl.ConstructURL("pages/%s/components", statusPageID)
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
