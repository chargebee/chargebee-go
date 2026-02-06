package enum

type Status string

const (
	StatusScheduled              Status = "scheduled"
	StatusSkipped                Status = "skipped"
	StatusInProgress             Status = "in_progress"
	StatusSuccess                Status = "success"
	StatusFailed                 Status = "failed"
	StatusRegistered             Status = "registered"
	StatusAccepted               Status = "accepted"
	StatusRejected               Status = "rejected"
	StatusMessageAcknowledgement Status = "message_acknowledgement"
	StatusInProcess              Status = "in_process"
	StatusUnderQuery             Status = "under_query"
	StatusConditionallyAccepted  Status = "conditionally_accepted"
	StatusPaid                   Status = "paid"
)
