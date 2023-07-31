package estimate

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/models/subscriptionestimate"
	"github.com/chargebee/chargebee-go/models/invoiceestimate"
	"github.com/chargebee/chargebee-go/models/creditnoteestimate"
	"github.com/chargebee/chargebee-go/models/unbilledcharge"
	contractTermEnum "github.com/chargebee/chargebee-go/models/contractterm/enum"
	paymentIntentEnum "github.com/chargebee/chargebee-go/models/paymentintent/enum"
)

type Estimate struct {
	CreatedAt               int64                                        `json:"created_at"`
	SubscriptionEstimate    *subscriptionestimate.SubscriptionEstimate   `json:"subscription_estimate"`
	SubscriptionEstimates   []*subscriptionestimate.SubscriptionEstimate `json:"subscription_estimates"`
	InvoiceEstimate         *invoiceestimate.InvoiceEstimate             `json:"invoice_estimate"`
	InvoiceEstimates        []*invoiceestimate.InvoiceEstimate           `json:"invoice_estimates"`
	NextInvoiceEstimate     *invoiceestimate.InvoiceEstimate             `json:"next_invoice_estimate"`
	CreditNoteEstimates     []*creditnoteestimate.CreditNoteEstimate     `json:"credit_note_estimates"`
	UnbilledChargeEstimates []*unbilledcharge.UnbilledCharge             `json:"unbilled_charge_estimates"`
	Object                  string                                       `json:"object"`
}
type CreateSubscriptionRequestParams struct {
	Subscription            *CreateSubscriptionSubscriptionParams      `json:"subscription,omitempty"`
	BillingCycles           *int32                                     `json:"billing_cycles,omitempty"`
	Addons                  []*CreateSubscriptionAddonParams           `json:"addons,omitempty"`
	EventBasedAddons        []*CreateSubscriptionEventBasedAddonParams `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove []string                                   `json:"mandatory_addons_to_remove,omitempty"`
	TermsToCharge           *int32                                     `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode                  `json:"billing_alignment_mode,omitempty"`
	CouponIds               []string                                   `json:"coupon_ids,omitempty"`
	BillingAddress          *CreateSubscriptionBillingAddressParams    `json:"billing_address,omitempty"`
	ShippingAddress         *CreateSubscriptionShippingAddressParams   `json:"shipping_address,omitempty"`
	Customer                *CreateSubscriptionCustomerParams          `json:"customer,omitempty"`
	InvoiceImmediately      *bool                                      `json:"invoice_immediately,omitempty"`
	InvoiceDate             *int64                                     `json:"invoice_date,omitempty"`
	ContractTerm            *CreateSubscriptionContractTermParams      `json:"contract_term,omitempty"`
	ClientProfileId         string                                     `json:"client_profile_id,omitempty"`
}
type CreateSubscriptionSubscriptionParams struct {
	Id                                string                    `json:"id,omitempty"`
	PlanId                            string                    `json:"plan_id"`
	PlanQuantity                      *int32                    `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                    `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int32                    `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                    `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int32                    `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	Coupon                            string                    `json:"coupon,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	FreePeriod                        *int32                    `json:"free_period,omitempty"`
	FreePeriodUnit                    enum.FreePeriodUnit       `json:"free_period_unit,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	TrialEndAction                    enum.TrialEndAction       `json:"trial_end_action,omitempty"`
}
type CreateSubscriptionAddonParams struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}
type CreateSubscriptionEventBasedAddonParams struct {
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
type CreateSubscriptionBillingAddressParams struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateSubscriptionShippingAddressParams struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateSubscriptionCustomerParams struct {
	VatNumber        string                   `json:"vat_number,omitempty"`
	VatNumberPrefix  string                   `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool                    `json:"registered_for_gst,omitempty"`
	Taxability       enum.Taxability          `json:"taxability,omitempty"`
	EntityCode       enum.EntityCode          `json:"entity_code,omitempty"`
	ExemptNumber     string                   `json:"exempt_number,omitempty"`
	ExemptionDetails []map[string]interface{} `json:"exemption_details,omitempty"`
	CustomerType     enum.CustomerType        `json:"customer_type,omitempty"`
}
type CreateSubscriptionContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type CreateSubItemEstimateRequestParams struct {
	Subscription           *CreateSubItemEstimateSubscriptionParams       `json:"subscription,omitempty"`
	BillingCycles          *int32                                         `json:"billing_cycles,omitempty"`
	SubscriptionItems      []*CreateSubItemEstimateSubscriptionItemParams `json:"subscription_items,omitempty"`
	Discounts              []*CreateSubItemEstimateDiscountParams         `json:"discounts,omitempty"`
	MandatoryItemsToRemove []string                                       `json:"mandatory_items_to_remove,omitempty"`
	ItemTiers              []*CreateSubItemEstimateItemTierParams         `json:"item_tiers,omitempty"`
	TermsToCharge          *int32                                         `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode   enum.BillingAlignmentMode                      `json:"billing_alignment_mode,omitempty"`
	CouponIds              []string                                       `json:"coupon_ids,omitempty"`
	BillingAddress         *CreateSubItemEstimateBillingAddressParams     `json:"billing_address,omitempty"`
	ShippingAddress        *CreateSubItemEstimateShippingAddressParams    `json:"shipping_address,omitempty"`
	Customer               *CreateSubItemEstimateCustomerParams           `json:"customer,omitempty"`
	InvoiceImmediately     *bool                                          `json:"invoice_immediately,omitempty"`
	InvoiceDate            *int64                                         `json:"invoice_date,omitempty"`
	ClientProfileId        string                                         `json:"client_profile_id,omitempty"`
	ContractTerm           *CreateSubItemEstimateContractTermParams       `json:"contract_term,omitempty"`
}
type CreateSubItemEstimateSubscriptionParams struct {
	Id                                string                    `json:"id,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	SetupFee                          *int32                    `json:"setup_fee,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	Coupon                            string                    `json:"coupon,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	FreePeriod                        *int32                    `json:"free_period,omitempty"`
	FreePeriodUnit                    enum.FreePeriodUnit       `json:"free_period_unit,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	TrialEndAction                    enum.TrialEndAction       `json:"trial_end_action,omitempty"`
}
type CreateSubItemEstimateSubscriptionItemParams struct {
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
type CreateSubItemEstimateDiscountParams struct {
	ApplyOn       enum.ApplyOn      `json:"apply_on"`
	DurationType  enum.DurationType `json:"duration_type"`
	Percentage    *float64          `json:"percentage,omitempty"`
	Amount        *int32            `json:"amount,omitempty"`
	Period        *int32            `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool             `json:"included_in_mrr,omitempty"`
	ItemPriceId   string            `json:"item_price_id,omitempty"`
}
type CreateSubItemEstimateItemTierParams struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type CreateSubItemEstimateBillingAddressParams struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateSubItemEstimateShippingAddressParams struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateSubItemEstimateCustomerParams struct {
	VatNumber        string                   `json:"vat_number,omitempty"`
	VatNumberPrefix  string                   `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool                    `json:"registered_for_gst,omitempty"`
	Taxability       enum.Taxability          `json:"taxability,omitempty"`
	EntityCode       enum.EntityCode          `json:"entity_code,omitempty"`
	ExemptNumber     string                   `json:"exempt_number,omitempty"`
	ExemptionDetails []map[string]interface{} `json:"exemption_details,omitempty"`
	CustomerType     enum.CustomerType        `json:"customer_type,omitempty"`
}
type CreateSubItemEstimateContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type CreateSubForCustomerEstimateRequestParams struct {
	UseExistingBalances     *bool                                                `json:"use_existing_balances,omitempty"`
	Subscription            *CreateSubForCustomerEstimateSubscriptionParams      `json:"subscription,omitempty"`
	InvoiceImmediately      *bool                                                `json:"invoice_immediately,omitempty"`
	BillingCycles           *int32                                               `json:"billing_cycles,omitempty"`
	Addons                  []*CreateSubForCustomerEstimateAddonParams           `json:"addons,omitempty"`
	EventBasedAddons        []*CreateSubForCustomerEstimateEventBasedAddonParams `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove []string                                             `json:"mandatory_addons_to_remove,omitempty"`
	TermsToCharge           *int32                                               `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode                            `json:"billing_alignment_mode,omitempty"`
	ShippingAddress         *CreateSubForCustomerEstimateShippingAddressParams   `json:"shipping_address,omitempty"`
	InvoiceDate             *int64                                               `json:"invoice_date,omitempty"`
	CouponIds               []string                                             `json:"coupon_ids,omitempty"`
	ContractTerm            *CreateSubForCustomerEstimateContractTermParams      `json:"contract_term,omitempty"`
}
type CreateSubForCustomerEstimateSubscriptionParams struct {
	Id                                string                    `json:"id,omitempty"`
	PlanId                            string                    `json:"plan_id"`
	PlanQuantity                      *int32                    `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                    `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int32                    `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                    `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int32                    `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	FreePeriod                        *int32                    `json:"free_period,omitempty"`
	FreePeriodUnit                    enum.FreePeriodUnit       `json:"free_period_unit,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	TrialEndAction                    enum.TrialEndAction       `json:"trial_end_action,omitempty"`
}
type CreateSubForCustomerEstimateAddonParams struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}
type CreateSubForCustomerEstimateEventBasedAddonParams struct {
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
type CreateSubForCustomerEstimateShippingAddressParams struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateSubForCustomerEstimateContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type CreateSubItemForCustomerEstimateRequestParams struct {
	UseExistingBalances    *bool                                                     `json:"use_existing_balances,omitempty"`
	Subscription           *CreateSubItemForCustomerEstimateSubscriptionParams       `json:"subscription,omitempty"`
	InvoiceImmediately     *bool                                                     `json:"invoice_immediately,omitempty"`
	BillingCycles          *int32                                                    `json:"billing_cycles,omitempty"`
	SubscriptionItems      []*CreateSubItemForCustomerEstimateSubscriptionItemParams `json:"subscription_items,omitempty"`
	Discounts              []*CreateSubItemForCustomerEstimateDiscountParams         `json:"discounts,omitempty"`
	MandatoryItemsToRemove []string                                                  `json:"mandatory_items_to_remove,omitempty"`
	ItemTiers              []*CreateSubItemForCustomerEstimateItemTierParams         `json:"item_tiers,omitempty"`
	TermsToCharge          *int32                                                    `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode   enum.BillingAlignmentMode                                 `json:"billing_alignment_mode,omitempty"`
	ShippingAddress        *CreateSubItemForCustomerEstimateShippingAddressParams    `json:"shipping_address,omitempty"`
	InvoiceDate            *int64                                                    `json:"invoice_date,omitempty"`
	CouponIds              []string                                                  `json:"coupon_ids,omitempty"`
	ContractTerm           *CreateSubItemForCustomerEstimateContractTermParams       `json:"contract_term,omitempty"`
}
type CreateSubItemForCustomerEstimateSubscriptionParams struct {
	Id                                string                    `json:"id,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	SetupFee                          *int32                    `json:"setup_fee,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	FreePeriod                        *int32                    `json:"free_period,omitempty"`
	FreePeriodUnit                    enum.FreePeriodUnit       `json:"free_period_unit,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	TrialEndAction                    enum.TrialEndAction       `json:"trial_end_action,omitempty"`
}
type CreateSubItemForCustomerEstimateSubscriptionItemParams struct {
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
type CreateSubItemForCustomerEstimateDiscountParams struct {
	ApplyOn       enum.ApplyOn      `json:"apply_on"`
	DurationType  enum.DurationType `json:"duration_type"`
	Percentage    *float64          `json:"percentage,omitempty"`
	Amount        *int32            `json:"amount,omitempty"`
	Period        *int32            `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool             `json:"included_in_mrr,omitempty"`
	ItemPriceId   string            `json:"item_price_id,omitempty"`
}
type CreateSubItemForCustomerEstimateItemTierParams struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type CreateSubItemForCustomerEstimateShippingAddressParams struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateSubItemForCustomerEstimateContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type UpdateSubscriptionRequestParams struct {
	Subscription            *UpdateSubscriptionSubscriptionParams      `json:"subscription,omitempty"`
	Addons                  []*UpdateSubscriptionAddonParams           `json:"addons,omitempty"`
	EventBasedAddons        []*UpdateSubscriptionEventBasedAddonParams `json:"event_based_addons,omitempty"`
	ReplaceAddonList        *bool                                      `json:"replace_addon_list,omitempty"`
	MandatoryAddonsToRemove []string                                   `json:"mandatory_addons_to_remove,omitempty"`
	InvoiceDate             *int64                                     `json:"invoice_date,omitempty"`
	BillingCycles           *int32                                     `json:"billing_cycles,omitempty"`
	TermsToCharge           *int32                                     `json:"terms_to_charge,omitempty"`
	ReactivateFrom          *int64                                     `json:"reactivate_from,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode                  `json:"billing_alignment_mode,omitempty"`
	CouponIds               []string                                   `json:"coupon_ids,omitempty"`
	ReplaceCouponList       *bool                                      `json:"replace_coupon_list,omitempty"`
	Prorate                 *bool                                      `json:"prorate,omitempty"`
	EndOfTerm               *bool                                      `json:"end_of_term,omitempty"`
	ForceTermReset          *bool                                      `json:"force_term_reset,omitempty"`
	Reactivate              *bool                                      `json:"reactivate,omitempty"`
	IncludeDelayedCharges   *bool                                      `json:"include_delayed_charges,omitempty"`
	UseExistingBalances     *bool                                      `json:"use_existing_balances,omitempty"`
	BillingAddress          *UpdateSubscriptionBillingAddressParams    `json:"billing_address,omitempty"`
	ShippingAddress         *UpdateSubscriptionShippingAddressParams   `json:"shipping_address,omitempty"`
	Customer                *UpdateSubscriptionCustomerParams          `json:"customer,omitempty"`
	InvoiceImmediately      *bool                                      `json:"invoice_immediately,omitempty"`
}
type UpdateSubscriptionSubscriptionParams struct {
	Id                     string                    `json:"id"`
	PlanId                 string                    `json:"plan_id,omitempty"`
	PlanQuantity           *int32                    `json:"plan_quantity,omitempty"`
	PlanUnitPrice          *int32                    `json:"plan_unit_price,omitempty"`
	SetupFee               *int32                    `json:"setup_fee,omitempty"`
	PlanQuantityInDecimal  string                    `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPriceInDecimal string                    `json:"plan_unit_price_in_decimal,omitempty"`
	StartDate              *int64                    `json:"start_date,omitempty"`
	TrialEnd               *int64                    `json:"trial_end,omitempty"`
	Coupon                 string                    `json:"coupon,omitempty"`
	AutoCollection         enum.AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod   enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	FreePeriod             *int32                    `json:"free_period,omitempty"`
	FreePeriodUnit         enum.FreePeriodUnit       `json:"free_period_unit,omitempty"`
	TrialEndAction         enum.TrialEndAction       `json:"trial_end_action,omitempty"`
}
type UpdateSubscriptionAddonParams struct {
	Id                 string             `json:"id,omitempty"`
	Quantity           *int32             `json:"quantity,omitempty"`
	UnitPrice          *int32             `json:"unit_price,omitempty"`
	BillingCycles      *int32             `json:"billing_cycles,omitempty"`
	QuantityInDecimal  string             `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string             `json:"unit_price_in_decimal,omitempty"`
	TrialEnd           *int64             `json:"trial_end,omitempty"`
	ProrationType      enum.ProrationType `json:"proration_type,omitempty"`
}
type UpdateSubscriptionEventBasedAddonParams struct {
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
type UpdateSubscriptionBillingAddressParams struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type UpdateSubscriptionShippingAddressParams struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type UpdateSubscriptionCustomerParams struct {
	VatNumber        string          `json:"vat_number,omitempty"`
	VatNumberPrefix  string          `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool           `json:"registered_for_gst,omitempty"`
	Taxability       enum.Taxability `json:"taxability,omitempty"`
}
type UpdateSubscriptionForItemsRequestParams struct {
	Subscription           *UpdateSubscriptionForItemsSubscriptionParams       `json:"subscription,omitempty"`
	SubscriptionItems      []*UpdateSubscriptionForItemsSubscriptionItemParams `json:"subscription_items,omitempty"`
	MandatoryItemsToRemove []string                                            `json:"mandatory_items_to_remove,omitempty"`
	ReplaceItemsList       *bool                                               `json:"replace_items_list,omitempty"`
	Discounts              []*UpdateSubscriptionForItemsDiscountParams         `json:"discounts,omitempty"`
	ItemTiers              []*UpdateSubscriptionForItemsItemTierParams         `json:"item_tiers,omitempty"`
	InvoiceDate            *int64                                              `json:"invoice_date,omitempty"`
	BillingCycles          *int32                                              `json:"billing_cycles,omitempty"`
	TermsToCharge          *int32                                              `json:"terms_to_charge,omitempty"`
	ReactivateFrom         *int64                                              `json:"reactivate_from,omitempty"`
	BillingAlignmentMode   enum.BillingAlignmentMode                           `json:"billing_alignment_mode,omitempty"`
	CouponIds              []string                                            `json:"coupon_ids,omitempty"`
	ReplaceCouponList      *bool                                               `json:"replace_coupon_list,omitempty"`
	Prorate                *bool                                               `json:"prorate,omitempty"`
	EndOfTerm              *bool                                               `json:"end_of_term,omitempty"`
	ForceTermReset         *bool                                               `json:"force_term_reset,omitempty"`
	Reactivate             *bool                                               `json:"reactivate,omitempty"`
	IncludeDelayedCharges  *bool                                               `json:"include_delayed_charges,omitempty"`
	UseExistingBalances    *bool                                               `json:"use_existing_balances,omitempty"`
	BillingAddress         *UpdateSubscriptionForItemsBillingAddressParams     `json:"billing_address,omitempty"`
	ShippingAddress        *UpdateSubscriptionForItemsShippingAddressParams    `json:"shipping_address,omitempty"`
	Customer               *UpdateSubscriptionForItemsCustomerParams           `json:"customer,omitempty"`
	InvoiceImmediately     *bool                                               `json:"invoice_immediately,omitempty"`
}
type UpdateSubscriptionForItemsSubscriptionParams struct {
	Id                   string                    `json:"id"`
	SetupFee             *int32                    `json:"setup_fee,omitempty"`
	StartDate            *int64                    `json:"start_date,omitempty"`
	TrialEnd             *int64                    `json:"trial_end,omitempty"`
	Coupon               string                    `json:"coupon,omitempty"`
	AutoCollection       enum.AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	FreePeriod           *int32                    `json:"free_period,omitempty"`
	FreePeriodUnit       enum.FreePeriodUnit       `json:"free_period_unit,omitempty"`
	TrialEndAction       enum.TrialEndAction       `json:"trial_end_action,omitempty"`
}
type UpdateSubscriptionForItemsSubscriptionItemParams struct {
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
type UpdateSubscriptionForItemsDiscountParams struct {
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
type UpdateSubscriptionForItemsItemTierParams struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type UpdateSubscriptionForItemsBillingAddressParams struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type UpdateSubscriptionForItemsShippingAddressParams struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type UpdateSubscriptionForItemsCustomerParams struct {
	VatNumber        string          `json:"vat_number,omitempty"`
	VatNumberPrefix  string          `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool           `json:"registered_for_gst,omitempty"`
	Taxability       enum.Taxability `json:"taxability,omitempty"`
}
type RenewalEstimateRequestParams struct {
	IncludeDelayedCharges       *bool `json:"include_delayed_charges,omitempty"`
	UseExistingBalances         *bool `json:"use_existing_balances,omitempty"`
	IgnoreScheduledCancellation *bool `json:"ignore_scheduled_cancellation,omitempty"`
	IgnoreScheduledChanges      *bool `json:"ignore_scheduled_changes,omitempty"`
}
type AdvanceInvoiceEstimateRequestParams struct {
	TermsToCharge         *int32                                               `json:"terms_to_charge,omitempty"`
	SpecificDatesSchedule []*AdvanceInvoiceEstimateSpecificDatesScheduleParams `json:"specific_dates_schedule,omitempty"`
	FixedIntervalSchedule *AdvanceInvoiceEstimateFixedIntervalScheduleParams   `json:"fixed_interval_schedule,omitempty"`
	InvoiceImmediately    *bool                                                `json:"invoice_immediately,omitempty"`
	ScheduleType          enum.ScheduleType                                    `json:"schedule_type,omitempty"`
}
type AdvanceInvoiceEstimateSpecificDatesScheduleParams struct {
	TermsToCharge *int32 `json:"terms_to_charge,omitempty"`
	Date          *int64 `json:"date,omitempty"`
}
type AdvanceInvoiceEstimateFixedIntervalScheduleParams struct {
	NumberOfOccurrences *int32             `json:"number_of_occurrences,omitempty"`
	DaysBeforeRenewal   *int32             `json:"days_before_renewal,omitempty"`
	EndScheduleOn       enum.EndScheduleOn `json:"end_schedule_on,omitempty"`
	EndDate             *int64             `json:"end_date,omitempty"`
}
type RegenerateInvoiceEstimateRequestParams struct {
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
	Prorate            *bool  `json:"prorate,omitempty"`
	InvoiceImmediately *bool  `json:"invoice_immediately,omitempty"`
}
type ChangeTermEndRequestParams struct {
	TermEndsAt         *int64 `json:"term_ends_at"`
	Prorate            *bool  `json:"prorate,omitempty"`
	InvoiceImmediately *bool  `json:"invoice_immediately,omitempty"`
}
type CancelSubscriptionRequestParams struct {
	EndOfTerm                         *bool                                      `json:"end_of_term,omitempty"`
	CancelAt                          *int64                                     `json:"cancel_at,omitempty"`
	CreditOptionForCurrentTermCharges enum.CreditOptionForCurrentTermCharges     `json:"credit_option_for_current_term_charges,omitempty"`
	UnbilledChargesOption             enum.UnbilledChargesOption                 `json:"unbilled_charges_option,omitempty"`
	AccountReceivablesHandling        enum.AccountReceivablesHandling            `json:"account_receivables_handling,omitempty"`
	RefundableCreditsHandling         enum.RefundableCreditsHandling             `json:"refundable_credits_handling,omitempty"`
	ContractTermCancelOption          enum.ContractTermCancelOption              `json:"contract_term_cancel_option,omitempty"`
	InvoiceDate                       *int64                                     `json:"invoice_date,omitempty"`
	EventBasedAddons                  []*CancelSubscriptionEventBasedAddonParams `json:"event_based_addons,omitempty"`
	CancelReasonCode                  string                                     `json:"cancel_reason_code,omitempty"`
}
type CancelSubscriptionEventBasedAddonParams struct {
	Id                  string `json:"id,omitempty"`
	Quantity            *int32 `json:"quantity,omitempty"`
	UnitPrice           *int32 `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32 `json:"service_period_in_days,omitempty"`
}
type CancelSubscriptionForItemsRequestParams struct {
	EndOfTerm                         *bool                                               `json:"end_of_term,omitempty"`
	CancelAt                          *int64                                              `json:"cancel_at,omitempty"`
	CreditOptionForCurrentTermCharges enum.CreditOptionForCurrentTermCharges              `json:"credit_option_for_current_term_charges,omitempty"`
	UnbilledChargesOption             enum.UnbilledChargesOption                          `json:"unbilled_charges_option,omitempty"`
	AccountReceivablesHandling        enum.AccountReceivablesHandling                     `json:"account_receivables_handling,omitempty"`
	RefundableCreditsHandling         enum.RefundableCreditsHandling                      `json:"refundable_credits_handling,omitempty"`
	ContractTermCancelOption          enum.ContractTermCancelOption                       `json:"contract_term_cancel_option,omitempty"`
	InvoiceDate                       *int64                                              `json:"invoice_date,omitempty"`
	SubscriptionItems                 []*CancelSubscriptionForItemsSubscriptionItemParams `json:"subscription_items,omitempty"`
	CancelReasonCode                  string                                              `json:"cancel_reason_code,omitempty"`
}
type CancelSubscriptionForItemsSubscriptionItemParams struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodDays  *int32 `json:"service_period_days,omitempty"`
}
type PauseSubscriptionRequestParams struct {
	PauseOption             enum.PauseOption                     `json:"pause_option,omitempty"`
	Subscription            *PauseSubscriptionSubscriptionParams `json:"subscription,omitempty"`
	UnbilledChargesHandling enum.UnbilledChargesHandling         `json:"unbilled_charges_handling,omitempty"`
}
type PauseSubscriptionSubscriptionParams struct {
	PauseDate         *int64 `json:"pause_date,omitempty"`
	ResumeDate        *int64 `json:"resume_date,omitempty"`
	SkipBillingCycles *int32 `json:"skip_billing_cycles,omitempty"`
}
type ResumeSubscriptionRequestParams struct {
	ResumeOption    enum.ResumeOption                     `json:"resume_option,omitempty"`
	Subscription    *ResumeSubscriptionSubscriptionParams `json:"subscription,omitempty"`
	ChargesHandling enum.ChargesHandling                  `json:"charges_handling,omitempty"`
}
type ResumeSubscriptionSubscriptionParams struct {
	ResumeDate *int64 `json:"resume_date,omitempty"`
}
type GiftSubscriptionRequestParams struct {
	Gift            *GiftSubscriptionGiftParams            `json:"gift,omitempty"`
	Gifter          *GiftSubscriptionGifterParams          `json:"gifter,omitempty"`
	GiftReceiver    *GiftSubscriptionGiftReceiverParams    `json:"gift_receiver,omitempty"`
	CouponIds       []string                               `json:"coupon_ids,omitempty"`
	PaymentIntent   *GiftSubscriptionPaymentIntentParams   `json:"payment_intent,omitempty"`
	ShippingAddress *GiftSubscriptionShippingAddressParams `json:"shipping_address,omitempty"`
	Subscription    *GiftSubscriptionSubscriptionParams    `json:"subscription,omitempty"`
	Addons          []*GiftSubscriptionAddonParams         `json:"addons,omitempty"`
}
type GiftSubscriptionGiftParams struct {
	ScheduledAt     *int64 `json:"scheduled_at,omitempty"`
	AutoClaim       *bool  `json:"auto_claim,omitempty"`
	NoExpiry        *bool  `json:"no_expiry,omitempty"`
	ClaimExpiryDate *int64 `json:"claim_expiry_date,omitempty"`
}
type GiftSubscriptionGifterParams struct {
	CustomerId   string `json:"customer_id"`
	Signature    string `json:"signature"`
	Note         string `json:"note,omitempty"`
	PaymentSrcId string `json:"payment_src_id,omitempty"`
}
type GiftSubscriptionGiftReceiverParams struct {
	CustomerId string `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
}
type GiftSubscriptionPaymentIntentParams struct {
	Id                    string                              `json:"id,omitempty"`
	GatewayAccountId      string                              `json:"gateway_account_id,omitempty"`
	GwToken               string                              `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntentEnum.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                              `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                              `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}              `json:"additional_information,omitempty"`
}
type GiftSubscriptionShippingAddressParams struct {
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
type GiftSubscriptionSubscriptionParams struct {
	PlanId                string `json:"plan_id"`
	PlanQuantity          *int32 `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal string `json:"plan_quantity_in_decimal,omitempty"`
}
type GiftSubscriptionAddonParams struct {
	Id                string `json:"id,omitempty"`
	Quantity          *int32 `json:"quantity,omitempty"`
	QuantityInDecimal string `json:"quantity_in_decimal,omitempty"`
}
type GiftSubscriptionForItemsRequestParams struct {
	Gift              *GiftSubscriptionForItemsGiftParams               `json:"gift,omitempty"`
	Gifter            *GiftSubscriptionForItemsGifterParams             `json:"gifter,omitempty"`
	GiftReceiver      *GiftSubscriptionForItemsGiftReceiverParams       `json:"gift_receiver,omitempty"`
	CouponIds         []string                                          `json:"coupon_ids,omitempty"`
	PaymentIntent     *GiftSubscriptionForItemsPaymentIntentParams      `json:"payment_intent,omitempty"`
	ShippingAddress   *GiftSubscriptionForItemsShippingAddressParams    `json:"shipping_address,omitempty"`
	SubscriptionItems []*GiftSubscriptionForItemsSubscriptionItemParams `json:"subscription_items,omitempty"`
}
type GiftSubscriptionForItemsGiftParams struct {
	ScheduledAt     *int64 `json:"scheduled_at,omitempty"`
	AutoClaim       *bool  `json:"auto_claim,omitempty"`
	NoExpiry        *bool  `json:"no_expiry,omitempty"`
	ClaimExpiryDate *int64 `json:"claim_expiry_date,omitempty"`
}
type GiftSubscriptionForItemsGifterParams struct {
	CustomerId   string `json:"customer_id"`
	Signature    string `json:"signature"`
	Note         string `json:"note,omitempty"`
	PaymentSrcId string `json:"payment_src_id,omitempty"`
}
type GiftSubscriptionForItemsGiftReceiverParams struct {
	CustomerId string `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
}
type GiftSubscriptionForItemsPaymentIntentParams struct {
	Id                    string                              `json:"id,omitempty"`
	GatewayAccountId      string                              `json:"gateway_account_id,omitempty"`
	GwToken               string                              `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntentEnum.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                              `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                              `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}              `json:"additional_information,omitempty"`
}
type GiftSubscriptionForItemsShippingAddressParams struct {
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
type GiftSubscriptionForItemsSubscriptionItemParams struct {
	ItemPriceId       string `json:"item_price_id,omitempty"`
	Quantity          *int32 `json:"quantity,omitempty"`
	QuantityInDecimal string `json:"quantity_in_decimal,omitempty"`
}
type CreateInvoiceRequestParams struct {
	Invoice                    *CreateInvoiceInvoiceParams         `json:"invoice,omitempty"`
	CurrencyCode               string                              `json:"currency_code,omitempty"`
	Addons                     []*CreateInvoiceAddonParams         `json:"addons,omitempty"`
	Charges                    []*CreateInvoiceChargeParams        `json:"charges,omitempty"`
	InvoiceNote                string                              `json:"invoice_note,omitempty"`
	RemoveGeneralNote          *bool                               `json:"remove_general_note,omitempty"`
	NotesToRemove              []*CreateInvoiceNotesToRemoveParams `json:"notes_to_remove,omitempty"`
	Coupon                     string                              `json:"coupon,omitempty"`
	CouponIds                  []string                            `json:"coupon_ids,omitempty"`
	AuthorizationTransactionId string                              `json:"authorization_transaction_id,omitempty"`
	PaymentSourceId            string                              `json:"payment_source_id,omitempty"`
	AutoCollection             enum.AutoCollection                 `json:"auto_collection,omitempty"`
	InvoiceDate                *int64                              `json:"invoice_date,omitempty"`
	ShippingAddress            *CreateInvoiceShippingAddressParams `json:"shipping_address,omitempty"`
}
type CreateInvoiceInvoiceParams struct {
	CustomerId     string `json:"customer_id,omitempty"`
	SubscriptionId string `json:"subscription_id,omitempty"`
	PoNumber       string `json:"po_number,omitempty"`
}
type CreateInvoiceAddonParams struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type CreateInvoiceChargeParams struct {
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
type CreateInvoiceNotesToRemoveParams struct {
	EntityType enum.EntityType `json:"entity_type,omitempty"`
	EntityId   string          `json:"entity_id,omitempty"`
}
type CreateInvoiceShippingAddressParams struct {
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
type CreateInvoiceForItemsRequestParams struct {
	Invoice                    *CreateInvoiceForItemsInvoiceParams         `json:"invoice,omitempty"`
	CurrencyCode               string                                      `json:"currency_code,omitempty"`
	ItemPrices                 []*CreateInvoiceForItemsItemPriceParams     `json:"item_prices,omitempty"`
	ItemTiers                  []*CreateInvoiceForItemsItemTierParams      `json:"item_tiers,omitempty"`
	Charges                    []*CreateInvoiceForItemsChargeParams        `json:"charges,omitempty"`
	InvoiceNote                string                                      `json:"invoice_note,omitempty"`
	RemoveGeneralNote          *bool                                       `json:"remove_general_note,omitempty"`
	NotesToRemove              []*CreateInvoiceForItemsNotesToRemoveParams `json:"notes_to_remove,omitempty"`
	Coupon                     string                                      `json:"coupon,omitempty"`
	CouponIds                  []string                                    `json:"coupon_ids,omitempty"`
	AuthorizationTransactionId string                                      `json:"authorization_transaction_id,omitempty"`
	PaymentSourceId            string                                      `json:"payment_source_id,omitempty"`
	AutoCollection             enum.AutoCollection                         `json:"auto_collection,omitempty"`
	Discounts                  []*CreateInvoiceForItemsDiscountParams      `json:"discounts,omitempty"`
	ShippingAddress            *CreateInvoiceForItemsShippingAddressParams `json:"shipping_address,omitempty"`
	InvoiceDate                *int64                                      `json:"invoice_date,omitempty"`
}
type CreateInvoiceForItemsInvoiceParams struct {
	CustomerId     string `json:"customer_id,omitempty"`
	SubscriptionId string `json:"subscription_id,omitempty"`
	PoNumber       string `json:"po_number,omitempty"`
}
type CreateInvoiceForItemsItemPriceParams struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type CreateInvoiceForItemsItemTierParams struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type CreateInvoiceForItemsChargeParams struct {
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
type CreateInvoiceForItemsNotesToRemoveParams struct {
	EntityType enum.EntityType `json:"entity_type,omitempty"`
	EntityId   string          `json:"entity_id,omitempty"`
}
type CreateInvoiceForItemsDiscountParams struct {
	Percentage  *float64     `json:"percentage,omitempty"`
	Amount      *int32       `json:"amount,omitempty"`
	ApplyOn     enum.ApplyOn `json:"apply_on"`
	ItemPriceId string       `json:"item_price_id,omitempty"`
}
type CreateInvoiceForItemsShippingAddressParams struct {
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
