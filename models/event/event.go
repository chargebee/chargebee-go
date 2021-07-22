package event

import (
	"encoding/json"
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	eventEnum "github.com/chargebee/chargebee-go/models/event/enum"
)

type Event struct {
	Id                   string                  `json:"id"`
	OccurredAt           int64                   `json:"occurred_at"`
	Source               enum.Source             `json:"source"`
	User                 string                  `json:"user"`
	WebhookStatus        eventEnum.WebhookStatus `json:"webhook_status"`
	WebhookFailureReason string                  `json:"webhook_failure_reason"`
	Webhooks             []*Webhook              `json:"webhooks"`
	EventType            enum.EventType          `json:"event_type"`
	ApiVersion           enum.ApiVersion         `json:"api_version"`
	Content              json.RawMessage         `json:"content"`
	Object               string                  `json:"object"`
}
type Webhook struct {
	Id            string                         `json:"id"`
	WebhookStatus eventEnum.WebhookWebhookStatus `json:"webhook_status"`
	Object        string                         `json:"object"`
}
type ListRequestParams struct {
	Limit         *int32                  `json:"limit,omitempty"`
	Offset        string                  `json:"offset,omitempty"`
	StartTime     *int64                  `json:"start_time,omitempty"`
	EndTime       *int64                  `json:"end_time,omitempty"`
	Id            *filter.StringFilter    `json:"id,omitempty"`
	WebhookStatus *filter.EnumFilter      `json:"webhook_status,omitempty"`
	EventType     *filter.EnumFilter      `json:"event_type,omitempty"`
	Source        *filter.EnumFilter      `json:"source,omitempty"`
	OccurredAt    *filter.TimestampFilter `json:"occurred_at,omitempty"`
	SortBy        *filter.SortFilter      `json:"sort_by,omitempty"`
}
