package chargebee

import (
	"encoding/json"
)

type PlanPeriodUnit string

const (
	PlanPeriodUnitDay   PlanPeriodUnit = "day"
	PlanPeriodUnitWeek  PlanPeriodUnit = "week"
	PlanPeriodUnitMonth PlanPeriodUnit = "month"
	PlanPeriodUnitYear  PlanPeriodUnit = "year"
)

type PlanTrialPeriodUnit string

const (
	PlanTrialPeriodUnitDay   PlanTrialPeriodUnit = "day"
	PlanTrialPeriodUnitMonth PlanTrialPeriodUnit = "month"
)

type PlanTrialEndAction string

const (
	PlanTrialEndActionSiteDefault          PlanTrialEndAction = "site_default"
	PlanTrialEndActionActivateSubscription PlanTrialEndAction = "activate_subscription"
	PlanTrialEndActionCancelSubscription   PlanTrialEndAction = "cancel_subscription"
)

type PlanPricingModel string

const (
	PlanPricingModelFlatFee   PlanPricingModel = "flat_fee"
	PlanPricingModelPerUnit   PlanPricingModel = "per_unit"
	PlanPricingModelTiered    PlanPricingModel = "tiered"
	PlanPricingModelVolume    PlanPricingModel = "volume"
	PlanPricingModelStairstep PlanPricingModel = "stairstep"
)

type PlanChargeModel string

const (
	PlanChargeModelFlatFee   PlanChargeModel = "flat_fee"
	PlanChargeModelPerUnit   PlanChargeModel = "per_unit"
	PlanChargeModelTiered    PlanChargeModel = "tiered"
	PlanChargeModelVolume    PlanChargeModel = "volume"
	PlanChargeModelStairstep PlanChargeModel = "stairstep"
)

type PlanStatus string

const (
	PlanStatusActive   PlanStatus = "active"
	PlanStatusArchived PlanStatus = "archived"
	PlanStatusDeleted  PlanStatus = "deleted"
)

type PlanAddonApplicability string

const (
	PlanAddonApplicabilityAll        PlanAddonApplicability = "all"
	PlanAddonApplicabilityRestricted PlanAddonApplicability = "restricted"
)

type PlanAvalaraSaleType string

const (
	PlanAvalaraSaleTypeWholesale PlanAvalaraSaleType = "wholesale"
	PlanAvalaraSaleTypeRetail    PlanAvalaraSaleType = "retail"
	PlanAvalaraSaleTypeConsumed  PlanAvalaraSaleType = "consumed"
	PlanAvalaraSaleTypeVendorUse PlanAvalaraSaleType = "vendor_use"
)

type PlanShippingFrequencyPeriodUnit string

const (
	PlanShippingFrequencyPeriodUnitYear  PlanShippingFrequencyPeriodUnit = "year"
	PlanShippingFrequencyPeriodUnitMonth PlanShippingFrequencyPeriodUnit = "month"
	PlanShippingFrequencyPeriodUnitWeek  PlanShippingFrequencyPeriodUnit = "week"
	PlanShippingFrequencyPeriodUnitDay   PlanShippingFrequencyPeriodUnit = "day"
)

type PlanChannel string

const (
	PlanChannelWeb       PlanChannel = "web"
	PlanChannelAppStore  PlanChannel = "app_store"
	PlanChannelPlayStore PlanChannel = "play_store"
)

type PlanTierPricingType string

const (
	PlanTierPricingTypePerUnit PlanTierPricingType = "per_unit"
	PlanTierPricingTypeFlatFee PlanTierPricingType = "flat_fee"
	PlanTierPricingTypePackage PlanTierPricingType = "package"
)

type PlanAttachedAddonType string

const (
	PlanAttachedAddonTypeRecommended PlanAttachedAddonType = "recommended"
	PlanAttachedAddonTypeMandatory   PlanAttachedAddonType = "mandatory"
)

type PlanEventBasedAddonOnEvent string

const (
	PlanEventBasedAddonOnEventSubscriptionCreation   PlanEventBasedAddonOnEvent = "subscription_creation"
	PlanEventBasedAddonOnEventSubscriptionTrialStart PlanEventBasedAddonOnEvent = "subscription_trial_start"
	PlanEventBasedAddonOnEventPlanActivation         PlanEventBasedAddonOnEvent = "plan_activation"
	PlanEventBasedAddonOnEventSubscriptionActivation PlanEventBasedAddonOnEvent = "subscription_activation"
	PlanEventBasedAddonOnEventContractTermination    PlanEventBasedAddonOnEvent = "contract_termination"
)

// just struct
type Plan struct {
	Id                          string                          `json:"id"`
	Name                        string                          `json:"name"`
	InvoiceName                 string                          `json:"invoice_name"`
	Description                 string                          `json:"description"`
	Price                       int64                           `json:"price"`
	CurrencyCode                string                          `json:"currency_code"`
	Period                      int32                           `json:"period"`
	PeriodUnit                  PlanPeriodUnit                  `json:"period_unit"`
	TrialPeriod                 int32                           `json:"trial_period"`
	TrialPeriodUnit             PlanTrialPeriodUnit             `json:"trial_period_unit"`
	TrialEndAction              PlanTrialEndAction              `json:"trial_end_action"`
	PricingModel                PlanPricingModel                `json:"pricing_model"`
	FreeQuantity                int32                           `json:"free_quantity"`
	SetupCost                   int64                           `json:"setup_cost"`
	Status                      PlanStatus                      `json:"status"`
	ArchivedAt                  int64                           `json:"archived_at"`
	BillingCycles               int32                           `json:"billing_cycles"`
	RedirectUrl                 string                          `json:"redirect_url"`
	EnabledInHostedPages        bool                            `json:"enabled_in_hosted_pages"`
	EnabledInPortal             bool                            `json:"enabled_in_portal"`
	AddonApplicability          PlanAddonApplicability          `json:"addon_applicability"`
	TaxCode                     string                          `json:"tax_code"`
	HsnCode                     string                          `json:"hsn_code"`
	TaxjarProductCode           string                          `json:"taxjar_product_code"`
	AvalaraSaleType             PlanAvalaraSaleType             `json:"avalara_sale_type"`
	AvalaraTransactionType      int32                           `json:"avalara_transaction_type"`
	AvalaraServiceType          int32                           `json:"avalara_service_type"`
	Sku                         string                          `json:"sku"`
	AccountingCode              string                          `json:"accounting_code"`
	AccountingCategory1         string                          `json:"accounting_category1"`
	AccountingCategory2         string                          `json:"accounting_category2"`
	AccountingCategory3         string                          `json:"accounting_category3"`
	AccountingCategory4         string                          `json:"accounting_category4"`
	IsShippable                 bool                            `json:"is_shippable"`
	ShippingFrequencyPeriod     int32                           `json:"shipping_frequency_period"`
	ShippingFrequencyPeriodUnit PlanShippingFrequencyPeriodUnit `json:"shipping_frequency_period_unit"`
	ResourceVersion             int64                           `json:"resource_version"`
	UpdatedAt                   int64                           `json:"updated_at"`
	Giftable                    bool                            `json:"giftable"`
	ClaimUrl                    string                          `json:"claim_url"`
	FreeQuantityInDecimal       string                          `json:"free_quantity_in_decimal"`
	PriceInDecimal              string                          `json:"price_in_decimal"`
	Channel                     PlanChannel                     `json:"channel"`
	InvoiceNotes                string                          `json:"invoice_notes"`
	Taxable                     bool                            `json:"taxable"`
	TaxProfileId                string                          `json:"tax_profile_id"`
	MetaData                    json.RawMessage                 `json:"meta_data"`
	Tiers                       []*PlanTier                     `json:"tiers"`
	TaxProvidersFields          []*PlanTaxProvidersField        `json:"tax_providers_fields"`
	ApplicableAddons            []*PlanApplicableAddon          `json:"applicable_addons"`
	AttachedAddons              []*PlanAttachedAddon            `json:"attached_addons"`
	EventBasedAddons            []*PlanEventBasedAddon          `json:"event_based_addons"`
	ShowDescriptionInInvoices   bool                            `json:"show_description_in_invoices"`
	ShowDescriptionInQuotes     bool                            `json:"show_description_in_quotes"`
	CustomField                 CustomField                     `json:"custom_field"`
	Object                      string                          `json:"object"`
}

// sub resources
type PlanTier struct {
	StartingUnit          int32               `json:"starting_unit"`
	EndingUnit            int32               `json:"ending_unit"`
	Price                 int64               `json:"price"`
	StartingUnitInDecimal string              `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string              `json:"ending_unit_in_decimal"`
	PriceInDecimal        string              `json:"price_in_decimal"`
	PricingType           PlanTierPricingType `json:"pricing_type"`
	PackageSize           int32               `json:"package_size"`
	Object                string              `json:"object"`
}
type PlanTaxProvidersField struct {
	ProviderName string `json:"provider_name"`
	FieldId      string `json:"field_id"`
	FieldValue   string `json:"field_value"`
	Object       string `json:"object"`
}
type PlanApplicableAddon struct {
	Id     string `json:"id"`
	Object string `json:"object"`
}
type PlanAttachedAddon struct {
	Id                string                `json:"id"`
	Quantity          int32                 `json:"quantity"`
	BillingCycles     int32                 `json:"billing_cycles"`
	Type              PlanAttachedAddonType `json:"type"`
	QuantityInDecimal string                `json:"quantity_in_decimal"`
	Object            string                `json:"object"`
}
type PlanEventBasedAddon struct {
	Id                string                     `json:"id"`
	Quantity          int32                      `json:"quantity"`
	OnEvent           PlanEventBasedAddonOnEvent `json:"on_event"`
	ChargeOnce        bool                       `json:"charge_once"`
	QuantityInDecimal string                     `json:"quantity_in_decimal"`
	Object            string                     `json:"object"`
}

// operations
// input params
type PlanCreateRequest struct {
	Id                          string                          `json:"id"`
	Name                        string                          `json:"name"`
	InvoiceName                 string                          `json:"invoice_name,omitempty"`
	Description                 string                          `json:"description,omitempty"`
	TrialPeriod                 *int32                          `json:"trial_period,omitempty"`
	TrialPeriodUnit             PlanTrialPeriodUnit             `json:"trial_period_unit,omitempty"`
	TrialEndAction              PlanTrialEndAction              `json:"trial_end_action,omitempty"`
	Period                      *int32                          `json:"period,omitempty"`
	PeriodUnit                  PlanPeriodUnit                  `json:"period_unit,omitempty"`
	SetupCost                   *int64                          `json:"setup_cost,omitempty"`
	Price                       *int64                          `json:"price,omitempty"`
	PriceInDecimal              string                          `json:"price_in_decimal,omitempty"`
	Tiers                       []*PlanCreateTier               `json:"tiers,omitempty"`
	CurrencyCode                string                          `json:"currency_code,omitempty"`
	BillingCycles               *int32                          `json:"billing_cycles,omitempty"`
	PricingModel                PlanPricingModel                `json:"pricing_model,omitempty"`
	FreeQuantity                *int32                          `json:"free_quantity,omitempty"`
	FreeQuantityInDecimal       string                          `json:"free_quantity_in_decimal,omitempty"`
	AddonApplicability          PlanAddonApplicability          `json:"addon_applicability,omitempty"`
	RedirectUrl                 string                          `json:"redirect_url,omitempty"`
	EnabledInHostedPages        *bool                           `json:"enabled_in_hosted_pages,omitempty"`
	EnabledInPortal             *bool                           `json:"enabled_in_portal,omitempty"`
	Taxable                     *bool                           `json:"taxable,omitempty"`
	TaxProfileId                string                          `json:"tax_profile_id,omitempty"`
	TaxCode                     string                          `json:"tax_code,omitempty"`
	HsnCode                     string                          `json:"hsn_code,omitempty"`
	TaxjarProductCode           string                          `json:"taxjar_product_code,omitempty"`
	AvalaraSaleType             PlanAvalaraSaleType             `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType      *int32                          `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType          *int32                          `json:"avalara_service_type,omitempty"`
	Sku                         string                          `json:"sku,omitempty"`
	AccountingCode              string                          `json:"accounting_code,omitempty"`
	AccountingCategory1         string                          `json:"accounting_category1,omitempty"`
	AccountingCategory2         string                          `json:"accounting_category2,omitempty"`
	AccountingCategory3         string                          `json:"accounting_category3,omitempty"`
	AccountingCategory4         string                          `json:"accounting_category4,omitempty"`
	IsShippable                 *bool                           `json:"is_shippable,omitempty"`
	ShippingFrequencyPeriod     *int32                          `json:"shipping_frequency_period,omitempty"`
	ShippingFrequencyPeriodUnit PlanShippingFrequencyPeriodUnit `json:"shipping_frequency_period_unit,omitempty"`
	TaxProvidersFields          []*PlanCreateTaxProvidersField  `json:"tax_providers_fields,omitempty"`
	ApplicableAddons            []*PlanCreateApplicableAddon    `json:"applicable_addons,omitempty"`
	EventBasedAddons            []*PlanCreateEventBasedAddon    `json:"event_based_addons,omitempty"`
	AttachedAddons              []*PlanCreateAttachedAddon      `json:"attached_addons,omitempty"`
	InvoiceNotes                string                          `json:"invoice_notes,omitempty"`
	MetaData                    map[string]interface{}          `json:"meta_data,omitempty"`
	ShowDescriptionInInvoices   *bool                           `json:"show_description_in_invoices,omitempty"`
	ShowDescriptionInQuotes     *bool                           `json:"show_description_in_quotes,omitempty"`
	Giftable                    *bool                           `json:"giftable,omitempty"`
	Status                      PlanStatus                      `json:"status,omitempty"`
	ClaimUrl                    string                          `json:"claim_url,omitempty"`
	apiRequest                  `json:"-" form:"-"`
}

func (r *PlanCreateRequest) payload() any { return r }

// input sub resource params multi
type PlanCreateTier struct {
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int64 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}

// input sub resource params multi
type PlanCreateTaxProvidersField struct {
	ProviderName string `json:"provider_name"`
	FieldId      string `json:"field_id"`
	FieldValue   string `json:"field_value"`
}

// input sub resource params multi
type PlanCreateApplicableAddon struct {
	Id string `json:"id,omitempty"`
}

// input sub resource params multi
type PlanCreateEventBasedAddon struct {
	Id                string  `json:"id,omitempty"`
	Quantity          *int32  `json:"quantity,omitempty"`
	QuantityInDecimal string  `json:"quantity_in_decimal,omitempty"`
	OnEvent           OnEvent `json:"on_event,omitempty"`
	ChargeOnce        *bool   `json:"charge_once,omitempty"`
}

// input sub resource params multi
type PlanCreateAttachedAddon struct {
	Id                string            `json:"id,omitempty"`
	Quantity          *int32            `json:"quantity,omitempty"`
	QuantityInDecimal string            `json:"quantity_in_decimal,omitempty"`
	BillingCycles     *int32            `json:"billing_cycles,omitempty"`
	Type              AttachedAddonType `json:"type,omitempty"`
}
type PlanUpdateRequest struct {
	Name                        string                          `json:"name,omitempty"`
	InvoiceName                 string                          `json:"invoice_name,omitempty"`
	Description                 string                          `json:"description,omitempty"`
	TrialPeriod                 *int32                          `json:"trial_period,omitempty"`
	TrialPeriodUnit             PlanTrialPeriodUnit             `json:"trial_period_unit,omitempty"`
	TrialEndAction              PlanTrialEndAction              `json:"trial_end_action,omitempty"`
	Period                      *int32                          `json:"period,omitempty"`
	PeriodUnit                  PlanPeriodUnit                  `json:"period_unit,omitempty"`
	SetupCost                   *int64                          `json:"setup_cost,omitempty"`
	Price                       *int64                          `json:"price,omitempty"`
	PriceInDecimal              string                          `json:"price_in_decimal,omitempty"`
	Tiers                       []*PlanUpdateTier               `json:"tiers,omitempty"`
	CurrencyCode                string                          `json:"currency_code,omitempty"`
	BillingCycles               *int32                          `json:"billing_cycles,omitempty"`
	PricingModel                PlanPricingModel                `json:"pricing_model,omitempty"`
	FreeQuantity                *int32                          `json:"free_quantity,omitempty"`
	FreeQuantityInDecimal       string                          `json:"free_quantity_in_decimal,omitempty"`
	AddonApplicability          PlanAddonApplicability          `json:"addon_applicability,omitempty"`
	RedirectUrl                 string                          `json:"redirect_url,omitempty"`
	EnabledInHostedPages        *bool                           `json:"enabled_in_hosted_pages,omitempty"`
	EnabledInPortal             *bool                           `json:"enabled_in_portal,omitempty"`
	Taxable                     *bool                           `json:"taxable,omitempty"`
	TaxProfileId                string                          `json:"tax_profile_id,omitempty"`
	TaxCode                     string                          `json:"tax_code,omitempty"`
	HsnCode                     string                          `json:"hsn_code,omitempty"`
	TaxjarProductCode           string                          `json:"taxjar_product_code,omitempty"`
	AvalaraSaleType             PlanAvalaraSaleType             `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType      *int32                          `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType          *int32                          `json:"avalara_service_type,omitempty"`
	Sku                         string                          `json:"sku,omitempty"`
	AccountingCode              string                          `json:"accounting_code,omitempty"`
	AccountingCategory1         string                          `json:"accounting_category1,omitempty"`
	AccountingCategory2         string                          `json:"accounting_category2,omitempty"`
	AccountingCategory3         string                          `json:"accounting_category3,omitempty"`
	AccountingCategory4         string                          `json:"accounting_category4,omitempty"`
	IsShippable                 *bool                           `json:"is_shippable,omitempty"`
	ShippingFrequencyPeriod     *int32                          `json:"shipping_frequency_period,omitempty"`
	ShippingFrequencyPeriodUnit PlanShippingFrequencyPeriodUnit `json:"shipping_frequency_period_unit,omitempty"`
	TaxProvidersFields          []*PlanUpdateTaxProvidersField  `json:"tax_providers_fields,omitempty"`
	ApplicableAddons            []*PlanUpdateApplicableAddon    `json:"applicable_addons,omitempty"`
	EventBasedAddons            []*PlanUpdateEventBasedAddon    `json:"event_based_addons,omitempty"`
	AttachedAddons              []*PlanUpdateAttachedAddon      `json:"attached_addons,omitempty"`
	InvoiceNotes                string                          `json:"invoice_notes,omitempty"`
	MetaData                    map[string]interface{}          `json:"meta_data,omitempty"`
	ShowDescriptionInInvoices   *bool                           `json:"show_description_in_invoices,omitempty"`
	ShowDescriptionInQuotes     *bool                           `json:"show_description_in_quotes,omitempty"`
	apiRequest                  `json:"-" form:"-"`
}

func (r *PlanUpdateRequest) payload() any { return r }

// input sub resource params multi
type PlanUpdateTier struct {
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int64 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}

// input sub resource params multi
type PlanUpdateTaxProvidersField struct {
	ProviderName string `json:"provider_name"`
	FieldId      string `json:"field_id"`
	FieldValue   string `json:"field_value"`
}

// input sub resource params multi
type PlanUpdateApplicableAddon struct {
	Id string `json:"id,omitempty"`
}

// input sub resource params multi
type PlanUpdateEventBasedAddon struct {
	Id                string  `json:"id,omitempty"`
	Quantity          *int32  `json:"quantity,omitempty"`
	QuantityInDecimal string  `json:"quantity_in_decimal,omitempty"`
	OnEvent           OnEvent `json:"on_event,omitempty"`
	ChargeOnce        *bool   `json:"charge_once,omitempty"`
}

// input sub resource params multi
type PlanUpdateAttachedAddon struct {
	Id                string            `json:"id,omitempty"`
	Quantity          *int32            `json:"quantity,omitempty"`
	QuantityInDecimal string            `json:"quantity_in_decimal,omitempty"`
	BillingCycles     *int32            `json:"billing_cycles,omitempty"`
	Type              AttachedAddonType `json:"type,omitempty"`
}
type PlanListRequest struct {
	Limit              *int32           `json:"limit,omitempty"`
	Offset             string           `json:"offset,omitempty"`
	Id                 *StringFilter    `json:"id,omitempty"`
	Name               *StringFilter    `json:"name,omitempty"`
	Price              *NumberFilter    `json:"price,omitempty"`
	Period             *NumberFilter    `json:"period,omitempty"`
	PeriodUnit         *EnumFilter      `json:"period_unit,omitempty"`
	TrialPeriod        *NumberFilter    `json:"trial_period,omitempty"`
	TrialPeriodUnit    *EnumFilter      `json:"trial_period_unit,omitempty"`
	AddonApplicability *EnumFilter      `json:"addon_applicability,omitempty"`
	Giftable           *BooleanFilter   `json:"giftable,omitempty"`
	PricingModel       *EnumFilter      `json:"pricing_model,omitempty"`
	Status             *EnumFilter      `json:"status,omitempty"`
	UpdatedAt          *TimestampFilter `json:"updated_at,omitempty"`
	CurrencyCode       *StringFilter    `json:"currency_code,omitempty"`
	Channel            *EnumFilter      `json:"channel,omitempty"`
	IncludeDeleted     *bool            `json:"include_deleted,omitempty"`
	apiRequest         `json:"-" form:"-"`
}

func (r *PlanListRequest) payload() any { return r }

type PlanCopyRequest struct {
	FromSite       string `json:"from_site"`
	IdAtFromSite   string `json:"id_at_from_site"`
	Id             string `json:"id,omitempty"`
	ForSiteMerging *bool  `json:"for_site_merging,omitempty"`
	apiRequest     `json:"-" form:"-"`
}

func (r *PlanCopyRequest) payload() any { return r }

// operation response
type PlanCreateResponse struct {
	Plan *Plan `json:"plan,omitempty"`
	apiResponse
}

// operation response
type PlanUpdateResponse struct {
	Plan *Plan `json:"plan,omitempty"`
	apiResponse
}

// operation sub response
type PlanListPlanResponse struct {
	Plan *Plan `json:"plan,omitempty"`
}

// operation response
type PlanListResponse struct {
	List       []*PlanListPlanResponse `json:"list,omitempty"`
	NextOffset string                  `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type PlanRetrieveResponse struct {
	Plan *Plan `json:"plan,omitempty"`
	apiResponse
}

// operation response
type PlanDeleteResponse struct {
	Plan *Plan `json:"plan,omitempty"`
	apiResponse
}

// operation response
type PlanCopyResponse struct {
	Plan *Plan `json:"plan,omitempty"`
	apiResponse
}

// operation response
type PlanUnarchiveResponse struct {
	Plan *Plan `json:"plan,omitempty"`
	apiResponse
}
