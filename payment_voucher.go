package chargebee

type PaymentVoucherPaymentVoucherType string

const (
	PaymentVoucherPaymentVoucherTypeBoleto PaymentVoucherPaymentVoucherType = "boleto"
)

type PaymentVoucherStatus string

const (
	PaymentVoucherStatusActive   PaymentVoucherStatus = "active"
	PaymentVoucherStatusConsumed PaymentVoucherStatus = "consumed"
	PaymentVoucherStatusExpired  PaymentVoucherStatus = "expired"
	PaymentVoucherStatusFailure  PaymentVoucherStatus = "failure"
)

type PaymentVoucherGateway string

const (
	PaymentVoucherGatewayChargebee             PaymentVoucherGateway = "chargebee"
	PaymentVoucherGatewayChargebeePayments     PaymentVoucherGateway = "chargebee_payments"
	PaymentVoucherGatewayAdyen                 PaymentVoucherGateway = "adyen"
	PaymentVoucherGatewayStripe                PaymentVoucherGateway = "stripe"
	PaymentVoucherGatewayWepay                 PaymentVoucherGateway = "wepay"
	PaymentVoucherGatewayBraintree             PaymentVoucherGateway = "braintree"
	PaymentVoucherGatewayAuthorizeNet          PaymentVoucherGateway = "authorize_net"
	PaymentVoucherGatewayPaypalPro             PaymentVoucherGateway = "paypal_pro"
	PaymentVoucherGatewayPin                   PaymentVoucherGateway = "pin"
	PaymentVoucherGatewayEway                  PaymentVoucherGateway = "eway"
	PaymentVoucherGatewayEwayRapid             PaymentVoucherGateway = "eway_rapid"
	PaymentVoucherGatewayWorldpay              PaymentVoucherGateway = "worldpay"
	PaymentVoucherGatewayBalancedPayments      PaymentVoucherGateway = "balanced_payments"
	PaymentVoucherGatewayBeanstream            PaymentVoucherGateway = "beanstream"
	PaymentVoucherGatewayBluepay               PaymentVoucherGateway = "bluepay"
	PaymentVoucherGatewayElavon                PaymentVoucherGateway = "elavon"
	PaymentVoucherGatewayFirstDataGlobal       PaymentVoucherGateway = "first_data_global"
	PaymentVoucherGatewayHdfc                  PaymentVoucherGateway = "hdfc"
	PaymentVoucherGatewayMigs                  PaymentVoucherGateway = "migs"
	PaymentVoucherGatewayNmi                   PaymentVoucherGateway = "nmi"
	PaymentVoucherGatewayOgone                 PaymentVoucherGateway = "ogone"
	PaymentVoucherGatewayPaymill               PaymentVoucherGateway = "paymill"
	PaymentVoucherGatewayPaypalPayflowPro      PaymentVoucherGateway = "paypal_payflow_pro"
	PaymentVoucherGatewaySagePay               PaymentVoucherGateway = "sage_pay"
	PaymentVoucherGatewayTco                   PaymentVoucherGateway = "tco"
	PaymentVoucherGatewayWirecard              PaymentVoucherGateway = "wirecard"
	PaymentVoucherGatewayAmazonPayments        PaymentVoucherGateway = "amazon_payments"
	PaymentVoucherGatewayPaypalExpressCheckout PaymentVoucherGateway = "paypal_express_checkout"
	PaymentVoucherGatewayGocardless            PaymentVoucherGateway = "gocardless"
	PaymentVoucherGatewayOrbital               PaymentVoucherGateway = "orbital"
	PaymentVoucherGatewayMonerisUs             PaymentVoucherGateway = "moneris_us"
	PaymentVoucherGatewayMoneris               PaymentVoucherGateway = "moneris"
	PaymentVoucherGatewayBluesnap              PaymentVoucherGateway = "bluesnap"
	PaymentVoucherGatewayCybersource           PaymentVoucherGateway = "cybersource"
	PaymentVoucherGatewayVantiv                PaymentVoucherGateway = "vantiv"
	PaymentVoucherGatewayCheckoutCom           PaymentVoucherGateway = "checkout_com"
	PaymentVoucherGatewayPaypal                PaymentVoucherGateway = "paypal"
	PaymentVoucherGatewayIngenicoDirect        PaymentVoucherGateway = "ingenico_direct"
	PaymentVoucherGatewayExact                 PaymentVoucherGateway = "exact"
	PaymentVoucherGatewayMollie                PaymentVoucherGateway = "mollie"
	PaymentVoucherGatewayQuickbooks            PaymentVoucherGateway = "quickbooks"
	PaymentVoucherGatewayRazorpay              PaymentVoucherGateway = "razorpay"
	PaymentVoucherGatewayGlobalPayments        PaymentVoucherGateway = "global_payments"
	PaymentVoucherGatewayBankOfAmerica         PaymentVoucherGateway = "bank_of_america"
	PaymentVoucherGatewayEcentric              PaymentVoucherGateway = "ecentric"
	PaymentVoucherGatewayMetricsGlobal         PaymentVoucherGateway = "metrics_global"
	PaymentVoucherGatewayWindcave              PaymentVoucherGateway = "windcave"
	PaymentVoucherGatewayPayCom                PaymentVoucherGateway = "pay_com"
	PaymentVoucherGatewayEbanx                 PaymentVoucherGateway = "ebanx"
	PaymentVoucherGatewayDlocal                PaymentVoucherGateway = "dlocal"
	PaymentVoucherGatewayNuvei                 PaymentVoucherGateway = "nuvei"
	PaymentVoucherGatewaySolidgate             PaymentVoucherGateway = "solidgate"
	PaymentVoucherGatewayPaystack              PaymentVoucherGateway = "paystack"
	PaymentVoucherGatewayJpMorgan              PaymentVoucherGateway = "jp_morgan"
	PaymentVoucherGatewayDeutscheBank          PaymentVoucherGateway = "deutsche_bank"
	PaymentVoucherGatewayEzidebit              PaymentVoucherGateway = "ezidebit"
	PaymentVoucherGatewayNotApplicable         PaymentVoucherGateway = "not_applicable"
)

// just struct
type PaymentVoucher struct {
	Id                 string                           `json:"id"`
	IdAtGateway        string                           `json:"id_at_gateway"`
	PaymentVoucherType PaymentVoucherPaymentVoucherType `json:"payment_voucher_type"`
	ExpiresAt          int64                            `json:"expires_at"`
	Status             PaymentVoucherStatus             `json:"status"`
	SubscriptionId     string                           `json:"subscription_id"`
	CurrencyCode       string                           `json:"currency_code"`
	Amount             int64                            `json:"amount"`
	GatewayAccountId   string                           `json:"gateway_account_id"`
	PaymentSourceId    string                           `json:"payment_source_id"`
	Gateway            PaymentVoucherGateway            `json:"gateway"`
	Payload            string                           `json:"payload"`
	ErrorCode          string                           `json:"error_code"`
	ErrorText          string                           `json:"error_text"`
	Url                string                           `json:"url"`
	Date               int64                            `json:"date"`
	ResourceVersion    int64                            `json:"resource_version"`
	UpdatedAt          int64                            `json:"updated_at"`
	CustomerId         string                           `json:"customer_id"`
	LinkedInvoices     []*PaymentVoucherLinkedInvoice   `json:"linked_invoices"`
	Object             string                           `json:"object"`
}

// sub resources
type PaymentVoucherLinkedInvoice struct {
	InvoiceId string `json:"invoice_id"`
	TxnId     string `json:"txn_id"`
	AppliedAt int64  `json:"applied_at"`
	Object    string `json:"object"`
}

// operations
// input params
type PaymentVoucherCreateRequest struct {
	CustomerId           string                                    `json:"customer_id"`
	VoucherPaymentSource *PaymentVoucherCreateVoucherPaymentSource `json:"voucher_payment_source,omitempty"`
	InvoiceAllocations   []*PaymentVoucherCreateInvoiceAllocation  `json:"invoice_allocations,omitempty"`
	PaymentSourceId      string                                    `json:"payment_source_id,omitempty"`
	apiRequest           `json:"-" form:"-"`
}

func (r *PaymentVoucherCreateRequest) payload() any { return r }

// input sub resource params single
type PaymentVoucherCreateVoucherPaymentSource struct {
	VoucherType PaymentVoucherVoucherPaymentSourceVoucherType `json:"voucher_type"`
}

// input sub resource params multi
type PaymentVoucherCreateInvoiceAllocation struct {
	InvoiceId string `json:"invoice_id"`
}
type PaymentVoucherPaymentVouchersForInvoiceRequest struct {
	Limit      *int32      `json:"limit,omitempty"`
	Offset     string      `json:"offset,omitempty"`
	Status     *EnumFilter `json:"status,omitempty"`
	SortBy     *SortFilter `json:"sort_by,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *PaymentVoucherPaymentVouchersForInvoiceRequest) payload() any { return r }

type PaymentVoucherPaymentVouchersForCustomerRequest struct {
	Limit      *int32      `json:"limit,omitempty"`
	Offset     string      `json:"offset,omitempty"`
	Status     *EnumFilter `json:"status,omitempty"`
	SortBy     *SortFilter `json:"sort_by,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *PaymentVoucherPaymentVouchersForCustomerRequest) payload() any { return r }

// operation response
type PaymentVoucherCreateResponse struct {
	PaymentVoucher *PaymentVoucher `json:"payment_voucher,omitempty"`
	apiResponse
}

// operation response
type PaymentVoucherRetrieveResponse struct {
	PaymentVoucher *PaymentVoucher `json:"payment_voucher,omitempty"`
	apiResponse
}

// operation sub response
type PaymentVoucherPaymentVouchersForInvoicePaymentVoucherResponse struct {
	PaymentVoucher *PaymentVoucher `json:"payment_voucher,omitempty"`
}

// operation response
type PaymentVoucherPaymentVouchersForInvoiceResponse struct {
	List       []*PaymentVoucherPaymentVouchersForInvoicePaymentVoucherResponse `json:"list,omitempty"`
	NextOffset string                                                           `json:"next_offset,omitempty"`
	apiResponse
}

// operation sub response
type PaymentVoucherPaymentVouchersForCustomerPaymentVoucherResponse struct {
	PaymentVoucher *PaymentVoucher `json:"payment_voucher,omitempty"`
}

// operation response
type PaymentVoucherPaymentVouchersForCustomerResponse struct {
	List       []*PaymentVoucherPaymentVouchersForCustomerPaymentVoucherResponse `json:"list,omitempty"`
	NextOffset string                                                            `json:"next_offset,omitempty"`
	apiResponse
}
