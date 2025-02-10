package usageevent

import (
	"encoding/json"
)

type UsageEvent struct {
	SubscriptionId  string          `json:"subscription_id"`
	DeduplicationId string          `json:"deduplication_id"`
	UsageTimestamp  int64           `json:"usage_timestamp"`
	Properties      json.RawMessage `json:"properties"`
	Object          string          `json:"object"`
}
type CreateRequestParams struct {
	DeduplicationId string                 `json:"deduplication_id"`
	SubscriptionId  string                 `json:"subscription_id"`
	UsageTimestamp  *int64                 `json:"usage_timestamp"`
	Properties      map[string]interface{} `json:"properties"`
}
type BatchIngestRequestParams struct {
	Events []*BatchIngestEventParams `json:"events,omitempty"`
}
type BatchIngestEventParams struct {
	DeduplicationId string                 `json:"deduplication_id"`
	SubscriptionId  string                 `json:"subscription_id"`
	UsageTimestamp  *int64                 `json:"usage_timestamp"`
	Properties      map[string]interface{} `json:"properties"`
}
