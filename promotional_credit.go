package chargebee

type Type string

const (
	TypeIncrement Type = "increment"
	TypeDecrement Type = "decrement"
)

type PromotionalCredit struct {
	Id               string          `json:"id"`
	CustomerId       string          `json:"customer_id"`
	Type             Type            `json:"type"`
	AmountInDecimal  string          `json:"amount_in_decimal"`
	Amount           int64           `json:"amount"`
	CurrencyCode     string          `json:"currency_code"`
	Description      string          `json:"description"`
	CreditType       enum.CreditType `json:"credit_type"`
	Reference        string          `json:"reference"`
	ClosingBalance   int64           `json:"closing_balance"`
	DoneBy           string          `json:"done_by"`
	CreatedAt        int64           `json:"created_at"`
	BusinessEntityId string          `json:"business_entity_id"`
	Object           string          `json:"object"`
}
type AddRequest struct {
	CustomerId      string          `json:"customer_id"`
	Amount          *int64          `json:"amount,omitempty"`
	AmountInDecimal string          `json:"amount_in_decimal,omitempty"`
	CurrencyCode    string          `json:"currency_code,omitempty"`
	Description     string          `json:"description"`
	CreditType      enum.CreditType `json:"credit_type,omitempty"`
	Reference       string          `json:"reference,omitempty"`
}
type DeductRequest struct {
	CustomerId      string          `json:"customer_id"`
	Amount          *int64          `json:"amount,omitempty"`
	AmountInDecimal string          `json:"amount_in_decimal,omitempty"`
	CurrencyCode    string          `json:"currency_code,omitempty"`
	Description     string          `json:"description"`
	CreditType      enum.CreditType `json:"credit_type,omitempty"`
	Reference       string          `json:"reference,omitempty"`
}
type SetRequest struct {
	CustomerId      string          `json:"customer_id"`
	Amount          *int64          `json:"amount,omitempty"`
	AmountInDecimal string          `json:"amount_in_decimal,omitempty"`
	CurrencyCode    string          `json:"currency_code,omitempty"`
	Description     string          `json:"description"`
	CreditType      enum.CreditType `json:"credit_type,omitempty"`
	Reference       string          `json:"reference,omitempty"`
}
type ListRequest struct {
	Limit      *int32                  `json:"limit,omitempty"`
	Offset     string                  `json:"offset,omitempty"`
	Id         *filter.StringFilter    `json:"id,omitempty"`
	CreatedAt  *filter.TimestampFilter `json:"created_at,omitempty"`
	Type       *filter.EnumFilter      `json:"type,omitempty"`
	CustomerId *filter.StringFilter    `json:"customer_id,omitempty"`
}

type AddResponse struct {
	Customer          *customer.Customer `json:"customer,omitempty"`
	PromotionalCredit *PromotionalCredit `json:"promotional_credit,omitempty"`
}

type DeductResponse struct {
	Customer          *customer.Customer `json:"customer,omitempty"`
	PromotionalCredit *PromotionalCredit `json:"promotional_credit,omitempty"`
}

type SetResponse struct {
	Customer          *customer.Customer `json:"customer,omitempty"`
	PromotionalCredit *PromotionalCredit `json:"promotional_credit,omitempty"`
}

type ListPromotionalCreditResponse struct {
	PromotionalCredit *PromotionalCredit `json:"promotional_credit,omitempty"`
}

type ListResponse struct {
	List       []*ListPromotionalCreditResponse `json:"list,omitempty"`
	NextOffset string                           `json:"next_offset,omitempty"`
}

type RetrieveResponse struct {
	PromotionalCredit *PromotionalCredit `json:"promotional_credit,omitempty"`
}
