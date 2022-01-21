package enum

type PaymentMethod string

const (
	PaymentMethodCash         PaymentMethod = "cash"
	PaymentMethodCheck        PaymentMethod = "check"
	PaymentMethodChargeback   PaymentMethod = "chargeback"
	PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
	PaymentMethodOther        PaymentMethod = "other"
)
