package chargebee

type ThirdPartyPaymentMethodType string

const (
	ThirdPartyPaymentMethodTypeCard                  ThirdPartyPaymentMethodType = "card"
	ThirdPartyPaymentMethodTypePaypalExpressCheckout ThirdPartyPaymentMethodType = "paypal_express_checkout"
	ThirdPartyPaymentMethodTypeAmazonPayments        ThirdPartyPaymentMethodType = "amazon_payments"
	ThirdPartyPaymentMethodTypeDirectDebit           ThirdPartyPaymentMethodType = "direct_debit"
	ThirdPartyPaymentMethodTypeGeneric               ThirdPartyPaymentMethodType = "generic"
	ThirdPartyPaymentMethodTypeAlipay                ThirdPartyPaymentMethodType = "alipay"
	ThirdPartyPaymentMethodTypeUnionpay              ThirdPartyPaymentMethodType = "unionpay"
	ThirdPartyPaymentMethodTypeApplePay              ThirdPartyPaymentMethodType = "apple_pay"
	ThirdPartyPaymentMethodTypeWechatPay             ThirdPartyPaymentMethodType = "wechat_pay"
	ThirdPartyPaymentMethodTypeIdeal                 ThirdPartyPaymentMethodType = "ideal"
	ThirdPartyPaymentMethodTypeGooglePay             ThirdPartyPaymentMethodType = "google_pay"
	ThirdPartyPaymentMethodTypeSofort                ThirdPartyPaymentMethodType = "sofort"
	ThirdPartyPaymentMethodTypeBancontact            ThirdPartyPaymentMethodType = "bancontact"
	ThirdPartyPaymentMethodTypeGiropay               ThirdPartyPaymentMethodType = "giropay"
	ThirdPartyPaymentMethodTypeDotpay                ThirdPartyPaymentMethodType = "dotpay"
	ThirdPartyPaymentMethodTypeUpi                   ThirdPartyPaymentMethodType = "upi"
	ThirdPartyPaymentMethodTypeNetbankingEmandates   ThirdPartyPaymentMethodType = "netbanking_emandates"
	ThirdPartyPaymentMethodTypeVenmo                 ThirdPartyPaymentMethodType = "venmo"
	ThirdPartyPaymentMethodTypePayTo                 ThirdPartyPaymentMethodType = "pay_to"
	ThirdPartyPaymentMethodTypeFasterPayments        ThirdPartyPaymentMethodType = "faster_payments"
	ThirdPartyPaymentMethodTypeSepaInstantTransfer   ThirdPartyPaymentMethodType = "sepa_instant_transfer"
	ThirdPartyPaymentMethodTypeAutomatedBankTransfer ThirdPartyPaymentMethodType = "automated_bank_transfer"
	ThirdPartyPaymentMethodTypeKlarnaPayNow          ThirdPartyPaymentMethodType = "klarna_pay_now"
	ThirdPartyPaymentMethodTypeOnlineBankingPoland   ThirdPartyPaymentMethodType = "online_banking_poland"
	ThirdPartyPaymentMethodTypePayconiqByBancontact  ThirdPartyPaymentMethodType = "payconiq_by_bancontact"
)

type ThirdPartyPaymentMethodGateway string

const (
	ThirdPartyPaymentMethodGatewayChargebee             ThirdPartyPaymentMethodGateway = "chargebee"
	ThirdPartyPaymentMethodGatewayChargebeePayments     ThirdPartyPaymentMethodGateway = "chargebee_payments"
	ThirdPartyPaymentMethodGatewayAdyen                 ThirdPartyPaymentMethodGateway = "adyen"
	ThirdPartyPaymentMethodGatewayStripe                ThirdPartyPaymentMethodGateway = "stripe"
	ThirdPartyPaymentMethodGatewayWepay                 ThirdPartyPaymentMethodGateway = "wepay"
	ThirdPartyPaymentMethodGatewayBraintree             ThirdPartyPaymentMethodGateway = "braintree"
	ThirdPartyPaymentMethodGatewayAuthorizeNet          ThirdPartyPaymentMethodGateway = "authorize_net"
	ThirdPartyPaymentMethodGatewayPaypalPro             ThirdPartyPaymentMethodGateway = "paypal_pro"
	ThirdPartyPaymentMethodGatewayPin                   ThirdPartyPaymentMethodGateway = "pin"
	ThirdPartyPaymentMethodGatewayEway                  ThirdPartyPaymentMethodGateway = "eway"
	ThirdPartyPaymentMethodGatewayEwayRapid             ThirdPartyPaymentMethodGateway = "eway_rapid"
	ThirdPartyPaymentMethodGatewayWorldpay              ThirdPartyPaymentMethodGateway = "worldpay"
	ThirdPartyPaymentMethodGatewayBalancedPayments      ThirdPartyPaymentMethodGateway = "balanced_payments"
	ThirdPartyPaymentMethodGatewayBeanstream            ThirdPartyPaymentMethodGateway = "beanstream"
	ThirdPartyPaymentMethodGatewayBluepay               ThirdPartyPaymentMethodGateway = "bluepay"
	ThirdPartyPaymentMethodGatewayElavon                ThirdPartyPaymentMethodGateway = "elavon"
	ThirdPartyPaymentMethodGatewayFirstDataGlobal       ThirdPartyPaymentMethodGateway = "first_data_global"
	ThirdPartyPaymentMethodGatewayHdfc                  ThirdPartyPaymentMethodGateway = "hdfc"
	ThirdPartyPaymentMethodGatewayMigs                  ThirdPartyPaymentMethodGateway = "migs"
	ThirdPartyPaymentMethodGatewayNmi                   ThirdPartyPaymentMethodGateway = "nmi"
	ThirdPartyPaymentMethodGatewayOgone                 ThirdPartyPaymentMethodGateway = "ogone"
	ThirdPartyPaymentMethodGatewayPaymill               ThirdPartyPaymentMethodGateway = "paymill"
	ThirdPartyPaymentMethodGatewayPaypalPayflowPro      ThirdPartyPaymentMethodGateway = "paypal_payflow_pro"
	ThirdPartyPaymentMethodGatewaySagePay               ThirdPartyPaymentMethodGateway = "sage_pay"
	ThirdPartyPaymentMethodGatewayTco                   ThirdPartyPaymentMethodGateway = "tco"
	ThirdPartyPaymentMethodGatewayWirecard              ThirdPartyPaymentMethodGateway = "wirecard"
	ThirdPartyPaymentMethodGatewayAmazonPayments        ThirdPartyPaymentMethodGateway = "amazon_payments"
	ThirdPartyPaymentMethodGatewayPaypalExpressCheckout ThirdPartyPaymentMethodGateway = "paypal_express_checkout"
	ThirdPartyPaymentMethodGatewayGocardless            ThirdPartyPaymentMethodGateway = "gocardless"
	ThirdPartyPaymentMethodGatewayOrbital               ThirdPartyPaymentMethodGateway = "orbital"
	ThirdPartyPaymentMethodGatewayMonerisUs             ThirdPartyPaymentMethodGateway = "moneris_us"
	ThirdPartyPaymentMethodGatewayMoneris               ThirdPartyPaymentMethodGateway = "moneris"
	ThirdPartyPaymentMethodGatewayBluesnap              ThirdPartyPaymentMethodGateway = "bluesnap"
	ThirdPartyPaymentMethodGatewayCybersource           ThirdPartyPaymentMethodGateway = "cybersource"
	ThirdPartyPaymentMethodGatewayVantiv                ThirdPartyPaymentMethodGateway = "vantiv"
	ThirdPartyPaymentMethodGatewayCheckoutCom           ThirdPartyPaymentMethodGateway = "checkout_com"
	ThirdPartyPaymentMethodGatewayPaypal                ThirdPartyPaymentMethodGateway = "paypal"
	ThirdPartyPaymentMethodGatewayIngenicoDirect        ThirdPartyPaymentMethodGateway = "ingenico_direct"
	ThirdPartyPaymentMethodGatewayExact                 ThirdPartyPaymentMethodGateway = "exact"
	ThirdPartyPaymentMethodGatewayMollie                ThirdPartyPaymentMethodGateway = "mollie"
	ThirdPartyPaymentMethodGatewayQuickbooks            ThirdPartyPaymentMethodGateway = "quickbooks"
	ThirdPartyPaymentMethodGatewayRazorpay              ThirdPartyPaymentMethodGateway = "razorpay"
	ThirdPartyPaymentMethodGatewayGlobalPayments        ThirdPartyPaymentMethodGateway = "global_payments"
	ThirdPartyPaymentMethodGatewayBankOfAmerica         ThirdPartyPaymentMethodGateway = "bank_of_america"
	ThirdPartyPaymentMethodGatewayEcentric              ThirdPartyPaymentMethodGateway = "ecentric"
	ThirdPartyPaymentMethodGatewayMetricsGlobal         ThirdPartyPaymentMethodGateway = "metrics_global"
	ThirdPartyPaymentMethodGatewayWindcave              ThirdPartyPaymentMethodGateway = "windcave"
	ThirdPartyPaymentMethodGatewayPayCom                ThirdPartyPaymentMethodGateway = "pay_com"
	ThirdPartyPaymentMethodGatewayEbanx                 ThirdPartyPaymentMethodGateway = "ebanx"
	ThirdPartyPaymentMethodGatewayDlocal                ThirdPartyPaymentMethodGateway = "dlocal"
	ThirdPartyPaymentMethodGatewayNuvei                 ThirdPartyPaymentMethodGateway = "nuvei"
	ThirdPartyPaymentMethodGatewaySolidgate             ThirdPartyPaymentMethodGateway = "solidgate"
	ThirdPartyPaymentMethodGatewayPaystack              ThirdPartyPaymentMethodGateway = "paystack"
	ThirdPartyPaymentMethodGatewayJpMorgan              ThirdPartyPaymentMethodGateway = "jp_morgan"
	ThirdPartyPaymentMethodGatewayDeutscheBank          ThirdPartyPaymentMethodGateway = "deutsche_bank"
	ThirdPartyPaymentMethodGatewayEzidebit              ThirdPartyPaymentMethodGateway = "ezidebit"
	ThirdPartyPaymentMethodGatewayNotApplicable         ThirdPartyPaymentMethodGateway = "not_applicable"
)

// just struct
type ThirdPartyPaymentMethod struct {
	Type             ThirdPartyPaymentMethodType    `json:"type"`
	Gateway          ThirdPartyPaymentMethodGateway `json:"gateway"`
	GatewayAccountId string                         `json:"gateway_account_id"`
	ReferenceId      string                         `json:"reference_id"`
	Object           string                         `json:"object"`
}

// sub resources
// operations
// input params
