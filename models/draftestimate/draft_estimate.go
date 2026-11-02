package draftestimate

import (
	"encoding/json"
	draftEstimateEnum "github.com/chargebee/chargebee-go/v3/models/draftestimate/enum"
)

type DraftEstimate struct {
	Status          draftEstimateEnum.Status `json:"status"`
	HttpStatus      int32                    `json:"http_status"`
	CreatedAt       int64                    `json:"created_at"`
	Result          json.RawMessage          `json:"result"`
	ErrorDetail     json.RawMessage          `json:"error_detail"`
	RevisionVersion int64                    `json:"revision_version"`
	Object          string                   `json:"object"`
}
