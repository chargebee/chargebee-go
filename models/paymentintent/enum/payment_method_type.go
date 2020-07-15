package enum

type PaymentMethodType string

const (
	PaymentMethodTypeCard       PaymentMethodType = "card"
	PaymentMethodTypeIdeal      PaymentMethodType = "ideal"
	PaymentMethodTypeSofort     PaymentMethodType = "sofort"
	PaymentMethodTypeBancontact PaymentMethodType = "bancontact"
)
