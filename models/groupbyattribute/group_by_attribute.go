package groupbyattribute

import (
	groupByAttributeEnum "github.com/chargebee/chargebee-go/v3/models/groupbyattribute/enum"
)

type GroupByAttribute struct {
	AttributeName       string                              `json:"attribute_name"`
	Rank                int32                               `json:"rank"`
	CoveragePercent     int32                               `json:"coverage_percent"`
	Cardinality         int64                               `json:"cardinality"`
	RemainingValueCount int64                               `json:"remaining_value_count"`
	Warning             bool                                `json:"warning"`
	SampleValues        []string                            `json:"sample_values"`
	WarningReasons      groupByAttributeEnum.WarningReasons `json:"warning_reasons"`
	Object              string                              `json:"object"`
}
