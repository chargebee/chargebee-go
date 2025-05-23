package itemprice

import (
	"encoding/json"
	"github.com/chargebee/chargebee-go/v3/enum"
	"github.com/chargebee/chargebee-go/v3/filter"
	itemPriceEnum "github.com/chargebee/chargebee-go/v3/models/itemprice/enum"
)

type ItemPrice struct {
	Id                              string                               `json:"id"`
	Name                            string                               `json:"name"`
	ItemFamilyId                    string                               `json:"item_family_id"`
	ItemId                          string                               `json:"item_id"`
	Description                     string                               `json:"description"`
	Status                          itemPriceEnum.Status                 `json:"status"`
	ExternalName                    string                               `json:"external_name"`
	PriceVariantId                  string                               `json:"price_variant_id"`
	ProrationType                   itemPriceEnum.ProrationType          `json:"proration_type"`
	PricingModel                    enum.PricingModel                    `json:"pricing_model"`
	Price                           int64                                `json:"price"`
	PriceInDecimal                  string                               `json:"price_in_decimal"`
	Period                          int32                                `json:"period"`
	CurrencyCode                    string                               `json:"currency_code"`
	PeriodUnit                      itemPriceEnum.PeriodUnit             `json:"period_unit"`
	TrialPeriod                     int32                                `json:"trial_period"`
	TrialPeriodUnit                 itemPriceEnum.TrialPeriodUnit        `json:"trial_period_unit"`
	TrialEndAction                  itemPriceEnum.TrialEndAction         `json:"trial_end_action"`
	ShippingPeriod                  int32                                `json:"shipping_period"`
	ShippingPeriodUnit              itemPriceEnum.ShippingPeriodUnit     `json:"shipping_period_unit"`
	BillingCycles                   int32                                `json:"billing_cycles"`
	FreeQuantity                    int32                                `json:"free_quantity"`
	FreeQuantityInDecimal           string                               `json:"free_quantity_in_decimal"`
	Channel                         enum.Channel                         `json:"channel"`
	ResourceVersion                 int64                                `json:"resource_version"`
	UpdatedAt                       int64                                `json:"updated_at"`
	CreatedAt                       int64                                `json:"created_at"`
	UsageAccumulationResetFrequency enum.UsageAccumulationResetFrequency `json:"usage_accumulation_reset_frequency"`
	ArchivedAt                      int64                                `json:"archived_at"`
	InvoiceNotes                    string                               `json:"invoice_notes"`
	Tiers                           []*Tier                              `json:"tiers"`
	IsTaxable                       bool                                 `json:"is_taxable"`
	TaxDetail                       *TaxDetail                           `json:"tax_detail"`
	TaxProvidersFields              []*TaxProvidersField                 `json:"tax_providers_fields"`
	AccountingDetail                *AccountingDetail                    `json:"accounting_detail"`
	Metadata                        json.RawMessage                      `json:"metadata"`
	ItemType                        enum.ItemType                        `json:"item_type"`
	Archivable                      bool                                 `json:"archivable"`
	ParentItemId                    string                               `json:"parent_item_id"`
	ShowDescriptionInInvoices       bool                                 `json:"show_description_in_invoices"`
	ShowDescriptionInQuotes         bool                                 `json:"show_description_in_quotes"`
	Deleted                         bool                                 `json:"deleted"`
	BusinessEntityId                string                               `json:"business_entity_id"`
	CustomField                     map[string]interface{}               `json:"custom_field"`
	Object                          string                               `json:"object"`
}
type Tier struct {
	StartingUnit          int32            `json:"starting_unit"`
	EndingUnit            int32            `json:"ending_unit"`
	Price                 int64            `json:"price"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal"`
	PriceInDecimal        string           `json:"price_in_decimal"`
	PricingType           enum.PricingType `json:"pricing_type"`
	PackageSize           int32            `json:"package_size"`
	Object                string           `json:"object"`
}
type TaxDetail struct {
	TaxProfileId           string               `json:"tax_profile_id"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type"`
	AvalaraTransactionType int32                `json:"avalara_transaction_type"`
	AvalaraServiceType     int32                `json:"avalara_service_type"`
	AvalaraTaxCode         string               `json:"avalara_tax_code"`
	HsnCode                string               `json:"hsn_code"`
	TaxjarProductCode      string               `json:"taxjar_product_code"`
	Object                 string               `json:"object"`
}
type TaxProvidersField struct {
	ProviderName string `json:"provider_name"`
	FieldId      string `json:"field_id"`
	FieldValue   string `json:"field_value"`
	Object       string `json:"object"`
}
type AccountingDetail struct {
	Sku                 string `json:"sku"`
	AccountingCode      string `json:"accounting_code"`
	AccountingCategory1 string `json:"accounting_category1"`
	AccountingCategory2 string `json:"accounting_category2"`
	AccountingCategory3 string `json:"accounting_category3"`
	AccountingCategory4 string `json:"accounting_category4"`
	Object              string `json:"object"`
}
type CreateRequestParams struct {
	Id                              string                               `json:"id"`
	Name                            string                               `json:"name"`
	Description                     string                               `json:"description,omitempty"`
	ItemId                          string                               `json:"item_id"`
	InvoiceNotes                    string                               `json:"invoice_notes,omitempty"`
	ProrationType                   itemPriceEnum.ProrationType          `json:"proration_type,omitempty"`
	ExternalName                    string                               `json:"external_name,omitempty"`
	CurrencyCode                    string                               `json:"currency_code,omitempty"`
	PriceVariantId                  string                               `json:"price_variant_id,omitempty"`
	IsTaxable                       *bool                                `json:"is_taxable,omitempty"`
	FreeQuantity                    *int32                               `json:"free_quantity,omitempty"`
	FreeQuantityInDecimal           string                               `json:"free_quantity_in_decimal,omitempty"`
	Metadata                        map[string]interface{}               `json:"metadata,omitempty"`
	ShowDescriptionInInvoices       *bool                                `json:"show_description_in_invoices,omitempty"`
	ShowDescriptionInQuotes         *bool                                `json:"show_description_in_quotes,omitempty"`
	UsageAccumulationResetFrequency enum.UsageAccumulationResetFrequency `json:"usage_accumulation_reset_frequency,omitempty"`
	BusinessEntityId                string                               `json:"business_entity_id,omitempty"`
	PricingModel                    enum.PricingModel                    `json:"pricing_model,omitempty"`
	Tiers                           []*CreateTierParams                  `json:"tiers,omitempty"`
	Price                           *int64                               `json:"price,omitempty"`
	PriceInDecimal                  string                               `json:"price_in_decimal,omitempty"`
	PeriodUnit                      itemPriceEnum.PeriodUnit             `json:"period_unit,omitempty"`
	Period                          *int32                               `json:"period,omitempty"`
	TrialPeriodUnit                 itemPriceEnum.TrialPeriodUnit        `json:"trial_period_unit,omitempty"`
	TrialPeriod                     *int32                               `json:"trial_period,omitempty"`
	ShippingPeriod                  *int32                               `json:"shipping_period,omitempty"`
	ShippingPeriodUnit              itemPriceEnum.ShippingPeriodUnit     `json:"shipping_period_unit,omitempty"`
	BillingCycles                   *int32                               `json:"billing_cycles,omitempty"`
	TrialEndAction                  itemPriceEnum.TrialEndAction         `json:"trial_end_action,omitempty"`
	TaxDetail                       *CreateTaxDetailParams               `json:"tax_detail,omitempty"`
	TaxProvidersFields              []*CreateTaxProvidersFieldParams     `json:"tax_providers_fields,omitempty"`
	AccountingDetail                *CreateAccountingDetailParams        `json:"accounting_detail,omitempty"`
}
type CreateTierParams struct {
	StartingUnit          *int32           `json:"starting_unit,omitempty"`
	EndingUnit            *int32           `json:"ending_unit,omitempty"`
	Price                 *int64           `json:"price,omitempty"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string           `json:"price_in_decimal,omitempty"`
	PricingType           enum.PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32           `json:"package_size,omitempty"`
}
type CreateTaxDetailParams struct {
	TaxProfileId           string               `json:"tax_profile_id,omitempty"`
	AvalaraTaxCode         string               `json:"avalara_tax_code,omitempty"`
	HsnCode                string               `json:"hsn_code,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	TaxjarProductCode      string               `json:"taxjar_product_code,omitempty"`
}
type CreateTaxProvidersFieldParams struct {
	ProviderName string `json:"provider_name"`
	FieldId      string `json:"field_id"`
	FieldValue   string `json:"field_value"`
}
type CreateAccountingDetailParams struct {
	Sku                 string `json:"sku,omitempty"`
	AccountingCode      string `json:"accounting_code,omitempty"`
	AccountingCategory1 string `json:"accounting_category1,omitempty"`
	AccountingCategory2 string `json:"accounting_category2,omitempty"`
	AccountingCategory3 string `json:"accounting_category3,omitempty"`
	AccountingCategory4 string `json:"accounting_category4,omitempty"`
}
type UpdateRequestParams struct {
	Name                            string                               `json:"name,omitempty"`
	Description                     string                               `json:"description,omitempty"`
	ProrationType                   itemPriceEnum.ProrationType          `json:"proration_type,omitempty"`
	PriceVariantId                  string                               `json:"price_variant_id,omitempty"`
	Status                          itemPriceEnum.Status                 `json:"status,omitempty"`
	ExternalName                    string                               `json:"external_name,omitempty"`
	UsageAccumulationResetFrequency enum.UsageAccumulationResetFrequency `json:"usage_accumulation_reset_frequency,omitempty"`
	CurrencyCode                    string                               `json:"currency_code,omitempty"`
	InvoiceNotes                    string                               `json:"invoice_notes,omitempty"`
	IsTaxable                       *bool                                `json:"is_taxable,omitempty"`
	FreeQuantity                    *int32                               `json:"free_quantity,omitempty"`
	FreeQuantityInDecimal           string                               `json:"free_quantity_in_decimal,omitempty"`
	Metadata                        map[string]interface{}               `json:"metadata,omitempty"`
	PricingModel                    enum.PricingModel                    `json:"pricing_model,omitempty"`
	Tiers                           []*UpdateTierParams                  `json:"tiers,omitempty"`
	Price                           *int64                               `json:"price,omitempty"`
	PriceInDecimal                  string                               `json:"price_in_decimal,omitempty"`
	PeriodUnit                      itemPriceEnum.PeriodUnit             `json:"period_unit,omitempty"`
	Period                          *int32                               `json:"period,omitempty"`
	TrialPeriodUnit                 itemPriceEnum.TrialPeriodUnit        `json:"trial_period_unit,omitempty"`
	TrialPeriod                     *int32                               `json:"trial_period,omitempty"`
	ShippingPeriod                  *int32                               `json:"shipping_period,omitempty"`
	ShippingPeriodUnit              itemPriceEnum.ShippingPeriodUnit     `json:"shipping_period_unit,omitempty"`
	BillingCycles                   *int32                               `json:"billing_cycles,omitempty"`
	TrialEndAction                  itemPriceEnum.TrialEndAction         `json:"trial_end_action,omitempty"`
	TaxDetail                       *UpdateTaxDetailParams               `json:"tax_detail,omitempty"`
	TaxProvidersFields              []*UpdateTaxProvidersFieldParams     `json:"tax_providers_fields,omitempty"`
	AccountingDetail                *UpdateAccountingDetailParams        `json:"accounting_detail,omitempty"`
	ShowDescriptionInInvoices       *bool                                `json:"show_description_in_invoices,omitempty"`
	ShowDescriptionInQuotes         *bool                                `json:"show_description_in_quotes,omitempty"`
}
type UpdateTierParams struct {
	StartingUnit          *int32           `json:"starting_unit,omitempty"`
	EndingUnit            *int32           `json:"ending_unit,omitempty"`
	Price                 *int64           `json:"price,omitempty"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string           `json:"price_in_decimal,omitempty"`
	PricingType           enum.PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32           `json:"package_size,omitempty"`
}
type UpdateTaxDetailParams struct {
	TaxProfileId           string               `json:"tax_profile_id,omitempty"`
	AvalaraTaxCode         string               `json:"avalara_tax_code,omitempty"`
	HsnCode                string               `json:"hsn_code,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	TaxjarProductCode      string               `json:"taxjar_product_code,omitempty"`
}
type UpdateTaxProvidersFieldParams struct {
	ProviderName string `json:"provider_name"`
	FieldId      string `json:"field_id"`
	FieldValue   string `json:"field_value"`
}
type UpdateAccountingDetailParams struct {
	Sku                 string `json:"sku,omitempty"`
	AccountingCode      string `json:"accounting_code,omitempty"`
	AccountingCategory1 string `json:"accounting_category1,omitempty"`
	AccountingCategory2 string `json:"accounting_category2,omitempty"`
	AccountingCategory3 string `json:"accounting_category3,omitempty"`
	AccountingCategory4 string `json:"accounting_category4,omitempty"`
}
type ListRequestParams struct {
	Limit                     *int32                  `json:"limit,omitempty"`
	Offset                    string                  `json:"offset,omitempty"`
	Id                        *filter.StringFilter    `json:"id,omitempty"`
	Name                      *filter.StringFilter    `json:"name,omitempty"`
	PricingModel              *filter.EnumFilter      `json:"pricing_model,omitempty"`
	ItemId                    *filter.StringFilter    `json:"item_id,omitempty"`
	ItemFamilyId              *filter.StringFilter    `json:"item_family_id,omitempty"`
	ItemType                  *filter.EnumFilter      `json:"item_type,omitempty"`
	CurrencyCode              *filter.StringFilter    `json:"currency_code,omitempty"`
	PriceVariantId            *filter.StringFilter    `json:"price_variant_id,omitempty"`
	TrialPeriod               *filter.NumberFilter    `json:"trial_period,omitempty"`
	TrialPeriodUnit           *filter.EnumFilter      `json:"trial_period_unit,omitempty"`
	Status                    *filter.EnumFilter      `json:"status,omitempty"`
	UpdatedAt                 *filter.TimestampFilter `json:"updated_at,omitempty"`
	BusinessEntityId          *filter.StringFilter    `json:"business_entity_id,omitempty"`
	IncludeSiteLevelResources *filter.BooleanFilter   `json:"include_site_level_resources,omitempty"`
	PeriodUnit                *filter.EnumFilter      `json:"period_unit,omitempty"`
	Period                    *filter.NumberFilter    `json:"period,omitempty"`
	Channel                   *filter.EnumFilter      `json:"channel,omitempty"`
	SortBy                    *filter.SortFilter      `json:"sort_by,omitempty"`
}
type FindApplicableItemsRequestParams struct {
	Limit  *int32             `json:"limit,omitempty"`
	Offset string             `json:"offset,omitempty"`
	SortBy *filter.SortFilter `json:"sort_by,omitempty"`
}
type FindApplicableItemPricesRequestParams struct {
	Limit  *int32             `json:"limit,omitempty"`
	Offset string             `json:"offset,omitempty"`
	ItemId string             `json:"item_id,omitempty"`
	SortBy *filter.SortFilter `json:"sort_by,omitempty"`
}
