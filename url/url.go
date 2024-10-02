package component

import (
	"fmt"
)

const (
	BaseURL        = "https://api.statuspage.io/v1/pages"
	ComponentsPath = "/components"
)

// ConstructURL constructs the full URL for the given page ID and path
func ConstructURL(pageID, path string) string {
	return fmt.Sprintf("%s/%s%s", BaseURL, pageID, path)
}
