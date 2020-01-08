package addon

import (
	"encoding/json"
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	addonEnum "github.com/chargebee/chargebee-go/models/addon/enum"
)

type Addon struct {
	Id                          string                                `json:"id"`
	Name                        string                                `json:"name"`
	InvoiceName                 string                                `json:"invoice_name"`
	Description                 string                                `json:"description"`
	PricingModel                enum.PricingModel                     `json:"pricing_model"`
	Type                        addonEnum.Type                        `json:"type"`
	ChargeType                  addonEnum.ChargeType                  `json:"charge_type"`
	Price                       int32                                 `json:"price"`
	CurrencyCode                string                                `json:"currency_code"`
	Period                      int32                                 `json:"period"`
	PeriodUnit                  addonEnum.PeriodUnit                  `json:"period_unit"`
	Unit                        string                                `json:"unit"`
	Status                      addonEnum.Status                      `json:"status"`
	ArchivedAt                  int64                                 `json:"archived_at"`
	EnabledInPortal             bool                                  `json:"enabled_in_portal"`
	TaxCode                     string                                `json:"tax_code"`
	TaxjarProductCode           string                                `json:"taxjar_product_code"`
	AvalaraSaleType             enum.AvalaraSaleType                  `json:"avalara_sale_type"`
	AvalaraTransactionType      int32                                 `json:"avalara_transaction_type"`
	AvalaraServiceType          int32                                 `json:"avalara_service_type"`
	Sku                         string                                `json:"sku"`
	AccountingCode              string                                `json:"accounting_code"`
	AccountingCategory1         string                                `json:"accounting_category1"`
	AccountingCategory2         string                                `json:"accounting_category2"`
	IsShippable                 bool                                  `json:"is_shippable"`
	ShippingFrequencyPeriod     int32                                 `json:"shipping_frequency_period"`
	ShippingFrequencyPeriodUnit addonEnum.ShippingFrequencyPeriodUnit `json:"shipping_frequency_period_unit"`
	ResourceVersion             int64                                 `json:"resource_version"`
	UpdatedAt                   int64                                 `json:"updated_at"`
	InvoiceNotes                string                                `json:"invoice_notes"`
	Taxable                     bool                                  `json:"taxable"`
	TaxProfileId                string                                `json:"tax_profile_id"`
	MetaData                    json.RawMessage                       `json:"meta_data"`
	Tiers                       []*Tier                               `json:"tiers"`
	CustomField                 map[string]interface{}                `json:"custom_field"`
	Object                      string                                `json:"object"`
}
type Tier struct {
	StartingUnit int32  `json:"starting_unit"`
	EndingUnit   int32  `json:"ending_unit"`
	Price        int32  `json:"price"`
	Object       string `json:"object"`
}
type CreateRequestParams struct {
	Id                          string                                `json:"id"`
	Name                        string                                `json:"name"`
	InvoiceName                 string                                `json:"invoice_name,omitempty"`
	Description                 string                                `json:"description,omitempty"`
	ChargeType                  addonEnum.ChargeType                  `json:"charge_type"`
	Price                       *int32                                `json:"price,omitempty"`
	Tiers                       []*CreateTierParams                   `json:"tiers,omitempty"`
	CurrencyCode                string                                `json:"currency_code,omitempty"`
	Period                      *int32                                `json:"period,omitempty"`
	PeriodUnit                  addonEnum.PeriodUnit                  `json:"period_unit,omitempty"`
	PricingModel                enum.PricingModel                     `json:"pricing_model,omitempty"`
	Type                        addonEnum.Type                        `json:"type,omitempty"`
	Unit                        string                                `json:"unit,omitempty"`
	EnabledInPortal             *bool                                 `json:"enabled_in_portal,omitempty"`
	Taxable                     *bool                                 `json:"taxable,omitempty"`
	TaxProfileId                string                                `json:"tax_profile_id,omitempty"`
	AvalaraSaleType             enum.AvalaraSaleType                  `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType      *int32                                `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType          *int32                                `json:"avalara_service_type,omitempty"`
	TaxCode                     string                                `json:"tax_code,omitempty"`
	TaxjarProductCode           string                                `json:"taxjar_product_code,omitempty"`
	InvoiceNotes                string                                `json:"invoice_notes,omitempty"`
	MetaData                    map[string]interface{}                `json:"meta_data,omitempty"`
	Sku                         string                                `json:"sku,omitempty"`
	AccountingCode              string                                `json:"accounting_code,omitempty"`
	AccountingCategory1         string                                `json:"accounting_category1,omitempty"`
	AccountingCategory2         string                                `json:"accounting_category2,omitempty"`
	IsShippable                 *bool                                 `json:"is_shippable,omitempty"`
	ShippingFrequencyPeriod     *int32                                `json:"shipping_frequency_period,omitempty"`
	ShippingFrequencyPeriodUnit addonEnum.ShippingFrequencyPeriodUnit `json:"shipping_frequency_period_unit,omitempty"`
	Status                      addonEnum.Status                      `json:"status,omitempty"`
}
type CreateTierParams struct {
	StartingUnit *int32 `json:"starting_unit,omitempty"`
	EndingUnit   *int32 `json:"ending_unit,omitempty"`
	Price        *int32 `json:"price,omitempty"`
}
type UpdateRequestParams struct {
	Name                        string                                `json:"name,omitempty"`
	InvoiceName                 string                                `json:"invoice_name,omitempty"`
	Description                 string                                `json:"description,omitempty"`
	ChargeType                  addonEnum.ChargeType                  `json:"charge_type,omitempty"`
	Price                       *int32                                `json:"price,omitempty"`
	Tiers                       []*UpdateTierParams                   `json:"tiers,omitempty"`
	CurrencyCode                string                                `json:"currency_code,omitempty"`
	Period                      *int32                                `json:"period,omitempty"`
	PeriodUnit                  addonEnum.PeriodUnit                  `json:"period_unit,omitempty"`
	PricingModel                enum.PricingModel                     `json:"pricing_model,omitempty"`
	Type                        addonEnum.Type                        `json:"type,omitempty"`
	Unit                        string                                `json:"unit,omitempty"`
	EnabledInPortal             *bool                                 `json:"enabled_in_portal,omitempty"`
	Taxable                     *bool                                 `json:"taxable,omitempty"`
	TaxProfileId                string                                `json:"tax_profile_id,omitempty"`
	AvalaraSaleType             enum.AvalaraSaleType                  `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType      *int32                                `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType          *int32                                `json:"avalara_service_type,omitempty"`
	TaxCode                     string                                `json:"tax_code,omitempty"`
	TaxjarProductCode           string                                `json:"taxjar_product_code,omitempty"`
	InvoiceNotes                string                                `json:"invoice_notes,omitempty"`
	MetaData                    map[string]interface{}                `json:"meta_data,omitempty"`
	Sku                         string                                `json:"sku,omitempty"`
	AccountingCode              string                                `json:"accounting_code,omitempty"`
	AccountingCategory1         string                                `json:"accounting_category1,omitempty"`
	AccountingCategory2         string                                `json:"accounting_category2,omitempty"`
	IsShippable                 *bool                                 `json:"is_shippable,omitempty"`
	ShippingFrequencyPeriod     *int32                                `json:"shipping_frequency_period,omitempty"`
	ShippingFrequencyPeriodUnit addonEnum.ShippingFrequencyPeriodUnit `json:"shipping_frequency_period_unit,omitempty"`
}
type UpdateTierParams struct {
	StartingUnit *int32 `json:"starting_unit,omitempty"`
	EndingUnit   *int32 `json:"ending_unit,omitempty"`
	Price        *int32 `json:"price,omitempty"`
}
type ListRequestParams struct {
	Limit        *int32                  `json:"limit,omitempty"`
	Offset       string                  `json:"offset,omitempty"`
	Id           *filter.StringFilter    `json:"id,omitempty"`
	Name         *filter.StringFilter    `json:"name,omitempty"`
	PricingModel *filter.EnumFilter      `json:"pricing_model,omitempty"`
	Type         *filter.EnumFilter      `json:"type,omitempty"`
	ChargeType   *filter.EnumFilter      `json:"charge_type,omitempty"`
	Price        *filter.NumberFilter    `json:"price,omitempty"`
	Period       *filter.NumberFilter    `json:"period,omitempty"`
	PeriodUnit   *filter.EnumFilter      `json:"period_unit,omitempty"`
	Status       *filter.EnumFilter      `json:"status,omitempty"`
	UpdatedAt    *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type CopyRequestParams struct {
	FromSite       string `json:"from_site"`
	IdAtFromSite   string `json:"id_at_from_site"`
	Id             string `json:"id,omitempty"`
	ForSiteMerging *bool  `json:"for_site_merging,omitempty"`
}
