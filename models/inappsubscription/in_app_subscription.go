package inappsubscription

type InAppSubscription struct {
	AppId          string `json:"app_id"`
	SubscriptionId string `json:"subscription_id"`
	CustomerId     string `json:"customer_id"`
	PlanId         string `json:"plan_id"`
	Object         string `json:"object"`
}
type ProcessReceiptRequestParams struct {
	Receipt  string                        `json:"receipt"`
	Product  *ProcessReceiptProductParams  `json:"product,omitempty"`
	Customer *ProcessReceiptCustomerParams `json:"customer,omitempty"`
}
type ProcessReceiptProductParams struct {
	Id             string `json:"id"`
	Name           string `json:"name,omitempty"`
	CurrencyCode   string `json:"currency_code"`
	Price          *int64 `json:"price"`
	PriceInDecimal string `json:"price_in_decimal,omitempty"`
	Period         string `json:"period,omitempty"`
	PeriodUnit     string `json:"period_unit,omitempty"`
}
type ProcessReceiptCustomerParams struct {
	Id string `json:"id,omitempty"`
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
