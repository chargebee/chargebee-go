package businessrulesetrule

type BusinessRulesetRule struct {
	RuleId   string `json:"rule_id"`
	Priority int32  `json:"priority"`
	Object   string `json:"object"`
}
