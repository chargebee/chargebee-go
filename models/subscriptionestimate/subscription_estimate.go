package subscriptionestimate

import (
	"github.com/chargebee/chargebee-go/enum"
	subscriptionEstimateEnum "github.com/chargebee/chargebee-go/models/subscriptionestimate/enum"
)

type SubscriptionEstimate struct {
	Id              string                          `json:"id"`
	CurrencyCode    string                          `json:"currency_code"`
	Status          subscriptionEstimateEnum.Status `json:"status"`
	TrialEndAction  enum.TrialEndAction             `json:"trial_end_action"`
	NextBillingAt   int64                           `json:"next_billing_at"`
	PauseDate       int64                           `json:"pause_date"`
	ResumeDate      int64                           `json:"resume_date"`
	ShippingAddress *ShippingAddress                `json:"shipping_address"`
	ContractTerm    *ContractTerm                   `json:"contract_term"`
	Object          string                          `json:"object"`
}
type ShippingAddress struct {
	FirstName        string                `json:"first_name"`
	LastName         string                `json:"last_name"`
	Email            string                `json:"email"`
	Company          string                `json:"company"`
	Phone            string                `json:"phone"`
	Line1            string                `json:"line1"`
	Line2            string                `json:"line2"`
	Line3            string                `json:"line3"`
	City             string                `json:"city"`
	StateCode        string                `json:"state_code"`
	State            string                `json:"state"`
	Country          string                `json:"country"`
	Zip              string                `json:"zip"`
	ValidationStatus enum.ValidationStatus `json:"validation_status"`
	Index            int32                 `json:"index"`
	Object           string                `json:"object"`
}
type ContractTerm struct {
	Id                          string                                               `json:"id"`
	Status                      subscriptionEstimateEnum.ContractTermStatus          `json:"status"`
	ContractStart               int64                                                `json:"contract_start"`
	ContractEnd                 int64                                                `json:"contract_end"`
	BillingCycle                int32                                                `json:"billing_cycle"`
	ActionAtTermEnd             subscriptionEstimateEnum.ContractTermActionAtTermEnd `json:"action_at_term_end"`
	TotalContractValue          int64                                                `json:"total_contract_value"`
	TotalContractValueBeforeTax int64                                                `json:"total_contract_value_before_tax"`
	CancellationCutoffPeriod    int32                                                `json:"cancellation_cutoff_period"`
	CreatedAt                   int64                                                `json:"created_at"`
	SubscriptionId              string                                               `json:"subscription_id"`
	RemainingBillingCycles      int32                                                `json:"remaining_billing_cycles"`
	Object                      string                                               `json:"object"`
}
