package contractterm

import (
	contractTermEnum "github.com/chargebee/chargebee-go/models/contractterm/enum"
)

type ContractTerm struct {
	Id                          string                           `json:"id"`
	Status                      contractTermEnum.Status          `json:"status"`
	ContractStart               int64                            `json:"contract_start"`
	ContractEnd                 int64                            `json:"contract_end"`
	BillingCycle                int32                            `json:"billing_cycle"`
	ActionAtTermEnd             contractTermEnum.ActionAtTermEnd `json:"action_at_term_end"`
	TotalContractValue          int64                            `json:"total_contract_value"`
	TotalContractValueBeforeTax int64                            `json:"total_contract_value_before_tax"`
	CancellationCutoffPeriod    int32                            `json:"cancellation_cutoff_period"`
	CreatedAt                   int64                            `json:"created_at"`
	SubscriptionId              string                           `json:"subscription_id"`
	RemainingBillingCycles      int32                            `json:"remaining_billing_cycles"`
	Object                      string                           `json:"object"`
}
