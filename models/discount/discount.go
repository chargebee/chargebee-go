package discount

import (
	"github.com/chargebee/chargebee-go/enum"
	discountEnum "github.com/chargebee/chargebee-go/models/discount/enum"
)

type Discount struct {
	Id            string            `json:"id"`
	InvoiceName   string            `json:"invoice_name"`
	Type          discountEnum.Type `json:"type"`
	Percentage    float64           `json:"percentage"`
	Amount        int32             `json:"amount"`
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
