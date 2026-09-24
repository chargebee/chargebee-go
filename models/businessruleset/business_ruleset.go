package businessruleset

import (
	"encoding/json"
	"github.com/chargebee/chargebee-go/v3/filter"
	businessRulesetEnum "github.com/chargebee/chargebee-go/v3/models/businessruleset/enum"
)

type BusinessRuleset struct {
	Id              string                          `json:"id"`
	Name            string                          `json:"name"`
	Description     string                          `json:"description"`
	Active          bool                            `json:"active"`
	ExecuteMode     businessRulesetEnum.ExecuteMode `json:"execute_mode"`
	UpdatedAt       int64                           `json:"updated_at"`
	UpdatedBy       string                          `json:"updated_by"`
	CreatedBy       string                          `json:"created_by"`
	CreatedAt       int64                           `json:"created_at"`
	Rules           json.RawMessage                 `json:"rules"`
	ResourceVersion int64                           `json:"resource_version"`
	Object          string                          `json:"object"`
}
type CreateRequestParams struct {
	Id          string                          `json:"id,omitempty"`
	Name        string                          `json:"name"`
	Description string                          `json:"description,omitempty"`
	ExecuteMode businessRulesetEnum.ExecuteMode `json:"execute_mode,omitempty"`
	Rules       []interface{}                   `json:"rules,omitempty"`
}
type UpdateRequestParams struct {
	Name        string                          `json:"name"`
	Description string                          `json:"description,omitempty"`
	ExecuteMode businessRulesetEnum.ExecuteMode `json:"execute_mode,omitempty"`
	Rules       []interface{}                   `json:"rules,omitempty"`
}
type AddRulesRequestParams struct {
	Rules []interface{} `json:"rules,omitempty"`
}
type RemoveRulesRequestParams struct {
	Rules []interface{} `json:"rules,omitempty"`
}
type ListRulesRequestParams struct {
	Limit  *int32                `json:"limit,omitempty"`
	Offset string                `json:"offset,omitempty"`
	Active *filter.BooleanFilter `json:"active,omitempty"`
}
type ListRequestParams struct {
	Limit  *int32                `json:"limit,omitempty"`
	Offset string                `json:"offset,omitempty"`
	Active *filter.BooleanFilter `json:"active,omitempty"`
}
