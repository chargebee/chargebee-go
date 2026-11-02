package collectionsremindersequence

import (
	collectionsReminderSequenceEnum "github.com/chargebee/chargebee-go/v3/models/collectionsremindersequence/enum"
)

type CollectionsReminderSequence struct {
	Id              string                                 `json:"id"`
	InvoiceId       string                                 `json:"invoice_id"`
	CustomerId      string                                 `json:"customer_id"`
	Status          collectionsReminderSequenceEnum.Status `json:"status"`
	PausedUntil     int64                                  `json:"paused_until"`
	PauseNote       string                                 `json:"pause_note"`
	StartedAt       int64                                  `json:"started_at"`
	CompletedAt     int64                                  `json:"completed_at"`
	CreatedAt       int64                                  `json:"created_at"`
	UpdatedAt       int64                                  `json:"updated_at"`
	ResourceVersion int64                                  `json:"resource_version"`
	ActionLogs      []*ActionLog                           `json:"action_logs"`
	Object          string                                 `json:"object"`
}
type ActionLog struct {
	ActionType  collectionsReminderSequenceEnum.CollectionsReminderSequenceActionLogActionType `json:"action_type"`
	Status      collectionsReminderSequenceEnum.CollectionsReminderSequenceActionLogStatus     `json:"status"`
	ReasonCode  collectionsReminderSequenceEnum.CollectionsReminderSequenceActionLogReasonCode `json:"reason_code"`
	ExecutedAt  int64                                                                          `json:"executed_at"`
	ScheduledAt int64                                                                          `json:"scheduled_at"`
	Object      string                                                                         `json:"object"`
}
