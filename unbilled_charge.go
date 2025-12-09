package chargebee

type EntityType string

const (
	EntityTypeAdhoc           EntityType = "adhoc"
	EntityTypePlanItemPrice   EntityType = "plan_item_price"
	EntityTypeAddonItemPrice  EntityType = "addon_item_price"
	EntityTypeChargeItemPrice EntityType = "charge_item_price"
	EntityTypePlanSetup       EntityType = "plan_setup"
	EntityTypePlan            EntityType = "plan"
	EntityTypeAddon           EntityType = "addon"
)

type UnbilledCharge struct {
	Id                  string            `json:"id"`
	CustomerId          string            `json:"customer_id"`
	SubscriptionId      string            `json:"subscription_id"`
	DateFrom            int64             `json:"date_from"`
	DateTo              int64             `json:"date_to"`
	UnitAmount          int64             `json:"unit_amount"`
	PricingModel        enum.PricingModel `json:"pricing_model"`
	Quantity            int32             `json:"quantity"`
	Amount              int64             `json:"amount"`
	CurrencyCode        string            `json:"currency_code"`
	DiscountAmount      int64             `json:"discount_amount"`
	Description         string            `json:"description"`
	EntityType          EntityType        `json:"entity_type"`
	EntityId            string            `json:"entity_id"`
	IsVoided            bool              `json:"is_voided"`
	VoidedAt            int64             `json:"voided_at"`
	UnitAmountInDecimal string            `json:"unit_amount_in_decimal"`
	QuantityInDecimal   string            `json:"quantity_in_decimal"`
	AmountInDecimal     string            `json:"amount_in_decimal"`
	UpdatedAt           int64             `json:"updated_at"`
	Tiers               []*Tier           `json:"tiers"`
	IsAdvanceCharge     bool              `json:"is_advance_charge"`
	BusinessEntityId    string            `json:"business_entity_id"`
	Deleted             bool              `json:"deleted"`
	Object              string            `json:"object"`
}
type Tier struct {
	StartingUnit          int32            `json:"starting_unit"`
	EndingUnit            int32            `json:"ending_unit"`
	QuantityUsed          int32            `json:"quantity_used"`
	UnitAmount            int64            `json:"unit_amount"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal"`
	QuantityUsedInDecimal string           `json:"quantity_used_in_decimal"`
	UnitAmountInDecimal   string           `json:"unit_amount_in_decimal"`
	PricingType           enum.PricingType `json:"pricing_type"`
	PackageSize           int32            `json:"package_size"`
	Object                string           `json:"object"`
}
type CreateUnbilledChargeRequest struct {
	SubscriptionId     string                                   `json:"subscription_id"`
	CurrencyCode       string                                   `json:"currency_code,omitempty"`
	Addons             []*CreateUnbilledChargeAddon             `json:"addons,omitempty"`
	Charges            []*CreateUnbilledChargeCharge            `json:"charges,omitempty"`
	TaxProvidersFields []*CreateUnbilledChargeTaxProvidersField `json:"tax_providers_fields,omitempty"`
}
type CreateUnbilledChargeAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type CreateUnbilledChargeCharge struct {
	Amount                 *int64               `json:"amount,omitempty"`
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
type CreateUnbilledChargeTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type CreateRequest struct {
	SubscriptionId     string                     `json:"subscription_id"`
	CurrencyCode       string                     `json:"currency_code,omitempty"`
	ItemPrices         []*CreateItemPrice         `json:"item_prices,omitempty"`
	ItemTiers          []*CreateItemTier          `json:"item_tiers,omitempty"`
	Charges            []*CreateCharge            `json:"charges,omitempty"`
	TaxProvidersFields []*CreateTaxProvidersField `json:"tax_providers_fields,omitempty"`
}
type CreateItemPrice struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type CreateItemTier struct {
	ItemPriceId           string           `json:"item_price_id,omitempty"`
	StartingUnit          *int32           `json:"starting_unit,omitempty"`
	EndingUnit            *int32           `json:"ending_unit,omitempty"`
	Price                 *int64           `json:"price,omitempty"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string           `json:"price_in_decimal,omitempty"`
	PricingType           enum.PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32           `json:"package_size,omitempty"`
}
type CreateCharge struct {
	Amount                 *int64               `json:"amount,omitempty"`
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
type CreateTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type InvoiceUnbilledChargesRequest struct {
	SubscriptionId string `json:"subscription_id,omitempty"`
	CustomerId     string `json:"customer_id,omitempty"`
}
type ListRequest struct {
	Limit          *int32               `json:"limit,omitempty"`
	Offset         string               `json:"offset,omitempty"`
	IncludeDeleted *bool                `json:"include_deleted,omitempty"`
	IsVoided       *bool                `json:"is_voided,omitempty"`
	SubscriptionId *filter.StringFilter `json:"subscription_id,omitempty"`
	CustomerId     *filter.StringFilter `json:"customer_id,omitempty"`
}
type InvoiceNowEstimateRequest struct {
	SubscriptionId string `json:"subscription_id,omitempty"`
	CustomerId     string `json:"customer_id,omitempty"`
}

type CreateUnbilledChargeResponse struct {
	UnbilledCharges []*UnbilledCharge `json:"unbilled_charges,omitempty"`
}

type CreateResponse struct {
	UnbilledCharges []*UnbilledCharge `json:"unbilled_charges,omitempty"`
}

type InvoiceUnbilledChargesResponse struct {
	Invoices []*invoice.Invoice `json:"invoices,omitempty"`
}

type DeleteResponse struct {
	UnbilledCharge *UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type ListUnbilledChargeResponse struct {
	UnbilledCharge *UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type ListResponse struct {
	List       []*ListUnbilledChargeResponse `json:"list,omitempty"`
	NextOffset string                        `json:"next_offset,omitempty"`
}

type InvoiceNowEstimateResponse struct {
	Estimate *estimate.Estimate `json:"estimate,omitempty"`
}
