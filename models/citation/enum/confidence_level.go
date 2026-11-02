package enum

type ConfidenceLevel string

const (
	ConfidenceLevelLow         ConfidenceLevel = "low"
	ConfidenceLevelMedium      ConfidenceLevel = "medium"
	ConfidenceLevelHigh        ConfidenceLevel = "high"
	ConfidenceLevelNotProvided ConfidenceLevel = "not_provided"
)
