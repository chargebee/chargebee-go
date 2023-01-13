package unbilledcharge

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	unbilledChargeEnum "github.com/chargebee/chargebee-go/models/unbilledcharge/enum"
)

type UnbilledCharge struct {
	Id                  string                        `json:"id"`
	CustomerId          string                        `json:"customer_id"`
	SubscriptionId      string                        `json:"subscription_id"`
	DateFrom            int64                         `json:"date_from"`
	DateTo              int64                         `json:"date_to"`
	UnitAmount          int32                         `json:"unit_amount"`
	PricingModel        enum.PricingModel             `json:"pricing_model"`
	Quantity            int32                         `json:"quantity"`
	Amount              int32                         `json:"amount"`
	CurrencyCode        string                        `json:"currency_code"`
	DiscountAmount      int32                         `json:"discount_amount"`
	Description         string                        `json:"description"`
	EntityType          unbilledChargeEnum.EntityType `json:"entity_type"`
	EntityId            string                        `json:"entity_id"`
	IsVoided            bool                          `json:"is_voided"`
	VoidedAt            int64                         `json:"voided_at"`
	UnitAmountInDecimal string                        `json:"unit_amount_in_decimal"`
	QuantityInDecimal   string                        `json:"quantity_in_decimal"`
	AmountInDecimal     string                        `json:"amount_in_decimal"`
	UpdatedAt           int64                         `json:"updated_at"`
	Tiers               []*Tier                       `json:"tiers"`
	IsAdvanceCharge     bool                          `json:"is_advance_charge"`
	Deleted             bool                          `json:"deleted"`
	Object              string                        `json:"object"`
}
type Tier struct {
	StartingUnit          int32  `json:"starting_unit"`
	EndingUnit            int32  `json:"ending_unit"`
	QuantityUsed          int32  `json:"quantity_used"`
	UnitAmount            int32  `json:"unit_amount"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal"`
	QuantityUsedInDecimal string `json:"quantity_used_in_decimal"`
	UnitAmountInDecimal   string `json:"unit_amount_in_decimal"`
	Object                string `json:"object"`
}
type CreateUnbilledChargeRequestParams struct {
	SubscriptionId string                              `json:"subscription_id"`
	CurrencyCode   string                              `json:"currency_code,omitempty"`
	Addons         []*CreateUnbilledChargeAddonParams  `json:"addons,omitempty"`
	Charges        []*CreateUnbilledChargeChargeParams `json:"charges,omitempty"`
}
type CreateUnbilledChargeAddonParams struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type CreateUnbilledChargeChargeParams struct {
	Amount                 *int32               `json:"amount,omitempty"`
	AmountInDecimal        string               `json:"amount_in_decimal,omitempty"`
	Description            string               `json:"description,omitempty"`
	Taxable                *bool                `json:"taxable,omitempty"`
	TaxProfileId           string               `json:"tax_profile_id,omitempty"`
	AvalaraTaxCode         string               `json:"avalara_tax_code,omitempty"`
	HsnCode                string               `json:"hsn_code,omitempty"`
	TaxjarProductCode      string               `json:"taxjar_product_code,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	DateFrom               *int64               `json:"date_from,omitempty"`
	DateTo                 *int64               `json:"date_to,omitempty"`
}
type CreateRequestParams struct {
	SubscriptionId string                   `json:"subscription_id"`
	CurrencyCode   string                   `json:"currency_code,omitempty"`
	ItemPrices     []*CreateItemPriceParams `json:"item_prices,omitempty"`
	ItemTiers      []*CreateItemTierParams  `json:"item_tiers,omitempty"`
	Charges        []*CreateChargeParams    `json:"charges,omitempty"`
}
type CreateItemPriceParams struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type CreateItemTierParams struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type CreateChargeParams struct {
	Amount                 *int32               `json:"amount,omitempty"`
	AmountInDecimal        string               `json:"amount_in_decimal,omitempty"`
	Description            string               `json:"description,omitempty"`
	Taxable                *bool                `json:"taxable,omitempty"`
	TaxProfileId           string               `json:"tax_profile_id,omitempty"`
	AvalaraTaxCode         string               `json:"avalara_tax_code,omitempty"`
	HsnCode                string               `json:"hsn_code,omitempty"`
	TaxjarProductCode      string               `json:"taxjar_product_code,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	DateFrom               *int64               `json:"date_from,omitempty"`
	DateTo                 *int64               `json:"date_to,omitempty"`
}
type InvoiceUnbilledChargesRequestParams struct {
	SubscriptionId string `json:"subscription_id,omitempty"`
	CustomerId     string `json:"customer_id,omitempty"`
}
type ListRequestParams struct {
	Limit          *int32               `json:"limit,omitempty"`
	Offset         string               `json:"offset,omitempty"`
	IncludeDeleted *bool                `json:"include_deleted,omitempty"`
	IsVoided       *bool                `json:"is_voided,omitempty"`
	SubscriptionId *filter.StringFilter `json:"subscription_id,omitempty"`
	CustomerId     *filter.StringFilter `json:"customer_id,omitempty"`
}
type InvoiceNowEstimateRequestParams struct {
	SubscriptionId string `json:"subscription_id,omitempty"`
	CustomerId     string `json:"customer_id,omitempty"`
}
