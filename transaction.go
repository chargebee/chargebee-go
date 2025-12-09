package chargebee

type Type string

const (
	TypeAuthorization   Type = "authorization"
	TypePayment         Type = "payment"
	TypeRefund          Type = "refund"
	TypePaymentReversal Type = "payment_reversal"
)

type Status string

const (
	StatusInProgress     Status = "in_progress"
	StatusSuccess        Status = "success"
	StatusVoided         Status = "voided"
	StatusFailure        Status = "failure"
	StatusTimeout        Status = "timeout"
	StatusNeedsAttention Status = "needs_attention"
	StatusLateFailure    Status = "late_failure"
)

type FraudFlag string

const (
	FraudFlagSafe       FraudFlag = "safe"
	FraudFlagSuspicious FraudFlag = "suspicious"
	FraudFlagFraudulent FraudFlag = "fraudulent"
)

type InitiatorType string

const (
	InitiatorTypeCustomer InitiatorType = "customer"
	InitiatorTypeMerchant InitiatorType = "merchant"
)

type AuthorizationReason string

const (
	AuthorizationReasonBlockingFunds    AuthorizationReason = "blocking_funds"
	AuthorizationReasonVerification     AuthorizationReason = "verification"
	AuthorizationReasonScheduledCapture AuthorizationReason = "scheduled_capture"
)

type LinkedPaymentStatus string

const (
	LinkedPaymentStatusInProgress     LinkedPaymentStatus = "in_progress"
	LinkedPaymentStatusSuccess        LinkedPaymentStatus = "success"
	LinkedPaymentStatusVoided         LinkedPaymentStatus = "voided"
	LinkedPaymentStatusFailure        LinkedPaymentStatus = "failure"
	LinkedPaymentStatusTimeout        LinkedPaymentStatus = "timeout"
	LinkedPaymentStatusNeedsAttention LinkedPaymentStatus = "needs_attention"
	LinkedPaymentStatusLateFailure    LinkedPaymentStatus = "late_failure"
)

type Transaction struct {
	Id                       string              `json:"id"`
	CustomerId               string              `json:"customer_id"`
	SubscriptionId           string              `json:"subscription_id"`
	GatewayAccountId         string              `json:"gateway_account_id"`
	PaymentSourceId          string              `json:"payment_source_id"`
	PaymentMethod            enum.PaymentMethod  `json:"payment_method"`
	ReferenceNumber          string              `json:"reference_number"`
	Gateway                  enum.Gateway        `json:"gateway"`
	Type                     Type                `json:"type"`
	Date                     int64               `json:"date"`
	SettledAt                int64               `json:"settled_at"`
	ExchangeRate             float64             `json:"exchange_rate"`
	CurrencyCode             string              `json:"currency_code"`
	Amount                   int64               `json:"amount"`
	IdAtGateway              string              `json:"id_at_gateway"`
	Status                   Status              `json:"status"`
	FraudFlag                FraudFlag           `json:"fraud_flag"`
	InitiatorType            InitiatorType       `json:"initiator_type"`
	ThreeDSecure             bool                `json:"three_d_secure"`
	AuthorizationReason      AuthorizationReason `json:"authorization_reason"`
	ErrorCode                string              `json:"error_code"`
	ErrorText                string              `json:"error_text"`
	VoidedAt                 int64               `json:"voided_at"`
	ResourceVersion          int64               `json:"resource_version"`
	UpdatedAt                int64               `json:"updated_at"`
	FraudReason              string              `json:"fraud_reason"`
	CustomPaymentMethodId    string              `json:"custom_payment_method_id"`
	AmountUnused             int64               `json:"amount_unused"`
	MaskedCardNumber         string              `json:"masked_card_number"`
	ReferenceTransactionId   string              `json:"reference_transaction_id"`
	RefundedTxnId            string              `json:"refunded_txn_id"`
	ReferenceAuthorizationId string              `json:"reference_authorization_id"`
	AmountCapturable         int64               `json:"amount_capturable"`
	ReversalTransactionId    string              `json:"reversal_transaction_id"`
	LinkedInvoices           []*LinkedInvoice    `json:"linked_invoices"`
	LinkedCreditNotes        []*LinkedCreditNote `json:"linked_credit_notes"`
	LinkedRefunds            []*LinkedRefund     `json:"linked_refunds"`
	LinkedPayments           []*LinkedPayment    `json:"linked_payments"`
	Deleted                  bool                `json:"deleted"`
	Iin                      string              `json:"iin"`
	Last4                    string              `json:"last4"`
	MerchantReferenceId      string              `json:"merchant_reference_id"`
	BusinessEntityId         string              `json:"business_entity_id"`
	PaymentMethodDetails     string              `json:"payment_method_details"`
	ErrorDetail              *GatewayErrorDetail `json:"error_detail"`
	CustomPaymentMethodName  string              `json:"custom_payment_method_name"`
	Object                   string              `json:"object"`
}
type LinkedInvoice struct {
	InvoiceId     string         `json:"invoice_id"`
	AppliedAmount int64          `json:"applied_amount"`
	AppliedAt     int64          `json:"applied_at"`
	InvoiceDate   int64          `json:"invoice_date"`
	InvoiceTotal  int64          `json:"invoice_total"`
	InvoiceStatus invoice.Status `json:"invoice_status"`
	Object        string         `json:"object"`
}
type LinkedCreditNote struct {
	CnId                 string                `json:"cn_id"`
	AppliedAmount        int64                 `json:"applied_amount"`
	AppliedAt            int64                 `json:"applied_at"`
	CnReasonCode         creditNote.ReasonCode `json:"cn_reason_code"`
	CnCreateReasonCode   string                `json:"cn_create_reason_code"`
	CnDate               int64                 `json:"cn_date"`
	CnTotal              int64                 `json:"cn_total"`
	CnStatus             creditNote.Status     `json:"cn_status"`
	CnReferenceInvoiceId string                `json:"cn_reference_invoice_id"`
	Object               string                `json:"object"`
}
type LinkedRefund struct {
	TxnId     string `json:"txn_id"`
	TxnStatus Status `json:"txn_status"`
	TxnDate   int64  `json:"txn_date"`
	TxnAmount int64  `json:"txn_amount"`
	Object    string `json:"object"`
}
type LinkedPayment struct {
	Id     string              `json:"id"`
	Status LinkedPaymentStatus `json:"status"`
	Amount int64               `json:"amount"`
	Date   int64               `json:"date"`
	Object string              `json:"object"`
}
type GatewayErrorDetail struct {
	RequestId             string `json:"request_id"`
	ErrorCategory         string `json:"error_category"`
	ErrorCode             string `json:"error_code"`
	ErrorMessage          string `json:"error_message"`
	DeclineCode           string `json:"decline_code"`
	DeclineMessage        string `json:"decline_message"`
	NetworkErrorCode      string `json:"network_error_code"`
	NetworkErrorMessage   string `json:"network_error_message"`
	ErrorField            string `json:"error_field"`
	RecommendationCode    string `json:"recommendation_code"`
	RecommendationMessage string `json:"recommendation_message"`
	ProcessorErrorCode    string `json:"processor_error_code"`
	ProcessorErrorMessage string `json:"processor_error_message"`
	ErrorCauseId          string `json:"error_cause_id"`
	ProcessorAdviceCode   string `json:"processor_advice_code"`
	Object                string `json:"object"`
}
type CreateAuthorizationRequest struct {
	CustomerId      string `json:"customer_id"`
	PaymentSourceId string `json:"payment_source_id,omitempty"`
	CurrencyCode    string `json:"currency_code,omitempty"`
	Amount          *int64 `json:"amount"`
}
type RecordRefundRequest struct {
	Amount                *int64             `json:"amount,omitempty"`
	PaymentMethod         enum.PaymentMethod `json:"payment_method"`
	Date                  *int64             `json:"date"`
	ReferenceNumber       string             `json:"reference_number,omitempty"`
	CustomPaymentMethodId string             `json:"custom_payment_method_id,omitempty"`
	Comment               string             `json:"comment,omitempty"`
}
type ReconcileRequest struct {
	IdAtGateway string `json:"id_at_gateway,omitempty"`
	CustomerId  string `json:"customer_id,omitempty"`
	Status      Status `json:"status,omitempty"`
}
type RefundRequest struct {
	Amount  *int64 `json:"amount,omitempty"`
	Comment string `json:"comment,omitempty"`
}
type ListRequest struct {
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
	AmountCapturable *filter.NumberFilter    `json:"amount_capturable,omitempty"`
	Status           *filter.EnumFilter      `json:"status,omitempty"`
	UpdatedAt        *filter.TimestampFilter `json:"updated_at,omitempty"`
	SortBy           *filter.SortFilter      `json:"sort_by,omitempty"`
}
type TransactionsForCustomerRequest struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
type TransactionsForSubscriptionRequest struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
type PaymentsForInvoiceRequest struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
type DeleteOfflineTransactionRequest struct {
	Comment string `json:"comment,omitempty"`
}

type CreateAuthorizationResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
}

type VoidTransactionResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
}

type RecordRefundResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
}

type ReconcileResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
}

type RefundResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
}

type ListTransactionResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
}

type ListResponse struct {
	List       []*ListTransactionResponse `json:"list,omitempty"`
	NextOffset string                     `json:"next_offset,omitempty"`
}

type TransactionsForCustomerTransactionResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
}

type TransactionsForCustomerResponse struct {
	List       []*TransactionsForCustomerTransactionResponse `json:"list,omitempty"`
	NextOffset string                                        `json:"next_offset,omitempty"`
}

type TransactionsForSubscriptionTransactionResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
}

type TransactionsForSubscriptionResponse struct {
	List       []*TransactionsForSubscriptionTransactionResponse `json:"list,omitempty"`
	NextOffset string                                            `json:"next_offset,omitempty"`
}

type PaymentsForInvoiceTransactionResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
}

type PaymentsForInvoiceResponse struct {
	List       []*PaymentsForInvoiceTransactionResponse `json:"list,omitempty"`
	NextOffset string                                   `json:"next_offset,omitempty"`
}

type RetrieveResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
}

type DeleteOfflineTransactionResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
}
