package chargebee

type ContractTermStatus string

const (
	ContractTermStatusActive     ContractTermStatus = "active"
	ContractTermStatusCompleted  ContractTermStatus = "completed"
	ContractTermStatusCancelled  ContractTermStatus = "cancelled"
	ContractTermStatusTerminated ContractTermStatus = "terminated"
)

type ContractTermActionAtTermEnd string

const (
	ContractTermActionAtTermEndRenew     ContractTermActionAtTermEnd = "renew"
	ContractTermActionAtTermEndEvergreen ContractTermActionAtTermEnd = "evergreen"
	ContractTermActionAtTermEndCancel    ContractTermActionAtTermEnd = "cancel"
	ContractTermActionAtTermEndRenewOnce ContractTermActionAtTermEnd = "renew_once"
)

// just struct
type ContractTerm struct {
	Id                          string                      `json:"id"`
	Status                      ContractTermStatus          `json:"status"`
	ContractStart               int64                       `json:"contract_start"`
	ContractEnd                 int64                       `json:"contract_end"`
	BillingCycle                int32                       `json:"billing_cycle"`
	ActionAtTermEnd             ContractTermActionAtTermEnd `json:"action_at_term_end"`
	TotalContractValue          int64                       `json:"total_contract_value"`
	TotalContractValueBeforeTax int64                       `json:"total_contract_value_before_tax"`
	CancellationCutoffPeriod    int32                       `json:"cancellation_cutoff_period"`
	CreatedAt                   int64                       `json:"created_at"`
	SubscriptionId              string                      `json:"subscription_id"`
	RemainingBillingCycles      int32                       `json:"remaining_billing_cycles"`
	Object                      string                      `json:"object"`
}

// sub resources
// operations
// input params
