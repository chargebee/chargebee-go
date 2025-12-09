package chargebee

type Type string

const (
	TypeFixedAmount   Type = "fixed_amount"
	TypePercentage    Type = "percentage"
	TypeOfferQuantity Type = "offer_quantity"
)

type Discount struct {
	Id            string            `json:"id"`
	InvoiceName   string            `json:"invoice_name"`
	Type          Type              `json:"type"`
	Percentage    float64           `json:"percentage"`
	Amount        int64             `json:"amount"`
	Quantity      int32             `json:"quantity"`
	CurrencyCode  string            `json:"currency_code"`
	DurationType  enum.DurationType `json:"duration_type"`
	Period        int32             `json:"period"`
	PeriodUnit    enum.PeriodUnit   `json:"period_unit"`
	IncludedInMrr bool              `json:"included_in_mrr"`
	ApplyOn       enum.ApplyOn      `json:"apply_on"`
	ItemPriceId   string            `json:"item_price_id"`
	CreatedAt     int64             `json:"created_at"`
	ApplyTill     int64             `json:"apply_till"`
	AppliedCount  int32             `json:"applied_count"`
	CouponId      string            `json:"coupon_id"`
	Index         int32             `json:"index"`
	Object        string            `json:"object"`
}
