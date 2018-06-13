package enum

type Status string

const (
	StatusNew        Status = "new"
	StatusProcessing Status = "processing"
	StatusComplete   Status = "complete"
	StatusCancelled  Status = "cancelled"
	StatusVoided     Status = "voided"
)
