package chargebee

type Status string

const (
	StatusActive   Status = "active"
	StatusConsumed Status = "consumed"
	StatusExpired  Status = "expired"
	StatusFailure  Status = "failure"
)

type PaymentVoucher struct {
	Id                 string                  `json:"id"`
	IdAtGateway        string                  `json:"id_at_gateway"`
	PaymentVoucherType enum.PaymentVoucherType `json:"payment_voucher_type"`
	ExpiresAt          int64                   `json:"expires_at"`
	Status             Status                  `json:"status"`
	SubscriptionId     string                  `json:"subscription_id"`
	CurrencyCode       string                  `json:"currency_code"`
	Amount             int64                   `json:"amount"`
	GatewayAccountId   string                  `json:"gateway_account_id"`
	PaymentSourceId    string                  `json:"payment_source_id"`
	Gateway            enum.Gateway            `json:"gateway"`
	Payload            string                  `json:"payload"`
	ErrorCode          string                  `json:"error_code"`
	ErrorText          string                  `json:"error_text"`
	Url                string                  `json:"url"`
	Date               int64                   `json:"date"`
	ResourceVersion    int64                   `json:"resource_version"`
	UpdatedAt          int64                   `json:"updated_at"`
	CustomerId         string                  `json:"customer_id"`
	LinkedInvoices     []*LinkedInvoice        `json:"linked_invoices"`
	Object             string                  `json:"object"`
}
type LinkedInvoice struct {
	InvoiceId string `json:"invoice_id"`
	TxnId     string `json:"txn_id"`
	AppliedAt int64  `json:"applied_at"`
	Object    string `json:"object"`
}
type CreateRequest struct {
	CustomerId           string                      `json:"customer_id"`
	VoucherPaymentSource *CreateVoucherPaymentSource `json:"voucher_payment_source,omitempty"`
	InvoiceAllocations   []*CreateInvoiceAllocation  `json:"invoice_allocations,omitempty"`
	PaymentSourceId      string                      `json:"payment_source_id,omitempty"`
}
type CreateVoucherPaymentSource struct {
	VoucherType enum.VoucherType `json:"voucher_type"`
}
type CreateInvoiceAllocation struct {
	InvoiceId string `json:"invoice_id"`
}
type PaymentVouchersForInvoiceRequest struct {
	Limit  *int32             `json:"limit,omitempty"`
	Offset string             `json:"offset,omitempty"`
	Status *filter.EnumFilter `json:"status,omitempty"`
	SortBy *filter.SortFilter `json:"sort_by,omitempty"`
}
type PaymentVouchersForCustomerRequest struct {
	Limit  *int32             `json:"limit,omitempty"`
	Offset string             `json:"offset,omitempty"`
	Status *filter.EnumFilter `json:"status,omitempty"`
	SortBy *filter.SortFilter `json:"sort_by,omitempty"`
}

type CreateResponse struct {
	PaymentVoucher *PaymentVoucher `json:"payment_voucher,omitempty"`
}

type RetrieveResponse struct {
	PaymentVoucher *PaymentVoucher `json:"payment_voucher,omitempty"`
}

type PaymentVouchersForInvoicePaymentVoucherResponse struct {
	PaymentVoucher *PaymentVoucher `json:"payment_voucher,omitempty"`
}

type PaymentVouchersForInvoiceResponse struct {
	List       []*PaymentVouchersForInvoicePaymentVoucherResponse `json:"list,omitempty"`
	NextOffset string                                             `json:"next_offset,omitempty"`
}

type PaymentVouchersForCustomerPaymentVoucherResponse struct {
	PaymentVoucher *PaymentVoucher `json:"payment_voucher,omitempty"`
}

type PaymentVouchersForCustomerResponse struct {
	List       []*PaymentVouchersForCustomerPaymentVoucherResponse `json:"list,omitempty"`
	NextOffset string                                              `json:"next_offset,omitempty"`
}
