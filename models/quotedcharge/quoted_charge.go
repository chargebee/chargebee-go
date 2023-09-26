package quotedcharge

import (
	"github.com/chargebee/chargebee-go/enum"
)

type QuotedCharge struct {
	Charges      []*Charge      `json:"charges"`
	Addons       []*Addon       `json:"addons"`
	InvoiceItems []*InvoiceItem `json:"invoice_items"`
	ItemTiers    []*ItemTier    `json:"item_tiers"`
	Coupons      []*Coupon      `json:"coupons"`
	Object       string         `json:"object"`
}
type Charge struct {
	Amount                 int32                `json:"amount"`
	AmountInDecimal        string               `json:"amount_in_decimal"`
	Description            string               `json:"description"`
	ServicePeriodInDays    int32                `json:"service_period_in_days"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type"`
	AvalaraTransactionType int32                `json:"avalara_transaction_type"`
	AvalaraServiceType     int32                `json:"avalara_service_type"`
	Object                 string               `json:"object"`
}
type Addon struct {
	Id                 string             `json:"id"`
	Quantity           int32              `json:"quantity"`
	UnitPrice          int32              `json:"unit_price"`
	QuantityInDecimal  string             `json:"quantity_in_decimal"`
	UnitPriceInDecimal string             `json:"unit_price_in_decimal"`
	ProrationType      enum.ProrationType `json:"proration_type"`
	ServicePeriod      int32              `json:"service_period"`
	Object             string             `json:"object"`
}
type InvoiceItem struct {
	ItemPriceId        string `json:"item_price_id"`
	Quantity           int32  `json:"quantity"`
	QuantityInDecimal  string `json:"quantity_in_decimal"`
	UnitPrice          int32  `json:"unit_price"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal"`
	ServicePeriodDays  int32  `json:"service_period_days"`
	Object             string `json:"object"`
}
type ItemTier struct {
	ItemPriceId           string `json:"item_price_id"`
	StartingUnit          int32  `json:"starting_unit"`
	EndingUnit            int32  `json:"ending_unit"`
	Price                 int32  `json:"price"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal"`
	PriceInDecimal        string `json:"price_in_decimal"`
	Index                 int32  `json:"index"`
	Object                string `json:"object"`
}
type Coupon struct {
	CouponId string `json:"coupon_id"`
	Object   string `json:"object"`
}
