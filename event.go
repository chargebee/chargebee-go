package chargebee

import (
	"encoding/json"
)

type EventSource string

const (
	EventSourceAdminConsole    EventSource = "admin_console"
	EventSourceApi             EventSource = "api"
	EventSourceScheduledJob    EventSource = "scheduled_job"
	EventSourceHostedPage      EventSource = "hosted_page"
	EventSourcePortal          EventSource = "portal"
	EventSourceSystem          EventSource = "system"
	EventSourceNone            EventSource = "none"
	EventSourceJsApi           EventSource = "js_api"
	EventSourceMigration       EventSource = "migration"
	EventSourceBulkOperation   EventSource = "bulk_operation"
	EventSourceExternalService EventSource = "external_service"
)

type EventWebhookStatus string

const (
	EventWebhookStatusNotConfigured EventWebhookStatus = "not_configured"
	EventWebhookStatusScheduled     EventWebhookStatus = "scheduled"
	EventWebhookStatusSucceeded     EventWebhookStatus = "succeeded"
	EventWebhookStatusReScheduled   EventWebhookStatus = "re_scheduled"
	EventWebhookStatusFailed        EventWebhookStatus = "failed"
	EventWebhookStatusSkipped       EventWebhookStatus = "skipped"
	EventWebhookStatusNotApplicable EventWebhookStatus = "not_applicable"
	EventWebhookStatusDisabled      EventWebhookStatus = "disabled"
)

type EventEventType string

const (
	EventEventTypeCouponCreated                                           EventEventType = "coupon_created"
	EventEventTypeCouponUpdated                                           EventEventType = "coupon_updated"
	EventEventTypeCouponDeleted                                           EventEventType = "coupon_deleted"
	EventEventTypeCouponSetCreated                                        EventEventType = "coupon_set_created"
	EventEventTypeCouponSetUpdated                                        EventEventType = "coupon_set_updated"
	EventEventTypeCouponSetDeleted                                        EventEventType = "coupon_set_deleted"
	EventEventTypeCouponCodesAdded                                        EventEventType = "coupon_codes_added"
	EventEventTypeCouponCodesDeleted                                      EventEventType = "coupon_codes_deleted"
	EventEventTypeCouponCodesUpdated                                      EventEventType = "coupon_codes_updated"
	EventEventTypeCustomerCreated                                         EventEventType = "customer_created"
	EventEventTypeCustomerChanged                                         EventEventType = "customer_changed"
	EventEventTypeCustomerDeleted                                         EventEventType = "customer_deleted"
	EventEventTypeCustomerMovedOut                                        EventEventType = "customer_moved_out"
	EventEventTypeCustomerMovedIn                                         EventEventType = "customer_moved_in"
	EventEventTypePromotionalCreditsAdded                                 EventEventType = "promotional_credits_added"
	EventEventTypePromotionalCreditsDeducted                              EventEventType = "promotional_credits_deducted"
	EventEventTypeSubscriptionCreated                                     EventEventType = "subscription_created"
	EventEventTypeSubscriptionCreatedWithBackdating                       EventEventType = "subscription_created_with_backdating"
	EventEventTypeSubscriptionStarted                                     EventEventType = "subscription_started"
	EventEventTypeSubscriptionTrialEndReminder                            EventEventType = "subscription_trial_end_reminder"
	EventEventTypeSubscriptionActivated                                   EventEventType = "subscription_activated"
	EventEventTypeSubscriptionActivatedWithBackdating                     EventEventType = "subscription_activated_with_backdating"
	EventEventTypeSubscriptionChanged                                     EventEventType = "subscription_changed"
	EventEventTypeSubscriptionTrialExtended                               EventEventType = "subscription_trial_extended"
	EventEventTypeMrrUpdated                                              EventEventType = "mrr_updated"
	EventEventTypeSubscriptionChangedWithBackdating                       EventEventType = "subscription_changed_with_backdating"
	EventEventTypeSubscriptionCancellationScheduled                       EventEventType = "subscription_cancellation_scheduled"
	EventEventTypeSubscriptionCancellationReminder                        EventEventType = "subscription_cancellation_reminder"
	EventEventTypeSubscriptionCancelled                                   EventEventType = "subscription_cancelled"
	EventEventTypeSubscriptionCanceledWithBackdating                      EventEventType = "subscription_canceled_with_backdating"
	EventEventTypeSubscriptionReactivated                                 EventEventType = "subscription_reactivated"
	EventEventTypeSubscriptionReactivatedWithBackdating                   EventEventType = "subscription_reactivated_with_backdating"
	EventEventTypeSubscriptionRenewed                                     EventEventType = "subscription_renewed"
	EventEventTypeSubscriptionItemsRenewed                                EventEventType = "subscription_items_renewed"
	EventEventTypeSubscriptionScheduledCancellationRemoved                EventEventType = "subscription_scheduled_cancellation_removed"
	EventEventTypeSubscriptionChangesScheduled                            EventEventType = "subscription_changes_scheduled"
	EventEventTypeSubscriptionScheduledChangesRemoved                     EventEventType = "subscription_scheduled_changes_removed"
	EventEventTypeSubscriptionShippingAddressUpdated                      EventEventType = "subscription_shipping_address_updated"
	EventEventTypeSubscriptionDeleted                                     EventEventType = "subscription_deleted"
	EventEventTypeSubscriptionPaused                                      EventEventType = "subscription_paused"
	EventEventTypeSubscriptionPauseScheduled                              EventEventType = "subscription_pause_scheduled"
	EventEventTypeSubscriptionScheduledPauseRemoved                       EventEventType = "subscription_scheduled_pause_removed"
	EventEventTypeSubscriptionResumed                                     EventEventType = "subscription_resumed"
	EventEventTypeSubscriptionResumptionScheduled                         EventEventType = "subscription_resumption_scheduled"
	EventEventTypeSubscriptionScheduledResumptionRemoved                  EventEventType = "subscription_scheduled_resumption_removed"
	EventEventTypeSubscriptionAdvanceInvoiceScheduleAdded                 EventEventType = "subscription_advance_invoice_schedule_added"
	EventEventTypeSubscriptionAdvanceInvoiceScheduleUpdated               EventEventType = "subscription_advance_invoice_schedule_updated"
	EventEventTypeSubscriptionAdvanceInvoiceScheduleRemoved               EventEventType = "subscription_advance_invoice_schedule_removed"
	EventEventTypePendingInvoiceCreated                                   EventEventType = "pending_invoice_created"
	EventEventTypePendingInvoiceUpdated                                   EventEventType = "pending_invoice_updated"
	EventEventTypeInvoiceGenerated                                        EventEventType = "invoice_generated"
	EventEventTypeInvoiceGeneratedWithBackdating                          EventEventType = "invoice_generated_with_backdating"
	EventEventTypeInvoiceUpdated                                          EventEventType = "invoice_updated"
	EventEventTypeInvoiceDeleted                                          EventEventType = "invoice_deleted"
	EventEventTypeCreditNoteCreated                                       EventEventType = "credit_note_created"
	EventEventTypeCreditNoteCreatedWithBackdating                         EventEventType = "credit_note_created_with_backdating"
	EventEventTypeCreditNoteUpdated                                       EventEventType = "credit_note_updated"
	EventEventTypeCreditNoteDeleted                                       EventEventType = "credit_note_deleted"
	EventEventTypePaymentSchedulesCreated                                 EventEventType = "payment_schedules_created"
	EventEventTypePaymentSchedulesUpdated                                 EventEventType = "payment_schedules_updated"
	EventEventTypePaymentScheduleSchemeCreated                            EventEventType = "payment_schedule_scheme_created"
	EventEventTypePaymentScheduleSchemeDeleted                            EventEventType = "payment_schedule_scheme_deleted"
	EventEventTypeSubscriptionRenewalReminder                             EventEventType = "subscription_renewal_reminder"
	EventEventTypeAddUsagesReminder                                       EventEventType = "add_usages_reminder"
	EventEventTypeTransactionCreated                                      EventEventType = "transaction_created"
	EventEventTypeTransactionUpdated                                      EventEventType = "transaction_updated"
	EventEventTypeTransactionDeleted                                      EventEventType = "transaction_deleted"
	EventEventTypePaymentSucceeded                                        EventEventType = "payment_succeeded"
	EventEventTypePaymentFailed                                           EventEventType = "payment_failed"
	EventEventTypeDunningUpdated                                          EventEventType = "dunning_updated"
	EventEventTypePaymentRefunded                                         EventEventType = "payment_refunded"
	EventEventTypePaymentInitiated                                        EventEventType = "payment_initiated"
	EventEventTypeRefundInitiated                                         EventEventType = "refund_initiated"
	EventEventTypeAuthorizationSucceeded                                  EventEventType = "authorization_succeeded"
	EventEventTypeAuthorizationVoided                                     EventEventType = "authorization_voided"
	EventEventTypeCardAdded                                               EventEventType = "card_added"
	EventEventTypeCardUpdated                                             EventEventType = "card_updated"
	EventEventTypeCardExpiryReminder                                      EventEventType = "card_expiry_reminder"
	EventEventTypeCardExpired                                             EventEventType = "card_expired"
	EventEventTypeCardDeleted                                             EventEventType = "card_deleted"
	EventEventTypePaymentSourceAdded                                      EventEventType = "payment_source_added"
	EventEventTypePaymentSourceUpdated                                    EventEventType = "payment_source_updated"
	EventEventTypePaymentSourceDeleted                                    EventEventType = "payment_source_deleted"
	EventEventTypePaymentSourceExpiring                                   EventEventType = "payment_source_expiring"
	EventEventTypePaymentSourceExpired                                    EventEventType = "payment_source_expired"
	EventEventTypePaymentSourceLocallyDeleted                             EventEventType = "payment_source_locally_deleted"
	EventEventTypeVirtualBankAccountAdded                                 EventEventType = "virtual_bank_account_added"
	EventEventTypeVirtualBankAccountUpdated                               EventEventType = "virtual_bank_account_updated"
	EventEventTypeVirtualBankAccountDeleted                               EventEventType = "virtual_bank_account_deleted"
	EventEventTypeTokenCreated                                            EventEventType = "token_created"
	EventEventTypeTokenConsumed                                           EventEventType = "token_consumed"
	EventEventTypeTokenExpired                                            EventEventType = "token_expired"
	EventEventTypeUnbilledChargesCreated                                  EventEventType = "unbilled_charges_created"
	EventEventTypeUnbilledChargesVoided                                   EventEventType = "unbilled_charges_voided"
	EventEventTypeUnbilledChargesDeleted                                  EventEventType = "unbilled_charges_deleted"
	EventEventTypeUnbilledChargesInvoiced                                 EventEventType = "unbilled_charges_invoiced"
	EventEventTypeOrderCreated                                            EventEventType = "order_created"
	EventEventTypeOrderUpdated                                            EventEventType = "order_updated"
	EventEventTypeOrderCancelled                                          EventEventType = "order_cancelled"
	EventEventTypeOrderDelivered                                          EventEventType = "order_delivered"
	EventEventTypeOrderReturned                                           EventEventType = "order_returned"
	EventEventTypeOrderReadyToProcess                                     EventEventType = "order_ready_to_process"
	EventEventTypeOrderReadyToShip                                        EventEventType = "order_ready_to_ship"
	EventEventTypeOrderDeleted                                            EventEventType = "order_deleted"
	EventEventTypeOrderResent                                             EventEventType = "order_resent"
	EventEventTypeQuoteCreated                                            EventEventType = "quote_created"
	EventEventTypeQuoteUpdated                                            EventEventType = "quote_updated"
	EventEventTypeQuoteDeleted                                            EventEventType = "quote_deleted"
	EventEventTypeTaxWithheldRecorded                                     EventEventType = "tax_withheld_recorded"
	EventEventTypeTaxWithheldDeleted                                      EventEventType = "tax_withheld_deleted"
	EventEventTypeTaxWithheldRefunded                                     EventEventType = "tax_withheld_refunded"
	EventEventTypeGiftScheduled                                           EventEventType = "gift_scheduled"
	EventEventTypeGiftUnclaimed                                           EventEventType = "gift_unclaimed"
	EventEventTypeGiftClaimed                                             EventEventType = "gift_claimed"
	EventEventTypeGiftExpired                                             EventEventType = "gift_expired"
	EventEventTypeGiftCancelled                                           EventEventType = "gift_cancelled"
	EventEventTypeGiftUpdated                                             EventEventType = "gift_updated"
	EventEventTypeHierarchyCreated                                        EventEventType = "hierarchy_created"
	EventEventTypeHierarchyDeleted                                        EventEventType = "hierarchy_deleted"
	EventEventTypePaymentIntentCreated                                    EventEventType = "payment_intent_created"
	EventEventTypePaymentIntentUpdated                                    EventEventType = "payment_intent_updated"
	EventEventTypeContractTermCreated                                     EventEventType = "contract_term_created"
	EventEventTypeContractTermRenewed                                     EventEventType = "contract_term_renewed"
	EventEventTypeContractTermTerminated                                  EventEventType = "contract_term_terminated"
	EventEventTypeContractTermCompleted                                   EventEventType = "contract_term_completed"
	EventEventTypeContractTermCancelled                                   EventEventType = "contract_term_cancelled"
	EventEventTypeItemFamilyCreated                                       EventEventType = "item_family_created"
	EventEventTypeItemFamilyUpdated                                       EventEventType = "item_family_updated"
	EventEventTypeItemFamilyDeleted                                       EventEventType = "item_family_deleted"
	EventEventTypeItemCreated                                             EventEventType = "item_created"
	EventEventTypeItemUpdated                                             EventEventType = "item_updated"
	EventEventTypeItemDeleted                                             EventEventType = "item_deleted"
	EventEventTypeItemPriceCreated                                        EventEventType = "item_price_created"
	EventEventTypeItemPriceUpdated                                        EventEventType = "item_price_updated"
	EventEventTypeItemPriceDeleted                                        EventEventType = "item_price_deleted"
	EventEventTypeAttachedItemCreated                                     EventEventType = "attached_item_created"
	EventEventTypeAttachedItemUpdated                                     EventEventType = "attached_item_updated"
	EventEventTypeAttachedItemDeleted                                     EventEventType = "attached_item_deleted"
	EventEventTypeDifferentialPriceCreated                                EventEventType = "differential_price_created"
	EventEventTypeDifferentialPriceUpdated                                EventEventType = "differential_price_updated"
	EventEventTypeDifferentialPriceDeleted                                EventEventType = "differential_price_deleted"
	EventEventTypeFeatureCreated                                          EventEventType = "feature_created"
	EventEventTypeFeatureUpdated                                          EventEventType = "feature_updated"
	EventEventTypeFeatureDeleted                                          EventEventType = "feature_deleted"
	EventEventTypeFeatureActivated                                        EventEventType = "feature_activated"
	EventEventTypeFeatureReactivated                                      EventEventType = "feature_reactivated"
	EventEventTypeFeatureArchived                                         EventEventType = "feature_archived"
	EventEventTypeItemEntitlementsUpdated                                 EventEventType = "item_entitlements_updated"
	EventEventTypeEntitlementOverridesUpdated                             EventEventType = "entitlement_overrides_updated"
	EventEventTypeEntitlementOverridesRemoved                             EventEventType = "entitlement_overrides_removed"
	EventEventTypeItemEntitlementsRemoved                                 EventEventType = "item_entitlements_removed"
	EventEventTypeEntitlementOverridesAutoRemoved                         EventEventType = "entitlement_overrides_auto_removed"
	EventEventTypeSubscriptionEntitlementsCreated                         EventEventType = "subscription_entitlements_created"
	EventEventTypeSubscriptionEntitlementsUpdated                         EventEventType = "subscription_entitlements_updated"
	EventEventTypeBusinessEntityCreated                                   EventEventType = "business_entity_created"
	EventEventTypeBusinessEntityUpdated                                   EventEventType = "business_entity_updated"
	EventEventTypeBusinessEntityDeleted                                   EventEventType = "business_entity_deleted"
	EventEventTypeCustomerBusinessEntityChanged                           EventEventType = "customer_business_entity_changed"
	EventEventTypeSubscriptionBusinessEntityChanged                       EventEventType = "subscription_business_entity_changed"
	EventEventTypePurchaseCreated                                         EventEventType = "purchase_created"
	EventEventTypeVoucherCreated                                          EventEventType = "voucher_created"
	EventEventTypeVoucherExpired                                          EventEventType = "voucher_expired"
	EventEventTypeVoucherCreateFailed                                     EventEventType = "voucher_create_failed"
	EventEventTypeItemPriceEntitlementsUpdated                            EventEventType = "item_price_entitlements_updated"
	EventEventTypeItemPriceEntitlementsRemoved                            EventEventType = "item_price_entitlements_removed"
	EventEventTypeSubscriptionRampCreated                                 EventEventType = "subscription_ramp_created"
	EventEventTypeSubscriptionRampDeleted                                 EventEventType = "subscription_ramp_deleted"
	EventEventTypeSubscriptionRampApplied                                 EventEventType = "subscription_ramp_applied"
	EventEventTypeSubscriptionRampDrafted                                 EventEventType = "subscription_ramp_drafted"
	EventEventTypeSubscriptionRampUpdated                                 EventEventType = "subscription_ramp_updated"
	EventEventTypePriceVariantCreated                                     EventEventType = "price_variant_created"
	EventEventTypePriceVariantUpdated                                     EventEventType = "price_variant_updated"
	EventEventTypePriceVariantDeleted                                     EventEventType = "price_variant_deleted"
	EventEventTypeCustomerEntitlementsUpdated                             EventEventType = "customer_entitlements_updated"
	EventEventTypeSubscriptionMovedIn                                     EventEventType = "subscription_moved_in"
	EventEventTypeSubscriptionMovedOut                                    EventEventType = "subscription_moved_out"
	EventEventTypeSubscriptionMovementFailed                              EventEventType = "subscription_movement_failed"
	EventEventTypeOmnichannelSubscriptionCreated                          EventEventType = "omnichannel_subscription_created"
	EventEventTypeOmnichannelSubscriptionItemRenewed                      EventEventType = "omnichannel_subscription_item_renewed"
	EventEventTypeOmnichannelSubscriptionItemDowngraded                   EventEventType = "omnichannel_subscription_item_downgraded"
	EventEventTypeOmnichannelSubscriptionItemExpired                      EventEventType = "omnichannel_subscription_item_expired"
	EventEventTypeOmnichannelSubscriptionItemCancellationScheduled        EventEventType = "omnichannel_subscription_item_cancellation_scheduled"
	EventEventTypeOmnichannelSubscriptionItemScheduledCancellationRemoved EventEventType = "omnichannel_subscription_item_scheduled_cancellation_removed"
	EventEventTypeOmnichannelSubscriptionItemResubscribed                 EventEventType = "omnichannel_subscription_item_resubscribed"
	EventEventTypeOmnichannelSubscriptionItemUpgraded                     EventEventType = "omnichannel_subscription_item_upgraded"
	EventEventTypeOmnichannelSubscriptionItemCancelled                    EventEventType = "omnichannel_subscription_item_cancelled"
	EventEventTypeOmnichannelSubscriptionImported                         EventEventType = "omnichannel_subscription_imported"
	EventEventTypeOmnichannelSubscriptionItemGracePeriodStarted           EventEventType = "omnichannel_subscription_item_grace_period_started"
	EventEventTypeOmnichannelSubscriptionItemGracePeriodExpired           EventEventType = "omnichannel_subscription_item_grace_period_expired"
	EventEventTypeOmnichannelSubscriptionItemDunningStarted               EventEventType = "omnichannel_subscription_item_dunning_started"
	EventEventTypeOmnichannelSubscriptionItemDunningExpired               EventEventType = "omnichannel_subscription_item_dunning_expired"
	EventEventTypeRuleCreated                                             EventEventType = "rule_created"
	EventEventTypeRuleUpdated                                             EventEventType = "rule_updated"
	EventEventTypeRuleDeleted                                             EventEventType = "rule_deleted"
	EventEventTypeRecordPurchaseFailed                                    EventEventType = "record_purchase_failed"
	EventEventTypeOmnichannelSubscriptionItemChangeScheduled              EventEventType = "omnichannel_subscription_item_change_scheduled"
	EventEventTypeOmnichannelSubscriptionItemScheduledChangeRemoved       EventEventType = "omnichannel_subscription_item_scheduled_change_removed"
	EventEventTypeOmnichannelSubscriptionItemReactivated                  EventEventType = "omnichannel_subscription_item_reactivated"
	EventEventTypeSalesOrderCreated                                       EventEventType = "sales_order_created"
	EventEventTypeSalesOrderUpdated                                       EventEventType = "sales_order_updated"
	EventEventTypeOmnichannelSubscriptionItemChanged                      EventEventType = "omnichannel_subscription_item_changed"
	EventEventTypeOmnichannelSubscriptionItemPaused                       EventEventType = "omnichannel_subscription_item_paused"
	EventEventTypeOmnichannelSubscriptionItemResumed                      EventEventType = "omnichannel_subscription_item_resumed"
	EventEventTypeOmnichannelOneTimeOrderCreated                          EventEventType = "omnichannel_one_time_order_created"
	EventEventTypeOmnichannelOneTimeOrderItemCancelled                    EventEventType = "omnichannel_one_time_order_item_cancelled"
	EventEventTypeUsageFileIngested                                       EventEventType = "usage_file_ingested"
	EventEventTypeOmnichannelSubscriptionItemPauseScheduled               EventEventType = "omnichannel_subscription_item_pause_scheduled"
	EventEventTypeOmnichannelSubscriptionMovedIn                          EventEventType = "omnichannel_subscription_moved_in"
	EventEventTypeOmnichannelTransactionCreated                           EventEventType = "omnichannel_transaction_created"
	EventEventTypePlanCreated                                             EventEventType = "plan_created"
	EventEventTypePlanUpdated                                             EventEventType = "plan_updated"
	EventEventTypePlanDeleted                                             EventEventType = "plan_deleted"
	EventEventTypeAddonCreated                                            EventEventType = "addon_created"
	EventEventTypeAddonUpdated                                            EventEventType = "addon_updated"
	EventEventTypeAddonDeleted                                            EventEventType = "addon_deleted"
	EventEventTypeNetdPaymentDueReminder                                  EventEventType = "netd_payment_due_reminder"
	EventEventTypeProductCreated                                          EventEventType = "product_created"
	EventEventTypeProductUpdated                                          EventEventType = "product_updated"
	EventEventTypeProductDeleted                                          EventEventType = "product_deleted"
	EventEventTypeVariantCreated                                          EventEventType = "variant_created"
	EventEventTypeVariantUpdated                                          EventEventType = "variant_updated"
	EventEventTypeVariantDeleted                                          EventEventType = "variant_deleted"
	EventEventTypeOmnichannelSubscriptionItemDowngradeScheduled           EventEventType = "omnichannel_subscription_item_downgrade_scheduled"
	EventEventTypeOmnichannelSubscriptionItemScheduledDowngradeRemoved    EventEventType = "omnichannel_subscription_item_scheduled_downgrade_removed"
)

type EventApiVersion string

const (
	EventApiVersionV1 EventApiVersion = "v1"
	EventApiVersionV2 EventApiVersion = "v2"
)

type EventWebhookWebhookStatus string

const (
	EventWebhookWebhookStatusNotConfigured EventWebhookWebhookStatus = "not_configured"
	EventWebhookWebhookStatusScheduled     EventWebhookWebhookStatus = "scheduled"
	EventWebhookWebhookStatusSucceeded     EventWebhookWebhookStatus = "succeeded"
	EventWebhookWebhookStatusReScheduled   EventWebhookWebhookStatus = "re_scheduled"
	EventWebhookWebhookStatusFailed        EventWebhookWebhookStatus = "failed"
	EventWebhookWebhookStatusSkipped       EventWebhookWebhookStatus = "skipped"
	EventWebhookWebhookStatusNotApplicable EventWebhookWebhookStatus = "not_applicable"
	EventWebhookWebhookStatusDisabled      EventWebhookWebhookStatus = "disabled"
)

// just struct
type Event struct {
	Id         string          `json:"id"`
	OccurredAt int64           `json:"occurred_at"`
	Source     EventSource     `json:"source"`
	User       string          `json:"user"`
	Webhooks   []*EventWebhook `json:"webhooks"`
	EventType  EventEventType  `json:"event_type"`
	ApiVersion EventApiVersion `json:"api_version"`
	Content    json.RawMessage `json:"content"`
	OriginUser string          `json:"origin_user"`
	Object     string          `json:"object"`
}

// sub resources
type EventWebhook struct {
	Id            string                    `json:"id"`
	WebhookStatus EventWebhookWebhookStatus `json:"webhook_status"`
	Object        string                    `json:"object"`
}

// operations
// input params
type EventListRequest struct {
	Limit         *int32           `json:"limit,omitempty"`
	Offset        string           `json:"offset,omitempty"`
	Id            *StringFilter    `json:"id,omitempty"`
	WebhookStatus *EnumFilter      `json:"webhook_status,omitempty"`
	EventType     *EnumFilter      `json:"event_type,omitempty"`
	Source        *EnumFilter      `json:"source,omitempty"`
	OccurredAt    *TimestampFilter `json:"occurred_at,omitempty"`
	SortBy        *SortFilter      `json:"sort_by,omitempty"`
	apiRequest    `json:"-" form:"-"`
}

func (r *EventListRequest) payload() any { return r }

// operation sub response
type EventListEventResponse struct {
	Event *Event `json:"event,omitempty"`
}

// operation response
type EventListResponse struct {
	List       []*EventListEventResponse `json:"list,omitempty"`
	NextOffset string                    `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type EventRetrieveResponse struct {
	Event *Event `json:"event,omitempty"`
	apiResponse
}
