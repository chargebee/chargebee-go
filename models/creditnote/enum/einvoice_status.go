package enum

type EinvoiceStatus string

const (
	EinvoiceStatusScheduled  EinvoiceStatus = "scheduled"
	EinvoiceStatusSkipped    EinvoiceStatus = "skipped"
	EinvoiceStatusInProgress EinvoiceStatus = "in_progress"
	EinvoiceStatusSuccess    EinvoiceStatus = "success"
	EinvoiceStatusFailed     EinvoiceStatus = "failed"
)
