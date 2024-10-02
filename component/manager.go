package component

import (
	"github.com/bhanurp/status-page/common"
)

func FetchAllComponents() ([]Component, error) {
	statusPageID := common.FetchStatusPageID()
	return fetchAllComponents(statusPageID)
}

func FetchComponentByComponentID(componentID string) (*Component, error) {
	statusPageID := common.FetchStatusPageID()
	return fetchComponentByComponentID(statusPageID, componentID)
}

func FetchComponentByComponentName(componentName string) (*Component, error) {
	statusPageID := common.FetchStatusPageID()
	return fetchComponentByComponentName(statusPageID, componentName)
}

func CreateComponent(name, description, status string) (*Component, error) {
	newComponent := Component{
		Name:        name,
		Description: description,
		Status:      status,
	}
	statusPageID := common.FetchStatusPageID()
	return createComponent(statusPageID, newComponent)
}

func DeleteComponent(componentID string) error {
	statusPageID := common.FetchStatusPageID()
	return deleteComponent(statusPageID, componentID)
}
