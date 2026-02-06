package enum

type EinvoiceStatus string

const (
	EinvoiceStatusScheduled              EinvoiceStatus = "scheduled"
	EinvoiceStatusSkipped                EinvoiceStatus = "skipped"
	EinvoiceStatusInProgress             EinvoiceStatus = "in_progress"
	EinvoiceStatusSuccess                EinvoiceStatus = "success"
	EinvoiceStatusFailed                 EinvoiceStatus = "failed"
	EinvoiceStatusRegistered             EinvoiceStatus = "registered"
	EinvoiceStatusAccepted               EinvoiceStatus = "accepted"
	EinvoiceStatusRejected               EinvoiceStatus = "rejected"
	EinvoiceStatusMessageAcknowledgement EinvoiceStatus = "message_acknowledgement"
	EinvoiceStatusInProcess              EinvoiceStatus = "in_process"
	EinvoiceStatusUnderQuery             EinvoiceStatus = "under_query"
	EinvoiceStatusConditionallyAccepted  EinvoiceStatus = "conditionally_accepted"
	EinvoiceStatusPaid                   EinvoiceStatus = "paid"
)
