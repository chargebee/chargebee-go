package enum

type ScheduleType string

const (
	ScheduleTypeFixedIntervals ScheduleType = "fixed_intervals"
	ScheduleTypeSpecificDates  ScheduleType = "specific_dates"
)
