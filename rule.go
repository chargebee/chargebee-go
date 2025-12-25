package chargebee

type RuleStatus string

const (
	RuleStatusActive   RuleStatus = "active"
	RuleStatusDisabled RuleStatus = "disabled"
)

// just struct
type Rule struct {
	Id         string     `json:"id"`
	Namespace  string     `json:"namespace"`
	RuleName   string     `json:"rule_name"`
	RuleOrder  int32      `json:"rule_order"`
	Status     RuleStatus `json:"status"`
	Conditions string     `json:"conditions"`
	Outcome    string     `json:"outcome"`
	Deleted    bool       `json:"deleted"`
	CreatedAt  int64      `json:"created_at"`
	ModifiedAt int64      `json:"modified_at"`
	Object     string     `json:"object"`
}

// sub resources
// operations
// input params

// operation response
type RuleRetrieveResponse struct {
	Rule *Rule `json:"rule,omitempty"`
	apiResponse
}
