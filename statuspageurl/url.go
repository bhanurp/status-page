package statuspageurl

import (
	"fmt"
)

const (
	BaseURL = "https://api.statuspage.io/v1/"
)

// ConstructURL constructs the full URL for the given path and parameters
func ConstructURL(path string, params ...interface{}) string {
	return fmt.Sprintf(BaseURL+path, params...)
}
