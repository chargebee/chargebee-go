package quote

import (
	"encoding/json"
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	contractTermEnum "github.com/chargebee/chargebee-go/models/contractterm/enum"
	quoteEnum "github.com/chargebee/chargebee-go/models/quote/enum"
)

type Quote struct {
	Id                         string                  `json:"id"`
	Name                       string                  `json:"name"`
	PoNumber                   string                  `json:"po_number"`
	CustomerId                 string                  `json:"customer_id"`
	SubscriptionId             string                  `json:"subscription_id"`
	InvoiceId                  string                  `json:"invoice_id"`
	Status                     quoteEnum.Status        `json:"status"`
	OperationType              quoteEnum.OperationType `json:"operation_type"`
	VatNumber                  string                  `json:"vat_number"`
	PriceType                  enum.PriceType          `json:"price_type"`
	ValidTill                  int64                   `json:"valid_till"`
	Date                       int64                   `json:"date"`
	TotalPayable               int64                   `json:"total_payable"`
	ChargeOnAcceptance         int32                   `json:"charge_on_acceptance"`
	SubTotal                   int32                   `json:"sub_total"`
	Total                      int32                   `json:"total"`
	CreditsApplied             int32                   `json:"credits_applied"`
	AmountPaid                 int32                   `json:"amount_paid"`
	AmountDue                  int32                   `json:"amount_due"`
	Version                    int32                   `json:"version"`
	ResourceVersion            int64                   `json:"resource_version"`
	UpdatedAt                  int64                   `json:"updated_at"`
	VatNumberPrefix            string                  `json:"vat_number_prefix"`
	LineItems                  []*LineItem             `json:"line_items"`
	Discounts                  []*Discount             `json:"discounts"`
	LineItemDiscounts          []*LineItemDiscount     `json:"line_item_discounts"`
	Taxes                      []*Tax                  `json:"taxes"`
	LineItemTaxes              []*LineItemTax          `json:"line_item_taxes"`
	LineItemTiers              []*LineItemTier         `json:"line_item_tiers"`
	TaxCategory                string                  `json:"tax_category"`
	CurrencyCode               string                  `json:"currency_code"`
	Notes                      json.RawMessage         `json:"notes"`
	ShippingAddress            *ShippingAddress        `json:"shipping_address"`
	BillingAddress             *BillingAddress         `json:"billing_address"`
	ContractTermStart          int64                   `json:"contract_term_start"`
	ContractTermEnd            int64                   `json:"contract_term_end"`
	ContractTermTerminationFee int32                   `json:"contract_term_termination_fee"`
	BusinessEntityId           string                  `json:"business_entity_id"`
	Object                     string                  `json:"object"`
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
	UnitAmountInDecimal     string                       `json:"unit_amount_in_decimal"`
	QuantityInDecimal       string                       `json:"quantity_in_decimal"`
	AmountInDecimal         string                       `json:"amount_in_decimal"`
	DiscountAmount          int32                        `json:"discount_amount"`
	ItemLevelDiscountAmount int32                        `json:"item_level_discount_amount"`
	ReferenceLineItemId     string                       `json:"reference_line_item_id"`
	Description             string                       `json:"description"`
	EntityDescription       string                       `json:"entity_description"`
	EntityType              quoteEnum.LineItemEntityType `json:"entity_type"`
	TaxExemptReason         enum.TaxExemptReason         `json:"tax_exempt_reason"`
	EntityId                string                       `json:"entity_id"`
	CustomerId              string                       `json:"customer_id"`
	Object                  string                       `json:"object"`
}
type Discount struct {
	Amount        int32                        `json:"amount"`
	Description   string                       `json:"description"`
	EntityType    quoteEnum.DiscountEntityType `json:"entity_type"`
	EntityId      string                       `json:"entity_id"`
	CouponSetCode string                       `json:"coupon_set_code"`
	Object        string                       `json:"object"`
}
type LineItemDiscount struct {
	LineItemId     string                                 `json:"line_item_id"`
	DiscountType   quoteEnum.LineItemDiscountDiscountType `json:"discount_type"`
	CouponId       string                                 `json:"coupon_id"`
	EntityId       string                                 `json:"entity_id"`
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
type LineItemTier struct {
	LineItemId            string `json:"line_item_id"`
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
	Index            int32                 `json:"index"`
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
	MandatoryAddonsToRemove []string                                          `json:"mandatory_addons_to_remove,omitempty"`
	TermsToCharge           *int32                                            `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode                         `json:"billing_alignment_mode,omitempty"`
	ShippingAddress         *CreateSubForCustomerQuoteShippingAddressParams   `json:"shipping_address,omitempty"`
	ContractTerm            *CreateSubForCustomerQuoteContractTermParams      `json:"contract_term,omitempty"`
	CouponIds               []string                                          `json:"coupon_ids,omitempty"`
}
type CreateSubForCustomerQuoteSubscriptionParams struct {
	Id                                string                    `json:"id,omitempty"`
	PoNumber                          string                    `json:"po_number,omitempty"`
	PlanId                            string                    `json:"plan_id"`
	PlanQuantity                      *int32                    `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                    `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int32                    `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                    `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int32                    `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type CreateSubForCustomerQuoteAddonParams struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}
type CreateSubForCustomerQuoteEventBasedAddonParams struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int32        `json:"unit_price,omitempty"`
	QuantityInDecimal   string        `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string        `json:"unit_price_in_decimal,omitempty"`
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
type CreateSubForCustomerQuoteContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type EditCreateSubForCustomerQuoteRequestParams struct {
	Notes                   string                                                `json:"notes,omitempty"`
	ExpiresAt               *int64                                                `json:"expires_at,omitempty"`
	Subscription            *EditCreateSubForCustomerQuoteSubscriptionParams      `json:"subscription,omitempty"`
	BillingCycles           *int32                                                `json:"billing_cycles,omitempty"`
	Addons                  []*EditCreateSubForCustomerQuoteAddonParams           `json:"addons,omitempty"`
	EventBasedAddons        []*EditCreateSubForCustomerQuoteEventBasedAddonParams `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove []string                                              `json:"mandatory_addons_to_remove,omitempty"`
	TermsToCharge           *int32                                                `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode                             `json:"billing_alignment_mode,omitempty"`
	ShippingAddress         *EditCreateSubForCustomerQuoteShippingAddressParams   `json:"shipping_address,omitempty"`
	ContractTerm            *EditCreateSubForCustomerQuoteContractTermParams      `json:"contract_term,omitempty"`
	CouponIds               []string                                              `json:"coupon_ids,omitempty"`
}
type EditCreateSubForCustomerQuoteSubscriptionParams struct {
	Id                                string                    `json:"id,omitempty"`
	PoNumber                          string                    `json:"po_number,omitempty"`
	PlanId                            string                    `json:"plan_id"`
	PlanQuantity                      *int32                    `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                    `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int32                    `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                    `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int32                    `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type EditCreateSubForCustomerQuoteAddonParams struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}
type EditCreateSubForCustomerQuoteEventBasedAddonParams struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int32        `json:"unit_price,omitempty"`
	QuantityInDecimal   string        `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string        `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
}
type EditCreateSubForCustomerQuoteShippingAddressParams struct {
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
type EditCreateSubForCustomerQuoteContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type UpdateSubscriptionQuoteRequestParams struct {
	Name                    string                                          `json:"name,omitempty"`
	Notes                   string                                          `json:"notes,omitempty"`
	ExpiresAt               *int64                                          `json:"expires_at,omitempty"`
	Subscription            *UpdateSubscriptionQuoteSubscriptionParams      `json:"subscription,omitempty"`
	Addons                  []*UpdateSubscriptionQuoteAddonParams           `json:"addons,omitempty"`
	EventBasedAddons        []*UpdateSubscriptionQuoteEventBasedAddonParams `json:"event_based_addons,omitempty"`
	ReplaceAddonList        *bool                                           `json:"replace_addon_list,omitempty"`
	MandatoryAddonsToRemove []string                                        `json:"mandatory_addons_to_remove,omitempty"`
	BillingCycles           *int32                                          `json:"billing_cycles,omitempty"`
	TermsToCharge           *int32                                          `json:"terms_to_charge,omitempty"`
	ReactivateFrom          *int64                                          `json:"reactivate_from,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode                       `json:"billing_alignment_mode,omitempty"`
	CouponIds               []string                                        `json:"coupon_ids,omitempty"`
	ReplaceCouponList       *bool                                           `json:"replace_coupon_list,omitempty"`
	ChangeOption            enum.ChangeOption                               `json:"change_option,omitempty"`
	ChangesScheduledAt      *int64                                          `json:"changes_scheduled_at,omitempty"`
	ForceTermReset          *bool                                           `json:"force_term_reset,omitempty"`
	Reactivate              *bool                                           `json:"reactivate,omitempty"`
	BillingAddress          *UpdateSubscriptionQuoteBillingAddressParams    `json:"billing_address,omitempty"`
	ShippingAddress         *UpdateSubscriptionQuoteShippingAddressParams   `json:"shipping_address,omitempty"`
	Customer                *UpdateSubscriptionQuoteCustomerParams          `json:"customer,omitempty"`
	ContractTerm            *UpdateSubscriptionQuoteContractTermParams      `json:"contract_term,omitempty"`
}
type UpdateSubscriptionQuoteSubscriptionParams struct {
	Id                                string                    `json:"id"`
	PlanId                            string                    `json:"plan_id,omitempty"`
	PlanQuantity                      *int32                    `json:"plan_quantity,omitempty"`
	PlanUnitPrice                     *int32                    `json:"plan_unit_price,omitempty"`
	SetupFee                          *int32                    `json:"setup_fee,omitempty"`
	PlanQuantityInDecimal             string                    `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPriceInDecimal            string                    `json:"plan_unit_price_in_decimal,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	Coupon                            string                    `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type UpdateSubscriptionQuoteAddonParams struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}
type UpdateSubscriptionQuoteEventBasedAddonParams struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int32        `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
	QuantityInDecimal   string        `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string        `json:"unit_price_in_decimal,omitempty"`
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
	VatNumberPrefix  string `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool  `json:"registered_for_gst,omitempty"`
}
type UpdateSubscriptionQuoteContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type EditUpdateSubscriptionQuoteRequestParams struct {
	Notes                   string                                              `json:"notes,omitempty"`
	ExpiresAt               *int64                                              `json:"expires_at,omitempty"`
	Subscription            *EditUpdateSubscriptionQuoteSubscriptionParams      `json:"subscription,omitempty"`
	Addons                  []*EditUpdateSubscriptionQuoteAddonParams           `json:"addons,omitempty"`
	EventBasedAddons        []*EditUpdateSubscriptionQuoteEventBasedAddonParams `json:"event_based_addons,omitempty"`
	ReplaceAddonList        *bool                                               `json:"replace_addon_list,omitempty"`
	MandatoryAddonsToRemove []string                                            `json:"mandatory_addons_to_remove,omitempty"`
	BillingCycles           *int32                                              `json:"billing_cycles,omitempty"`
	TermsToCharge           *int32                                              `json:"terms_to_charge,omitempty"`
	ReactivateFrom          *int64                                              `json:"reactivate_from,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode                           `json:"billing_alignment_mode,omitempty"`
	CouponIds               []string                                            `json:"coupon_ids,omitempty"`
	ReplaceCouponList       *bool                                               `json:"replace_coupon_list,omitempty"`
	ChangeOption            enum.ChangeOption                                   `json:"change_option,omitempty"`
	ChangesScheduledAt      *int64                                              `json:"changes_scheduled_at,omitempty"`
	ForceTermReset          *bool                                               `json:"force_term_reset,omitempty"`
	Reactivate              *bool                                               `json:"reactivate,omitempty"`
	BillingAddress          *EditUpdateSubscriptionQuoteBillingAddressParams    `json:"billing_address,omitempty"`
	ShippingAddress         *EditUpdateSubscriptionQuoteShippingAddressParams   `json:"shipping_address,omitempty"`
	Customer                *EditUpdateSubscriptionQuoteCustomerParams          `json:"customer,omitempty"`
	ContractTerm            *EditUpdateSubscriptionQuoteContractTermParams      `json:"contract_term,omitempty"`
}
type EditUpdateSubscriptionQuoteSubscriptionParams struct {
	PlanId                            string                    `json:"plan_id,omitempty"`
	PlanQuantity                      *int32                    `json:"plan_quantity,omitempty"`
	PlanUnitPrice                     *int32                    `json:"plan_unit_price,omitempty"`
	SetupFee                          *int32                    `json:"setup_fee,omitempty"`
	PlanQuantityInDecimal             string                    `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPriceInDecimal            string                    `json:"plan_unit_price_in_decimal,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	Coupon                            string                    `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type EditUpdateSubscriptionQuoteAddonParams struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}
type EditUpdateSubscriptionQuoteEventBasedAddonParams struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int32        `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
	QuantityInDecimal   string        `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string        `json:"unit_price_in_decimal,omitempty"`
}
type EditUpdateSubscriptionQuoteBillingAddressParams struct {
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
type EditUpdateSubscriptionQuoteShippingAddressParams struct {
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
type EditUpdateSubscriptionQuoteCustomerParams struct {
	VatNumber        string `json:"vat_number,omitempty"`
	VatNumberPrefix  string `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool  `json:"registered_for_gst,omitempty"`
}
type EditUpdateSubscriptionQuoteContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type CreateForOnetimeChargesRequestParams struct {
	Name            string                                        `json:"name,omitempty"`
	CustomerId      string                                        `json:"customer_id"`
	PoNumber        string                                        `json:"po_number,omitempty"`
	Notes           string                                        `json:"notes,omitempty"`
	ExpiresAt       *int64                                        `json:"expires_at,omitempty"`
	CurrencyCode    string                                        `json:"currency_code,omitempty"`
	Addons          []*CreateForOnetimeChargesAddonParams         `json:"addons,omitempty"`
	Charges         []*CreateForOnetimeChargesChargeParams        `json:"charges,omitempty"`
	Coupon          string                                        `json:"coupon,omitempty"`
	CouponIds       []string                                      `json:"coupon_ids,omitempty"`
	ShippingAddress *CreateForOnetimeChargesShippingAddressParams `json:"shipping_address,omitempty"`
}
type CreateForOnetimeChargesAddonParams struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	ServicePeriod      *int32 `json:"service_period,omitempty"`
}
type CreateForOnetimeChargesChargeParams struct {
	Amount                 *int32               `json:"amount,omitempty"`
	AmountInDecimal        string               `json:"amount_in_decimal,omitempty"`
	Description            string               `json:"description,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	ServicePeriod          *int32               `json:"service_period,omitempty"`
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
type EditOneTimeQuoteRequestParams struct {
	PoNumber        string                                 `json:"po_number,omitempty"`
	Notes           string                                 `json:"notes,omitempty"`
	ExpiresAt       *int64                                 `json:"expires_at,omitempty"`
	CurrencyCode    string                                 `json:"currency_code,omitempty"`
	Addons          []*EditOneTimeQuoteAddonParams         `json:"addons,omitempty"`
	Charges         []*EditOneTimeQuoteChargeParams        `json:"charges,omitempty"`
	Coupon          string                                 `json:"coupon,omitempty"`
	CouponIds       []string                               `json:"coupon_ids,omitempty"`
	ShippingAddress *EditOneTimeQuoteShippingAddressParams `json:"shipping_address,omitempty"`
}
type EditOneTimeQuoteAddonParams struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	ServicePeriod      *int32 `json:"service_period,omitempty"`
}
type EditOneTimeQuoteChargeParams struct {
	Amount                 *int32               `json:"amount,omitempty"`
	AmountInDecimal        string               `json:"amount_in_decimal,omitempty"`
	Description            string               `json:"description,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	ServicePeriod          *int32               `json:"service_period,omitempty"`
}
type EditOneTimeQuoteShippingAddressParams struct {
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
type CreateSubItemsForCustomerQuoteRequestParams struct {
	Name                   string                                                  `json:"name,omitempty"`
	Notes                  string                                                  `json:"notes,omitempty"`
	ExpiresAt              *int64                                                  `json:"expires_at,omitempty"`
	Subscription           *CreateSubItemsForCustomerQuoteSubscriptionParams       `json:"subscription,omitempty"`
	BillingCycles          *int32                                                  `json:"billing_cycles,omitempty"`
	SubscriptionItems      []*CreateSubItemsForCustomerQuoteSubscriptionItemParams `json:"subscription_items,omitempty"`
	Discounts              []*CreateSubItemsForCustomerQuoteDiscountParams         `json:"discounts,omitempty"`
	MandatoryItemsToRemove []string                                                `json:"mandatory_items_to_remove,omitempty"`
	ItemTiers              []*CreateSubItemsForCustomerQuoteItemTierParams         `json:"item_tiers,omitempty"`
	TermsToCharge          *int32                                                  `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode   enum.BillingAlignmentMode                               `json:"billing_alignment_mode,omitempty"`
	ShippingAddress        *CreateSubItemsForCustomerQuoteShippingAddressParams    `json:"shipping_address,omitempty"`
	ContractTerm           *CreateSubItemsForCustomerQuoteContractTermParams       `json:"contract_term,omitempty"`
	CouponIds              []string                                                `json:"coupon_ids,omitempty"`
}
type CreateSubItemsForCustomerQuoteSubscriptionParams struct {
	Id                                string                    `json:"id,omitempty"`
	PoNumber                          string                    `json:"po_number,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	SetupFee                          *int32                    `json:"setup_fee,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type CreateSubItemsForCustomerQuoteSubscriptionItemParams struct {
	ItemPriceId        string              `json:"item_price_id"`
	Quantity           *int32              `json:"quantity,omitempty"`
	QuantityInDecimal  string              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32              `json:"billing_cycles,omitempty"`
	TrialEnd           *int64              `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32              `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool               `json:"charge_once,omitempty"`
	ItemType           enum.ItemType       `json:"item_type,omitempty"`
	ChargeOnOption     enum.ChargeOnOption `json:"charge_on_option,omitempty"`
}
type CreateSubItemsForCustomerQuoteDiscountParams struct {
	ApplyOn       enum.ApplyOn      `json:"apply_on"`
	DurationType  enum.DurationType `json:"duration_type"`
	Percentage    *float64          `json:"percentage,omitempty"`
	Amount        *int32            `json:"amount,omitempty"`
	Period        *int32            `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool             `json:"included_in_mrr,omitempty"`
	ItemPriceId   string            `json:"item_price_id,omitempty"`
}
type CreateSubItemsForCustomerQuoteItemTierParams struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type CreateSubItemsForCustomerQuoteShippingAddressParams struct {
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
type CreateSubItemsForCustomerQuoteContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type EditCreateSubCustomerQuoteForItemsRequestParams struct {
	Notes                  string                                                      `json:"notes,omitempty"`
	ExpiresAt              *int64                                                      `json:"expires_at,omitempty"`
	Subscription           *EditCreateSubCustomerQuoteForItemsSubscriptionParams       `json:"subscription,omitempty"`
	BillingCycles          *int32                                                      `json:"billing_cycles,omitempty"`
	SubscriptionItems      []*EditCreateSubCustomerQuoteForItemsSubscriptionItemParams `json:"subscription_items,omitempty"`
	Discounts              []*EditCreateSubCustomerQuoteForItemsDiscountParams         `json:"discounts,omitempty"`
	MandatoryItemsToRemove []string                                                    `json:"mandatory_items_to_remove,omitempty"`
	ItemTiers              []*EditCreateSubCustomerQuoteForItemsItemTierParams         `json:"item_tiers,omitempty"`
	TermsToCharge          *int32                                                      `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode   enum.BillingAlignmentMode                                   `json:"billing_alignment_mode,omitempty"`
	ShippingAddress        *EditCreateSubCustomerQuoteForItemsShippingAddressParams    `json:"shipping_address,omitempty"`
	ContractTerm           *EditCreateSubCustomerQuoteForItemsContractTermParams       `json:"contract_term,omitempty"`
	CouponIds              []string                                                    `json:"coupon_ids,omitempty"`
}
type EditCreateSubCustomerQuoteForItemsSubscriptionParams struct {
	Id                                string                    `json:"id,omitempty"`
	PoNumber                          string                    `json:"po_number,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	SetupFee                          *int32                    `json:"setup_fee,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type EditCreateSubCustomerQuoteForItemsSubscriptionItemParams struct {
	ItemPriceId        string              `json:"item_price_id"`
	Quantity           *int32              `json:"quantity,omitempty"`
	QuantityInDecimal  string              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32              `json:"billing_cycles,omitempty"`
	TrialEnd           *int64              `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32              `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool               `json:"charge_once,omitempty"`
	ItemType           enum.ItemType       `json:"item_type,omitempty"`
	ChargeOnOption     enum.ChargeOnOption `json:"charge_on_option,omitempty"`
}
type EditCreateSubCustomerQuoteForItemsDiscountParams struct {
	ApplyOn       enum.ApplyOn      `json:"apply_on"`
	DurationType  enum.DurationType `json:"duration_type"`
	Percentage    *float64          `json:"percentage,omitempty"`
	Amount        *int32            `json:"amount,omitempty"`
	Period        *int32            `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool             `json:"included_in_mrr,omitempty"`
	ItemPriceId   string            `json:"item_price_id,omitempty"`
}
type EditCreateSubCustomerQuoteForItemsItemTierParams struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type EditCreateSubCustomerQuoteForItemsShippingAddressParams struct {
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
type EditCreateSubCustomerQuoteForItemsContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type UpdateSubscriptionQuoteForItemsRequestParams struct {
	Name                   string                                                   `json:"name,omitempty"`
	Notes                  string                                                   `json:"notes,omitempty"`
	ExpiresAt              *int64                                                   `json:"expires_at,omitempty"`
	Subscription           *UpdateSubscriptionQuoteForItemsSubscriptionParams       `json:"subscription,omitempty"`
	SubscriptionItems      []*UpdateSubscriptionQuoteForItemsSubscriptionItemParams `json:"subscription_items,omitempty"`
	MandatoryItemsToRemove []string                                                 `json:"mandatory_items_to_remove,omitempty"`
	ReplaceItemsList       *bool                                                    `json:"replace_items_list,omitempty"`
	Discounts              []*UpdateSubscriptionQuoteForItemsDiscountParams         `json:"discounts,omitempty"`
	ItemTiers              []*UpdateSubscriptionQuoteForItemsItemTierParams         `json:"item_tiers,omitempty"`
	BillingCycles          *int32                                                   `json:"billing_cycles,omitempty"`
	TermsToCharge          *int32                                                   `json:"terms_to_charge,omitempty"`
	ReactivateFrom         *int64                                                   `json:"reactivate_from,omitempty"`
	BillingAlignmentMode   enum.BillingAlignmentMode                                `json:"billing_alignment_mode,omitempty"`
	CouponIds              []string                                                 `json:"coupon_ids,omitempty"`
	ReplaceCouponList      *bool                                                    `json:"replace_coupon_list,omitempty"`
	ChangeOption           enum.ChangeOption                                        `json:"change_option,omitempty"`
	ChangesScheduledAt     *int64                                                   `json:"changes_scheduled_at,omitempty"`
	ForceTermReset         *bool                                                    `json:"force_term_reset,omitempty"`
	Reactivate             *bool                                                    `json:"reactivate,omitempty"`
	BillingAddress         *UpdateSubscriptionQuoteForItemsBillingAddressParams     `json:"billing_address,omitempty"`
	ShippingAddress        *UpdateSubscriptionQuoteForItemsShippingAddressParams    `json:"shipping_address,omitempty"`
	Customer               *UpdateSubscriptionQuoteForItemsCustomerParams           `json:"customer,omitempty"`
	ContractTerm           *UpdateSubscriptionQuoteForItemsContractTermParams       `json:"contract_term,omitempty"`
}
type UpdateSubscriptionQuoteForItemsSubscriptionParams struct {
	Id                                string                    `json:"id"`
	SetupFee                          *int32                    `json:"setup_fee,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	Coupon                            string                    `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type UpdateSubscriptionQuoteForItemsSubscriptionItemParams struct {
	ItemPriceId        string              `json:"item_price_id"`
	Quantity           *int32              `json:"quantity,omitempty"`
	QuantityInDecimal  string              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32              `json:"billing_cycles,omitempty"`
	TrialEnd           *int64              `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32              `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool               `json:"charge_once,omitempty"`
	ChargeOnOption     enum.ChargeOnOption `json:"charge_on_option,omitempty"`
	ItemType           enum.ItemType       `json:"item_type,omitempty"`
}
type UpdateSubscriptionQuoteForItemsDiscountParams struct {
	ApplyOn       enum.ApplyOn       `json:"apply_on"`
	DurationType  enum.DurationType  `json:"duration_type"`
	Percentage    *float64           `json:"percentage,omitempty"`
	Amount        *int32             `json:"amount,omitempty"`
	Period        *int32             `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit    `json:"period_unit,omitempty"`
	IncludedInMrr *bool              `json:"included_in_mrr,omitempty"`
	ItemPriceId   string             `json:"item_price_id,omitempty"`
	OperationType enum.OperationType `json:"operation_type"`
	Id            string             `json:"id,omitempty"`
}
type UpdateSubscriptionQuoteForItemsItemTierParams struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type UpdateSubscriptionQuoteForItemsBillingAddressParams struct {
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
type UpdateSubscriptionQuoteForItemsShippingAddressParams struct {
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
type UpdateSubscriptionQuoteForItemsCustomerParams struct {
	VatNumber        string `json:"vat_number,omitempty"`
	VatNumberPrefix  string `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool  `json:"registered_for_gst,omitempty"`
}
type UpdateSubscriptionQuoteForItemsContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type EditUpdateSubscriptionQuoteForItemsRequestParams struct {
	Notes                  string                                                       `json:"notes,omitempty"`
	ExpiresAt              *int64                                                       `json:"expires_at,omitempty"`
	SubscriptionItems      []*EditUpdateSubscriptionQuoteForItemsSubscriptionItemParams `json:"subscription_items,omitempty"`
	MandatoryItemsToRemove []string                                                     `json:"mandatory_items_to_remove,omitempty"`
	ReplaceItemsList       *bool                                                        `json:"replace_items_list,omitempty"`
	Subscription           *EditUpdateSubscriptionQuoteForItemsSubscriptionParams       `json:"subscription,omitempty"`
	Discounts              []*EditUpdateSubscriptionQuoteForItemsDiscountParams         `json:"discounts,omitempty"`
	ItemTiers              []*EditUpdateSubscriptionQuoteForItemsItemTierParams         `json:"item_tiers,omitempty"`
	BillingCycles          *int32                                                       `json:"billing_cycles,omitempty"`
	TermsToCharge          *int32                                                       `json:"terms_to_charge,omitempty"`
	ReactivateFrom         *int64                                                       `json:"reactivate_from,omitempty"`
	BillingAlignmentMode   enum.BillingAlignmentMode                                    `json:"billing_alignment_mode,omitempty"`
	CouponIds              []string                                                     `json:"coupon_ids,omitempty"`
	ReplaceCouponList      *bool                                                        `json:"replace_coupon_list,omitempty"`
	ChangeOption           enum.ChangeOption                                            `json:"change_option,omitempty"`
	ChangesScheduledAt     *int64                                                       `json:"changes_scheduled_at,omitempty"`
	ForceTermReset         *bool                                                        `json:"force_term_reset,omitempty"`
	Reactivate             *bool                                                        `json:"reactivate,omitempty"`
	BillingAddress         *EditUpdateSubscriptionQuoteForItemsBillingAddressParams     `json:"billing_address,omitempty"`
	ShippingAddress        *EditUpdateSubscriptionQuoteForItemsShippingAddressParams    `json:"shipping_address,omitempty"`
	Customer               *EditUpdateSubscriptionQuoteForItemsCustomerParams           `json:"customer,omitempty"`
	ContractTerm           *EditUpdateSubscriptionQuoteForItemsContractTermParams       `json:"contract_term,omitempty"`
}
type EditUpdateSubscriptionQuoteForItemsSubscriptionItemParams struct {
	ItemPriceId        string              `json:"item_price_id"`
	Quantity           *int32              `json:"quantity,omitempty"`
	QuantityInDecimal  string              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32              `json:"billing_cycles,omitempty"`
	TrialEnd           *int64              `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32              `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool               `json:"charge_once,omitempty"`
	ChargeOnOption     enum.ChargeOnOption `json:"charge_on_option,omitempty"`
	ItemType           enum.ItemType       `json:"item_type,omitempty"`
}
type EditUpdateSubscriptionQuoteForItemsSubscriptionParams struct {
	SetupFee                          *int32                    `json:"setup_fee,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	Coupon                            string                    `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type EditUpdateSubscriptionQuoteForItemsDiscountParams struct {
	ApplyOn       enum.ApplyOn       `json:"apply_on"`
	DurationType  enum.DurationType  `json:"duration_type"`
	Percentage    *float64           `json:"percentage,omitempty"`
	Amount        *int32             `json:"amount,omitempty"`
	Period        *int32             `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit    `json:"period_unit,omitempty"`
	IncludedInMrr *bool              `json:"included_in_mrr,omitempty"`
	ItemPriceId   string             `json:"item_price_id,omitempty"`
	OperationType enum.OperationType `json:"operation_type"`
	Id            string             `json:"id,omitempty"`
}
type EditUpdateSubscriptionQuoteForItemsItemTierParams struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type EditUpdateSubscriptionQuoteForItemsBillingAddressParams struct {
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
type EditUpdateSubscriptionQuoteForItemsShippingAddressParams struct {
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
type EditUpdateSubscriptionQuoteForItemsCustomerParams struct {
	VatNumber        string `json:"vat_number,omitempty"`
	VatNumberPrefix  string `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool  `json:"registered_for_gst,omitempty"`
}
type EditUpdateSubscriptionQuoteForItemsContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type CreateForChargeItemsAndChargesRequestParams struct {
	Name            string                                               `json:"name,omitempty"`
	CustomerId      string                                               `json:"customer_id"`
	PoNumber        string                                               `json:"po_number,omitempty"`
	Notes           string                                               `json:"notes,omitempty"`
	ExpiresAt       *int64                                               `json:"expires_at,omitempty"`
	CurrencyCode    string                                               `json:"currency_code,omitempty"`
	ItemPrices      []*CreateForChargeItemsAndChargesItemPriceParams     `json:"item_prices,omitempty"`
	ItemTiers       []*CreateForChargeItemsAndChargesItemTierParams      `json:"item_tiers,omitempty"`
	Charges         []*CreateForChargeItemsAndChargesChargeParams        `json:"charges,omitempty"`
	Coupon          string                                               `json:"coupon,omitempty"`
	CouponIds       []string                                             `json:"coupon_ids,omitempty"`
	ShippingAddress *CreateForChargeItemsAndChargesShippingAddressParams `json:"shipping_address,omitempty"`
	Discounts       []*CreateForChargeItemsAndChargesDiscountParams      `json:"discounts,omitempty"`
}
type CreateForChargeItemsAndChargesItemPriceParams struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodDays  *int32 `json:"service_period_days,omitempty"`
}
type CreateForChargeItemsAndChargesItemTierParams struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type CreateForChargeItemsAndChargesChargeParams struct {
	Amount                 *int32               `json:"amount,omitempty"`
	AmountInDecimal        string               `json:"amount_in_decimal,omitempty"`
	Description            string               `json:"description,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	ServicePeriod          *int32               `json:"service_period,omitempty"`
}
type CreateForChargeItemsAndChargesShippingAddressParams struct {
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
type CreateForChargeItemsAndChargesDiscountParams struct {
	Percentage  *float64     `json:"percentage,omitempty"`
	Amount      *int32       `json:"amount,omitempty"`
	ApplyOn     enum.ApplyOn `json:"apply_on"`
	ItemPriceId string       `json:"item_price_id,omitempty"`
}
type EditForChargeItemsAndChargesRequestParams struct {
	PoNumber        string                                             `json:"po_number,omitempty"`
	Notes           string                                             `json:"notes,omitempty"`
	ExpiresAt       *int64                                             `json:"expires_at,omitempty"`
	CurrencyCode    string                                             `json:"currency_code,omitempty"`
	ItemPrices      []*EditForChargeItemsAndChargesItemPriceParams     `json:"item_prices,omitempty"`
	ItemTiers       []*EditForChargeItemsAndChargesItemTierParams      `json:"item_tiers,omitempty"`
	Charges         []*EditForChargeItemsAndChargesChargeParams        `json:"charges,omitempty"`
	Coupon          string                                             `json:"coupon,omitempty"`
	CouponIds       []string                                           `json:"coupon_ids,omitempty"`
	ShippingAddress *EditForChargeItemsAndChargesShippingAddressParams `json:"shipping_address,omitempty"`
	Discounts       []*EditForChargeItemsAndChargesDiscountParams      `json:"discounts,omitempty"`
}
type EditForChargeItemsAndChargesItemPriceParams struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodDays  *int32 `json:"service_period_days,omitempty"`
}
type EditForChargeItemsAndChargesItemTierParams struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type EditForChargeItemsAndChargesChargeParams struct {
	Amount                 *int32               `json:"amount,omitempty"`
	AmountInDecimal        string               `json:"amount_in_decimal,omitempty"`
	Description            string               `json:"description,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	ServicePeriod          *int32               `json:"service_period,omitempty"`
}
type EditForChargeItemsAndChargesShippingAddressParams struct {
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
type EditForChargeItemsAndChargesDiscountParams struct {
	Percentage  *float64     `json:"percentage,omitempty"`
	Amount      *int32       `json:"amount,omitempty"`
	ApplyOn     enum.ApplyOn `json:"apply_on"`
	ItemPriceId string       `json:"item_price_id,omitempty"`
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
type QuoteLineGroupsForQuoteRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
type ConvertRequestParams struct {
	Subscription          *ConvertSubscriptionParams `json:"subscription,omitempty"`
	InvoiceDate           *int64                     `json:"invoice_date,omitempty"`
	CreatePendingInvoices *bool                      `json:"create_pending_invoices,omitempty"`
	FirstInvoicePending   *bool                      `json:"first_invoice_pending,omitempty"`
}
type ConvertSubscriptionParams struct {
	Id                string              `json:"id,omitempty"`
	AutoCollection    enum.AutoCollection `json:"auto_collection,omitempty"`
	PoNumber          string              `json:"po_number,omitempty"`
	AutoCloseInvoices *bool               `json:"auto_close_invoices,omitempty"`
}
type UpdateStatusRequestParams struct {
	Status  quoteEnum.Status `json:"status"`
	Comment string           `json:"comment,omitempty"`
}
type ExtendExpiryDateRequestParams struct {
	ValidTill *int64 `json:"valid_till"`
}
type DeleteRequestParams struct {
	Comment string `json:"comment,omitempty"`
}
type PdfRequestParams struct {
	ConsolidatedView *bool                `json:"consolidated_view,omitempty"`
	DispositionType  enum.DispositionType `json:"disposition_type,omitempty"`
}
