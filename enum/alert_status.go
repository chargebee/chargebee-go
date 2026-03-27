package enum

type AlertStatus string

const (
	AlertStatusWithinLimit AlertStatus = "within_limit"
	AlertStatusInAlarm     AlertStatus = "in_alarm"
)
