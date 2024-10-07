package component

import (
	"encoding/json"

	"github.com/bhanurp/rest"
	"github.com/bhanurp/status-page/common"
	"github.com/bhanurp/status-page/logger"
	"github.com/bhanurp/status-page/statuspageurl"
)

func updateComponent(updateComponent Component) (*Component, error) {
	restClient := rest.PutRequest{}
	url := statuspageurl.ConstructURL("pages/%s/components/%s", updateComponent.PageID, updateComponent.ID)
	jsonData, err := fetchComponentData(updateComponent)
	if err != nil {
		return nil, err
	}
	resp, err := restClient.Do(url, jsonData, common.CreateHeaders(), 10)
	if err != nil {
		return nil, err
	}
	logger.Debug("Updated component successfully")
	var updatedComponent Component
	err = json.Unmarshal(resp.Body, &updatedComponent)
	if err != nil {
		return nil, err
	}
	return &updatedComponent, nil
}
