package chargebee


type AccountHolderType string

const (
    AccountHolderTypeIndividual AccountHolderType = "individual"
    AccountHolderTypeCompany AccountHolderType = "company"
)

type AccountReceivablesHandling string

const (
    AccountReceivablesHandlingNoAction AccountReceivablesHandling = "no_action"
    AccountReceivablesHandlingSchedulePaymentCollection AccountReceivablesHandling = "schedule_payment_collection"
    AccountReceivablesHandlingWriteOff AccountReceivablesHandling = "write_off"
)

type AccountType string

const (
    AccountTypeChecking AccountType = "checking"
    AccountTypeSavings AccountType = "savings"
    AccountTypeBusinessChecking AccountType = "business_checking"
    AccountTypeCurrent AccountType = "current"
)

type Action string

const (
    ActionUpsert Action = "upsert"
    ActionRemove Action = "remove"
)

type ApiVersion string

const (
    ApiVersionV1 ApiVersion = "v1"
    ApiVersionV2 ApiVersion = "v2"
)

type ApplyOn string

const (
    ApplyOnInvoiceAmount ApplyOn = "invoice_amount"
    ApplyOnSpecificItemPrice ApplyOn = "specific_item_price"
)

type AutoCollection string

const (
    AutoCollectionOn AutoCollection = "on"
    AutoCollectionOff AutoCollection = "off"
)

type AvalaraSaleType string

const (
    AvalaraSaleTypeWholesale AvalaraSaleType = "wholesale"
    AvalaraSaleTypeRetail AvalaraSaleType = "retail"
    AvalaraSaleTypeConsumed AvalaraSaleType = "consumed"
    AvalaraSaleTypeVendorUse AvalaraSaleType = "vendor_use"
)

type BillingAlignmentMode string

const (
    BillingAlignmentModeImmediate BillingAlignmentMode = "immediate"
    BillingAlignmentModeDelayed BillingAlignmentMode = "delayed"
)

type BillingDateMode string

const (
    BillingDateModeUsingDefaults BillingDateMode = "using_defaults"
    BillingDateModeManuallySet BillingDateMode = "manually_set"
)

type BillingDayOfWeekMode string

const (
    BillingDayOfWeekModeUsingDefaults BillingDayOfWeekMode = "using_defaults"
    BillingDayOfWeekModeManuallySet BillingDayOfWeekMode = "manually_set"
)

type BillingPeriodUnit string

const (
    BillingPeriodUnitDay BillingPeriodUnit = "day"
    BillingPeriodUnitWeek BillingPeriodUnit = "week"
    BillingPeriodUnitMonth BillingPeriodUnit = "month"
    BillingPeriodUnitYear BillingPeriodUnit = "year"
)

type BillingStartOption string

const (
    BillingStartOptionImmediately BillingStartOption = "immediately"
    BillingStartOptionOnSpecificDate BillingStartOption = "on_specific_date"
)

type CancelOption string

const (
    CancelOptionImmediately CancelOption = "immediately"
    CancelOptionEndOfTerm CancelOption = "end_of_term"
    CancelOptionSpecificDate CancelOption = "specific_date"
    CancelOptionEndOfBillingTerm CancelOption = "end_of_billing_term"
)

type Category string

const (
    CategoryIntroductory Category = "introductory"
    CategoryPromotional Category = "promotional"
    CategoryDeveloperDetermined Category = "developer_determined"
)

type ChangeOption string

const (
    ChangeOptionImmediately ChangeOption = "immediately"
    ChangeOptionEndOfTerm ChangeOption = "end_of_term"
    ChangeOptionSpecificDate ChangeOption = "specific_date"
)

type Channel string

const (
    ChannelWeb Channel = "web"
    ChannelAppStore Channel = "app_store"
    ChannelPlayStore Channel = "play_store"
)

type ChargeModel string

const (
    ChargeModelFullCharge ChargeModel = "full_charge"
    ChargeModelProrate ChargeModel = "prorate"
)

type ChargeOn string

const (
    ChargeOnImmediately ChargeOn = "immediately"
    ChargeOnOnEvent ChargeOn = "on_event"
)

type ChargeOnEvent string

const (
    ChargeOnEventSubscriptionCreation ChargeOnEvent = "subscription_creation"
    ChargeOnEventSubscriptionTrialStart ChargeOnEvent = "subscription_trial_start"
    ChargeOnEventPlanActivation ChargeOnEvent = "plan_activation"
    ChargeOnEventSubscriptionActivation ChargeOnEvent = "subscription_activation"
    ChargeOnEventContractTermination ChargeOnEvent = "contract_termination"
    ChargeOnEventOnDemand ChargeOnEvent = "on_demand"
)

type ChargeOnOption string

const (
    ChargeOnOptionImmediately ChargeOnOption = "immediately"
    ChargeOnOptionOnEvent ChargeOnOption = "on_event"
)

type ChargebeeResponseSchemaType string

const (
    ChargebeeResponseSchemaTypePlansAddons ChargebeeResponseSchemaType = "plans_addons"
    ChargebeeResponseSchemaTypeItems ChargebeeResponseSchemaType = "items"
    ChargebeeResponseSchemaTypeCompat ChargebeeResponseSchemaType = "compat"
)

type ChargesHandling string

const (
    ChargesHandlingInvoiceImmediately ChargesHandling = "invoice_immediately"
    ChargesHandlingAddToUnbilledCharges ChargesHandling = "add_to_unbilled_charges"
)

type ContractTermCancelOption string

const (
    ContractTermCancelOptionTerminateImmediately ContractTermCancelOption = "terminate_immediately"
    ContractTermCancelOptionEndOfContractTerm ContractTermCancelOption = "end_of_contract_term"
    ContractTermCancelOptionSpecificDate ContractTermCancelOption = "specific_date"
    ContractTermCancelOptionEndOfSubscriptionBillingTerm ContractTermCancelOption = "end_of_subscription_billing_term"
)

type CreditOptionForCurrentTermCharges string

const (
    CreditOptionForCurrentTermChargesNone CreditOptionForCurrentTermCharges = "none"
    CreditOptionForCurrentTermChargesProrate CreditOptionForCurrentTermCharges = "prorate"
    CreditOptionForCurrentTermChargesFull CreditOptionForCurrentTermCharges = "full"
)

type CreditType string

const (
    CreditTypeLoyaltyCredits CreditType = "loyalty_credits"
    CreditTypeReferralRewards CreditType = "referral_rewards"
    CreditTypeGeneral CreditType = "general"
)

type CustomerType string

const (
    CustomerTypeResidential CustomerType = "residential"
    CustomerTypeBusiness CustomerType = "business"
    CustomerTypeSeniorCitizen CustomerType = "senior_citizen"
    CustomerTypeIndustrial CustomerType = "industrial"
)

type DedupeOption string

const (
    DedupeOptionSkip DedupeOption = "skip"
    DedupeOptionUpdateExisting DedupeOption = "update_existing"
)

type DirectDebitScheme string

const (
    DirectDebitSchemeAch DirectDebitScheme = "ach"
    DirectDebitSchemeBacs DirectDebitScheme = "bacs"
    DirectDebitSchemeSepaCore DirectDebitScheme = "sepa_core"
    DirectDebitSchemeAutogiro DirectDebitScheme = "autogiro"
    DirectDebitSchemeBecs DirectDebitScheme = "becs"
    DirectDebitSchemeBecsNz DirectDebitScheme = "becs_nz"
    DirectDebitSchemePad DirectDebitScheme = "pad"
    DirectDebitSchemeNotApplicable DirectDebitScheme = "not_applicable"
)

type DiscountType string

const (
    DiscountTypeFixedAmount DiscountType = "fixed_amount"
    DiscountTypePercentage DiscountType = "percentage"
    DiscountTypePrice DiscountType = "price"
)

type DispositionType string

const (
    DispositionTypeAttachment DispositionType = "attachment"
    DispositionTypeInline DispositionType = "inline"
)

type DunningType string

const (
    DunningTypeAutoCollect DunningType = "auto_collect"
    DunningTypeOffline DunningType = "offline"
    DunningTypeDirectDebit DunningType = "direct_debit"
)

type DurationType string

const (
    DurationTypeOneTime DurationType = "one_time"
    DurationTypeForever DurationType = "forever"
    DurationTypeLimitedPeriod DurationType = "limited_period"
)

type EcheckType string

const (
    EcheckTypeWeb EcheckType = "web"
    EcheckTypePpd EcheckType = "ppd"
    EcheckTypeCcd EcheckType = "ccd"
)

type EinvoicingMethod string

const (
    EinvoicingMethodAutomatic EinvoicingMethod = "automatic"
    EinvoicingMethodManual EinvoicingMethod = "manual"
    EinvoicingMethodSiteDefault EinvoicingMethod = "site_default"
)

type EndScheduleOn string

const (
    EndScheduleOnAfterNumberOfIntervals EndScheduleOn = "after_number_of_intervals"
    EndScheduleOnSpecificDate EndScheduleOn = "specific_date"
    EndScheduleOnSubscriptionEnd EndScheduleOn = "subscription_end"
)

type EntityCode string

const (
    EntityCodeA EntityCode = "a"
    EntityCodeB EntityCode = "b"
    EntityCodeC EntityCode = "c"
    EntityCodeD EntityCode = "d"
    EntityCodeE EntityCode = "e"
    EntityCodeF EntityCode = "f"
    EntityCodeG EntityCode = "g"
    EntityCodeH EntityCode = "h"
    EntityCodeI EntityCode = "i"
    EntityCodeJ EntityCode = "j"
    EntityCodeK EntityCode = "k"
    EntityCodeL EntityCode = "l"
    EntityCodeM EntityCode = "m"
    EntityCodeN EntityCode = "n"
    EntityCodeP EntityCode = "p"
    EntityCodeQ EntityCode = "q"
    EntityCodeR EntityCode = "r"
    EntityCodeMed1 EntityCode = "med1"
    EntityCodeMed2 EntityCode = "med2"
)

type EntityType string

const (
    EntityTypeCustomer EntityType = "customer"
    EntityTypeSubscription EntityType = "subscription"
    EntityTypeCoupon EntityType = "coupon"
    EntityTypePlanItemPrice EntityType = "plan_item_price"
    EntityTypeAddonItemPrice EntityType = "addon_item_price"
    EntityTypeChargeItemPrice EntityType = "charge_item_price"
    EntityTypeInvoice EntityType = "invoice"
    EntityTypeQuote EntityType = "quote"
    EntityTypeCreditNote EntityType = "credit_note"
    EntityTypeTransaction EntityType = "transaction"
    EntityTypePlan EntityType = "plan"
    EntityTypeAddon EntityType = "addon"
    EntityTypeOrder EntityType = "order"
    EntityTypeItemFamily EntityType = "item_family"
    EntityTypeItem EntityType = "item"
    EntityTypeItemPrice EntityType = "item_price"
    EntityTypePlanItem EntityType = "plan_item"
    EntityTypeAddonItem EntityType = "addon_item"
    EntityTypeChargeItem EntityType = "charge_item"
    EntityTypePlanPrice EntityType = "plan_price"
    EntityTypeAddonPrice EntityType = "addon_price"
    EntityTypeChargePrice EntityType = "charge_price"
    EntityTypeDifferentialPrice EntityType = "differential_price"
    EntityTypeAttachedItem EntityType = "attached_item"
    EntityTypeFeature EntityType = "feature"
    EntityTypeSubscriptionEntitlement EntityType = "subscription_entitlement"
    EntityTypeItemEntitlement EntityType = "item_entitlement"
    EntityTypeBusinessEntity EntityType = "business_entity"
    EntityTypePriceVariant EntityType = "price_variant"
    EntityTypeOmnichannelSubscription EntityType = "omnichannel_subscription"
    EntityTypeOmnichannelSubscriptionItem EntityType = "omnichannel_subscription_item"
    EntityTypeOmnichannelTransaction EntityType = "omnichannel_transaction"
    EntityTypeRecordedPurchase EntityType = "recorded_purchase"
    EntityTypeOmnichannelSubscriptionItemScheduledChange EntityType = "omnichannel_subscription_item_scheduled_change"
    EntityTypeSalesOrder EntityType = "sales_order"
    EntityTypeOmnichannelOneTimeOrder EntityType = "omnichannel_one_time_order"
    EntityTypeOmnichannelOneTimeOrderItem EntityType = "omnichannel_one_time_order_item"
    EntityTypeUsageFile EntityType = "usage_file"
    EntityTypeBusinessRule EntityType = "business_rule"
    EntityTypeRuleset EntityType = "ruleset"
)

type EventName string

const (
    EventNameCancellationPageLoaded EventName = "cancellation_page_loaded"
)

type EventType string

const (
    EventTypeCouponCreated EventType = "coupon_created"
    EventTypeCouponUpdated EventType = "coupon_updated"
    EventTypeCouponDeleted EventType = "coupon_deleted"
    EventTypeCouponSetCreated EventType = "coupon_set_created"
    EventTypeCouponSetUpdated EventType = "coupon_set_updated"
    EventTypeCouponSetDeleted EventType = "coupon_set_deleted"
    EventTypeCouponCodesAdded EventType = "coupon_codes_added"
    EventTypeCouponCodesDeleted EventType = "coupon_codes_deleted"
    EventTypeCouponCodesUpdated EventType = "coupon_codes_updated"
    EventTypeCustomerCreated EventType = "customer_created"
    EventTypeCustomerChanged EventType = "customer_changed"
    EventTypeCustomerDeleted EventType = "customer_deleted"
    EventTypeCustomerMovedOut EventType = "customer_moved_out"
    EventTypeCustomerMovedIn EventType = "customer_moved_in"
    EventTypePromotionalCreditsAdded EventType = "promotional_credits_added"
    EventTypePromotionalCreditsDeducted EventType = "promotional_credits_deducted"
    EventTypeSubscriptionCreated EventType = "subscription_created"
    EventTypeSubscriptionCreatedWithBackdating EventType = "subscription_created_with_backdating"
    EventTypeSubscriptionStarted EventType = "subscription_started"
    EventTypeSubscriptionTrialEndReminder EventType = "subscription_trial_end_reminder"
    EventTypeSubscriptionActivated EventType = "subscription_activated"
    EventTypeSubscriptionActivatedWithBackdating EventType = "subscription_activated_with_backdating"
    EventTypeSubscriptionChanged EventType = "subscription_changed"
    EventTypeSubscriptionTrialExtended EventType = "subscription_trial_extended"
    EventTypeMrrUpdated EventType = "mrr_updated"
    EventTypeSubscriptionChangedWithBackdating EventType = "subscription_changed_with_backdating"
    EventTypeSubscriptionCancellationScheduled EventType = "subscription_cancellation_scheduled"
    EventTypeSubscriptionCancellationReminder EventType = "subscription_cancellation_reminder"
    EventTypeSubscriptionCancelled EventType = "subscription_cancelled"
    EventTypeSubscriptionCanceledWithBackdating EventType = "subscription_canceled_with_backdating"
    EventTypeSubscriptionReactivated EventType = "subscription_reactivated"
    EventTypeSubscriptionReactivatedWithBackdating EventType = "subscription_reactivated_with_backdating"
    EventTypeSubscriptionRenewed EventType = "subscription_renewed"
    EventTypeSubscriptionItemsRenewed EventType = "subscription_items_renewed"
    EventTypeSubscriptionScheduledCancellationRemoved EventType = "subscription_scheduled_cancellation_removed"
    EventTypeSubscriptionChangesScheduled EventType = "subscription_changes_scheduled"
    EventTypeSubscriptionScheduledChangesRemoved EventType = "subscription_scheduled_changes_removed"
    EventTypeSubscriptionShippingAddressUpdated EventType = "subscription_shipping_address_updated"
    EventTypeSubscriptionDeleted EventType = "subscription_deleted"
    EventTypeSubscriptionPaused EventType = "subscription_paused"
    EventTypeSubscriptionPauseScheduled EventType = "subscription_pause_scheduled"
    EventTypeSubscriptionScheduledPauseRemoved EventType = "subscription_scheduled_pause_removed"
    EventTypeSubscriptionResumed EventType = "subscription_resumed"
    EventTypeSubscriptionResumptionScheduled EventType = "subscription_resumption_scheduled"
    EventTypeSubscriptionScheduledResumptionRemoved EventType = "subscription_scheduled_resumption_removed"
    EventTypeSubscriptionAdvanceInvoiceScheduleAdded EventType = "subscription_advance_invoice_schedule_added"
    EventTypeSubscriptionAdvanceInvoiceScheduleUpdated EventType = "subscription_advance_invoice_schedule_updated"
    EventTypeSubscriptionAdvanceInvoiceScheduleRemoved EventType = "subscription_advance_invoice_schedule_removed"
    EventTypePendingInvoiceCreated EventType = "pending_invoice_created"
    EventTypePendingInvoiceUpdated EventType = "pending_invoice_updated"
    EventTypeInvoiceGenerated EventType = "invoice_generated"
    EventTypeInvoiceGeneratedWithBackdating EventType = "invoice_generated_with_backdating"
    EventTypeInvoiceUpdated EventType = "invoice_updated"
    EventTypeInvoiceDeleted EventType = "invoice_deleted"
    EventTypeCreditNoteCreated EventType = "credit_note_created"
    EventTypeCreditNoteCreatedWithBackdating EventType = "credit_note_created_with_backdating"
    EventTypeCreditNoteUpdated EventType = "credit_note_updated"
    EventTypeCreditNoteDeleted EventType = "credit_note_deleted"
    EventTypePaymentSchedulesCreated EventType = "payment_schedules_created"
    EventTypePaymentSchedulesUpdated EventType = "payment_schedules_updated"
    EventTypePaymentScheduleSchemeCreated EventType = "payment_schedule_scheme_created"
    EventTypePaymentScheduleSchemeDeleted EventType = "payment_schedule_scheme_deleted"
    EventTypeSubscriptionRenewalReminder EventType = "subscription_renewal_reminder"
    EventTypeAddUsagesReminder EventType = "add_usages_reminder"
    EventTypeTransactionCreated EventType = "transaction_created"
    EventTypeTransactionUpdated EventType = "transaction_updated"
    EventTypeTransactionDeleted EventType = "transaction_deleted"
    EventTypePaymentSucceeded EventType = "payment_succeeded"
    EventTypePaymentFailed EventType = "payment_failed"
    EventTypeDunningUpdated EventType = "dunning_updated"
    EventTypePaymentRefunded EventType = "payment_refunded"
    EventTypePaymentInitiated EventType = "payment_initiated"
    EventTypeRefundInitiated EventType = "refund_initiated"
    EventTypeAuthorizationSucceeded EventType = "authorization_succeeded"
    EventTypeAuthorizationVoided EventType = "authorization_voided"
    EventTypeCardAdded EventType = "card_added"
    EventTypeCardUpdated EventType = "card_updated"
    EventTypeCardExpiryReminder EventType = "card_expiry_reminder"
    EventTypeCardExpired EventType = "card_expired"
    EventTypeCardDeleted EventType = "card_deleted"
    EventTypePaymentSourceAdded EventType = "payment_source_added"
    EventTypePaymentSourceUpdated EventType = "payment_source_updated"
    EventTypePaymentSourceDeleted EventType = "payment_source_deleted"
    EventTypePaymentSourceExpiring EventType = "payment_source_expiring"
    EventTypePaymentSourceExpired EventType = "payment_source_expired"
    EventTypePaymentSourceLocallyDeleted EventType = "payment_source_locally_deleted"
    EventTypeVirtualBankAccountAdded EventType = "virtual_bank_account_added"
    EventTypeVirtualBankAccountUpdated EventType = "virtual_bank_account_updated"
    EventTypeVirtualBankAccountDeleted EventType = "virtual_bank_account_deleted"
    EventTypeTokenCreated EventType = "token_created"
    EventTypeTokenConsumed EventType = "token_consumed"
    EventTypeTokenExpired EventType = "token_expired"
    EventTypeUnbilledChargesCreated EventType = "unbilled_charges_created"
    EventTypeUnbilledChargesVoided EventType = "unbilled_charges_voided"
    EventTypeUnbilledChargesDeleted EventType = "unbilled_charges_deleted"
    EventTypeUnbilledChargesInvoiced EventType = "unbilled_charges_invoiced"
    EventTypeOrderCreated EventType = "order_created"
    EventTypeOrderUpdated EventType = "order_updated"
    EventTypeOrderCancelled EventType = "order_cancelled"
    EventTypeOrderDelivered EventType = "order_delivered"
    EventTypeOrderReturned EventType = "order_returned"
    EventTypeOrderReadyToProcess EventType = "order_ready_to_process"
    EventTypeOrderReadyToShip EventType = "order_ready_to_ship"
    EventTypeOrderDeleted EventType = "order_deleted"
    EventTypeOrderResent EventType = "order_resent"
    EventTypeQuoteCreated EventType = "quote_created"
    EventTypeQuoteUpdated EventType = "quote_updated"
    EventTypeQuoteDeleted EventType = "quote_deleted"
    EventTypeTaxWithheldRecorded EventType = "tax_withheld_recorded"
    EventTypeTaxWithheldDeleted EventType = "tax_withheld_deleted"
    EventTypeTaxWithheldRefunded EventType = "tax_withheld_refunded"
    EventTypeGiftScheduled EventType = "gift_scheduled"
    EventTypeGiftUnclaimed EventType = "gift_unclaimed"
    EventTypeGiftClaimed EventType = "gift_claimed"
    EventTypeGiftExpired EventType = "gift_expired"
    EventTypeGiftCancelled EventType = "gift_cancelled"
    EventTypeGiftUpdated EventType = "gift_updated"
    EventTypeHierarchyCreated EventType = "hierarchy_created"
    EventTypeHierarchyDeleted EventType = "hierarchy_deleted"
    EventTypePaymentIntentCreated EventType = "payment_intent_created"
    EventTypePaymentIntentUpdated EventType = "payment_intent_updated"
    EventTypeContractTermCreated EventType = "contract_term_created"
    EventTypeContractTermRenewed EventType = "contract_term_renewed"
    EventTypeContractTermTerminated EventType = "contract_term_terminated"
    EventTypeContractTermCompleted EventType = "contract_term_completed"
    EventTypeContractTermCancelled EventType = "contract_term_cancelled"
    EventTypeItemFamilyCreated EventType = "item_family_created"
    EventTypeItemFamilyUpdated EventType = "item_family_updated"
    EventTypeItemFamilyDeleted EventType = "item_family_deleted"
    EventTypeItemCreated EventType = "item_created"
    EventTypeItemUpdated EventType = "item_updated"
    EventTypeItemDeleted EventType = "item_deleted"
    EventTypeItemPriceCreated EventType = "item_price_created"
    EventTypeItemPriceUpdated EventType = "item_price_updated"
    EventTypeItemPriceDeleted EventType = "item_price_deleted"
    EventTypeAttachedItemCreated EventType = "attached_item_created"
    EventTypeAttachedItemUpdated EventType = "attached_item_updated"
    EventTypeAttachedItemDeleted EventType = "attached_item_deleted"
    EventTypeDifferentialPriceCreated EventType = "differential_price_created"
    EventTypeDifferentialPriceUpdated EventType = "differential_price_updated"
    EventTypeDifferentialPriceDeleted EventType = "differential_price_deleted"
    EventTypeFeatureCreated EventType = "feature_created"
    EventTypeFeatureUpdated EventType = "feature_updated"
    EventTypeFeatureDeleted EventType = "feature_deleted"
    EventTypeFeatureActivated EventType = "feature_activated"
    EventTypeFeatureReactivated EventType = "feature_reactivated"
    EventTypeFeatureArchived EventType = "feature_archived"
    EventTypeItemEntitlementsUpdated EventType = "item_entitlements_updated"
    EventTypeEntitlementOverridesUpdated EventType = "entitlement_overrides_updated"
    EventTypeEntitlementOverridesRemoved EventType = "entitlement_overrides_removed"
    EventTypeItemEntitlementsRemoved EventType = "item_entitlements_removed"
    EventTypeEntitlementOverridesAutoRemoved EventType = "entitlement_overrides_auto_removed"
    EventTypeSubscriptionEntitlementsCreated EventType = "subscription_entitlements_created"
    EventTypeSubscriptionEntitlementsUpdated EventType = "subscription_entitlements_updated"
    EventTypeBusinessEntityCreated EventType = "business_entity_created"
    EventTypeBusinessEntityUpdated EventType = "business_entity_updated"
    EventTypeBusinessEntityDeleted EventType = "business_entity_deleted"
    EventTypeCustomerBusinessEntityChanged EventType = "customer_business_entity_changed"
    EventTypeSubscriptionBusinessEntityChanged EventType = "subscription_business_entity_changed"
    EventTypePurchaseCreated EventType = "purchase_created"
    EventTypeVoucherCreated EventType = "voucher_created"
    EventTypeVoucherExpired EventType = "voucher_expired"
    EventTypeVoucherCreateFailed EventType = "voucher_create_failed"
    EventTypeItemPriceEntitlementsUpdated EventType = "item_price_entitlements_updated"
    EventTypeItemPriceEntitlementsRemoved EventType = "item_price_entitlements_removed"
    EventTypeSubscriptionRampCreated EventType = "subscription_ramp_created"
    EventTypeSubscriptionRampDeleted EventType = "subscription_ramp_deleted"
    EventTypeSubscriptionRampApplied EventType = "subscription_ramp_applied"
    EventTypeSubscriptionRampDrafted EventType = "subscription_ramp_drafted"
    EventTypeSubscriptionRampUpdated EventType = "subscription_ramp_updated"
    EventTypePriceVariantCreated EventType = "price_variant_created"
    EventTypePriceVariantUpdated EventType = "price_variant_updated"
    EventTypePriceVariantDeleted EventType = "price_variant_deleted"
    EventTypeCustomerEntitlementsUpdated EventType = "customer_entitlements_updated"
    EventTypeSubscriptionMovedIn EventType = "subscription_moved_in"
    EventTypeSubscriptionMovedOut EventType = "subscription_moved_out"
    EventTypeSubscriptionMovementFailed EventType = "subscription_movement_failed"
    EventTypeOmnichannelSubscriptionCreated EventType = "omnichannel_subscription_created"
    EventTypeOmnichannelSubscriptionItemRenewed EventType = "omnichannel_subscription_item_renewed"
    EventTypeOmnichannelSubscriptionItemDowngraded EventType = "omnichannel_subscription_item_downgraded"
    EventTypeOmnichannelSubscriptionItemExpired EventType = "omnichannel_subscription_item_expired"
    EventTypeOmnichannelSubscriptionItemCancellationScheduled EventType = "omnichannel_subscription_item_cancellation_scheduled"
    EventTypeOmnichannelSubscriptionItemScheduledCancellationRemoved EventType = "omnichannel_subscription_item_scheduled_cancellation_removed"
    EventTypeOmnichannelSubscriptionItemResubscribed EventType = "omnichannel_subscription_item_resubscribed"
    EventTypeOmnichannelSubscriptionItemUpgraded EventType = "omnichannel_subscription_item_upgraded"
    EventTypeOmnichannelSubscriptionItemCancelled EventType = "omnichannel_subscription_item_cancelled"
    EventTypeOmnichannelSubscriptionImported EventType = "omnichannel_subscription_imported"
    EventTypeOmnichannelSubscriptionItemGracePeriodStarted EventType = "omnichannel_subscription_item_grace_period_started"
    EventTypeOmnichannelSubscriptionItemGracePeriodExpired EventType = "omnichannel_subscription_item_grace_period_expired"
    EventTypeOmnichannelSubscriptionItemDunningStarted EventType = "omnichannel_subscription_item_dunning_started"
    EventTypeOmnichannelSubscriptionItemDunningExpired EventType = "omnichannel_subscription_item_dunning_expired"
    EventTypeRuleCreated EventType = "rule_created"
    EventTypeRuleUpdated EventType = "rule_updated"
    EventTypeRuleDeleted EventType = "rule_deleted"
    EventTypeRecordPurchaseFailed EventType = "record_purchase_failed"
    EventTypeOmnichannelSubscriptionItemChangeScheduled EventType = "omnichannel_subscription_item_change_scheduled"
    EventTypeOmnichannelSubscriptionItemScheduledChangeRemoved EventType = "omnichannel_subscription_item_scheduled_change_removed"
    EventTypeOmnichannelSubscriptionItemReactivated EventType = "omnichannel_subscription_item_reactivated"
    EventTypeSalesOrderCreated EventType = "sales_order_created"
    EventTypeSalesOrderUpdated EventType = "sales_order_updated"
    EventTypeOmnichannelSubscriptionItemChanged EventType = "omnichannel_subscription_item_changed"
    EventTypeOmnichannelSubscriptionItemPaused EventType = "omnichannel_subscription_item_paused"
    EventTypeOmnichannelSubscriptionItemResumed EventType = "omnichannel_subscription_item_resumed"
    EventTypeOmnichannelOneTimeOrderCreated EventType = "omnichannel_one_time_order_created"
    EventTypeOmnichannelOneTimeOrderItemCancelled EventType = "omnichannel_one_time_order_item_cancelled"
    EventTypeUsageFileIngested EventType = "usage_file_ingested"
    EventTypeOmnichannelSubscriptionItemPauseScheduled EventType = "omnichannel_subscription_item_pause_scheduled"
    EventTypeOmnichannelSubscriptionMovedIn EventType = "omnichannel_subscription_moved_in"
    EventTypeOmnichannelTransactionCreated EventType = "omnichannel_transaction_created"
    EventTypePlanCreated EventType = "plan_created"
    EventTypePlanUpdated EventType = "plan_updated"
    EventTypePlanDeleted EventType = "plan_deleted"
    EventTypeAddonCreated EventType = "addon_created"
    EventTypeAddonUpdated EventType = "addon_updated"
    EventTypeAddonDeleted EventType = "addon_deleted"
)

type ExportType string

const (
    ExportTypeData ExportType = "data"
    ExportTypeImportFriendlyData ExportType = "import_friendly_data"
)

type FreePeriodUnit string

const (
    FreePeriodUnitDay FreePeriodUnit = "day"
    FreePeriodUnitWeek FreePeriodUnit = "week"
    FreePeriodUnitMonth FreePeriodUnit = "month"
    FreePeriodUnitYear FreePeriodUnit = "year"
)

type FriendOfferType string

const (
    FriendOfferTypeNone FriendOfferType = "none"
    FriendOfferTypeCoupon FriendOfferType = "coupon"
    FriendOfferTypeCouponCode FriendOfferType = "coupon_code"
)

type Gateway string

const (
    GatewayChargebee Gateway = "chargebee"
    GatewayChargebeePayments Gateway = "chargebee_payments"
    GatewayAdyen Gateway = "adyen"
    GatewayStripe Gateway = "stripe"
    GatewayWepay Gateway = "wepay"
    GatewayBraintree Gateway = "braintree"
    GatewayAuthorizeNet Gateway = "authorize_net"
    GatewayPaypalPro Gateway = "paypal_pro"
    GatewayPin Gateway = "pin"
    GatewayEway Gateway = "eway"
    GatewayEwayRapid Gateway = "eway_rapid"
    GatewayWorldpay Gateway = "worldpay"
    GatewayBalancedPayments Gateway = "balanced_payments"
    GatewayBeanstream Gateway = "beanstream"
    GatewayBluepay Gateway = "bluepay"
    GatewayElavon Gateway = "elavon"
    GatewayFirstDataGlobal Gateway = "first_data_global"
    GatewayHdfc Gateway = "hdfc"
    GatewayMigs Gateway = "migs"
    GatewayNmi Gateway = "nmi"
    GatewayOgone Gateway = "ogone"
    GatewayPaymill Gateway = "paymill"
    GatewayPaypalPayflowPro Gateway = "paypal_payflow_pro"
    GatewaySagePay Gateway = "sage_pay"
    GatewayTco Gateway = "tco"
    GatewayWirecard Gateway = "wirecard"
    GatewayAmazonPayments Gateway = "amazon_payments"
    GatewayPaypalExpressCheckout Gateway = "paypal_express_checkout"
    GatewayOrbital Gateway = "orbital"
    GatewayMonerisUs Gateway = "moneris_us"
    GatewayMoneris Gateway = "moneris"
    GatewayBluesnap Gateway = "bluesnap"
    GatewayCybersource Gateway = "cybersource"
    GatewayVantiv Gateway = "vantiv"
    GatewayCheckoutCom Gateway = "checkout_com"
    GatewayPaypal Gateway = "paypal"
    GatewayIngenicoDirect Gateway = "ingenico_direct"
    GatewayExact Gateway = "exact"
    GatewayMollie Gateway = "mollie"
    GatewayQuickbooks Gateway = "quickbooks"
    GatewayRazorpay Gateway = "razorpay"
    GatewayGlobalPayments Gateway = "global_payments"
    GatewayBankOfAmerica Gateway = "bank_of_america"
    GatewayEcentric Gateway = "ecentric"
    GatewayMetricsGlobal Gateway = "metrics_global"
    GatewayWindcave Gateway = "windcave"
    GatewayPayCom Gateway = "pay_com"
    GatewayEbanx Gateway = "ebanx"
    GatewayDlocal Gateway = "dlocal"
    GatewayNuvei Gateway = "nuvei"
    GatewaySolidgate Gateway = "solidgate"
    GatewayPaystack Gateway = "paystack"
    GatewayJpMorgan Gateway = "jp_morgan"
    GatewayDeutscheBank Gateway = "deutsche_bank"
    GatewayEzidebit Gateway = "ezidebit"
    GatewayGocardless Gateway = "gocardless"
    GatewayNotApplicable Gateway = "not_applicable"
)

type HierarchyOperationType string

const (
    HierarchyOperationTypeCompleteHierarchy HierarchyOperationType = "complete_hierarchy"
    HierarchyOperationTypeSubordinates HierarchyOperationType = "subordinates"
    HierarchyOperationTypePathToRoot HierarchyOperationType = "path_to_root"
)

type InvoiceDunningHandling string

const (
    InvoiceDunningHandlingContinue InvoiceDunningHandling = "continue"
    InvoiceDunningHandlingStop InvoiceDunningHandling = "stop"
)

type ItemType string

const (
    ItemTypePlan ItemType = "plan"
    ItemTypeAddon ItemType = "addon"
    ItemTypeCharge ItemType = "charge"
)

type Layout string

const (
    LayoutInApp Layout = "in_app"
    LayoutFullPage Layout = "full_page"
)

type NotifyReferralSystem string

const (
    NotifyReferralSystemNone NotifyReferralSystem = "none"
    NotifyReferralSystemFirstPaidConversion NotifyReferralSystem = "first_paid_conversion"
    NotifyReferralSystemAllInvoices NotifyReferralSystem = "all_invoices"
)

type OfflinePaymentMethod string

const (
    OfflinePaymentMethodNoPreference OfflinePaymentMethod = "no_preference"
    OfflinePaymentMethodCash OfflinePaymentMethod = "cash"
    OfflinePaymentMethodCheck OfflinePaymentMethod = "check"
    OfflinePaymentMethodBankTransfer OfflinePaymentMethod = "bank_transfer"
    OfflinePaymentMethodAchCredit OfflinePaymentMethod = "ach_credit"
    OfflinePaymentMethodSepaCredit OfflinePaymentMethod = "sepa_credit"
    OfflinePaymentMethodBoleto OfflinePaymentMethod = "boleto"
    OfflinePaymentMethodUsAutomatedBankTransfer OfflinePaymentMethod = "us_automated_bank_transfer"
    OfflinePaymentMethodEuAutomatedBankTransfer OfflinePaymentMethod = "eu_automated_bank_transfer"
    OfflinePaymentMethodUkAutomatedBankTransfer OfflinePaymentMethod = "uk_automated_bank_transfer"
    OfflinePaymentMethodJpAutomatedBankTransfer OfflinePaymentMethod = "jp_automated_bank_transfer"
    OfflinePaymentMethodMxAutomatedBankTransfer OfflinePaymentMethod = "mx_automated_bank_transfer"
    OfflinePaymentMethodCustom OfflinePaymentMethod = "custom"
)

type OnEvent string

const (
    OnEventSubscriptionCreation OnEvent = "subscription_creation"
    OnEventSubscriptionTrialStart OnEvent = "subscription_trial_start"
    OnEventPlanActivation OnEvent = "plan_activation"
    OnEventSubscriptionActivation OnEvent = "subscription_activation"
    OnEventContractTermination OnEvent = "contract_termination"
)

type Operation string

const (
    OperationCreate Operation = "create"
    OperationUpdate Operation = "update"
    OperationDelete Operation = "delete"
)

type OperationType string

const (
    OperationTypeAdd OperationType = "add"
    OperationTypeRemove OperationType = "remove"
)

type PauseOption string

const (
    PauseOptionImmediately PauseOption = "immediately"
    PauseOptionEndOfTerm PauseOption = "end_of_term"
    PauseOptionSpecificDate PauseOption = "specific_date"
    PauseOptionBillingCycles PauseOption = "billing_cycles"
)

type PaymentInitiator string

const (
    PaymentInitiatorCustomer PaymentInitiator = "customer"
    PaymentInitiatorMerchant PaymentInitiator = "merchant"
)

type PaymentMethod string

const (
    PaymentMethodCash PaymentMethod = "cash"
    PaymentMethodCheck PaymentMethod = "check"
    PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
    PaymentMethodOther PaymentMethod = "other"
    PaymentMethodCustom PaymentMethod = "custom"
    PaymentMethodChargeback PaymentMethod = "chargeback"
    PaymentMethodCard PaymentMethod = "card"
    PaymentMethodAmazonPayments PaymentMethod = "amazon_payments"
    PaymentMethodPaypalExpressCheckout PaymentMethod = "paypal_express_checkout"
    PaymentMethodDirectDebit PaymentMethod = "direct_debit"
    PaymentMethodAlipay PaymentMethod = "alipay"
    PaymentMethodUnionpay PaymentMethod = "unionpay"
    PaymentMethodApplePay PaymentMethod = "apple_pay"
    PaymentMethodWechatPay PaymentMethod = "wechat_pay"
    PaymentMethodAchCredit PaymentMethod = "ach_credit"
    PaymentMethodSepaCredit PaymentMethod = "sepa_credit"
    PaymentMethodIdeal PaymentMethod = "ideal"
    PaymentMethodGooglePay PaymentMethod = "google_pay"
    PaymentMethodSofort PaymentMethod = "sofort"
    PaymentMethodBancontact PaymentMethod = "bancontact"
    PaymentMethodGiropay PaymentMethod = "giropay"
    PaymentMethodDotpay PaymentMethod = "dotpay"
    PaymentMethodUpi PaymentMethod = "upi"
    PaymentMethodNetbankingEmandates PaymentMethod = "netbanking_emandates"
    PaymentMethodBoleto PaymentMethod = "boleto"
    PaymentMethodVenmo PaymentMethod = "venmo"
    PaymentMethodPayTo PaymentMethod = "pay_to"
    PaymentMethodFasterPayments PaymentMethod = "faster_payments"
    PaymentMethodSepaInstantTransfer PaymentMethod = "sepa_instant_transfer"
    PaymentMethodAutomatedBankTransfer PaymentMethod = "automated_bank_transfer"
    PaymentMethodKlarnaPayNow PaymentMethod = "klarna_pay_now"
    PaymentMethodOnlineBankingPoland PaymentMethod = "online_banking_poland"
    PaymentMethodPayconiqByBancontact PaymentMethod = "payconiq_by_bancontact"
)

type PaymentMethodType string

const (
    PaymentMethodTypeCard PaymentMethodType = "card"
    PaymentMethodTypePaypalExpressCheckout PaymentMethodType = "paypal_express_checkout"
    PaymentMethodTypeAmazonPayments PaymentMethodType = "amazon_payments"
    PaymentMethodTypeDirectDebit PaymentMethodType = "direct_debit"
    PaymentMethodTypeGeneric PaymentMethodType = "generic"
    PaymentMethodTypeAlipay PaymentMethodType = "alipay"
    PaymentMethodTypeUnionpay PaymentMethodType = "unionpay"
    PaymentMethodTypeApplePay PaymentMethodType = "apple_pay"
    PaymentMethodTypeWechatPay PaymentMethodType = "wechat_pay"
    PaymentMethodTypeIdeal PaymentMethodType = "ideal"
    PaymentMethodTypeGooglePay PaymentMethodType = "google_pay"
    PaymentMethodTypeSofort PaymentMethodType = "sofort"
    PaymentMethodTypeBancontact PaymentMethodType = "bancontact"
    PaymentMethodTypeGiropay PaymentMethodType = "giropay"
    PaymentMethodTypeDotpay PaymentMethodType = "dotpay"
    PaymentMethodTypeUpi PaymentMethodType = "upi"
    PaymentMethodTypeNetbankingEmandates PaymentMethodType = "netbanking_emandates"
    PaymentMethodTypeVenmo PaymentMethodType = "venmo"
    PaymentMethodTypePayTo PaymentMethodType = "pay_to"
    PaymentMethodTypeFasterPayments PaymentMethodType = "faster_payments"
    PaymentMethodTypeSepaInstantTransfer PaymentMethodType = "sepa_instant_transfer"
    PaymentMethodTypeAutomatedBankTransfer PaymentMethodType = "automated_bank_transfer"
    PaymentMethodTypeKlarnaPayNow PaymentMethodType = "klarna_pay_now"
    PaymentMethodTypeOnlineBankingPoland PaymentMethodType = "online_banking_poland"
    PaymentMethodTypePayconiqByBancontact PaymentMethodType = "payconiq_by_bancontact"
)

type PaymentVoucherType string

const (
    PaymentVoucherTypeBoleto PaymentVoucherType = "boleto"
)

type PeriodUnit string

const (
    PeriodUnitDay PeriodUnit = "day"
    PeriodUnitWeek PeriodUnit = "week"
    PeriodUnitMonth PeriodUnit = "month"
    PeriodUnitYear PeriodUnit = "year"
)

type PriceType string

const (
    PriceTypeTaxExclusive PriceType = "tax_exclusive"
    PriceTypeTaxInclusive PriceType = "tax_inclusive"
)

type PricingModel string

const (
    PricingModelFlatFee PricingModel = "flat_fee"
    PricingModelPerUnit PricingModel = "per_unit"
    PricingModelTiered PricingModel = "tiered"
    PricingModelVolume PricingModel = "volume"
    PricingModelStairstep PricingModel = "stairstep"
)

type PricingType string

const (
    PricingTypePerUnit PricingType = "per_unit"
    PricingTypeFlatFee PricingType = "flat_fee"
    PricingTypePackage PricingType = "package"
)

type ProductCatalogVersion string

const (
    ProductCatalogVersionV1 ProductCatalogVersion = "v1"
    ProductCatalogVersionV2 ProductCatalogVersion = "v2"
)

type ProrationType string

const (
    ProrationTypeFullTerm ProrationType = "full_term"
    ProrationTypePartialTerm ProrationType = "partial_term"
    ProrationTypeNone ProrationType = "none"
)

type ReferralSystem string

const (
    ReferralSystemReferralCandy ReferralSystem = "referral_candy"
    ReferralSystemReferralSaasquatch ReferralSystem = "referral_saasquatch"
    ReferralSystemFriendbuy ReferralSystem = "friendbuy"
)

type ReferrerRewardType string

const (
    ReferrerRewardTypeNone ReferrerRewardType = "none"
    ReferrerRewardTypeReferralDirectReward ReferrerRewardType = "referral_direct_reward"
    ReferrerRewardTypeCustomPromotionalCredit ReferrerRewardType = "custom_promotional_credit"
    ReferrerRewardTypeCustomRevenuePercentBased ReferrerRewardType = "custom_revenue_percent_based"
)

type RefundableCreditsHandling string

const (
    RefundableCreditsHandlingNoAction RefundableCreditsHandling = "no_action"
    RefundableCreditsHandlingScheduleRefund RefundableCreditsHandling = "schedule_refund"
)

type ReportBy string

const (
    ReportByCustomer ReportBy = "customer"
    ReportByInvoice ReportBy = "invoice"
    ReportByProduct ReportBy = "product"
    ReportBySubscription ReportBy = "subscription"
)

type ResumeOption string

const (
    ResumeOptionImmediately ResumeOption = "immediately"
    ResumeOptionSpecificDate ResumeOption = "specific_date"
)

type Role string

const (
    RolePrimary Role = "primary"
    RoleBackup Role = "backup"
    RoleNone Role = "none"
)

type ScheduleType string

const (
    ScheduleTypeImmediate ScheduleType = "immediate"
    ScheduleTypeSpecificDates ScheduleType = "specific_dates"
    ScheduleTypeFixedIntervals ScheduleType = "fixed_intervals"
)

type Source string

const (
    SourceAdminConsole Source = "admin_console"
    SourceApi Source = "api"
    SourceBulkOperation Source = "bulk_operation"
    SourceScheduledJob Source = "scheduled_job"
    SourceHostedPage Source = "hosted_page"
    SourcePortal Source = "portal"
    SourceSystem Source = "system"
    SourceNone Source = "none"
    SourceJsApi Source = "js_api"
    SourceMigration Source = "migration"
    SourceExternalService Source = "external_service"
)

type TaxExemptReason string

const (
    TaxExemptReasonTaxNotConfigured TaxExemptReason = "tax_not_configured"
    TaxExemptReasonRegionNonTaxable TaxExemptReason = "region_non_taxable"
    TaxExemptReasonExport TaxExemptReason = "export"
    TaxExemptReasonCustomerExempt TaxExemptReason = "customer_exempt"
    TaxExemptReasonProductExempt TaxExemptReason = "product_exempt"
    TaxExemptReasonZeroRated TaxExemptReason = "zero_rated"
    TaxExemptReasonReverseCharge TaxExemptReason = "reverse_charge"
    TaxExemptReasonHighValuePhysicalGoods TaxExemptReason = "high_value_physical_goods"
    TaxExemptReasonZeroValueItem TaxExemptReason = "zero_value_item"
    TaxExemptReasonTaxNotConfiguredExternalProvider TaxExemptReason = "tax_not_configured_external_provider"
)

type TaxJurisType string

const (
    TaxJurisTypeCountry TaxJurisType = "country"
    TaxJurisTypeFederal TaxJurisType = "federal"
    TaxJurisTypeState TaxJurisType = "state"
    TaxJurisTypeCounty TaxJurisType = "county"
    TaxJurisTypeCity TaxJurisType = "city"
    TaxJurisTypeSpecial TaxJurisType = "special"
    TaxJurisTypeUnincorporated TaxJurisType = "unincorporated"
    TaxJurisTypeOther TaxJurisType = "other"
)

type TaxOverrideReason string

const (
    TaxOverrideReasonIdExempt TaxOverrideReason = "id_exempt"
    TaxOverrideReasonCustomerExempt TaxOverrideReason = "customer_exempt"
    TaxOverrideReasonExport TaxOverrideReason = "export"
)

type Taxability string

const (
    TaxabilityTaxable Taxability = "taxable"
    TaxabilityExempt Taxability = "exempt"
)

type TaxjarExemptionCategory string

const (
    TaxjarExemptionCategoryWholesale TaxjarExemptionCategory = "wholesale"
    TaxjarExemptionCategoryGovernment TaxjarExemptionCategory = "government"
    TaxjarExemptionCategoryOther TaxjarExemptionCategory = "other"
)

type TrialEndAction string

const (
    TrialEndActionSiteDefault TrialEndAction = "site_default"
    TrialEndActionPlanDefault TrialEndAction = "plan_default"
    TrialEndActionActivateSubscription TrialEndAction = "activate_subscription"
    TrialEndActionCancelSubscription TrialEndAction = "cancel_subscription"
)

type Type string

const (
    TypeCard Type = "card"
    TypePaypalExpressCheckout Type = "paypal_express_checkout"
    TypeAmazonPayments Type = "amazon_payments"
    TypeDirectDebit Type = "direct_debit"
    TypeGeneric Type = "generic"
    TypeAlipay Type = "alipay"
    TypeUnionpay Type = "unionpay"
    TypeApplePay Type = "apple_pay"
    TypeWechatPay Type = "wechat_pay"
    TypeIdeal Type = "ideal"
    TypeGooglePay Type = "google_pay"
    TypeSofort Type = "sofort"
    TypeBancontact Type = "bancontact"
    TypeGiropay Type = "giropay"
    TypeDotpay Type = "dotpay"
    TypeUpi Type = "upi"
    TypeNetbankingEmandates Type = "netbanking_emandates"
    TypeVenmo Type = "venmo"
    TypePayTo Type = "pay_to"
    TypeFasterPayments Type = "faster_payments"
    TypeSepaInstantTransfer Type = "sepa_instant_transfer"
    TypeAutomatedBankTransfer Type = "automated_bank_transfer"
    TypeKlarnaPayNow Type = "klarna_pay_now"
    TypeOnlineBankingPoland Type = "online_banking_poland"
    TypePayconiqByBancontact Type = "payconiq_by_bancontact"
    TypeFreeTrial Type = "free_trial"
    TypePayUpFront Type = "pay_up_front"
    TypePayAsYouGo Type = "pay_as_you_go"
)

type UnbilledChargesHandling string

const (
    UnbilledChargesHandlingNoAction UnbilledChargesHandling = "no_action"
    UnbilledChargesHandlingInvoice UnbilledChargesHandling = "invoice"
)

type UnbilledChargesOption string

const (
    UnbilledChargesOptionInvoice UnbilledChargesOption = "invoice"
    UnbilledChargesOptionDelete UnbilledChargesOption = "delete"
)

type UnpaidInvoicesHandling string

const (
    UnpaidInvoicesHandlingNoAction UnpaidInvoicesHandling = "no_action"
    UnpaidInvoicesHandlingSchedulePaymentCollection UnpaidInvoicesHandling = "schedule_payment_collection"
)

type UsageAccumulationResetFrequency string

const (
    UsageAccumulationResetFrequencyNever UsageAccumulationResetFrequency = "never"
    UsageAccumulationResetFrequencySubscriptionBillingFrequency UsageAccumulationResetFrequency = "subscription_billing_frequency"
)

type ValidationStatus string

const (
    ValidationStatusNotValidated ValidationStatus = "not_validated"
    ValidationStatusValid ValidationStatus = "valid"
    ValidationStatusPartiallyValid ValidationStatus = "partially_valid"
    ValidationStatusInvalid ValidationStatus = "invalid"
)

type VoucherType string

const (
    VoucherTypeBoleto VoucherType = "boleto"
)

