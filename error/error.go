package error

import (
	"encoding/json"
	"fmt"
)

// NotionError represents error response of Notion
type NotionError struct {
	Status  int    `json:"status"`
	Code    Code   `json:"code"`
	Message string `json:"message"`
	Object  string `json:"object"`
}

// Error returns error message which is contained in Notion API response
func (ne *NotionError) Error() string {
	return fmt.Sprintf("%s: %s", ne.Code, ne.Message)
}

// Parse returns NotionError which is parsed from response json bytes
func Parse(response []byte) (error, error) {
	parsed := NotionError{}
	if err := json.Unmarshal(response, &parsed); err != nil {
		return nil, err
	}

	return &parsed, nil
}
