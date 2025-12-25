package chargebee

type PromotionalCreditType string

const (
	PromotionalCreditTypeIncrement PromotionalCreditType = "increment"
	PromotionalCreditTypeDecrement PromotionalCreditType = "decrement"
)

type PromotionalCreditCreditType string

const (
	PromotionalCreditCreditTypeLoyaltyCredits  PromotionalCreditCreditType = "loyalty_credits"
	PromotionalCreditCreditTypeReferralRewards PromotionalCreditCreditType = "referral_rewards"
	PromotionalCreditCreditTypeGeneral         PromotionalCreditCreditType = "general"
)

// just struct
type PromotionalCredit struct {
	Id               string                      `json:"id"`
	CustomerId       string                      `json:"customer_id"`
	Type             PromotionalCreditType       `json:"type"`
	AmountInDecimal  string                      `json:"amount_in_decimal"`
	Amount           int64                       `json:"amount"`
	CurrencyCode     string                      `json:"currency_code"`
	Description      string                      `json:"description"`
	CreditType       PromotionalCreditCreditType `json:"credit_type"`
	Reference        string                      `json:"reference"`
	ClosingBalance   int64                       `json:"closing_balance"`
	DoneBy           string                      `json:"done_by"`
	CreatedAt        int64                       `json:"created_at"`
	BusinessEntityId string                      `json:"business_entity_id"`
	Object           string                      `json:"object"`
}

// sub resources
// operations
// input params
type PromotionalCreditAddRequest struct {
	CustomerId      string                      `json:"customer_id"`
	Amount          *int64                      `json:"amount,omitempty"`
	AmountInDecimal string                      `json:"amount_in_decimal,omitempty"`
	CurrencyCode    string                      `json:"currency_code,omitempty"`
	Description     string                      `json:"description"`
	CreditType      PromotionalCreditCreditType `json:"credit_type,omitempty"`
	Reference       string                      `json:"reference,omitempty"`
	apiRequest      `json:"-" form:"-"`
}

func (r *PromotionalCreditAddRequest) payload() any { return r }

type PromotionalCreditDeductRequest struct {
	CustomerId      string                      `json:"customer_id"`
	Amount          *int64                      `json:"amount,omitempty"`
	AmountInDecimal string                      `json:"amount_in_decimal,omitempty"`
	CurrencyCode    string                      `json:"currency_code,omitempty"`
	Description     string                      `json:"description"`
	CreditType      PromotionalCreditCreditType `json:"credit_type,omitempty"`
	Reference       string                      `json:"reference,omitempty"`
	apiRequest      `json:"-" form:"-"`
}

func (r *PromotionalCreditDeductRequest) payload() any { return r }

type PromotionalCreditSetRequest struct {
	CustomerId      string                      `json:"customer_id"`
	Amount          *int64                      `json:"amount,omitempty"`
	AmountInDecimal string                      `json:"amount_in_decimal,omitempty"`
	CurrencyCode    string                      `json:"currency_code,omitempty"`
	Description     string                      `json:"description"`
	CreditType      PromotionalCreditCreditType `json:"credit_type,omitempty"`
	Reference       string                      `json:"reference,omitempty"`
	apiRequest      `json:"-" form:"-"`
}

func (r *PromotionalCreditSetRequest) payload() any { return r }

type PromotionalCreditListRequest struct {
	Limit      *int32           `json:"limit,omitempty"`
	Offset     string           `json:"offset,omitempty"`
	Id         *StringFilter    `json:"id,omitempty"`
	CreatedAt  *TimestampFilter `json:"created_at,omitempty"`
	Type       *EnumFilter      `json:"type,omitempty"`
	CustomerId *StringFilter    `json:"customer_id,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *PromotionalCreditListRequest) payload() any { return r }

// operation response
type PromotionalCreditAddResponse struct {
	Customer          Customer           `json:"customer,omitempty"`
	PromotionalCredit *PromotionalCredit `json:"promotional_credit,omitempty"`
	apiResponse
}

// operation response
type PromotionalCreditDeductResponse struct {
	Customer          Customer           `json:"customer,omitempty"`
	PromotionalCredit *PromotionalCredit `json:"promotional_credit,omitempty"`
	apiResponse
}

// operation response
type PromotionalCreditSetResponse struct {
	Customer          Customer           `json:"customer,omitempty"`
	PromotionalCredit *PromotionalCredit `json:"promotional_credit,omitempty"`
	apiResponse
}

// operation sub response
type PromotionalCreditListPromotionalCreditResponse struct {
	PromotionalCredit *PromotionalCredit `json:"promotional_credit,omitempty"`
}

// operation response
type PromotionalCreditListResponse struct {
	List       []*PromotionalCreditListPromotionalCreditResponse `json:"list,omitempty"`
	NextOffset string                                            `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type PromotionalCreditRetrieveResponse struct {
	PromotionalCredit *PromotionalCredit `json:"promotional_credit,omitempty"`
	apiResponse
}
