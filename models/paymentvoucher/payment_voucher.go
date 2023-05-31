package paymentvoucher

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	paymentVoucherEnum "github.com/chargebee/chargebee-go/models/paymentvoucher/enum"
)

type PaymentVoucher struct {
	Id                 string                    `json:"id"`
	IdAtGateway        string                    `json:"id_at_gateway"`
	PaymentVoucherType enum.PaymentVoucherType   `json:"payment_voucher_type"`
	ExpiresAt          int64                     `json:"expires_at"`
	Status             paymentVoucherEnum.Status `json:"status"`
	SubscriptionId     string                    `json:"subscription_id"`
	CurrencyCode       string                    `json:"currency_code"`
	Amount             int32                     `json:"amount"`
	GatewayAccountId   string                    `json:"gateway_account_id"`
	PaymentSourceId    string                    `json:"payment_source_id"`
	Gateway            enum.Gateway              `json:"gateway"`
	Payload            string                    `json:"payload"`
	ErrorCode          string                    `json:"error_code"`
	ErrorText          string                    `json:"error_text"`
	Url                string                    `json:"url"`
	Date               int64                     `json:"date"`
	ResourceVersion    int64                     `json:"resource_version"`
	UpdatedAt          int64                     `json:"updated_at"`
	CustomerId         string                    `json:"customer_id"`
	LinkedInvoices     []*LinkedInvoice          `json:"linked_invoices"`
	Object             string                    `json:"object"`
}
type LinkedInvoice struct {
	InvoiceId string `json:"invoice_id"`
	TxnId     string `json:"txn_id"`
	AppliedAt int64  `json:"applied_at"`
	Object    string `json:"object"`
}
type CreateRequestParams struct {
	CustomerId           string                            `json:"customer_id"`
	VoucherPaymentSource *CreateVoucherPaymentSourceParams `json:"voucher_payment_source,omitempty"`
	InvoiceAllocations   []*CreateInvoiceAllocationParams  `json:"invoice_allocations,omitempty"`
	PaymentSourceId      string                            `json:"payment_source_id,omitempty"`
}
type CreateVoucherPaymentSourceParams struct {
	VoucherType enum.VoucherType `json:"voucher_type"`
}
type CreateInvoiceAllocationParams struct {
	InvoiceId string `json:"invoice_id"`
}
type PaymentVouchersForInvoiceRequestParams struct {
	Limit  *int32             `json:"limit,omitempty"`
	Offset string             `json:"offset,omitempty"`
	Status *filter.EnumFilter `json:"status,omitempty"`
	SortBy *filter.SortFilter `json:"sort_by,omitempty"`
}
type PaymentVouchersForCustomerRequestParams struct {
	Limit  *int32             `json:"limit,omitempty"`
	Offset string             `json:"offset,omitempty"`
	Status *filter.EnumFilter `json:"status,omitempty"`
	SortBy *filter.SortFilter `json:"sort_by,omitempty"`
}
