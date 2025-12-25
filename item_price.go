package chargebee

import (
	"encoding/json"
)

type ItemPriceStatus string

const (
	ItemPriceStatusActive   ItemPriceStatus = "active"
	ItemPriceStatusArchived ItemPriceStatus = "archived"
	ItemPriceStatusDeleted  ItemPriceStatus = "deleted"
)

type ItemPriceProrationType string

const (
	ItemPriceProrationTypeSiteDefault ItemPriceProrationType = "site_default"
	ItemPriceProrationTypePartialTerm ItemPriceProrationType = "partial_term"
	ItemPriceProrationTypeFullTerm    ItemPriceProrationType = "full_term"
)

type ItemPricePricingModel string

const (
	ItemPricePricingModelFlatFee   ItemPricePricingModel = "flat_fee"
	ItemPricePricingModelPerUnit   ItemPricePricingModel = "per_unit"
	ItemPricePricingModelTiered    ItemPricePricingModel = "tiered"
	ItemPricePricingModelVolume    ItemPricePricingModel = "volume"
	ItemPricePricingModelStairstep ItemPricePricingModel = "stairstep"
)

type ItemPricePeriodUnit string

const (
	ItemPricePeriodUnitDay   ItemPricePeriodUnit = "day"
	ItemPricePeriodUnitWeek  ItemPricePeriodUnit = "week"
	ItemPricePeriodUnitMonth ItemPricePeriodUnit = "month"
	ItemPricePeriodUnitYear  ItemPricePeriodUnit = "year"
)

type ItemPriceTrialPeriodUnit string

const (
	ItemPriceTrialPeriodUnitDay   ItemPriceTrialPeriodUnit = "day"
	ItemPriceTrialPeriodUnitMonth ItemPriceTrialPeriodUnit = "month"
)

type ItemPriceTrialEndAction string

const (
	ItemPriceTrialEndActionSiteDefault          ItemPriceTrialEndAction = "site_default"
	ItemPriceTrialEndActionActivateSubscription ItemPriceTrialEndAction = "activate_subscription"
	ItemPriceTrialEndActionCancelSubscription   ItemPriceTrialEndAction = "cancel_subscription"
)

type ItemPriceShippingPeriodUnit string

const (
	ItemPriceShippingPeriodUnitDay   ItemPriceShippingPeriodUnit = "day"
	ItemPriceShippingPeriodUnitWeek  ItemPriceShippingPeriodUnit = "week"
	ItemPriceShippingPeriodUnitMonth ItemPriceShippingPeriodUnit = "month"
	ItemPriceShippingPeriodUnitYear  ItemPriceShippingPeriodUnit = "year"
)

type ItemPriceChannel string

const (
	ItemPriceChannelWeb       ItemPriceChannel = "web"
	ItemPriceChannelAppStore  ItemPriceChannel = "app_store"
	ItemPriceChannelPlayStore ItemPriceChannel = "play_store"
)

type ItemPriceUsageAccumulationResetFrequency string

const (
	ItemPriceUsageAccumulationResetFrequencyNever                        ItemPriceUsageAccumulationResetFrequency = "never"
	ItemPriceUsageAccumulationResetFrequencySubscriptionBillingFrequency ItemPriceUsageAccumulationResetFrequency = "subscription_billing_frequency"
)

type ItemPriceItemType string

const (
	ItemPriceItemTypePlan   ItemPriceItemType = "plan"
	ItemPriceItemTypeAddon  ItemPriceItemType = "addon"
	ItemPriceItemTypeCharge ItemPriceItemType = "charge"
)

type ItemPriceTierPricingType string

const (
	ItemPriceTierPricingTypePerUnit ItemPriceTierPricingType = "per_unit"
	ItemPriceTierPricingTypeFlatFee ItemPriceTierPricingType = "flat_fee"
	ItemPriceTierPricingTypePackage ItemPriceTierPricingType = "package"
)

type ItemPriceTaxDetailAvalaraSaleType string

const (
	ItemPriceTaxDetailAvalaraSaleTypeWholesale ItemPriceTaxDetailAvalaraSaleType = "wholesale"
	ItemPriceTaxDetailAvalaraSaleTypeRetail    ItemPriceTaxDetailAvalaraSaleType = "retail"
	ItemPriceTaxDetailAvalaraSaleTypeConsumed  ItemPriceTaxDetailAvalaraSaleType = "consumed"
	ItemPriceTaxDetailAvalaraSaleTypeVendorUse ItemPriceTaxDetailAvalaraSaleType = "vendor_use"
)

// just struct
type ItemPrice struct {
	Id                              string                                   `json:"id"`
	Name                            string                                   `json:"name"`
	ItemFamilyId                    string                                   `json:"item_family_id"`
	ItemId                          string                                   `json:"item_id"`
	Description                     string                                   `json:"description"`
	Status                          ItemPriceStatus                          `json:"status"`
	ExternalName                    string                                   `json:"external_name"`
	PriceVariantId                  string                                   `json:"price_variant_id"`
	ProrationType                   ItemPriceProrationType                   `json:"proration_type"`
	PricingModel                    ItemPricePricingModel                    `json:"pricing_model"`
	Price                           int64                                    `json:"price"`
	PriceInDecimal                  string                                   `json:"price_in_decimal"`
	Period                          int32                                    `json:"period"`
	CurrencyCode                    string                                   `json:"currency_code"`
	PeriodUnit                      ItemPricePeriodUnit                      `json:"period_unit"`
	TrialPeriod                     int32                                    `json:"trial_period"`
	TrialPeriodUnit                 ItemPriceTrialPeriodUnit                 `json:"trial_period_unit"`
	TrialEndAction                  ItemPriceTrialEndAction                  `json:"trial_end_action"`
	ShippingPeriod                  int32                                    `json:"shipping_period"`
	ShippingPeriodUnit              ItemPriceShippingPeriodUnit              `json:"shipping_period_unit"`
	BillingCycles                   int32                                    `json:"billing_cycles"`
	FreeQuantity                    int32                                    `json:"free_quantity"`
	FreeQuantityInDecimal           string                                   `json:"free_quantity_in_decimal"`
	Channel                         ItemPriceChannel                         `json:"channel"`
	ResourceVersion                 int64                                    `json:"resource_version"`
	UpdatedAt                       int64                                    `json:"updated_at"`
	CreatedAt                       int64                                    `json:"created_at"`
	UsageAccumulationResetFrequency ItemPriceUsageAccumulationResetFrequency `json:"usage_accumulation_reset_frequency"`
	ArchivedAt                      int64                                    `json:"archived_at"`
	InvoiceNotes                    string                                   `json:"invoice_notes"`
	Tiers                           []*ItemPriceTier                         `json:"tiers"`
	IsTaxable                       bool                                     `json:"is_taxable"`
	TaxDetail                       *TaxDetail                               `json:"tax_detail"`
	TaxProvidersFields              []*ItemPriceTaxProvidersField            `json:"tax_providers_fields"`
	AccountingDetail                *AccountingDetail                        `json:"accounting_detail"`
	Metadata                        json.RawMessage                          `json:"metadata"`
	ItemType                        ItemPriceItemType                        `json:"item_type"`
	ShowDescriptionInInvoices       bool                                     `json:"show_description_in_invoices"`
	ShowDescriptionInQuotes         bool                                     `json:"show_description_in_quotes"`
	Deleted                         bool                                     `json:"deleted"`
	BusinessEntityId                string                                   `json:"business_entity_id"`
	CustomField                     CustomField                              `json:"custom_field"`
	Object                          string                                   `json:"object"`
}

// sub resources
type ItemPriceTier struct {
	StartingUnit          int32                    `json:"starting_unit"`
	EndingUnit            int32                    `json:"ending_unit"`
	Price                 int64                    `json:"price"`
	StartingUnitInDecimal string                   `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string                   `json:"ending_unit_in_decimal"`
	PriceInDecimal        string                   `json:"price_in_decimal"`
	PricingType           ItemPriceTierPricingType `json:"pricing_type"`
	PackageSize           int32                    `json:"package_size"`
	Object                string                   `json:"object"`
}
type ItemPriceTaxDetail struct {
	TaxProfileId           string                            `json:"tax_profile_id"`
	AvalaraSaleType        ItemPriceTaxDetailAvalaraSaleType `json:"avalara_sale_type"`
	AvalaraTransactionType int32                             `json:"avalara_transaction_type"`
	AvalaraServiceType     int32                             `json:"avalara_service_type"`
	AvalaraTaxCode         string                            `json:"avalara_tax_code"`
	HsnCode                string                            `json:"hsn_code"`
	TaxjarProductCode      string                            `json:"taxjar_product_code"`
	Object                 string                            `json:"object"`
}
type ItemPriceTaxProvidersField struct {
	ProviderName string `json:"provider_name"`
	FieldId      string `json:"field_id"`
	FieldValue   string `json:"field_value"`
	Object       string `json:"object"`
}
type ItemPriceAccountingDetail struct {
	Sku                 string `json:"sku"`
	AccountingCode      string `json:"accounting_code"`
	AccountingCategory1 string `json:"accounting_category1"`
	AccountingCategory2 string `json:"accounting_category2"`
	AccountingCategory3 string `json:"accounting_category3"`
	AccountingCategory4 string `json:"accounting_category4"`
	Object              string `json:"object"`
}

// operations
// input params
type ItemPriceCreateRequest struct {
	Id                              string                                   `json:"id"`
	Name                            string                                   `json:"name"`
	Description                     string                                   `json:"description,omitempty"`
	ItemId                          string                                   `json:"item_id"`
	InvoiceNotes                    string                                   `json:"invoice_notes,omitempty"`
	ProrationType                   ItemPriceProrationType                   `json:"proration_type,omitempty"`
	ExternalName                    string                                   `json:"external_name,omitempty"`
	CurrencyCode                    string                                   `json:"currency_code,omitempty"`
	PriceVariantId                  string                                   `json:"price_variant_id,omitempty"`
	IsTaxable                       *bool                                    `json:"is_taxable,omitempty"`
	FreeQuantity                    *int32                                   `json:"free_quantity,omitempty"`
	FreeQuantityInDecimal           string                                   `json:"free_quantity_in_decimal,omitempty"`
	Metadata                        map[string]interface{}                   `json:"metadata,omitempty"`
	ShowDescriptionInInvoices       *bool                                    `json:"show_description_in_invoices,omitempty"`
	ShowDescriptionInQuotes         *bool                                    `json:"show_description_in_quotes,omitempty"`
	UsageAccumulationResetFrequency ItemPriceUsageAccumulationResetFrequency `json:"usage_accumulation_reset_frequency,omitempty"`
	BusinessEntityId                string                                   `json:"business_entity_id,omitempty"`
	PricingModel                    ItemPricePricingModel                    `json:"pricing_model,omitempty"`
	Tiers                           []*ItemPriceCreateTier                   `json:"tiers,omitempty"`
	Price                           *int64                                   `json:"price,omitempty"`
	PriceInDecimal                  string                                   `json:"price_in_decimal,omitempty"`
	PeriodUnit                      ItemPricePeriodUnit                      `json:"period_unit,omitempty"`
	Period                          *int32                                   `json:"period,omitempty"`
	TrialPeriodUnit                 ItemPriceTrialPeriodUnit                 `json:"trial_period_unit,omitempty"`
	TrialPeriod                     *int32                                   `json:"trial_period,omitempty"`
	ShippingPeriod                  *int32                                   `json:"shipping_period,omitempty"`
	ShippingPeriodUnit              ItemPriceShippingPeriodUnit              `json:"shipping_period_unit,omitempty"`
	BillingCycles                   *int32                                   `json:"billing_cycles,omitempty"`
	TrialEndAction                  ItemPriceTrialEndAction                  `json:"trial_end_action,omitempty"`
	TaxDetail                       *ItemPriceCreateTaxDetail                `json:"tax_detail,omitempty"`
	TaxProvidersFields              []*ItemPriceCreateTaxProvidersField      `json:"tax_providers_fields,omitempty"`
	AccountingDetail                *ItemPriceCreateAccountingDetail         `json:"accounting_detail,omitempty"`
	apiRequest                      `json:"-" form:"-"`
}

func (r *ItemPriceCreateRequest) payload() any { return r }

// input sub resource params multi
type ItemPriceCreateTier struct {
	StartingUnit          *int32      `json:"starting_unit,omitempty"`
	EndingUnit            *int32      `json:"ending_unit,omitempty"`
	Price                 *int64      `json:"price,omitempty"`
	StartingUnitInDecimal string      `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string      `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string      `json:"price_in_decimal,omitempty"`
	PricingType           PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32      `json:"package_size,omitempty"`
}

// input sub resource params single
type ItemPriceCreateTaxDetail struct {
	TaxProfileId           string          `json:"tax_profile_id,omitempty"`
	AvalaraTaxCode         string          `json:"avalara_tax_code,omitempty"`
	HsnCode                string          `json:"hsn_code,omitempty"`
	AvalaraSaleType        AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32          `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32          `json:"avalara_service_type,omitempty"`
	TaxjarProductCode      string          `json:"taxjar_product_code,omitempty"`
}

// input sub resource params multi
type ItemPriceCreateTaxProvidersField struct {
	ProviderName string `json:"provider_name"`
	FieldId      string `json:"field_id"`
	FieldValue   string `json:"field_value"`
}

// input sub resource params single
type ItemPriceCreateAccountingDetail struct {
	Sku                 string `json:"sku,omitempty"`
	AccountingCode      string `json:"accounting_code,omitempty"`
	AccountingCategory1 string `json:"accounting_category1,omitempty"`
	AccountingCategory2 string `json:"accounting_category2,omitempty"`
	AccountingCategory3 string `json:"accounting_category3,omitempty"`
	AccountingCategory4 string `json:"accounting_category4,omitempty"`
}
type ItemPriceUpdateRequest struct {
	Name                            string                                   `json:"name,omitempty"`
	Description                     string                                   `json:"description,omitempty"`
	ProrationType                   ItemPriceProrationType                   `json:"proration_type,omitempty"`
	PriceVariantId                  string                                   `json:"price_variant_id,omitempty"`
	Status                          ItemPriceStatus                          `json:"status,omitempty"`
	ExternalName                    string                                   `json:"external_name,omitempty"`
	UsageAccumulationResetFrequency ItemPriceUsageAccumulationResetFrequency `json:"usage_accumulation_reset_frequency,omitempty"`
	CurrencyCode                    string                                   `json:"currency_code,omitempty"`
	InvoiceNotes                    string                                   `json:"invoice_notes,omitempty"`
	IsTaxable                       *bool                                    `json:"is_taxable,omitempty"`
	FreeQuantity                    *int32                                   `json:"free_quantity,omitempty"`
	FreeQuantityInDecimal           string                                   `json:"free_quantity_in_decimal,omitempty"`
	Metadata                        map[string]interface{}                   `json:"metadata,omitempty"`
	PricingModel                    ItemPricePricingModel                    `json:"pricing_model,omitempty"`
	Tiers                           []*ItemPriceUpdateTier                   `json:"tiers,omitempty"`
	Price                           *int64                                   `json:"price,omitempty"`
	PriceInDecimal                  string                                   `json:"price_in_decimal,omitempty"`
	PeriodUnit                      ItemPricePeriodUnit                      `json:"period_unit,omitempty"`
	Period                          *int32                                   `json:"period,omitempty"`
	TrialPeriodUnit                 ItemPriceTrialPeriodUnit                 `json:"trial_period_unit,omitempty"`
	TrialPeriod                     *int32                                   `json:"trial_period,omitempty"`
	ShippingPeriod                  *int32                                   `json:"shipping_period,omitempty"`
	ShippingPeriodUnit              ItemPriceShippingPeriodUnit              `json:"shipping_period_unit,omitempty"`
	BillingCycles                   *int32                                   `json:"billing_cycles,omitempty"`
	TrialEndAction                  ItemPriceTrialEndAction                  `json:"trial_end_action,omitempty"`
	TaxDetail                       *ItemPriceUpdateTaxDetail                `json:"tax_detail,omitempty"`
	TaxProvidersFields              []*ItemPriceUpdateTaxProvidersField      `json:"tax_providers_fields,omitempty"`
	AccountingDetail                *ItemPriceUpdateAccountingDetail         `json:"accounting_detail,omitempty"`
	ShowDescriptionInInvoices       *bool                                    `json:"show_description_in_invoices,omitempty"`
	ShowDescriptionInQuotes         *bool                                    `json:"show_description_in_quotes,omitempty"`
	apiRequest                      `json:"-" form:"-"`
}

func (r *ItemPriceUpdateRequest) payload() any { return r }

// input sub resource params multi
type ItemPriceUpdateTier struct {
	StartingUnit          *int32      `json:"starting_unit,omitempty"`
	EndingUnit            *int32      `json:"ending_unit,omitempty"`
	Price                 *int64      `json:"price,omitempty"`
	StartingUnitInDecimal string      `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string      `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string      `json:"price_in_decimal,omitempty"`
	PricingType           PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32      `json:"package_size,omitempty"`
}

// input sub resource params single
type ItemPriceUpdateTaxDetail struct {
	TaxProfileId           string          `json:"tax_profile_id,omitempty"`
	AvalaraTaxCode         string          `json:"avalara_tax_code,omitempty"`
	HsnCode                string          `json:"hsn_code,omitempty"`
	AvalaraSaleType        AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32          `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32          `json:"avalara_service_type,omitempty"`
	TaxjarProductCode      string          `json:"taxjar_product_code,omitempty"`
}

// input sub resource params multi
type ItemPriceUpdateTaxProvidersField struct {
	ProviderName string `json:"provider_name"`
	FieldId      string `json:"field_id"`
	FieldValue   string `json:"field_value"`
}

// input sub resource params single
type ItemPriceUpdateAccountingDetail struct {
	Sku                 string `json:"sku,omitempty"`
	AccountingCode      string `json:"accounting_code,omitempty"`
	AccountingCategory1 string `json:"accounting_category1,omitempty"`
	AccountingCategory2 string `json:"accounting_category2,omitempty"`
	AccountingCategory3 string `json:"accounting_category3,omitempty"`
	AccountingCategory4 string `json:"accounting_category4,omitempty"`
}
type ItemPriceListRequest struct {
	Limit                     *int32           `json:"limit,omitempty"`
	Offset                    string           `json:"offset,omitempty"`
	Id                        *StringFilter    `json:"id,omitempty"`
	Name                      *StringFilter    `json:"name,omitempty"`
	PricingModel              *EnumFilter      `json:"pricing_model,omitempty"`
	ItemId                    *StringFilter    `json:"item_id,omitempty"`
	ItemFamilyId              *StringFilter    `json:"item_family_id,omitempty"`
	ItemType                  *EnumFilter      `json:"item_type,omitempty"`
	CurrencyCode              *StringFilter    `json:"currency_code,omitempty"`
	PriceVariantId            *StringFilter    `json:"price_variant_id,omitempty"`
	TrialPeriod               *NumberFilter    `json:"trial_period,omitempty"`
	TrialPeriodUnit           *EnumFilter      `json:"trial_period_unit,omitempty"`
	Status                    *EnumFilter      `json:"status,omitempty"`
	UpdatedAt                 *TimestampFilter `json:"updated_at,omitempty"`
	BusinessEntityId          *StringFilter    `json:"business_entity_id,omitempty"`
	IncludeSiteLevelResources *BooleanFilter   `json:"include_site_level_resources,omitempty"`
	PeriodUnit                *EnumFilter      `json:"period_unit,omitempty"`
	Period                    *NumberFilter    `json:"period,omitempty"`
	Channel                   *EnumFilter      `json:"channel,omitempty"`
	SortBy                    *SortFilter      `json:"sort_by,omitempty"`
	apiRequest                `json:"-" form:"-"`
}

func (r *ItemPriceListRequest) payload() any { return r }

type ItemPriceFindApplicableItemsRequest struct {
	Limit      *int32      `json:"limit,omitempty"`
	Offset     string      `json:"offset,omitempty"`
	SortBy     *SortFilter `json:"sort_by,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *ItemPriceFindApplicableItemsRequest) payload() any { return r }

type ItemPriceFindApplicableItemPricesRequest struct {
	Limit      *int32      `json:"limit,omitempty"`
	Offset     string      `json:"offset,omitempty"`
	ItemId     string      `json:"item_id,omitempty"`
	SortBy     *SortFilter `json:"sort_by,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *ItemPriceFindApplicableItemPricesRequest) payload() any { return r }

type ItemPriceMoveItemPriceRequest struct {
	DestinationItemId string `json:"destination_item_id"`
	apiRequest        `json:"-" form:"-"`
}

func (r *ItemPriceMoveItemPriceRequest) payload() any { return r }

// operation response
type ItemPriceCreateResponse struct {
	ItemPrice *ItemPrice `json:"item_price,omitempty"`
	apiResponse
}

// operation response
type ItemPriceRetrieveResponse struct {
	ItemPrice *ItemPrice `json:"item_price,omitempty"`
	apiResponse
}

// operation response
type ItemPriceUpdateResponse struct {
	ItemPrice *ItemPrice `json:"item_price,omitempty"`
	apiResponse
}

// operation sub response
type ItemPriceListItemPriceResponse struct {
	ItemPrice *ItemPrice `json:"item_price,omitempty"`
}

// operation response
type ItemPriceListResponse struct {
	List       []*ItemPriceListItemPriceResponse `json:"list,omitempty"`
	NextOffset string                            `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type ItemPriceDeleteResponse struct {
	ItemPrice *ItemPrice `json:"item_price,omitempty"`
	apiResponse
}

// operation sub response
type ItemPriceFindApplicableItemsItemPriceResponse struct {
	Item Item `json:"item,omitempty"`
}

// operation response
type ItemPriceFindApplicableItemsResponse struct {
	List       []*ItemPriceFindApplicableItemsItemPriceResponse `json:"list,omitempty"`
	NextOffset string                                           `json:"next_offset,omitempty"`
	apiResponse
}

// operation sub response
type ItemPriceFindApplicableItemPricesItemPriceResponse struct {
	ItemPrice *ItemPrice `json:"item_price,omitempty"`
}

// operation response
type ItemPriceFindApplicableItemPricesResponse struct {
	List       []*ItemPriceFindApplicableItemPricesItemPriceResponse `json:"list,omitempty"`
	NextOffset string                                                `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type ItemPriceMoveItemPriceResponse struct {
	ItemPrice *ItemPrice `json:"item_price,omitempty"`
	apiResponse
}
