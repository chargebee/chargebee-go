package enum

type Gateway string

const (
	GatewayChargebee             Gateway = "chargebee"
	GatewayChargebeePayments     Gateway = "chargebee_payments"
	GatewayStripe                Gateway = "stripe"
	GatewayWepay                 Gateway = "wepay"
	GatewayBraintree             Gateway = "braintree"
	GatewayAuthorizeNet          Gateway = "authorize_net"
	GatewayPaypalPro             Gateway = "paypal_pro"
	GatewayPin                   Gateway = "pin"
	GatewayEway                  Gateway = "eway"
	GatewayEwayRapid             Gateway = "eway_rapid"
	GatewayWorldpay              Gateway = "worldpay"
	GatewayBalancedPayments      Gateway = "balanced_payments"
	GatewayBeanstream            Gateway = "beanstream"
	GatewayBluepay               Gateway = "bluepay"
	GatewayElavon                Gateway = "elavon"
	GatewayFirstDataGlobal       Gateway = "first_data_global"
	GatewayHdfc                  Gateway = "hdfc"
	GatewayMigs                  Gateway = "migs"
	GatewayNmi                   Gateway = "nmi"
	GatewayOgone                 Gateway = "ogone"
	GatewayPaymill               Gateway = "paymill"
	GatewayPaypalPayflowPro      Gateway = "paypal_payflow_pro"
	GatewaySagePay               Gateway = "sage_pay"
	GatewayTco                   Gateway = "tco"
	GatewayWirecard              Gateway = "wirecard"
	GatewayAmazonPayments        Gateway = "amazon_payments"
	GatewayPaypalExpressCheckout Gateway = "paypal_express_checkout"
	GatewayGocardless            Gateway = "gocardless"
	GatewayAdyen                 Gateway = "adyen"
	GatewayOrbital               Gateway = "orbital"
	GatewayMonerisUs             Gateway = "moneris_us"
	GatewayMoneris               Gateway = "moneris"
	GatewayBluesnap              Gateway = "bluesnap"
	GatewayCybersource           Gateway = "cybersource"
	GatewayVantiv                Gateway = "vantiv"
	GatewayCheckoutCom           Gateway = "checkout_com"
	GatewayPaypal                Gateway = "paypal"
	GatewayIngenicoDirect        Gateway = "ingenico_direct"
	GatewayExact                 Gateway = "exact"
	GatewayMollie                Gateway = "mollie"
	GatewayQuickbooks            Gateway = "quickbooks"
	GatewayRazorpay              Gateway = "razorpay"
	GatewayGlobalPayments        Gateway = "global_payments"
	GatewayBankOfAmerica         Gateway = "bank_of_america"
	GatewayNotApplicable         Gateway = "not_applicable"
)
