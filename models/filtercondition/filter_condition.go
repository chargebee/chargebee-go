package filtercondition

import (
	filterConditionEnum "github.com/chargebee/chargebee-go/v3/models/filtercondition/enum"
)

type FilterCondition struct {
	Field    filterConditionEnum.Field    `json:"field"`
	Operator filterConditionEnum.Operator `json:"operator"`
	Value    string                       `json:"value"`
	Object   string                       `json:"object"`
}
