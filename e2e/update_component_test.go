package e2e

import (
	"testing"

	"github.com/bhanurp/status-page/component"
	"github.com/bhanurp/status-page/e2e/utils"
)

func TestCreateComponentAndUpdate(t *testing.T) {
	utils.Inite2e()
	statusPageID := utils.FetchStatusPageID()
	// Create a test component
	testComponent := component.Component{
		Name:        "Test Component",
		Description: "This is a test component",
		Status:      "operational",
	}
	createdComponent, err := component.CreateComponent(testComponent.Name, testComponent.Description, testComponent.Status)
	if err != nil {
		t.Fatalf("Failed to create component: %v", err)
	}
	// Collect the component ID
	componentID := createdComponent.ID
	updatedComponent, err := component.UpdateComponent(componentID, "Updated Component", "This is an updated component", "degraded_performance", "", statusPageID)
	if err != nil {
		t.Fatalf("Failed to update component: %v", err)
	}
	if updatedComponent.Name != "Updated Component" {
		t.Fatalf("Failed to update component name")
	}

}
