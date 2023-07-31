package enum

type ProrationType string

const (
	ProrationTypeSiteDefault ProrationType = "site_default"
	ProrationTypePartialTerm ProrationType = "partial_term"
	ProrationTypeFullTerm    ProrationType = "full_term"
)
