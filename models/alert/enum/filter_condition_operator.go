package enum

type FilterConditionOperator string

const (
	FilterConditionOperatorEquals    FilterConditionOperator = "equals"
	FilterConditionOperatorNotEquals FilterConditionOperator = "not_equals"
)
