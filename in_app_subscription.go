package chargebee

type InAppSubscriptionStoreStatus string

const (
	InAppSubscriptionStoreStatusInTrial   InAppSubscriptionStoreStatus = "in_trial"
	InAppSubscriptionStoreStatusActive    InAppSubscriptionStoreStatus = "active"
	InAppSubscriptionStoreStatusCancelled InAppSubscriptionStoreStatus = "cancelled"
	InAppSubscriptionStoreStatusPaused    InAppSubscriptionStoreStatus = "paused"
)

// just struct
type InAppSubscription struct {
	SubscriptionId string                       `json:"subscription_id"`
	CustomerId     string                       `json:"customer_id"`
	PlanId         string                       `json:"plan_id"`
	StoreStatus    InAppSubscriptionStoreStatus `json:"store_status"`
	InvoiceId      string                       `json:"invoice_id"`
	Object         string                       `json:"object"`
}

// sub resources
// operations
// input params
type InAppSubscriptionProcessReceiptRequest struct {
	Receipt    string                                   `json:"receipt"`
	Product    *InAppSubscriptionProcessReceiptProduct  `json:"product,omitempty"`
	Customer   *InAppSubscriptionProcessReceiptCustomer `json:"customer,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *InAppSubscriptionProcessReceiptRequest) payload() any { return r }

// input sub resource params single
type InAppSubscriptionProcessReceiptProduct struct {
	Id             string `json:"id"`
	CurrencyCode   string `json:"currency_code"`
	Price          *int64 `json:"price"`
	Name           string `json:"name,omitempty"`
	PriceInDecimal string `json:"price_in_decimal,omitempty"`
	Period         string `json:"period,omitempty"`
	PeriodUnit     string `json:"period_unit,omitempty"`
}

// input sub resource params single
type InAppSubscriptionProcessReceiptCustomer struct {
	Id        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

type InAppSubscriptionImportReceiptRequest struct {
	Receipt    string                                  `json:"receipt"`
	Product    *InAppSubscriptionImportReceiptProduct  `json:"product,omitempty"`
	Customer   *InAppSubscriptionImportReceiptCustomer `json:"customer,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *InAppSubscriptionImportReceiptRequest) payload() any { return r }

// input sub resource params single
type InAppSubscriptionImportReceiptProduct struct {
	CurrencyCode string `json:"currency_code"`
}

// input sub resource params single
type InAppSubscriptionImportReceiptCustomer struct {
	Id    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

type InAppSubscriptionImportSubscriptionRequest struct {
	Subscription *InAppSubscriptionImportSubscriptionSubscription `json:"subscription,omitempty"`
	Customer     *InAppSubscriptionImportSubscriptionCustomer     `json:"customer,omitempty"`
	apiRequest   `json:"-" form:"-"`
}

func (r *InAppSubscriptionImportSubscriptionRequest) payload() any { return r }

// input sub resource params single
type InAppSubscriptionImportSubscriptionSubscription struct {
	Id            string `json:"id"`
	StartedAt     *int64 `json:"started_at"`
	TermStart     *int64 `json:"term_start"`
	TermEnd       *int64 `json:"term_end"`
	ProductId     string `json:"product_id"`
	CurrencyCode  string `json:"currency_code"`
	TransactionId string `json:"transaction_id"`
	IsTrial       *bool  `json:"is_trial,omitempty"`
}

// input sub resource params single
type InAppSubscriptionImportSubscriptionCustomer struct {
	Id    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

type InAppSubscriptionRetrieveStoreSubsRequest struct {
	Receipt    string `json:"receipt"`
	apiRequest `json:"-" form:"-"`
}

func (r *InAppSubscriptionRetrieveStoreSubsRequest) payload() any { return r }

// operation response
type InAppSubscriptionProcessReceiptResponse struct {
	InAppSubscription *InAppSubscription `json:"in_app_subscription,omitempty"`
	apiResponse
}

// operation response
type InAppSubscriptionImportReceiptResponse struct {
	InAppSubscriptions []*InAppSubscription `json:"in_app_subscriptions,omitempty"`
	apiResponse
}

// operation response
type InAppSubscriptionImportSubscriptionResponse struct {
	InAppSubscription *InAppSubscription `json:"in_app_subscription,omitempty"`
	apiResponse
}

// operation response
type InAppSubscriptionRetrieveStoreSubsResponse struct {
	InAppSubscriptions []*InAppSubscription `json:"in_app_subscriptions,omitempty"`
	apiResponse
}
