package enum

type PaymentInitiator string

const (
	PaymentInitiatorCustomer PaymentInitiator = "customer"
	PaymentInitiatorMerchant PaymentInitiator = "merchant"
)
