package chargebee

type SubscriptionEstimateStatus string

const (
	SubscriptionEstimateStatusFuture      SubscriptionEstimateStatus = "future"
	SubscriptionEstimateStatusInTrial     SubscriptionEstimateStatus = "in_trial"
	SubscriptionEstimateStatusActive      SubscriptionEstimateStatus = "active"
	SubscriptionEstimateStatusNonRenewing SubscriptionEstimateStatus = "non_renewing"
	SubscriptionEstimateStatusPaused      SubscriptionEstimateStatus = "paused"
	SubscriptionEstimateStatusCancelled   SubscriptionEstimateStatus = "cancelled"
	SubscriptionEstimateStatusTransferred SubscriptionEstimateStatus = "transferred"
)

type SubscriptionEstimateContractTermStatus string

const (
	SubscriptionEstimateContractTermStatusActive     SubscriptionEstimateContractTermStatus = "active"
	SubscriptionEstimateContractTermStatusCompleted  SubscriptionEstimateContractTermStatus = "completed"
	SubscriptionEstimateContractTermStatusCancelled  SubscriptionEstimateContractTermStatus = "cancelled"
	SubscriptionEstimateContractTermStatusTerminated SubscriptionEstimateContractTermStatus = "terminated"
)

type SubscriptionEstimateContractTermActionAtTermEnd string

const (
	SubscriptionEstimateContractTermActionAtTermEndRenew     SubscriptionEstimateContractTermActionAtTermEnd = "renew"
	SubscriptionEstimateContractTermActionAtTermEndEvergreen SubscriptionEstimateContractTermActionAtTermEnd = "evergreen"
	SubscriptionEstimateContractTermActionAtTermEndCancel    SubscriptionEstimateContractTermActionAtTermEnd = "cancel"
	SubscriptionEstimateContractTermActionAtTermEndRenewOnce SubscriptionEstimateContractTermActionAtTermEnd = "renew_once"
)

// just struct
type SubscriptionEstimate struct {
	Id              string                               `json:"id"`
	CurrencyCode    string                               `json:"currency_code"`
	Status          SubscriptionEstimateStatus           `json:"status"`
	TrialEndAction  TrialEndAction                       `json:"trial_end_action"`
	NextBillingAt   int64                                `json:"next_billing_at"`
	PauseDate       int64                                `json:"pause_date"`
	ResumeDate      int64                                `json:"resume_date"`
	ShippingAddress *SubscriptionEstimateShippingAddress `json:"shipping_address"`
	ContractTerm    *SubscriptionEstimateContractTerm    `json:"contract_term"`
	Object          string                               `json:"object"`
}

// sub resources
type SubscriptionEstimateShippingAddress struct {
	FirstName        string           `json:"first_name"`
	LastName         string           `json:"last_name"`
	Email            string           `json:"email"`
	Company          string           `json:"company"`
	Phone            string           `json:"phone"`
	Line1            string           `json:"line1"`
	Line2            string           `json:"line2"`
	Line3            string           `json:"line3"`
	City             string           `json:"city"`
	StateCode        string           `json:"state_code"`
	State            string           `json:"state"`
	Country          string           `json:"country"`
	Zip              string           `json:"zip"`
	ValidationStatus ValidationStatus `json:"validation_status"`
	Object           string           `json:"object"`
}

type SubscriptionEstimateContractTerm struct {
	Id                          string                                          `json:"id"`
	Status                      SubscriptionEstimateContractTermStatus          `json:"status"`
	ContractStart               int64                                           `json:"contract_start"`
	ContractEnd                 int64                                           `json:"contract_end"`
	BillingCycle                int32                                           `json:"billing_cycle"`
	ActionAtTermEnd             SubscriptionEstimateContractTermActionAtTermEnd `json:"action_at_term_end"`
	TotalContractValue          int64                                           `json:"total_contract_value"`
	TotalContractValueBeforeTax int64                                           `json:"total_contract_value_before_tax"`
	CancellationCutoffPeriod    int32                                           `json:"cancellation_cutoff_period"`
	CreatedAt                   int64                                           `json:"created_at"`
	SubscriptionId              string                                          `json:"subscription_id"`
	RemainingBillingCycles      int32                                           `json:"remaining_billing_cycles"`
	Object                      string                                          `json:"object"`
}

// operations
// input params
