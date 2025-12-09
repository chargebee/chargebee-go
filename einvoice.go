package chargebee

type Status string

const (
	StatusScheduled  Status = "scheduled"
	StatusSkipped    Status = "skipped"
	StatusInProgress Status = "in_progress"
	StatusSuccess    Status = "success"
	StatusFailed     Status = "failed"
	StatusRegistered Status = "registered"
)

type Einvoice struct {
	Id              string `json:"id"`
	ReferenceNumber string `json:"reference_number"`
	Status          Status `json:"status"`
	Message         string `json:"message"`
	Object          string `json:"object"`
}
