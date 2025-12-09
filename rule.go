package chargebee

type Status string

const (
	StatusActive   Status = "active"
	StatusDisabled Status = "disabled"
)

type Rule struct {
	Id         string `json:"id"`
	Namespace  string `json:"namespace"`
	RuleName   string `json:"rule_name"`
	RuleOrder  int32  `json:"rule_order"`
	Status     Status `json:"status"`
	Conditions string `json:"conditions"`
	Outcome    string `json:"outcome"`
	Deleted    bool   `json:"deleted"`
	CreatedAt  int64  `json:"created_at"`
	ModifiedAt int64  `json:"modified_at"`
	Object     string `json:"object"`
}

type RetrieveResponse struct {
	Rule *Rule `json:"rule,omitempty"`
}
