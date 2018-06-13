package enum

type TimeTravelStatus string

const (
	TimeTravelStatusNotEnabled TimeTravelStatus = "not_enabled"
	TimeTravelStatusInProgress TimeTravelStatus = "in_progress"
	TimeTravelStatusSucceeded  TimeTravelStatus = "succeeded"
	TimeTravelStatusFailed     TimeTravelStatus = "failed"
)
