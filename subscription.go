package chargebee

import (
	"encoding/json"
)

type SubscriptionStatus string

const (
	SubscriptionStatusFuture      SubscriptionStatus = "future"
	SubscriptionStatusInTrial     SubscriptionStatus = "in_trial"
	SubscriptionStatusActive      SubscriptionStatus = "active"
	SubscriptionStatusNonRenewing SubscriptionStatus = "non_renewing"
	SubscriptionStatusPaused      SubscriptionStatus = "paused"
	SubscriptionStatusCancelled   SubscriptionStatus = "cancelled"
	SubscriptionStatusTransferred SubscriptionStatus = "transferred"
)

type SubscriptionTrialEndAction string

const (
	SubscriptionTrialEndActionSiteDefault          SubscriptionTrialEndAction = "site_default"
	SubscriptionTrialEndActionPlanDefault          SubscriptionTrialEndAction = "plan_default"
	SubscriptionTrialEndActionActivateSubscription SubscriptionTrialEndAction = "activate_subscription"
	SubscriptionTrialEndActionCancelSubscription   SubscriptionTrialEndAction = "cancel_subscription"
)

type SubscriptionCancelReason string

const (
	SubscriptionCancelReasonNotPaid                         SubscriptionCancelReason = "not_paid"
	SubscriptionCancelReasonNoCard                          SubscriptionCancelReason = "no_card"
	SubscriptionCancelReasonFraudReviewFailed               SubscriptionCancelReason = "fraud_review_failed"
	SubscriptionCancelReasonNonCompliantEuCustomer          SubscriptionCancelReason = "non_compliant_eu_customer"
	SubscriptionCancelReasonTaxCalculationFailed            SubscriptionCancelReason = "tax_calculation_failed"
	SubscriptionCancelReasonCurrencyIncompatibleWithGateway SubscriptionCancelReason = "currency_incompatible_with_gateway"
	SubscriptionCancelReasonNonCompliantCustomer            SubscriptionCancelReason = "non_compliant_customer"
)

type SubscriptionOfflinePaymentMethod string

const (
	SubscriptionOfflinePaymentMethodNoPreference            SubscriptionOfflinePaymentMethod = "no_preference"
	SubscriptionOfflinePaymentMethodCash                    SubscriptionOfflinePaymentMethod = "cash"
	SubscriptionOfflinePaymentMethodCheck                   SubscriptionOfflinePaymentMethod = "check"
	SubscriptionOfflinePaymentMethodBankTransfer            SubscriptionOfflinePaymentMethod = "bank_transfer"
	SubscriptionOfflinePaymentMethodAchCredit               SubscriptionOfflinePaymentMethod = "ach_credit"
	SubscriptionOfflinePaymentMethodSepaCredit              SubscriptionOfflinePaymentMethod = "sepa_credit"
	SubscriptionOfflinePaymentMethodBoleto                  SubscriptionOfflinePaymentMethod = "boleto"
	SubscriptionOfflinePaymentMethodUsAutomatedBankTransfer SubscriptionOfflinePaymentMethod = "us_automated_bank_transfer"
	SubscriptionOfflinePaymentMethodEuAutomatedBankTransfer SubscriptionOfflinePaymentMethod = "eu_automated_bank_transfer"
	SubscriptionOfflinePaymentMethodUkAutomatedBankTransfer SubscriptionOfflinePaymentMethod = "uk_automated_bank_transfer"
	SubscriptionOfflinePaymentMethodJpAutomatedBankTransfer SubscriptionOfflinePaymentMethod = "jp_automated_bank_transfer"
	SubscriptionOfflinePaymentMethodMxAutomatedBankTransfer SubscriptionOfflinePaymentMethod = "mx_automated_bank_transfer"
	SubscriptionOfflinePaymentMethodCustom                  SubscriptionOfflinePaymentMethod = "custom"
)

type SubscriptionChannel string

const (
	SubscriptionChannelWeb       SubscriptionChannel = "web"
	SubscriptionChannelAppStore  SubscriptionChannel = "app_store"
	SubscriptionChannelPlayStore SubscriptionChannel = "play_store"
)

type SubscriptionFreePeriodUnit string

const (
	SubscriptionFreePeriodUnitDay   SubscriptionFreePeriodUnit = "day"
	SubscriptionFreePeriodUnitWeek  SubscriptionFreePeriodUnit = "week"
	SubscriptionFreePeriodUnitMonth SubscriptionFreePeriodUnit = "month"
	SubscriptionFreePeriodUnitYear  SubscriptionFreePeriodUnit = "year"
)

type SubscriptionBillingPeriodUnit string

const (
	SubscriptionBillingPeriodUnitDay   SubscriptionBillingPeriodUnit = "day"
	SubscriptionBillingPeriodUnitWeek  SubscriptionBillingPeriodUnit = "week"
	SubscriptionBillingPeriodUnitMonth SubscriptionBillingPeriodUnit = "month"
	SubscriptionBillingPeriodUnitYear  SubscriptionBillingPeriodUnit = "year"
)

type SubscriptionAutoCollection string

const (
	SubscriptionAutoCollectionOn  SubscriptionAutoCollection = "on"
	SubscriptionAutoCollectionOff SubscriptionAutoCollection = "off"
)

type SubscriptionSubscriptionItemItemType string

const (
	SubscriptionSubscriptionItemItemTypePlan   SubscriptionSubscriptionItemItemType = "plan"
	SubscriptionSubscriptionItemItemTypeAddon  SubscriptionSubscriptionItemItemType = "addon"
	SubscriptionSubscriptionItemItemTypeCharge SubscriptionSubscriptionItemItemType = "charge"
)

type SubscriptionSubscriptionItemBillingPeriodUnit string

const (
	SubscriptionSubscriptionItemBillingPeriodUnitDay   SubscriptionSubscriptionItemBillingPeriodUnit = "day"
	SubscriptionSubscriptionItemBillingPeriodUnitWeek  SubscriptionSubscriptionItemBillingPeriodUnit = "week"
	SubscriptionSubscriptionItemBillingPeriodUnitMonth SubscriptionSubscriptionItemBillingPeriodUnit = "month"
	SubscriptionSubscriptionItemBillingPeriodUnitYear  SubscriptionSubscriptionItemBillingPeriodUnit = "year"
)

type SubscriptionSubscriptionItemChargeOnEvent string

const (
	SubscriptionSubscriptionItemChargeOnEventSubscriptionCreation   SubscriptionSubscriptionItemChargeOnEvent = "subscription_creation"
	SubscriptionSubscriptionItemChargeOnEventSubscriptionTrialStart SubscriptionSubscriptionItemChargeOnEvent = "subscription_trial_start"
	SubscriptionSubscriptionItemChargeOnEventPlanActivation         SubscriptionSubscriptionItemChargeOnEvent = "plan_activation"
	SubscriptionSubscriptionItemChargeOnEventSubscriptionActivation SubscriptionSubscriptionItemChargeOnEvent = "subscription_activation"
	SubscriptionSubscriptionItemChargeOnEventContractTermination    SubscriptionSubscriptionItemChargeOnEvent = "contract_termination"
)

type SubscriptionSubscriptionItemChargeOnOption string

const (
	SubscriptionSubscriptionItemChargeOnOptionImmediately SubscriptionSubscriptionItemChargeOnOption = "immediately"
	SubscriptionSubscriptionItemChargeOnOptionOnEvent     SubscriptionSubscriptionItemChargeOnOption = "on_event"
)

type SubscriptionSubscriptionItemProrationType string

const (
	SubscriptionSubscriptionItemProrationTypeFullTerm    SubscriptionSubscriptionItemProrationType = "full_term"
	SubscriptionSubscriptionItemProrationTypePartialTerm SubscriptionSubscriptionItemProrationType = "partial_term"
	SubscriptionSubscriptionItemProrationTypeNone        SubscriptionSubscriptionItemProrationType = "none"
)

type SubscriptionSubscriptionItemUsageAccumulationResetFrequency string

const (
	SubscriptionSubscriptionItemUsageAccumulationResetFrequencyNever                        SubscriptionSubscriptionItemUsageAccumulationResetFrequency = "never"
	SubscriptionSubscriptionItemUsageAccumulationResetFrequencySubscriptionBillingFrequency SubscriptionSubscriptionItemUsageAccumulationResetFrequency = "subscription_billing_frequency"
)

type SubscriptionItemTierPricingType string

const (
	SubscriptionItemTierPricingTypePerUnit SubscriptionItemTierPricingType = "per_unit"
	SubscriptionItemTierPricingTypeFlatFee SubscriptionItemTierPricingType = "flat_fee"
	SubscriptionItemTierPricingTypePackage SubscriptionItemTierPricingType = "package"
)

type SubscriptionShippingAddressValidationStatus string

const (
	SubscriptionShippingAddressValidationStatusNotValidated   SubscriptionShippingAddressValidationStatus = "not_validated"
	SubscriptionShippingAddressValidationStatusValid          SubscriptionShippingAddressValidationStatus = "valid"
	SubscriptionShippingAddressValidationStatusPartiallyValid SubscriptionShippingAddressValidationStatus = "partially_valid"
	SubscriptionShippingAddressValidationStatusInvalid        SubscriptionShippingAddressValidationStatus = "invalid"
)

type SubscriptionReferralInfoRewardStatus string

const (
	SubscriptionReferralInfoRewardStatusPending SubscriptionReferralInfoRewardStatus = "pending"
	SubscriptionReferralInfoRewardStatusPaid    SubscriptionReferralInfoRewardStatus = "paid"
	SubscriptionReferralInfoRewardStatusInvalid SubscriptionReferralInfoRewardStatus = "invalid"
)

type SubscriptionReferralInfoReferralSystem string

const (
	SubscriptionReferralInfoReferralSystemReferralCandy      SubscriptionReferralInfoReferralSystem = "referral_candy"
	SubscriptionReferralInfoReferralSystemReferralSaasquatch SubscriptionReferralInfoReferralSystem = "referral_saasquatch"
	SubscriptionReferralInfoReferralSystemFriendbuy          SubscriptionReferralInfoReferralSystem = "friendbuy"
)

type SubscriptionReferralInfoFriendOfferType string

const (
	SubscriptionReferralInfoFriendOfferTypeNone       SubscriptionReferralInfoFriendOfferType = "none"
	SubscriptionReferralInfoFriendOfferTypeCoupon     SubscriptionReferralInfoFriendOfferType = "coupon"
	SubscriptionReferralInfoFriendOfferTypeCouponCode SubscriptionReferralInfoFriendOfferType = "coupon_code"
)

type SubscriptionReferralInfoReferrerRewardType string

const (
	SubscriptionReferralInfoReferrerRewardTypeNone                      SubscriptionReferralInfoReferrerRewardType = "none"
	SubscriptionReferralInfoReferrerRewardTypeReferralDirectReward      SubscriptionReferralInfoReferrerRewardType = "referral_direct_reward"
	SubscriptionReferralInfoReferrerRewardTypeCustomPromotionalCredit   SubscriptionReferralInfoReferrerRewardType = "custom_promotional_credit"
	SubscriptionReferralInfoReferrerRewardTypeCustomRevenuePercentBased SubscriptionReferralInfoReferrerRewardType = "custom_revenue_percent_based"
)

type SubscriptionReferralInfoNotifyReferralSystem string

const (
	SubscriptionReferralInfoNotifyReferralSystemNone                SubscriptionReferralInfoNotifyReferralSystem = "none"
	SubscriptionReferralInfoNotifyReferralSystemFirstPaidConversion SubscriptionReferralInfoNotifyReferralSystem = "first_paid_conversion"
	SubscriptionReferralInfoNotifyReferralSystemAllInvoices         SubscriptionReferralInfoNotifyReferralSystem = "all_invoices"
)

type SubscriptionContractTermStatus string

const (
	SubscriptionContractTermStatusActive     SubscriptionContractTermStatus = "active"
	SubscriptionContractTermStatusCompleted  SubscriptionContractTermStatus = "completed"
	SubscriptionContractTermStatusCancelled  SubscriptionContractTermStatus = "cancelled"
	SubscriptionContractTermStatusTerminated SubscriptionContractTermStatus = "terminated"
)

type SubscriptionContractTermActionAtTermEnd string

const (
	SubscriptionContractTermActionAtTermEndRenew     SubscriptionContractTermActionAtTermEnd = "renew"
	SubscriptionContractTermActionAtTermEndEvergreen SubscriptionContractTermActionAtTermEnd = "evergreen"
	SubscriptionContractTermActionAtTermEndCancel    SubscriptionContractTermActionAtTermEnd = "cancel"
	SubscriptionContractTermActionAtTermEndRenewOnce SubscriptionContractTermActionAtTermEnd = "renew_once"
)

type SubscriptionDiscountType string

const (
	SubscriptionDiscountTypeFixedAmount   SubscriptionDiscountType = "fixed_amount"
	SubscriptionDiscountTypePercentage    SubscriptionDiscountType = "percentage"
	SubscriptionDiscountTypeOfferQuantity SubscriptionDiscountType = "offer_quantity"
)

type SubscriptionDiscountDurationType string

const (
	SubscriptionDiscountDurationTypeOneTime       SubscriptionDiscountDurationType = "one_time"
	SubscriptionDiscountDurationTypeForever       SubscriptionDiscountDurationType = "forever"
	SubscriptionDiscountDurationTypeLimitedPeriod SubscriptionDiscountDurationType = "limited_period"
)

type SubscriptionDiscountPeriodUnit string

const (
	SubscriptionDiscountPeriodUnitDay   SubscriptionDiscountPeriodUnit = "day"
	SubscriptionDiscountPeriodUnitWeek  SubscriptionDiscountPeriodUnit = "week"
	SubscriptionDiscountPeriodUnitMonth SubscriptionDiscountPeriodUnit = "month"
	SubscriptionDiscountPeriodUnitYear  SubscriptionDiscountPeriodUnit = "year"
)

type SubscriptionDiscountApplyOn string

const (
	SubscriptionDiscountApplyOnInvoiceAmount     SubscriptionDiscountApplyOn = "invoice_amount"
	SubscriptionDiscountApplyOnSpecificItemPrice SubscriptionDiscountApplyOn = "specific_item_price"
)

type SubscriptionAddonProrationType string

const (
	SubscriptionAddonProrationTypeFullTerm    SubscriptionAddonProrationType = "full_term"
	SubscriptionAddonProrationTypePartialTerm SubscriptionAddonProrationType = "partial_term"
	SubscriptionAddonProrationTypeNone        SubscriptionAddonProrationType = "none"
)

type SubscriptionEventBasedAddonOnEvent string

const (
	SubscriptionEventBasedAddonOnEventSubscriptionCreation   SubscriptionEventBasedAddonOnEvent = "subscription_creation"
	SubscriptionEventBasedAddonOnEventSubscriptionTrialStart SubscriptionEventBasedAddonOnEvent = "subscription_trial_start"
	SubscriptionEventBasedAddonOnEventPlanActivation         SubscriptionEventBasedAddonOnEvent = "plan_activation"
	SubscriptionEventBasedAddonOnEventSubscriptionActivation SubscriptionEventBasedAddonOnEvent = "subscription_activation"
	SubscriptionEventBasedAddonOnEventContractTermination    SubscriptionEventBasedAddonOnEvent = "contract_termination"
)

type SubscriptionBillingAlignmentMode string

const (
	SubscriptionBillingAlignmentModeImmediate SubscriptionBillingAlignmentMode = "immediate"
	SubscriptionBillingAlignmentModeDelayed   SubscriptionBillingAlignmentMode = "delayed"
)

type SubscriptionPaymentInitiator string

const (
	SubscriptionPaymentInitiatorCustomer SubscriptionPaymentInitiator = "customer"
	SubscriptionPaymentInitiatorMerchant SubscriptionPaymentInitiator = "merchant"
)

type SubscriptionChangeOption string

const (
	SubscriptionChangeOptionImmediately  SubscriptionChangeOption = "immediately"
	SubscriptionChangeOptionEndOfTerm    SubscriptionChangeOption = "end_of_term"
	SubscriptionChangeOptionSpecificDate SubscriptionChangeOption = "specific_date"
)

type SubscriptionAvalaraSaleType string

const (
	SubscriptionAvalaraSaleTypeWholesale SubscriptionAvalaraSaleType = "wholesale"
	SubscriptionAvalaraSaleTypeRetail    SubscriptionAvalaraSaleType = "retail"
	SubscriptionAvalaraSaleTypeConsumed  SubscriptionAvalaraSaleType = "consumed"
	SubscriptionAvalaraSaleTypeVendorUse SubscriptionAvalaraSaleType = "vendor_use"
)

type SubscriptionScheduleType string

const (
	SubscriptionScheduleTypeImmediate      SubscriptionScheduleType = "immediate"
	SubscriptionScheduleTypeSpecificDates  SubscriptionScheduleType = "specific_dates"
	SubscriptionScheduleTypeFixedIntervals SubscriptionScheduleType = "fixed_intervals"
)

type SubscriptionScheduleType string

const (
	SubscriptionScheduleTypeSpecificDates  SubscriptionScheduleType = "specific_dates"
	SubscriptionScheduleTypeFixedIntervals SubscriptionScheduleType = "fixed_intervals"
)

type SubscriptionPauseOption string

const (
	SubscriptionPauseOptionImmediately   SubscriptionPauseOption = "immediately"
	SubscriptionPauseOptionEndOfTerm     SubscriptionPauseOption = "end_of_term"
	SubscriptionPauseOptionSpecificDate  SubscriptionPauseOption = "specific_date"
	SubscriptionPauseOptionBillingCycles SubscriptionPauseOption = "billing_cycles"
)

type SubscriptionUnbilledChargesHandling string

const (
	SubscriptionUnbilledChargesHandlingNoAction SubscriptionUnbilledChargesHandling = "no_action"
	SubscriptionUnbilledChargesHandlingInvoice  SubscriptionUnbilledChargesHandling = "invoice"
)

type SubscriptionInvoiceDunningHandling string

const (
	SubscriptionInvoiceDunningHandlingContinue SubscriptionInvoiceDunningHandling = "continue"
	SubscriptionInvoiceDunningHandlingStop     SubscriptionInvoiceDunningHandling = "stop"
)

type SubscriptionCancelOption string

const (
	SubscriptionCancelOptionImmediately      SubscriptionCancelOption = "immediately"
	SubscriptionCancelOptionEndOfTerm        SubscriptionCancelOption = "end_of_term"
	SubscriptionCancelOptionSpecificDate     SubscriptionCancelOption = "specific_date"
	SubscriptionCancelOptionEndOfBillingTerm SubscriptionCancelOption = "end_of_billing_term"
)

type SubscriptionCreditOptionForCurrentTermCharges string

const (
	SubscriptionCreditOptionForCurrentTermChargesNone    SubscriptionCreditOptionForCurrentTermCharges = "none"
	SubscriptionCreditOptionForCurrentTermChargesProrate SubscriptionCreditOptionForCurrentTermCharges = "prorate"
	SubscriptionCreditOptionForCurrentTermChargesFull    SubscriptionCreditOptionForCurrentTermCharges = "full"
)

type SubscriptionUnbilledChargesOption string

const (
	SubscriptionUnbilledChargesOptionInvoice SubscriptionUnbilledChargesOption = "invoice"
	SubscriptionUnbilledChargesOptionDelete  SubscriptionUnbilledChargesOption = "delete"
)

type SubscriptionAccountReceivablesHandling string

const (
	SubscriptionAccountReceivablesHandlingNoAction                  SubscriptionAccountReceivablesHandling = "no_action"
	SubscriptionAccountReceivablesHandlingSchedulePaymentCollection SubscriptionAccountReceivablesHandling = "schedule_payment_collection"
	SubscriptionAccountReceivablesHandlingWriteOff                  SubscriptionAccountReceivablesHandling = "write_off"
)

type SubscriptionRefundableCreditsHandling string

const (
	SubscriptionRefundableCreditsHandlingNoAction       SubscriptionRefundableCreditsHandling = "no_action"
	SubscriptionRefundableCreditsHandlingScheduleRefund SubscriptionRefundableCreditsHandling = "schedule_refund"
)

type SubscriptionContractTermCancelOption string

const (
	SubscriptionContractTermCancelOptionTerminateImmediately         SubscriptionContractTermCancelOption = "terminate_immediately"
	SubscriptionContractTermCancelOptionEndOfContractTerm            SubscriptionContractTermCancelOption = "end_of_contract_term"
	SubscriptionContractTermCancelOptionSpecificDate                 SubscriptionContractTermCancelOption = "specific_date"
	SubscriptionContractTermCancelOptionEndOfSubscriptionBillingTerm SubscriptionContractTermCancelOption = "end_of_subscription_billing_term"
)

type SubscriptionResumeOption string

const (
	SubscriptionResumeOptionImmediately  SubscriptionResumeOption = "immediately"
	SubscriptionResumeOptionSpecificDate SubscriptionResumeOption = "specific_date"
)

type SubscriptionChargesHandling string

const (
	SubscriptionChargesHandlingInvoiceImmediately   SubscriptionChargesHandling = "invoice_immediately"
	SubscriptionChargesHandlingAddToUnbilledCharges SubscriptionChargesHandling = "add_to_unbilled_charges"
)

type SubscriptionUnpaidInvoicesHandling string

const (
	SubscriptionUnpaidInvoicesHandlingNoAction                  SubscriptionUnpaidInvoicesHandling = "no_action"
	SubscriptionUnpaidInvoicesHandlingSchedulePaymentCollection SubscriptionUnpaidInvoicesHandling = "schedule_payment_collection"
)

// just struct
type Subscription struct {
	Id                                string                                `json:"id"`
	CurrencyCode                      string                                `json:"currency_code"`
	PlanId                            string                                `json:"plan_id"`
	PlanQuantity                      int32                                 `json:"plan_quantity"`
	PlanUnitPrice                     int64                                 `json:"plan_unit_price"`
	SetupFee                          int64                                 `json:"setup_fee"`
	BillingPeriod                     int32                                 `json:"billing_period"`
	BillingPeriodUnit                 SubscriptionBillingPeriodUnit         `json:"billing_period_unit"`
	StartDate                         int64                                 `json:"start_date"`
	TrialEnd                          int64                                 `json:"trial_end"`
	RemainingBillingCycles            int32                                 `json:"remaining_billing_cycles"`
	PoNumber                          string                                `json:"po_number"`
	AutoCollection                    SubscriptionAutoCollection            `json:"auto_collection"`
	PlanQuantityInDecimal             string                                `json:"plan_quantity_in_decimal"`
	PlanUnitPriceInDecimal            string                                `json:"plan_unit_price_in_decimal"`
	CustomerId                        string                                `json:"customer_id"`
	PlanAmount                        int64                                 `json:"plan_amount"`
	PlanFreeQuantity                  int32                                 `json:"plan_free_quantity"`
	Status                            SubscriptionStatus                    `json:"status"`
	TrialStart                        int64                                 `json:"trial_start"`
	TrialEndAction                    SubscriptionTrialEndAction            `json:"trial_end_action"`
	CurrentTermStart                  int64                                 `json:"current_term_start"`
	CurrentTermEnd                    int64                                 `json:"current_term_end"`
	NextBillingAt                     int64                                 `json:"next_billing_at"`
	CreatedAt                         int64                                 `json:"created_at"`
	StartedAt                         int64                                 `json:"started_at"`
	ActivatedAt                       int64                                 `json:"activated_at"`
	GiftId                            string                                `json:"gift_id"`
	ContractTermBillingCycleOnRenewal int32                                 `json:"contract_term_billing_cycle_on_renewal"`
	OverrideRelationship              bool                                  `json:"override_relationship"`
	PauseDate                         int64                                 `json:"pause_date"`
	ResumeDate                        int64                                 `json:"resume_date"`
	CancelledAt                       int64                                 `json:"cancelled_at"`
	CancelReason                      SubscriptionCancelReason              `json:"cancel_reason"`
	AffiliateToken                    string                                `json:"affiliate_token"`
	CreatedFromIp                     string                                `json:"created_from_ip"`
	ResourceVersion                   int64                                 `json:"resource_version"`
	UpdatedAt                         int64                                 `json:"updated_at"`
	HasScheduledAdvanceInvoices       bool                                  `json:"has_scheduled_advance_invoices"`
	HasScheduledChanges               bool                                  `json:"has_scheduled_changes"`
	PaymentSourceId                   string                                `json:"payment_source_id"`
	PlanFreeQuantityInDecimal         string                                `json:"plan_free_quantity_in_decimal"`
	PlanAmountInDecimal               string                                `json:"plan_amount_in_decimal"`
	CancelScheduleCreatedAt           int64                                 `json:"cancel_schedule_created_at"`
	OfflinePaymentMethod              SubscriptionOfflinePaymentMethod      `json:"offline_payment_method"`
	Channel                           SubscriptionChannel                   `json:"channel"`
	NetTermDays                       int32                                 `json:"net_term_days"`
	ActiveId                          string                                `json:"active_id"`
	SubscriptionItems                 []*SubscriptionSubscriptionItem       `json:"subscription_items"`
	ItemTiers                         []*SubscriptionItemTier               `json:"item_tiers"`
	ChargedItems                      []*SubscriptionChargedItem            `json:"charged_items"`
	DueInvoicesCount                  int32                                 `json:"due_invoices_count"`
	DueSince                          int64                                 `json:"due_since"`
	TotalDues                         int64                                 `json:"total_dues"`
	Mrr                               int64                                 `json:"mrr"`
	ExchangeRate                      float64                               `json:"exchange_rate"`
	BaseCurrencyCode                  string                                `json:"base_currency_code"`
	Addons                            []*SubscriptionAddon                  `json:"addons"`
	EventBasedAddons                  []*SubscriptionEventBasedAddon        `json:"event_based_addons"`
	ChargedEventBasedAddons           []*SubscriptionChargedEventBasedAddon `json:"charged_event_based_addons"`
	Coupons                           []*SubscriptionCoupon                 `json:"coupons"`
	ShippingAddress                   *SubscriptionShippingAddress          `json:"shipping_address"`
	ReferralInfo                      *SubscriptionReferralInfo             `json:"referral_info"`
	BillingOverride                   *SubscriptionBillingOverride          `json:"billing_override"`
	InvoiceNotes                      string                                `json:"invoice_notes"`
	MetaData                          json.RawMessage                       `json:"meta_data"`
	Deleted                           bool                                  `json:"deleted"`
	ChangesScheduledAt                int64                                 `json:"changes_scheduled_at"`
	ContractTerm                      *SubscriptionContractTerm             `json:"contract_term"`
	CancelReasonCode                  string                                `json:"cancel_reason_code"`
	FreePeriod                        int32                                 `json:"free_period"`
	FreePeriodUnit                    SubscriptionFreePeriodUnit            `json:"free_period_unit"`
	CreatePendingInvoices             bool                                  `json:"create_pending_invoices"`
	AutoCloseInvoices                 bool                                  `json:"auto_close_invoices"`
	Discounts                         []*SubscriptionDiscount               `json:"discounts"`
	BusinessEntityId                  string                                `json:"business_entity_id"`
	CustomField                       CustomField                           `json:"custom_field"`
	Object                            string                                `json:"object"`

	// Deprecated: MetaData is deprecated please use MetaData instead.
	Metadata json.RawMessage `json:"metadata"`
}

// sub resources
type SubscriptionSubscriptionItem struct {
	ItemPriceId                     string                                                      `json:"item_price_id"`
	ItemType                        SubscriptionSubscriptionItemItemType                        `json:"item_type"`
	Quantity                        int32                                                       `json:"quantity"`
	QuantityInDecimal               string                                                      `json:"quantity_in_decimal"`
	MeteredQuantity                 string                                                      `json:"metered_quantity"`
	LastCalculatedAt                int64                                                       `json:"last_calculated_at"`
	UnitPrice                       int64                                                       `json:"unit_price"`
	UnitPriceInDecimal              string                                                      `json:"unit_price_in_decimal"`
	Amount                          int64                                                       `json:"amount"`
	CurrentTermStart                int64                                                       `json:"current_term_start"`
	CurrentTermEnd                  int64                                                       `json:"current_term_end"`
	NextBillingAt                   int64                                                       `json:"next_billing_at"`
	AmountInDecimal                 string                                                      `json:"amount_in_decimal"`
	BillingPeriod                   int32                                                       `json:"billing_period"`
	BillingPeriodUnit               SubscriptionSubscriptionItemBillingPeriodUnit               `json:"billing_period_unit"`
	FreeQuantity                    int32                                                       `json:"free_quantity"`
	FreeQuantityInDecimal           string                                                      `json:"free_quantity_in_decimal"`
	TrialEnd                        int64                                                       `json:"trial_end"`
	BillingCycles                   int32                                                       `json:"billing_cycles"`
	ServicePeriodDays               int32                                                       `json:"service_period_days"`
	ChargeOnEvent                   SubscriptionSubscriptionItemChargeOnEvent                   `json:"charge_on_event"`
	ChargeOnce                      bool                                                        `json:"charge_once"`
	ChargeOnOption                  SubscriptionSubscriptionItemChargeOnOption                  `json:"charge_on_option"`
	ProrationType                   SubscriptionSubscriptionItemProrationType                   `json:"proration_type"`
	UsageAccumulationResetFrequency SubscriptionSubscriptionItemUsageAccumulationResetFrequency `json:"usage_accumulation_reset_frequency"`
	Object                          string                                                      `json:"object"`
}

type SubscriptionItemTier struct {
	ItemPriceId           string                          `json:"item_price_id"`
	StartingUnit          int32                           `json:"starting_unit"`
	EndingUnit            int32                           `json:"ending_unit"`
	Price                 int64                           `json:"price"`
	StartingUnitInDecimal string                          `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string                          `json:"ending_unit_in_decimal"`
	PriceInDecimal        string                          `json:"price_in_decimal"`
	PricingType           SubscriptionItemTierPricingType `json:"pricing_type"`
	PackageSize           int32                           `json:"package_size"`
	Index                 int32                           `json:"index"`
	Object                string                          `json:"object"`
}

type SubscriptionChargedItem struct {
	ItemPriceId   string `json:"item_price_id"`
	LastChargedAt int64  `json:"last_charged_at"`
	Object        string `json:"object"`
}

type SubscriptionCoupon struct {
	CouponId     string `json:"coupon_id"`
	ApplyTill    int64  `json:"apply_till"`
	AppliedCount int32  `json:"applied_count"`
	CouponCode   string `json:"coupon_code"`
	Object       string `json:"object"`
}

type SubscriptionShippingAddress struct {
	FirstName        string                                      `json:"first_name"`
	LastName         string                                      `json:"last_name"`
	Email            string                                      `json:"email"`
	Company          string                                      `json:"company"`
	Phone            string                                      `json:"phone"`
	Line1            string                                      `json:"line1"`
	Line2            string                                      `json:"line2"`
	Line3            string                                      `json:"line3"`
	City             string                                      `json:"city"`
	StateCode        string                                      `json:"state_code"`
	State            string                                      `json:"state"`
	Country          string                                      `json:"country"`
	Zip              string                                      `json:"zip"`
	ValidationStatus SubscriptionShippingAddressValidationStatus `json:"validation_status"`
	Object           string                                      `json:"object"`
}

type SubscriptionReferralInfo struct {
	ReferralCode              string                                       `json:"referral_code"`
	CouponCode                string                                       `json:"coupon_code"`
	ReferrerId                string                                       `json:"referrer_id"`
	ExternalReferenceId       string                                       `json:"external_reference_id"`
	RewardStatus              SubscriptionReferralInfoRewardStatus         `json:"reward_status"`
	ReferralSystem            SubscriptionReferralInfoReferralSystem       `json:"referral_system"`
	AccountId                 string                                       `json:"account_id"`
	CampaignId                string                                       `json:"campaign_id"`
	ExternalCampaignId        string                                       `json:"external_campaign_id"`
	FriendOfferType           SubscriptionReferralInfoFriendOfferType      `json:"friend_offer_type"`
	ReferrerRewardType        SubscriptionReferralInfoReferrerRewardType   `json:"referrer_reward_type"`
	NotifyReferralSystem      SubscriptionReferralInfoNotifyReferralSystem `json:"notify_referral_system"`
	DestinationUrl            string                                       `json:"destination_url"`
	PostPurchaseWidgetEnabled bool                                         `json:"post_purchase_widget_enabled"`
	Object                    string                                       `json:"object"`
}

type SubscriptionBillingOverride struct {
	MaxExcessPaymentUsage     int64  `json:"max_excess_payment_usage"`
	MaxRefundableCreditsUsage int64  `json:"max_refundable_credits_usage"`
	Object                    string `json:"object"`
}

type SubscriptionContractTerm struct {
	Id                          string                                  `json:"id"`
	Status                      SubscriptionContractTermStatus          `json:"status"`
	ContractStart               int64                                   `json:"contract_start"`
	ContractEnd                 int64                                   `json:"contract_end"`
	BillingCycle                int32                                   `json:"billing_cycle"`
	ActionAtTermEnd             SubscriptionContractTermActionAtTermEnd `json:"action_at_term_end"`
	TotalContractValue          int64                                   `json:"total_contract_value"`
	TotalContractValueBeforeTax int64                                   `json:"total_contract_value_before_tax"`
	CancellationCutoffPeriod    int32                                   `json:"cancellation_cutoff_period"`
	CreatedAt                   int64                                   `json:"created_at"`
	SubscriptionId              string                                  `json:"subscription_id"`
	RemainingBillingCycles      int32                                   `json:"remaining_billing_cycles"`
	Object                      string                                  `json:"object"`
}

type SubscriptionDiscount struct {
	Id            string                           `json:"id"`
	InvoiceName   string                           `json:"invoice_name"`
	Type          SubscriptionDiscountType         `json:"type"`
	Percentage    float64                          `json:"percentage"`
	Amount        int64                            `json:"amount"`
	Quantity      int32                            `json:"quantity"`
	CurrencyCode  string                           `json:"currency_code"`
	DurationType  SubscriptionDiscountDurationType `json:"duration_type"`
	Period        int32                            `json:"period"`
	PeriodUnit    SubscriptionDiscountPeriodUnit   `json:"period_unit"`
	IncludedInMrr bool                             `json:"included_in_mrr"`
	ApplyOn       SubscriptionDiscountApplyOn      `json:"apply_on"`
	ItemPriceId   string                           `json:"item_price_id"`
	CreatedAt     int64                            `json:"created_at"`
	ApplyTill     int64                            `json:"apply_till"`
	AppliedCount  int32                            `json:"applied_count"`
	CouponId      string                           `json:"coupon_id"`
	Index         int32                            `json:"index"`
	Object        string                           `json:"object"`
}

type SubscriptionAddon struct {
	Id                     string                         `json:"id"`
	Quantity               int32                          `json:"quantity"`
	UnitPrice              int64                          `json:"unit_price"`
	Amount                 int64                          `json:"amount"`
	TrialEnd               int64                          `json:"trial_end"`
	RemainingBillingCycles int32                          `json:"remaining_billing_cycles"`
	QuantityInDecimal      string                         `json:"quantity_in_decimal"`
	UnitPriceInDecimal     string                         `json:"unit_price_in_decimal"`
	AmountInDecimal        string                         `json:"amount_in_decimal"`
	ProrationType          SubscriptionAddonProrationType `json:"proration_type"`
	Object                 string                         `json:"object"`
}

type SubscriptionChargedEventBasedAddon struct {
	Id            string `json:"id"`
	LastChargedAt int64  `json:"last_charged_at"`
	Object        string `json:"object"`
}

type SubscriptionEventBasedAddon struct {
	Id                  string                             `json:"id"`
	Quantity            int32                              `json:"quantity"`
	UnitPrice           int64                              `json:"unit_price"`
	ServicePeriodInDays int32                              `json:"service_period_in_days"`
	OnEvent             SubscriptionEventBasedAddonOnEvent `json:"on_event"`
	ChargeOnce          bool                               `json:"charge_once"`
	QuantityInDecimal   string                             `json:"quantity_in_decimal"`
	UnitPriceInDecimal  string                             `json:"unit_price_in_decimal"`
	Object              string                             `json:"object"`
}

// operations
// input params
type SubscriptionCreateRequest struct {
	Id                                string                                 `json:"id,omitempty"`
	Customer                          *SubscriptionCreateCustomer            `json:"customer,omitempty"`
	EntityIdentifiers                 []*SubscriptionCreateEntityIdentifier  `json:"entity_identifiers,omitempty"`
	TaxProvidersFields                []*SubscriptionCreateTaxProvidersField `json:"tax_providers_fields,omitempty"`
	PlanId                            string                                 `json:"plan_id"`
	PlanQuantity                      *int32                                 `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                                 `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int64                                 `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                                 `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int64                                 `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                                 `json:"trial_end,omitempty"`
	BillingCycles                     *int32                                 `json:"billing_cycles,omitempty"`
	Addons                            []*SubscriptionCreateAddon             `json:"addons,omitempty"`
	EventBasedAddons                  []*SubscriptionCreateEventBasedAddon   `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove           []string                               `json:"mandatory_addons_to_remove,omitempty"`
	StartDate                         *int64                                 `json:"start_date,omitempty"`
	AutoCollection                    SubscriptionAutoCollection             `json:"auto_collection,omitempty"`
	TermsToCharge                     *int32                                 `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode              SubscriptionBillingAlignmentMode       `json:"billing_alignment_mode,omitempty"`
	OfflinePaymentMethod              SubscriptionOfflinePaymentMethod       `json:"offline_payment_method,omitempty"`
	PoNumber                          string                                 `json:"po_number,omitempty"`
	CouponIds                         []string                               `json:"coupon_ids,omitempty"`
	Card                              *SubscriptionCreateCard                `json:"card,omitempty"`
	BankAccount                       *SubscriptionCreateBankAccount         `json:"bank_account,omitempty"`
	TokenId                           string                                 `json:"token_id,omitempty"`
	PaymentMethod                     *SubscriptionCreatePaymentMethod       `json:"payment_method,omitempty"`
	PaymentIntent                     *SubscriptionCreatePaymentIntent       `json:"payment_intent,omitempty"`
	BillingAddress                    *SubscriptionCreateBillingAddress      `json:"billing_address,omitempty"`
	ShippingAddress                   *SubscriptionCreateShippingAddress     `json:"shipping_address,omitempty"`
	StatementDescriptor               *SubscriptionCreateStatementDescriptor `json:"statement_descriptor,omitempty"`
	AffiliateToken                    string                                 `json:"affiliate_token,omitempty"`
	InvoiceNotes                      string                                 `json:"invoice_notes,omitempty"`
	InvoiceDate                       *int64                                 `json:"invoice_date,omitempty"`
	MetaData                          map[string]interface{}                 `json:"meta_data,omitempty"`
	InvoiceImmediately                *bool                                  `json:"invoice_immediately,omitempty"`
	FreePeriod                        *int32                                 `json:"free_period,omitempty"`
	FreePeriodUnit                    SubscriptionFreePeriodUnit             `json:"free_period_unit,omitempty"`
	ContractTerm                      *SubscriptionCreateContractTerm        `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                 `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	TrialEndAction                    SubscriptionTrialEndAction             `json:"trial_end_action,omitempty"`
	ClientProfileId                   string                                 `json:"client_profile_id,omitempty"`
	PaymentInitiator                  SubscriptionPaymentInitiator           `json:"payment_initiator,omitempty"`
	Coupons                           []*SubscriptionCreateCoupon            `json:"coupons,omitempty"`
	apiRequest                        `json:"-" form:"-"`
}

func (r *SubscriptionCreateRequest) payload() any { return r }

// input sub resource params single
type SubscriptionCreateCustomer struct {
	Id                               string                                      `json:"id,omitempty"`
	Email                            string                                      `json:"email,omitempty"`
	FirstName                        string                                      `json:"first_name,omitempty"`
	LastName                         string                                      `json:"last_name,omitempty"`
	Company                          string                                      `json:"company,omitempty"`
	Phone                            string                                      `json:"phone,omitempty"`
	Locale                           string                                      `json:"locale,omitempty"`
	Taxability                       SubscriptionCustomerTaxability              `json:"taxability,omitempty"`
	EntityCode                       SubscriptionCustomerEntityCode              `json:"entity_code,omitempty"`
	ExemptNumber                     string                                      `json:"exempt_number,omitempty"`
	NetTermDays                      *int32                                      `json:"net_term_days,omitempty"`
	TaxjarExemptionCategory          SubscriptionCustomerTaxjarExemptionCategory `json:"taxjar_exemption_category,omitempty"`
	AutoCollection                   SubscriptionCustomerAutoCollection          `json:"auto_collection,omitempty"`
	OfflinePaymentMethod             SubscriptionCustomerOfflinePaymentMethod    `json:"offline_payment_method,omitempty"`
	AllowDirectDebit                 *bool                                       `json:"allow_direct_debit,omitempty"`
	ConsolidatedInvoicing            *bool                                       `json:"consolidated_invoicing,omitempty"`
	VatNumber                        string                                      `json:"vat_number,omitempty"`
	VatNumberPrefix                  string                                      `json:"vat_number_prefix,omitempty"`
	EntityIdentifierScheme           string                                      `json:"entity_identifier_scheme,omitempty"`
	EntityIdentifierStandard         string                                      `json:"entity_identifier_standard,omitempty"`
	IsEinvoiceEnabled                *bool                                       `json:"is_einvoice_enabled,omitempty"`
	EinvoicingMethod                 SubscriptionCustomerEinvoicingMethod        `json:"einvoicing_method,omitempty"`
	RegisteredForGst                 *bool                                       `json:"registered_for_gst,omitempty"`
	BusinessCustomerWithoutVatNumber *bool                                       `json:"business_customer_without_vat_number,omitempty"`
	ExemptionDetails                 []map[string]interface{}                    `json:"exemption_details,omitempty"`
	CustomerType                     SubscriptionCustomerCustomerType            `json:"customer_type,omitempty"`
}

// input sub resource params multi
type SubscriptionCreateEntityIdentifier struct {
	Id       string `json:"id,omitempty"`
	Scheme   string `json:"scheme,omitempty"`
	Value    string `json:"value,omitempty"`
	Standard string `json:"standard,omitempty"`
}

// input sub resource params multi
type SubscriptionCreateTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}

// input sub resource params multi
type SubscriptionCreateAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}

// input sub resource params multi
type SubscriptionCreateEventBasedAddon struct {
	Id                  string                              `json:"id,omitempty"`
	Quantity            *int32                              `json:"quantity,omitempty"`
	UnitPrice           *int64                              `json:"unit_price,omitempty"`
	QuantityInDecimal   string                              `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string                              `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodInDays *int32                              `json:"service_period_in_days,omitempty"`
	OnEvent             SubscriptionEventBasedAddonOnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool                               `json:"charge_once,omitempty"`
	ChargeOn            SubscriptionEventBasedAddonChargeOn `json:"charge_on,omitempty"`
}

// input sub resource params single
type SubscriptionCreateCard struct {
	Gateway               SubscriptionCardGateway         `json:"gateway,omitempty"`
	GatewayAccountId      string                          `json:"gateway_account_id,omitempty"`
	TmpToken              string                          `json:"tmp_token,omitempty"`
	FirstName             string                          `json:"first_name,omitempty"`
	LastName              string                          `json:"last_name,omitempty"`
	Number                string                          `json:"number,omitempty"`
	ExpiryMonth           *int32                          `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                          `json:"expiry_year,omitempty"`
	Cvv                   string                          `json:"cvv,omitempty"`
	PreferredScheme       SubscriptionCardPreferredScheme `json:"preferred_scheme,omitempty"`
	BillingAddr1          string                          `json:"billing_addr1,omitempty"`
	BillingAddr2          string                          `json:"billing_addr2,omitempty"`
	BillingCity           string                          `json:"billing_city,omitempty"`
	BillingStateCode      string                          `json:"billing_state_code,omitempty"`
	BillingState          string                          `json:"billing_state,omitempty"`
	BillingZip            string                          `json:"billing_zip,omitempty"`
	BillingCountry        string                          `json:"billing_country,omitempty"`
	IpAddress             string                          `json:"ip_address,omitempty"`
	AdditionalInformation map[string]interface{}          `json:"additional_information,omitempty"`
}

// input sub resource params single
type SubscriptionCreateBankAccount struct {
	GatewayAccountId      string                                   `json:"gateway_account_id,omitempty"`
	Iban                  string                                   `json:"iban,omitempty"`
	FirstName             string                                   `json:"first_name,omitempty"`
	LastName              string                                   `json:"last_name,omitempty"`
	Company               string                                   `json:"company,omitempty"`
	Email                 string                                   `json:"email,omitempty"`
	Phone                 string                                   `json:"phone,omitempty"`
	BankName              string                                   `json:"bank_name,omitempty"`
	AccountNumber         string                                   `json:"account_number,omitempty"`
	RoutingNumber         string                                   `json:"routing_number,omitempty"`
	BankCode              string                                   `json:"bank_code,omitempty"`
	AccountType           SubscriptionBankAccountAccountType       `json:"account_type,omitempty"`
	AccountHolderType     SubscriptionBankAccountAccountHolderType `json:"account_holder_type,omitempty"`
	EcheckType            SubscriptionBankAccountEcheckType        `json:"echeck_type,omitempty"`
	IssuingCountry        string                                   `json:"issuing_country,omitempty"`
	SwedishIdentityNumber string                                   `json:"swedish_identity_number,omitempty"`
	BillingAddress        map[string]interface{}                   `json:"billing_address,omitempty"`
}

// input sub resource params single
type SubscriptionCreatePaymentMethod struct {
	Type                  SubscriptionPaymentMethodType    `json:"type,omitempty"`
	Gateway               SubscriptionPaymentMethodGateway `json:"gateway,omitempty"`
	GatewayAccountId      string                           `json:"gateway_account_id,omitempty"`
	ReferenceId           string                           `json:"reference_id,omitempty"`
	TmpToken              string                           `json:"tmp_token,omitempty"`
	IssuingCountry        string                           `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{}           `json:"additional_information,omitempty"`
}

// input sub resource params single
type SubscriptionCreatePaymentIntent struct {
	Id                    string                                     `json:"id,omitempty"`
	GatewayAccountId      string                                     `json:"gateway_account_id,omitempty"`
	GwToken               string                                     `json:"gw_token,omitempty"`
	PaymentMethodType     SubscriptionPaymentIntentPaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                                     `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                                     `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}                     `json:"additional_information,omitempty"`
}

// input sub resource params single
type SubscriptionCreateBillingAddress struct {
	FirstName        string                                     `json:"first_name,omitempty"`
	LastName         string                                     `json:"last_name,omitempty"`
	Email            string                                     `json:"email,omitempty"`
	Company          string                                     `json:"company,omitempty"`
	Phone            string                                     `json:"phone,omitempty"`
	Line1            string                                     `json:"line1,omitempty"`
	Line2            string                                     `json:"line2,omitempty"`
	Line3            string                                     `json:"line3,omitempty"`
	City             string                                     `json:"city,omitempty"`
	StateCode        string                                     `json:"state_code,omitempty"`
	State            string                                     `json:"state,omitempty"`
	Zip              string                                     `json:"zip,omitempty"`
	Country          string                                     `json:"country,omitempty"`
	ValidationStatus SubscriptionBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type SubscriptionCreateShippingAddress struct {
	FirstName        string                                      `json:"first_name,omitempty"`
	LastName         string                                      `json:"last_name,omitempty"`
	Email            string                                      `json:"email,omitempty"`
	Company          string                                      `json:"company,omitempty"`
	Phone            string                                      `json:"phone,omitempty"`
	Line1            string                                      `json:"line1,omitempty"`
	Line2            string                                      `json:"line2,omitempty"`
	Line3            string                                      `json:"line3,omitempty"`
	City             string                                      `json:"city,omitempty"`
	StateCode        string                                      `json:"state_code,omitempty"`
	State            string                                      `json:"state,omitempty"`
	Zip              string                                      `json:"zip,omitempty"`
	Country          string                                      `json:"country,omitempty"`
	ValidationStatus SubscriptionShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type SubscriptionCreateStatementDescriptor struct {
	Descriptor string `json:"descriptor,omitempty"`
}

// input sub resource params single
type SubscriptionCreateContractTerm struct {
	ActionAtTermEnd          SubscriptionContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                                  `json:"cancellation_cutoff_period,omitempty"`
}

// input sub resource params multi
type SubscriptionCreateCoupon struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}

type SubscriptionCreateForCustomerRequest struct {
	Id                                string                                            `json:"id,omitempty"`
	PlanId                            string                                            `json:"plan_id"`
	PlanQuantity                      *int32                                            `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                                            `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int64                                            `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                                            `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int64                                            `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                                            `json:"trial_end,omitempty"`
	BillingCycles                     *int32                                            `json:"billing_cycles,omitempty"`
	Addons                            []*SubscriptionCreateForCustomerAddon             `json:"addons,omitempty"`
	EventBasedAddons                  []*SubscriptionCreateForCustomerEventBasedAddon   `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove           []string                                          `json:"mandatory_addons_to_remove,omitempty"`
	StartDate                         *int64                                            `json:"start_date,omitempty"`
	AutoCollection                    SubscriptionAutoCollection                        `json:"auto_collection,omitempty"`
	TermsToCharge                     *int32                                            `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode              SubscriptionBillingAlignmentMode                  `json:"billing_alignment_mode,omitempty"`
	OfflinePaymentMethod              SubscriptionOfflinePaymentMethod                  `json:"offline_payment_method,omitempty"`
	PoNumber                          string                                            `json:"po_number,omitempty"`
	CouponIds                         []string                                          `json:"coupon_ids,omitempty"`
	PaymentSourceId                   string                                            `json:"payment_source_id,omitempty"`
	OverrideRelationship              *bool                                             `json:"override_relationship,omitempty"`
	ShippingAddress                   *SubscriptionCreateForCustomerShippingAddress     `json:"shipping_address,omitempty"`
	StatementDescriptor               *SubscriptionCreateForCustomerStatementDescriptor `json:"statement_descriptor,omitempty"`
	InvoiceNotes                      string                                            `json:"invoice_notes,omitempty"`
	InvoiceDate                       *int64                                            `json:"invoice_date,omitempty"`
	MetaData                          map[string]interface{}                            `json:"meta_data,omitempty"`
	InvoiceImmediately                *bool                                             `json:"invoice_immediately,omitempty"`
	ReplacePrimaryPaymentSource       *bool                                             `json:"replace_primary_payment_source,omitempty"`
	PaymentIntent                     *SubscriptionCreateForCustomerPaymentIntent       `json:"payment_intent,omitempty"`
	FreePeriod                        *int32                                            `json:"free_period,omitempty"`
	FreePeriodUnit                    SubscriptionFreePeriodUnit                        `json:"free_period_unit,omitempty"`
	ContractTerm                      *SubscriptionCreateForCustomerContractTerm        `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                            `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	TrialEndAction                    SubscriptionTrialEndAction                        `json:"trial_end_action,omitempty"`
	PaymentInitiator                  SubscriptionPaymentInitiator                      `json:"payment_initiator,omitempty"`
	Coupons                           []*SubscriptionCreateForCustomerCoupon            `json:"coupons,omitempty"`
	apiRequest                        `json:"-" form:"-"`
}

func (r *SubscriptionCreateForCustomerRequest) payload() any { return r }

// input sub resource params multi
type SubscriptionCreateForCustomerAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}

// input sub resource params multi
type SubscriptionCreateForCustomerEventBasedAddon struct {
	Id                  string                              `json:"id,omitempty"`
	Quantity            *int32                              `json:"quantity,omitempty"`
	UnitPrice           *int64                              `json:"unit_price,omitempty"`
	QuantityInDecimal   string                              `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string                              `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodInDays *int32                              `json:"service_period_in_days,omitempty"`
	OnEvent             SubscriptionEventBasedAddonOnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool                               `json:"charge_once,omitempty"`
	ChargeOn            SubscriptionEventBasedAddonChargeOn `json:"charge_on,omitempty"`
}

// input sub resource params single
type SubscriptionCreateForCustomerShippingAddress struct {
	FirstName        string                                      `json:"first_name,omitempty"`
	LastName         string                                      `json:"last_name,omitempty"`
	Email            string                                      `json:"email,omitempty"`
	Company          string                                      `json:"company,omitempty"`
	Phone            string                                      `json:"phone,omitempty"`
	Line1            string                                      `json:"line1,omitempty"`
	Line2            string                                      `json:"line2,omitempty"`
	Line3            string                                      `json:"line3,omitempty"`
	City             string                                      `json:"city,omitempty"`
	StateCode        string                                      `json:"state_code,omitempty"`
	State            string                                      `json:"state,omitempty"`
	Zip              string                                      `json:"zip,omitempty"`
	Country          string                                      `json:"country,omitempty"`
	ValidationStatus SubscriptionShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type SubscriptionCreateForCustomerStatementDescriptor struct {
	Descriptor string `json:"descriptor,omitempty"`
}

// input sub resource params single
type SubscriptionCreateForCustomerPaymentIntent struct {
	Id                    string                                     `json:"id,omitempty"`
	GatewayAccountId      string                                     `json:"gateway_account_id,omitempty"`
	GwToken               string                                     `json:"gw_token,omitempty"`
	PaymentMethodType     SubscriptionPaymentIntentPaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                                     `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                                     `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}                     `json:"additional_information,omitempty"`
}

// input sub resource params single
type SubscriptionCreateForCustomerContractTerm struct {
	ActionAtTermEnd          SubscriptionContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                                  `json:"cancellation_cutoff_period,omitempty"`
}

// input sub resource params multi
type SubscriptionCreateForCustomerCoupon struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}

type SubscriptionCreateWithItemsRequest struct {
	Id                                string                                          `json:"id,omitempty"`
	BusinessEntityId                  string                                          `json:"business_entity_id,omitempty"`
	TrialEnd                          *int64                                          `json:"trial_end,omitempty"`
	BillingCycles                     *int32                                          `json:"billing_cycles,omitempty"`
	SubscriptionItems                 []*SubscriptionCreateWithItemsSubscriptionItem  `json:"subscription_items,omitempty"`
	Discounts                         []*SubscriptionCreateWithItemsDiscount          `json:"discounts,omitempty"`
	MandatoryItemsToRemove            []string                                        `json:"mandatory_items_to_remove,omitempty"`
	ItemTiers                         []*SubscriptionCreateWithItemsItemTier          `json:"item_tiers,omitempty"`
	NetTermDays                       *int32                                          `json:"net_term_days,omitempty"`
	StartDate                         *int64                                          `json:"start_date,omitempty"`
	AutoCollection                    SubscriptionAutoCollection                      `json:"auto_collection,omitempty"`
	TermsToCharge                     *int32                                          `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode              SubscriptionBillingAlignmentMode                `json:"billing_alignment_mode,omitempty"`
	OfflinePaymentMethod              SubscriptionOfflinePaymentMethod                `json:"offline_payment_method,omitempty"`
	PoNumber                          string                                          `json:"po_number,omitempty"`
	CouponIds                         []string                                        `json:"coupon_ids,omitempty"`
	PaymentSourceId                   string                                          `json:"payment_source_id,omitempty"`
	OverrideRelationship              *bool                                           `json:"override_relationship,omitempty"`
	ShippingAddress                   *SubscriptionCreateWithItemsShippingAddress     `json:"shipping_address,omitempty"`
	StatementDescriptor               *SubscriptionCreateWithItemsStatementDescriptor `json:"statement_descriptor,omitempty"`
	InvoiceNotes                      string                                          `json:"invoice_notes,omitempty"`
	InvoiceDate                       *int64                                          `json:"invoice_date,omitempty"`
	MetaData                          map[string]interface{}                          `json:"meta_data,omitempty"`
	InvoiceImmediately                *bool                                           `json:"invoice_immediately,omitempty"`
	ReplacePrimaryPaymentSource       *bool                                           `json:"replace_primary_payment_source,omitempty"`
	PaymentIntent                     *SubscriptionCreateWithItemsPaymentIntent       `json:"payment_intent,omitempty"`
	FreePeriod                        *int32                                          `json:"free_period,omitempty"`
	FreePeriodUnit                    SubscriptionFreePeriodUnit                      `json:"free_period_unit,omitempty"`
	ContractTerm                      *SubscriptionCreateWithItemsContractTerm        `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                          `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	CreatePendingInvoices             *bool                                           `json:"create_pending_invoices,omitempty"`
	AutoCloseInvoices                 *bool                                           `json:"auto_close_invoices,omitempty"`
	FirstInvoicePending               *bool                                           `json:"first_invoice_pending,omitempty"`
	TrialEndAction                    SubscriptionTrialEndAction                      `json:"trial_end_action,omitempty"`
	PaymentInitiator                  SubscriptionPaymentInitiator                    `json:"payment_initiator,omitempty"`
	Coupons                           []*SubscriptionCreateWithItemsCoupon            `json:"coupons,omitempty"`
	BillingOverride                   *SubscriptionCreateWithItemsBillingOverride     `json:"billing_override,omitempty"`
	apiRequest                        `json:"-" form:"-"`
}

func (r *SubscriptionCreateWithItemsRequest) payload() any { return r }

// input sub resource params multi
type SubscriptionCreateWithItemsSubscriptionItem struct {
	ItemPriceId                     string                                                      `json:"item_price_id"`
	Quantity                        *int32                                                      `json:"quantity,omitempty"`
	QuantityInDecimal               string                                                      `json:"quantity_in_decimal,omitempty"`
	UnitPrice                       *int64                                                      `json:"unit_price,omitempty"`
	UnitPriceInDecimal              string                                                      `json:"unit_price_in_decimal,omitempty"`
	BillingCycles                   *int32                                                      `json:"billing_cycles,omitempty"`
	TrialEnd                        *int64                                                      `json:"trial_end,omitempty"`
	ServicePeriodDays               *int32                                                      `json:"service_period_days,omitempty"`
	ChargeOnEvent                   SubscriptionSubscriptionItemChargeOnEvent                   `json:"charge_on_event,omitempty"`
	ChargeOnce                      *bool                                                       `json:"charge_once,omitempty"`
	ItemType                        SubscriptionSubscriptionItemItemType                        `json:"item_type,omitempty"`
	ChargeOnOption                  SubscriptionSubscriptionItemChargeOnOption                  `json:"charge_on_option,omitempty"`
	UsageAccumulationResetFrequency SubscriptionSubscriptionItemUsageAccumulationResetFrequency `json:"usage_accumulation_reset_frequency,omitempty"`
}

// input sub resource params multi
type SubscriptionCreateWithItemsDiscount struct {
	ApplyOn       SubscriptionDiscountApplyOn      `json:"apply_on,omitempty"`
	DurationType  SubscriptionDiscountDurationType `json:"duration_type"`
	Percentage    *float64                         `json:"percentage,omitempty"`
	Amount        *int64                           `json:"amount,omitempty"`
	Period        *int32                           `json:"period,omitempty"`
	PeriodUnit    SubscriptionDiscountPeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool                            `json:"included_in_mrr,omitempty"`
	ItemPriceId   string                           `json:"item_price_id,omitempty"`
	Quantity      *int32                           `json:"quantity,omitempty"`
}

// input sub resource params multi
type SubscriptionCreateWithItemsItemTier struct {
	ItemPriceId           string                          `json:"item_price_id,omitempty"`
	StartingUnit          *int32                          `json:"starting_unit,omitempty"`
	EndingUnit            *int32                          `json:"ending_unit,omitempty"`
	Price                 *int64                          `json:"price,omitempty"`
	StartingUnitInDecimal string                          `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string                          `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string                          `json:"price_in_decimal,omitempty"`
	PricingType           SubscriptionItemTierPricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32                          `json:"package_size,omitempty"`
}

// input sub resource params single
type SubscriptionCreateWithItemsShippingAddress struct {
	FirstName        string                                      `json:"first_name,omitempty"`
	LastName         string                                      `json:"last_name,omitempty"`
	Email            string                                      `json:"email,omitempty"`
	Company          string                                      `json:"company,omitempty"`
	Phone            string                                      `json:"phone,omitempty"`
	Line1            string                                      `json:"line1,omitempty"`
	Line2            string                                      `json:"line2,omitempty"`
	Line3            string                                      `json:"line3,omitempty"`
	City             string                                      `json:"city,omitempty"`
	StateCode        string                                      `json:"state_code,omitempty"`
	State            string                                      `json:"state,omitempty"`
	Zip              string                                      `json:"zip,omitempty"`
	Country          string                                      `json:"country,omitempty"`
	ValidationStatus SubscriptionShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type SubscriptionCreateWithItemsStatementDescriptor struct {
	Descriptor string `json:"descriptor,omitempty"`
}

// input sub resource params single
type SubscriptionCreateWithItemsPaymentIntent struct {
	Id                    string                                     `json:"id,omitempty"`
	GatewayAccountId      string                                     `json:"gateway_account_id,omitempty"`
	GwToken               string                                     `json:"gw_token,omitempty"`
	PaymentMethodType     SubscriptionPaymentIntentPaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                                     `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                                     `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}                     `json:"additional_information,omitempty"`
}

// input sub resource params single
type SubscriptionCreateWithItemsContractTerm struct {
	ActionAtTermEnd          SubscriptionContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	ContractStart            *int64                                  `json:"contract_start,omitempty"`
	CancellationCutoffPeriod *int32                                  `json:"cancellation_cutoff_period,omitempty"`
}

// input sub resource params multi
type SubscriptionCreateWithItemsCoupon struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}

// input sub resource params single
type SubscriptionCreateWithItemsBillingOverride struct {
	MaxExcessPaymentUsage     *int64 `json:"max_excess_payment_usage,omitempty"`
	MaxRefundableCreditsUsage *int64 `json:"max_refundable_credits_usage,omitempty"`
}

type SubscriptionListRequest struct {
	Limit                  *int32           `json:"limit,omitempty"`
	Offset                 string           `json:"offset,omitempty"`
	IncludeDeleted         *bool            `json:"include_deleted,omitempty"`
	Id                     *StringFilter    `json:"id,omitempty"`
	CustomerId             *StringFilter    `json:"customer_id,omitempty"`
	PlanId                 *StringFilter    `json:"plan_id,omitempty"`
	ItemId                 *StringFilter    `json:"item_id,omitempty"`
	ItemPriceId            *StringFilter    `json:"item_price_id,omitempty"`
	Status                 *EnumFilter      `json:"status,omitempty"`
	CancelReason           *EnumFilter      `json:"cancel_reason,omitempty"`
	CancelReasonCode       *StringFilter    `json:"cancel_reason_code,omitempty"`
	RemainingBillingCycles *NumberFilter    `json:"remaining_billing_cycles,omitempty"`
	CreatedAt              *TimestampFilter `json:"created_at,omitempty"`
	ActivatedAt            *TimestampFilter `json:"activated_at,omitempty"`
	NextBillingAt          *TimestampFilter `json:"next_billing_at,omitempty"`
	CancelledAt            *TimestampFilter `json:"cancelled_at,omitempty"`
	HasScheduledChanges    *BooleanFilter   `json:"has_scheduled_changes,omitempty"`
	UpdatedAt              *TimestampFilter `json:"updated_at,omitempty"`
	OfflinePaymentMethod   *EnumFilter      `json:"offline_payment_method,omitempty"`
	AutoCloseInvoices      *BooleanFilter   `json:"auto_close_invoices,omitempty"`
	OverrideRelationship   *BooleanFilter   `json:"override_relationship,omitempty"`
	SortBy                 *SortFilter      `json:"sort_by,omitempty"`
	BusinessEntityId       *StringFilter    `json:"business_entity_id,omitempty"`
	Channel                *EnumFilter      `json:"channel,omitempty"`
	apiRequest             `json:"-" form:"-"`
}

func (r *SubscriptionListRequest) payload() any { return r }

type SubscriptionSubscriptionsForCustomerRequest struct {
	Limit      *int32 `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *SubscriptionSubscriptionsForCustomerRequest) payload() any { return r }

type SubscriptionContractTermsForSubscriptionRequest struct {
	Limit      *int32      `json:"limit,omitempty"`
	Offset     string      `json:"offset,omitempty"`
	SortBy     *SortFilter `json:"sort_by,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *SubscriptionContractTermsForSubscriptionRequest) payload() any { return r }

type SubscriptionListDiscountsRequest struct {
	Limit      *int32 `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *SubscriptionListDiscountsRequest) payload() any { return r }

type SubscriptionRemoveScheduledCancellationRequest struct {
	BillingCycles                     *int32                                               `json:"billing_cycles,omitempty"`
	ContractTerm                      *SubscriptionRemoveScheduledCancellationContractTerm `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                               `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	apiRequest                        `json:"-" form:"-"`
}

func (r *SubscriptionRemoveScheduledCancellationRequest) payload() any { return r }

// input sub resource params single
type SubscriptionRemoveScheduledCancellationContractTerm struct {
	ActionAtTermEnd          SubscriptionContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                                  `json:"cancellation_cutoff_period,omitempty"`
}

type SubscriptionRemoveCouponsRequest struct {
	CouponIds  []string `json:"coupon_ids,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *SubscriptionRemoveCouponsRequest) payload() any { return r }

type SubscriptionUpdateRequest struct {
	PlanId                            string                                 `json:"plan_id,omitempty"`
	PlanQuantity                      *int32                                 `json:"plan_quantity,omitempty"`
	PlanUnitPrice                     *int64                                 `json:"plan_unit_price,omitempty"`
	SetupFee                          *int64                                 `json:"setup_fee,omitempty"`
	Addons                            []*SubscriptionUpdateAddon             `json:"addons,omitempty"`
	EventBasedAddons                  []*SubscriptionUpdateEventBasedAddon   `json:"event_based_addons,omitempty"`
	ReplaceAddonList                  *bool                                  `json:"replace_addon_list,omitempty"`
	MandatoryAddonsToRemove           []string                               `json:"mandatory_addons_to_remove,omitempty"`
	PlanQuantityInDecimal             string                                 `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPriceInDecimal            string                                 `json:"plan_unit_price_in_decimal,omitempty"`
	InvoiceDate                       *int64                                 `json:"invoice_date,omitempty"`
	StartDate                         *int64                                 `json:"start_date,omitempty"`
	TrialEnd                          *int64                                 `json:"trial_end,omitempty"`
	BillingCycles                     *int32                                 `json:"billing_cycles,omitempty"`
	TermsToCharge                     *int32                                 `json:"terms_to_charge,omitempty"`
	ReactivateFrom                    *int64                                 `json:"reactivate_from,omitempty"`
	BillingAlignmentMode              SubscriptionBillingAlignmentMode       `json:"billing_alignment_mode,omitempty"`
	AutoCollection                    SubscriptionAutoCollection             `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              SubscriptionOfflinePaymentMethod       `json:"offline_payment_method,omitempty"`
	PoNumber                          string                                 `json:"po_number,omitempty"`
	CouponIds                         []string                               `json:"coupon_ids,omitempty"`
	ReplaceCouponList                 *bool                                  `json:"replace_coupon_list,omitempty"`
	Prorate                           *bool                                  `json:"prorate,omitempty"`
	EndOfTerm                         *bool                                  `json:"end_of_term,omitempty"`
	ForceTermReset                    *bool                                  `json:"force_term_reset,omitempty"`
	Reactivate                        *bool                                  `json:"reactivate,omitempty"`
	Card                              *SubscriptionUpdateCard                `json:"card,omitempty"`
	TokenId                           string                                 `json:"token_id,omitempty"`
	PaymentMethod                     *SubscriptionUpdatePaymentMethod       `json:"payment_method,omitempty"`
	PaymentIntent                     *SubscriptionUpdatePaymentIntent       `json:"payment_intent,omitempty"`
	BillingAddress                    *SubscriptionUpdateBillingAddress      `json:"billing_address,omitempty"`
	ShippingAddress                   *SubscriptionUpdateShippingAddress     `json:"shipping_address,omitempty"`
	StatementDescriptor               *SubscriptionUpdateStatementDescriptor `json:"statement_descriptor,omitempty"`
	Customer                          *SubscriptionUpdateCustomer            `json:"customer,omitempty"`
	InvoiceNotes                      string                                 `json:"invoice_notes,omitempty"`
	MetaData                          map[string]interface{}                 `json:"meta_data,omitempty"`
	InvoiceImmediately                *bool                                  `json:"invoice_immediately,omitempty"`
	OverrideRelationship              *bool                                  `json:"override_relationship,omitempty"`
	ChangesScheduledAt                *int64                                 `json:"changes_scheduled_at,omitempty"`
	ChangeOption                      SubscriptionChangeOption               `json:"change_option,omitempty"`
	ContractTerm                      *SubscriptionUpdateContractTerm        `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                 `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	FreePeriod                        *int32                                 `json:"free_period,omitempty"`
	FreePeriodUnit                    SubscriptionFreePeriodUnit             `json:"free_period_unit,omitempty"`
	TrialEndAction                    SubscriptionTrialEndAction             `json:"trial_end_action,omitempty"`
	Coupons                           []*SubscriptionUpdateCoupon            `json:"coupons,omitempty"`
	apiRequest                        `json:"-" form:"-"`
}

func (r *SubscriptionUpdateRequest) payload() any { return r }

// input sub resource params multi
type SubscriptionUpdateAddon struct {
	Id                 string                         `json:"id,omitempty"`
	Quantity           *int32                         `json:"quantity,omitempty"`
	UnitPrice          *int64                         `json:"unit_price,omitempty"`
	BillingCycles      *int32                         `json:"billing_cycles,omitempty"`
	QuantityInDecimal  string                         `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string                         `json:"unit_price_in_decimal,omitempty"`
	TrialEnd           *int64                         `json:"trial_end,omitempty"`
	ProrationType      SubscriptionAddonProrationType `json:"proration_type,omitempty"`
}

// input sub resource params multi
type SubscriptionUpdateEventBasedAddon struct {
	Id                  string                              `json:"id,omitempty"`
	Quantity            *int32                              `json:"quantity,omitempty"`
	UnitPrice           *int64                              `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32                              `json:"service_period_in_days,omitempty"`
	ChargeOn            SubscriptionEventBasedAddonChargeOn `json:"charge_on,omitempty"`
	OnEvent             SubscriptionEventBasedAddonOnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool                               `json:"charge_once,omitempty"`
	QuantityInDecimal   string                              `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string                              `json:"unit_price_in_decimal,omitempty"`
}

// input sub resource params single
type SubscriptionUpdateCard struct {
	Gateway               SubscriptionCardGateway         `json:"gateway,omitempty"`
	GatewayAccountId      string                          `json:"gateway_account_id,omitempty"`
	TmpToken              string                          `json:"tmp_token,omitempty"`
	FirstName             string                          `json:"first_name,omitempty"`
	LastName              string                          `json:"last_name,omitempty"`
	Number                string                          `json:"number,omitempty"`
	ExpiryMonth           *int32                          `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                          `json:"expiry_year,omitempty"`
	Cvv                   string                          `json:"cvv,omitempty"`
	PreferredScheme       SubscriptionCardPreferredScheme `json:"preferred_scheme,omitempty"`
	BillingAddr1          string                          `json:"billing_addr1,omitempty"`
	BillingAddr2          string                          `json:"billing_addr2,omitempty"`
	BillingCity           string                          `json:"billing_city,omitempty"`
	BillingStateCode      string                          `json:"billing_state_code,omitempty"`
	BillingState          string                          `json:"billing_state,omitempty"`
	BillingZip            string                          `json:"billing_zip,omitempty"`
	BillingCountry        string                          `json:"billing_country,omitempty"`
	IpAddress             string                          `json:"ip_address,omitempty"`
	AdditionalInformation map[string]interface{}          `json:"additional_information,omitempty"`
}

// input sub resource params single
type SubscriptionUpdatePaymentMethod struct {
	Type                  SubscriptionPaymentMethodType    `json:"type,omitempty"`
	Gateway               SubscriptionPaymentMethodGateway `json:"gateway,omitempty"`
	GatewayAccountId      string                           `json:"gateway_account_id,omitempty"`
	ReferenceId           string                           `json:"reference_id,omitempty"`
	TmpToken              string                           `json:"tmp_token,omitempty"`
	IssuingCountry        string                           `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{}           `json:"additional_information,omitempty"`
}

// input sub resource params single
type SubscriptionUpdatePaymentIntent struct {
	Id                    string                                     `json:"id,omitempty"`
	GatewayAccountId      string                                     `json:"gateway_account_id,omitempty"`
	GwToken               string                                     `json:"gw_token,omitempty"`
	PaymentMethodType     SubscriptionPaymentIntentPaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                                     `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                                     `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}                     `json:"additional_information,omitempty"`
}

// input sub resource params single
type SubscriptionUpdateBillingAddress struct {
	FirstName        string                                     `json:"first_name,omitempty"`
	LastName         string                                     `json:"last_name,omitempty"`
	Email            string                                     `json:"email,omitempty"`
	Company          string                                     `json:"company,omitempty"`
	Phone            string                                     `json:"phone,omitempty"`
	Line1            string                                     `json:"line1,omitempty"`
	Line2            string                                     `json:"line2,omitempty"`
	Line3            string                                     `json:"line3,omitempty"`
	City             string                                     `json:"city,omitempty"`
	StateCode        string                                     `json:"state_code,omitempty"`
	State            string                                     `json:"state,omitempty"`
	Zip              string                                     `json:"zip,omitempty"`
	Country          string                                     `json:"country,omitempty"`
	ValidationStatus SubscriptionBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type SubscriptionUpdateShippingAddress struct {
	FirstName        string                                      `json:"first_name,omitempty"`
	LastName         string                                      `json:"last_name,omitempty"`
	Email            string                                      `json:"email,omitempty"`
	Company          string                                      `json:"company,omitempty"`
	Phone            string                                      `json:"phone,omitempty"`
	Line1            string                                      `json:"line1,omitempty"`
	Line2            string                                      `json:"line2,omitempty"`
	Line3            string                                      `json:"line3,omitempty"`
	City             string                                      `json:"city,omitempty"`
	StateCode        string                                      `json:"state_code,omitempty"`
	State            string                                      `json:"state,omitempty"`
	Zip              string                                      `json:"zip,omitempty"`
	Country          string                                      `json:"country,omitempty"`
	ValidationStatus SubscriptionShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type SubscriptionUpdateStatementDescriptor struct {
	Descriptor string `json:"descriptor,omitempty"`
}

// input sub resource params single
type SubscriptionUpdateCustomer struct {
	VatNumber                        string                               `json:"vat_number,omitempty"`
	VatNumberPrefix                  string                               `json:"vat_number_prefix,omitempty"`
	EntityIdentifierScheme           string                               `json:"entity_identifier_scheme,omitempty"`
	IsEinvoiceEnabled                *bool                                `json:"is_einvoice_enabled,omitempty"`
	EinvoicingMethod                 SubscriptionCustomerEinvoicingMethod `json:"einvoicing_method,omitempty"`
	EntityIdentifierStandard         string                               `json:"entity_identifier_standard,omitempty"`
	BusinessCustomerWithoutVatNumber *bool                                `json:"business_customer_without_vat_number,omitempty"`
	RegisteredForGst                 *bool                                `json:"registered_for_gst,omitempty"`
}

// input sub resource params single
type SubscriptionUpdateContractTerm struct {
	ActionAtTermEnd          SubscriptionContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                                  `json:"cancellation_cutoff_period,omitempty"`
}

// input sub resource params multi
type SubscriptionUpdateCoupon struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}

type SubscriptionUpdateForItemsRequest struct {
	SubscriptionItems                 []*SubscriptionUpdateForItemsSubscriptionItem  `json:"subscription_items,omitempty"`
	MandatoryItemsToRemove            []string                                       `json:"mandatory_items_to_remove,omitempty"`
	ReplaceItemsList                  *bool                                          `json:"replace_items_list,omitempty"`
	Discounts                         []*SubscriptionUpdateForItemsDiscount          `json:"discounts,omitempty"`
	ItemTiers                         []*SubscriptionUpdateForItemsItemTier          `json:"item_tiers,omitempty"`
	NetTermDays                       *int32                                         `json:"net_term_days,omitempty"`
	InvoiceDate                       *int64                                         `json:"invoice_date,omitempty"`
	StartDate                         *int64                                         `json:"start_date,omitempty"`
	TrialEnd                          *int64                                         `json:"trial_end,omitempty"`
	BillingCycles                     *int32                                         `json:"billing_cycles,omitempty"`
	TermsToCharge                     *int32                                         `json:"terms_to_charge,omitempty"`
	ReactivateFrom                    *int64                                         `json:"reactivate_from,omitempty"`
	BillingAlignmentMode              SubscriptionBillingAlignmentMode               `json:"billing_alignment_mode,omitempty"`
	AutoCollection                    SubscriptionAutoCollection                     `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              SubscriptionOfflinePaymentMethod               `json:"offline_payment_method,omitempty"`
	PoNumber                          string                                         `json:"po_number,omitempty"`
	CouponIds                         []string                                       `json:"coupon_ids,omitempty"`
	ReplaceCouponList                 *bool                                          `json:"replace_coupon_list,omitempty"`
	Prorate                           *bool                                          `json:"prorate,omitempty"`
	EndOfTerm                         *bool                                          `json:"end_of_term,omitempty"`
	ForceTermReset                    *bool                                          `json:"force_term_reset,omitempty"`
	Reactivate                        *bool                                          `json:"reactivate,omitempty"`
	Card                              *SubscriptionUpdateForItemsCard                `json:"card,omitempty"`
	TokenId                           string                                         `json:"token_id,omitempty"`
	PaymentMethod                     *SubscriptionUpdateForItemsPaymentMethod       `json:"payment_method,omitempty"`
	PaymentIntent                     *SubscriptionUpdateForItemsPaymentIntent       `json:"payment_intent,omitempty"`
	BillingAddress                    *SubscriptionUpdateForItemsBillingAddress      `json:"billing_address,omitempty"`
	ShippingAddress                   *SubscriptionUpdateForItemsShippingAddress     `json:"shipping_address,omitempty"`
	StatementDescriptor               *SubscriptionUpdateForItemsStatementDescriptor `json:"statement_descriptor,omitempty"`
	Customer                          *SubscriptionUpdateForItemsCustomer            `json:"customer,omitempty"`
	InvoiceNotes                      string                                         `json:"invoice_notes,omitempty"`
	MetaData                          map[string]interface{}                         `json:"meta_data,omitempty"`
	InvoiceImmediately                *bool                                          `json:"invoice_immediately,omitempty"`
	OverrideRelationship              *bool                                          `json:"override_relationship,omitempty"`
	ChangesScheduledAt                *int64                                         `json:"changes_scheduled_at,omitempty"`
	ChangeOption                      SubscriptionChangeOption                       `json:"change_option,omitempty"`
	ContractTerm                      *SubscriptionUpdateForItemsContractTerm        `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                         `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	FreePeriod                        *int32                                         `json:"free_period,omitempty"`
	FreePeriodUnit                    SubscriptionFreePeriodUnit                     `json:"free_period_unit,omitempty"`
	CreatePendingInvoices             *bool                                          `json:"create_pending_invoices,omitempty"`
	AutoCloseInvoices                 *bool                                          `json:"auto_close_invoices,omitempty"`
	TrialEndAction                    SubscriptionTrialEndAction                     `json:"trial_end_action,omitempty"`
	PaymentInitiator                  SubscriptionPaymentInitiator                   `json:"payment_initiator,omitempty"`
	Coupons                           []*SubscriptionUpdateForItemsCoupon            `json:"coupons,omitempty"`
	InvoiceUsages                     *bool                                          `json:"invoice_usages,omitempty"`
	BillingOverride                   *SubscriptionUpdateForItemsBillingOverride     `json:"billing_override,omitempty"`
	apiRequest                        `json:"-" form:"-"`
}

func (r *SubscriptionUpdateForItemsRequest) payload() any { return r }

// input sub resource params multi
type SubscriptionUpdateForItemsSubscriptionItem struct {
	ItemPriceId                     string                                                      `json:"item_price_id"`
	Quantity                        *int32                                                      `json:"quantity,omitempty"`
	QuantityInDecimal               string                                                      `json:"quantity_in_decimal,omitempty"`
	UnitPrice                       *int64                                                      `json:"unit_price,omitempty"`
	UnitPriceInDecimal              string                                                      `json:"unit_price_in_decimal,omitempty"`
	BillingCycles                   *int32                                                      `json:"billing_cycles,omitempty"`
	TrialEnd                        *int64                                                      `json:"trial_end,omitempty"`
	ServicePeriodDays               *int32                                                      `json:"service_period_days,omitempty"`
	ChargeOnEvent                   SubscriptionSubscriptionItemChargeOnEvent                   `json:"charge_on_event,omitempty"`
	ChargeOnce                      *bool                                                       `json:"charge_once,omitempty"`
	ChargeOnOption                  SubscriptionSubscriptionItemChargeOnOption                  `json:"charge_on_option,omitempty"`
	ItemType                        SubscriptionSubscriptionItemItemType                        `json:"item_type,omitempty"`
	ProrationType                   SubscriptionSubscriptionItemProrationType                   `json:"proration_type,omitempty"`
	UsageAccumulationResetFrequency SubscriptionSubscriptionItemUsageAccumulationResetFrequency `json:"usage_accumulation_reset_frequency,omitempty"`
}

// input sub resource params multi
type SubscriptionUpdateForItemsDiscount struct {
	ApplyOn       SubscriptionDiscountApplyOn       `json:"apply_on,omitempty"`
	DurationType  SubscriptionDiscountDurationType  `json:"duration_type"`
	Percentage    *float64                          `json:"percentage,omitempty"`
	Amount        *int64                            `json:"amount,omitempty"`
	Period        *int32                            `json:"period,omitempty"`
	PeriodUnit    SubscriptionDiscountPeriodUnit    `json:"period_unit,omitempty"`
	IncludedInMrr *bool                             `json:"included_in_mrr,omitempty"`
	ItemPriceId   string                            `json:"item_price_id,omitempty"`
	Quantity      *int32                            `json:"quantity,omitempty"`
	OperationType SubscriptionDiscountOperationType `json:"operation_type"`
	Id            string                            `json:"id,omitempty"`
}

// input sub resource params multi
type SubscriptionUpdateForItemsItemTier struct {
	ItemPriceId           string                          `json:"item_price_id,omitempty"`
	StartingUnit          *int32                          `json:"starting_unit,omitempty"`
	EndingUnit            *int32                          `json:"ending_unit,omitempty"`
	Price                 *int64                          `json:"price,omitempty"`
	StartingUnitInDecimal string                          `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string                          `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string                          `json:"price_in_decimal,omitempty"`
	PricingType           SubscriptionItemTierPricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32                          `json:"package_size,omitempty"`
}

// input sub resource params single
type SubscriptionUpdateForItemsCard struct {
	Gateway               SubscriptionCardGateway         `json:"gateway,omitempty"`
	GatewayAccountId      string                          `json:"gateway_account_id,omitempty"`
	TmpToken              string                          `json:"tmp_token,omitempty"`
	FirstName             string                          `json:"first_name,omitempty"`
	LastName              string                          `json:"last_name,omitempty"`
	Number                string                          `json:"number,omitempty"`
	ExpiryMonth           *int32                          `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                          `json:"expiry_year,omitempty"`
	Cvv                   string                          `json:"cvv,omitempty"`
	PreferredScheme       SubscriptionCardPreferredScheme `json:"preferred_scheme,omitempty"`
	BillingAddr1          string                          `json:"billing_addr1,omitempty"`
	BillingAddr2          string                          `json:"billing_addr2,omitempty"`
	BillingCity           string                          `json:"billing_city,omitempty"`
	BillingStateCode      string                          `json:"billing_state_code,omitempty"`
	BillingState          string                          `json:"billing_state,omitempty"`
	BillingZip            string                          `json:"billing_zip,omitempty"`
	BillingCountry        string                          `json:"billing_country,omitempty"`
	IpAddress             string                          `json:"ip_address,omitempty"`
	AdditionalInformation map[string]interface{}          `json:"additional_information,omitempty"`
}

// input sub resource params single
type SubscriptionUpdateForItemsPaymentMethod struct {
	Type                  SubscriptionPaymentMethodType    `json:"type,omitempty"`
	Gateway               SubscriptionPaymentMethodGateway `json:"gateway,omitempty"`
	GatewayAccountId      string                           `json:"gateway_account_id,omitempty"`
	ReferenceId           string                           `json:"reference_id,omitempty"`
	TmpToken              string                           `json:"tmp_token,omitempty"`
	IssuingCountry        string                           `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{}           `json:"additional_information,omitempty"`
}

// input sub resource params single
type SubscriptionUpdateForItemsPaymentIntent struct {
	Id                    string                                     `json:"id,omitempty"`
	GatewayAccountId      string                                     `json:"gateway_account_id,omitempty"`
	GwToken               string                                     `json:"gw_token,omitempty"`
	PaymentMethodType     SubscriptionPaymentIntentPaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                                     `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                                     `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}                     `json:"additional_information,omitempty"`
}

// input sub resource params single
type SubscriptionUpdateForItemsBillingAddress struct {
	FirstName        string                                     `json:"first_name,omitempty"`
	LastName         string                                     `json:"last_name,omitempty"`
	Email            string                                     `json:"email,omitempty"`
	Company          string                                     `json:"company,omitempty"`
	Phone            string                                     `json:"phone,omitempty"`
	Line1            string                                     `json:"line1,omitempty"`
	Line2            string                                     `json:"line2,omitempty"`
	Line3            string                                     `json:"line3,omitempty"`
	City             string                                     `json:"city,omitempty"`
	StateCode        string                                     `json:"state_code,omitempty"`
	State            string                                     `json:"state,omitempty"`
	Zip              string                                     `json:"zip,omitempty"`
	Country          string                                     `json:"country,omitempty"`
	ValidationStatus SubscriptionBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type SubscriptionUpdateForItemsShippingAddress struct {
	FirstName        string                                      `json:"first_name,omitempty"`
	LastName         string                                      `json:"last_name,omitempty"`
	Email            string                                      `json:"email,omitempty"`
	Company          string                                      `json:"company,omitempty"`
	Phone            string                                      `json:"phone,omitempty"`
	Line1            string                                      `json:"line1,omitempty"`
	Line2            string                                      `json:"line2,omitempty"`
	Line3            string                                      `json:"line3,omitempty"`
	City             string                                      `json:"city,omitempty"`
	StateCode        string                                      `json:"state_code,omitempty"`
	State            string                                      `json:"state,omitempty"`
	Zip              string                                      `json:"zip,omitempty"`
	Country          string                                      `json:"country,omitempty"`
	ValidationStatus SubscriptionShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type SubscriptionUpdateForItemsStatementDescriptor struct {
	Descriptor string `json:"descriptor,omitempty"`
}

// input sub resource params single
type SubscriptionUpdateForItemsCustomer struct {
	VatNumber                        string                               `json:"vat_number,omitempty"`
	VatNumberPrefix                  string                               `json:"vat_number_prefix,omitempty"`
	EntityIdentifierScheme           string                               `json:"entity_identifier_scheme,omitempty"`
	IsEinvoiceEnabled                *bool                                `json:"is_einvoice_enabled,omitempty"`
	EinvoicingMethod                 SubscriptionCustomerEinvoicingMethod `json:"einvoicing_method,omitempty"`
	EntityIdentifierStandard         string                               `json:"entity_identifier_standard,omitempty"`
	BusinessCustomerWithoutVatNumber *bool                                `json:"business_customer_without_vat_number,omitempty"`
	RegisteredForGst                 *bool                                `json:"registered_for_gst,omitempty"`
}

// input sub resource params single
type SubscriptionUpdateForItemsContractTerm struct {
	ActionAtTermEnd          SubscriptionContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                                  `json:"cancellation_cutoff_period,omitempty"`
	ContractStart            *int64                                  `json:"contract_start,omitempty"`
}

// input sub resource params multi
type SubscriptionUpdateForItemsCoupon struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}

// input sub resource params single
type SubscriptionUpdateForItemsBillingOverride struct {
	MaxExcessPaymentUsage     *int64 `json:"max_excess_payment_usage,omitempty"`
	MaxRefundableCreditsUsage *int64 `json:"max_refundable_credits_usage,omitempty"`
}

type SubscriptionChangeTermEndRequest struct {
	TermEndsAt         *int64 `json:"term_ends_at"`
	Prorate            *bool  `json:"prorate,omitempty"`
	InvoiceImmediately *bool  `json:"invoice_immediately,omitempty"`
	apiRequest         `json:"-" form:"-"`
}

func (r *SubscriptionChangeTermEndRequest) payload() any { return r }

type SubscriptionReactivateRequest struct {
	TrialEnd                          *int64                                     `json:"trial_end,omitempty"`
	BillingCycles                     *int32                                     `json:"billing_cycles,omitempty"`
	ReactivateFrom                    *int64                                     `json:"reactivate_from,omitempty"`
	InvoiceImmediately                *bool                                      `json:"invoice_immediately,omitempty"`
	BillingAlignmentMode              SubscriptionBillingAlignmentMode           `json:"billing_alignment_mode,omitempty"`
	TermsToCharge                     *int32                                     `json:"terms_to_charge,omitempty"`
	InvoiceDate                       *int64                                     `json:"invoice_date,omitempty"`
	ContractTerm                      *SubscriptionReactivateContractTerm        `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                     `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	PaymentInitiator                  SubscriptionPaymentInitiator               `json:"payment_initiator,omitempty"`
	StatementDescriptor               *SubscriptionReactivateStatementDescriptor `json:"statement_descriptor,omitempty"`
	PaymentIntent                     *SubscriptionReactivatePaymentIntent       `json:"payment_intent,omitempty"`
	apiRequest                        `json:"-" form:"-"`
}

func (r *SubscriptionReactivateRequest) payload() any { return r }

// input sub resource params single
type SubscriptionReactivateContractTerm struct {
	ActionAtTermEnd          SubscriptionContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                                  `json:"cancellation_cutoff_period,omitempty"`
}

// input sub resource params single
type SubscriptionReactivateStatementDescriptor struct {
	Descriptor string `json:"descriptor,omitempty"`
}

// input sub resource params single
type SubscriptionReactivatePaymentIntent struct {
	Id                    string                                     `json:"id,omitempty"`
	GatewayAccountId      string                                     `json:"gateway_account_id,omitempty"`
	GwToken               string                                     `json:"gw_token,omitempty"`
	PaymentMethodType     SubscriptionPaymentIntentPaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                                     `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                                     `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}                     `json:"additional_information,omitempty"`
}

type SubscriptionAddChargeAtTermEndRequest struct {
	Amount                 *int64                      `json:"amount,omitempty"`
	Description            string                      `json:"description"`
	AmountInDecimal        string                      `json:"amount_in_decimal,omitempty"`
	AvalaraSaleType        SubscriptionAvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32                      `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32                      `json:"avalara_service_type,omitempty"`
	DateFrom               *int64                      `json:"date_from,omitempty"`
	DateTo                 *int64                      `json:"date_to,omitempty"`
	apiRequest             `json:"-" form:"-"`
}

func (r *SubscriptionAddChargeAtTermEndRequest) payload() any { return r }

type SubscriptionChargeAddonAtTermEndRequest struct {
	AddonId                 string `json:"addon_id"`
	AddonQuantity           *int32 `json:"addon_quantity,omitempty"`
	AddonUnitPrice          *int64 `json:"addon_unit_price,omitempty"`
	AddonQuantityInDecimal  string `json:"addon_quantity_in_decimal,omitempty"`
	AddonUnitPriceInDecimal string `json:"addon_unit_price_in_decimal,omitempty"`
	DateFrom                *int64 `json:"date_from,omitempty"`
	DateTo                  *int64 `json:"date_to,omitempty"`
	apiRequest              `json:"-" form:"-"`
}

func (r *SubscriptionChargeAddonAtTermEndRequest) payload() any { return r }

type SubscriptionChargeFutureRenewalsRequest struct {
	TermsToCharge         *int32                                                   `json:"terms_to_charge,omitempty"`
	SpecificDatesSchedule []*SubscriptionChargeFutureRenewalsSpecificDatesSchedule `json:"specific_dates_schedule,omitempty"`
	FixedIntervalSchedule *SubscriptionChargeFutureRenewalsFixedIntervalSchedule   `json:"fixed_interval_schedule,omitempty"`
	InvoiceImmediately    *bool                                                    `json:"invoice_immediately,omitempty"`
	ScheduleType          SubscriptionScheduleType                                 `json:"schedule_type,omitempty"`
	apiRequest            `json:"-" form:"-"`
}

func (r *SubscriptionChargeFutureRenewalsRequest) payload() any { return r }

// input sub resource params multi
type SubscriptionChargeFutureRenewalsSpecificDatesSchedule struct {
	TermsToCharge *int32 `json:"terms_to_charge,omitempty"`
	Date          *int64 `json:"date,omitempty"`
}

// input sub resource params single
type SubscriptionChargeFutureRenewalsFixedIntervalSchedule struct {
	NumberOfOccurrences *int32                                         `json:"number_of_occurrences,omitempty"`
	DaysBeforeRenewal   *int32                                         `json:"days_before_renewal,omitempty"`
	EndScheduleOn       SubscriptionFixedIntervalScheduleEndScheduleOn `json:"end_schedule_on,omitempty"`
	EndDate             *int64                                         `json:"end_date,omitempty"`
}

type SubscriptionEditAdvanceInvoiceScheduleRequest struct {
	TermsToCharge         *int32                                                         `json:"terms_to_charge,omitempty"`
	ScheduleType          SubscriptionScheduleType                                       `json:"schedule_type,omitempty"`
	SpecificDatesSchedule []*SubscriptionEditAdvanceInvoiceScheduleSpecificDatesSchedule `json:"specific_dates_schedule,omitempty"`
	FixedIntervalSchedule *SubscriptionEditAdvanceInvoiceScheduleFixedIntervalSchedule   `json:"fixed_interval_schedule,omitempty"`
	apiRequest            `json:"-" form:"-"`
}

func (r *SubscriptionEditAdvanceInvoiceScheduleRequest) payload() any { return r }

// input sub resource params multi
type SubscriptionEditAdvanceInvoiceScheduleSpecificDatesSchedule struct {
	Id            string `json:"id,omitempty"`
	TermsToCharge *int32 `json:"terms_to_charge,omitempty"`
	Date          *int64 `json:"date,omitempty"`
}

// input sub resource params single
type SubscriptionEditAdvanceInvoiceScheduleFixedIntervalSchedule struct {
	NumberOfOccurrences *int32                                         `json:"number_of_occurrences,omitempty"`
	DaysBeforeRenewal   *int32                                         `json:"days_before_renewal,omitempty"`
	EndScheduleOn       SubscriptionFixedIntervalScheduleEndScheduleOn `json:"end_schedule_on,omitempty"`
	EndDate             *int64                                         `json:"end_date,omitempty"`
}

type SubscriptionRemoveAdvanceInvoiceScheduleRequest struct {
	SpecificDatesSchedule []*SubscriptionRemoveAdvanceInvoiceScheduleSpecificDatesSchedule `json:"specific_dates_schedule,omitempty"`
	apiRequest            `json:"-" form:"-"`
}

func (r *SubscriptionRemoveAdvanceInvoiceScheduleRequest) payload() any { return r }

// input sub resource params multi
type SubscriptionRemoveAdvanceInvoiceScheduleSpecificDatesSchedule struct {
	Id string `json:"id,omitempty"`
}

type SubscriptionRegenerateInvoiceRequest struct {
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
	Prorate            *bool  `json:"prorate,omitempty"`
	InvoiceImmediately *bool  `json:"invoice_immediately,omitempty"`
	apiRequest         `json:"-" form:"-"`
}

func (r *SubscriptionRegenerateInvoiceRequest) payload() any { return r }

type SubscriptionImportSubscriptionRequest struct {
	Id                                string                                                  `json:"id,omitempty"`
	Customer                          *SubscriptionImportSubscriptionCustomer                 `json:"customer,omitempty"`
	ClientProfileId                   string                                                  `json:"client_profile_id,omitempty"`
	PlanId                            string                                                  `json:"plan_id"`
	PlanQuantity                      *int32                                                  `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                                                  `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int64                                                  `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                                                  `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int64                                                  `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                                                  `json:"trial_end,omitempty"`
	BillingCycles                     *int32                                                  `json:"billing_cycles,omitempty"`
	Addons                            []*SubscriptionImportSubscriptionAddon                  `json:"addons,omitempty"`
	EventBasedAddons                  []*SubscriptionImportSubscriptionEventBasedAddon        `json:"event_based_addons,omitempty"`
	ChargedEventBasedAddons           []*SubscriptionImportSubscriptionChargedEventBasedAddon `json:"charged_event_based_addons,omitempty"`
	StartDate                         *int64                                                  `json:"start_date,omitempty"`
	AutoCollection                    SubscriptionAutoCollection                              `json:"auto_collection,omitempty"`
	PoNumber                          string                                                  `json:"po_number,omitempty"`
	CouponIds                         []string                                                `json:"coupon_ids,omitempty"`
	ContractTerm                      *SubscriptionImportSubscriptionContractTerm             `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                                  `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	Status                            SubscriptionStatus                                      `json:"status"`
	CurrentTermEnd                    *int64                                                  `json:"current_term_end,omitempty"`
	CurrentTermStart                  *int64                                                  `json:"current_term_start,omitempty"`
	TrialStart                        *int64                                                  `json:"trial_start,omitempty"`
	CancelledAt                       *int64                                                  `json:"cancelled_at,omitempty"`
	StartedAt                         *int64                                                  `json:"started_at,omitempty"`
	ActivatedAt                       *int64                                                  `json:"activated_at,omitempty"`
	PauseDate                         *int64                                                  `json:"pause_date,omitempty"`
	ResumeDate                        *int64                                                  `json:"resume_date,omitempty"`
	CreateCurrentTermInvoice          *bool                                                   `json:"create_current_term_invoice,omitempty"`
	Card                              *SubscriptionImportSubscriptionCard                     `json:"card,omitempty"`
	PaymentMethod                     *SubscriptionImportSubscriptionPaymentMethod            `json:"payment_method,omitempty"`
	BillingAddress                    *SubscriptionImportSubscriptionBillingAddress           `json:"billing_address,omitempty"`
	ShippingAddress                   *SubscriptionImportSubscriptionShippingAddress          `json:"shipping_address,omitempty"`
	AffiliateToken                    string                                                  `json:"affiliate_token,omitempty"`
	InvoiceNotes                      string                                                  `json:"invoice_notes,omitempty"`
	MetaData                          map[string]interface{}                                  `json:"meta_data,omitempty"`
	Transaction                       *SubscriptionImportSubscriptionTransaction              `json:"transaction,omitempty"`
	Coupons                           []*SubscriptionImportSubscriptionCoupon                 `json:"coupons,omitempty"`
	apiRequest                        `json:"-" form:"-"`
}

func (r *SubscriptionImportSubscriptionRequest) payload() any { return r }

// input sub resource params single
type SubscriptionImportSubscriptionCustomer struct {
	Id                      string                                      `json:"id,omitempty"`
	Email                   string                                      `json:"email,omitempty"`
	FirstName               string                                      `json:"first_name,omitempty"`
	LastName                string                                      `json:"last_name,omitempty"`
	Company                 string                                      `json:"company,omitempty"`
	Phone                   string                                      `json:"phone,omitempty"`
	Locale                  string                                      `json:"locale,omitempty"`
	Taxability              SubscriptionCustomerTaxability              `json:"taxability,omitempty"`
	EntityCode              SubscriptionCustomerEntityCode              `json:"entity_code,omitempty"`
	ExemptNumber            string                                      `json:"exempt_number,omitempty"`
	NetTermDays             *int32                                      `json:"net_term_days,omitempty"`
	TaxjarExemptionCategory SubscriptionCustomerTaxjarExemptionCategory `json:"taxjar_exemption_category,omitempty"`
	CustomerType            SubscriptionCustomerCustomerType            `json:"customer_type,omitempty"`
	AutoCollection          SubscriptionCustomerAutoCollection          `json:"auto_collection,omitempty"`
	AllowDirectDebit        *bool                                       `json:"allow_direct_debit,omitempty"`
	VatNumber               string                                      `json:"vat_number,omitempty"`
	VatNumberPrefix         string                                      `json:"vat_number_prefix,omitempty"`
}

// input sub resource params multi
type SubscriptionImportSubscriptionAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
}

// input sub resource params multi
type SubscriptionImportSubscriptionEventBasedAddon struct {
	Id                  string                             `json:"id,omitempty"`
	Quantity            *int32                             `json:"quantity,omitempty"`
	UnitPrice           *int64                             `json:"unit_price,omitempty"`
	QuantityInDecimal   string                             `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string                             `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodInDays *int32                             `json:"service_period_in_days,omitempty"`
	OnEvent             SubscriptionEventBasedAddonOnEvent `json:"on_event,omitempty"`
	ChargeOnce          *bool                              `json:"charge_once,omitempty"`
}

// input sub resource params multi
type SubscriptionImportSubscriptionChargedEventBasedAddon struct {
	Id            string `json:"id,omitempty"`
	LastChargedAt *int64 `json:"last_charged_at,omitempty"`
}

// input sub resource params single
type SubscriptionImportSubscriptionContractTerm struct {
	Id                         string                                  `json:"id,omitempty"`
	CreatedAt                  *int64                                  `json:"created_at,omitempty"`
	ContractStart              *int64                                  `json:"contract_start,omitempty"`
	BillingCycle               *int32                                  `json:"billing_cycle,omitempty"`
	TotalAmountRaised          *int64                                  `json:"total_amount_raised,omitempty"`
	TotalAmountRaisedBeforeTax *int64                                  `json:"total_amount_raised_before_tax,omitempty"`
	ActionAtTermEnd            SubscriptionContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod   *int32                                  `json:"cancellation_cutoff_period,omitempty"`
}

// input sub resource params single
type SubscriptionImportSubscriptionCard struct {
	Gateway               SubscriptionCardGateway         `json:"gateway,omitempty"`
	GatewayAccountId      string                          `json:"gateway_account_id,omitempty"`
	TmpToken              string                          `json:"tmp_token,omitempty"`
	FirstName             string                          `json:"first_name,omitempty"`
	LastName              string                          `json:"last_name,omitempty"`
	Number                string                          `json:"number,omitempty"`
	ExpiryMonth           *int32                          `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                          `json:"expiry_year,omitempty"`
	Cvv                   string                          `json:"cvv,omitempty"`
	PreferredScheme       SubscriptionCardPreferredScheme `json:"preferred_scheme,omitempty"`
	BillingAddr1          string                          `json:"billing_addr1,omitempty"`
	BillingAddr2          string                          `json:"billing_addr2,omitempty"`
	BillingCity           string                          `json:"billing_city,omitempty"`
	BillingStateCode      string                          `json:"billing_state_code,omitempty"`
	BillingState          string                          `json:"billing_state,omitempty"`
	BillingZip            string                          `json:"billing_zip,omitempty"`
	BillingCountry        string                          `json:"billing_country,omitempty"`
	AdditionalInformation map[string]interface{}          `json:"additional_information,omitempty"`
}

// input sub resource params single
type SubscriptionImportSubscriptionPaymentMethod struct {
	Type                  SubscriptionPaymentMethodType    `json:"type,omitempty"`
	Gateway               SubscriptionPaymentMethodGateway `json:"gateway,omitempty"`
	GatewayAccountId      string                           `json:"gateway_account_id,omitempty"`
	ReferenceId           string                           `json:"reference_id,omitempty"`
	IssuingCountry        string                           `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{}           `json:"additional_information,omitempty"`
}

// input sub resource params single
type SubscriptionImportSubscriptionBillingAddress struct {
	FirstName        string                                     `json:"first_name,omitempty"`
	LastName         string                                     `json:"last_name,omitempty"`
	Email            string                                     `json:"email,omitempty"`
	Company          string                                     `json:"company,omitempty"`
	Phone            string                                     `json:"phone,omitempty"`
	Line1            string                                     `json:"line1,omitempty"`
	Line2            string                                     `json:"line2,omitempty"`
	Line3            string                                     `json:"line3,omitempty"`
	City             string                                     `json:"city,omitempty"`
	StateCode        string                                     `json:"state_code,omitempty"`
	State            string                                     `json:"state,omitempty"`
	Zip              string                                     `json:"zip,omitempty"`
	Country          string                                     `json:"country,omitempty"`
	ValidationStatus SubscriptionBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type SubscriptionImportSubscriptionShippingAddress struct {
	FirstName        string                                      `json:"first_name,omitempty"`
	LastName         string                                      `json:"last_name,omitempty"`
	Email            string                                      `json:"email,omitempty"`
	Company          string                                      `json:"company,omitempty"`
	Phone            string                                      `json:"phone,omitempty"`
	Line1            string                                      `json:"line1,omitempty"`
	Line2            string                                      `json:"line2,omitempty"`
	Line3            string                                      `json:"line3,omitempty"`
	City             string                                      `json:"city,omitempty"`
	StateCode        string                                      `json:"state_code,omitempty"`
	State            string                                      `json:"state,omitempty"`
	Zip              string                                      `json:"zip,omitempty"`
	Country          string                                      `json:"country,omitempty"`
	ValidationStatus SubscriptionShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type SubscriptionImportSubscriptionTransaction struct {
	Amount          *int64                               `json:"amount,omitempty"`
	PaymentMethod   SubscriptionTransactionPaymentMethod `json:"payment_method,omitempty"`
	ReferenceNumber string                               `json:"reference_number,omitempty"`
	Date            *int64                               `json:"date,omitempty"`
}

// input sub resource params multi
type SubscriptionImportSubscriptionCoupon struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}

type SubscriptionImportForCustomerRequest struct {
	Id                                string                                                 `json:"id,omitempty"`
	PlanId                            string                                                 `json:"plan_id"`
	PlanQuantity                      *int32                                                 `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                                                 `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int64                                                 `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                                                 `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int64                                                 `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                                                 `json:"trial_end,omitempty"`
	BillingCycles                     *int32                                                 `json:"billing_cycles,omitempty"`
	Addons                            []*SubscriptionImportForCustomerAddon                  `json:"addons,omitempty"`
	EventBasedAddons                  []*SubscriptionImportForCustomerEventBasedAddon        `json:"event_based_addons,omitempty"`
	ChargedEventBasedAddons           []*SubscriptionImportForCustomerChargedEventBasedAddon `json:"charged_event_based_addons,omitempty"`
	StartDate                         *int64                                                 `json:"start_date,omitempty"`
	AutoCollection                    SubscriptionAutoCollection                             `json:"auto_collection,omitempty"`
	PoNumber                          string                                                 `json:"po_number,omitempty"`
	CouponIds                         []string                                               `json:"coupon_ids,omitempty"`
	PaymentSourceId                   string                                                 `json:"payment_source_id,omitempty"`
	Status                            SubscriptionStatus                                     `json:"status"`
	CurrentTermEnd                    *int64                                                 `json:"current_term_end,omitempty"`
	CurrentTermStart                  *int64                                                 `json:"current_term_start,omitempty"`
	TrialStart                        *int64                                                 `json:"trial_start,omitempty"`
	CancelledAt                       *int64                                                 `json:"cancelled_at,omitempty"`
	StartedAt                         *int64                                                 `json:"started_at,omitempty"`
	ActivatedAt                       *int64                                                 `json:"activated_at,omitempty"`
	PauseDate                         *int64                                                 `json:"pause_date,omitempty"`
	ResumeDate                        *int64                                                 `json:"resume_date,omitempty"`
	ContractTerm                      *SubscriptionImportForCustomerContractTerm             `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                                 `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	CreateCurrentTermInvoice          *bool                                                  `json:"create_current_term_invoice,omitempty"`
	Transaction                       *SubscriptionImportForCustomerTransaction              `json:"transaction,omitempty"`
	ShippingAddress                   *SubscriptionImportForCustomerShippingAddress          `json:"shipping_address,omitempty"`
	InvoiceNotes                      string                                                 `json:"invoice_notes,omitempty"`
	MetaData                          map[string]interface{}                                 `json:"meta_data,omitempty"`
	Coupons                           []*SubscriptionImportForCustomerCoupon                 `json:"coupons,omitempty"`
	apiRequest                        `json:"-" form:"-"`
}

func (r *SubscriptionImportForCustomerRequest) payload() any { return r }

// input sub resource params multi
type SubscriptionImportForCustomerAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
}

// input sub resource params multi
type SubscriptionImportForCustomerEventBasedAddon struct {
	Id                  string                             `json:"id,omitempty"`
	Quantity            *int32                             `json:"quantity,omitempty"`
	UnitPrice           *int64                             `json:"unit_price,omitempty"`
	QuantityInDecimal   string                             `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string                             `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodInDays *int32                             `json:"service_period_in_days,omitempty"`
	OnEvent             SubscriptionEventBasedAddonOnEvent `json:"on_event,omitempty"`
	ChargeOnce          *bool                              `json:"charge_once,omitempty"`
}

// input sub resource params multi
type SubscriptionImportForCustomerChargedEventBasedAddon struct {
	Id            string `json:"id,omitempty"`
	LastChargedAt *int64 `json:"last_charged_at,omitempty"`
}

// input sub resource params single
type SubscriptionImportForCustomerContractTerm struct {
	Id                         string                                  `json:"id,omitempty"`
	CreatedAt                  *int64                                  `json:"created_at,omitempty"`
	ContractStart              *int64                                  `json:"contract_start,omitempty"`
	BillingCycle               *int32                                  `json:"billing_cycle,omitempty"`
	TotalAmountRaised          *int64                                  `json:"total_amount_raised,omitempty"`
	TotalAmountRaisedBeforeTax *int64                                  `json:"total_amount_raised_before_tax,omitempty"`
	ActionAtTermEnd            SubscriptionContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod   *int32                                  `json:"cancellation_cutoff_period,omitempty"`
}

// input sub resource params single
type SubscriptionImportForCustomerTransaction struct {
	Amount          *int64                               `json:"amount,omitempty"`
	PaymentMethod   SubscriptionTransactionPaymentMethod `json:"payment_method,omitempty"`
	ReferenceNumber string                               `json:"reference_number,omitempty"`
	Date            *int64                               `json:"date,omitempty"`
}

// input sub resource params single
type SubscriptionImportForCustomerShippingAddress struct {
	FirstName        string                                      `json:"first_name,omitempty"`
	LastName         string                                      `json:"last_name,omitempty"`
	Email            string                                      `json:"email,omitempty"`
	Company          string                                      `json:"company,omitempty"`
	Phone            string                                      `json:"phone,omitempty"`
	Line1            string                                      `json:"line1,omitempty"`
	Line2            string                                      `json:"line2,omitempty"`
	Line3            string                                      `json:"line3,omitempty"`
	City             string                                      `json:"city,omitempty"`
	StateCode        string                                      `json:"state_code,omitempty"`
	State            string                                      `json:"state,omitempty"`
	Zip              string                                      `json:"zip,omitempty"`
	Country          string                                      `json:"country,omitempty"`
	ValidationStatus SubscriptionShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params multi
type SubscriptionImportForCustomerCoupon struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}

type SubscriptionImportContractTermRequest struct {
	ContractTerm                      *SubscriptionImportContractTermContractTerm `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                      `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	apiRequest                        `json:"-" form:"-"`
}

func (r *SubscriptionImportContractTermRequest) payload() any { return r }

// input sub resource params single
type SubscriptionImportContractTermContractTerm struct {
	Id                          string                                  `json:"id,omitempty"`
	CreatedAt                   *int64                                  `json:"created_at,omitempty"`
	ContractStart               *int64                                  `json:"contract_start,omitempty"`
	ContractEnd                 *int64                                  `json:"contract_end,omitempty"`
	Status                      SubscriptionContractTermStatus          `json:"status,omitempty"`
	TotalAmountRaised           *int64                                  `json:"total_amount_raised,omitempty"`
	TotalAmountRaisedBeforeTax  *int64                                  `json:"total_amount_raised_before_tax,omitempty"`
	TotalContractValue          *int64                                  `json:"total_contract_value,omitempty"`
	TotalContractValueBeforeTax *int64                                  `json:"total_contract_value_before_tax,omitempty"`
	BillingCycle                *int32                                  `json:"billing_cycle,omitempty"`
	ActionAtTermEnd             SubscriptionContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod    *int32                                  `json:"cancellation_cutoff_period,omitempty"`
}

type SubscriptionImportUnbilledChargesRequest struct {
	UnbilledCharges []*SubscriptionImportUnbilledChargesUnbilledCharge `json:"unbilled_charges,omitempty"`
	Discounts       []*SubscriptionImportUnbilledChargesDiscount       `json:"discounts,omitempty"`
	Tiers           []*SubscriptionImportUnbilledChargesTier           `json:"tiers,omitempty"`
	apiRequest      `json:"-" form:"-"`
}

func (r *SubscriptionImportUnbilledChargesRequest) payload() any { return r }

// input sub resource params multi
type SubscriptionImportUnbilledChargesUnbilledCharge struct {
	Id                  string                               `json:"id,omitempty"`
	DateFrom            *int64                               `json:"date_from"`
	DateTo              *int64                               `json:"date_to"`
	EntityType          SubscriptionUnbilledChargeEntityType `json:"entity_type"`
	EntityId            string                               `json:"entity_id,omitempty"`
	Description         string                               `json:"description,omitempty"`
	UnitAmount          *int64                               `json:"unit_amount,omitempty"`
	Quantity            *int32                               `json:"quantity,omitempty"`
	Amount              *int64                               `json:"amount,omitempty"`
	UnitAmountInDecimal string                               `json:"unit_amount_in_decimal,omitempty"`
	QuantityInDecimal   string                               `json:"quantity_in_decimal,omitempty"`
	AmountInDecimal     string                               `json:"amount_in_decimal,omitempty"`
	DiscountAmount      *int64                               `json:"discount_amount,omitempty"`
	UseForProration     *bool                                `json:"use_for_proration,omitempty"`
	IsAdvanceCharge     *bool                                `json:"is_advance_charge,omitempty"`
}

// input sub resource params multi
type SubscriptionImportUnbilledChargesDiscount struct {
	UnbilledChargeId string                         `json:"unbilled_charge_id,omitempty"`
	EntityType       SubscriptionDiscountEntityType `json:"entity_type,omitempty"`
	EntityId         string                         `json:"entity_id,omitempty"`
	Description      string                         `json:"description,omitempty"`
	Amount           *int64                         `json:"amount"`
}

// input sub resource params multi
type SubscriptionImportUnbilledChargesTier struct {
	UnbilledChargeId      string `json:"unbilled_charge_id"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	QuantityUsed          *int32 `json:"quantity_used,omitempty"`
	UnitAmount            *int64 `json:"unit_amount,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	QuantityUsedInDecimal string `json:"quantity_used_in_decimal,omitempty"`
	UnitAmountInDecimal   string `json:"unit_amount_in_decimal,omitempty"`
}

type SubscriptionImportForItemsRequest struct {
	ExhaustedCouponIds                []string                                      `json:"exhausted_coupon_ids,omitempty"`
	Id                                string                                        `json:"id,omitempty"`
	TrialEnd                          *int64                                        `json:"trial_end,omitempty"`
	BillingCycles                     *int32                                        `json:"billing_cycles,omitempty"`
	SubscriptionItems                 []*SubscriptionImportForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	Discounts                         []*SubscriptionImportForItemsDiscount         `json:"discounts,omitempty"`
	ChargedItems                      []*SubscriptionImportForItemsChargedItem      `json:"charged_items,omitempty"`
	ItemTiers                         []*SubscriptionImportForItemsItemTier         `json:"item_tiers,omitempty"`
	NetTermDays                       *int32                                        `json:"net_term_days,omitempty"`
	StartDate                         *int64                                        `json:"start_date,omitempty"`
	AutoCollection                    SubscriptionAutoCollection                    `json:"auto_collection,omitempty"`
	PoNumber                          string                                        `json:"po_number,omitempty"`
	CouponIds                         []string                                      `json:"coupon_ids,omitempty"`
	PaymentSourceId                   string                                        `json:"payment_source_id,omitempty"`
	Status                            SubscriptionStatus                            `json:"status"`
	CurrentTermEnd                    *int64                                        `json:"current_term_end,omitempty"`
	CurrentTermStart                  *int64                                        `json:"current_term_start,omitempty"`
	TrialStart                        *int64                                        `json:"trial_start,omitempty"`
	CancelledAt                       *int64                                        `json:"cancelled_at,omitempty"`
	StartedAt                         *int64                                        `json:"started_at,omitempty"`
	ActivatedAt                       *int64                                        `json:"activated_at,omitempty"`
	PauseDate                         *int64                                        `json:"pause_date,omitempty"`
	ResumeDate                        *int64                                        `json:"resume_date,omitempty"`
	ContractTerm                      *SubscriptionImportForItemsContractTerm       `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                        `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	CreateCurrentTermInvoice          *bool                                         `json:"create_current_term_invoice,omitempty"`
	Transaction                       *SubscriptionImportForItemsTransaction        `json:"transaction,omitempty"`
	ShippingAddress                   *SubscriptionImportForItemsShippingAddress    `json:"shipping_address,omitempty"`
	InvoiceNotes                      string                                        `json:"invoice_notes,omitempty"`
	MetaData                          map[string]interface{}                        `json:"meta_data,omitempty"`
	CancelReasonCode                  string                                        `json:"cancel_reason_code,omitempty"`
	CreatePendingInvoices             *bool                                         `json:"create_pending_invoices,omitempty"`
	AutoCloseInvoices                 *bool                                         `json:"auto_close_invoices,omitempty"`
	Coupons                           []*SubscriptionImportForItemsCoupon           `json:"coupons,omitempty"`
	apiRequest                        `json:"-" form:"-"`
}

func (r *SubscriptionImportForItemsRequest) payload() any { return r }

// input sub resource params multi
type SubscriptionImportForItemsSubscriptionItem struct {
	ItemPriceId        string                                    `json:"item_price_id"`
	Quantity           *int32                                    `json:"quantity,omitempty"`
	QuantityInDecimal  string                                    `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64                                    `json:"unit_price,omitempty"`
	UnitPriceInDecimal string                                    `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32                                    `json:"billing_cycles,omitempty"`
	TrialEnd           *int64                                    `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32                                    `json:"service_period_days,omitempty"`
	ChargeOnEvent      SubscriptionSubscriptionItemChargeOnEvent `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool                                     `json:"charge_once,omitempty"`
	ItemType           SubscriptionSubscriptionItemItemType      `json:"item_type,omitempty"`
}

// input sub resource params multi
type SubscriptionImportForItemsDiscount struct {
	ApplyOn       SubscriptionDiscountApplyOn      `json:"apply_on,omitempty"`
	DurationType  SubscriptionDiscountDurationType `json:"duration_type"`
	Percentage    *float64                         `json:"percentage,omitempty"`
	Amount        *int64                           `json:"amount,omitempty"`
	Period        *int32                           `json:"period,omitempty"`
	PeriodUnit    SubscriptionDiscountPeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool                            `json:"included_in_mrr,omitempty"`
	ItemPriceId   string                           `json:"item_price_id,omitempty"`
	Quantity      *int32                           `json:"quantity,omitempty"`
}

// input sub resource params multi
type SubscriptionImportForItemsChargedItem struct {
	ItemPriceId   string `json:"item_price_id,omitempty"`
	LastChargedAt *int64 `json:"last_charged_at,omitempty"`
}

// input sub resource params multi
type SubscriptionImportForItemsItemTier struct {
	ItemPriceId           string                          `json:"item_price_id,omitempty"`
	StartingUnit          *int32                          `json:"starting_unit,omitempty"`
	EndingUnit            *int32                          `json:"ending_unit,omitempty"`
	Price                 *int64                          `json:"price,omitempty"`
	StartingUnitInDecimal string                          `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string                          `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string                          `json:"price_in_decimal,omitempty"`
	PricingType           SubscriptionItemTierPricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32                          `json:"package_size,omitempty"`
}

// input sub resource params single
type SubscriptionImportForItemsContractTerm struct {
	Id                         string                                  `json:"id,omitempty"`
	CreatedAt                  *int64                                  `json:"created_at,omitempty"`
	ContractStart              *int64                                  `json:"contract_start,omitempty"`
	BillingCycle               *int32                                  `json:"billing_cycle,omitempty"`
	TotalAmountRaised          *int64                                  `json:"total_amount_raised,omitempty"`
	TotalAmountRaisedBeforeTax *int64                                  `json:"total_amount_raised_before_tax,omitempty"`
	ActionAtTermEnd            SubscriptionContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod   *int32                                  `json:"cancellation_cutoff_period,omitempty"`
}

// input sub resource params single
type SubscriptionImportForItemsTransaction struct {
	Amount          *int64                               `json:"amount,omitempty"`
	PaymentMethod   SubscriptionTransactionPaymentMethod `json:"payment_method,omitempty"`
	ReferenceNumber string                               `json:"reference_number,omitempty"`
	Date            *int64                               `json:"date,omitempty"`
}

// input sub resource params single
type SubscriptionImportForItemsShippingAddress struct {
	FirstName        string                                      `json:"first_name,omitempty"`
	LastName         string                                      `json:"last_name,omitempty"`
	Email            string                                      `json:"email,omitempty"`
	Company          string                                      `json:"company,omitempty"`
	Phone            string                                      `json:"phone,omitempty"`
	Line1            string                                      `json:"line1,omitempty"`
	Line2            string                                      `json:"line2,omitempty"`
	Line3            string                                      `json:"line3,omitempty"`
	City             string                                      `json:"city,omitempty"`
	StateCode        string                                      `json:"state_code,omitempty"`
	State            string                                      `json:"state,omitempty"`
	Zip              string                                      `json:"zip,omitempty"`
	Country          string                                      `json:"country,omitempty"`
	ValidationStatus SubscriptionShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params multi
type SubscriptionImportForItemsCoupon struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}

type SubscriptionOverrideBillingProfileRequest struct {
	PaymentSourceId string                     `json:"payment_source_id,omitempty"`
	AutoCollection  SubscriptionAutoCollection `json:"auto_collection,omitempty"`
	apiRequest      `json:"-" form:"-"`
}

func (r *SubscriptionOverrideBillingProfileRequest) payload() any { return r }

type SubscriptionPauseRequest struct {
	PauseOption             SubscriptionPauseOption             `json:"pause_option,omitempty"`
	PauseDate               *int64                              `json:"pause_date,omitempty"`
	UnbilledChargesHandling SubscriptionUnbilledChargesHandling `json:"unbilled_charges_handling,omitempty"`
	InvoiceDunningHandling  SubscriptionInvoiceDunningHandling  `json:"invoice_dunning_handling,omitempty"`
	SkipBillingCycles       *int32                              `json:"skip_billing_cycles,omitempty"`
	ResumeDate              *int64                              `json:"resume_date,omitempty"`
	apiRequest              `json:"-" form:"-"`
}

func (r *SubscriptionPauseRequest) payload() any { return r }

type SubscriptionCancelRequest struct {
	CancelOption                      SubscriptionCancelOption                      `json:"cancel_option,omitempty"`
	EndOfTerm                         *bool                                         `json:"end_of_term,omitempty"`
	CancelAt                          *int64                                        `json:"cancel_at,omitempty"`
	CreditOptionForCurrentTermCharges SubscriptionCreditOptionForCurrentTermCharges `json:"credit_option_for_current_term_charges,omitempty"`
	UnbilledChargesOption             SubscriptionUnbilledChargesOption             `json:"unbilled_charges_option,omitempty"`
	AccountReceivablesHandling        SubscriptionAccountReceivablesHandling        `json:"account_receivables_handling,omitempty"`
	RefundableCreditsHandling         SubscriptionRefundableCreditsHandling         `json:"refundable_credits_handling,omitempty"`
	ContractTermCancelOption          SubscriptionContractTermCancelOption          `json:"contract_term_cancel_option,omitempty"`
	InvoiceDate                       *int64                                        `json:"invoice_date,omitempty"`
	EventBasedAddons                  []*SubscriptionCancelEventBasedAddon          `json:"event_based_addons,omitempty"`
	CancelReasonCode                  string                                        `json:"cancel_reason_code,omitempty"`
	apiRequest                        `json:"-" form:"-"`
}

func (r *SubscriptionCancelRequest) payload() any { return r }

// input sub resource params multi
type SubscriptionCancelEventBasedAddon struct {
	Id                  string `json:"id,omitempty"`
	Quantity            *int32 `json:"quantity,omitempty"`
	UnitPrice           *int64 `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32 `json:"service_period_in_days,omitempty"`
}

type SubscriptionCancelForItemsRequest struct {
	CancelOption                      SubscriptionCancelOption                      `json:"cancel_option,omitempty"`
	EndOfTerm                         *bool                                         `json:"end_of_term,omitempty"`
	CancelAt                          *int64                                        `json:"cancel_at,omitempty"`
	CreditOptionForCurrentTermCharges SubscriptionCreditOptionForCurrentTermCharges `json:"credit_option_for_current_term_charges,omitempty"`
	UnbilledChargesOption             SubscriptionUnbilledChargesOption             `json:"unbilled_charges_option,omitempty"`
	AccountReceivablesHandling        SubscriptionAccountReceivablesHandling        `json:"account_receivables_handling,omitempty"`
	RefundableCreditsHandling         SubscriptionRefundableCreditsHandling         `json:"refundable_credits_handling,omitempty"`
	ContractTermCancelOption          SubscriptionContractTermCancelOption          `json:"contract_term_cancel_option,omitempty"`
	InvoiceDate                       *int64                                        `json:"invoice_date,omitempty"`
	SubscriptionItems                 []*SubscriptionCancelForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	CancelReasonCode                  string                                        `json:"cancel_reason_code,omitempty"`
	apiRequest                        `json:"-" form:"-"`
}

func (r *SubscriptionCancelForItemsRequest) payload() any { return r }

// input sub resource params multi
type SubscriptionCancelForItemsSubscriptionItem struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodDays  *int32 `json:"service_period_days,omitempty"`
}

type SubscriptionResumeRequest struct {
	ResumeOption           SubscriptionResumeOption           `json:"resume_option,omitempty"`
	ResumeDate             *int64                             `json:"resume_date,omitempty"`
	ChargesHandling        SubscriptionChargesHandling        `json:"charges_handling,omitempty"`
	UnpaidInvoicesHandling SubscriptionUnpaidInvoicesHandling `json:"unpaid_invoices_handling,omitempty"`
	PaymentInitiator       SubscriptionPaymentInitiator       `json:"payment_initiator,omitempty"`
	PaymentIntent          *SubscriptionResumePaymentIntent   `json:"payment_intent,omitempty"`
	apiRequest             `json:"-" form:"-"`
}

func (r *SubscriptionResumeRequest) payload() any { return r }

// input sub resource params single
type SubscriptionResumePaymentIntent struct {
	Id                    string                                     `json:"id,omitempty"`
	GatewayAccountId      string                                     `json:"gateway_account_id,omitempty"`
	GwToken               string                                     `json:"gw_token,omitempty"`
	PaymentMethodType     SubscriptionPaymentIntentPaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                                     `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                                     `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}                     `json:"additional_information,omitempty"`
}

type SubscriptionMoveRequest struct {
	ToCustomerId      string `json:"to_customer_id"`
	CopyPaymentSource *bool  `json:"copy_payment_source,omitempty"`
	apiRequest        `json:"-" form:"-"`
}

func (r *SubscriptionMoveRequest) payload() any { return r }

// operation response
type SubscriptionCreateResponse struct {
	Subscription    *Subscription    `json:"subscription,omitempty"`
	Customer        Customer         `json:"customer,omitempty"`
	Card            Card             `json:"card,omitempty"`
	Invoice         Invoice          `json:"invoice,omitempty"`
	UnbilledCharges []UnbilledCharge `json:"unbilled_charges,omitempty"`
	apiResponse
}

// operation response
type SubscriptionCreateForCustomerResponse struct {
	Subscription    *Subscription    `json:"subscription,omitempty"`
	Customer        Customer         `json:"customer,omitempty"`
	Card            Card             `json:"card,omitempty"`
	Invoice         Invoice          `json:"invoice,omitempty"`
	UnbilledCharges []UnbilledCharge `json:"unbilled_charges,omitempty"`
	apiResponse
}

// operation response
type SubscriptionCreateWithItemsResponse struct {
	Subscription    *Subscription    `json:"subscription,omitempty"`
	Customer        Customer         `json:"customer,omitempty"`
	Card            Card             `json:"card,omitempty"`
	Invoice         Invoice          `json:"invoice,omitempty"`
	UnbilledCharges []UnbilledCharge `json:"unbilled_charges,omitempty"`
	apiResponse
}

// operation sub response
type SubscriptionListSubscriptionResponse struct {
	Subscription *Subscription `json:"subscription,omitempty"`
	Customer     Customer      `json:"customer,omitempty"`
	Card         Card          `json:"card,omitempty"`
}

// operation response
type SubscriptionListResponse struct {
	List       []*SubscriptionListSubscriptionResponse `json:"list,omitempty"`
	NextOffset string                                  `json:"next_offset,omitempty"`
	apiResponse
}

// operation sub response
type SubscriptionSubscriptionsForCustomerSubscriptionResponse struct {
	Subscription *Subscription `json:"subscription,omitempty"`
}

// operation response
type SubscriptionSubscriptionsForCustomerResponse struct {
	List       []*SubscriptionSubscriptionsForCustomerSubscriptionResponse `json:"list,omitempty"`
	NextOffset string                                                      `json:"next_offset,omitempty"`
	apiResponse
}

// operation sub response
type SubscriptionContractTermsForSubscriptionSubscriptionResponse struct {
	ContractTerm ContractTerm `json:"contract_term,omitempty"`
}

// operation response
type SubscriptionContractTermsForSubscriptionResponse struct {
	List       []*SubscriptionContractTermsForSubscriptionSubscriptionResponse `json:"list,omitempty"`
	NextOffset string                                                          `json:"next_offset,omitempty"`
	apiResponse
}

// operation sub response
type SubscriptionListDiscountsSubscriptionResponse struct {
	Discount Discount `json:"discount,omitempty"`
}

// operation response
type SubscriptionListDiscountsResponse struct {
	List       []*SubscriptionListDiscountsSubscriptionResponse `json:"list,omitempty"`
	NextOffset string                                           `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type SubscriptionRetrieveResponse struct {
	Subscription *Subscription `json:"subscription,omitempty"`
	Customer     Customer      `json:"customer,omitempty"`
	Card         Card          `json:"card,omitempty"`
	apiResponse
}

// operation response
type SubscriptionRetrieveWithScheduledChangesResponse struct {
	Subscription *Subscription `json:"subscription,omitempty"`
	Customer     Customer      `json:"customer,omitempty"`
	Card         Card          `json:"card,omitempty"`
	apiResponse
}

// operation response
type SubscriptionRemoveScheduledChangesResponse struct {
	Subscription *Subscription `json:"subscription,omitempty"`
	Customer     Customer      `json:"customer,omitempty"`
	Card         Card          `json:"card,omitempty"`
	CreditNotes  []CreditNote  `json:"credit_notes,omitempty"`
	apiResponse
}

// operation response
type SubscriptionRemoveScheduledCancellationResponse struct {
	Subscription *Subscription `json:"subscription,omitempty"`
	Customer     Customer      `json:"customer,omitempty"`
	Card         Card          `json:"card,omitempty"`
	apiResponse
}

// operation response
type SubscriptionRemoveCouponsResponse struct {
	Subscription *Subscription `json:"subscription,omitempty"`
	Customer     Customer      `json:"customer,omitempty"`
	Card         Card          `json:"card,omitempty"`
	apiResponse
}

// operation response
type SubscriptionUpdateResponse struct {
	Subscription    *Subscription    `json:"subscription,omitempty"`
	Customer        Customer         `json:"customer,omitempty"`
	Card            Card             `json:"card,omitempty"`
	Invoice         Invoice          `json:"invoice,omitempty"`
	UnbilledCharges []UnbilledCharge `json:"unbilled_charges,omitempty"`
	CreditNotes     []CreditNote     `json:"credit_notes,omitempty"`
	apiResponse
}

// operation response
type SubscriptionUpdateForItemsResponse struct {
	Subscription    *Subscription    `json:"subscription,omitempty"`
	Customer        Customer         `json:"customer,omitempty"`
	Card            Card             `json:"card,omitempty"`
	Invoice         Invoice          `json:"invoice,omitempty"`
	UnbilledCharges []UnbilledCharge `json:"unbilled_charges,omitempty"`
	CreditNotes     []CreditNote     `json:"credit_notes,omitempty"`
	apiResponse
}

// operation response
type SubscriptionChangeTermEndResponse struct {
	Subscription    *Subscription    `json:"subscription,omitempty"`
	Customer        Customer         `json:"customer,omitempty"`
	Card            Card             `json:"card,omitempty"`
	Invoice         Invoice          `json:"invoice,omitempty"`
	UnbilledCharges []UnbilledCharge `json:"unbilled_charges,omitempty"`
	CreditNotes     []CreditNote     `json:"credit_notes,omitempty"`
	apiResponse
}

// operation response
type SubscriptionReactivateResponse struct {
	Subscription    *Subscription    `json:"subscription,omitempty"`
	Customer        Customer         `json:"customer,omitempty"`
	Card            Card             `json:"card,omitempty"`
	Invoice         Invoice          `json:"invoice,omitempty"`
	UnbilledCharges []UnbilledCharge `json:"unbilled_charges,omitempty"`
	apiResponse
}

// operation response
type SubscriptionAddChargeAtTermEndResponse struct {
	Estimate Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type SubscriptionChargeAddonAtTermEndResponse struct {
	Estimate Estimate `json:"estimate,omitempty"`
	apiResponse
}

// operation response
type SubscriptionChargeFutureRenewalsResponse struct {
	Subscription            *Subscription            `json:"subscription,omitempty"`
	Customer                Customer                 `json:"customer,omitempty"`
	Card                    Card                     `json:"card,omitempty"`
	Invoice                 Invoice                  `json:"invoice,omitempty"`
	AdvanceInvoiceSchedules []AdvanceInvoiceSchedule `json:"advance_invoice_schedules,omitempty"`
	apiResponse
}

// operation response
type SubscriptionEditAdvanceInvoiceScheduleResponse struct {
	AdvanceInvoiceSchedules []AdvanceInvoiceSchedule `json:"advance_invoice_schedules,omitempty"`
	apiResponse
}

// operation response
type SubscriptionRetrieveAdvanceInvoiceScheduleResponse struct {
	AdvanceInvoiceSchedules []AdvanceInvoiceSchedule `json:"advance_invoice_schedules,omitempty"`
	apiResponse
}

// operation response
type SubscriptionRemoveAdvanceInvoiceScheduleResponse struct {
	Subscription            *Subscription            `json:"subscription,omitempty"`
	AdvanceInvoiceSchedules []AdvanceInvoiceSchedule `json:"advance_invoice_schedules,omitempty"`
	apiResponse
}

// operation response
type SubscriptionRegenerateInvoiceResponse struct {
	Invoice         Invoice          `json:"invoice,omitempty"`
	UnbilledCharges []UnbilledCharge `json:"unbilled_charges,omitempty"`
	apiResponse
}

// operation response
type SubscriptionImportSubscriptionResponse struct {
	Subscription *Subscription `json:"subscription,omitempty"`
	Customer     Customer      `json:"customer,omitempty"`
	Card         Card          `json:"card,omitempty"`
	Invoice      Invoice       `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type SubscriptionImportForCustomerResponse struct {
	Subscription *Subscription `json:"subscription,omitempty"`
	Customer     Customer      `json:"customer,omitempty"`
	Card         Card          `json:"card,omitempty"`
	Invoice      Invoice       `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type SubscriptionImportContractTermResponse struct {
	ContractTerm ContractTerm `json:"contract_term,omitempty"`
	apiResponse
}

// operation response
type SubscriptionImportUnbilledChargesResponse struct {
	UnbilledCharges []UnbilledCharge `json:"unbilled_charges,omitempty"`
	apiResponse
}

// operation response
type SubscriptionImportForItemsResponse struct {
	Subscription *Subscription `json:"subscription,omitempty"`
	Customer     Customer      `json:"customer,omitempty"`
	Card         Card          `json:"card,omitempty"`
	Invoice      Invoice       `json:"invoice,omitempty"`
	apiResponse
}

// operation response
type SubscriptionOverrideBillingProfileResponse struct {
	Subscription  *Subscription `json:"subscription,omitempty"`
	PaymentSource PaymentSource `json:"payment_source,omitempty"`
	apiResponse
}

// operation response
type SubscriptionDeleteResponse struct {
	Subscription *Subscription `json:"subscription,omitempty"`
	Customer     Customer      `json:"customer,omitempty"`
	Card         Card          `json:"card,omitempty"`
	apiResponse
}

// operation response
type SubscriptionPauseResponse struct {
	Subscription    *Subscription    `json:"subscription,omitempty"`
	Customer        Customer         `json:"customer,omitempty"`
	Card            Card             `json:"card,omitempty"`
	Invoice         Invoice          `json:"invoice,omitempty"`
	UnbilledCharges []UnbilledCharge `json:"unbilled_charges,omitempty"`
	CreditNotes     []CreditNote     `json:"credit_notes,omitempty"`
	apiResponse
}

// operation response
type SubscriptionCancelResponse struct {
	Subscription    *Subscription    `json:"subscription,omitempty"`
	Customer        Customer         `json:"customer,omitempty"`
	Card            Card             `json:"card,omitempty"`
	Invoice         Invoice          `json:"invoice,omitempty"`
	UnbilledCharges []UnbilledCharge `json:"unbilled_charges,omitempty"`
	CreditNotes     []CreditNote     `json:"credit_notes,omitempty"`
	apiResponse
}

// operation response
type SubscriptionCancelForItemsResponse struct {
	Subscription    *Subscription    `json:"subscription,omitempty"`
	Customer        Customer         `json:"customer,omitempty"`
	Card            Card             `json:"card,omitempty"`
	Invoice         Invoice          `json:"invoice,omitempty"`
	UnbilledCharges []UnbilledCharge `json:"unbilled_charges,omitempty"`
	CreditNotes     []CreditNote     `json:"credit_notes,omitempty"`
	apiResponse
}

// operation response
type SubscriptionResumeResponse struct {
	Subscription    *Subscription    `json:"subscription,omitempty"`
	Customer        Customer         `json:"customer,omitempty"`
	Card            Card             `json:"card,omitempty"`
	Invoice         Invoice          `json:"invoice,omitempty"`
	UnbilledCharges []UnbilledCharge `json:"unbilled_charges,omitempty"`
	apiResponse
}

// operation response
type SubscriptionRemoveScheduledPauseResponse struct {
	Subscription *Subscription `json:"subscription,omitempty"`
	Customer     Customer      `json:"customer,omitempty"`
	Card         Card          `json:"card,omitempty"`
	apiResponse
}

// operation response
type SubscriptionRemoveScheduledResumptionResponse struct {
	Subscription *Subscription `json:"subscription,omitempty"`
	Customer     Customer      `json:"customer,omitempty"`
	Card         Card          `json:"card,omitempty"`
	apiResponse
}

// operation response
type SubscriptionMoveResponse struct {
	Subscription *Subscription `json:"subscription,omitempty"`
	apiResponse
}
