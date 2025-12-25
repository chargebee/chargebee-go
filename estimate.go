package chargebee

type EstimateSubscriptionEstimateStatus string

const (
	EstimateSubscriptionEstimateStatusFuture      EstimateSubscriptionEstimateStatus = "future"
	EstimateSubscriptionEstimateStatusInTrial     EstimateSubscriptionEstimateStatus = "in_trial"
	EstimateSubscriptionEstimateStatusActive      EstimateSubscriptionEstimateStatus = "active"
	EstimateSubscriptionEstimateStatusNonRenewing EstimateSubscriptionEstimateStatus = "non_renewing"
	EstimateSubscriptionEstimateStatusPaused      EstimateSubscriptionEstimateStatus = "paused"
	EstimateSubscriptionEstimateStatusCancelled   EstimateSubscriptionEstimateStatus = "cancelled"
	EstimateSubscriptionEstimateStatusTransferred EstimateSubscriptionEstimateStatus = "transferred"
)

type EstimateSubscriptionEstimateTrialEndAction string

const (
	EstimateSubscriptionEstimateTrialEndActionSiteDefault          EstimateSubscriptionEstimateTrialEndAction = "site_default"
	EstimateSubscriptionEstimateTrialEndActionPlanDefault          EstimateSubscriptionEstimateTrialEndAction = "plan_default"
	EstimateSubscriptionEstimateTrialEndActionActivateSubscription EstimateSubscriptionEstimateTrialEndAction = "activate_subscription"
	EstimateSubscriptionEstimateTrialEndActionCancelSubscription   EstimateSubscriptionEstimateTrialEndAction = "cancel_subscription"
)

type EstimateInvoiceEstimatePriceType string

const (
	EstimateInvoiceEstimatePriceTypeTaxExclusive EstimateInvoiceEstimatePriceType = "tax_exclusive"
	EstimateInvoiceEstimatePriceTypeTaxInclusive EstimateInvoiceEstimatePriceType = "tax_inclusive"
)

type EstimatePaymentScheduleEstimateEntityType string

const (
	EstimatePaymentScheduleEstimateEntityTypeInvoice EstimatePaymentScheduleEstimateEntityType = "invoice"
)

type EstimateCreditNoteEstimateType string

const (
	EstimateCreditNoteEstimateTypeAdjustment EstimateCreditNoteEstimateType = "adjustment"
	EstimateCreditNoteEstimateTypeRefundable EstimateCreditNoteEstimateType = "refundable"
	EstimateCreditNoteEstimateTypeStore      EstimateCreditNoteEstimateType = "store"
)

type EstimateCreditNoteEstimatePriceType string

const (
	EstimateCreditNoteEstimatePriceTypeTaxExclusive EstimateCreditNoteEstimatePriceType = "tax_exclusive"
	EstimateCreditNoteEstimatePriceTypeTaxInclusive EstimateCreditNoteEstimatePriceType = "tax_inclusive"
)

type EstimateUnbilledChargePricingModel string

const (
	EstimateUnbilledChargePricingModelFlatFee   EstimateUnbilledChargePricingModel = "flat_fee"
	EstimateUnbilledChargePricingModelPerUnit   EstimateUnbilledChargePricingModel = "per_unit"
	EstimateUnbilledChargePricingModelTiered    EstimateUnbilledChargePricingModel = "tiered"
	EstimateUnbilledChargePricingModelVolume    EstimateUnbilledChargePricingModel = "volume"
	EstimateUnbilledChargePricingModelStairstep EstimateUnbilledChargePricingModel = "stairstep"
)

type EstimateUnbilledChargeEntityType string

const (
	EstimateUnbilledChargeEntityTypeAdhoc           EstimateUnbilledChargeEntityType = "adhoc"
	EstimateUnbilledChargeEntityTypePlanItemPrice   EstimateUnbilledChargeEntityType = "plan_item_price"
	EstimateUnbilledChargeEntityTypeAddonItemPrice  EstimateUnbilledChargeEntityType = "addon_item_price"
	EstimateUnbilledChargeEntityTypeChargeItemPrice EstimateUnbilledChargeEntityType = "charge_item_price"
	EstimateUnbilledChargeEntityTypePlanSetup       EstimateUnbilledChargeEntityType = "plan_setup"
	EstimateUnbilledChargeEntityTypePlan            EstimateUnbilledChargeEntityType = "plan"
	EstimateUnbilledChargeEntityTypeAddon           EstimateUnbilledChargeEntityType = "addon"
)

type EstimateBillingAlignmentMode string

const (
	EstimateBillingAlignmentModeImmediate EstimateBillingAlignmentMode = "immediate"
	EstimateBillingAlignmentModeDelayed   EstimateBillingAlignmentMode = "delayed"
)

type EstimateChangeOption string

const (
	EstimateChangeOptionImmediately  EstimateChangeOption = "immediately"
	EstimateChangeOptionEndOfTerm    EstimateChangeOption = "end_of_term"
	EstimateChangeOptionSpecificDate EstimateChangeOption = "specific_date"
)

type EstimateScheduleType string

const (
	EstimateScheduleTypeImmediate      EstimateScheduleType = "immediate"
	EstimateScheduleTypeSpecificDates  EstimateScheduleType = "specific_dates"
	EstimateScheduleTypeFixedIntervals EstimateScheduleType = "fixed_intervals"
)

type EstimateCancelOption string

const (
	EstimateCancelOptionImmediately      EstimateCancelOption = "immediately"
	EstimateCancelOptionEndOfTerm        EstimateCancelOption = "end_of_term"
	EstimateCancelOptionSpecificDate     EstimateCancelOption = "specific_date"
	EstimateCancelOptionEndOfBillingTerm EstimateCancelOption = "end_of_billing_term"
)

type EstimateCreditOptionForCurrentTermCharges string

const (
	EstimateCreditOptionForCurrentTermChargesNone    EstimateCreditOptionForCurrentTermCharges = "none"
	EstimateCreditOptionForCurrentTermChargesProrate EstimateCreditOptionForCurrentTermCharges = "prorate"
	EstimateCreditOptionForCurrentTermChargesFull    EstimateCreditOptionForCurrentTermCharges = "full"
)

type EstimateUnbilledChargesOption string

const (
	EstimateUnbilledChargesOptionInvoice EstimateUnbilledChargesOption = "invoice"
	EstimateUnbilledChargesOptionDelete  EstimateUnbilledChargesOption = "delete"
)

type EstimateAccountReceivablesHandling string

const (
	EstimateAccountReceivablesHandlingNoAction                  EstimateAccountReceivablesHandling = "no_action"
	EstimateAccountReceivablesHandlingSchedulePaymentCollection EstimateAccountReceivablesHandling = "schedule_payment_collection"
	EstimateAccountReceivablesHandlingWriteOff                  EstimateAccountReceivablesHandling = "write_off"
)

type EstimateRefundableCreditsHandling string

const (
	EstimateRefundableCreditsHandlingNoAction       EstimateRefundableCreditsHandling = "no_action"
	EstimateRefundableCreditsHandlingScheduleRefund EstimateRefundableCreditsHandling = "schedule_refund"
)

type EstimateContractTermCancelOption string

const (
	EstimateContractTermCancelOptionTerminateImmediately         EstimateContractTermCancelOption = "terminate_immediately"
	EstimateContractTermCancelOptionEndOfContractTerm            EstimateContractTermCancelOption = "end_of_contract_term"
	EstimateContractTermCancelOptionSpecificDate                 EstimateContractTermCancelOption = "specific_date"
	EstimateContractTermCancelOptionEndOfSubscriptionBillingTerm EstimateContractTermCancelOption = "end_of_subscription_billing_term"
)

type EstimatePauseOption string

const (
	EstimatePauseOptionImmediately   EstimatePauseOption = "immediately"
	EstimatePauseOptionEndOfTerm     EstimatePauseOption = "end_of_term"
	EstimatePauseOptionSpecificDate  EstimatePauseOption = "specific_date"
	EstimatePauseOptionBillingCycles EstimatePauseOption = "billing_cycles"
)

type EstimateUnbilledChargesHandling string

const (
	EstimateUnbilledChargesHandlingNoAction EstimateUnbilledChargesHandling = "no_action"
	EstimateUnbilledChargesHandlingInvoice  EstimateUnbilledChargesHandling = "invoice"
)

type EstimateResumeOption string

const (
	EstimateResumeOptionImmediately  EstimateResumeOption = "immediately"
	EstimateResumeOptionSpecificDate EstimateResumeOption = "specific_date"
)

type EstimateChargesHandling string

const (
	EstimateChargesHandlingInvoiceImmediately   EstimateChargesHandling = "invoice_immediately"
	EstimateChargesHandlingAddToUnbilledCharges EstimateChargesHandling = "add_to_unbilled_charges"
)

type EstimateAutoCollection string

const (
	EstimateAutoCollectionOn  EstimateAutoCollection = "on"
	EstimateAutoCollectionOff EstimateAutoCollection = "off"
)

// just struct
type Estimate struct {
	CreatedAt                int64                      `json:"created_at"`
	SubscriptionEstimate     *SubscriptionEstimate      `json:"subscription_estimate"`
	SubscriptionEstimates    []*SubscriptionEstimate    `json:"subscription_estimates"`
	InvoiceEstimate          *InvoiceEstimate           `json:"invoice_estimate"`
	InvoiceEstimates         []*InvoiceEstimate         `json:"invoice_estimates"`
	PaymentScheduleEstimates []*PaymentScheduleEstimate `json:"payment_schedule_estimates"`
	NextInvoiceEstimate      *InvoiceEstimate           `json:"next_invoice_estimate"`
	CreditNoteEstimates      []*CreditNoteEstimate      `json:"credit_note_estimates"`
	UnbilledChargeEstimates  []*UnbilledCharge          `json:"unbilled_charge_estimates"`
	Object                   string                     `json:"object"`
}

// sub resources
// operations
// input params
type EstimateCreateSubscriptionRequest struct {
	Subscription            *EstimateCreateSubscriptionSubscription        `json:"subscription,omitempty"`
	BillingCycles           *int32                                         `json:"billing_cycles,omitempty"`
	Addons                  []*EstimateCreateSubscriptionAddon             `json:"addons,omitempty"`
	EventBasedAddons        []*EstimateCreateSubscriptionEventBasedAddon   `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove []string                                       `json:"mandatory_addons_to_remove,omitempty"`
	TermsToCharge           *int32                                         `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode    EstimateBillingAlignmentMode                   `json:"billing_alignment_mode,omitempty"`
	CouponIds               []string                                       `json:"coupon_ids,omitempty"`
	BillingAddress          *EstimateCreateSubscriptionBillingAddress      `json:"billing_address,omitempty"`
	ShippingAddress         *EstimateCreateSubscriptionShippingAddress     `json:"shipping_address,omitempty"`
	Customer                *EstimateCreateSubscriptionCustomer            `json:"customer,omitempty"`
	InvoiceImmediately      *bool                                          `json:"invoice_immediately,omitempty"`
	InvoiceDate             *int64                                         `json:"invoice_date,omitempty"`
	ContractTerm            *EstimateCreateSubscriptionContractTerm        `json:"contract_term,omitempty"`
	TaxProvidersFields      []*EstimateCreateSubscriptionTaxProvidersField `json:"tax_providers_fields,omitempty"`
	ClientProfileId         string                                         `json:"client_profile_id,omitempty"`
	apiRequest              `json:"-" form:"-"`
}

func (r *EstimateCreateSubscriptionRequest) payload() any { return r }

// input sub resource params single
type EstimateCreateSubscriptionSubscription struct {
	Id                                string               `json:"id,omitempty"`
	PlanId                            string               `json:"plan_id"`
	PlanQuantity                      *int32               `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string               `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int64               `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string               `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int64               `json:"setup_fee,omitempty"`
	TrialEnd                          *int64               `json:"trial_end,omitempty"`
	StartDate                         *int64               `json:"start_date,omitempty"`
	Coupon                            string               `json:"coupon,omitempty"`
	OfflinePaymentMethod              OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	FreePeriod                        *int32               `json:"free_period,omitempty"`
	FreePeriodUnit                    FreePeriodUnit       `json:"free_period_unit,omitempty"`
	ContractTermBillingCycleOnRenewal *int32               `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	TrialEndAction                    TrialEndAction       `json:"trial_end_action,omitempty"`
}

// input sub resource params multi
type EstimateCreateSubscriptionAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}

// input sub resource params multi
type EstimateCreateSubscriptionEventBasedAddon struct {
	Id                  string   `json:"id,omitempty"`
	Quantity            *int32   `json:"quantity,omitempty"`
	UnitPrice           *int64   `json:"unit_price,omitempty"`
	QuantityInDecimal   string   `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string   `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodInDays *int32   `json:"service_period_in_days,omitempty"`
	OnEvent             OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool    `json:"charge_once,omitempty"`
	ChargeOn            ChargeOn `json:"charge_on,omitempty"`
}

// input sub resource params single
type EstimateCreateSubscriptionBillingAddress struct {
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type EstimateCreateSubscriptionShippingAddress struct {
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type EstimateCreateSubscriptionCustomer struct {
	VatNumber        string                   `json:"vat_number,omitempty"`
	VatNumberPrefix  string                   `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool                    `json:"registered_for_gst,omitempty"`
	Taxability       Taxability               `json:"taxability,omitempty"`
	EntityCode       EntityCode               `json:"entity_code,omitempty"`
	ExemptNumber     string                   `json:"exempt_number,omitempty"`
	ExemptionDetails []map[string]interface{} `json:"exemption_details,omitempty"`
	CustomerType     CustomerType             `json:"customer_type,omitempty"`
}

// input sub resource params single
type EstimateCreateSubscriptionContractTerm struct {
	ActionAtTermEnd          ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32          `json:"cancellation_cutoff_period,omitempty"`
}

// input sub resource params multi
type EstimateCreateSubscriptionTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type EstimateCreateSubItemEstimateRequest struct {
	Subscription           *EstimateCreateSubItemEstimateSubscription        `json:"subscription,omitempty"`
	BillingCycles          *int32                                            `json:"billing_cycles,omitempty"`
	SubscriptionItems      []*EstimateCreateSubItemEstimateSubscriptionItem  `json:"subscription_items,omitempty"`
	Discounts              []*EstimateCreateSubItemEstimateDiscount          `json:"discounts,omitempty"`
	MandatoryItemsToRemove []string                                          `json:"mandatory_items_to_remove,omitempty"`
	ItemTiers              []*EstimateCreateSubItemEstimateItemTier          `json:"item_tiers,omitempty"`
	TermsToCharge          *int32                                            `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode   EstimateBillingAlignmentMode                      `json:"billing_alignment_mode,omitempty"`
	CouponIds              []string                                          `json:"coupon_ids,omitempty"`
	BillingAddress         *EstimateCreateSubItemEstimateBillingAddress      `json:"billing_address,omitempty"`
	ShippingAddress        *EstimateCreateSubItemEstimateShippingAddress     `json:"shipping_address,omitempty"`
	Customer               *EstimateCreateSubItemEstimateCustomer            `json:"customer,omitempty"`
	InvoiceImmediately     *bool                                             `json:"invoice_immediately,omitempty"`
	InvoiceDate            *int64                                            `json:"invoice_date,omitempty"`
	ClientProfileId        string                                            `json:"client_profile_id,omitempty"`
	ContractTerm           *EstimateCreateSubItemEstimateContractTerm        `json:"contract_term,omitempty"`
	TaxProvidersFields     []*EstimateCreateSubItemEstimateTaxProvidersField `json:"tax_providers_fields,omitempty"`
	apiRequest             `json:"-" form:"-"`
}

func (r *EstimateCreateSubItemEstimateRequest) payload() any { return r }

// input sub resource params single
type EstimateCreateSubItemEstimateSubscription struct {
	Id                                string               `json:"id,omitempty"`
	TrialEnd                          *int64               `json:"trial_end,omitempty"`
	SetupFee                          *int64               `json:"setup_fee,omitempty"`
	StartDate                         *int64               `json:"start_date,omitempty"`
	Coupon                            string               `json:"coupon,omitempty"`
	OfflinePaymentMethod              OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	FreePeriod                        *int32               `json:"free_period,omitempty"`
	FreePeriodUnit                    FreePeriodUnit       `json:"free_period_unit,omitempty"`
	ContractTermBillingCycleOnRenewal *int32               `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	TrialEndAction                    TrialEndAction       `json:"trial_end_action,omitempty"`
}

// input sub resource params multi
type EstimateCreateSubItemEstimateSubscriptionItem struct {
	ItemPriceId        string         `json:"item_price_id"`
	Quantity           *int32         `json:"quantity,omitempty"`
	QuantityInDecimal  string         `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64         `json:"unit_price,omitempty"`
	UnitPriceInDecimal string         `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32         `json:"billing_cycles,omitempty"`
	TrialEnd           *int64         `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32         `json:"service_period_days,omitempty"`
	ChargeOnEvent      ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool          `json:"charge_once,omitempty"`
	ItemType           ItemType       `json:"item_type,omitempty"`
	ChargeOnOption     ChargeOnOption `json:"charge_on_option,omitempty"`
}

// input sub resource params multi
type EstimateCreateSubItemEstimateDiscount struct {
	ApplyOn       ApplyOn      `json:"apply_on,omitempty"`
	DurationType  DurationType `json:"duration_type"`
	Percentage    *float64     `json:"percentage,omitempty"`
	Amount        *int64       `json:"amount,omitempty"`
	Period        *int32       `json:"period,omitempty"`
	PeriodUnit    PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool        `json:"included_in_mrr,omitempty"`
	ItemPriceId   string       `json:"item_price_id,omitempty"`
	Quantity      *int32       `json:"quantity,omitempty"`
}

// input sub resource params multi
type EstimateCreateSubItemEstimateItemTier struct {
	ItemPriceId           string      `json:"item_price_id,omitempty"`
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
type EstimateCreateSubItemEstimateBillingAddress struct {
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type EstimateCreateSubItemEstimateShippingAddress struct {
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type EstimateCreateSubItemEstimateCustomer struct {
	VatNumber        string                   `json:"vat_number,omitempty"`
	VatNumberPrefix  string                   `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool                    `json:"registered_for_gst,omitempty"`
	Taxability       Taxability               `json:"taxability,omitempty"`
	EntityCode       EntityCode               `json:"entity_code,omitempty"`
	ExemptNumber     string                   `json:"exempt_number,omitempty"`
	ExemptionDetails []map[string]interface{} `json:"exemption_details,omitempty"`
	CustomerType     CustomerType             `json:"customer_type,omitempty"`
}

// input sub resource params single
type EstimateCreateSubItemEstimateContractTerm struct {
	ActionAtTermEnd          ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	ContractStart            *int64          `json:"contract_start,omitempty"`
	CancellationCutoffPeriod *int32          `json:"cancellation_cutoff_period,omitempty"`
}

// input sub resource params multi
type EstimateCreateSubItemEstimateTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type EstimateCreateSubForCustomerEstimateRequest struct {
	UseExistingBalances     *bool                                                  `json:"use_existing_balances,omitempty"`
	Subscription            *EstimateCreateSubForCustomerEstimateSubscription      `json:"subscription,omitempty"`
	InvoiceImmediately      *bool                                                  `json:"invoice_immediately,omitempty"`
	BillingCycles           *int32                                                 `json:"billing_cycles,omitempty"`
	Addons                  []*EstimateCreateSubForCustomerEstimateAddon           `json:"addons,omitempty"`
	EventBasedAddons        []*EstimateCreateSubForCustomerEstimateEventBasedAddon `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove []string                                               `json:"mandatory_addons_to_remove,omitempty"`
	TermsToCharge           *int32                                                 `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode    EstimateBillingAlignmentMode                           `json:"billing_alignment_mode,omitempty"`
	ShippingAddress         *EstimateCreateSubForCustomerEstimateShippingAddress   `json:"shipping_address,omitempty"`
	InvoiceDate             *int64                                                 `json:"invoice_date,omitempty"`
	CouponIds               []string                                               `json:"coupon_ids,omitempty"`
	ContractTerm            *EstimateCreateSubForCustomerEstimateContractTerm      `json:"contract_term,omitempty"`
	apiRequest              `json:"-" form:"-"`
}

func (r *EstimateCreateSubForCustomerEstimateRequest) payload() any { return r }

// input sub resource params single
type EstimateCreateSubForCustomerEstimateSubscription struct {
	Id                                string               `json:"id,omitempty"`
	PlanId                            string               `json:"plan_id"`
	PlanQuantity                      *int32               `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string               `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int64               `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string               `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int64               `json:"setup_fee,omitempty"`
	TrialEnd                          *int64               `json:"trial_end,omitempty"`
	StartDate                         *int64               `json:"start_date,omitempty"`
	OfflinePaymentMethod              OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	FreePeriod                        *int32               `json:"free_period,omitempty"`
	FreePeriodUnit                    FreePeriodUnit       `json:"free_period_unit,omitempty"`
	ContractTermBillingCycleOnRenewal *int32               `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	TrialEndAction                    TrialEndAction       `json:"trial_end_action,omitempty"`
}

// input sub resource params multi
type EstimateCreateSubForCustomerEstimateAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}

// input sub resource params multi
type EstimateCreateSubForCustomerEstimateEventBasedAddon struct {
	Id                  string   `json:"id,omitempty"`
	Quantity            *int32   `json:"quantity,omitempty"`
	UnitPrice           *int64   `json:"unit_price,omitempty"`
	QuantityInDecimal   string   `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string   `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodInDays *int32   `json:"service_period_in_days,omitempty"`
	OnEvent             OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool    `json:"charge_once,omitempty"`
	ChargeOn            ChargeOn `json:"charge_on,omitempty"`
}

// input sub resource params single
type EstimateCreateSubForCustomerEstimateShippingAddress struct {
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type EstimateCreateSubForCustomerEstimateContractTerm struct {
	ActionAtTermEnd          ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32          `json:"cancellation_cutoff_period,omitempty"`
}
type EstimateCreateSubItemForCustomerEstimateRequest struct {
	UseExistingBalances    *bool                                                       `json:"use_existing_balances,omitempty"`
	Subscription           *EstimateCreateSubItemForCustomerEstimateSubscription       `json:"subscription,omitempty"`
	InvoiceImmediately     *bool                                                       `json:"invoice_immediately,omitempty"`
	BillingCycles          *int32                                                      `json:"billing_cycles,omitempty"`
	SubscriptionItems      []*EstimateCreateSubItemForCustomerEstimateSubscriptionItem `json:"subscription_items,omitempty"`
	Discounts              []*EstimateCreateSubItemForCustomerEstimateDiscount         `json:"discounts,omitempty"`
	MandatoryItemsToRemove []string                                                    `json:"mandatory_items_to_remove,omitempty"`
	ItemTiers              []*EstimateCreateSubItemForCustomerEstimateItemTier         `json:"item_tiers,omitempty"`
	TermsToCharge          *int32                                                      `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode   EstimateBillingAlignmentMode                                `json:"billing_alignment_mode,omitempty"`
	ShippingAddress        *EstimateCreateSubItemForCustomerEstimateShippingAddress    `json:"shipping_address,omitempty"`
	BillingAddress         *EstimateCreateSubItemForCustomerEstimateBillingAddress     `json:"billing_address,omitempty"`
	InvoiceDate            *int64                                                      `json:"invoice_date,omitempty"`
	CouponIds              []string                                                    `json:"coupon_ids,omitempty"`
	ContractTerm           *EstimateCreateSubItemForCustomerEstimateContractTerm       `json:"contract_term,omitempty"`
	BillingOverride        *EstimateCreateSubItemForCustomerEstimateBillingOverride    `json:"billing_override,omitempty"`
	apiRequest             `json:"-" form:"-"`
}

func (r *EstimateCreateSubItemForCustomerEstimateRequest) payload() any { return r }

// input sub resource params single
type EstimateCreateSubItemForCustomerEstimateSubscription struct {
	Id                                string               `json:"id,omitempty"`
	TrialEnd                          *int64               `json:"trial_end,omitempty"`
	SetupFee                          *int64               `json:"setup_fee,omitempty"`
	StartDate                         *int64               `json:"start_date,omitempty"`
	OfflinePaymentMethod              OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	FreePeriod                        *int32               `json:"free_period,omitempty"`
	FreePeriodUnit                    FreePeriodUnit       `json:"free_period_unit,omitempty"`
	ContractTermBillingCycleOnRenewal *int32               `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	TrialEndAction                    TrialEndAction       `json:"trial_end_action,omitempty"`
}

// input sub resource params multi
type EstimateCreateSubItemForCustomerEstimateSubscriptionItem struct {
	ItemPriceId        string         `json:"item_price_id"`
	Quantity           *int32         `json:"quantity,omitempty"`
	QuantityInDecimal  string         `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64         `json:"unit_price,omitempty"`
	UnitPriceInDecimal string         `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32         `json:"billing_cycles,omitempty"`
	TrialEnd           *int64         `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32         `json:"service_period_days,omitempty"`
	ChargeOnEvent      ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool          `json:"charge_once,omitempty"`
	ItemType           ItemType       `json:"item_type,omitempty"`
	ChargeOnOption     ChargeOnOption `json:"charge_on_option,omitempty"`
}

// input sub resource params multi
type EstimateCreateSubItemForCustomerEstimateDiscount struct {
	ApplyOn       ApplyOn      `json:"apply_on,omitempty"`
	DurationType  DurationType `json:"duration_type"`
	Percentage    *float64     `json:"percentage,omitempty"`
	Amount        *int64       `json:"amount,omitempty"`
	Period        *int32       `json:"period,omitempty"`
	PeriodUnit    PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool        `json:"included_in_mrr,omitempty"`
	ItemPriceId   string       `json:"item_price_id,omitempty"`
	Quantity      *int32       `json:"quantity,omitempty"`
}

// input sub resource params multi
type EstimateCreateSubItemForCustomerEstimateItemTier struct {
	ItemPriceId           string      `json:"item_price_id,omitempty"`
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
type EstimateCreateSubItemForCustomerEstimateShippingAddress struct {
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type EstimateCreateSubItemForCustomerEstimateBillingAddress struct {
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type EstimateCreateSubItemForCustomerEstimateContractTerm struct {
	ActionAtTermEnd          ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	ContractStart            *int64          `json:"contract_start,omitempty"`
	CancellationCutoffPeriod *int32          `json:"cancellation_cutoff_period,omitempty"`
}

// input sub resource params single
type EstimateCreateSubItemForCustomerEstimateBillingOverride struct {
	MaxExcessPaymentUsage     *int64 `json:"max_excess_payment_usage,omitempty"`
	MaxRefundableCreditsUsage *int64 `json:"max_refundable_credits_usage,omitempty"`
}
type EstimateUpdateSubscriptionRequest struct {
	Subscription            *EstimateUpdateSubscriptionSubscription      `json:"subscription,omitempty"`
	ChangesScheduledAt      *int64                                       `json:"changes_scheduled_at,omitempty"`
	ChangeOption            EstimateChangeOption                         `json:"change_option,omitempty"`
	Addons                  []*EstimateUpdateSubscriptionAddon           `json:"addons,omitempty"`
	EventBasedAddons        []*EstimateUpdateSubscriptionEventBasedAddon `json:"event_based_addons,omitempty"`
	ReplaceAddonList        *bool                                        `json:"replace_addon_list,omitempty"`
	MandatoryAddonsToRemove []string                                     `json:"mandatory_addons_to_remove,omitempty"`
	InvoiceDate             *int64                                       `json:"invoice_date,omitempty"`
	BillingCycles           *int32                                       `json:"billing_cycles,omitempty"`
	TermsToCharge           *int32                                       `json:"terms_to_charge,omitempty"`
	ReactivateFrom          *int64                                       `json:"reactivate_from,omitempty"`
	BillingAlignmentMode    EstimateBillingAlignmentMode                 `json:"billing_alignment_mode,omitempty"`
	CouponIds               []string                                     `json:"coupon_ids,omitempty"`
	ReplaceCouponList       *bool                                        `json:"replace_coupon_list,omitempty"`
	Prorate                 *bool                                        `json:"prorate,omitempty"`
	EndOfTerm               *bool                                        `json:"end_of_term,omitempty"`
	ForceTermReset          *bool                                        `json:"force_term_reset,omitempty"`
	Reactivate              *bool                                        `json:"reactivate,omitempty"`
	IncludeDelayedCharges   *bool                                        `json:"include_delayed_charges,omitempty"`
	UseExistingBalances     *bool                                        `json:"use_existing_balances,omitempty"`
	BillingAddress          *EstimateUpdateSubscriptionBillingAddress    `json:"billing_address,omitempty"`
	ShippingAddress         *EstimateUpdateSubscriptionShippingAddress   `json:"shipping_address,omitempty"`
	Customer                *EstimateUpdateSubscriptionCustomer          `json:"customer,omitempty"`
	InvoiceImmediately      *bool                                        `json:"invoice_immediately,omitempty"`
	apiRequest              `json:"-" form:"-"`
}

func (r *EstimateUpdateSubscriptionRequest) payload() any { return r }

// input sub resource params single
type EstimateUpdateSubscriptionSubscription struct {
	Id                     string               `json:"id"`
	PlanId                 string               `json:"plan_id,omitempty"`
	PlanQuantity           *int32               `json:"plan_quantity,omitempty"`
	PlanUnitPrice          *int64               `json:"plan_unit_price,omitempty"`
	SetupFee               *int64               `json:"setup_fee,omitempty"`
	PlanQuantityInDecimal  string               `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPriceInDecimal string               `json:"plan_unit_price_in_decimal,omitempty"`
	StartDate              *int64               `json:"start_date,omitempty"`
	TrialEnd               *int64               `json:"trial_end,omitempty"`
	Coupon                 string               `json:"coupon,omitempty"`
	AutoCollection         AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod   OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	FreePeriod             *int32               `json:"free_period,omitempty"`
	FreePeriodUnit         FreePeriodUnit       `json:"free_period_unit,omitempty"`
	TrialEndAction         TrialEndAction       `json:"trial_end_action,omitempty"`
}

// input sub resource params multi
type EstimateUpdateSubscriptionAddon struct {
	Id                 string        `json:"id,omitempty"`
	Quantity           *int32        `json:"quantity,omitempty"`
	UnitPrice          *int64        `json:"unit_price,omitempty"`
	BillingCycles      *int32        `json:"billing_cycles,omitempty"`
	QuantityInDecimal  string        `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string        `json:"unit_price_in_decimal,omitempty"`
	TrialEnd           *int64        `json:"trial_end,omitempty"`
	ProrationType      ProrationType `json:"proration_type,omitempty"`
}

// input sub resource params multi
type EstimateUpdateSubscriptionEventBasedAddon struct {
	Id                  string   `json:"id,omitempty"`
	Quantity            *int32   `json:"quantity,omitempty"`
	UnitPrice           *int64   `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32   `json:"service_period_in_days,omitempty"`
	ChargeOn            ChargeOn `json:"charge_on,omitempty"`
	OnEvent             OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool    `json:"charge_once,omitempty"`
	QuantityInDecimal   string   `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string   `json:"unit_price_in_decimal,omitempty"`
}

// input sub resource params single
type EstimateUpdateSubscriptionBillingAddress struct {
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type EstimateUpdateSubscriptionShippingAddress struct {
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type EstimateUpdateSubscriptionCustomer struct {
	VatNumber        string     `json:"vat_number,omitempty"`
	VatNumberPrefix  string     `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool      `json:"registered_for_gst,omitempty"`
	Taxability       Taxability `json:"taxability,omitempty"`
}
type EstimateUpdateSubscriptionForItemsRequest struct {
	Subscription           *EstimateUpdateSubscriptionForItemsSubscription       `json:"subscription,omitempty"`
	ChangesScheduledAt     *int64                                                `json:"changes_scheduled_at,omitempty"`
	ChangeOption           EstimateChangeOption                                  `json:"change_option,omitempty"`
	SubscriptionItems      []*EstimateUpdateSubscriptionForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	MandatoryItemsToRemove []string                                              `json:"mandatory_items_to_remove,omitempty"`
	ReplaceItemsList       *bool                                                 `json:"replace_items_list,omitempty"`
	Discounts              []*EstimateUpdateSubscriptionForItemsDiscount         `json:"discounts,omitempty"`
	ItemTiers              []*EstimateUpdateSubscriptionForItemsItemTier         `json:"item_tiers,omitempty"`
	InvoiceDate            *int64                                                `json:"invoice_date,omitempty"`
	BillingCycles          *int32                                                `json:"billing_cycles,omitempty"`
	TermsToCharge          *int32                                                `json:"terms_to_charge,omitempty"`
	ReactivateFrom         *int64                                                `json:"reactivate_from,omitempty"`
	BillingAlignmentMode   EstimateBillingAlignmentMode                          `json:"billing_alignment_mode,omitempty"`
	CouponIds              []string                                              `json:"coupon_ids,omitempty"`
	ReplaceCouponList      *bool                                                 `json:"replace_coupon_list,omitempty"`
	Prorate                *bool                                                 `json:"prorate,omitempty"`
	EndOfTerm              *bool                                                 `json:"end_of_term,omitempty"`
	ForceTermReset         *bool                                                 `json:"force_term_reset,omitempty"`
	Reactivate             *bool                                                 `json:"reactivate,omitempty"`
	IncludeDelayedCharges  *bool                                                 `json:"include_delayed_charges,omitempty"`
	UseExistingBalances    *bool                                                 `json:"use_existing_balances,omitempty"`
	BillingAddress         *EstimateUpdateSubscriptionForItemsBillingAddress     `json:"billing_address,omitempty"`
	ShippingAddress        *EstimateUpdateSubscriptionForItemsShippingAddress    `json:"shipping_address,omitempty"`
	Customer               *EstimateUpdateSubscriptionForItemsCustomer           `json:"customer,omitempty"`
	InvoiceImmediately     *bool                                                 `json:"invoice_immediately,omitempty"`
	InvoiceUsages          *bool                                                 `json:"invoice_usages,omitempty"`
	BillingOverride        *EstimateUpdateSubscriptionForItemsBillingOverride    `json:"billing_override,omitempty"`
	apiRequest             `json:"-" form:"-"`
}

func (r *EstimateUpdateSubscriptionForItemsRequest) payload() any { return r }

// input sub resource params single
type EstimateUpdateSubscriptionForItemsSubscription struct {
	Id                   string               `json:"id"`
	SetupFee             *int64               `json:"setup_fee,omitempty"`
	StartDate            *int64               `json:"start_date,omitempty"`
	TrialEnd             *int64               `json:"trial_end,omitempty"`
	Coupon               string               `json:"coupon,omitempty"`
	AutoCollection       AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	FreePeriod           *int32               `json:"free_period,omitempty"`
	FreePeriodUnit       FreePeriodUnit       `json:"free_period_unit,omitempty"`
	TrialEndAction       TrialEndAction       `json:"trial_end_action,omitempty"`
}

// input sub resource params multi
type EstimateUpdateSubscriptionForItemsSubscriptionItem struct {
	ItemPriceId        string         `json:"item_price_id"`
	Quantity           *int32         `json:"quantity,omitempty"`
	QuantityInDecimal  string         `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64         `json:"unit_price,omitempty"`
	UnitPriceInDecimal string         `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32         `json:"billing_cycles,omitempty"`
	TrialEnd           *int64         `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32         `json:"service_period_days,omitempty"`
	ChargeOnEvent      ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool          `json:"charge_once,omitempty"`
	ChargeOnOption     ChargeOnOption `json:"charge_on_option,omitempty"`
	ItemType           ItemType       `json:"item_type,omitempty"`
	ProrationType      ProrationType  `json:"proration_type,omitempty"`
}

// input sub resource params multi
type EstimateUpdateSubscriptionForItemsDiscount struct {
	ApplyOn       ApplyOn       `json:"apply_on,omitempty"`
	DurationType  DurationType  `json:"duration_type"`
	Percentage    *float64      `json:"percentage,omitempty"`
	Amount        *int64        `json:"amount,omitempty"`
	Period        *int32        `json:"period,omitempty"`
	PeriodUnit    PeriodUnit    `json:"period_unit,omitempty"`
	IncludedInMrr *bool         `json:"included_in_mrr,omitempty"`
	ItemPriceId   string        `json:"item_price_id,omitempty"`
	Quantity      *int32        `json:"quantity,omitempty"`
	OperationType OperationType `json:"operation_type"`
	Id            string        `json:"id,omitempty"`
}

// input sub resource params multi
type EstimateUpdateSubscriptionForItemsItemTier struct {
	ItemPriceId           string      `json:"item_price_id,omitempty"`
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
type EstimateUpdateSubscriptionForItemsBillingAddress struct {
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type EstimateUpdateSubscriptionForItemsShippingAddress struct {
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type EstimateUpdateSubscriptionForItemsCustomer struct {
	VatNumber        string     `json:"vat_number,omitempty"`
	VatNumberPrefix  string     `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool      `json:"registered_for_gst,omitempty"`
	Taxability       Taxability `json:"taxability,omitempty"`
}

// input sub resource params single
type EstimateUpdateSubscriptionForItemsBillingOverride struct {
	MaxExcessPaymentUsage     *int64 `json:"max_excess_payment_usage,omitempty"`
	MaxRefundableCreditsUsage *int64 `json:"max_refundable_credits_usage,omitempty"`
}
type EstimateRenewalEstimateRequest struct {
	IncludeDelayedCharges       *bool `json:"include_delayed_charges,omitempty"`
	UseExistingBalances         *bool `json:"use_existing_balances,omitempty"`
	IgnoreScheduledCancellation *bool `json:"ignore_scheduled_cancellation,omitempty"`
	IgnoreScheduledChanges      *bool `json:"ignore_scheduled_changes,omitempty"`
	apiRequest                  `json:"-" form:"-"`
}

func (r *EstimateRenewalEstimateRequest) payload() any { return r }

type EstimateAdvanceInvoiceEstimateRequest struct {
	TermsToCharge         *int32                                                 `json:"terms_to_charge,omitempty"`
	SpecificDatesSchedule []*EstimateAdvanceInvoiceEstimateSpecificDatesSchedule `json:"specific_dates_schedule,omitempty"`
	FixedIntervalSchedule *EstimateAdvanceInvoiceEstimateFixedIntervalSchedule   `json:"fixed_interval_schedule,omitempty"`
	InvoiceImmediately    *bool                                                  `json:"invoice_immediately,omitempty"`
	ScheduleType          EstimateScheduleType                                   `json:"schedule_type,omitempty"`
	apiRequest            `json:"-" form:"-"`
}

func (r *EstimateAdvanceInvoiceEstimateRequest) payload() any { return r }

// input sub resource params multi
type EstimateAdvanceInvoiceEstimateSpecificDatesSchedule struct {
	TermsToCharge *int32 `json:"terms_to_charge,omitempty"`
	Date          *int64 `json:"date,omitempty"`
}

// input sub resource params single
type EstimateAdvanceInvoiceEstimateFixedIntervalSchedule struct {
	NumberOfOccurrences *int32        `json:"number_of_occurrences,omitempty"`
	DaysBeforeRenewal   *int32        `json:"days_before_renewal,omitempty"`
	EndScheduleOn       EndScheduleOn `json:"end_schedule_on,omitempty"`
	EndDate             *int64        `json:"end_date,omitempty"`
}
type EstimateRegenerateInvoiceEstimateRequest struct {
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
	Prorate            *bool  `json:"prorate,omitempty"`
	InvoiceImmediately *bool  `json:"invoice_immediately,omitempty"`
	apiRequest         `json:"-" form:"-"`
}

func (r *EstimateRegenerateInvoiceEstimateRequest) payload() any { return r }

type EstimateChangeTermEndRequest struct {
	TermEndsAt         *int64 `json:"term_ends_at"`
	Prorate            *bool  `json:"prorate,omitempty"`
	InvoiceImmediately *bool  `json:"invoice_immediately,omitempty"`
	apiRequest         `json:"-" form:"-"`
}

func (r *EstimateChangeTermEndRequest) payload() any { return r }

type EstimateCancelSubscriptionRequest struct {
	CancelOption                      EstimateCancelOption                         `json:"cancel_option,omitempty"`
	EndOfTerm                         *bool                                        `json:"end_of_term,omitempty"`
	CancelAt                          *int64                                       `json:"cancel_at,omitempty"`
	CreditOptionForCurrentTermCharges EstimateCreditOptionForCurrentTermCharges    `json:"credit_option_for_current_term_charges,omitempty"`
	UnbilledChargesOption             EstimateUnbilledChargesOption                `json:"unbilled_charges_option,omitempty"`
	AccountReceivablesHandling        EstimateAccountReceivablesHandling           `json:"account_receivables_handling,omitempty"`
	RefundableCreditsHandling         EstimateRefundableCreditsHandling            `json:"refundable_credits_handling,omitempty"`
	ContractTermCancelOption          EstimateContractTermCancelOption             `json:"contract_term_cancel_option,omitempty"`
	InvoiceDate                       *int64                                       `json:"invoice_date,omitempty"`
	EventBasedAddons                  []*EstimateCancelSubscriptionEventBasedAddon `json:"event_based_addons,omitempty"`
	CancelReasonCode                  string                                       `json:"cancel_reason_code,omitempty"`
	apiRequest                        `json:"-" form:"-"`
}

func (r *EstimateCancelSubscriptionRequest) payload() any { return r }

// input sub resource params multi
type EstimateCancelSubscriptionEventBasedAddon struct {
	Id                  string `json:"id,omitempty"`
	Quantity            *int32 `json:"quantity,omitempty"`
	UnitPrice           *int64 `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32 `json:"service_period_in_days,omitempty"`
}
type EstimateCancelSubscriptionForItemsRequest struct {
	CancelOption                      EstimateCancelOption                                  `json:"cancel_option,omitempty"`
	EndOfTerm                         *bool                                                 `json:"end_of_term,omitempty"`
	CancelAt                          *int64                                                `json:"cancel_at,omitempty"`
	CreditOptionForCurrentTermCharges EstimateCreditOptionForCurrentTermCharges             `json:"credit_option_for_current_term_charges,omitempty"`
	UnbilledChargesOption             EstimateUnbilledChargesOption                         `json:"unbilled_charges_option,omitempty"`
	AccountReceivablesHandling        EstimateAccountReceivablesHandling                    `json:"account_receivables_handling,omitempty"`
	RefundableCreditsHandling         EstimateRefundableCreditsHandling                     `json:"refundable_credits_handling,omitempty"`
	ContractTermCancelOption          EstimateContractTermCancelOption                      `json:"contract_term_cancel_option,omitempty"`
	InvoiceDate                       *int64                                                `json:"invoice_date,omitempty"`
	SubscriptionItems                 []*EstimateCancelSubscriptionForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	CancelReasonCode                  string                                                `json:"cancel_reason_code,omitempty"`
	apiRequest                        `json:"-" form:"-"`
}

func (r *EstimateCancelSubscriptionForItemsRequest) payload() any { return r }

// input sub resource params multi
type EstimateCancelSubscriptionForItemsSubscriptionItem struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodDays  *int32 `json:"service_period_days,omitempty"`
}
type EstimatePauseSubscriptionRequest struct {
	PauseOption             EstimatePauseOption                    `json:"pause_option,omitempty"`
	Subscription            *EstimatePauseSubscriptionSubscription `json:"subscription,omitempty"`
	UnbilledChargesHandling EstimateUnbilledChargesHandling        `json:"unbilled_charges_handling,omitempty"`
	apiRequest              `json:"-" form:"-"`
}

func (r *EstimatePauseSubscriptionRequest) payload() any { return r }

// input sub resource params single
type EstimatePauseSubscriptionSubscription struct {
	PauseDate         *int64 `json:"pause_date,omitempty"`
	ResumeDate        *int64 `json:"resume_date,omitempty"`
	SkipBillingCycles *int32 `json:"skip_billing_cycles,omitempty"`
}
type EstimateResumeSubscriptionRequest struct {
	ResumeOption    EstimateResumeOption                    `json:"resume_option,omitempty"`
	Subscription    *EstimateResumeSubscriptionSubscription `json:"subscription,omitempty"`
	ChargesHandling EstimateChargesHandling                 `json:"charges_handling,omitempty"`
	apiRequest      `json:"-" form:"-"`
}

func (r *EstimateResumeSubscriptionRequest) payload() any { return r }

// input sub resource params single
type EstimateResumeSubscriptionSubscription struct {
	ResumeDate *int64 `json:"resume_date,omitempty"`
}
type EstimateGiftSubscriptionRequest struct {
	Gift            *EstimateGiftSubscriptionGift            `json:"gift,omitempty"`
	Gifter          *EstimateGiftSubscriptionGifter          `json:"gifter,omitempty"`
	GiftReceiver    *EstimateGiftSubscriptionGiftReceiver    `json:"gift_receiver,omitempty"`
	CouponIds       []string                                 `json:"coupon_ids,omitempty"`
	PaymentIntent   *EstimateGiftSubscriptionPaymentIntent   `json:"payment_intent,omitempty"`
	ShippingAddress *EstimateGiftSubscriptionShippingAddress `json:"shipping_address,omitempty"`
	Subscription    *EstimateGiftSubscriptionSubscription    `json:"subscription,omitempty"`
	Addons          []*EstimateGiftSubscriptionAddon         `json:"addons,omitempty"`
	apiRequest      `json:"-" form:"-"`
}

func (r *EstimateGiftSubscriptionRequest) payload() any { return r }

// input sub resource params single
type EstimateGiftSubscriptionGift struct {
	ScheduledAt     *int64 `json:"scheduled_at,omitempty"`
	AutoClaim       *bool  `json:"auto_claim,omitempty"`
	NoExpiry        *bool  `json:"no_expiry,omitempty"`
	ClaimExpiryDate *int64 `json:"claim_expiry_date,omitempty"`
}

// input sub resource params single
type EstimateGiftSubscriptionGifter struct {
	CustomerId   string `json:"customer_id"`
	Signature    string `json:"signature"`
	Note         string `json:"note,omitempty"`
	PaymentSrcId string `json:"payment_src_id,omitempty"`
}

// input sub resource params single
type EstimateGiftSubscriptionGiftReceiver struct {
	CustomerId string `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
}

// input sub resource params single
type EstimateGiftSubscriptionPaymentIntent struct {
	Id                    string                 `json:"id,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	GwToken               string                 `json:"gw_token,omitempty"`
	PaymentMethodType     PaymentMethodType      `json:"payment_method_type,omitempty"`
	ReferenceId           string                 `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                 `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}

// input sub resource params single
type EstimateGiftSubscriptionShippingAddress struct {
	FirstName        string           `json:"first_name,omitempty"`
	LastName         string           `json:"last_name,omitempty"`
	Email            string           `json:"email,omitempty"`
	Company          string           `json:"company,omitempty"`
	Phone            string           `json:"phone,omitempty"`
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	State            string           `json:"state,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type EstimateGiftSubscriptionSubscription struct {
	PlanId                string `json:"plan_id"`
	PlanQuantity          *int32 `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal string `json:"plan_quantity_in_decimal,omitempty"`
}

// input sub resource params multi
type EstimateGiftSubscriptionAddon struct {
	Id                string `json:"id,omitempty"`
	Quantity          *int32 `json:"quantity,omitempty"`
	QuantityInDecimal string `json:"quantity_in_decimal,omitempty"`
}
type EstimateGiftSubscriptionForItemsRequest struct {
	Gift              *EstimateGiftSubscriptionForItemsGift               `json:"gift,omitempty"`
	Gifter            *EstimateGiftSubscriptionForItemsGifter             `json:"gifter,omitempty"`
	GiftReceiver      *EstimateGiftSubscriptionForItemsGiftReceiver       `json:"gift_receiver,omitempty"`
	CouponIds         []string                                            `json:"coupon_ids,omitempty"`
	PaymentIntent     *EstimateGiftSubscriptionForItemsPaymentIntent      `json:"payment_intent,omitempty"`
	ShippingAddress   *EstimateGiftSubscriptionForItemsShippingAddress    `json:"shipping_address,omitempty"`
	SubscriptionItems []*EstimateGiftSubscriptionForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	ItemTiers         []*EstimateGiftSubscriptionForItemsItemTier         `json:"item_tiers,omitempty"`
	apiRequest        `json:"-" form:"-"`
}

func (r *EstimateGiftSubscriptionForItemsRequest) payload() any { return r }

// input sub resource params single
type EstimateGiftSubscriptionForItemsGift struct {
	ScheduledAt     *int64 `json:"scheduled_at,omitempty"`
	AutoClaim       *bool  `json:"auto_claim,omitempty"`
	NoExpiry        *bool  `json:"no_expiry,omitempty"`
	ClaimExpiryDate *int64 `json:"claim_expiry_date,omitempty"`
}

// input sub resource params single
type EstimateGiftSubscriptionForItemsGifter struct {
	CustomerId   string `json:"customer_id"`
	Signature    string `json:"signature"`
	Note         string `json:"note,omitempty"`
	PaymentSrcId string `json:"payment_src_id,omitempty"`
}

// input sub resource params single
type EstimateGiftSubscriptionForItemsGiftReceiver struct {
	CustomerId string `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
}

// input sub resource params single
type EstimateGiftSubscriptionForItemsPaymentIntent struct {
	Id                    string                 `json:"id,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	GwToken               string                 `json:"gw_token,omitempty"`
	PaymentMethodType     PaymentMethodType      `json:"payment_method_type,omitempty"`
	ReferenceId           string                 `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                 `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}

// input sub resource params single
type EstimateGiftSubscriptionForItemsShippingAddress struct {
	FirstName        string           `json:"first_name,omitempty"`
	LastName         string           `json:"last_name,omitempty"`
	Email            string           `json:"email,omitempty"`
	Company          string           `json:"company,omitempty"`
	Phone            string           `json:"phone,omitempty"`
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	State            string           `json:"state,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params multi
type EstimateGiftSubscriptionForItemsSubscriptionItem struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
}

// input sub resource params multi
type EstimateGiftSubscriptionForItemsItemTier struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int64 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type EstimateCreateInvoiceRequest struct {
	Invoice                    *EstimateCreateInvoiceInvoice             `json:"invoice,omitempty"`
	CurrencyCode               string                                    `json:"currency_code,omitempty"`
	Addons                     []*EstimateCreateInvoiceAddon             `json:"addons,omitempty"`
	Charges                    []*EstimateCreateInvoiceCharge            `json:"charges,omitempty"`
	InvoiceNote                string                                    `json:"invoice_note,omitempty"`
	RemoveGeneralNote          *bool                                     `json:"remove_general_note,omitempty"`
	NotesToRemove              []*EstimateCreateInvoiceNotesToRemove     `json:"notes_to_remove,omitempty"`
	CouponIds                  []string                                  `json:"coupon_ids,omitempty"`
	AuthorizationTransactionId string                                    `json:"authorization_transaction_id,omitempty"`
	PaymentSourceId            string                                    `json:"payment_source_id,omitempty"`
	AutoCollection             EstimateAutoCollection                    `json:"auto_collection,omitempty"`
	InvoiceDate                *int64                                    `json:"invoice_date,omitempty"`
	ShippingAddress            *EstimateCreateInvoiceShippingAddress     `json:"shipping_address,omitempty"`
	TaxProvidersFields         []*EstimateCreateInvoiceTaxProvidersField `json:"tax_providers_fields,omitempty"`
	apiRequest                 `json:"-" form:"-"`
}

func (r *EstimateCreateInvoiceRequest) payload() any { return r }

// input sub resource params single
type EstimateCreateInvoiceInvoice struct {
	CustomerId     string `json:"customer_id,omitempty"`
	SubscriptionId string `json:"subscription_id,omitempty"`
	PoNumber       string `json:"po_number,omitempty"`
}

// input sub resource params multi
type EstimateCreateInvoiceAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}

// input sub resource params multi
type EstimateCreateInvoiceCharge struct {
	Amount                 *int64          `json:"amount,omitempty"`
	AmountInDecimal        string          `json:"amount_in_decimal,omitempty"`
	Description            string          `json:"description,omitempty"`
	Taxable                *bool           `json:"taxable,omitempty"`
	TaxProfileId           string          `json:"tax_profile_id,omitempty"`
	AvalaraTaxCode         string          `json:"avalara_tax_code,omitempty"`
	HsnCode                string          `json:"hsn_code,omitempty"`
	TaxjarProductCode      string          `json:"taxjar_product_code,omitempty"`
	AvalaraSaleType        AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32          `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32          `json:"avalara_service_type,omitempty"`
	DateFrom               *int64          `json:"date_from,omitempty"`
	DateTo                 *int64          `json:"date_to,omitempty"`
}

// input sub resource params multi
type EstimateCreateInvoiceNotesToRemove struct {
	EntityType EntityType `json:"entity_type,omitempty"`
	EntityId   string     `json:"entity_id,omitempty"`
}

// input sub resource params single
type EstimateCreateInvoiceShippingAddress struct {
	FirstName        string           `json:"first_name,omitempty"`
	LastName         string           `json:"last_name,omitempty"`
	Email            string           `json:"email,omitempty"`
	Company          string           `json:"company,omitempty"`
	Phone            string           `json:"phone,omitempty"`
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	State            string           `json:"state,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params multi
type EstimateCreateInvoiceTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type EstimateCreateInvoiceForItemsRequest struct {
	Invoice                    *EstimateCreateInvoiceForItemsInvoice             `json:"invoice,omitempty"`
	CurrencyCode               string                                            `json:"currency_code,omitempty"`
	ItemPrices                 []*EstimateCreateInvoiceForItemsItemPrice         `json:"item_prices,omitempty"`
	ItemTiers                  []*EstimateCreateInvoiceForItemsItemTier          `json:"item_tiers,omitempty"`
	Charges                    []*EstimateCreateInvoiceForItemsCharge            `json:"charges,omitempty"`
	InvoiceNote                string                                            `json:"invoice_note,omitempty"`
	RemoveGeneralNote          *bool                                             `json:"remove_general_note,omitempty"`
	NotesToRemove              []*EstimateCreateInvoiceForItemsNotesToRemove     `json:"notes_to_remove,omitempty"`
	CouponIds                  []string                                          `json:"coupon_ids,omitempty"`
	AuthorizationTransactionId string                                            `json:"authorization_transaction_id,omitempty"`
	PaymentSourceId            string                                            `json:"payment_source_id,omitempty"`
	AutoCollection             EstimateAutoCollection                            `json:"auto_collection,omitempty"`
	Discounts                  []*EstimateCreateInvoiceForItemsDiscount          `json:"discounts,omitempty"`
	ShippingAddress            *EstimateCreateInvoiceForItemsShippingAddress     `json:"shipping_address,omitempty"`
	TaxProvidersFields         []*EstimateCreateInvoiceForItemsTaxProvidersField `json:"tax_providers_fields,omitempty"`
	InvoiceDate                *int64                                            `json:"invoice_date,omitempty"`
	BillingAddress             *EstimateCreateInvoiceForItemsBillingAddress      `json:"billing_address,omitempty"`
	apiRequest                 `json:"-" form:"-"`
}

func (r *EstimateCreateInvoiceForItemsRequest) payload() any { return r }

// input sub resource params single
type EstimateCreateInvoiceForItemsInvoice struct {
	CustomerId     string `json:"customer_id,omitempty"`
	SubscriptionId string `json:"subscription_id,omitempty"`
	PoNumber       string `json:"po_number,omitempty"`
}

// input sub resource params multi
type EstimateCreateInvoiceForItemsItemPrice struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}

// input sub resource params multi
type EstimateCreateInvoiceForItemsItemTier struct {
	ItemPriceId           string      `json:"item_price_id,omitempty"`
	StartingUnit          *int32      `json:"starting_unit,omitempty"`
	EndingUnit            *int32      `json:"ending_unit,omitempty"`
	Price                 *int64      `json:"price,omitempty"`
	StartingUnitInDecimal string      `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string      `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string      `json:"price_in_decimal,omitempty"`
	PricingType           PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32      `json:"package_size,omitempty"`
}

// input sub resource params multi
type EstimateCreateInvoiceForItemsCharge struct {
	Amount                 *int64          `json:"amount,omitempty"`
	AmountInDecimal        string          `json:"amount_in_decimal,omitempty"`
	Description            string          `json:"description,omitempty"`
	Taxable                *bool           `json:"taxable,omitempty"`
	TaxProfileId           string          `json:"tax_profile_id,omitempty"`
	AvalaraTaxCode         string          `json:"avalara_tax_code,omitempty"`
	HsnCode                string          `json:"hsn_code,omitempty"`
	TaxjarProductCode      string          `json:"taxjar_product_code,omitempty"`
	AvalaraSaleType        AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32          `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32          `json:"avalara_service_type,omitempty"`
	DateFrom               *int64          `json:"date_from,omitempty"`
	DateTo                 *int64          `json:"date_to,omitempty"`
}

// input sub resource params multi
type EstimateCreateInvoiceForItemsNotesToRemove struct {
	EntityType EntityType `json:"entity_type,omitempty"`
	EntityId   string     `json:"entity_id,omitempty"`
}

// input sub resource params multi
type EstimateCreateInvoiceForItemsDiscount struct {
	Percentage  *float64 `json:"percentage,omitempty"`
	Amount      *int64   `json:"amount,omitempty"`
	Quantity    *int32   `json:"quantity,omitempty"`
	ApplyOn     ApplyOn  `json:"apply_on"`
	ItemPriceId string   `json:"item_price_id,omitempty"`
}

// input sub resource params single
type EstimateCreateInvoiceForItemsShippingAddress struct {
	FirstName        string           `json:"first_name,omitempty"`
	LastName         string           `json:"last_name,omitempty"`
	Email            string           `json:"email,omitempty"`
	Company          string           `json:"company,omitempty"`
	Phone            string           `json:"phone,omitempty"`
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	State            string           `json:"state,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params multi
type EstimateCreateInvoiceForItemsTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}

// input sub resource params single
type EstimateCreateInvoiceForItemsBillingAddress struct {
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}
type EstimatePaymentSchedulesRequest struct {
	SchemeId                 string `json:"scheme_id"`
	Amount                   *int64 `json:"amount,omitempty"`
	InvoiceId                string `json:"invoice_id,omitempty"`
	PaymentScheduleStartDate *int64 `json:"payment_schedule_start_date,omitempty"`
	apiRequest               `json:"-" form:"-"`
}

func (r *EstimatePaymentSchedulesRequest) payload() any { return r }

// operation response
type EstimateCreateSubscriptionResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type EstimateCreateSubItemEstimateResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type EstimateCreateSubForCustomerEstimateResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type EstimateCreateSubItemForCustomerEstimateResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type EstimateUpdateSubscriptionResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type EstimateUpdateSubscriptionForItemsResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type EstimateRenewalEstimateResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type EstimateAdvanceInvoiceEstimateResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type EstimateRegenerateInvoiceEstimateResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type EstimateUpcomingInvoicesEstimateResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type EstimateChangeTermEndResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type EstimateCancelSubscriptionResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type EstimateCancelSubscriptionForItemsResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type EstimatePauseSubscriptionResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type EstimateResumeSubscriptionResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type EstimateGiftSubscriptionResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type EstimateGiftSubscriptionForItemsResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type EstimateCreateInvoiceResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type EstimateCreateInvoiceForItemsResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type EstimatePaymentSchedulesResponse struct {
	Estimate *Estimate `json:"estimate,omitempty"`
	apiResponse
}
