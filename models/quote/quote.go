package quote

import (
	"encoding/json"
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	quoteEnum "github.com/chargebee/chargebee-go/models/quote/enum"
)

type Quote struct {
	Id                string                  `json:"id"`
	Name              string                  `json:"name"`
	PoNumber          string                  `json:"po_number"`
	CustomerId        string                  `json:"customer_id"`
	SubscriptionId    string                  `json:"subscription_id"`
	InvoiceId         string                  `json:"invoice_id"`
	Status            quoteEnum.Status        `json:"status"`
	OperationType     quoteEnum.OperationType `json:"operation_type"`
	VatNumber         string                  `json:"vat_number"`
	PriceType         enum.PriceType          `json:"price_type"`
	ValidTill         int64                   `json:"valid_till"`
	Date              int64                   `json:"date"`
	SubTotal          int32                   `json:"sub_total"`
	Total             int32                   `json:"total"`
	CreditsApplied    int32                   `json:"credits_applied"`
	AmountPaid        int32                   `json:"amount_paid"`
	AmountDue         int32                   `json:"amount_due"`
	ResourceVersion   int64                   `json:"resource_version"`
	UpdatedAt         int64                   `json:"updated_at"`
	CurrencyCode      string                  `json:"currency_code"`
	LineItems         []*LineItem             `json:"line_items"`
	Discounts         []*Discount             `json:"discounts"`
	LineItemDiscounts []*LineItemDiscount     `json:"line_item_discounts"`
	Taxes             []*Tax                  `json:"taxes"`
	LineItemTaxes     []*LineItemTax          `json:"line_item_taxes"`
	Notes             json.RawMessage         `json:"notes"`
	ShippingAddress   *ShippingAddress        `json:"shipping_address"`
	BillingAddress    *BillingAddress         `json:"billing_address"`
	Object            string                  `json:"object"`
}
type LineItem struct {
	Id                      string                       `json:"id"`
	SubscriptionId          string                       `json:"subscription_id"`
	DateFrom                int64                        `json:"date_from"`
	DateTo                  int64                        `json:"date_to"`
	UnitAmount              int32                        `json:"unit_amount"`
	Quantity                int32                        `json:"quantity"`
	Amount                  int32                        `json:"amount"`
	PricingModel            enum.PricingModel            `json:"pricing_model"`
	IsTaxed                 bool                         `json:"is_taxed"`
	TaxAmount               int32                        `json:"tax_amount"`
	TaxRate                 float64                      `json:"tax_rate"`
	DiscountAmount          int32                        `json:"discount_amount"`
	ItemLevelDiscountAmount int32                        `json:"item_level_discount_amount"`
	Description             string                       `json:"description"`
	EntityType              quoteEnum.LineItemEntityType `json:"entity_type"`
	TaxExemptReason         enum.TaxExemptReason         `json:"tax_exempt_reason"`
	EntityId                string                       `json:"entity_id"`
	CustomerId              string                       `json:"customer_id"`
	Object                  string                       `json:"object"`
}
type Discount struct {
	Amount      int32                        `json:"amount"`
	Description string                       `json:"description"`
	EntityType  quoteEnum.DiscountEntityType `json:"entity_type"`
	EntityId    string                       `json:"entity_id"`
	Object      string                       `json:"object"`
}
type LineItemDiscount struct {
	LineItemId     string                                 `json:"line_item_id"`
	DiscountType   quoteEnum.LineItemDiscountDiscountType `json:"discount_type"`
	CouponId       string                                 `json:"coupon_id"`
	DiscountAmount int32                                  `json:"discount_amount"`
	Object         string                                 `json:"object"`
}
type Tax struct {
	Name        string `json:"name"`
	Amount      int32  `json:"amount"`
	Description string `json:"description"`
	Object      string `json:"object"`
}
type LineItemTax struct {
	LineItemId               string            `json:"line_item_id"`
	TaxName                  string            `json:"tax_name"`
	TaxRate                  float64           `json:"tax_rate"`
	IsPartialTaxApplied      bool              `json:"is_partial_tax_applied"`
	IsNonComplianceTax       bool              `json:"is_non_compliance_tax"`
	TaxableAmount            int32             `json:"taxable_amount"`
	TaxAmount                int32             `json:"tax_amount"`
	TaxJurisType             enum.TaxJurisType `json:"tax_juris_type"`
	TaxJurisName             string            `json:"tax_juris_name"`
	TaxJurisCode             string            `json:"tax_juris_code"`
	TaxAmountInLocalCurrency int32             `json:"tax_amount_in_local_currency"`
	LocalCurrencyCode        string            `json:"local_currency_code"`
	Object                   string            `json:"object"`
}
type ShippingAddress struct {
	FirstName        string                `json:"first_name"`
	LastName         string                `json:"last_name"`
	Email            string                `json:"email"`
	Company          string                `json:"company"`
	Phone            string                `json:"phone"`
	Line1            string                `json:"line1"`
	Line2            string                `json:"line2"`
	Line3            string                `json:"line3"`
	City             string                `json:"city"`
	StateCode        string                `json:"state_code"`
	State            string                `json:"state"`
	Country          string                `json:"country"`
	Zip              string                `json:"zip"`
	ValidationStatus enum.ValidationStatus `json:"validation_status"`
	Object           string                `json:"object"`
}
type BillingAddress struct {
	FirstName        string                `json:"first_name"`
	LastName         string                `json:"last_name"`
	Email            string                `json:"email"`
	Company          string                `json:"company"`
	Phone            string                `json:"phone"`
	Line1            string                `json:"line1"`
	Line2            string                `json:"line2"`
	Line3            string                `json:"line3"`
	City             string                `json:"city"`
	StateCode        string                `json:"state_code"`
	State            string                `json:"state"`
	Country          string                `json:"country"`
	Zip              string                `json:"zip"`
	ValidationStatus enum.ValidationStatus `json:"validation_status"`
	Object           string                `json:"object"`
}
type CreateSubForCustomerQuoteRequestParams struct {
	Name                    string                                            `json:"name,omitempty"`
	Notes                   string                                            `json:"notes,omitempty"`
	ExpiresAt               *int64                                            `json:"expires_at,omitempty"`
	Subscription            *CreateSubForCustomerQuoteSubscriptionParams      `json:"subscription,omitempty"`
	BillingCycles           *int32                                            `json:"billing_cycles,omitempty"`
	Addons                  []*CreateSubForCustomerQuoteAddonParams           `json:"addons,omitempty"`
	EventBasedAddons        []*CreateSubForCustomerQuoteEventBasedAddonParams `json:"event_based_addons,omitempty"`
	TermsToCharge           *int32                                            `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode                         `json:"billing_alignment_mode,omitempty"`
	MandatoryAddonsToRemove []string                                          `json:"mandatory_addons_to_remove,omitempty"`
	ShippingAddress         *CreateSubForCustomerQuoteShippingAddressParams   `json:"shipping_address,omitempty"`
	CouponIds               []string                                          `json:"coupon_ids,omitempty"`
}
type CreateSubForCustomerQuoteSubscriptionParams struct {
	Id            string `json:"id,omitempty"`
	PlanId        string `json:"plan_id"`
	PlanQuantity  *int32 `json:"plan_quantity,omitempty"`
	PlanUnitPrice *int32 `json:"plan_unit_price,omitempty"`
	SetupFee      *int32 `json:"setup_fee,omitempty"`
	StartDate     *int64 `json:"start_date,omitempty"`
	TrialEnd      *int64 `json:"trial_end,omitempty"`
}
type CreateSubForCustomerQuoteAddonParams struct {
	Id            string `json:"id,omitempty"`
	Quantity      *int32 `json:"quantity,omitempty"`
	UnitPrice     *int32 `json:"unit_price,omitempty"`
	BillingCycles *int32 `json:"billing_cycles,omitempty"`
	TrialEnd      *int64 `json:"trial_end,omitempty"`
}
type CreateSubForCustomerQuoteEventBasedAddonParams struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int32        `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
}
type CreateSubForCustomerQuoteShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type UpdateSubscriptionQuoteRequestParams struct {
	Name                    string                                          `json:"name,omitempty"`
	Notes                   string                                          `json:"notes,omitempty"`
	ExpiresAt               *int64                                          `json:"expires_at,omitempty"`
	Subscription            *UpdateSubscriptionQuoteSubscriptionParams      `json:"subscription,omitempty"`
	BillingCycles           *int32                                          `json:"billing_cycles,omitempty"`
	Addons                  []*UpdateSubscriptionQuoteAddonParams           `json:"addons,omitempty"`
	EventBasedAddons        []*UpdateSubscriptionQuoteEventBasedAddonParams `json:"event_based_addons,omitempty"`
	ReplaceAddonList        *bool                                           `json:"replace_addon_list,omitempty"`
	MandatoryAddonsToRemove []string                                        `json:"mandatory_addons_to_remove,omitempty"`
	TermsToCharge           *int32                                          `json:"terms_to_charge,omitempty"`
	ReactivateFrom          *int64                                          `json:"reactivate_from,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode                       `json:"billing_alignment_mode,omitempty"`
	CouponIds               []string                                        `json:"coupon_ids,omitempty"`
	ReplaceCouponList       *bool                                           `json:"replace_coupon_list,omitempty"`
	ForceTermReset          *bool                                           `json:"force_term_reset,omitempty"`
	Reactivate              *bool                                           `json:"reactivate,omitempty"`
	BillingAddress          *UpdateSubscriptionQuoteBillingAddressParams    `json:"billing_address,omitempty"`
	ShippingAddress         *UpdateSubscriptionQuoteShippingAddressParams   `json:"shipping_address,omitempty"`
	Customer                *UpdateSubscriptionQuoteCustomerParams          `json:"customer,omitempty"`
}
type UpdateSubscriptionQuoteSubscriptionParams struct {
	Id            string `json:"id"`
	PlanId        string `json:"plan_id,omitempty"`
	PlanQuantity  *int32 `json:"plan_quantity,omitempty"`
	PlanUnitPrice *int32 `json:"plan_unit_price,omitempty"`
	SetupFee      *int32 `json:"setup_fee,omitempty"`
	StartDate     *int64 `json:"start_date,omitempty"`
	TrialEnd      *int64 `json:"trial_end,omitempty"`
	Coupon        string `json:"coupon,omitempty"`
}
type UpdateSubscriptionQuoteAddonParams struct {
	Id            string `json:"id,omitempty"`
	Quantity      *int32 `json:"quantity,omitempty"`
	UnitPrice     *int32 `json:"unit_price,omitempty"`
	BillingCycles *int32 `json:"billing_cycles,omitempty"`
	TrialEnd      *int64 `json:"trial_end,omitempty"`
}
type UpdateSubscriptionQuoteEventBasedAddonParams struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int32        `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
}
type UpdateSubscriptionQuoteBillingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type UpdateSubscriptionQuoteShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type UpdateSubscriptionQuoteCustomerParams struct {
	VatNumber        string `json:"vat_number,omitempty"`
	RegisteredForGst *bool  `json:"registered_for_gst,omitempty"`
}
type CreateForOnetimeChargesRequestParams struct {
	CustomerId      string                                        `json:"customer_id"`
	PoNumber        string                                        `json:"po_number,omitempty"`
	Name            string                                        `json:"name,omitempty"`
	Notes           string                                        `json:"notes,omitempty"`
	ExpiresAt       *int64                                        `json:"expires_at,omitempty"`
	CurrencyCode    string                                        `json:"currency_code,omitempty"`
	Addons          []*CreateForOnetimeChargesAddonParams         `json:"addons,omitempty"`
	Charges         []*CreateForOnetimeChargesChargeParams        `json:"charges,omitempty"`
	Coupon          string                                        `json:"coupon,omitempty"`
	ShippingAddress *CreateForOnetimeChargesShippingAddressParams `json:"shipping_address,omitempty"`
}
type CreateForOnetimeChargesAddonParams struct {
	Id        string `json:"id,omitempty"`
	Quantity  *int32 `json:"quantity,omitempty"`
	UnitPrice *int32 `json:"unit_price,omitempty"`
	DateFrom  *int64 `json:"date_from,omitempty"`
	DateTo    *int64 `json:"date_to,omitempty"`
}
type CreateForOnetimeChargesChargeParams struct {
	Amount                 *int32               `json:"amount,omitempty"`
	Description            string               `json:"description,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	DateFrom               *int64               `json:"date_from,omitempty"`
	DateTo                 *int64               `json:"date_to,omitempty"`
}
type CreateForOnetimeChargesShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type ListRequestParams struct {
	Limit          *int32                  `json:"limit,omitempty"`
	Offset         string                  `json:"offset,omitempty"`
	IncludeDeleted *bool                   `json:"include_deleted,omitempty"`
	Id             *filter.StringFilter    `json:"id,omitempty"`
	CustomerId     *filter.StringFilter    `json:"customer_id,omitempty"`
	SubscriptionId *filter.StringFilter    `json:"subscription_id,omitempty"`
	Status         *filter.EnumFilter      `json:"status,omitempty"`
	Date           *filter.TimestampFilter `json:"date,omitempty"`
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
	SortBy         *filter.SortFilter      `json:"sort_by,omitempty"`
}
type ConvertRequestParams struct {
	Subscription *ConvertSubscriptionParams `json:"subscription,omitempty"`
}
type ConvertSubscriptionParams struct {
	Id             string              `json:"id,omitempty"`
	AutoCollection enum.AutoCollection `json:"auto_collection,omitempty"`
	PoNumber       string              `json:"po_number,omitempty"`
}
type UpdateStatusRequestParams struct {
	Status  quoteEnum.Status `json:"status"`
	Comment string           `json:"comment,omitempty"`
}
type DeleteRequestParams struct {
	Comment string `json:"comment,omitempty"`
}
type PdfRequestParams struct {
	DispositionType enum.DispositionType `json:"disposition_type,omitempty"`
}
