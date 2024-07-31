package enum

type ScheduleStatus string

const (
	ScheduleStatusActivated ScheduleStatus = "activated"
	ScheduleStatusScheduled ScheduleStatus = "scheduled"
	ScheduleStatusFailed    ScheduleStatus = "failed"
)
