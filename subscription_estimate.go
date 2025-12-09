package chargebee

type Status string

const (
	StatusFuture      Status = "future"
	StatusInTrial     Status = "in_trial"
	StatusActive      Status = "active"
	StatusNonRenewing Status = "non_renewing"
	StatusPaused      Status = "paused"
	StatusCancelled   Status = "cancelled"
	StatusTransferred Status = "transferred"
)

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

type SubscriptionEstimate struct {
	Id              string              `json:"id"`
	CurrencyCode    string              `json:"currency_code"`
	Status          Status              `json:"status"`
	TrialEndAction  enum.TrialEndAction `json:"trial_end_action"`
	NextBillingAt   int64               `json:"next_billing_at"`
	PauseDate       int64               `json:"pause_date"`
	ResumeDate      int64               `json:"resume_date"`
	ShippingAddress *ShippingAddress    `json:"shipping_address"`
	ContractTerm    *ContractTerm       `json:"contract_term"`
	Object          string              `json:"object"`
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
	Object           string                `json:"object"`
}
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
