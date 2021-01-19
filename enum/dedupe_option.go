package enum

type DedupeOption string

const (
	DedupeOptionSkip           DedupeOption = "skip"
	DedupeOptionUpdateExisting DedupeOption = "update_existing"
)
