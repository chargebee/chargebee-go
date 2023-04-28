package inappsubscription

import (
	inAppSubscriptionEnum "github.com/chargebee/chargebee-go/models/inappsubscription/enum"
)

type InAppSubscription struct {
	AppId          string                            `json:"app_id"`
	SubscriptionId string                            `json:"subscription_id"`
	CustomerId     string                            `json:"customer_id"`
	PlanId         string                            `json:"plan_id"`
	StoreStatus    inAppSubscriptionEnum.StoreStatus `json:"store_status"`
	InvoiceId      string                            `json:"invoice_id"`
	Object         string                            `json:"object"`
}
type ProcessReceiptRequestParams struct {
	Receipt  string                        `json:"receipt"`
	Product  *ProcessReceiptProductParams  `json:"product,omitempty"`
	Customer *ProcessReceiptCustomerParams `json:"customer,omitempty"`
}
type ProcessReceiptProductParams struct {
	Id             string `json:"id"`
	CurrencyCode   string `json:"currency_code"`
	Price          *int32 `json:"price"`
	Name           string `json:"name,omitempty"`
	PriceInDecimal string `json:"price_in_decimal,omitempty"`
	Period         string `json:"period,omitempty"`
	PeriodUnit     string `json:"period_unit,omitempty"`
}
type ProcessReceiptCustomerParams struct {
	Id        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}
type ImportReceiptRequestParams struct {
	Receipt  string                       `json:"receipt"`
	Product  *ImportReceiptProductParams  `json:"product,omitempty"`
	Customer *ImportReceiptCustomerParams `json:"customer,omitempty"`
}
type ImportReceiptProductParams struct {
	CurrencyCode string `json:"currency_code"`
}
type ImportReceiptCustomerParams struct {
	Id    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}
type ImportSubscriptionRequestParams struct {
	Subscription *ImportSubscriptionSubscriptionParams `json:"subscription,omitempty"`
	Customer     *ImportSubscriptionCustomerParams     `json:"customer,omitempty"`
}
type ImportSubscriptionSubscriptionParams struct {
	Id            string `json:"id"`
	StartedAt     *int64 `json:"started_at"`
	TermStart     *int64 `json:"term_start"`
	TermEnd       *int64 `json:"term_end"`
	ProductId     string `json:"product_id"`
	CurrencyCode  string `json:"currency_code"`
	TransactionId string `json:"transaction_id"`
	IsTrial       *bool  `json:"is_trial,omitempty"`
}
type ImportSubscriptionCustomerParams struct {
	Id    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}
type RetrieveStoreSubsRequestParams struct {
	Receipt string `json:"receipt"`
}
