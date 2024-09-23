package e2e

import (
	"log"
	"testing"

	"github.com/bhanurp/status-page/component"
	"github.com/bhanurp/status-page/e2e/utils"
)

// TestCreateAndDeleteComponent tests the creation and deletion of a component
func TestCreateAndDeleteComponent(t *testing.T) {
	utils.Inite2e()
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
	// Delete the created component
	err = component.DeleteComponent(componentID)
	if err != nil {
		t.Fatalf("Failed to delete component: %v", err)
	}

	// Validate the component is no longer available
	_, err = component.FetchComponentByComponentID(componentID)
	if err == nil {
		t.Fatalf("Component was not deleted")
	}
	log.Println("Component created and deleted successfully")
}
