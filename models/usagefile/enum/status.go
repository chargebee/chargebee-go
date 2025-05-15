package enum

type Status string

const (
	StatusQueued     Status = "queued"
	StatusImported   Status = "imported"
	StatusProcessing Status = "processing"
	StatusProcessed  Status = "processed"
	StatusFailed     Status = "failed"
)
