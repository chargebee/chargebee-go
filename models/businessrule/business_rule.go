package businessrule

import (
	"encoding/json"
	"github.com/chargebee/chargebee-go/v3/filter"
)

type BusinessRule struct {
	Id                   string          `json:"id"`
	Name                 string          `json:"name"`
	Description          string          `json:"description"`
	LatestVersion        int32           `json:"latest_version"`
	Active               bool            `json:"active"`
	ReleasedAt           int64           `json:"released_at"`
	ReleasedBy           string          `json:"released_by"`
	UpdatedAt            int64           `json:"updated_at"`
	UpdatedBy            string          `json:"updated_by"`
	CreatedBy            string          `json:"created_by"`
	CreatedAt            int64           `json:"created_at"`
	Tags                 json.RawMessage `json:"tags"`
	StructuredExpression json.RawMessage `json:"structured_expression"`
	ActionsOnSuccess     json.RawMessage `json:"actions_on_success"`
	ResourceVersion      int64           `json:"resource_version"`
	Object               string          `json:"object"`
}
type CreateRequestParams struct {
	Id                   string                 `json:"id,omitempty"`
	Name                 string                 `json:"name"`
	Description          string                 `json:"description,omitempty"`
	Tags                 []interface{}          `json:"tags,omitempty"`
	StructuredExpression map[string]interface{} `json:"structured_expression"`
	ActionsOnSuccess     []interface{}          `json:"actions_on_success,omitempty"`
}
type UpdateDraftRequestParams struct {
	Name                 string                 `json:"name"`
	Description          string                 `json:"description,omitempty"`
	Tags                 []interface{}          `json:"tags,omitempty"`
	StructuredExpression map[string]interface{} `json:"structured_expression"`
	ActionsOnSuccess     []interface{}          `json:"actions_on_success,omitempty"`
}
type ListRequestParams struct {
	Limit  *int32                `json:"limit,omitempty"`
	Offset string                `json:"offset,omitempty"`
	Draft  *filter.BooleanFilter `json:"draft,omitempty"`
	Active *filter.BooleanFilter `json:"active,omitempty"`
}
type ApplyRulesRequestParams struct {
	Evaluate             *bool                  `json:"evaluate,omitempty"`
	RuleId               string                 `json:"rule_id,omitempty"`
	RulesetId            string                 `json:"ruleset_id,omitempty"`
	SkipFailedRules      *bool                  `json:"skip_failed_rules,omitempty"`
	StructuredExpression map[string]interface{} `json:"structured_expression,omitempty"`
	Context              map[string]interface{} `json:"context,omitempty"`
}
