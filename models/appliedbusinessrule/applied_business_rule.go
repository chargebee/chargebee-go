package appliedbusinessrule

import (
	appliedBusinessRuleEnum "github.com/chargebee/chargebee-go/v3/models/appliedbusinessrule/enum"
)

type AppliedBusinessRule struct {
	Handle        string                             `json:"handle"`
	EntityType    appliedBusinessRuleEnum.EntityType `json:"entity_type"`
	EntityId      int64                              `json:"entity_id"`
	EntityVersion int32                              `json:"entity_version"`
	RuleId        string                             `json:"rule_id"`
	Version       int32                              `json:"version"`
	CreatedAt     int64                              `json:"created_at"`
	ModifiedAt    int64                              `json:"modified_at"`
	Object        string                             `json:"object"`
}
