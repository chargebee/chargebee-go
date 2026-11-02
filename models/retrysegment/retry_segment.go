package retrysegment

import (
	"encoding/json"
	retrySegmentEnum "github.com/chargebee/chargebee-go/v3/models/retrysegment/enum"
)

type RetrySegment struct {
	Id             string                          `json:"id"`
	Name           string                          `json:"name"`
	InvoiceId      string                          `json:"invoice_id"`
	Enabled        bool                            `json:"enabled"`
	Priority       int32                           `json:"priority"`
	ConditionLogic retrySegmentEnum.ConditionLogic `json:"condition_logic"`
	Criteria       json.RawMessage                 `json:"criteria"`
	RetrySettings  json.RawMessage                 `json:"retry_settings"`
	Object         string                          `json:"object"`
}
type RetrieveRequestParams struct {
	InvoiceId string `json:"invoice_id,omitempty"`
}
