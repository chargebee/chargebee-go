package enum

type PaymentMethodSavePolicy string

const (
	PaymentMethodSavePolicyAlways PaymentMethodSavePolicy = "always"
	PaymentMethodSavePolicyAsk    PaymentMethodSavePolicy = "ask"
	PaymentMethodSavePolicyNever  PaymentMethodSavePolicy = "never"
)
