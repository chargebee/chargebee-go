package transaction

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	creditNoteEnum "github.com/chargebee/chargebee-go/models/creditnote/enum"
	invoiceEnum "github.com/chargebee/chargebee-go/models/invoice/enum"
	transactionEnum "github.com/chargebee/chargebee-go/models/transaction/enum"
)

type Transaction struct {
	Id                     string                    `json:"id"`
	CustomerId             string                    `json:"customer_id"`
	SubscriptionId         string                    `json:"subscription_id"`
	GatewayAccountId       string                    `json:"gateway_account_id"`
	PaymentSourceId        string                    `json:"payment_source_id"`
	PaymentMethod          enum.PaymentMethod        `json:"payment_method"`
	ReferenceNumber        string                    `json:"reference_number"`
	Gateway                enum.Gateway              `json:"gateway"`
	Type                   transactionEnum.Type      `json:"type"`
	Date                   int64                     `json:"date"`
	SettledAt              int64                     `json:"settled_at"`
	ExchangeRate           float64                   `json:"exchange_rate"`
	CurrencyCode           string                    `json:"currency_code"`
	Amount                 int32                     `json:"amount"`
	IdAtGateway            string                    `json:"id_at_gateway"`
	Status                 transactionEnum.Status    `json:"status"`
	FraudFlag              transactionEnum.FraudFlag `json:"fraud_flag"`
	ErrorCode              string                    `json:"error_code"`
	ErrorText              string                    `json:"error_text"`
	VoidedAt               int64                     `json:"voided_at"`
	ResourceVersion        int64                     `json:"resource_version"`
	UpdatedAt              int64                     `json:"updated_at"`
	FraudReason            string                    `json:"fraud_reason"`
	AmountUnused           int32                     `json:"amount_unused"`
	MaskedCardNumber       string                    `json:"masked_card_number"`
	ReferenceTransactionId string                    `json:"reference_transaction_id"`
	RefundedTxnId          string                    `json:"refunded_txn_id"`
	ReversalTransactionId  string                    `json:"reversal_transaction_id"`
	LinkedInvoices         []*LinkedInvoice          `json:"linked_invoices"`
	LinkedCreditNotes      []*LinkedCreditNote       `json:"linked_credit_notes"`
	LinkedRefunds          []*LinkedRefund           `json:"linked_refunds"`
	Deleted                bool                      `json:"deleted"`
	BaseCurrencyCode       string                    `json:"base_currency_code"`
	Object                 string                    `json:"object"`
}
type LinkedInvoice struct {
	InvoiceId     string             `json:"invoice_id"`
	AppliedAmount int32              `json:"applied_amount"`
	AppliedAt     int64              `json:"applied_at"`
	InvoiceDate   int64              `json:"invoice_date"`
	InvoiceTotal  int32              `json:"invoice_total"`
	InvoiceStatus invoiceEnum.Status `json:"invoice_status"`
	Object        string             `json:"object"`
}
type LinkedCreditNote struct {
	CnId                 string                    `json:"cn_id"`
	AppliedAmount        int32                     `json:"applied_amount"`
	AppliedAt            int64                     `json:"applied_at"`
	CnReasonCode         creditNoteEnum.ReasonCode `json:"cn_reason_code"`
	CnDate               int64                     `json:"cn_date"`
	CnTotal              int32                     `json:"cn_total"`
	CnStatus             creditNoteEnum.Status     `json:"cn_status"`
	CnReferenceInvoiceId string                    `json:"cn_reference_invoice_id"`
	Object               string                    `json:"object"`
}
type LinkedRefund struct {
	TxnId     string                 `json:"txn_id"`
	TxnStatus transactionEnum.Status `json:"txn_status"`
	TxnDate   int64                  `json:"txn_date"`
	TxnAmount int32                  `json:"txn_amount"`
	Object    string                 `json:"object"`
}
type ListRequestParams struct {
	Limit            *int32                  `json:"limit,omitempty"`
	Offset           string                  `json:"offset,omitempty"`
	IncludeDeleted   *bool                   `json:"include_deleted,omitempty"`
	Id               *filter.StringFilter    `json:"id,omitempty"`
	CustomerId       *filter.StringFilter    `json:"customer_id,omitempty"`
	SubscriptionId   *filter.StringFilter    `json:"subscription_id,omitempty"`
	PaymentSourceId  *filter.StringFilter    `json:"payment_source_id,omitempty"`
	PaymentMethod    *filter.EnumFilter      `json:"payment_method,omitempty"`
	Gateway          *filter.EnumFilter      `json:"gateway,omitempty"`
	GatewayAccountId *filter.StringFilter    `json:"gateway_account_id,omitempty"`
	IdAtGateway      *filter.StringFilter    `json:"id_at_gateway,omitempty"`
	ReferenceNumber  *filter.StringFilter    `json:"reference_number,omitempty"`
	Type             *filter.EnumFilter      `json:"type,omitempty"`
	Date             *filter.TimestampFilter `json:"date,omitempty"`
	Amount           *filter.NumberFilter    `json:"amount,omitempty"`
	Status           *filter.EnumFilter      `json:"status,omitempty"`
	UpdatedAt        *filter.TimestampFilter `json:"updated_at,omitempty"`
	SortBy           *filter.SortFilter      `json:"sort_by,omitempty"`
}

type TransactionsForCustomerRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}

type TransactionsForSubscriptionRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}

type PaymentsForInvoiceRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
