package chargebee

type PaymentSourceType string

const (
	PaymentSourceTypeCard                  PaymentSourceType = "card"
	PaymentSourceTypePaypalExpressCheckout PaymentSourceType = "paypal_express_checkout"
	PaymentSourceTypeAmazonPayments        PaymentSourceType = "amazon_payments"
	PaymentSourceTypeDirectDebit           PaymentSourceType = "direct_debit"
	PaymentSourceTypeGeneric               PaymentSourceType = "generic"
	PaymentSourceTypeAlipay                PaymentSourceType = "alipay"
	PaymentSourceTypeUnionpay              PaymentSourceType = "unionpay"
	PaymentSourceTypeApplePay              PaymentSourceType = "apple_pay"
	PaymentSourceTypeWechatPay             PaymentSourceType = "wechat_pay"
	PaymentSourceTypeIdeal                 PaymentSourceType = "ideal"
	PaymentSourceTypeGooglePay             PaymentSourceType = "google_pay"
	PaymentSourceTypeSofort                PaymentSourceType = "sofort"
	PaymentSourceTypeBancontact            PaymentSourceType = "bancontact"
	PaymentSourceTypeGiropay               PaymentSourceType = "giropay"
	PaymentSourceTypeDotpay                PaymentSourceType = "dotpay"
	PaymentSourceTypeUpi                   PaymentSourceType = "upi"
	PaymentSourceTypeNetbankingEmandates   PaymentSourceType = "netbanking_emandates"
	PaymentSourceTypeVenmo                 PaymentSourceType = "venmo"
	PaymentSourceTypePayTo                 PaymentSourceType = "pay_to"
	PaymentSourceTypeFasterPayments        PaymentSourceType = "faster_payments"
	PaymentSourceTypeSepaInstantTransfer   PaymentSourceType = "sepa_instant_transfer"
	PaymentSourceTypeAutomatedBankTransfer PaymentSourceType = "automated_bank_transfer"
	PaymentSourceTypeKlarnaPayNow          PaymentSourceType = "klarna_pay_now"
	PaymentSourceTypeOnlineBankingPoland   PaymentSourceType = "online_banking_poland"
	PaymentSourceTypePayconiqByBancontact  PaymentSourceType = "payconiq_by_bancontact"
)

type PaymentSourceStatus string

const (
	PaymentSourceStatusValid               PaymentSourceStatus = "valid"
	PaymentSourceStatusExpiring            PaymentSourceStatus = "expiring"
	PaymentSourceStatusExpired             PaymentSourceStatus = "expired"
	PaymentSourceStatusInvalid             PaymentSourceStatus = "invalid"
	PaymentSourceStatusPendingVerification PaymentSourceStatus = "pending_verification"
)

type PaymentSourceGateway string

const (
	PaymentSourceGatewayChargebee             PaymentSourceGateway = "chargebee"
	PaymentSourceGatewayChargebeePayments     PaymentSourceGateway = "chargebee_payments"
	PaymentSourceGatewayAdyen                 PaymentSourceGateway = "adyen"
	PaymentSourceGatewayStripe                PaymentSourceGateway = "stripe"
	PaymentSourceGatewayWepay                 PaymentSourceGateway = "wepay"
	PaymentSourceGatewayBraintree             PaymentSourceGateway = "braintree"
	PaymentSourceGatewayAuthorizeNet          PaymentSourceGateway = "authorize_net"
	PaymentSourceGatewayPaypalPro             PaymentSourceGateway = "paypal_pro"
	PaymentSourceGatewayPin                   PaymentSourceGateway = "pin"
	PaymentSourceGatewayEway                  PaymentSourceGateway = "eway"
	PaymentSourceGatewayEwayRapid             PaymentSourceGateway = "eway_rapid"
	PaymentSourceGatewayWorldpay              PaymentSourceGateway = "worldpay"
	PaymentSourceGatewayBalancedPayments      PaymentSourceGateway = "balanced_payments"
	PaymentSourceGatewayBeanstream            PaymentSourceGateway = "beanstream"
	PaymentSourceGatewayBluepay               PaymentSourceGateway = "bluepay"
	PaymentSourceGatewayElavon                PaymentSourceGateway = "elavon"
	PaymentSourceGatewayFirstDataGlobal       PaymentSourceGateway = "first_data_global"
	PaymentSourceGatewayHdfc                  PaymentSourceGateway = "hdfc"
	PaymentSourceGatewayMigs                  PaymentSourceGateway = "migs"
	PaymentSourceGatewayNmi                   PaymentSourceGateway = "nmi"
	PaymentSourceGatewayOgone                 PaymentSourceGateway = "ogone"
	PaymentSourceGatewayPaymill               PaymentSourceGateway = "paymill"
	PaymentSourceGatewayPaypalPayflowPro      PaymentSourceGateway = "paypal_payflow_pro"
	PaymentSourceGatewaySagePay               PaymentSourceGateway = "sage_pay"
	PaymentSourceGatewayTco                   PaymentSourceGateway = "tco"
	PaymentSourceGatewayWirecard              PaymentSourceGateway = "wirecard"
	PaymentSourceGatewayAmazonPayments        PaymentSourceGateway = "amazon_payments"
	PaymentSourceGatewayPaypalExpressCheckout PaymentSourceGateway = "paypal_express_checkout"
	PaymentSourceGatewayGocardless            PaymentSourceGateway = "gocardless"
	PaymentSourceGatewayOrbital               PaymentSourceGateway = "orbital"
	PaymentSourceGatewayMonerisUs             PaymentSourceGateway = "moneris_us"
	PaymentSourceGatewayMoneris               PaymentSourceGateway = "moneris"
	PaymentSourceGatewayBluesnap              PaymentSourceGateway = "bluesnap"
	PaymentSourceGatewayCybersource           PaymentSourceGateway = "cybersource"
	PaymentSourceGatewayVantiv                PaymentSourceGateway = "vantiv"
	PaymentSourceGatewayCheckoutCom           PaymentSourceGateway = "checkout_com"
	PaymentSourceGatewayPaypal                PaymentSourceGateway = "paypal"
	PaymentSourceGatewayIngenicoDirect        PaymentSourceGateway = "ingenico_direct"
	PaymentSourceGatewayExact                 PaymentSourceGateway = "exact"
	PaymentSourceGatewayMollie                PaymentSourceGateway = "mollie"
	PaymentSourceGatewayQuickbooks            PaymentSourceGateway = "quickbooks"
	PaymentSourceGatewayRazorpay              PaymentSourceGateway = "razorpay"
	PaymentSourceGatewayGlobalPayments        PaymentSourceGateway = "global_payments"
	PaymentSourceGatewayBankOfAmerica         PaymentSourceGateway = "bank_of_america"
	PaymentSourceGatewayEcentric              PaymentSourceGateway = "ecentric"
	PaymentSourceGatewayMetricsGlobal         PaymentSourceGateway = "metrics_global"
	PaymentSourceGatewayWindcave              PaymentSourceGateway = "windcave"
	PaymentSourceGatewayPayCom                PaymentSourceGateway = "pay_com"
	PaymentSourceGatewayEbanx                 PaymentSourceGateway = "ebanx"
	PaymentSourceGatewayDlocal                PaymentSourceGateway = "dlocal"
	PaymentSourceGatewayNuvei                 PaymentSourceGateway = "nuvei"
	PaymentSourceGatewaySolidgate             PaymentSourceGateway = "solidgate"
	PaymentSourceGatewayPaystack              PaymentSourceGateway = "paystack"
	PaymentSourceGatewayJpMorgan              PaymentSourceGateway = "jp_morgan"
	PaymentSourceGatewayDeutscheBank          PaymentSourceGateway = "deutsche_bank"
	PaymentSourceGatewayEzidebit              PaymentSourceGateway = "ezidebit"
	PaymentSourceGatewayNotApplicable         PaymentSourceGateway = "not_applicable"
)

type PaymentSourceCardBrand string

const (
	PaymentSourceCardBrandVisa            PaymentSourceCardBrand = "visa"
	PaymentSourceCardBrandMastercard      PaymentSourceCardBrand = "mastercard"
	PaymentSourceCardBrandAmericanExpress PaymentSourceCardBrand = "american_express"
	PaymentSourceCardBrandDiscover        PaymentSourceCardBrand = "discover"
	PaymentSourceCardBrandJcb             PaymentSourceCardBrand = "jcb"
	PaymentSourceCardBrandDinersClub      PaymentSourceCardBrand = "diners_club"
	PaymentSourceCardBrandOther           PaymentSourceCardBrand = "other"
	PaymentSourceCardBrandBancontact      PaymentSourceCardBrand = "bancontact"
	PaymentSourceCardBrandCmrFalabella    PaymentSourceCardBrand = "cmr_falabella"
	PaymentSourceCardBrandTarjetaNaranja  PaymentSourceCardBrand = "tarjeta_naranja"
	PaymentSourceCardBrandNativa          PaymentSourceCardBrand = "nativa"
	PaymentSourceCardBrandCencosud        PaymentSourceCardBrand = "cencosud"
	PaymentSourceCardBrandCabal           PaymentSourceCardBrand = "cabal"
	PaymentSourceCardBrandArgencard       PaymentSourceCardBrand = "argencard"
	PaymentSourceCardBrandElo             PaymentSourceCardBrand = "elo"
	PaymentSourceCardBrandHipercard       PaymentSourceCardBrand = "hipercard"
	PaymentSourceCardBrandCarnet          PaymentSourceCardBrand = "carnet"
	PaymentSourceCardBrandRupay           PaymentSourceCardBrand = "rupay"
	PaymentSourceCardBrandMaestro         PaymentSourceCardBrand = "maestro"
	PaymentSourceCardBrandDankort         PaymentSourceCardBrand = "dankort"
	PaymentSourceCardBrandCartesBancaires PaymentSourceCardBrand = "cartes_bancaires"
	PaymentSourceCardBrandNotApplicable   PaymentSourceCardBrand = "not_applicable"
)

type PaymentSourceCardFundingType string

const (
	PaymentSourceCardFundingTypeCredit        PaymentSourceCardFundingType = "credit"
	PaymentSourceCardFundingTypeDebit         PaymentSourceCardFundingType = "debit"
	PaymentSourceCardFundingTypePrepaid       PaymentSourceCardFundingType = "prepaid"
	PaymentSourceCardFundingTypeNotKnown      PaymentSourceCardFundingType = "not_known"
	PaymentSourceCardFundingTypeNotApplicable PaymentSourceCardFundingType = "not_applicable"
)

type PaymentSourceBankAccountDirectDebitScheme string

const (
	PaymentSourceBankAccountDirectDebitSchemeAch           PaymentSourceBankAccountDirectDebitScheme = "ach"
	PaymentSourceBankAccountDirectDebitSchemeBacs          PaymentSourceBankAccountDirectDebitScheme = "bacs"
	PaymentSourceBankAccountDirectDebitSchemeSepaCore      PaymentSourceBankAccountDirectDebitScheme = "sepa_core"
	PaymentSourceBankAccountDirectDebitSchemeAutogiro      PaymentSourceBankAccountDirectDebitScheme = "autogiro"
	PaymentSourceBankAccountDirectDebitSchemeBecs          PaymentSourceBankAccountDirectDebitScheme = "becs"
	PaymentSourceBankAccountDirectDebitSchemeBecsNz        PaymentSourceBankAccountDirectDebitScheme = "becs_nz"
	PaymentSourceBankAccountDirectDebitSchemePad           PaymentSourceBankAccountDirectDebitScheme = "pad"
	PaymentSourceBankAccountDirectDebitSchemeNotApplicable PaymentSourceBankAccountDirectDebitScheme = "not_applicable"
)

type PaymentSourceBankAccountAccountType string

const (
	PaymentSourceBankAccountAccountTypeChecking         PaymentSourceBankAccountAccountType = "checking"
	PaymentSourceBankAccountAccountTypeSavings          PaymentSourceBankAccountAccountType = "savings"
	PaymentSourceBankAccountAccountTypeBusinessChecking PaymentSourceBankAccountAccountType = "business_checking"
	PaymentSourceBankAccountAccountTypeCurrent          PaymentSourceBankAccountAccountType = "current"
)

type PaymentSourceBankAccountEcheckType string

const (
	PaymentSourceBankAccountEcheckTypeWeb PaymentSourceBankAccountEcheckType = "web"
	PaymentSourceBankAccountEcheckTypePpd PaymentSourceBankAccountEcheckType = "ppd"
	PaymentSourceBankAccountEcheckTypeCcd PaymentSourceBankAccountEcheckType = "ccd"
)

type PaymentSourceBankAccountAccountHolderType string

const (
	PaymentSourceBankAccountAccountHolderTypeIndividual PaymentSourceBankAccountAccountHolderType = "individual"
	PaymentSourceBankAccountAccountHolderTypeCompany    PaymentSourceBankAccountAccountHolderType = "company"
)

type PaymentSourceBillingAddressValidationStatus string

const (
	PaymentSourceBillingAddressValidationStatusNotValidated   PaymentSourceBillingAddressValidationStatus = "not_validated"
	PaymentSourceBillingAddressValidationStatusValid          PaymentSourceBillingAddressValidationStatus = "valid"
	PaymentSourceBillingAddressValidationStatusPartiallyValid PaymentSourceBillingAddressValidationStatus = "partially_valid"
	PaymentSourceBillingAddressValidationStatusInvalid        PaymentSourceBillingAddressValidationStatus = "invalid"
)

// just struct
type PaymentSource struct {
	Id               string                          `json:"id"`
	ResourceVersion  int64                           `json:"resource_version"`
	UpdatedAt        int64                           `json:"updated_at"`
	CreatedAt        int64                           `json:"created_at"`
	CustomerId       string                          `json:"customer_id"`
	Type             PaymentSourceType               `json:"type"`
	ReferenceId      string                          `json:"reference_id"`
	Status           PaymentSourceStatus             `json:"status"`
	Gateway          PaymentSourceGateway            `json:"gateway"`
	GatewayAccountId string                          `json:"gateway_account_id"`
	IpAddress        string                          `json:"ip_address"`
	IssuingCountry   string                          `json:"issuing_country"`
	Card             *PaymentSourceCard              `json:"card"`
	BankAccount      *PaymentSourceBankAccount       `json:"bank_account"`
	Boleto           *PaymentSourceCustVoucherSource `json:"boleto"`
	BillingAddress   *PaymentSourceBillingAddress    `json:"billing_address"`
	AmazonPayment    *PaymentSourceAmazonPayment     `json:"amazon_payment"`
	Upi              *PaymentSourceUpi               `json:"upi"`
	Paypal           *PaymentSourcePaypal            `json:"paypal"`
	Venmo            *PaymentSourceVenmo             `json:"venmo"`
	KlarnaPayNow     *PaymentSourceKlarnaPayNow      `json:"klarna_pay_now"`
	Mandates         []*PaymentSourceMandate         `json:"mandates"`
	Deleted          bool                            `json:"deleted"`
	BusinessEntityId string                          `json:"business_entity_id"`
	Object           string                          `json:"object"`
}

// sub resources
type PaymentSourceCard struct {
	FirstName        string                       `json:"first_name"`
	LastName         string                       `json:"last_name"`
	Iin              string                       `json:"iin"`
	Last4            string                       `json:"last4"`
	Brand            PaymentSourceCardBrand       `json:"brand"`
	FundingType      PaymentSourceCardFundingType `json:"funding_type"`
	ExpiryMonth      int32                        `json:"expiry_month"`
	ExpiryYear       int32                        `json:"expiry_year"`
	BillingAddr1     string                       `json:"billing_addr1"`
	BillingAddr2     string                       `json:"billing_addr2"`
	BillingCity      string                       `json:"billing_city"`
	BillingStateCode string                       `json:"billing_state_code"`
	BillingState     string                       `json:"billing_state"`
	BillingCountry   string                       `json:"billing_country"`
	BillingZip       string                       `json:"billing_zip"`
	MaskedNumber     string                       `json:"masked_number"`
	Object           string                       `json:"object"`
}
type PaymentSourceBankAccount struct {
	Last4             string                                    `json:"last4"`
	NameOnAccount     string                                    `json:"name_on_account"`
	FirstName         string                                    `json:"first_name"`
	LastName          string                                    `json:"last_name"`
	DirectDebitScheme PaymentSourceBankAccountDirectDebitScheme `json:"direct_debit_scheme"`
	BankName          string                                    `json:"bank_name"`
	MandateId         string                                    `json:"mandate_id"`
	AccountType       PaymentSourceBankAccountAccountType       `json:"account_type"`
	EcheckType        PaymentSourceBankAccountEcheckType        `json:"echeck_type"`
	AccountHolderType PaymentSourceBankAccountAccountHolderType `json:"account_holder_type"`
	Email             string                                    `json:"email"`
	Object            string                                    `json:"object"`
}
type PaymentSourceCustVoucherSource struct {
	Last4     string `json:"last4"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Object    string `json:"object"`
}
type PaymentSourceBillingAddress struct {
	FirstName        string                                      `json:"first_name"`
	LastName         string                                      `json:"last_name"`
	Email            string                                      `json:"email"`
	Company          string                                      `json:"company"`
	Phone            string                                      `json:"phone"`
	Line1            string                                      `json:"line1"`
	Line2            string                                      `json:"line2"`
	Line3            string                                      `json:"line3"`
	City             string                                      `json:"city"`
	StateCode        string                                      `json:"state_code"`
	State            string                                      `json:"state"`
	Country          string                                      `json:"country"`
	Zip              string                                      `json:"zip"`
	ValidationStatus PaymentSourceBillingAddressValidationStatus `json:"validation_status"`
	Object           string                                      `json:"object"`
}
type PaymentSourceAmazonPayment struct {
	Email       string `json:"email"`
	AgreementId string `json:"agreement_id"`
	Object      string `json:"object"`
}
type PaymentSourceUpi struct {
	Vpa    string `json:"vpa"`
	Object string `json:"object"`
}
type PaymentSourcePaypal struct {
	Email       string `json:"email"`
	AgreementId string `json:"agreement_id"`
	Object      string `json:"object"`
}
type PaymentSourceVenmo struct {
	UserName string `json:"user_name"`
	Object   string `json:"object"`
}
type PaymentSourceKlarnaPayNow struct {
	Email  string `json:"email"`
	Object string `json:"object"`
}
type PaymentSourceMandate struct {
	Id             string `json:"id"`
	SubscriptionId string `json:"subscription_id"`
	CreatedAt      int64  `json:"created_at"`
	Object         string `json:"object"`
}

// operations
// input params
type PaymentSourceCreateUsingTempTokenRequest struct {
	CustomerId                  string                 `json:"customer_id"`
	GatewayAccountId            string                 `json:"gateway_account_id,omitempty"`
	Type                        PaymentSourceType      `json:"type"`
	TmpToken                    string                 `json:"tmp_token"`
	IssuingCountry              string                 `json:"issuing_country,omitempty"`
	ReplacePrimaryPaymentSource *bool                  `json:"replace_primary_payment_source,omitempty"`
	AdditionalInformation       map[string]interface{} `json:"additional_information,omitempty"`
	apiRequest                  `json:"-" form:"-"`
}

func (r *PaymentSourceCreateUsingTempTokenRequest) payload() any { return r }

type PaymentSourceCreateUsingPermanentTokenRequest struct {
	CustomerId                  string                                                `json:"customer_id"`
	Type                        PaymentSourceType                                     `json:"type"`
	GatewayAccountId            string                                                `json:"gateway_account_id,omitempty"`
	ReferenceId                 string                                                `json:"reference_id,omitempty"`
	IssuingCountry              string                                                `json:"issuing_country,omitempty"`
	ReplacePrimaryPaymentSource *bool                                                 `json:"replace_primary_payment_source,omitempty"`
	PaymentMethodToken          string                                                `json:"payment_method_token,omitempty"`
	CustomerProfileToken        string                                                `json:"customer_profile_token,omitempty"`
	NetworkTransactionId        string                                                `json:"network_transaction_id,omitempty"`
	MandateId                   string                                                `json:"mandate_id,omitempty"`
	SkipRetrieval               *bool                                                 `json:"skip_retrieval,omitempty"`
	Card                        *PaymentSourceCreateUsingPermanentTokenCard           `json:"card,omitempty"`
	BillingAddress              *PaymentSourceCreateUsingPermanentTokenBillingAddress `json:"billing_address,omitempty"`
	AdditionalInformation       map[string]interface{}                                `json:"additional_information,omitempty"`
	apiRequest                  `json:"-" form:"-"`
}

func (r *PaymentSourceCreateUsingPermanentTokenRequest) payload() any { return r }

// input sub resource params single
type PaymentSourceCreateUsingPermanentTokenCard struct {
	Last4       string          `json:"last4,omitempty"`
	Iin         string          `json:"iin,omitempty"`
	ExpiryMonth *int32          `json:"expiry_month,omitempty"`
	ExpiryYear  *int32          `json:"expiry_year,omitempty"`
	Brand       CardBrand       `json:"brand,omitempty"`
	FundingType CardFundingType `json:"funding_type,omitempty"`
}

// input sub resource params single
type PaymentSourceCreateUsingPermanentTokenBillingAddress struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Line1     string `json:"line1,omitempty"`
	Line2     string `json:"line2,omitempty"`
	Line3     string `json:"line3,omitempty"`
	City      string `json:"city,omitempty"`
	StateCode string `json:"state_code,omitempty"`
	State     string `json:"state,omitempty"`
	Zip       string `json:"zip,omitempty"`
	Country   string `json:"country,omitempty"`
}
type PaymentSourceCreateUsingTokenRequest struct {
	CustomerId                  string `json:"customer_id"`
	ReplacePrimaryPaymentSource *bool  `json:"replace_primary_payment_source,omitempty"`
	TokenId                     string `json:"token_id"`
	apiRequest                  `json:"-" form:"-"`
}

func (r *PaymentSourceCreateUsingTokenRequest) payload() any { return r }

type PaymentSourceCreateUsingPaymentIntentRequest struct {
	CustomerId                  string                                              `json:"customer_id"`
	PaymentIntent               *PaymentSourceCreateUsingPaymentIntentPaymentIntent `json:"payment_intent,omitempty"`
	ReplacePrimaryPaymentSource *bool                                               `json:"replace_primary_payment_source,omitempty"`
	apiRequest                  `json:"-" form:"-"`
}

func (r *PaymentSourceCreateUsingPaymentIntentRequest) payload() any { return r }

// input sub resource params single
type PaymentSourceCreateUsingPaymentIntentPaymentIntent struct {
	Id                    string                 `json:"id,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	GwToken               string                 `json:"gw_token,omitempty"`
	PaymentMethodType     PaymentMethodType      `json:"payment_method_type,omitempty"`
	ReferenceId           string                 `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                 `json:"gw_payment_method_id,omitempty"`
	AdditionalInfo        map[string]interface{} `json:"additional_info,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type PaymentSourceCreateVoucherPaymentSourceRequest struct {
	CustomerId           string                                                       `json:"customer_id"`
	VoucherPaymentSource *PaymentSourceCreateVoucherPaymentSourceVoucherPaymentSource `json:"voucher_payment_source,omitempty"`
	apiRequest           `json:"-" form:"-"`
}

func (r *PaymentSourceCreateVoucherPaymentSourceRequest) payload() any { return r }

// input sub resource params single
type PaymentSourceCreateVoucherPaymentSourceVoucherPaymentSource struct {
	VoucherType      VoucherType            `json:"voucher_type"`
	GatewayAccountId string                 `json:"gateway_account_id,omitempty"`
	TaxId            string                 `json:"tax_id,omitempty"`
	BillingAddress   map[string]interface{} `json:"billing_address,omitempty"`
}
type PaymentSourceCreateCardRequest struct {
	CustomerId                  string                       `json:"customer_id"`
	Card                        *PaymentSourceCreateCardCard `json:"card,omitempty"`
	ReplacePrimaryPaymentSource *bool                        `json:"replace_primary_payment_source,omitempty"`
	apiRequest                  `json:"-" form:"-"`
}

func (r *PaymentSourceCreateCardRequest) payload() any { return r }

// input sub resource params single
type PaymentSourceCreateCardCard struct {
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	Number                string                 `json:"number"`
	ExpiryMonth           *int32                 `json:"expiry_month"`
	ExpiryYear            *int32                 `json:"expiry_year"`
	Cvv                   string                 `json:"cvv,omitempty"`
	PreferredScheme       PreferredScheme        `json:"preferred_scheme,omitempty"`
	BillingAddr1          string                 `json:"billing_addr1,omitempty"`
	BillingAddr2          string                 `json:"billing_addr2,omitempty"`
	BillingCity           string                 `json:"billing_city,omitempty"`
	BillingStateCode      string                 `json:"billing_state_code,omitempty"`
	BillingState          string                 `json:"billing_state,omitempty"`
	BillingZip            string                 `json:"billing_zip,omitempty"`
	BillingCountry        string                 `json:"billing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type PaymentSourceCreateBankAccountRequest struct {
	CustomerId                  string                                     `json:"customer_id"`
	BankAccount                 *PaymentSourceCreateBankAccountBankAccount `json:"bank_account,omitempty"`
	IssuingCountry              string                                     `json:"issuing_country,omitempty"`
	ReplacePrimaryPaymentSource *bool                                      `json:"replace_primary_payment_source,omitempty"`
	apiRequest                  `json:"-" form:"-"`
}

func (r *PaymentSourceCreateBankAccountRequest) payload() any { return r }

// input sub resource params single
type PaymentSourceCreateBankAccountBankAccount struct {
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	Iban                  string                 `json:"iban,omitempty"`
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	Company               string                 `json:"company,omitempty"`
	Email                 string                 `json:"email,omitempty"`
	Phone                 string                 `json:"phone,omitempty"`
	BankName              string                 `json:"bank_name,omitempty"`
	AccountNumber         string                 `json:"account_number,omitempty"`
	RoutingNumber         string                 `json:"routing_number,omitempty"`
	BankCode              string                 `json:"bank_code,omitempty"`
	AccountType           AccountType            `json:"account_type,omitempty"`
	AccountHolderType     AccountHolderType      `json:"account_holder_type,omitempty"`
	EcheckType            EcheckType             `json:"echeck_type,omitempty"`
	SwedishIdentityNumber string                 `json:"swedish_identity_number,omitempty"`
	BillingAddress        map[string]interface{} `json:"billing_address,omitempty"`
}
type PaymentSourceUpdateCardRequest struct {
	Card                 *PaymentSourceUpdateCardCard `json:"card,omitempty"`
	GatewayMetaData      map[string]interface{}       `json:"gateway_meta_data,omitempty"`
	ReferenceTransaction string                       `json:"reference_transaction,omitempty"`
	apiRequest           `json:"-" form:"-"`
}

func (r *PaymentSourceUpdateCardRequest) payload() any { return r }

// input sub resource params single
type PaymentSourceUpdateCardCard struct {
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	ExpiryMonth           *int32                 `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                 `json:"expiry_year,omitempty"`
	BillingAddr1          string                 `json:"billing_addr1,omitempty"`
	BillingAddr2          string                 `json:"billing_addr2,omitempty"`
	BillingCity           string                 `json:"billing_city,omitempty"`
	BillingZip            string                 `json:"billing_zip,omitempty"`
	BillingStateCode      string                 `json:"billing_state_code,omitempty"`
	BillingState          string                 `json:"billing_state,omitempty"`
	BillingCountry        string                 `json:"billing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type PaymentSourceUpdateBankAccountRequest struct {
	BankAccount *PaymentSourceUpdateBankAccountBankAccount `json:"bank_account,omitempty"`
	apiRequest  `json:"-" form:"-"`
}

func (r *PaymentSourceUpdateBankAccountRequest) payload() any { return r }

// input sub resource params single
type PaymentSourceUpdateBankAccountBankAccount struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}
type PaymentSourceVerifyBankAccountRequest struct {
	Amount1    *int64 `json:"amount1"`
	Amount2    *int64 `json:"amount2"`
	apiRequest `json:"-" form:"-"`
}

func (r *PaymentSourceVerifyBankAccountRequest) payload() any { return r }

type PaymentSourceListRequest struct {
	Limit          *int32           `json:"limit,omitempty"`
	Offset         string           `json:"offset,omitempty"`
	SubscriptionId string           `json:"subscription_id,omitempty"`
	IncludeDeleted *bool            `json:"include_deleted,omitempty"`
	CustomerId     *StringFilter    `json:"customer_id,omitempty"`
	Type           *EnumFilter      `json:"type,omitempty"`
	Status         *EnumFilter      `json:"status,omitempty"`
	UpdatedAt      *TimestampFilter `json:"updated_at,omitempty"`
	CreatedAt      *TimestampFilter `json:"created_at,omitempty"`
	SortBy         *SortFilter      `json:"sort_by,omitempty"`
	apiRequest     `json:"-" form:"-"`
}

func (r *PaymentSourceListRequest) payload() any { return r }

type PaymentSourceSwitchGatewayAccountRequest struct {
	GatewayAccountId string `json:"gateway_account_id"`
	apiRequest       `json:"-" form:"-"`
}

func (r *PaymentSourceSwitchGatewayAccountRequest) payload() any { return r }

type PaymentSourceExportPaymentSourceRequest struct {
	GatewayAccountId string `json:"gateway_account_id"`
	apiRequest       `json:"-" form:"-"`
}

func (r *PaymentSourceExportPaymentSourceRequest) payload() any { return r }

// operation response
type PaymentSourceCreateUsingTempTokenResponse struct {
	Customer      Customer       `json:"customer,omitempty"`
	PaymentSource *PaymentSource `json:"payment_source,omitempty"`
	apiResponse
}

// operation response
type PaymentSourceCreateUsingPermanentTokenResponse struct {
	Customer      Customer       `json:"customer,omitempty"`
	PaymentSource *PaymentSource `json:"payment_source,omitempty"`
	apiResponse
}

// operation response
type PaymentSourceCreateUsingTokenResponse struct {
	Customer      Customer       `json:"customer,omitempty"`
	PaymentSource *PaymentSource `json:"payment_source,omitempty"`
	apiResponse
}

// operation response
type PaymentSourceCreateUsingPaymentIntentResponse struct {
	Customer      Customer       `json:"customer,omitempty"`
	PaymentSource *PaymentSource `json:"payment_source,omitempty"`
	apiResponse
}

// operation response
type PaymentSourceCreateVoucherPaymentSourceResponse struct {
	Customer      Customer       `json:"customer,omitempty"`
	PaymentSource *PaymentSource `json:"payment_source,omitempty"`
	apiResponse
}

// operation response
type PaymentSourceCreateCardResponse struct {
	Customer      Customer       `json:"customer,omitempty"`
	PaymentSource *PaymentSource `json:"payment_source,omitempty"`
	apiResponse
}

// operation response
type PaymentSourceCreateBankAccountResponse struct {
	Customer      Customer       `json:"customer,omitempty"`
	PaymentSource *PaymentSource `json:"payment_source,omitempty"`
	apiResponse
}

// operation response
type PaymentSourceUpdateCardResponse struct {
	Customer      Customer       `json:"customer,omitempty"`
	PaymentSource *PaymentSource `json:"payment_source,omitempty"`
	apiResponse
}

// operation response
type PaymentSourceUpdateBankAccountResponse struct {
	Customer      Customer       `json:"customer,omitempty"`
	PaymentSource *PaymentSource `json:"payment_source,omitempty"`
	apiResponse
}

// operation response
type PaymentSourceVerifyBankAccountResponse struct {
	PaymentSource *PaymentSource `json:"payment_source,omitempty"`
	apiResponse
}

// operation response
type PaymentSourceRetrieveResponse struct {
	PaymentSource *PaymentSource `json:"payment_source,omitempty"`
	apiResponse
}

// operation sub response
type PaymentSourceListPaymentSourceResponse struct {
	PaymentSource *PaymentSource `json:"payment_source,omitempty"`
}

// operation response
type PaymentSourceListResponse struct {
	List       []*PaymentSourceListPaymentSourceResponse `json:"list,omitempty"`
	NextOffset string                                    `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type PaymentSourceSwitchGatewayAccountResponse struct {
	Customer      Customer       `json:"customer,omitempty"`
	PaymentSource *PaymentSource `json:"payment_source,omitempty"`
	apiResponse
}

// operation response
type PaymentSourceExportPaymentSourceResponse struct {
	ThirdPartyPaymentMethod ThirdPartyPaymentMethod `json:"third_party_payment_method,omitempty"`
	apiResponse
}

// operation response
type PaymentSourceDeleteResponse struct {
	Customer      Customer       `json:"customer,omitempty"`
	PaymentSource *PaymentSource `json:"payment_source,omitempty"`
	apiResponse
}

// operation response
type PaymentSourceDeleteLocalResponse struct {
	Customer      Customer       `json:"customer,omitempty"`
	PaymentSource *PaymentSource `json:"payment_source,omitempty"`
	apiResponse
}
