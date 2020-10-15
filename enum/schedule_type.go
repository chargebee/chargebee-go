package enum

type ScheduleType string

const (
	ScheduleTypeImmediate      ScheduleType = "immediate"
	ScheduleTypeSpecificDates  ScheduleType = "specific_dates"
	ScheduleTypeFixedIntervals ScheduleType = "fixed_intervals"
)
