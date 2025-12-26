package chargebee

import (
	"encoding/json"
)

type EventWebhookStatus string

const (
	EventWebhookStatusNotConfigured EventWebhookStatus = "not_configured"
	EventWebhookStatusScheduled     EventWebhookStatus = "scheduled"
	EventWebhookStatusSucceeded     EventWebhookStatus = "succeeded"
	EventWebhookStatusReScheduled   EventWebhookStatus = "re_scheduled"
	EventWebhookStatusFailed        EventWebhookStatus = "failed"
	EventWebhookStatusSkipped       EventWebhookStatus = "skipped"
	EventWebhookStatusNotApplicable EventWebhookStatus = "not_applicable"
	EventWebhookStatusDisabled      EventWebhookStatus = "disabled"
)

type EventWebhookWebhookStatus string

const (
	EventWebhookWebhookStatusNotConfigured EventWebhookWebhookStatus = "not_configured"
	EventWebhookWebhookStatusScheduled     EventWebhookWebhookStatus = "scheduled"
	EventWebhookWebhookStatusSucceeded     EventWebhookWebhookStatus = "succeeded"
	EventWebhookWebhookStatusReScheduled   EventWebhookWebhookStatus = "re_scheduled"
	EventWebhookWebhookStatusFailed        EventWebhookWebhookStatus = "failed"
	EventWebhookWebhookStatusSkipped       EventWebhookWebhookStatus = "skipped"
	EventWebhookWebhookStatusNotApplicable EventWebhookWebhookStatus = "not_applicable"
	EventWebhookWebhookStatusDisabled      EventWebhookWebhookStatus = "disabled"
)

// just struct
type Event struct {
	Id         string          `json:"id"`
	OccurredAt int64           `json:"occurred_at"`
	Source     Source          `json:"source"`
	User       string          `json:"user"`
	Webhooks   []*EventWebhook `json:"webhooks"`
	EventType  EventType       `json:"event_type"`
	ApiVersion ApiVersion      `json:"api_version"`
	Content    json.RawMessage `json:"content"`
	OriginUser string          `json:"origin_user"`
	Object     string          `json:"object"`
}

// sub resources
type EventWebhook struct {
	Id            string                    `json:"id"`
	WebhookStatus EventWebhookWebhookStatus `json:"webhook_status"`
	Object        string                    `json:"object"`
}

// operations
// input params
type EventListRequest struct {
	Limit         *int32           `json:"limit,omitempty"`
	Offset        string           `json:"offset,omitempty"`
	Id            *StringFilter    `json:"id,omitempty"`
	WebhookStatus *EnumFilter      `json:"webhook_status,omitempty"`
	EventType     *EnumFilter      `json:"event_type,omitempty"`
	Source        *EnumFilter      `json:"source,omitempty"`
	OccurredAt    *TimestampFilter `json:"occurred_at,omitempty"`
	SortBy        *SortFilter      `json:"sort_by,omitempty"`
	apiRequest    `json:"-" form:"-"`
}

func (r *EventListRequest) payload() any { return r }

// operation sub response
type EventListEventResponse struct {
	Event *Event `json:"event,omitempty"`
}

// operation response
type EventListResponse struct {
	List       []*EventListEventResponse `json:"list,omitempty"`
	NextOffset string                    `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type EventRetrieveResponse struct {
	Event *Event `json:"event,omitempty"`
	apiResponse
}
