package chargebee

type TransactionType string

const (
	TransactionTypeAuthorization   TransactionType = "authorization"
	TransactionTypePayment         TransactionType = "payment"
	TransactionTypeRefund          TransactionType = "refund"
	TransactionTypePaymentReversal TransactionType = "payment_reversal"
)

type TransactionStatus string

const (
	TransactionStatusInProgress     TransactionStatus = "in_progress"
	TransactionStatusSuccess        TransactionStatus = "success"
	TransactionStatusVoided         TransactionStatus = "voided"
	TransactionStatusFailure        TransactionStatus = "failure"
	TransactionStatusTimeout        TransactionStatus = "timeout"
	TransactionStatusNeedsAttention TransactionStatus = "needs_attention"
	TransactionStatusLateFailure    TransactionStatus = "late_failure"
)

type TransactionFraudFlag string

const (
	TransactionFraudFlagSafe       TransactionFraudFlag = "safe"
	TransactionFraudFlagSuspicious TransactionFraudFlag = "suspicious"
	TransactionFraudFlagFraudulent TransactionFraudFlag = "fraudulent"
)

type TransactionInitiatorType string

const (
	TransactionInitiatorTypeCustomer TransactionInitiatorType = "customer"
	TransactionInitiatorTypeMerchant TransactionInitiatorType = "merchant"
)

type TransactionAuthorizationReason string

const (
	TransactionAuthorizationReasonBlockingFunds    TransactionAuthorizationReason = "blocking_funds"
	TransactionAuthorizationReasonVerification     TransactionAuthorizationReason = "verification"
	TransactionAuthorizationReasonScheduledCapture TransactionAuthorizationReason = "scheduled_capture"
)

type TransactionLinkedInvoiceInvoiceStatus string

const (
	TransactionLinkedInvoiceInvoiceStatusPaid       TransactionLinkedInvoiceInvoiceStatus = "paid"
	TransactionLinkedInvoiceInvoiceStatusPosted     TransactionLinkedInvoiceInvoiceStatus = "posted"
	TransactionLinkedInvoiceInvoiceStatusPaymentDue TransactionLinkedInvoiceInvoiceStatus = "payment_due"
	TransactionLinkedInvoiceInvoiceStatusNotPaid    TransactionLinkedInvoiceInvoiceStatus = "not_paid"
	TransactionLinkedInvoiceInvoiceStatusVoided     TransactionLinkedInvoiceInvoiceStatus = "voided"
	TransactionLinkedInvoiceInvoiceStatusPending    TransactionLinkedInvoiceInvoiceStatus = "pending"
)

type TransactionLinkedCreditNoteCnReasonCode string

const (
	TransactionLinkedCreditNoteCnReasonCodeWriteOff                 TransactionLinkedCreditNoteCnReasonCode = "write_off"
	TransactionLinkedCreditNoteCnReasonCodeSubscriptionChange       TransactionLinkedCreditNoteCnReasonCode = "subscription_change"
	TransactionLinkedCreditNoteCnReasonCodeSubscriptionCancellation TransactionLinkedCreditNoteCnReasonCode = "subscription_cancellation"
	TransactionLinkedCreditNoteCnReasonCodeSubscriptionPause        TransactionLinkedCreditNoteCnReasonCode = "subscription_pause"
	TransactionLinkedCreditNoteCnReasonCodeChargeback               TransactionLinkedCreditNoteCnReasonCode = "chargeback"
	TransactionLinkedCreditNoteCnReasonCodeProductUnsatisfactory    TransactionLinkedCreditNoteCnReasonCode = "product_unsatisfactory"
	TransactionLinkedCreditNoteCnReasonCodeServiceUnsatisfactory    TransactionLinkedCreditNoteCnReasonCode = "service_unsatisfactory"
	TransactionLinkedCreditNoteCnReasonCodeOrderChange              TransactionLinkedCreditNoteCnReasonCode = "order_change"
	TransactionLinkedCreditNoteCnReasonCodeOrderCancellation        TransactionLinkedCreditNoteCnReasonCode = "order_cancellation"
	TransactionLinkedCreditNoteCnReasonCodeWaiver                   TransactionLinkedCreditNoteCnReasonCode = "waiver"
	TransactionLinkedCreditNoteCnReasonCodeOther                    TransactionLinkedCreditNoteCnReasonCode = "other"
	TransactionLinkedCreditNoteCnReasonCodeFraudulent               TransactionLinkedCreditNoteCnReasonCode = "fraudulent"
)

type TransactionLinkedCreditNoteCnStatus string

const (
	TransactionLinkedCreditNoteCnStatusAdjusted  TransactionLinkedCreditNoteCnStatus = "adjusted"
	TransactionLinkedCreditNoteCnStatusRefunded  TransactionLinkedCreditNoteCnStatus = "refunded"
	TransactionLinkedCreditNoteCnStatusRefundDue TransactionLinkedCreditNoteCnStatus = "refund_due"
	TransactionLinkedCreditNoteCnStatusVoided    TransactionLinkedCreditNoteCnStatus = "voided"
)

type TransactionLinkedRefundTxnStatus string

const (
	TransactionLinkedRefundTxnStatusInProgress     TransactionLinkedRefundTxnStatus = "in_progress"
	TransactionLinkedRefundTxnStatusSuccess        TransactionLinkedRefundTxnStatus = "success"
	TransactionLinkedRefundTxnStatusVoided         TransactionLinkedRefundTxnStatus = "voided"
	TransactionLinkedRefundTxnStatusFailure        TransactionLinkedRefundTxnStatus = "failure"
	TransactionLinkedRefundTxnStatusTimeout        TransactionLinkedRefundTxnStatus = "timeout"
	TransactionLinkedRefundTxnStatusNeedsAttention TransactionLinkedRefundTxnStatus = "needs_attention"
	TransactionLinkedRefundTxnStatusLateFailure    TransactionLinkedRefundTxnStatus = "late_failure"
)

type TransactionLinkedPaymentStatus string

const (
	TransactionLinkedPaymentStatusInProgress     TransactionLinkedPaymentStatus = "in_progress"
	TransactionLinkedPaymentStatusSuccess        TransactionLinkedPaymentStatus = "success"
	TransactionLinkedPaymentStatusVoided         TransactionLinkedPaymentStatus = "voided"
	TransactionLinkedPaymentStatusFailure        TransactionLinkedPaymentStatus = "failure"
	TransactionLinkedPaymentStatusTimeout        TransactionLinkedPaymentStatus = "timeout"
	TransactionLinkedPaymentStatusNeedsAttention TransactionLinkedPaymentStatus = "needs_attention"
	TransactionLinkedPaymentStatusLateFailure    TransactionLinkedPaymentStatus = "late_failure"
)

// just struct
type Transaction struct {
	Id                       string                         `json:"id"`
	CustomerId               string                         `json:"customer_id"`
	SubscriptionId           string                         `json:"subscription_id"`
	GatewayAccountId         string                         `json:"gateway_account_id"`
	PaymentSourceId          string                         `json:"payment_source_id"`
	PaymentMethod            PaymentMethod                  `json:"payment_method"`
	ReferenceNumber          string                         `json:"reference_number"`
	Gateway                  Gateway                        `json:"gateway"`
	Type                     TransactionType                `json:"type"`
	Date                     int64                          `json:"date"`
	SettledAt                int64                          `json:"settled_at"`
	ExchangeRate             float64                        `json:"exchange_rate"`
	CurrencyCode             string                         `json:"currency_code"`
	Amount                   int64                          `json:"amount"`
	IdAtGateway              string                         `json:"id_at_gateway"`
	Status                   TransactionStatus              `json:"status"`
	FraudFlag                TransactionFraudFlag           `json:"fraud_flag"`
	InitiatorType            TransactionInitiatorType       `json:"initiator_type"`
	ThreeDSecure             bool                           `json:"three_d_secure"`
	AuthorizationReason      TransactionAuthorizationReason `json:"authorization_reason"`
	ErrorCode                string                         `json:"error_code"`
	ErrorText                string                         `json:"error_text"`
	VoidedAt                 int64                          `json:"voided_at"`
	ResourceVersion          int64                          `json:"resource_version"`
	UpdatedAt                int64                          `json:"updated_at"`
	FraudReason              string                         `json:"fraud_reason"`
	CustomPaymentMethodId    string                         `json:"custom_payment_method_id"`
	AmountUnused             int64                          `json:"amount_unused"`
	MaskedCardNumber         string                         `json:"masked_card_number"`
	ReferenceTransactionId   string                         `json:"reference_transaction_id"`
	RefundedTxnId            string                         `json:"refunded_txn_id"`
	ReferenceAuthorizationId string                         `json:"reference_authorization_id"`
	AmountCapturable         int64                          `json:"amount_capturable"`
	ReversalTransactionId    string                         `json:"reversal_transaction_id"`
	LinkedInvoices           []*TransactionLinkedInvoice    `json:"linked_invoices"`
	LinkedCreditNotes        []*TransactionLinkedCreditNote `json:"linked_credit_notes"`
	LinkedRefunds            []*TransactionLinkedRefund     `json:"linked_refunds"`
	LinkedPayments           []*TransactionLinkedPayment    `json:"linked_payments"`
	Deleted                  bool                           `json:"deleted"`
	Iin                      string                         `json:"iin"`
	Last4                    string                         `json:"last4"`
	MerchantReferenceId      string                         `json:"merchant_reference_id"`
	BusinessEntityId         string                         `json:"business_entity_id"`
	PaymentMethodDetails     string                         `json:"payment_method_details"`
	ErrorDetail              *TransactionGatewayErrorDetail `json:"error_detail"`
	CustomPaymentMethodName  string                         `json:"custom_payment_method_name"`
	Object                   string                         `json:"object"`
}

// sub resources
type TransactionLinkedInvoice struct {
	InvoiceId     string                                `json:"invoice_id"`
	AppliedAmount int64                                 `json:"applied_amount"`
	AppliedAt     int64                                 `json:"applied_at"`
	InvoiceDate   int64                                 `json:"invoice_date"`
	InvoiceTotal  int64                                 `json:"invoice_total"`
	InvoiceStatus TransactionLinkedInvoiceInvoiceStatus `json:"invoice_status"`
	Object        string                                `json:"object"`
}

type TransactionLinkedCreditNote struct {
	CnId                 string                                  `json:"cn_id"`
	AppliedAmount        int64                                   `json:"applied_amount"`
	AppliedAt            int64                                   `json:"applied_at"`
	CnReasonCode         TransactionLinkedCreditNoteCnReasonCode `json:"cn_reason_code"`
	CnCreateReasonCode   string                                  `json:"cn_create_reason_code"`
	CnDate               int64                                   `json:"cn_date"`
	CnTotal              int64                                   `json:"cn_total"`
	CnStatus             TransactionLinkedCreditNoteCnStatus     `json:"cn_status"`
	CnReferenceInvoiceId string                                  `json:"cn_reference_invoice_id"`
	Object               string                                  `json:"object"`
}

type TransactionLinkedRefund struct {
	TxnId     string                           `json:"txn_id"`
	TxnStatus TransactionLinkedRefundTxnStatus `json:"txn_status"`
	TxnDate   int64                            `json:"txn_date"`
	TxnAmount int64                            `json:"txn_amount"`
	Object    string                           `json:"object"`
}

type TransactionLinkedPayment struct {
	Id     string                         `json:"id"`
	Status TransactionLinkedPaymentStatus `json:"status"`
	Amount int64                          `json:"amount"`
	Date   int64                          `json:"date"`
	Object string                         `json:"object"`
}

type TransactionGatewayErrorDetail struct {
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

// operations
// input params
type TransactionCreateAuthorizationRequest struct {
	CustomerId      string `json:"customer_id"`
	PaymentSourceId string `json:"payment_source_id,omitempty"`
	CurrencyCode    string `json:"currency_code,omitempty"`
	Amount          *int64 `json:"amount"`
	apiRequest      `json:"-" form:"-"`
}

func (r *TransactionCreateAuthorizationRequest) payload() any { return r }

type TransactionRecordRefundRequest struct {
	Amount                *int64        `json:"amount,omitempty"`
	PaymentMethod         PaymentMethod `json:"payment_method"`
	Date                  *int64        `json:"date"`
	ReferenceNumber       string        `json:"reference_number,omitempty"`
	CustomPaymentMethodId string        `json:"custom_payment_method_id,omitempty"`
	Comment               string        `json:"comment,omitempty"`
	apiRequest            `json:"-" form:"-"`
}

func (r *TransactionRecordRefundRequest) payload() any { return r }

type TransactionReconcileRequest struct {
	IdAtGateway string            `json:"id_at_gateway,omitempty"`
	CustomerId  string            `json:"customer_id,omitempty"`
	Status      TransactionStatus `json:"status,omitempty"`
	apiRequest  `json:"-" form:"-"`
}

func (r *TransactionReconcileRequest) payload() any { return r }

type TransactionRefundRequest struct {
	Amount     *int64 `json:"amount,omitempty"`
	Comment    string `json:"comment,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *TransactionRefundRequest) payload() any { return r }

type TransactionListRequest struct {
	Limit            *int32           `json:"limit,omitempty"`
	Offset           string           `json:"offset,omitempty"`
	IncludeDeleted   *bool            `json:"include_deleted,omitempty"`
	Id               *StringFilter    `json:"id,omitempty"`
	CustomerId       *StringFilter    `json:"customer_id,omitempty"`
	SubscriptionId   *StringFilter    `json:"subscription_id,omitempty"`
	PaymentSourceId  *StringFilter    `json:"payment_source_id,omitempty"`
	PaymentMethod    *EnumFilter      `json:"payment_method,omitempty"`
	Gateway          *EnumFilter      `json:"gateway,omitempty"`
	GatewayAccountId *StringFilter    `json:"gateway_account_id,omitempty"`
	IdAtGateway      *StringFilter    `json:"id_at_gateway,omitempty"`
	ReferenceNumber  *StringFilter    `json:"reference_number,omitempty"`
	Type             *EnumFilter      `json:"type,omitempty"`
	Date             *TimestampFilter `json:"date,omitempty"`
	Amount           *NumberFilter    `json:"amount,omitempty"`
	AmountCapturable *NumberFilter    `json:"amount_capturable,omitempty"`
	Status           *EnumFilter      `json:"status,omitempty"`
	UpdatedAt        *TimestampFilter `json:"updated_at,omitempty"`
	SortBy           *SortFilter      `json:"sort_by,omitempty"`
	apiRequest       `json:"-" form:"-"`
}

func (r *TransactionListRequest) payload() any { return r }

type TransactionTransactionsForCustomerRequest struct {
	Limit      *int32 `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *TransactionTransactionsForCustomerRequest) payload() any { return r }

type TransactionTransactionsForSubscriptionRequest struct {
	Limit      *int32 `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *TransactionTransactionsForSubscriptionRequest) payload() any { return r }

type TransactionPaymentsForInvoiceRequest struct {
	Limit      *int32 `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *TransactionPaymentsForInvoiceRequest) payload() any { return r }

type TransactionDeleteOfflineTransactionRequest struct {
	Comment    string `json:"comment,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *TransactionDeleteOfflineTransactionRequest) payload() any { return r }

// operation response
type TransactionCreateAuthorizationResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
	apiResponse
}

// operation response
type TransactionVoidTransactionResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
	apiResponse
}

// operation response
type TransactionRecordRefundResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
	apiResponse
}

// operation response
type TransactionReconcileResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
	apiResponse
}

// operation response
type TransactionRefundResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
	apiResponse
}

// operation sub response
type TransactionListTransactionResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
}

// operation response
type TransactionListResponse struct {
	List       []*TransactionListTransactionResponse `json:"list,omitempty"`
	NextOffset string                                `json:"next_offset,omitempty"`
	apiResponse
}

// operation sub response
type TransactionTransactionsForCustomerTransactionResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
}

// operation response
type TransactionTransactionsForCustomerResponse struct {
	List       []*TransactionTransactionsForCustomerTransactionResponse `json:"list,omitempty"`
	NextOffset string                                                   `json:"next_offset,omitempty"`
	apiResponse
}

// operation sub response
type TransactionTransactionsForSubscriptionTransactionResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
}

// operation response
type TransactionTransactionsForSubscriptionResponse struct {
	List       []*TransactionTransactionsForSubscriptionTransactionResponse `json:"list,omitempty"`
	NextOffset string                                                       `json:"next_offset,omitempty"`
	apiResponse
}

// operation sub response
type TransactionPaymentsForInvoiceTransactionResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
}

// operation response
type TransactionPaymentsForInvoiceResponse struct {
	List       []*TransactionPaymentsForInvoiceTransactionResponse `json:"list,omitempty"`
	NextOffset string                                              `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type TransactionRetrieveResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
	apiResponse
}

// operation response
type TransactionDeleteOfflineTransactionResponse struct {
	Transaction *Transaction `json:"transaction,omitempty"`
	apiResponse
}
