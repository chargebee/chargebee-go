package chargebee

type Type string

const (
	TypePayment Type = "payment"
	TypeRefund  Type = "refund"
)

type PaymentMethod string

const (
	PaymentMethodCash         PaymentMethod = "cash"
	PaymentMethodCheck        PaymentMethod = "check"
	PaymentMethodChargeback   PaymentMethod = "chargeback"
	PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
	PaymentMethodOther        PaymentMethod = "other"
)

type TaxWithheld struct {
	Id              string        `json:"id"`
	User            string        `json:"user"`
	ReferenceNumber string        `json:"reference_number"`
	Description     string        `json:"description"`
	Type            Type          `json:"type"`
	PaymentMethod   PaymentMethod `json:"payment_method"`
	Date            int64         `json:"date"`
	CurrencyCode    string        `json:"currency_code"`
	Amount          int64         `json:"amount"`
	ResourceVersion int64         `json:"resource_version"`
	UpdatedAt       int64         `json:"updated_at"`
	ExchangeRate    float64       `json:"exchange_rate"`
	Object          string        `json:"object"`
}
