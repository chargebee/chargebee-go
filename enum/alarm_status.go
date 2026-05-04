package enum

type AlarmStatus string

const (
	AlarmStatusWithinLimit AlarmStatus = "within_limit"
	AlarmStatusInAlarm     AlarmStatus = "in_alarm"
)
