package enum

type ProrationType string

const (
	ProrationTypeFullTerm    ProrationType = "full_term"
	ProrationTypePartialTerm ProrationType = "partial_term"
	ProrationTypeNone        ProrationType = "none"
)
