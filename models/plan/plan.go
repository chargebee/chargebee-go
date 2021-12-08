package plan

import (
	"encoding/json"
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	planEnum "github.com/chargebee/chargebee-go/models/plan/enum"
)

type Plan struct {
	Id                          string                               `json:"id"`
	Name                        string                               `json:"name"`
	InvoiceName                 string                               `json:"invoice_name"`
	Description                 string                               `json:"description"`
	Price                       int32                                `json:"price"`
	CurrencyCode                string                               `json:"currency_code"`
	Period                      int32                                `json:"period"`
	PeriodUnit                  planEnum.PeriodUnit                  `json:"period_unit"`
	TrialPeriod                 int32                                `json:"trial_period"`
	TrialPeriodUnit             planEnum.TrialPeriodUnit             `json:"trial_period_unit"`
	TrialEndAction              planEnum.TrialEndAction              `json:"trial_end_action"`
	PricingModel                enum.PricingModel                    `json:"pricing_model"`
	ChargeModel                 planEnum.ChargeModel                 `json:"charge_model"`
	FreeQuantity                int32                                `json:"free_quantity"`
	SetupCost                   int32                                `json:"setup_cost"`
	DowngradePenalty            float64                              `json:"downgrade_penalty"`
	Status                      planEnum.Status                      `json:"status"`
	ArchivedAt                  int64                                `json:"archived_at"`
	BillingCycles               int32                                `json:"billing_cycles"`
	RedirectUrl                 string                               `json:"redirect_url"`
	EnabledInHostedPages        bool                                 `json:"enabled_in_hosted_pages"`
	EnabledInPortal             bool                                 `json:"enabled_in_portal"`
	AddonApplicability          planEnum.AddonApplicability          `json:"addon_applicability"`
	TaxCode                     string                               `json:"tax_code"`
	HsnCode                     string                               `json:"hsn_code"`
	TaxjarProductCode           string                               `json:"taxjar_product_code"`
	AvalaraSaleType             enum.AvalaraSaleType                 `json:"avalara_sale_type"`
	AvalaraTransactionType      int32                                `json:"avalara_transaction_type"`
	AvalaraServiceType          int32                                `json:"avalara_service_type"`
	Sku                         string                               `json:"sku"`
	AccountingCode              string                               `json:"accounting_code"`
	AccountingCategory1         string                               `json:"accounting_category1"`
	AccountingCategory2         string                               `json:"accounting_category2"`
	AccountingCategory3         string                               `json:"accounting_category3"`
	AccountingCategory4         string                               `json:"accounting_category4"`
	IsShippable                 bool                                 `json:"is_shippable"`
	ShippingFrequencyPeriod     int32                                `json:"shipping_frequency_period"`
	ShippingFrequencyPeriodUnit planEnum.ShippingFrequencyPeriodUnit `json:"shipping_frequency_period_unit"`
	ResourceVersion             int64                                `json:"resource_version"`
	UpdatedAt                   int64                                `json:"updated_at"`
	Giftable                    bool                                 `json:"giftable"`
	ClaimUrl                    string                               `json:"claim_url"`
	FreeQuantityInDecimal       string                               `json:"free_quantity_in_decimal"`
	PriceInDecimal              string                               `json:"price_in_decimal"`
	InvoiceNotes                string                               `json:"invoice_notes"`
	Taxable                     bool                                 `json:"taxable"`
	TaxProfileId                string                               `json:"tax_profile_id"`
	MetaData                    json.RawMessage                      `json:"meta_data"`
	Tiers                       []*Tier                              `json:"tiers"`
	ApplicableAddons            []*ApplicableAddon                   `json:"applicable_addons"`
	AttachedAddons              []*AttachedAddon                     `json:"attached_addons"`
	EventBasedAddons            []*EventBasedAddon                   `json:"event_based_addons"`
	ShowDescriptionInInvoices   bool                                 `json:"show_description_in_invoices"`
	ShowDescriptionInQuotes     bool                                 `json:"show_description_in_quotes"`
	CustomField                 map[string]interface{}               `json:"custom_field"`
	Object                      string                               `json:"object"`
}
type Tier struct {
	StartingUnit          int32  `json:"starting_unit"`
	EndingUnit            int32  `json:"ending_unit"`
	Price                 int32  `json:"price"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal"`
	PriceInDecimal        string `json:"price_in_decimal"`
	Object                string `json:"object"`
}
type ApplicableAddon struct {
	Id     string `json:"id"`
	Object string `json:"object"`
}
type AttachedAddon struct {
	Id                string                     `json:"id"`
	Quantity          int32                      `json:"quantity"`
	BillingCycles     int32                      `json:"billing_cycles"`
	Type              planEnum.AttachedAddonType `json:"type"`
	QuantityInDecimal string                     `json:"quantity_in_decimal"`
	Object            string                     `json:"object"`
}
type EventBasedAddon struct {
	Id                string       `json:"id"`
	Quantity          int32        `json:"quantity"`
	OnEvent           enum.OnEvent `json:"on_event"`
	ChargeOnce        bool         `json:"charge_once"`
	QuantityInDecimal string       `json:"quantity_in_decimal"`
	Object            string       `json:"object"`
}
type CreateRequestParams struct {
	Id                          string                               `json:"id"`
	Name                        string                               `json:"name"`
	InvoiceName                 string                               `json:"invoice_name,omitempty"`
	Description                 string                               `json:"description,omitempty"`
	TrialPeriod                 *int32                               `json:"trial_period,omitempty"`
	TrialPeriodUnit             planEnum.TrialPeriodUnit             `json:"trial_period_unit,omitempty"`
	TrialEndAction              planEnum.TrialEndAction              `json:"trial_end_action,omitempty"`
	Period                      *int32                               `json:"period,omitempty"`
	PeriodUnit                  planEnum.PeriodUnit                  `json:"period_unit,omitempty"`
	SetupCost                   *int32                               `json:"setup_cost,omitempty"`
	Price                       *int32                               `json:"price,omitempty"`
	PriceInDecimal              string                               `json:"price_in_decimal,omitempty"`
	Tiers                       []*CreateTierParams                  `json:"tiers,omitempty"`
	CurrencyCode                string                               `json:"currency_code,omitempty"`
	BillingCycles               *int32                               `json:"billing_cycles,omitempty"`
	PricingModel                enum.PricingModel                    `json:"pricing_model,omitempty"`
	ChargeModel                 planEnum.ChargeModel                 `json:"charge_model,omitempty"`
	FreeQuantity                *int32                               `json:"free_quantity,omitempty"`
	FreeQuantityInDecimal       string                               `json:"free_quantity_in_decimal,omitempty"`
	AddonApplicability          planEnum.AddonApplicability          `json:"addon_applicability,omitempty"`
	DowngradePenalty            *float64                             `json:"downgrade_penalty,omitempty"`
	RedirectUrl                 string                               `json:"redirect_url,omitempty"`
	EnabledInHostedPages        *bool                                `json:"enabled_in_hosted_pages,omitempty"`
	EnabledInPortal             *bool                                `json:"enabled_in_portal,omitempty"`
	Taxable                     *bool                                `json:"taxable,omitempty"`
	TaxProfileId                string                               `json:"tax_profile_id,omitempty"`
	TaxCode                     string                               `json:"tax_code,omitempty"`
	HsnCode                     string                               `json:"hsn_code,omitempty"`
	TaxjarProductCode           string                               `json:"taxjar_product_code,omitempty"`
	AvalaraSaleType             enum.AvalaraSaleType                 `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType      *int32                               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType          *int32                               `json:"avalara_service_type,omitempty"`
	Sku                         string                               `json:"sku,omitempty"`
	AccountingCode              string                               `json:"accounting_code,omitempty"`
	AccountingCategory1         string                               `json:"accounting_category1,omitempty"`
	AccountingCategory2         string                               `json:"accounting_category2,omitempty"`
	AccountingCategory3         string                               `json:"accounting_category3,omitempty"`
	AccountingCategory4         string                               `json:"accounting_category4,omitempty"`
	IsShippable                 *bool                                `json:"is_shippable,omitempty"`
	ShippingFrequencyPeriod     *int32                               `json:"shipping_frequency_period,omitempty"`
	ShippingFrequencyPeriodUnit planEnum.ShippingFrequencyPeriodUnit `json:"shipping_frequency_period_unit,omitempty"`
	ApplicableAddons            []*CreateApplicableAddonParams       `json:"applicable_addons,omitempty"`
	EventBasedAddons            []*CreateEventBasedAddonParams       `json:"event_based_addons,omitempty"`
	AttachedAddons              []*CreateAttachedAddonParams         `json:"attached_addons,omitempty"`
	InvoiceNotes                string                               `json:"invoice_notes,omitempty"`
	MetaData                    map[string]interface{}               `json:"meta_data,omitempty"`
	ShowDescriptionInInvoices   *bool                                `json:"show_description_in_invoices,omitempty"`
	ShowDescriptionInQuotes     *bool                                `json:"show_description_in_quotes,omitempty"`
	Giftable                    *bool                                `json:"giftable,omitempty"`
	Status                      planEnum.Status                      `json:"status,omitempty"`
	ClaimUrl                    string                               `json:"claim_url,omitempty"`
}
type CreateTierParams struct {
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type CreateApplicableAddonParams struct {
	Id string `json:"id,omitempty"`
}
type CreateEventBasedAddonParams struct {
	Id                string       `json:"id,omitempty"`
	Quantity          *int32       `json:"quantity,omitempty"`
	QuantityInDecimal string       `json:"quantity_in_decimal,omitempty"`
	OnEvent           enum.OnEvent `json:"on_event,omitempty"`
	ChargeOnce        *bool        `json:"charge_once,omitempty"`
}
type CreateAttachedAddonParams struct {
	Id                string                     `json:"id,omitempty"`
	Quantity          *int32                     `json:"quantity,omitempty"`
	QuantityInDecimal string                     `json:"quantity_in_decimal,omitempty"`
	BillingCycles     *int32                     `json:"billing_cycles,omitempty"`
	Type              planEnum.AttachedAddonType `json:"type,omitempty"`
}
type UpdateRequestParams struct {
	Name                        string                               `json:"name,omitempty"`
	InvoiceName                 string                               `json:"invoice_name,omitempty"`
	Description                 string                               `json:"description,omitempty"`
	TrialPeriod                 *int32                               `json:"trial_period,omitempty"`
	TrialPeriodUnit             planEnum.TrialPeriodUnit             `json:"trial_period_unit,omitempty"`
	TrialEndAction              planEnum.TrialEndAction              `json:"trial_end_action,omitempty"`
	Period                      *int32                               `json:"period,omitempty"`
	PeriodUnit                  planEnum.PeriodUnit                  `json:"period_unit,omitempty"`
	SetupCost                   *int32                               `json:"setup_cost,omitempty"`
	Price                       *int32                               `json:"price,omitempty"`
	PriceInDecimal              string                               `json:"price_in_decimal,omitempty"`
	Tiers                       []*UpdateTierParams                  `json:"tiers,omitempty"`
	CurrencyCode                string                               `json:"currency_code,omitempty"`
	BillingCycles               *int32                               `json:"billing_cycles,omitempty"`
	PricingModel                enum.PricingModel                    `json:"pricing_model,omitempty"`
	ChargeModel                 planEnum.ChargeModel                 `json:"charge_model,omitempty"`
	FreeQuantity                *int32                               `json:"free_quantity,omitempty"`
	FreeQuantityInDecimal       string                               `json:"free_quantity_in_decimal,omitempty"`
	AddonApplicability          planEnum.AddonApplicability          `json:"addon_applicability,omitempty"`
	DowngradePenalty            *float64                             `json:"downgrade_penalty,omitempty"`
	RedirectUrl                 string                               `json:"redirect_url,omitempty"`
	EnabledInHostedPages        *bool                                `json:"enabled_in_hosted_pages,omitempty"`
	EnabledInPortal             *bool                                `json:"enabled_in_portal,omitempty"`
	Taxable                     *bool                                `json:"taxable,omitempty"`
	TaxProfileId                string                               `json:"tax_profile_id,omitempty"`
	TaxCode                     string                               `json:"tax_code,omitempty"`
	HsnCode                     string                               `json:"hsn_code,omitempty"`
	TaxjarProductCode           string                               `json:"taxjar_product_code,omitempty"`
	AvalaraSaleType             enum.AvalaraSaleType                 `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType      *int32                               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType          *int32                               `json:"avalara_service_type,omitempty"`
	Sku                         string                               `json:"sku,omitempty"`
	AccountingCode              string                               `json:"accounting_code,omitempty"`
	AccountingCategory1         string                               `json:"accounting_category1,omitempty"`
	AccountingCategory2         string                               `json:"accounting_category2,omitempty"`
	AccountingCategory3         string                               `json:"accounting_category3,omitempty"`
	AccountingCategory4         string                               `json:"accounting_category4,omitempty"`
	IsShippable                 *bool                                `json:"is_shippable,omitempty"`
	ShippingFrequencyPeriod     *int32                               `json:"shipping_frequency_period,omitempty"`
	ShippingFrequencyPeriodUnit planEnum.ShippingFrequencyPeriodUnit `json:"shipping_frequency_period_unit,omitempty"`
	ApplicableAddons            []*UpdateApplicableAddonParams       `json:"applicable_addons,omitempty"`
	EventBasedAddons            []*UpdateEventBasedAddonParams       `json:"event_based_addons,omitempty"`
	AttachedAddons              []*UpdateAttachedAddonParams         `json:"attached_addons,omitempty"`
	InvoiceNotes                string                               `json:"invoice_notes,omitempty"`
	MetaData                    map[string]interface{}               `json:"meta_data,omitempty"`
	ShowDescriptionInInvoices   *bool                                `json:"show_description_in_invoices,omitempty"`
	ShowDescriptionInQuotes     *bool                                `json:"show_description_in_quotes,omitempty"`
}
type UpdateTierParams struct {
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type UpdateApplicableAddonParams struct {
	Id string `json:"id,omitempty"`
}
type UpdateEventBasedAddonParams struct {
	Id                string       `json:"id,omitempty"`
	Quantity          *int32       `json:"quantity,omitempty"`
	QuantityInDecimal string       `json:"quantity_in_decimal,omitempty"`
	OnEvent           enum.OnEvent `json:"on_event,omitempty"`
	ChargeOnce        *bool        `json:"charge_once,omitempty"`
}
type UpdateAttachedAddonParams struct {
	Id                string                     `json:"id,omitempty"`
	Quantity          *int32                     `json:"quantity,omitempty"`
	QuantityInDecimal string                     `json:"quantity_in_decimal,omitempty"`
	BillingCycles     *int32                     `json:"billing_cycles,omitempty"`
	Type              planEnum.AttachedAddonType `json:"type,omitempty"`
}
type ListRequestParams struct {
	Limit              *int32                  `json:"limit,omitempty"`
	Offset             string                  `json:"offset,omitempty"`
	Id                 *filter.StringFilter    `json:"id,omitempty"`
	Name               *filter.StringFilter    `json:"name,omitempty"`
	Price              *filter.NumberFilter    `json:"price,omitempty"`
	Period             *filter.NumberFilter    `json:"period,omitempty"`
	PeriodUnit         *filter.EnumFilter      `json:"period_unit,omitempty"`
	TrialPeriod        *filter.NumberFilter    `json:"trial_period,omitempty"`
	TrialPeriodUnit    *filter.EnumFilter      `json:"trial_period_unit,omitempty"`
	AddonApplicability *filter.EnumFilter      `json:"addon_applicability,omitempty"`
	Giftable           *filter.BooleanFilter   `json:"giftable,omitempty"`
	ChargeModel        *filter.EnumFilter      `json:"charge_model,omitempty"`
	PricingModel       *filter.EnumFilter      `json:"pricing_model,omitempty"`
	Status             *filter.EnumFilter      `json:"status,omitempty"`
	UpdatedAt          *filter.TimestampFilter `json:"updated_at,omitempty"`
	CurrencyCode       *filter.StringFilter    `json:"currency_code,omitempty"`
	IncludeDeleted     *bool                   `json:"include_deleted,omitempty"`
}
type CopyRequestParams struct {
	FromSite       string `json:"from_site"`
	IdAtFromSite   string `json:"id_at_from_site"`
	Id             string `json:"id,omitempty"`
	ForSiteMerging *bool  `json:"for_site_merging,omitempty"`
}
