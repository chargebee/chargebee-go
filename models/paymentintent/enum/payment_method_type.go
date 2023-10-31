package enum

type PaymentMethodType string

const (
	PaymentMethodTypeCard                  PaymentMethodType = "card"
	PaymentMethodTypeIdeal                 PaymentMethodType = "ideal"
	PaymentMethodTypeSofort                PaymentMethodType = "sofort"
	PaymentMethodTypeBancontact            PaymentMethodType = "bancontact"
	PaymentMethodTypeGooglePay             PaymentMethodType = "google_pay"
	PaymentMethodTypeDotpay                PaymentMethodType = "dotpay"
	PaymentMethodTypeGiropay               PaymentMethodType = "giropay"
	PaymentMethodTypeApplePay              PaymentMethodType = "apple_pay"
	PaymentMethodTypeUpi                   PaymentMethodType = "upi"
	PaymentMethodTypeNetbankingEmandates   PaymentMethodType = "netbanking_emandates"
	PaymentMethodTypePaypalExpressCheckout PaymentMethodType = "paypal_express_checkout"
	PaymentMethodTypeDirectDebit           PaymentMethodType = "direct_debit"
	PaymentMethodTypeBoleto                PaymentMethodType = "boleto"
	PaymentMethodTypeVenmo                 PaymentMethodType = "venmo"
	PaymentMethodTypeAmazonPayments        PaymentMethodType = "amazon_payments"
	PaymentMethodTypePayTo                 PaymentMethodType = "pay_to"
	PaymentMethodTypeFasterPayments        PaymentMethodType = "faster_payments"
	PaymentMethodTypeSepaInstantTransfer   PaymentMethodType = "sepa_instant_transfer"
)
