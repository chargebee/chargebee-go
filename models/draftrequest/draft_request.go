package draftrequest

import (
	"encoding/json"
	draftRequestEnum "github.com/chargebee/chargebee-go/v3/models/draftrequest/enum"
)

type DraftRequest struct {
	Id              string                  `json:"id"`
	PathParam       string                  `json:"path_param"`
	RequestUri      string                  `json:"request_uri"`
	OperationId     string                  `json:"operation_id"`
	Status          draftRequestEnum.Status `json:"status"`
	ExecutedAt      int64                   `json:"executed_at"`
	CreatedAt       int64                   `json:"created_at"`
	ResourceVersion int64                   `json:"resource_version"`
	UpdatedAt       int64                   `json:"updated_at"`
	CurrentRevision *DraftRevision          `json:"current_revision"`
	Validations     []*Validation           `json:"validations"`
	HasComments     bool                    `json:"has_comments"`
	HasCitations    bool                    `json:"has_citations"`
	Executions      []*Execution            `json:"executions"`
	Object          string                  `json:"object"`
}
type DraftRevision struct {
	Version   int64           `json:"version"`
	CreatedAt int64           `json:"created_at"`
	Payload   json.RawMessage `json:"payload"`
	Object    string          `json:"object"`
}
type Validation struct {
	Category  string `json:"category"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	ParamName string `json:"param_name"`
	Type      string `json:"type"`
	Object    string `json:"object"`
}
type Execution struct {
	Status          draftRequestEnum.ExecutionStatus `json:"status"`
	CreatedAt       int64                            `json:"created_at"`
	Result          json.RawMessage                  `json:"result"`
	ErrorDetail     json.RawMessage                  `json:"error_detail"`
	RevisionVersion int64                            `json:"revision_version"`
	Object          string                           `json:"object"`
}
