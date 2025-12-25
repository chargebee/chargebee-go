package chargebee

type WebhookEndpointApiVersion string

const (
	WebhookEndpointApiVersionV1 WebhookEndpointApiVersion = "v1"
	WebhookEndpointApiVersionV2 WebhookEndpointApiVersion = "v2"
)

type WebhookEndpointChargebeeResponseSchemaType string

const (
	WebhookEndpointChargebeeResponseSchemaTypePlansAddons WebhookEndpointChargebeeResponseSchemaType = "plans_addons"
	WebhookEndpointChargebeeResponseSchemaTypeItems       WebhookEndpointChargebeeResponseSchemaType = "items"
	WebhookEndpointChargebeeResponseSchemaTypeCompat      WebhookEndpointChargebeeResponseSchemaType = "compat"
)

type WebhookEndpointEnabledEvents string

const (
	WebhookEndpointEnabledEventsCouponCreated                                           WebhookEndpointEnabledEvents = "coupon_created"
	WebhookEndpointEnabledEventsCouponUpdated                                           WebhookEndpointEnabledEvents = "coupon_updated"
	WebhookEndpointEnabledEventsCouponDeleted                                           WebhookEndpointEnabledEvents = "coupon_deleted"
	WebhookEndpointEnabledEventsCouponSetCreated                                        WebhookEndpointEnabledEvents = "coupon_set_created"
	WebhookEndpointEnabledEventsCouponSetUpdated                                        WebhookEndpointEnabledEvents = "coupon_set_updated"
	WebhookEndpointEnabledEventsCouponSetDeleted                                        WebhookEndpointEnabledEvents = "coupon_set_deleted"
	WebhookEndpointEnabledEventsCouponCodesAdded                                        WebhookEndpointEnabledEvents = "coupon_codes_added"
	WebhookEndpointEnabledEventsCouponCodesDeleted                                      WebhookEndpointEnabledEvents = "coupon_codes_deleted"
	WebhookEndpointEnabledEventsCouponCodesUpdated                                      WebhookEndpointEnabledEvents = "coupon_codes_updated"
	WebhookEndpointEnabledEventsCustomerCreated                                         WebhookEndpointEnabledEvents = "customer_created"
	WebhookEndpointEnabledEventsCustomerChanged                                         WebhookEndpointEnabledEvents = "customer_changed"
	WebhookEndpointEnabledEventsCustomerDeleted                                         WebhookEndpointEnabledEvents = "customer_deleted"
	WebhookEndpointEnabledEventsCustomerMovedOut                                        WebhookEndpointEnabledEvents = "customer_moved_out"
	WebhookEndpointEnabledEventsCustomerMovedIn                                         WebhookEndpointEnabledEvents = "customer_moved_in"
	WebhookEndpointEnabledEventsPromotionalCreditsAdded                                 WebhookEndpointEnabledEvents = "promotional_credits_added"
	WebhookEndpointEnabledEventsPromotionalCreditsDeducted                              WebhookEndpointEnabledEvents = "promotional_credits_deducted"
	WebhookEndpointEnabledEventsSubscriptionCreated                                     WebhookEndpointEnabledEvents = "subscription_created"
	WebhookEndpointEnabledEventsSubscriptionCreatedWithBackdating                       WebhookEndpointEnabledEvents = "subscription_created_with_backdating"
	WebhookEndpointEnabledEventsSubscriptionStarted                                     WebhookEndpointEnabledEvents = "subscription_started"
	WebhookEndpointEnabledEventsSubscriptionTrialEndReminder                            WebhookEndpointEnabledEvents = "subscription_trial_end_reminder"
	WebhookEndpointEnabledEventsSubscriptionActivated                                   WebhookEndpointEnabledEvents = "subscription_activated"
	WebhookEndpointEnabledEventsSubscriptionActivatedWithBackdating                     WebhookEndpointEnabledEvents = "subscription_activated_with_backdating"
	WebhookEndpointEnabledEventsSubscriptionChanged                                     WebhookEndpointEnabledEvents = "subscription_changed"
	WebhookEndpointEnabledEventsSubscriptionTrialExtended                               WebhookEndpointEnabledEvents = "subscription_trial_extended"
	WebhookEndpointEnabledEventsMrrUpdated                                              WebhookEndpointEnabledEvents = "mrr_updated"
	WebhookEndpointEnabledEventsSubscriptionChangedWithBackdating                       WebhookEndpointEnabledEvents = "subscription_changed_with_backdating"
	WebhookEndpointEnabledEventsSubscriptionCancellationScheduled                       WebhookEndpointEnabledEvents = "subscription_cancellation_scheduled"
	WebhookEndpointEnabledEventsSubscriptionCancellationReminder                        WebhookEndpointEnabledEvents = "subscription_cancellation_reminder"
	WebhookEndpointEnabledEventsSubscriptionCancelled                                   WebhookEndpointEnabledEvents = "subscription_cancelled"
	WebhookEndpointEnabledEventsSubscriptionCanceledWithBackdating                      WebhookEndpointEnabledEvents = "subscription_canceled_with_backdating"
	WebhookEndpointEnabledEventsSubscriptionReactivated                                 WebhookEndpointEnabledEvents = "subscription_reactivated"
	WebhookEndpointEnabledEventsSubscriptionReactivatedWithBackdating                   WebhookEndpointEnabledEvents = "subscription_reactivated_with_backdating"
	WebhookEndpointEnabledEventsSubscriptionRenewed                                     WebhookEndpointEnabledEvents = "subscription_renewed"
	WebhookEndpointEnabledEventsSubscriptionItemsRenewed                                WebhookEndpointEnabledEvents = "subscription_items_renewed"
	WebhookEndpointEnabledEventsSubscriptionScheduledCancellationRemoved                WebhookEndpointEnabledEvents = "subscription_scheduled_cancellation_removed"
	WebhookEndpointEnabledEventsSubscriptionChangesScheduled                            WebhookEndpointEnabledEvents = "subscription_changes_scheduled"
	WebhookEndpointEnabledEventsSubscriptionScheduledChangesRemoved                     WebhookEndpointEnabledEvents = "subscription_scheduled_changes_removed"
	WebhookEndpointEnabledEventsSubscriptionShippingAddressUpdated                      WebhookEndpointEnabledEvents = "subscription_shipping_address_updated"
	WebhookEndpointEnabledEventsSubscriptionDeleted                                     WebhookEndpointEnabledEvents = "subscription_deleted"
	WebhookEndpointEnabledEventsSubscriptionPaused                                      WebhookEndpointEnabledEvents = "subscription_paused"
	WebhookEndpointEnabledEventsSubscriptionPauseScheduled                              WebhookEndpointEnabledEvents = "subscription_pause_scheduled"
	WebhookEndpointEnabledEventsSubscriptionScheduledPauseRemoved                       WebhookEndpointEnabledEvents = "subscription_scheduled_pause_removed"
	WebhookEndpointEnabledEventsSubscriptionResumed                                     WebhookEndpointEnabledEvents = "subscription_resumed"
	WebhookEndpointEnabledEventsSubscriptionResumptionScheduled                         WebhookEndpointEnabledEvents = "subscription_resumption_scheduled"
	WebhookEndpointEnabledEventsSubscriptionScheduledResumptionRemoved                  WebhookEndpointEnabledEvents = "subscription_scheduled_resumption_removed"
	WebhookEndpointEnabledEventsSubscriptionAdvanceInvoiceScheduleAdded                 WebhookEndpointEnabledEvents = "subscription_advance_invoice_schedule_added"
	WebhookEndpointEnabledEventsSubscriptionAdvanceInvoiceScheduleUpdated               WebhookEndpointEnabledEvents = "subscription_advance_invoice_schedule_updated"
	WebhookEndpointEnabledEventsSubscriptionAdvanceInvoiceScheduleRemoved               WebhookEndpointEnabledEvents = "subscription_advance_invoice_schedule_removed"
	WebhookEndpointEnabledEventsPendingInvoiceCreated                                   WebhookEndpointEnabledEvents = "pending_invoice_created"
	WebhookEndpointEnabledEventsPendingInvoiceUpdated                                   WebhookEndpointEnabledEvents = "pending_invoice_updated"
	WebhookEndpointEnabledEventsInvoiceGenerated                                        WebhookEndpointEnabledEvents = "invoice_generated"
	WebhookEndpointEnabledEventsInvoiceGeneratedWithBackdating                          WebhookEndpointEnabledEvents = "invoice_generated_with_backdating"
	WebhookEndpointEnabledEventsInvoiceUpdated                                          WebhookEndpointEnabledEvents = "invoice_updated"
	WebhookEndpointEnabledEventsInvoiceDeleted                                          WebhookEndpointEnabledEvents = "invoice_deleted"
	WebhookEndpointEnabledEventsCreditNoteCreated                                       WebhookEndpointEnabledEvents = "credit_note_created"
	WebhookEndpointEnabledEventsCreditNoteCreatedWithBackdating                         WebhookEndpointEnabledEvents = "credit_note_created_with_backdating"
	WebhookEndpointEnabledEventsCreditNoteUpdated                                       WebhookEndpointEnabledEvents = "credit_note_updated"
	WebhookEndpointEnabledEventsCreditNoteDeleted                                       WebhookEndpointEnabledEvents = "credit_note_deleted"
	WebhookEndpointEnabledEventsPaymentSchedulesCreated                                 WebhookEndpointEnabledEvents = "payment_schedules_created"
	WebhookEndpointEnabledEventsPaymentSchedulesUpdated                                 WebhookEndpointEnabledEvents = "payment_schedules_updated"
	WebhookEndpointEnabledEventsPaymentScheduleSchemeCreated                            WebhookEndpointEnabledEvents = "payment_schedule_scheme_created"
	WebhookEndpointEnabledEventsPaymentScheduleSchemeDeleted                            WebhookEndpointEnabledEvents = "payment_schedule_scheme_deleted"
	WebhookEndpointEnabledEventsSubscriptionRenewalReminder                             WebhookEndpointEnabledEvents = "subscription_renewal_reminder"
	WebhookEndpointEnabledEventsAddUsagesReminder                                       WebhookEndpointEnabledEvents = "add_usages_reminder"
	WebhookEndpointEnabledEventsTransactionCreated                                      WebhookEndpointEnabledEvents = "transaction_created"
	WebhookEndpointEnabledEventsTransactionUpdated                                      WebhookEndpointEnabledEvents = "transaction_updated"
	WebhookEndpointEnabledEventsTransactionDeleted                                      WebhookEndpointEnabledEvents = "transaction_deleted"
	WebhookEndpointEnabledEventsPaymentSucceeded                                        WebhookEndpointEnabledEvents = "payment_succeeded"
	WebhookEndpointEnabledEventsPaymentFailed                                           WebhookEndpointEnabledEvents = "payment_failed"
	WebhookEndpointEnabledEventsDunningUpdated                                          WebhookEndpointEnabledEvents = "dunning_updated"
	WebhookEndpointEnabledEventsPaymentRefunded                                         WebhookEndpointEnabledEvents = "payment_refunded"
	WebhookEndpointEnabledEventsPaymentInitiated                                        WebhookEndpointEnabledEvents = "payment_initiated"
	WebhookEndpointEnabledEventsRefundInitiated                                         WebhookEndpointEnabledEvents = "refund_initiated"
	WebhookEndpointEnabledEventsNetdPaymentDueReminder                                  WebhookEndpointEnabledEvents = "netd_payment_due_reminder"
	WebhookEndpointEnabledEventsAuthorizationSucceeded                                  WebhookEndpointEnabledEvents = "authorization_succeeded"
	WebhookEndpointEnabledEventsAuthorizationVoided                                     WebhookEndpointEnabledEvents = "authorization_voided"
	WebhookEndpointEnabledEventsCardAdded                                               WebhookEndpointEnabledEvents = "card_added"
	WebhookEndpointEnabledEventsCardUpdated                                             WebhookEndpointEnabledEvents = "card_updated"
	WebhookEndpointEnabledEventsCardExpiryReminder                                      WebhookEndpointEnabledEvents = "card_expiry_reminder"
	WebhookEndpointEnabledEventsCardExpired                                             WebhookEndpointEnabledEvents = "card_expired"
	WebhookEndpointEnabledEventsCardDeleted                                             WebhookEndpointEnabledEvents = "card_deleted"
	WebhookEndpointEnabledEventsPaymentSourceAdded                                      WebhookEndpointEnabledEvents = "payment_source_added"
	WebhookEndpointEnabledEventsPaymentSourceUpdated                                    WebhookEndpointEnabledEvents = "payment_source_updated"
	WebhookEndpointEnabledEventsPaymentSourceDeleted                                    WebhookEndpointEnabledEvents = "payment_source_deleted"
	WebhookEndpointEnabledEventsPaymentSourceExpiring                                   WebhookEndpointEnabledEvents = "payment_source_expiring"
	WebhookEndpointEnabledEventsPaymentSourceExpired                                    WebhookEndpointEnabledEvents = "payment_source_expired"
	WebhookEndpointEnabledEventsPaymentSourceLocallyDeleted                             WebhookEndpointEnabledEvents = "payment_source_locally_deleted"
	WebhookEndpointEnabledEventsVirtualBankAccountAdded                                 WebhookEndpointEnabledEvents = "virtual_bank_account_added"
	WebhookEndpointEnabledEventsVirtualBankAccountUpdated                               WebhookEndpointEnabledEvents = "virtual_bank_account_updated"
	WebhookEndpointEnabledEventsVirtualBankAccountDeleted                               WebhookEndpointEnabledEvents = "virtual_bank_account_deleted"
	WebhookEndpointEnabledEventsTokenCreated                                            WebhookEndpointEnabledEvents = "token_created"
	WebhookEndpointEnabledEventsTokenConsumed                                           WebhookEndpointEnabledEvents = "token_consumed"
	WebhookEndpointEnabledEventsTokenExpired                                            WebhookEndpointEnabledEvents = "token_expired"
	WebhookEndpointEnabledEventsUnbilledChargesCreated                                  WebhookEndpointEnabledEvents = "unbilled_charges_created"
	WebhookEndpointEnabledEventsUnbilledChargesVoided                                   WebhookEndpointEnabledEvents = "unbilled_charges_voided"
	WebhookEndpointEnabledEventsUnbilledChargesDeleted                                  WebhookEndpointEnabledEvents = "unbilled_charges_deleted"
	WebhookEndpointEnabledEventsUnbilledChargesInvoiced                                 WebhookEndpointEnabledEvents = "unbilled_charges_invoiced"
	WebhookEndpointEnabledEventsOrderCreated                                            WebhookEndpointEnabledEvents = "order_created"
	WebhookEndpointEnabledEventsOrderUpdated                                            WebhookEndpointEnabledEvents = "order_updated"
	WebhookEndpointEnabledEventsOrderCancelled                                          WebhookEndpointEnabledEvents = "order_cancelled"
	WebhookEndpointEnabledEventsOrderDelivered                                          WebhookEndpointEnabledEvents = "order_delivered"
	WebhookEndpointEnabledEventsOrderReturned                                           WebhookEndpointEnabledEvents = "order_returned"
	WebhookEndpointEnabledEventsOrderReadyToProcess                                     WebhookEndpointEnabledEvents = "order_ready_to_process"
	WebhookEndpointEnabledEventsOrderReadyToShip                                        WebhookEndpointEnabledEvents = "order_ready_to_ship"
	WebhookEndpointEnabledEventsOrderDeleted                                            WebhookEndpointEnabledEvents = "order_deleted"
	WebhookEndpointEnabledEventsOrderResent                                             WebhookEndpointEnabledEvents = "order_resent"
	WebhookEndpointEnabledEventsQuoteCreated                                            WebhookEndpointEnabledEvents = "quote_created"
	WebhookEndpointEnabledEventsQuoteUpdated                                            WebhookEndpointEnabledEvents = "quote_updated"
	WebhookEndpointEnabledEventsQuoteDeleted                                            WebhookEndpointEnabledEvents = "quote_deleted"
	WebhookEndpointEnabledEventsTaxWithheldRecorded                                     WebhookEndpointEnabledEvents = "tax_withheld_recorded"
	WebhookEndpointEnabledEventsTaxWithheldDeleted                                      WebhookEndpointEnabledEvents = "tax_withheld_deleted"
	WebhookEndpointEnabledEventsTaxWithheldRefunded                                     WebhookEndpointEnabledEvents = "tax_withheld_refunded"
	WebhookEndpointEnabledEventsGiftScheduled                                           WebhookEndpointEnabledEvents = "gift_scheduled"
	WebhookEndpointEnabledEventsGiftUnclaimed                                           WebhookEndpointEnabledEvents = "gift_unclaimed"
	WebhookEndpointEnabledEventsGiftClaimed                                             WebhookEndpointEnabledEvents = "gift_claimed"
	WebhookEndpointEnabledEventsGiftExpired                                             WebhookEndpointEnabledEvents = "gift_expired"
	WebhookEndpointEnabledEventsGiftCancelled                                           WebhookEndpointEnabledEvents = "gift_cancelled"
	WebhookEndpointEnabledEventsGiftUpdated                                             WebhookEndpointEnabledEvents = "gift_updated"
	WebhookEndpointEnabledEventsHierarchyCreated                                        WebhookEndpointEnabledEvents = "hierarchy_created"
	WebhookEndpointEnabledEventsHierarchyDeleted                                        WebhookEndpointEnabledEvents = "hierarchy_deleted"
	WebhookEndpointEnabledEventsPaymentIntentCreated                                    WebhookEndpointEnabledEvents = "payment_intent_created"
	WebhookEndpointEnabledEventsPaymentIntentUpdated                                    WebhookEndpointEnabledEvents = "payment_intent_updated"
	WebhookEndpointEnabledEventsContractTermCreated                                     WebhookEndpointEnabledEvents = "contract_term_created"
	WebhookEndpointEnabledEventsContractTermRenewed                                     WebhookEndpointEnabledEvents = "contract_term_renewed"
	WebhookEndpointEnabledEventsContractTermTerminated                                  WebhookEndpointEnabledEvents = "contract_term_terminated"
	WebhookEndpointEnabledEventsContractTermCompleted                                   WebhookEndpointEnabledEvents = "contract_term_completed"
	WebhookEndpointEnabledEventsContractTermCancelled                                   WebhookEndpointEnabledEvents = "contract_term_cancelled"
	WebhookEndpointEnabledEventsItemFamilyCreated                                       WebhookEndpointEnabledEvents = "item_family_created"
	WebhookEndpointEnabledEventsItemFamilyUpdated                                       WebhookEndpointEnabledEvents = "item_family_updated"
	WebhookEndpointEnabledEventsItemFamilyDeleted                                       WebhookEndpointEnabledEvents = "item_family_deleted"
	WebhookEndpointEnabledEventsItemCreated                                             WebhookEndpointEnabledEvents = "item_created"
	WebhookEndpointEnabledEventsItemUpdated                                             WebhookEndpointEnabledEvents = "item_updated"
	WebhookEndpointEnabledEventsItemDeleted                                             WebhookEndpointEnabledEvents = "item_deleted"
	WebhookEndpointEnabledEventsItemPriceCreated                                        WebhookEndpointEnabledEvents = "item_price_created"
	WebhookEndpointEnabledEventsItemPriceUpdated                                        WebhookEndpointEnabledEvents = "item_price_updated"
	WebhookEndpointEnabledEventsItemPriceDeleted                                        WebhookEndpointEnabledEvents = "item_price_deleted"
	WebhookEndpointEnabledEventsAttachedItemCreated                                     WebhookEndpointEnabledEvents = "attached_item_created"
	WebhookEndpointEnabledEventsAttachedItemUpdated                                     WebhookEndpointEnabledEvents = "attached_item_updated"
	WebhookEndpointEnabledEventsAttachedItemDeleted                                     WebhookEndpointEnabledEvents = "attached_item_deleted"
	WebhookEndpointEnabledEventsDifferentialPriceCreated                                WebhookEndpointEnabledEvents = "differential_price_created"
	WebhookEndpointEnabledEventsDifferentialPriceUpdated                                WebhookEndpointEnabledEvents = "differential_price_updated"
	WebhookEndpointEnabledEventsDifferentialPriceDeleted                                WebhookEndpointEnabledEvents = "differential_price_deleted"
	WebhookEndpointEnabledEventsFeatureCreated                                          WebhookEndpointEnabledEvents = "feature_created"
	WebhookEndpointEnabledEventsFeatureUpdated                                          WebhookEndpointEnabledEvents = "feature_updated"
	WebhookEndpointEnabledEventsFeatureDeleted                                          WebhookEndpointEnabledEvents = "feature_deleted"
	WebhookEndpointEnabledEventsFeatureActivated                                        WebhookEndpointEnabledEvents = "feature_activated"
	WebhookEndpointEnabledEventsFeatureReactivated                                      WebhookEndpointEnabledEvents = "feature_reactivated"
	WebhookEndpointEnabledEventsFeatureArchived                                         WebhookEndpointEnabledEvents = "feature_archived"
	WebhookEndpointEnabledEventsItemEntitlementsUpdated                                 WebhookEndpointEnabledEvents = "item_entitlements_updated"
	WebhookEndpointEnabledEventsEntitlementOverridesUpdated                             WebhookEndpointEnabledEvents = "entitlement_overrides_updated"
	WebhookEndpointEnabledEventsEntitlementOverridesRemoved                             WebhookEndpointEnabledEvents = "entitlement_overrides_removed"
	WebhookEndpointEnabledEventsItemEntitlementsRemoved                                 WebhookEndpointEnabledEvents = "item_entitlements_removed"
	WebhookEndpointEnabledEventsEntitlementOverridesAutoRemoved                         WebhookEndpointEnabledEvents = "entitlement_overrides_auto_removed"
	WebhookEndpointEnabledEventsSubscriptionEntitlementsCreated                         WebhookEndpointEnabledEvents = "subscription_entitlements_created"
	WebhookEndpointEnabledEventsSubscriptionEntitlementsUpdated                         WebhookEndpointEnabledEvents = "subscription_entitlements_updated"
	WebhookEndpointEnabledEventsBusinessEntityCreated                                   WebhookEndpointEnabledEvents = "business_entity_created"
	WebhookEndpointEnabledEventsBusinessEntityUpdated                                   WebhookEndpointEnabledEvents = "business_entity_updated"
	WebhookEndpointEnabledEventsBusinessEntityDeleted                                   WebhookEndpointEnabledEvents = "business_entity_deleted"
	WebhookEndpointEnabledEventsCustomerBusinessEntityChanged                           WebhookEndpointEnabledEvents = "customer_business_entity_changed"
	WebhookEndpointEnabledEventsSubscriptionBusinessEntityChanged                       WebhookEndpointEnabledEvents = "subscription_business_entity_changed"
	WebhookEndpointEnabledEventsPurchaseCreated                                         WebhookEndpointEnabledEvents = "purchase_created"
	WebhookEndpointEnabledEventsVoucherCreated                                          WebhookEndpointEnabledEvents = "voucher_created"
	WebhookEndpointEnabledEventsVoucherExpired                                          WebhookEndpointEnabledEvents = "voucher_expired"
	WebhookEndpointEnabledEventsVoucherCreateFailed                                     WebhookEndpointEnabledEvents = "voucher_create_failed"
	WebhookEndpointEnabledEventsProductCreated                                          WebhookEndpointEnabledEvents = "product_created"
	WebhookEndpointEnabledEventsProductUpdated                                          WebhookEndpointEnabledEvents = "product_updated"
	WebhookEndpointEnabledEventsProductDeleted                                          WebhookEndpointEnabledEvents = "product_deleted"
	WebhookEndpointEnabledEventsVariantCreated                                          WebhookEndpointEnabledEvents = "variant_created"
	WebhookEndpointEnabledEventsVariantUpdated                                          WebhookEndpointEnabledEvents = "variant_updated"
	WebhookEndpointEnabledEventsVariantDeleted                                          WebhookEndpointEnabledEvents = "variant_deleted"
	WebhookEndpointEnabledEventsItemPriceEntitlementsUpdated                            WebhookEndpointEnabledEvents = "item_price_entitlements_updated"
	WebhookEndpointEnabledEventsItemPriceEntitlementsRemoved                            WebhookEndpointEnabledEvents = "item_price_entitlements_removed"
	WebhookEndpointEnabledEventsSubscriptionRampCreated                                 WebhookEndpointEnabledEvents = "subscription_ramp_created"
	WebhookEndpointEnabledEventsSubscriptionRampDeleted                                 WebhookEndpointEnabledEvents = "subscription_ramp_deleted"
	WebhookEndpointEnabledEventsSubscriptionRampApplied                                 WebhookEndpointEnabledEvents = "subscription_ramp_applied"
	WebhookEndpointEnabledEventsSubscriptionRampDrafted                                 WebhookEndpointEnabledEvents = "subscription_ramp_drafted"
	WebhookEndpointEnabledEventsSubscriptionRampUpdated                                 WebhookEndpointEnabledEvents = "subscription_ramp_updated"
	WebhookEndpointEnabledEventsPriceVariantCreated                                     WebhookEndpointEnabledEvents = "price_variant_created"
	WebhookEndpointEnabledEventsPriceVariantUpdated                                     WebhookEndpointEnabledEvents = "price_variant_updated"
	WebhookEndpointEnabledEventsPriceVariantDeleted                                     WebhookEndpointEnabledEvents = "price_variant_deleted"
	WebhookEndpointEnabledEventsCustomerEntitlementsUpdated                             WebhookEndpointEnabledEvents = "customer_entitlements_updated"
	WebhookEndpointEnabledEventsSubscriptionMovedIn                                     WebhookEndpointEnabledEvents = "subscription_moved_in"
	WebhookEndpointEnabledEventsSubscriptionMovedOut                                    WebhookEndpointEnabledEvents = "subscription_moved_out"
	WebhookEndpointEnabledEventsSubscriptionMovementFailed                              WebhookEndpointEnabledEvents = "subscription_movement_failed"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionCreated                          WebhookEndpointEnabledEvents = "omnichannel_subscription_created"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemRenewed                      WebhookEndpointEnabledEvents = "omnichannel_subscription_item_renewed"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemDowngradeScheduled           WebhookEndpointEnabledEvents = "omnichannel_subscription_item_downgrade_scheduled"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemScheduledDowngradeRemoved    WebhookEndpointEnabledEvents = "omnichannel_subscription_item_scheduled_downgrade_removed"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemDowngraded                   WebhookEndpointEnabledEvents = "omnichannel_subscription_item_downgraded"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemExpired                      WebhookEndpointEnabledEvents = "omnichannel_subscription_item_expired"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemCancellationScheduled        WebhookEndpointEnabledEvents = "omnichannel_subscription_item_cancellation_scheduled"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemScheduledCancellationRemoved WebhookEndpointEnabledEvents = "omnichannel_subscription_item_scheduled_cancellation_removed"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemResubscribed                 WebhookEndpointEnabledEvents = "omnichannel_subscription_item_resubscribed"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemUpgraded                     WebhookEndpointEnabledEvents = "omnichannel_subscription_item_upgraded"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemCancelled                    WebhookEndpointEnabledEvents = "omnichannel_subscription_item_cancelled"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionImported                         WebhookEndpointEnabledEvents = "omnichannel_subscription_imported"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemGracePeriodStarted           WebhookEndpointEnabledEvents = "omnichannel_subscription_item_grace_period_started"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemGracePeriodExpired           WebhookEndpointEnabledEvents = "omnichannel_subscription_item_grace_period_expired"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemDunningStarted               WebhookEndpointEnabledEvents = "omnichannel_subscription_item_dunning_started"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemDunningExpired               WebhookEndpointEnabledEvents = "omnichannel_subscription_item_dunning_expired"
	WebhookEndpointEnabledEventsRuleCreated                                             WebhookEndpointEnabledEvents = "rule_created"
	WebhookEndpointEnabledEventsRuleUpdated                                             WebhookEndpointEnabledEvents = "rule_updated"
	WebhookEndpointEnabledEventsRuleDeleted                                             WebhookEndpointEnabledEvents = "rule_deleted"
	WebhookEndpointEnabledEventsRecordPurchaseFailed                                    WebhookEndpointEnabledEvents = "record_purchase_failed"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemChangeScheduled              WebhookEndpointEnabledEvents = "omnichannel_subscription_item_change_scheduled"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemScheduledChangeRemoved       WebhookEndpointEnabledEvents = "omnichannel_subscription_item_scheduled_change_removed"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemReactivated                  WebhookEndpointEnabledEvents = "omnichannel_subscription_item_reactivated"
	WebhookEndpointEnabledEventsSalesOrderCreated                                       WebhookEndpointEnabledEvents = "sales_order_created"
	WebhookEndpointEnabledEventsSalesOrderUpdated                                       WebhookEndpointEnabledEvents = "sales_order_updated"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemChanged                      WebhookEndpointEnabledEvents = "omnichannel_subscription_item_changed"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemPaused                       WebhookEndpointEnabledEvents = "omnichannel_subscription_item_paused"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemResumed                      WebhookEndpointEnabledEvents = "omnichannel_subscription_item_resumed"
	WebhookEndpointEnabledEventsOmnichannelOneTimeOrderCreated                          WebhookEndpointEnabledEvents = "omnichannel_one_time_order_created"
	WebhookEndpointEnabledEventsOmnichannelOneTimeOrderItemCancelled                    WebhookEndpointEnabledEvents = "omnichannel_one_time_order_item_cancelled"
	WebhookEndpointEnabledEventsUsageFileIngested                                       WebhookEndpointEnabledEvents = "usage_file_ingested"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionItemPauseScheduled               WebhookEndpointEnabledEvents = "omnichannel_subscription_item_pause_scheduled"
	WebhookEndpointEnabledEventsOmnichannelSubscriptionMovedIn                          WebhookEndpointEnabledEvents = "omnichannel_subscription_moved_in"
	WebhookEndpointEnabledEventsOmnichannelTransactionCreated                           WebhookEndpointEnabledEvents = "omnichannel_transaction_created"
	WebhookEndpointEnabledEventsPlanCreated                                             WebhookEndpointEnabledEvents = "plan_created"
	WebhookEndpointEnabledEventsPlanUpdated                                             WebhookEndpointEnabledEvents = "plan_updated"
	WebhookEndpointEnabledEventsPlanDeleted                                             WebhookEndpointEnabledEvents = "plan_deleted"
	WebhookEndpointEnabledEventsAddonCreated                                            WebhookEndpointEnabledEvents = "addon_created"
	WebhookEndpointEnabledEventsAddonUpdated                                            WebhookEndpointEnabledEvents = "addon_updated"
	WebhookEndpointEnabledEventsAddonDeleted                                            WebhookEndpointEnabledEvents = "addon_deleted"
)

// just struct
type WebhookEndpoint struct {
	Id                          string                                     `json:"id"`
	Name                        string                                     `json:"name"`
	Url                         string                                     `json:"url"`
	SendCardResource            bool                                       `json:"send_card_resource"`
	Disabled                    bool                                       `json:"disabled"`
	PrimaryUrl                  bool                                       `json:"primary_url"`
	ApiVersion                  WebhookEndpointApiVersion                  `json:"api_version"`
	ChargebeeResponseSchemaType WebhookEndpointChargebeeResponseSchemaType `json:"chargebee_response_schema_type"`
	EnabledEvents               []EventType                                `json:"enabled_events"`
	Object                      string                                     `json:"object"`
}

// sub resources
// operations
// input params
type WebhookEndpointCreateRequest struct {
	Name                        string                                     `json:"name"`
	ApiVersion                  WebhookEndpointApiVersion                  `json:"api_version,omitempty"`
	Url                         string                                     `json:"url"`
	PrimaryUrl                  *bool                                      `json:"primary_url,omitempty"`
	Disabled                    *bool                                      `json:"disabled,omitempty"`
	BasicAuthPassword           string                                     `json:"basic_auth_password,omitempty"`
	BasicAuthUsername           string                                     `json:"basic_auth_username,omitempty"`
	SendCardResource            *bool                                      `json:"send_card_resource,omitempty"`
	ChargebeeResponseSchemaType WebhookEndpointChargebeeResponseSchemaType `json:"chargebee_response_schema_type,omitempty"`
	EnabledEvents               EventType                                  `json:"enabled_events,omitempty"`
	apiRequest                  `json:"-" form:"-"`
}

func (r *WebhookEndpointCreateRequest) payload() any { return r }

type WebhookEndpointUpdateRequest struct {
	Name              string                    `json:"name,omitempty"`
	ApiVersion        WebhookEndpointApiVersion `json:"api_version,omitempty"`
	Url               string                    `json:"url,omitempty"`
	PrimaryUrl        *bool                     `json:"primary_url,omitempty"`
	SendCardResource  *bool                     `json:"send_card_resource,omitempty"`
	BasicAuthPassword string                    `json:"basic_auth_password,omitempty"`
	BasicAuthUsername string                    `json:"basic_auth_username,omitempty"`
	Disabled          *bool                     `json:"disabled,omitempty"`
	EnabledEvents     EventType                 `json:"enabled_events,omitempty"`
	apiRequest        `json:"-" form:"-"`
}

func (r *WebhookEndpointUpdateRequest) payload() any { return r }

type WebhookEndpointListRequest struct {
	Limit      *int32 `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *WebhookEndpointListRequest) payload() any { return r }

// operation response
type WebhookEndpointCreateResponse struct {
	WebhookEndpoint *WebhookEndpoint `json:"webhook_endpoint,omitempty"`
	apiResponse
}

// operation response
type WebhookEndpointUpdateResponse struct {
	WebhookEndpoint *WebhookEndpoint `json:"webhook_endpoint,omitempty"`
	apiResponse
}

// operation response
type WebhookEndpointRetrieveResponse struct {
	WebhookEndpoint *WebhookEndpoint `json:"webhook_endpoint,omitempty"`
	apiResponse
}

// operation response
type WebhookEndpointDeleteResponse struct {
	WebhookEndpoint *WebhookEndpoint `json:"webhook_endpoint,omitempty"`
	apiResponse
}

// operation sub response
type WebhookEndpointListWebhookEndpointResponse struct {
	WebhookEndpoint *WebhookEndpoint `json:"webhook_endpoint,omitempty"`
}

// operation response
type WebhookEndpointListResponse struct {
	List       []*WebhookEndpointListWebhookEndpointResponse `json:"list,omitempty"`
	NextOffset string                                        `json:"next_offset,omitempty"`
	apiResponse
}
