package enum

type EventType string

const (
	EventTypeCouponCreated                                           EventType = "coupon_created"
	EventTypeCouponUpdated                                           EventType = "coupon_updated"
	EventTypeCouponDeleted                                           EventType = "coupon_deleted"
	EventTypeCouponSetCreated                                        EventType = "coupon_set_created"
	EventTypeCouponSetUpdated                                        EventType = "coupon_set_updated"
	EventTypeCouponSetDeleted                                        EventType = "coupon_set_deleted"
	EventTypeCouponCodesAdded                                        EventType = "coupon_codes_added"
	EventTypeCouponCodesDeleted                                      EventType = "coupon_codes_deleted"
	EventTypeCouponCodesUpdated                                      EventType = "coupon_codes_updated"
	EventTypeCustomerCreated                                         EventType = "customer_created"
	EventTypeCustomerChanged                                         EventType = "customer_changed"
	EventTypeCustomerDeleted                                         EventType = "customer_deleted"
	EventTypeCustomerMovedOut                                        EventType = "customer_moved_out"
	EventTypeCustomerMovedIn                                         EventType = "customer_moved_in"
	EventTypePromotionalCreditsAdded                                 EventType = "promotional_credits_added"
	EventTypePromotionalCreditsDeducted                              EventType = "promotional_credits_deducted"
	EventTypeSubscriptionCreated                                     EventType = "subscription_created"
	EventTypeSubscriptionCreatedWithBackdating                       EventType = "subscription_created_with_backdating"
	EventTypeSubscriptionStarted                                     EventType = "subscription_started"
	EventTypeSubscriptionTrialEndReminder                            EventType = "subscription_trial_end_reminder"
	EventTypeSubscriptionActivated                                   EventType = "subscription_activated"
	EventTypeSubscriptionActivatedWithBackdating                     EventType = "subscription_activated_with_backdating"
	EventTypeSubscriptionChanged                                     EventType = "subscription_changed"
	EventTypeSubscriptionTrialExtended                               EventType = "subscription_trial_extended"
	EventTypeMrrUpdated                                              EventType = "mrr_updated"
	EventTypeSubscriptionChangedWithBackdating                       EventType = "subscription_changed_with_backdating"
	EventTypeSubscriptionCancellationScheduled                       EventType = "subscription_cancellation_scheduled"
	EventTypeSubscriptionCancellationReminder                        EventType = "subscription_cancellation_reminder"
	EventTypeSubscriptionCancelled                                   EventType = "subscription_cancelled"
	EventTypeSubscriptionCanceledWithBackdating                      EventType = "subscription_canceled_with_backdating"
	EventTypeSubscriptionReactivated                                 EventType = "subscription_reactivated"
	EventTypeSubscriptionReactivatedWithBackdating                   EventType = "subscription_reactivated_with_backdating"
	EventTypeSubscriptionRenewed                                     EventType = "subscription_renewed"
	EventTypeSubscriptionItemsRenewed                                EventType = "subscription_items_renewed"
	EventTypeSubscriptionScheduledCancellationRemoved                EventType = "subscription_scheduled_cancellation_removed"
	EventTypeSubscriptionChangesScheduled                            EventType = "subscription_changes_scheduled"
	EventTypeSubscriptionScheduledChangesRemoved                     EventType = "subscription_scheduled_changes_removed"
	EventTypeSubscriptionShippingAddressUpdated                      EventType = "subscription_shipping_address_updated"
	EventTypeSubscriptionDeleted                                     EventType = "subscription_deleted"
	EventTypeSubscriptionPaused                                      EventType = "subscription_paused"
	EventTypeSubscriptionPauseScheduled                              EventType = "subscription_pause_scheduled"
	EventTypeSubscriptionScheduledPauseRemoved                       EventType = "subscription_scheduled_pause_removed"
	EventTypeSubscriptionResumed                                     EventType = "subscription_resumed"
	EventTypeSubscriptionResumptionScheduled                         EventType = "subscription_resumption_scheduled"
	EventTypeSubscriptionScheduledResumptionRemoved                  EventType = "subscription_scheduled_resumption_removed"
	EventTypeSubscriptionAdvanceInvoiceScheduleAdded                 EventType = "subscription_advance_invoice_schedule_added"
	EventTypeSubscriptionAdvanceInvoiceScheduleUpdated               EventType = "subscription_advance_invoice_schedule_updated"
	EventTypeSubscriptionAdvanceInvoiceScheduleRemoved               EventType = "subscription_advance_invoice_schedule_removed"
	EventTypePendingInvoiceCreated                                   EventType = "pending_invoice_created"
	EventTypePendingInvoiceUpdated                                   EventType = "pending_invoice_updated"
	EventTypeInvoiceGenerated                                        EventType = "invoice_generated"
	EventTypeInvoiceGeneratedWithBackdating                          EventType = "invoice_generated_with_backdating"
	EventTypeInvoiceUpdated                                          EventType = "invoice_updated"
	EventTypeInvoiceDeleted                                          EventType = "invoice_deleted"
	EventTypeCreditNoteCreated                                       EventType = "credit_note_created"
	EventTypeCreditNoteCreatedWithBackdating                         EventType = "credit_note_created_with_backdating"
	EventTypeCreditNoteUpdated                                       EventType = "credit_note_updated"
	EventTypeCreditNoteDeleted                                       EventType = "credit_note_deleted"
	EventTypePaymentSchedulesCreated                                 EventType = "payment_schedules_created"
	EventTypePaymentSchedulesUpdated                                 EventType = "payment_schedules_updated"
	EventTypePaymentScheduleSchemeCreated                            EventType = "payment_schedule_scheme_created"
	EventTypePaymentScheduleSchemeDeleted                            EventType = "payment_schedule_scheme_deleted"
	EventTypeSubscriptionRenewalReminder                             EventType = "subscription_renewal_reminder"
	EventTypeAddUsagesReminder                                       EventType = "add_usages_reminder"
	EventTypeTransactionCreated                                      EventType = "transaction_created"
	EventTypeTransactionUpdated                                      EventType = "transaction_updated"
	EventTypeTransactionDeleted                                      EventType = "transaction_deleted"
	EventTypePaymentSucceeded                                        EventType = "payment_succeeded"
	EventTypePaymentFailed                                           EventType = "payment_failed"
	EventTypePaymentRefunded                                         EventType = "payment_refunded"
	EventTypePaymentInitiated                                        EventType = "payment_initiated"
	EventTypeRefundInitiated                                         EventType = "refund_initiated"
	EventTypeAuthorizationSucceeded                                  EventType = "authorization_succeeded"
	EventTypeAuthorizationVoided                                     EventType = "authorization_voided"
	EventTypeCardAdded                                               EventType = "card_added"
	EventTypeCardUpdated                                             EventType = "card_updated"
	EventTypeCardExpiryReminder                                      EventType = "card_expiry_reminder"
	EventTypeCardExpired                                             EventType = "card_expired"
	EventTypeCardDeleted                                             EventType = "card_deleted"
	EventTypePaymentSourceAdded                                      EventType = "payment_source_added"
	EventTypePaymentSourceUpdated                                    EventType = "payment_source_updated"
	EventTypePaymentSourceDeleted                                    EventType = "payment_source_deleted"
	EventTypePaymentSourceExpiring                                   EventType = "payment_source_expiring"
	EventTypePaymentSourceExpired                                    EventType = "payment_source_expired"
	EventTypePaymentSourceLocallyDeleted                             EventType = "payment_source_locally_deleted"
	EventTypeVirtualBankAccountAdded                                 EventType = "virtual_bank_account_added"
	EventTypeVirtualBankAccountUpdated                               EventType = "virtual_bank_account_updated"
	EventTypeVirtualBankAccountDeleted                               EventType = "virtual_bank_account_deleted"
	EventTypeTokenCreated                                            EventType = "token_created"
	EventTypeTokenConsumed                                           EventType = "token_consumed"
	EventTypeTokenExpired                                            EventType = "token_expired"
	EventTypeUnbilledChargesCreated                                  EventType = "unbilled_charges_created"
	EventTypeUnbilledChargesVoided                                   EventType = "unbilled_charges_voided"
	EventTypeUnbilledChargesDeleted                                  EventType = "unbilled_charges_deleted"
	EventTypeUnbilledChargesInvoiced                                 EventType = "unbilled_charges_invoiced"
	EventTypeOrderCreated                                            EventType = "order_created"
	EventTypeOrderUpdated                                            EventType = "order_updated"
	EventTypeOrderCancelled                                          EventType = "order_cancelled"
	EventTypeOrderDelivered                                          EventType = "order_delivered"
	EventTypeOrderReturned                                           EventType = "order_returned"
	EventTypeOrderReadyToProcess                                     EventType = "order_ready_to_process"
	EventTypeOrderReadyToShip                                        EventType = "order_ready_to_ship"
	EventTypeOrderDeleted                                            EventType = "order_deleted"
	EventTypeOrderResent                                             EventType = "order_resent"
	EventTypeQuoteCreated                                            EventType = "quote_created"
	EventTypeQuoteUpdated                                            EventType = "quote_updated"
	EventTypeQuoteDeleted                                            EventType = "quote_deleted"
	EventTypeTaxWithheldRecorded                                     EventType = "tax_withheld_recorded"
	EventTypeTaxWithheldDeleted                                      EventType = "tax_withheld_deleted"
	EventTypeTaxWithheldRefunded                                     EventType = "tax_withheld_refunded"
	EventTypeGiftScheduled                                           EventType = "gift_scheduled"
	EventTypeGiftUnclaimed                                           EventType = "gift_unclaimed"
	EventTypeGiftClaimed                                             EventType = "gift_claimed"
	EventTypeGiftExpired                                             EventType = "gift_expired"
	EventTypeGiftCancelled                                           EventType = "gift_cancelled"
	EventTypeGiftUpdated                                             EventType = "gift_updated"
	EventTypeHierarchyCreated                                        EventType = "hierarchy_created"
	EventTypeHierarchyDeleted                                        EventType = "hierarchy_deleted"
	EventTypePaymentIntentCreated                                    EventType = "payment_intent_created"
	EventTypePaymentIntentUpdated                                    EventType = "payment_intent_updated"
	EventTypeContractTermCreated                                     EventType = "contract_term_created"
	EventTypeContractTermRenewed                                     EventType = "contract_term_renewed"
	EventTypeContractTermTerminated                                  EventType = "contract_term_terminated"
	EventTypeContractTermCompleted                                   EventType = "contract_term_completed"
	EventTypeContractTermCancelled                                   EventType = "contract_term_cancelled"
	EventTypeItemFamilyCreated                                       EventType = "item_family_created"
	EventTypeItemFamilyUpdated                                       EventType = "item_family_updated"
	EventTypeItemFamilyDeleted                                       EventType = "item_family_deleted"
	EventTypeItemCreated                                             EventType = "item_created"
	EventTypeItemUpdated                                             EventType = "item_updated"
	EventTypeItemDeleted                                             EventType = "item_deleted"
	EventTypeItemPriceCreated                                        EventType = "item_price_created"
	EventTypeItemPriceUpdated                                        EventType = "item_price_updated"
	EventTypeItemPriceDeleted                                        EventType = "item_price_deleted"
	EventTypeAttachedItemCreated                                     EventType = "attached_item_created"
	EventTypeAttachedItemUpdated                                     EventType = "attached_item_updated"
	EventTypeAttachedItemDeleted                                     EventType = "attached_item_deleted"
	EventTypeDifferentialPriceCreated                                EventType = "differential_price_created"
	EventTypeDifferentialPriceUpdated                                EventType = "differential_price_updated"
	EventTypeDifferentialPriceDeleted                                EventType = "differential_price_deleted"
	EventTypeFeatureCreated                                          EventType = "feature_created"
	EventTypeFeatureUpdated                                          EventType = "feature_updated"
	EventTypeFeatureDeleted                                          EventType = "feature_deleted"
	EventTypeFeatureActivated                                        EventType = "feature_activated"
	EventTypeFeatureReactivated                                      EventType = "feature_reactivated"
	EventTypeFeatureArchived                                         EventType = "feature_archived"
	EventTypeItemEntitlementsUpdated                                 EventType = "item_entitlements_updated"
	EventTypeEntitlementOverridesUpdated                             EventType = "entitlement_overrides_updated"
	EventTypeEntitlementOverridesRemoved                             EventType = "entitlement_overrides_removed"
	EventTypeItemEntitlementsRemoved                                 EventType = "item_entitlements_removed"
	EventTypeEntitlementOverridesAutoRemoved                         EventType = "entitlement_overrides_auto_removed"
	EventTypeSubscriptionEntitlementsCreated                         EventType = "subscription_entitlements_created"
	EventTypeBusinessEntityCreated                                   EventType = "business_entity_created"
	EventTypeBusinessEntityUpdated                                   EventType = "business_entity_updated"
	EventTypeBusinessEntityDeleted                                   EventType = "business_entity_deleted"
	EventTypeCustomerBusinessEntityChanged                           EventType = "customer_business_entity_changed"
	EventTypeSubscriptionBusinessEntityChanged                       EventType = "subscription_business_entity_changed"
	EventTypePurchaseCreated                                         EventType = "purchase_created"
	EventTypeVoucherCreated                                          EventType = "voucher_created"
	EventTypeVoucherExpired                                          EventType = "voucher_expired"
	EventTypeVoucherCreateFailed                                     EventType = "voucher_create_failed"
	EventTypeItemPriceEntitlementsUpdated                            EventType = "item_price_entitlements_updated"
	EventTypeItemPriceEntitlementsRemoved                            EventType = "item_price_entitlements_removed"
	EventTypeSubscriptionRampCreated                                 EventType = "subscription_ramp_created"
	EventTypeSubscriptionRampDeleted                                 EventType = "subscription_ramp_deleted"
	EventTypeSubscriptionRampApplied                                 EventType = "subscription_ramp_applied"
	EventTypeSubscriptionRampDrafted                                 EventType = "subscription_ramp_drafted"
	EventTypeSubscriptionRampUpdated                                 EventType = "subscription_ramp_updated"
	EventTypePriceVariantCreated                                     EventType = "price_variant_created"
	EventTypePriceVariantUpdated                                     EventType = "price_variant_updated"
	EventTypePriceVariantDeleted                                     EventType = "price_variant_deleted"
	EventTypeCustomerEntitlementsUpdated                             EventType = "customer_entitlements_updated"
	EventTypeSubscriptionMovedIn                                     EventType = "subscription_moved_in"
	EventTypeSubscriptionMovedOut                                    EventType = "subscription_moved_out"
	EventTypeSubscriptionMovementFailed                              EventType = "subscription_movement_failed"
	EventTypeOmnichannelSubscriptionCreated                          EventType = "omnichannel_subscription_created"
	EventTypeOmnichannelSubscriptionItemRenewed                      EventType = "omnichannel_subscription_item_renewed"
	EventTypeOmnichannelSubscriptionItemDowngradeScheduled           EventType = "omnichannel_subscription_item_downgrade_scheduled"
	EventTypeOmnichannelSubscriptionItemScheduledDowngradeRemoved    EventType = "omnichannel_subscription_item_scheduled_downgrade_removed"
	EventTypeOmnichannelSubscriptionItemDowngraded                   EventType = "omnichannel_subscription_item_downgraded"
	EventTypeOmnichannelSubscriptionItemExpired                      EventType = "omnichannel_subscription_item_expired"
	EventTypeOmnichannelSubscriptionItemCancellationScheduled        EventType = "omnichannel_subscription_item_cancellation_scheduled"
	EventTypeOmnichannelSubscriptionItemScheduledCancellationRemoved EventType = "omnichannel_subscription_item_scheduled_cancellation_removed"
	EventTypeOmnichannelSubscriptionItemResubscribed                 EventType = "omnichannel_subscription_item_resubscribed"
	EventTypeOmnichannelSubscriptionItemUpgraded                     EventType = "omnichannel_subscription_item_upgraded"
	EventTypeOmnichannelSubscriptionItemCancelled                    EventType = "omnichannel_subscription_item_cancelled"
	EventTypePlanCreated                                             EventType = "plan_created"
	EventTypePlanUpdated                                             EventType = "plan_updated"
	EventTypePlanDeleted                                             EventType = "plan_deleted"
	EventTypeAddonCreated                                            EventType = "addon_created"
	EventTypeAddonUpdated                                            EventType = "addon_updated"
	EventTypeAddonDeleted                                            EventType = "addon_deleted"
)
