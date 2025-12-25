package chargebee

type TokenGateway string

const (
	TokenGatewayChargebee             TokenGateway = "chargebee"
	TokenGatewayChargebeePayments     TokenGateway = "chargebee_payments"
	TokenGatewayAdyen                 TokenGateway = "adyen"
	TokenGatewayStripe                TokenGateway = "stripe"
	TokenGatewayWepay                 TokenGateway = "wepay"
	TokenGatewayBraintree             TokenGateway = "braintree"
	TokenGatewayAuthorizeNet          TokenGateway = "authorize_net"
	TokenGatewayPaypalPro             TokenGateway = "paypal_pro"
	TokenGatewayPin                   TokenGateway = "pin"
	TokenGatewayEway                  TokenGateway = "eway"
	TokenGatewayEwayRapid             TokenGateway = "eway_rapid"
	TokenGatewayWorldpay              TokenGateway = "worldpay"
	TokenGatewayBalancedPayments      TokenGateway = "balanced_payments"
	TokenGatewayBeanstream            TokenGateway = "beanstream"
	TokenGatewayBluepay               TokenGateway = "bluepay"
	TokenGatewayElavon                TokenGateway = "elavon"
	TokenGatewayFirstDataGlobal       TokenGateway = "first_data_global"
	TokenGatewayHdfc                  TokenGateway = "hdfc"
	TokenGatewayMigs                  TokenGateway = "migs"
	TokenGatewayNmi                   TokenGateway = "nmi"
	TokenGatewayOgone                 TokenGateway = "ogone"
	TokenGatewayPaymill               TokenGateway = "paymill"
	TokenGatewayPaypalPayflowPro      TokenGateway = "paypal_payflow_pro"
	TokenGatewaySagePay               TokenGateway = "sage_pay"
	TokenGatewayTco                   TokenGateway = "tco"
	TokenGatewayWirecard              TokenGateway = "wirecard"
	TokenGatewayAmazonPayments        TokenGateway = "amazon_payments"
	TokenGatewayPaypalExpressCheckout TokenGateway = "paypal_express_checkout"
	TokenGatewayGocardless            TokenGateway = "gocardless"
	TokenGatewayOrbital               TokenGateway = "orbital"
	TokenGatewayMonerisUs             TokenGateway = "moneris_us"
	TokenGatewayMoneris               TokenGateway = "moneris"
	TokenGatewayBluesnap              TokenGateway = "bluesnap"
	TokenGatewayCybersource           TokenGateway = "cybersource"
	TokenGatewayVantiv                TokenGateway = "vantiv"
	TokenGatewayCheckoutCom           TokenGateway = "checkout_com"
	TokenGatewayPaypal                TokenGateway = "paypal"
	TokenGatewayIngenicoDirect        TokenGateway = "ingenico_direct"
	TokenGatewayExact                 TokenGateway = "exact"
	TokenGatewayMollie                TokenGateway = "mollie"
	TokenGatewayQuickbooks            TokenGateway = "quickbooks"
	TokenGatewayRazorpay              TokenGateway = "razorpay"
	TokenGatewayGlobalPayments        TokenGateway = "global_payments"
	TokenGatewayBankOfAmerica         TokenGateway = "bank_of_america"
	TokenGatewayEcentric              TokenGateway = "ecentric"
	TokenGatewayMetricsGlobal         TokenGateway = "metrics_global"
	TokenGatewayWindcave              TokenGateway = "windcave"
	TokenGatewayPayCom                TokenGateway = "pay_com"
	TokenGatewayEbanx                 TokenGateway = "ebanx"
	TokenGatewayDlocal                TokenGateway = "dlocal"
	TokenGatewayNuvei                 TokenGateway = "nuvei"
	TokenGatewaySolidgate             TokenGateway = "solidgate"
	TokenGatewayPaystack              TokenGateway = "paystack"
	TokenGatewayJpMorgan              TokenGateway = "jp_morgan"
	TokenGatewayDeutscheBank          TokenGateway = "deutsche_bank"
	TokenGatewayEzidebit              TokenGateway = "ezidebit"
	TokenGatewayNotApplicable         TokenGateway = "not_applicable"
)

type TokenPaymentMethodType string

const (
	TokenPaymentMethodTypeCard                  TokenPaymentMethodType = "card"
	TokenPaymentMethodTypePaypalExpressCheckout TokenPaymentMethodType = "paypal_express_checkout"
	TokenPaymentMethodTypeAmazonPayments        TokenPaymentMethodType = "amazon_payments"
	TokenPaymentMethodTypeDirectDebit           TokenPaymentMethodType = "direct_debit"
	TokenPaymentMethodTypeGeneric               TokenPaymentMethodType = "generic"
	TokenPaymentMethodTypeAlipay                TokenPaymentMethodType = "alipay"
	TokenPaymentMethodTypeUnionpay              TokenPaymentMethodType = "unionpay"
	TokenPaymentMethodTypeApplePay              TokenPaymentMethodType = "apple_pay"
	TokenPaymentMethodTypeWechatPay             TokenPaymentMethodType = "wechat_pay"
	TokenPaymentMethodTypeIdeal                 TokenPaymentMethodType = "ideal"
	TokenPaymentMethodTypeGooglePay             TokenPaymentMethodType = "google_pay"
	TokenPaymentMethodTypeSofort                TokenPaymentMethodType = "sofort"
	TokenPaymentMethodTypeBancontact            TokenPaymentMethodType = "bancontact"
	TokenPaymentMethodTypeGiropay               TokenPaymentMethodType = "giropay"
	TokenPaymentMethodTypeDotpay                TokenPaymentMethodType = "dotpay"
	TokenPaymentMethodTypeUpi                   TokenPaymentMethodType = "upi"
	TokenPaymentMethodTypeNetbankingEmandates   TokenPaymentMethodType = "netbanking_emandates"
	TokenPaymentMethodTypeVenmo                 TokenPaymentMethodType = "venmo"
	TokenPaymentMethodTypePayTo                 TokenPaymentMethodType = "pay_to"
	TokenPaymentMethodTypeFasterPayments        TokenPaymentMethodType = "faster_payments"
	TokenPaymentMethodTypeSepaInstantTransfer   TokenPaymentMethodType = "sepa_instant_transfer"
	TokenPaymentMethodTypeAutomatedBankTransfer TokenPaymentMethodType = "automated_bank_transfer"
	TokenPaymentMethodTypeKlarnaPayNow          TokenPaymentMethodType = "klarna_pay_now"
	TokenPaymentMethodTypeOnlineBankingPoland   TokenPaymentMethodType = "online_banking_poland"
	TokenPaymentMethodTypePayconiqByBancontact  TokenPaymentMethodType = "payconiq_by_bancontact"
)

type TokenStatus string

const (
	TokenStatusNew      TokenStatus = "new"
	TokenStatusExpired  TokenStatus = "expired"
	TokenStatusConsumed TokenStatus = "consumed"
)

type TokenVault string

const (
	TokenVaultSpreedly TokenVault = "spreedly"
	TokenVaultGateway  TokenVault = "gateway"
)

// just struct
type Token struct {
	Id                string                 `json:"id"`
	Gateway           TokenGateway           `json:"gateway"`
	GatewayAccountId  string                 `json:"gateway_account_id"`
	PaymentMethodType TokenPaymentMethodType `json:"payment_method_type"`
	Status            TokenStatus            `json:"status"`
	IdAtVault         string                 `json:"id_at_vault"`
	Vault             TokenVault             `json:"vault"`
	IpAddress         string                 `json:"ip_address"`
	ResourceVersion   int64                  `json:"resource_version"`
	UpdatedAt         int64                  `json:"updated_at"`
	CreatedAt         int64                  `json:"created_at"`
	ExpiredAt         int64                  `json:"expired_at"`
	Object            string                 `json:"object"`
}

// sub resources
// operations
// input params
