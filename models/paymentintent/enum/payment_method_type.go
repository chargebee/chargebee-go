package enum

type PaymentMethodType string

const (
	PaymentMethodTypeCard                PaymentMethodType = "card"
	PaymentMethodTypeIdeal               PaymentMethodType = "ideal"
	PaymentMethodTypeSofort              PaymentMethodType = "sofort"
	PaymentMethodTypeBancontact          PaymentMethodType = "bancontact"
	PaymentMethodTypeGooglePay           PaymentMethodType = "google_pay"
	PaymentMethodTypeDotpay              PaymentMethodType = "dotpay"
	PaymentMethodTypeGiropay             PaymentMethodType = "giropay"
	PaymentMethodTypeApplePay            PaymentMethodType = "apple_pay"
	PaymentMethodTypeUpi                 PaymentMethodType = "upi"
	PaymentMethodTypeNetbankingEmandates PaymentMethodType = "netbanking_emandates"
)
