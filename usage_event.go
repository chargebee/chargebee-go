package chargebee

type UsageEvent struct {
	SubscriptionId  string          `json:"subscription_id"`
	DeduplicationId string          `json:"deduplication_id"`
	UsageTimestamp  int64           `json:"usage_timestamp"`
	Properties      json.RawMessage `json:"properties"`
	Object          string          `json:"object"`
}
type CreateRequest struct {
	DeduplicationId string                 `json:"deduplication_id"`
	SubscriptionId  string                 `json:"subscription_id"`
	UsageTimestamp  *int64                 `json:"usage_timestamp"`
	Properties      map[string]interface{} `json:"properties"`
}
type BatchIngestRequest struct {
	Events []*BatchIngestEvent `json:"events,omitempty"`
}
type BatchIngestEvent struct {
	DeduplicationId string                 `json:"deduplication_id"`
	SubscriptionId  string                 `json:"subscription_id"`
	UsageTimestamp  *int64                 `json:"usage_timestamp"`
	Properties      map[string]interface{} `json:"properties"`
}

type CreateResponse struct {
	UsageEvent *UsageEvent `json:"usage_event,omitempty"`
}

type BatchIngestResponse struct {
	BatchId      string   `json:"batch_id,omitempty"`
	FailedEvents []string `json:"failed_events,omitempty"`
}
