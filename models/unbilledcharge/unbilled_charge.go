package unbilledcharge

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	unbilledChargeEnum "github.com/chargebee/chargebee-go/models/unbilledcharge/enum"
)

type UnbilledCharge struct {
	Id             string                        `json:"id"`
	CustomerId     string                        `json:"customer_id"`
	SubscriptionId string                        `json:"subscription_id"`
	DateFrom       int64                         `json:"date_from"`
	DateTo         int64                         `json:"date_to"`
	UnitAmount     int32                         `json:"unit_amount"`
	PricingModel   enum.PricingModel             `json:"pricing_model"`
	Quantity       int32                         `json:"quantity"`
	Amount         int32                         `json:"amount"`
	CurrencyCode   string                        `json:"currency_code"`
	DiscountAmount int32                         `json:"discount_amount"`
	Description    string                        `json:"description"`
	EntityType     unbilledChargeEnum.EntityType `json:"entity_type"`
	EntityId       string                        `json:"entity_id"`
	IsVoided       bool                          `json:"is_voided"`
	VoidedAt       int64                         `json:"voided_at"`
	Tiers          []*Tier                       `json:"tiers"`
	Deleted        bool                          `json:"deleted"`
	Object         string                        `json:"object"`
}
type Tier struct {
	StartingUnit int32  `json:"starting_unit"`
	EndingUnit   int32  `json:"ending_unit"`
	QuantityUsed int32  `json:"quantity_used"`
	UnitAmount   int32  `json:"unit_amount"`
	Object       string `json:"object"`
}
type InvoiceUnbilledChargesRequestParams struct {
	SubscriptionId string `json:"subscription_id,omitempty"`
	CustomerId     string `json:"customer_id,omitempty"`
}
type ListRequestParams struct {
	Limit          *int32               `json:"limit,omitempty"`
	Offset         string               `json:"offset,omitempty"`
	IncludeDeleted *bool                `json:"include_deleted,omitempty"`
	SubscriptionId *filter.StringFilter `json:"subscription_id,omitempty"`
	CustomerId     *filter.StringFilter `json:"customer_id,omitempty"`
}
type InvoiceNowEstimateRequestParams struct {
	SubscriptionId string `json:"subscription_id,omitempty"`
	CustomerId     string `json:"customer_id,omitempty"`
}
