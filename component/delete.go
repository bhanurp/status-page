package component

import (
	"fmt"
	"net/http"

	"github.com/bhanurp/rest"
	"github.com/bhanurp/status-page/common"
	"github.com/bhanurp/status-page/logger"
	"github.com/bhanurp/status-page/statuspageurl"
	"go.uber.org/zap"
)

func deleteComponent(pageID, componentID string) error {
	logger.Debug("Deleting component", zap.String("componentID", componentID))
	componentURL := statuspageurl.ConstructURL("pages/%s/components/%s", pageID, componentID)
	commonHeaders := common.CreateHeaders()
	restClient := rest.DeleteRequest{}
	resp, err := restClient.Do(componentURL, []byte{}, commonHeaders, 10)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete component: %s", resp.Body)
	}
	fmt.Println("Component deleted successfully")
	return nil
}
