package rule

import (
	ruleEnum "github.com/chargebee/chargebee-go/v3/models/rule/enum"
)

type Rule struct {
	Id         string          `json:"id"`
	Namespace  string          `json:"namespace"`
	RuleName   string          `json:"rule_name"`
	RuleOrder  int32           `json:"rule_order"`
	Status     ruleEnum.Status `json:"status"`
	Conditions string          `json:"conditions"`
	Outcome    string          `json:"outcome"`
	Deleted    bool            `json:"deleted"`
	CreatedAt  int64           `json:"created_at"`
	ModifiedAt int64           `json:"modified_at"`
	Object     string          `json:"object"`
}
