package enum

type PiiCleared string

const (
	PiiClearedActive            PiiCleared = "active"
	PiiClearedScheduledForClear PiiCleared = "scheduled_for_clear"
	PiiClearedCleared           PiiCleared = "cleared"
)
