package chargebee

import (
	"encoding/json"
)

// just struct
type UsageEvent struct {
	SubscriptionId  string          `json:"subscription_id"`
	DeduplicationId string          `json:"deduplication_id"`
	UsageTimestamp  int64           `json:"usage_timestamp"`
	Properties      json.RawMessage `json:"properties"`
	Object          string          `json:"object"`
}

// sub resources
// operations
// input params
type UsageEventCreateRequest struct {
	DeduplicationId string                 `json:"deduplication_id"`
	SubscriptionId  string                 `json:"subscription_id"`
	UsageTimestamp  *int64                 `json:"usage_timestamp"`
	Properties      map[string]interface{} `json:"properties"`
	apiRequest      `json:"-" form:"-"`
}

func (r *UsageEventCreateRequest) payload() any { return r }

type UsageEventBatchIngestRequest struct {
	Events     []*UsageEventBatchIngestEvent `json:"events,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *UsageEventBatchIngestRequest) payload() any { return r }

// input sub resource params multi
type UsageEventBatchIngestEvent struct {
	DeduplicationId string                 `json:"deduplication_id"`
	SubscriptionId  string                 `json:"subscription_id"`
	UsageTimestamp  *int64                 `json:"usage_timestamp"`
	Properties      map[string]interface{} `json:"properties"`
}

// operation response
type UsageEventCreateResponse struct {
	UsageEvent *UsageEvent `json:"usage_event,omitempty"`
	apiResponse
}

// operation response
type UsageEventBatchIngestResponse struct {
	BatchId      string   `json:"batch_id,omitempty"`
	FailedEvents []string `json:"failed_events,omitempty"`
	apiResponse
}
