package enum

type WarningReasons string

const (
	WarningReasonsLowCoverage     WarningReasons = "low_coverage"
	WarningReasonsHighCardinality WarningReasons = "high_cardinality"
)
