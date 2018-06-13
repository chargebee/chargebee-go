package enum

type DunningStatus string

const (
	DunningStatusInProgress DunningStatus = "in_progress"
	DunningStatusExhausted  DunningStatus = "exhausted"
	DunningStatusStopped    DunningStatus = "stopped"
	DunningStatusSuccess    DunningStatus = "success"
)
