package component

import (
	"github.com/bhanurp/status-page/common"
)

func FetchAllComponents() ([]Component, error) {
	_, statusPageID, _, hostName := common.FetchStatusPageDetails()
	return fetchAllComponents(hostName, statusPageID)
}

func FetchComponentByComponentID(componentID string) (*Component, error) {
	_, statusPageID, _, hostName := common.FetchStatusPageDetails()
	return fetchComponentByComponentID(hostName, statusPageID, componentID)
}

func FetchComponentByComponentName(componentName string) (*Component, error) {
	_, statusPageID, _, hostName := common.FetchStatusPageDetails()
	return fetchComponentByComponentName(hostName, statusPageID, componentName)
}

func CreateComponent(name, description, status string) (*Component, error) {

	newComponent := Component{
		Name:        name,
		Description: description,
		Status:      status,
	}
	_, statusPageID, _, hostName := common.FetchStatusPageDetails()
	return createComponent(hostName, statusPageID, newComponent)
}

func DeleteComponent(componentID string) error {
	_, statusPageID, _, _ := common.FetchStatusPageDetails()
	return deleteComponent(statusPageID, componentID)
}
