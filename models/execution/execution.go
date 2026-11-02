package execution

import (
	"encoding/json"
	executionEnum "github.com/chargebee/chargebee-go/v3/models/execution/enum"
)

type Execution struct {
	Status          executionEnum.Status `json:"status"`
	CreatedAt       int64                `json:"created_at"`
	Result          json.RawMessage      `json:"result"`
	ErrorDetail     json.RawMessage      `json:"error_detail"`
	RevisionVersion int64                `json:"revision_version"`
	Object          string               `json:"object"`
}
