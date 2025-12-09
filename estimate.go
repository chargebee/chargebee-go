package chargebee

type PaymentScheduleEstimateEntityType string

const (
	PaymentScheduleEstimateEntityTypeInvoice PaymentScheduleEstimateEntityType = "invoice"
)

type UnbilledChargeEntityType string

const (
	UnbilledChargeEntityTypeAdhoc           UnbilledChargeEntityType = "adhoc"
	UnbilledChargeEntityTypePlanItemPrice   UnbilledChargeEntityType = "plan_item_price"
	UnbilledChargeEntityTypeAddonItemPrice  UnbilledChargeEntityType = "addon_item_price"
	UnbilledChargeEntityTypeChargeItemPrice UnbilledChargeEntityType = "charge_item_price"
	UnbilledChargeEntityTypePlanSetup       UnbilledChargeEntityType = "plan_setup"
	UnbilledChargeEntityTypePlan            UnbilledChargeEntityType = "plan"
	UnbilledChargeEntityTypeAddon           UnbilledChargeEntityType = "addon"
)

type Estimate struct {
	CreatedAt                int64                                              `json:"created_at"`
	SubscriptionEstimate     *subscriptionestimate.SubscriptionEstimate         `json:"subscription_estimate"`
	SubscriptionEstimates    []*subscriptionestimate.SubscriptionEstimate       `json:"subscription_estimates"`
	InvoiceEstimate          *invoiceestimate.InvoiceEstimate                   `json:"invoice_estimate"`
	InvoiceEstimates         []*invoiceestimate.InvoiceEstimate                 `json:"invoice_estimates"`
	PaymentScheduleEstimates []*paymentscheduleestimate.PaymentScheduleEstimate `json:"payment_schedule_estimates"`
	NextInvoiceEstimate      *invoiceestimate.InvoiceEstimate                   `json:"next_invoice_estimate"`
	CreditNoteEstimates      []*creditnoteestimate.CreditNoteEstimate           `json:"credit_note_estimates"`
	UnbilledChargeEstimates  []*unbilledcharge.UnbilledCharge                   `json:"unbilled_charge_estimates"`
	Object                   string                                             `json:"object"`
}
type CreateSubscriptionRequest struct {
	Subscription            *CreateSubscriptionSubscription        `json:"subscription,omitempty"`
	BillingCycles           *int32                                 `json:"billing_cycles,omitempty"`
	Addons                  []*CreateSubscriptionAddon             `json:"addons,omitempty"`
	EventBasedAddons        []*CreateSubscriptionEventBasedAddon   `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove []string                               `json:"mandatory_addons_to_remove,omitempty"`
	TermsToCharge           *int32                                 `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode              `json:"billing_alignment_mode,omitempty"`
	CouponIds               []string                               `json:"coupon_ids,omitempty"`
	BillingAddress          *CreateSubscriptionBillingAddress      `json:"billing_address,omitempty"`
	ShippingAddress         *CreateSubscriptionShippingAddress     `json:"shipping_address,omitempty"`
	Customer                *CreateSubscriptionCustomer            `json:"customer,omitempty"`
	InvoiceImmediately      *bool                                  `json:"invoice_immediately,omitempty"`
	InvoiceDate             *int64                                 `json:"invoice_date,omitempty"`
	ContractTerm            *CreateSubscriptionContractTerm        `json:"contract_term,omitempty"`
	TaxProvidersFields      []*CreateSubscriptionTaxProvidersField `json:"tax_providers_fields,omitempty"`
	ClientProfileId         string                                 `json:"client_profile_id,omitempty"`
}
type CreateSubscriptionSubscription struct {
	Id                                string                    `json:"id,omitempty"`
	PlanId                            string                    `json:"plan_id"`
	PlanQuantity                      *int32                    `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                    `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int64                    `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                    `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int64                    `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	Coupon                            string                    `json:"coupon,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	FreePeriod                        *int32                    `json:"free_period,omitempty"`
	FreePeriodUnit                    enum.FreePeriodUnit       `json:"free_period_unit,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	TrialEndAction                    enum.TrialEndAction       `json:"trial_end_action,omitempty"`
}
type CreateSubscriptionAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}
type CreateSubscriptionEventBasedAddon struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int64        `json:"unit_price,omitempty"`
	QuantityInDecimal   string        `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string        `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
}
type CreateSubscriptionBillingAddress struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateSubscriptionShippingAddress struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateSubscriptionCustomer struct {
	VatNumber        string                   `json:"vat_number,omitempty"`
	VatNumberPrefix  string                   `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool                    `json:"registered_for_gst,omitempty"`
	Taxability       enum.Taxability          `json:"taxability,omitempty"`
	EntityCode       enum.EntityCode          `json:"entity_code,omitempty"`
	ExemptNumber     string                   `json:"exempt_number,omitempty"`
	ExemptionDetails []map[string]interface{} `json:"exemption_details,omitempty"`
	CustomerType     enum.CustomerType        `json:"customer_type,omitempty"`
}
type CreateSubscriptionContractTerm struct {
	ActionAtTermEnd          contractTerm.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                       `json:"cancellation_cutoff_period,omitempty"`
}
type CreateSubscriptionTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type CreateSubItemEstimateRequest struct {
	Subscription           *CreateSubItemEstimateSubscription        `json:"subscription,omitempty"`
	BillingCycles          *int32                                    `json:"billing_cycles,omitempty"`
	SubscriptionItems      []*CreateSubItemEstimateSubscriptionItem  `json:"subscription_items,omitempty"`
	Discounts              []*CreateSubItemEstimateDiscount          `json:"discounts,omitempty"`
	MandatoryItemsToRemove []string                                  `json:"mandatory_items_to_remove,omitempty"`
	ItemTiers              []*CreateSubItemEstimateItemTier          `json:"item_tiers,omitempty"`
	TermsToCharge          *int32                                    `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode   enum.BillingAlignmentMode                 `json:"billing_alignment_mode,omitempty"`
	CouponIds              []string                                  `json:"coupon_ids,omitempty"`
	BillingAddress         *CreateSubItemEstimateBillingAddress      `json:"billing_address,omitempty"`
	ShippingAddress        *CreateSubItemEstimateShippingAddress     `json:"shipping_address,omitempty"`
	Customer               *CreateSubItemEstimateCustomer            `json:"customer,omitempty"`
	InvoiceImmediately     *bool                                     `json:"invoice_immediately,omitempty"`
	InvoiceDate            *int64                                    `json:"invoice_date,omitempty"`
	ClientProfileId        string                                    `json:"client_profile_id,omitempty"`
	ContractTerm           *CreateSubItemEstimateContractTerm        `json:"contract_term,omitempty"`
	TaxProvidersFields     []*CreateSubItemEstimateTaxProvidersField `json:"tax_providers_fields,omitempty"`
}
type CreateSubItemEstimateSubscription struct {
	Id                                string                    `json:"id,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	SetupFee                          *int64                    `json:"setup_fee,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	Coupon                            string                    `json:"coupon,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	FreePeriod                        *int32                    `json:"free_period,omitempty"`
	FreePeriodUnit                    enum.FreePeriodUnit       `json:"free_period_unit,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	TrialEndAction                    enum.TrialEndAction       `json:"trial_end_action,omitempty"`
}
type CreateSubItemEstimateSubscriptionItem struct {
	ItemPriceId        string              `json:"item_price_id"`
	Quantity           *int32              `json:"quantity,omitempty"`
	QuantityInDecimal  string              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32              `json:"billing_cycles,omitempty"`
	TrialEnd           *int64              `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32              `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool               `json:"charge_once,omitempty"`
	ItemType           enum.ItemType       `json:"item_type,omitempty"`
	ChargeOnOption     enum.ChargeOnOption `json:"charge_on_option,omitempty"`
}
type CreateSubItemEstimateDiscount struct {
	ApplyOn       enum.ApplyOn      `json:"apply_on,omitempty"`
	DurationType  enum.DurationType `json:"duration_type"`
	Percentage    *float64          `json:"percentage,omitempty"`
	Amount        *int64            `json:"amount,omitempty"`
	Period        *int32            `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool             `json:"included_in_mrr,omitempty"`
	ItemPriceId   string            `json:"item_price_id,omitempty"`
	Quantity      *int32            `json:"quantity,omitempty"`
}
type CreateSubItemEstimateItemTier struct {
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
type CreateSubItemEstimateBillingAddress struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateSubItemEstimateShippingAddress struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateSubItemEstimateCustomer struct {
	VatNumber        string                   `json:"vat_number,omitempty"`
	VatNumberPrefix  string                   `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool                    `json:"registered_for_gst,omitempty"`
	Taxability       enum.Taxability          `json:"taxability,omitempty"`
	EntityCode       enum.EntityCode          `json:"entity_code,omitempty"`
	ExemptNumber     string                   `json:"exempt_number,omitempty"`
	ExemptionDetails []map[string]interface{} `json:"exemption_details,omitempty"`
	CustomerType     enum.CustomerType        `json:"customer_type,omitempty"`
}
type CreateSubItemEstimateContractTerm struct {
	ActionAtTermEnd          contractTerm.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	ContractStart            *int64                       `json:"contract_start,omitempty"`
	CancellationCutoffPeriod *int32                       `json:"cancellation_cutoff_period,omitempty"`
}
type CreateSubItemEstimateTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type CreateSubForCustomerEstimateRequest struct {
	UseExistingBalances     *bool                                          `json:"use_existing_balances,omitempty"`
	Subscription            *CreateSubForCustomerEstimateSubscription      `json:"subscription,omitempty"`
	InvoiceImmediately      *bool                                          `json:"invoice_immediately,omitempty"`
	BillingCycles           *int32                                         `json:"billing_cycles,omitempty"`
	Addons                  []*CreateSubForCustomerEstimateAddon           `json:"addons,omitempty"`
	EventBasedAddons        []*CreateSubForCustomerEstimateEventBasedAddon `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove []string                                       `json:"mandatory_addons_to_remove,omitempty"`
	TermsToCharge           *int32                                         `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode                      `json:"billing_alignment_mode,omitempty"`
	ShippingAddress         *CreateSubForCustomerEstimateShippingAddress   `json:"shipping_address,omitempty"`
	InvoiceDate             *int64                                         `json:"invoice_date,omitempty"`
	CouponIds               []string                                       `json:"coupon_ids,omitempty"`
	ContractTerm            *CreateSubForCustomerEstimateContractTerm      `json:"contract_term,omitempty"`
}
type CreateSubForCustomerEstimateSubscription struct {
	Id                                string                    `json:"id,omitempty"`
	PlanId                            string                    `json:"plan_id"`
	PlanQuantity                      *int32                    `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                    `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int64                    `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                    `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int64                    `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	FreePeriod                        *int32                    `json:"free_period,omitempty"`
	FreePeriodUnit                    enum.FreePeriodUnit       `json:"free_period_unit,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	TrialEndAction                    enum.TrialEndAction       `json:"trial_end_action,omitempty"`
}
type CreateSubForCustomerEstimateAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}
type CreateSubForCustomerEstimateEventBasedAddon struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int64        `json:"unit_price,omitempty"`
	QuantityInDecimal   string        `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string        `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
}
type CreateSubForCustomerEstimateShippingAddress struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateSubForCustomerEstimateContractTerm struct {
	ActionAtTermEnd          contractTerm.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                       `json:"cancellation_cutoff_period,omitempty"`
}
type CreateSubItemForCustomerEstimateRequest struct {
	UseExistingBalances    *bool                                               `json:"use_existing_balances,omitempty"`
	Subscription           *CreateSubItemForCustomerEstimateSubscription       `json:"subscription,omitempty"`
	InvoiceImmediately     *bool                                               `json:"invoice_immediately,omitempty"`
	BillingCycles          *int32                                              `json:"billing_cycles,omitempty"`
	SubscriptionItems      []*CreateSubItemForCustomerEstimateSubscriptionItem `json:"subscription_items,omitempty"`
	Discounts              []*CreateSubItemForCustomerEstimateDiscount         `json:"discounts,omitempty"`
	MandatoryItemsToRemove []string                                            `json:"mandatory_items_to_remove,omitempty"`
	ItemTiers              []*CreateSubItemForCustomerEstimateItemTier         `json:"item_tiers,omitempty"`
	TermsToCharge          *int32                                              `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode   enum.BillingAlignmentMode                           `json:"billing_alignment_mode,omitempty"`
	ShippingAddress        *CreateSubItemForCustomerEstimateShippingAddress    `json:"shipping_address,omitempty"`
	BillingAddress         *CreateSubItemForCustomerEstimateBillingAddress     `json:"billing_address,omitempty"`
	InvoiceDate            *int64                                              `json:"invoice_date,omitempty"`
	CouponIds              []string                                            `json:"coupon_ids,omitempty"`
	ContractTerm           *CreateSubItemForCustomerEstimateContractTerm       `json:"contract_term,omitempty"`
	BillingOverride        *CreateSubItemForCustomerEstimateBillingOverride    `json:"billing_override,omitempty"`
}
type CreateSubItemForCustomerEstimateSubscription struct {
	Id                                string                    `json:"id,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	SetupFee                          *int64                    `json:"setup_fee,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	FreePeriod                        *int32                    `json:"free_period,omitempty"`
	FreePeriodUnit                    enum.FreePeriodUnit       `json:"free_period_unit,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	TrialEndAction                    enum.TrialEndAction       `json:"trial_end_action,omitempty"`
}
type CreateSubItemForCustomerEstimateSubscriptionItem struct {
	ItemPriceId        string              `json:"item_price_id"`
	Quantity           *int32              `json:"quantity,omitempty"`
	QuantityInDecimal  string              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32              `json:"billing_cycles,omitempty"`
	TrialEnd           *int64              `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32              `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool               `json:"charge_once,omitempty"`
	ItemType           enum.ItemType       `json:"item_type,omitempty"`
	ChargeOnOption     enum.ChargeOnOption `json:"charge_on_option,omitempty"`
}
type CreateSubItemForCustomerEstimateDiscount struct {
	ApplyOn       enum.ApplyOn      `json:"apply_on,omitempty"`
	DurationType  enum.DurationType `json:"duration_type"`
	Percentage    *float64          `json:"percentage,omitempty"`
	Amount        *int64            `json:"amount,omitempty"`
	Period        *int32            `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool             `json:"included_in_mrr,omitempty"`
	ItemPriceId   string            `json:"item_price_id,omitempty"`
	Quantity      *int32            `json:"quantity,omitempty"`
}
type CreateSubItemForCustomerEstimateItemTier struct {
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
type CreateSubItemForCustomerEstimateShippingAddress struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateSubItemForCustomerEstimateBillingAddress struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateSubItemForCustomerEstimateContractTerm struct {
	ActionAtTermEnd          contractTerm.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	ContractStart            *int64                       `json:"contract_start,omitempty"`
	CancellationCutoffPeriod *int32                       `json:"cancellation_cutoff_period,omitempty"`
}
type CreateSubItemForCustomerEstimateBillingOverride struct {
	MaxExcessPaymentUsage     *int64 `json:"max_excess_payment_usage,omitempty"`
	MaxRefundableCreditsUsage *int64 `json:"max_refundable_credits_usage,omitempty"`
}
type UpdateSubscriptionRequest struct {
	Subscription            *UpdateSubscriptionSubscription      `json:"subscription,omitempty"`
	ChangesScheduledAt      *int64                               `json:"changes_scheduled_at,omitempty"`
	ChangeOption            enum.ChangeOption                    `json:"change_option,omitempty"`
	Addons                  []*UpdateSubscriptionAddon           `json:"addons,omitempty"`
	EventBasedAddons        []*UpdateSubscriptionEventBasedAddon `json:"event_based_addons,omitempty"`
	ReplaceAddonList        *bool                                `json:"replace_addon_list,omitempty"`
	MandatoryAddonsToRemove []string                             `json:"mandatory_addons_to_remove,omitempty"`
	InvoiceDate             *int64                               `json:"invoice_date,omitempty"`
	BillingCycles           *int32                               `json:"billing_cycles,omitempty"`
	TermsToCharge           *int32                               `json:"terms_to_charge,omitempty"`
	ReactivateFrom          *int64                               `json:"reactivate_from,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode            `json:"billing_alignment_mode,omitempty"`
	CouponIds               []string                             `json:"coupon_ids,omitempty"`
	ReplaceCouponList       *bool                                `json:"replace_coupon_list,omitempty"`
	Prorate                 *bool                                `json:"prorate,omitempty"`
	EndOfTerm               *bool                                `json:"end_of_term,omitempty"`
	ForceTermReset          *bool                                `json:"force_term_reset,omitempty"`
	Reactivate              *bool                                `json:"reactivate,omitempty"`
	IncludeDelayedCharges   *bool                                `json:"include_delayed_charges,omitempty"`
	UseExistingBalances     *bool                                `json:"use_existing_balances,omitempty"`
	BillingAddress          *UpdateSubscriptionBillingAddress    `json:"billing_address,omitempty"`
	ShippingAddress         *UpdateSubscriptionShippingAddress   `json:"shipping_address,omitempty"`
	Customer                *UpdateSubscriptionCustomer          `json:"customer,omitempty"`
	InvoiceImmediately      *bool                                `json:"invoice_immediately,omitempty"`
}
type UpdateSubscriptionSubscription struct {
	Id                     string                    `json:"id"`
	PlanId                 string                    `json:"plan_id,omitempty"`
	PlanQuantity           *int32                    `json:"plan_quantity,omitempty"`
	PlanUnitPrice          *int64                    `json:"plan_unit_price,omitempty"`
	SetupFee               *int64                    `json:"setup_fee,omitempty"`
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
type UpdateSubscriptionAddon struct {
	Id                 string             `json:"id,omitempty"`
	Quantity           *int32             `json:"quantity,omitempty"`
	UnitPrice          *int64             `json:"unit_price,omitempty"`
	BillingCycles      *int32             `json:"billing_cycles,omitempty"`
	QuantityInDecimal  string             `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string             `json:"unit_price_in_decimal,omitempty"`
	TrialEnd           *int64             `json:"trial_end,omitempty"`
	ProrationType      enum.ProrationType `json:"proration_type,omitempty"`
}
type UpdateSubscriptionEventBasedAddon struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int64        `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
	QuantityInDecimal   string        `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string        `json:"unit_price_in_decimal,omitempty"`
}
type UpdateSubscriptionBillingAddress struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type UpdateSubscriptionShippingAddress struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type UpdateSubscriptionCustomer struct {
	VatNumber        string          `json:"vat_number,omitempty"`
	VatNumberPrefix  string          `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool           `json:"registered_for_gst,omitempty"`
	Taxability       enum.Taxability `json:"taxability,omitempty"`
}
type UpdateSubscriptionForItemsRequest struct {
	Subscription           *UpdateSubscriptionForItemsSubscription       `json:"subscription,omitempty"`
	ChangesScheduledAt     *int64                                        `json:"changes_scheduled_at,omitempty"`
	ChangeOption           enum.ChangeOption                             `json:"change_option,omitempty"`
	SubscriptionItems      []*UpdateSubscriptionForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	MandatoryItemsToRemove []string                                      `json:"mandatory_items_to_remove,omitempty"`
	ReplaceItemsList       *bool                                         `json:"replace_items_list,omitempty"`
	Discounts              []*UpdateSubscriptionForItemsDiscount         `json:"discounts,omitempty"`
	ItemTiers              []*UpdateSubscriptionForItemsItemTier         `json:"item_tiers,omitempty"`
	InvoiceDate            *int64                                        `json:"invoice_date,omitempty"`
	BillingCycles          *int32                                        `json:"billing_cycles,omitempty"`
	TermsToCharge          *int32                                        `json:"terms_to_charge,omitempty"`
	ReactivateFrom         *int64                                        `json:"reactivate_from,omitempty"`
	BillingAlignmentMode   enum.BillingAlignmentMode                     `json:"billing_alignment_mode,omitempty"`
	CouponIds              []string                                      `json:"coupon_ids,omitempty"`
	ReplaceCouponList      *bool                                         `json:"replace_coupon_list,omitempty"`
	Prorate                *bool                                         `json:"prorate,omitempty"`
	EndOfTerm              *bool                                         `json:"end_of_term,omitempty"`
	ForceTermReset         *bool                                         `json:"force_term_reset,omitempty"`
	Reactivate             *bool                                         `json:"reactivate,omitempty"`
	IncludeDelayedCharges  *bool                                         `json:"include_delayed_charges,omitempty"`
	UseExistingBalances    *bool                                         `json:"use_existing_balances,omitempty"`
	BillingAddress         *UpdateSubscriptionForItemsBillingAddress     `json:"billing_address,omitempty"`
	ShippingAddress        *UpdateSubscriptionForItemsShippingAddress    `json:"shipping_address,omitempty"`
	Customer               *UpdateSubscriptionForItemsCustomer           `json:"customer,omitempty"`
	InvoiceImmediately     *bool                                         `json:"invoice_immediately,omitempty"`
	InvoiceUsages          *bool                                         `json:"invoice_usages,omitempty"`
	BillingOverride        *UpdateSubscriptionForItemsBillingOverride    `json:"billing_override,omitempty"`
}
type UpdateSubscriptionForItemsSubscription struct {
	Id                   string                    `json:"id"`
	SetupFee             *int64                    `json:"setup_fee,omitempty"`
	StartDate            *int64                    `json:"start_date,omitempty"`
	TrialEnd             *int64                    `json:"trial_end,omitempty"`
	Coupon               string                    `json:"coupon,omitempty"`
	AutoCollection       enum.AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	FreePeriod           *int32                    `json:"free_period,omitempty"`
	FreePeriodUnit       enum.FreePeriodUnit       `json:"free_period_unit,omitempty"`
	TrialEndAction       enum.TrialEndAction       `json:"trial_end_action,omitempty"`
}
type UpdateSubscriptionForItemsSubscriptionItem struct {
	ItemPriceId        string              `json:"item_price_id"`
	Quantity           *int32              `json:"quantity,omitempty"`
	QuantityInDecimal  string              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32              `json:"billing_cycles,omitempty"`
	TrialEnd           *int64              `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32              `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool               `json:"charge_once,omitempty"`
	ChargeOnOption     enum.ChargeOnOption `json:"charge_on_option,omitempty"`
	ItemType           enum.ItemType       `json:"item_type,omitempty"`
	ProrationType      enum.ProrationType  `json:"proration_type,omitempty"`
}
type UpdateSubscriptionForItemsDiscount struct {
	ApplyOn       enum.ApplyOn       `json:"apply_on,omitempty"`
	DurationType  enum.DurationType  `json:"duration_type"`
	Percentage    *float64           `json:"percentage,omitempty"`
	Amount        *int64             `json:"amount,omitempty"`
	Period        *int32             `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit    `json:"period_unit,omitempty"`
	IncludedInMrr *bool              `json:"included_in_mrr,omitempty"`
	ItemPriceId   string             `json:"item_price_id,omitempty"`
	Quantity      *int32             `json:"quantity,omitempty"`
	OperationType enum.OperationType `json:"operation_type"`
	Id            string             `json:"id,omitempty"`
}
type UpdateSubscriptionForItemsItemTier struct {
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
type UpdateSubscriptionForItemsBillingAddress struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type UpdateSubscriptionForItemsShippingAddress struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type UpdateSubscriptionForItemsCustomer struct {
	VatNumber        string          `json:"vat_number,omitempty"`
	VatNumberPrefix  string          `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool           `json:"registered_for_gst,omitempty"`
	Taxability       enum.Taxability `json:"taxability,omitempty"`
}
type UpdateSubscriptionForItemsBillingOverride struct {
	MaxExcessPaymentUsage     *int64 `json:"max_excess_payment_usage,omitempty"`
	MaxRefundableCreditsUsage *int64 `json:"max_refundable_credits_usage,omitempty"`
}
type RenewalEstimateRequest struct {
	IncludeDelayedCharges       *bool `json:"include_delayed_charges,omitempty"`
	UseExistingBalances         *bool `json:"use_existing_balances,omitempty"`
	IgnoreScheduledCancellation *bool `json:"ignore_scheduled_cancellation,omitempty"`
	IgnoreScheduledChanges      *bool `json:"ignore_scheduled_changes,omitempty"`
}
type AdvanceInvoiceEstimateRequest struct {
	TermsToCharge         *int32                                         `json:"terms_to_charge,omitempty"`
	SpecificDatesSchedule []*AdvanceInvoiceEstimateSpecificDatesSchedule `json:"specific_dates_schedule,omitempty"`
	FixedIntervalSchedule *AdvanceInvoiceEstimateFixedIntervalSchedule   `json:"fixed_interval_schedule,omitempty"`
	InvoiceImmediately    *bool                                          `json:"invoice_immediately,omitempty"`
	ScheduleType          enum.ScheduleType                              `json:"schedule_type,omitempty"`
}
type AdvanceInvoiceEstimateSpecificDatesSchedule struct {
	TermsToCharge *int32 `json:"terms_to_charge,omitempty"`
	Date          *int64 `json:"date,omitempty"`
}
type AdvanceInvoiceEstimateFixedIntervalSchedule struct {
	NumberOfOccurrences *int32             `json:"number_of_occurrences,omitempty"`
	DaysBeforeRenewal   *int32             `json:"days_before_renewal,omitempty"`
	EndScheduleOn       enum.EndScheduleOn `json:"end_schedule_on,omitempty"`
	EndDate             *int64             `json:"end_date,omitempty"`
}
type RegenerateInvoiceEstimateRequest struct {
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
	Prorate            *bool  `json:"prorate,omitempty"`
	InvoiceImmediately *bool  `json:"invoice_immediately,omitempty"`
}
type ChangeTermEndRequest struct {
	TermEndsAt         *int64 `json:"term_ends_at"`
	Prorate            *bool  `json:"prorate,omitempty"`
	InvoiceImmediately *bool  `json:"invoice_immediately,omitempty"`
}
type CancelSubscriptionRequest struct {
	CancelOption                      enum.CancelOption                      `json:"cancel_option,omitempty"`
	EndOfTerm                         *bool                                  `json:"end_of_term,omitempty"`
	CancelAt                          *int64                                 `json:"cancel_at,omitempty"`
	CreditOptionForCurrentTermCharges enum.CreditOptionForCurrentTermCharges `json:"credit_option_for_current_term_charges,omitempty"`
	UnbilledChargesOption             enum.UnbilledChargesOption             `json:"unbilled_charges_option,omitempty"`
	AccountReceivablesHandling        enum.AccountReceivablesHandling        `json:"account_receivables_handling,omitempty"`
	RefundableCreditsHandling         enum.RefundableCreditsHandling         `json:"refundable_credits_handling,omitempty"`
	ContractTermCancelOption          enum.ContractTermCancelOption          `json:"contract_term_cancel_option,omitempty"`
	InvoiceDate                       *int64                                 `json:"invoice_date,omitempty"`
	EventBasedAddons                  []*CancelSubscriptionEventBasedAddon   `json:"event_based_addons,omitempty"`
	CancelReasonCode                  string                                 `json:"cancel_reason_code,omitempty"`
}
type CancelSubscriptionEventBasedAddon struct {
	Id                  string `json:"id,omitempty"`
	Quantity            *int32 `json:"quantity,omitempty"`
	UnitPrice           *int64 `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32 `json:"service_period_in_days,omitempty"`
}
type CancelSubscriptionForItemsRequest struct {
	CancelOption                      enum.CancelOption                             `json:"cancel_option,omitempty"`
	EndOfTerm                         *bool                                         `json:"end_of_term,omitempty"`
	CancelAt                          *int64                                        `json:"cancel_at,omitempty"`
	CreditOptionForCurrentTermCharges enum.CreditOptionForCurrentTermCharges        `json:"credit_option_for_current_term_charges,omitempty"`
	UnbilledChargesOption             enum.UnbilledChargesOption                    `json:"unbilled_charges_option,omitempty"`
	AccountReceivablesHandling        enum.AccountReceivablesHandling               `json:"account_receivables_handling,omitempty"`
	RefundableCreditsHandling         enum.RefundableCreditsHandling                `json:"refundable_credits_handling,omitempty"`
	ContractTermCancelOption          enum.ContractTermCancelOption                 `json:"contract_term_cancel_option,omitempty"`
	InvoiceDate                       *int64                                        `json:"invoice_date,omitempty"`
	SubscriptionItems                 []*CancelSubscriptionForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	CancelReasonCode                  string                                        `json:"cancel_reason_code,omitempty"`
}
type CancelSubscriptionForItemsSubscriptionItem struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodDays  *int32 `json:"service_period_days,omitempty"`
}
type PauseSubscriptionRequest struct {
	PauseOption             enum.PauseOption               `json:"pause_option,omitempty"`
	Subscription            *PauseSubscriptionSubscription `json:"subscription,omitempty"`
	UnbilledChargesHandling enum.UnbilledChargesHandling   `json:"unbilled_charges_handling,omitempty"`
}
type PauseSubscriptionSubscription struct {
	PauseDate         *int64 `json:"pause_date,omitempty"`
	ResumeDate        *int64 `json:"resume_date,omitempty"`
	SkipBillingCycles *int32 `json:"skip_billing_cycles,omitempty"`
}
type ResumeSubscriptionRequest struct {
	ResumeOption    enum.ResumeOption               `json:"resume_option,omitempty"`
	Subscription    *ResumeSubscriptionSubscription `json:"subscription,omitempty"`
	ChargesHandling enum.ChargesHandling            `json:"charges_handling,omitempty"`
}
type ResumeSubscriptionSubscription struct {
	ResumeDate *int64 `json:"resume_date,omitempty"`
}
type GiftSubscriptionRequest struct {
	Gift            *GiftSubscriptionGift            `json:"gift,omitempty"`
	Gifter          *GiftSubscriptionGifter          `json:"gifter,omitempty"`
	GiftReceiver    *GiftSubscriptionGiftReceiver    `json:"gift_receiver,omitempty"`
	CouponIds       []string                         `json:"coupon_ids,omitempty"`
	PaymentIntent   *GiftSubscriptionPaymentIntent   `json:"payment_intent,omitempty"`
	ShippingAddress *GiftSubscriptionShippingAddress `json:"shipping_address,omitempty"`
	Subscription    *GiftSubscriptionSubscription    `json:"subscription,omitempty"`
	Addons          []*GiftSubscriptionAddon         `json:"addons,omitempty"`
}
type GiftSubscriptionGift struct {
	ScheduledAt     *int64 `json:"scheduled_at,omitempty"`
	AutoClaim       *bool  `json:"auto_claim,omitempty"`
	NoExpiry        *bool  `json:"no_expiry,omitempty"`
	ClaimExpiryDate *int64 `json:"claim_expiry_date,omitempty"`
}
type GiftSubscriptionGifter struct {
	CustomerId   string `json:"customer_id"`
	Signature    string `json:"signature"`
	Note         string `json:"note,omitempty"`
	PaymentSrcId string `json:"payment_src_id,omitempty"`
}
type GiftSubscriptionGiftReceiver struct {
	CustomerId string `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
}
type GiftSubscriptionPaymentIntent struct {
	Id                    string                          `json:"id,omitempty"`
	GatewayAccountId      string                          `json:"gateway_account_id,omitempty"`
	GwToken               string                          `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntent.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                          `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                          `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}          `json:"additional_information,omitempty"`
}
type GiftSubscriptionShippingAddress struct {
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
type GiftSubscriptionSubscription struct {
	PlanId                string `json:"plan_id"`
	PlanQuantity          *int32 `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal string `json:"plan_quantity_in_decimal,omitempty"`
}
type GiftSubscriptionAddon struct {
	Id                string `json:"id,omitempty"`
	Quantity          *int32 `json:"quantity,omitempty"`
	QuantityInDecimal string `json:"quantity_in_decimal,omitempty"`
}
type GiftSubscriptionForItemsRequest struct {
	Gift              *GiftSubscriptionForItemsGift               `json:"gift,omitempty"`
	Gifter            *GiftSubscriptionForItemsGifter             `json:"gifter,omitempty"`
	GiftReceiver      *GiftSubscriptionForItemsGiftReceiver       `json:"gift_receiver,omitempty"`
	CouponIds         []string                                    `json:"coupon_ids,omitempty"`
	PaymentIntent     *GiftSubscriptionForItemsPaymentIntent      `json:"payment_intent,omitempty"`
	ShippingAddress   *GiftSubscriptionForItemsShippingAddress    `json:"shipping_address,omitempty"`
	SubscriptionItems []*GiftSubscriptionForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	ItemTiers         []*GiftSubscriptionForItemsItemTier         `json:"item_tiers,omitempty"`
}
type GiftSubscriptionForItemsGift struct {
	ScheduledAt     *int64 `json:"scheduled_at,omitempty"`
	AutoClaim       *bool  `json:"auto_claim,omitempty"`
	NoExpiry        *bool  `json:"no_expiry,omitempty"`
	ClaimExpiryDate *int64 `json:"claim_expiry_date,omitempty"`
}
type GiftSubscriptionForItemsGifter struct {
	CustomerId   string `json:"customer_id"`
	Signature    string `json:"signature"`
	Note         string `json:"note,omitempty"`
	PaymentSrcId string `json:"payment_src_id,omitempty"`
}
type GiftSubscriptionForItemsGiftReceiver struct {
	CustomerId string `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
}
type GiftSubscriptionForItemsPaymentIntent struct {
	Id                    string                          `json:"id,omitempty"`
	GatewayAccountId      string                          `json:"gateway_account_id,omitempty"`
	GwToken               string                          `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntent.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                          `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                          `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}          `json:"additional_information,omitempty"`
}
type GiftSubscriptionForItemsShippingAddress struct {
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
type GiftSubscriptionForItemsSubscriptionItem struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
}
type GiftSubscriptionForItemsItemTier struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int64 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type CreateInvoiceRequest struct {
	Invoice                    *CreateInvoiceInvoice             `json:"invoice,omitempty"`
	CurrencyCode               string                            `json:"currency_code,omitempty"`
	Addons                     []*CreateInvoiceAddon             `json:"addons,omitempty"`
	Charges                    []*CreateInvoiceCharge            `json:"charges,omitempty"`
	InvoiceNote                string                            `json:"invoice_note,omitempty"`
	RemoveGeneralNote          *bool                             `json:"remove_general_note,omitempty"`
	NotesToRemove              []*CreateInvoiceNotesToRemove     `json:"notes_to_remove,omitempty"`
	Coupon                     string                            `json:"coupon,omitempty"`
	CouponIds                  []string                          `json:"coupon_ids,omitempty"`
	AuthorizationTransactionId string                            `json:"authorization_transaction_id,omitempty"`
	PaymentSourceId            string                            `json:"payment_source_id,omitempty"`
	AutoCollection             enum.AutoCollection               `json:"auto_collection,omitempty"`
	InvoiceDate                *int64                            `json:"invoice_date,omitempty"`
	ShippingAddress            *CreateInvoiceShippingAddress     `json:"shipping_address,omitempty"`
	TaxProvidersFields         []*CreateInvoiceTaxProvidersField `json:"tax_providers_fields,omitempty"`
}
type CreateInvoiceInvoice struct {
	CustomerId     string `json:"customer_id,omitempty"`
	SubscriptionId string `json:"subscription_id,omitempty"`
	PoNumber       string `json:"po_number,omitempty"`
}
type CreateInvoiceAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type CreateInvoiceCharge struct {
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
type CreateInvoiceNotesToRemove struct {
	EntityType enum.EntityType `json:"entity_type,omitempty"`
	EntityId   string          `json:"entity_id,omitempty"`
}
type CreateInvoiceShippingAddress struct {
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
type CreateInvoiceTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type CreateInvoiceForItemsRequest struct {
	Invoice                    *CreateInvoiceForItemsInvoice             `json:"invoice,omitempty"`
	CurrencyCode               string                                    `json:"currency_code,omitempty"`
	ItemPrices                 []*CreateInvoiceForItemsItemPrice         `json:"item_prices,omitempty"`
	ItemTiers                  []*CreateInvoiceForItemsItemTier          `json:"item_tiers,omitempty"`
	Charges                    []*CreateInvoiceForItemsCharge            `json:"charges,omitempty"`
	InvoiceNote                string                                    `json:"invoice_note,omitempty"`
	RemoveGeneralNote          *bool                                     `json:"remove_general_note,omitempty"`
	NotesToRemove              []*CreateInvoiceForItemsNotesToRemove     `json:"notes_to_remove,omitempty"`
	Coupon                     string                                    `json:"coupon,omitempty"`
	CouponIds                  []string                                  `json:"coupon_ids,omitempty"`
	AuthorizationTransactionId string                                    `json:"authorization_transaction_id,omitempty"`
	PaymentSourceId            string                                    `json:"payment_source_id,omitempty"`
	AutoCollection             enum.AutoCollection                       `json:"auto_collection,omitempty"`
	Discounts                  []*CreateInvoiceForItemsDiscount          `json:"discounts,omitempty"`
	ShippingAddress            *CreateInvoiceForItemsShippingAddress     `json:"shipping_address,omitempty"`
	TaxProvidersFields         []*CreateInvoiceForItemsTaxProvidersField `json:"tax_providers_fields,omitempty"`
	InvoiceDate                *int64                                    `json:"invoice_date,omitempty"`
	BillingAddress             *CreateInvoiceForItemsBillingAddress      `json:"billing_address,omitempty"`
}
type CreateInvoiceForItemsInvoice struct {
	CustomerId     string `json:"customer_id,omitempty"`
	SubscriptionId string `json:"subscription_id,omitempty"`
	PoNumber       string `json:"po_number,omitempty"`
}
type CreateInvoiceForItemsItemPrice struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type CreateInvoiceForItemsItemTier struct {
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
type CreateInvoiceForItemsCharge struct {
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
type CreateInvoiceForItemsNotesToRemove struct {
	EntityType enum.EntityType `json:"entity_type,omitempty"`
	EntityId   string          `json:"entity_id,omitempty"`
}
type CreateInvoiceForItemsDiscount struct {
	Percentage  *float64     `json:"percentage,omitempty"`
	Amount      *int64       `json:"amount,omitempty"`
	Quantity    *int32       `json:"quantity,omitempty"`
	ApplyOn     enum.ApplyOn `json:"apply_on"`
	ItemPriceId string       `json:"item_price_id,omitempty"`
}
type CreateInvoiceForItemsShippingAddress struct {
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
type CreateInvoiceForItemsTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type CreateInvoiceForItemsBillingAddress struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type PaymentSchedulesRequest struct {
	SchemeId                 string `json:"scheme_id"`
	Amount                   *int64 `json:"amount,omitempty"`
	InvoiceId                string `json:"invoice_id,omitempty"`
	PaymentScheduleStartDate *int64 `json:"payment_schedule_start_date,omitempty"`
}

type CreateSubscriptionResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}

type CreateSubItemEstimateResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}

type CreateSubForCustomerEstimateResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}

type CreateSubItemForCustomerEstimateResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}

type UpdateSubscriptionResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}

type UpdateSubscriptionForItemsResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}

type RenewalEstimateResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}

type AdvanceInvoiceEstimateResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}

type RegenerateInvoiceEstimateResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}

type UpcomingInvoicesEstimateResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}

type ChangeTermEndResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}

type CancelSubscriptionResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}

type CancelSubscriptionForItemsResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}

type PauseSubscriptionResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}

type ResumeSubscriptionResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}

type GiftSubscriptionResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}

type GiftSubscriptionForItemsResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}

type CreateInvoiceResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}

type CreateInvoiceForItemsResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}

type PaymentSchedulesResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
}
