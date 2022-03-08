package enum

type PaymentMethodType string

const (
	PaymentMethodTypeCard                  PaymentMethodType = "card"
	PaymentMethodTypePaypalExpressCheckout PaymentMethodType = "paypal_express_checkout"
	PaymentMethodTypeAmazonPayments        PaymentMethodType = "amazon_payments"
	PaymentMethodTypeDirectDebit           PaymentMethodType = "direct_debit"
	PaymentMethodTypeGeneric               PaymentMethodType = "generic"
	PaymentMethodTypeAlipay                PaymentMethodType = "alipay"
	PaymentMethodTypeUnionpay              PaymentMethodType = "unionpay"
	PaymentMethodTypeApplePay              PaymentMethodType = "apple_pay"
	PaymentMethodTypeWechatPay             PaymentMethodType = "wechat_pay"
	PaymentMethodTypeIdeal                 PaymentMethodType = "ideal"
	PaymentMethodTypeGooglePay             PaymentMethodType = "google_pay"
	PaymentMethodTypeSofort                PaymentMethodType = "sofort"
	PaymentMethodTypeBancontact            PaymentMethodType = "bancontact"
	PaymentMethodTypeGiropay               PaymentMethodType = "giropay"
	PaymentMethodTypeDotpay                PaymentMethodType = "dotpay"
	PaymentMethodTypeUpi                   PaymentMethodType = "upi"
	PaymentMethodTypeNetbankingEmandates   PaymentMethodType = "netbanking_emandates"
)
