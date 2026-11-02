package collectionspaymentpromise

import (
	"github.com/chargebee/chargebee-go/v3/filter"
	collectionsPaymentPromiseEnum "github.com/chargebee/chargebee-go/v3/models/collectionspaymentpromise/enum"
)

type CollectionsPaymentPromise struct {
	Id              string                               `json:"id"`
	CustomerId      string                               `json:"customer_id"`
	InvoiceId       string                               `json:"invoice_id"`
	Status          collectionsPaymentPromiseEnum.Status `json:"status"`
	Reason          string                               `json:"reason"`
	PromisedBy      string                               `json:"promised_by"`
	PromisedOn      int64                                `json:"promised_on"`
	CreatedAt       int64                                `json:"created_at"`
	UpdatedAt       int64                                `json:"updated_at"`
	ResourceVersion int64                                `json:"resource_version"`
	Object          string                               `json:"object"`
}
type ListCollectionsPaymentPromisesRequestParams struct {
	Limit      *int32               `json:"limit,omitempty"`
	Offset     string               `json:"offset,omitempty"`
	CustomerId *filter.StringFilter `json:"customer_id,omitempty"`
	InvoiceId  *filter.StringFilter `json:"invoice_id,omitempty"`
	Status     *filter.EnumFilter   `json:"status,omitempty"`
}
type CreateCollectionsPaymentPromiseRequestParams struct {
	CustomerId   string `json:"customer_id"`
	InvoiceId    string `json:"invoice_id,omitempty"`
	Amount       *int64 `json:"amount"`
	CurrencyCode string `json:"currency_code,omitempty"`
	PromisedOn   *int64 `json:"promised_on"`
	Reason       string `json:"reason,omitempty"`
	PromisedBy   string `json:"promised_by,omitempty"`
}
type UpdateCollectionsPaymentPromiseRequestParams struct {
	PromisedOn *int64 `json:"promised_on,omitempty"`
	Reason     string `json:"reason,omitempty"`
	PromisedBy string `json:"promised_by,omitempty"`
}
