package webhook

import (
	"encoding/json"
	"time"
)

// WebhookEvent represents the standard structure of all Chargebee webhook events
type WebhookEvent struct {
	Id         string          `json:"id"`
	OccurredAt int64           `json:"occurred_at"`
	Source     string          `json:"source"`
	User       string          `json:"user"`
	EventType  string          `json:"event_type"`
	ApiVersion string          `json:"api_version"`
	Content    json.RawMessage `json:"content"`
}

// GetOccurredAtTime returns the occurred_at timestamp as a time.Time
func (e *WebhookEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// ParseWebhookEvent parses a webhook payload into a WebhookEvent
func ParseWebhookEvent(payload []byte) (*WebhookEvent, error) {
	var event WebhookEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

// ParseContent unmarshals the event content into the provided struct
func (e *WebhookEvent) ParseContent(v interface{}) error {
	return json.Unmarshal(e.Content, v)
}
