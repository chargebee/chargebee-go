package chargebee

type StoreStatus string

const (
	StoreStatusInTrial   StoreStatus = "in_trial"
	StoreStatusActive    StoreStatus = "active"
	StoreStatusCancelled StoreStatus = "cancelled"
	StoreStatusPaused    StoreStatus = "paused"
)

type InAppSubscription struct {
	AppId          string      `json:"app_id"`
	SubscriptionId string      `json:"subscription_id"`
	CustomerId     string      `json:"customer_id"`
	PlanId         string      `json:"plan_id"`
	StoreStatus    StoreStatus `json:"store_status"`
	InvoiceId      string      `json:"invoice_id"`
	Object         string      `json:"object"`
}
type ProcessReceiptRequest struct {
	Receipt  string                  `json:"receipt"`
	Product  *ProcessReceiptProduct  `json:"product,omitempty"`
	Customer *ProcessReceiptCustomer `json:"customer,omitempty"`
}
type ProcessReceiptProduct struct {
	Id             string `json:"id"`
	CurrencyCode   string `json:"currency_code"`
	Price          *int64 `json:"price"`
	Name           string `json:"name,omitempty"`
	PriceInDecimal string `json:"price_in_decimal,omitempty"`
	Period         string `json:"period,omitempty"`
	PeriodUnit     string `json:"period_unit,omitempty"`
}
type ProcessReceiptCustomer struct {
	Id        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}
type ImportReceiptRequest struct {
	Receipt  string                 `json:"receipt"`
	Product  *ImportReceiptProduct  `json:"product,omitempty"`
	Customer *ImportReceiptCustomer `json:"customer,omitempty"`
}
type ImportReceiptProduct struct {
	CurrencyCode string `json:"currency_code"`
}
type ImportReceiptCustomer struct {
	Id    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}
type ImportSubscriptionRequest struct {
	Subscription *ImportSubscriptionSubscription `json:"subscription,omitempty"`
	Customer     *ImportSubscriptionCustomer     `json:"customer,omitempty"`
}
type ImportSubscriptionSubscription struct {
	Id            string `json:"id"`
	StartedAt     *int64 `json:"started_at"`
	TermStart     *int64 `json:"term_start"`
	TermEnd       *int64 `json:"term_end"`
	ProductId     string `json:"product_id"`
	CurrencyCode  string `json:"currency_code"`
	TransactionId string `json:"transaction_id"`
	IsTrial       *bool  `json:"is_trial,omitempty"`
}
type ImportSubscriptionCustomer struct {
	Id    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}
type RetrieveStoreSubsRequest struct {
	Receipt string `json:"receipt"`
}

type ProcessReceiptResponse struct {
	InAppSubscription *InAppSubscription `json:"in_app_subscription,omitempty"`
}

type ImportReceiptResponse struct {
	InAppSubscriptions []*InAppSubscription `json:"in_app_subscriptions,omitempty"`
}

type ImportSubscriptionResponse struct {
	InAppSubscription *InAppSubscription `json:"in_app_subscription,omitempty"`
}

type RetrieveStoreSubsResponse struct {
	InAppSubscriptions []*InAppSubscription `json:"in_app_subscriptions,omitempty"`
}
