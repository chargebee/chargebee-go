package enum

type Status string

const (
	StatusInProgress     Status = "in_progress"
	StatusSuccess        Status = "success"
	StatusVoided         Status = "voided"
	StatusFailure        Status = "failure"
	StatusTimeout        Status = "timeout"
	StatusNeedsAttention Status = "needs_attention"
)
