package enum

type ChangeType string

const (
	ChangeTypeDowngrade ChangeType = "downgrade"
	ChangeTypePause     ChangeType = "pause"
)
