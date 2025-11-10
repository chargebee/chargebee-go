package webhook

import (
	"encoding/json"
	"time"

	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"

	"github.com/chargebee/chargebee-go/v3/models/attacheditem"

	"github.com/chargebee/chargebee-go/v3/models/attribute"

	"github.com/chargebee/chargebee-go/v3/models/businessentity"

	"github.com/chargebee/chargebee-go/v3/models/businessentitytransfer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/contractterm"

	"github.com/chargebee/chargebee-go/v3/models/coupon"

	"github.com/chargebee/chargebee-go/v3/models/couponcode"

	"github.com/chargebee/chargebee-go/v3/models/couponset"

	"github.com/chargebee/chargebee-go/v3/models/creditnote"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/differentialprice"

	"github.com/chargebee/chargebee-go/v3/models/feature"

	"github.com/chargebee/chargebee-go/v3/models/gift"

	"github.com/chargebee/chargebee-go/v3/models/impactedcustomer"

	"github.com/chargebee/chargebee-go/v3/models/impacteditem"

	"github.com/chargebee/chargebee-go/v3/models/impacteditemprice"

	"github.com/chargebee/chargebee-go/v3/models/impactedsubscription"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/item"

	"github.com/chargebee/chargebee-go/v3/models/itemfamily"

	"github.com/chargebee/chargebee-go/v3/models/itemprice"

	"github.com/chargebee/chargebee-go/v3/models/metadata"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelonetimeorder"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelonetimeorderitem"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitemscheduledchange"

	"github.com/chargebee/chargebee-go/v3/models/omnichanneltransaction"

	"github.com/chargebee/chargebee-go/v3/models/order"

	"github.com/chargebee/chargebee-go/v3/models/paymentintent"

	"github.com/chargebee/chargebee-go/v3/models/paymentschedule"

	"github.com/chargebee/chargebee-go/v3/models/paymentschedulescheme"

	"github.com/chargebee/chargebee-go/v3/models/paymentsource"

	"github.com/chargebee/chargebee-go/v3/models/paymentvoucher"

	"github.com/chargebee/chargebee-go/v3/models/pricevariant"

	"github.com/chargebee/chargebee-go/v3/models/promotionalcredit"

	"github.com/chargebee/chargebee-go/v3/models/purchase"

	"github.com/chargebee/chargebee-go/v3/models/quote"

	"github.com/chargebee/chargebee-go/v3/models/ramp"

	"github.com/chargebee/chargebee-go/v3/models/recordedpurchase"

	"github.com/chargebee/chargebee-go/v3/models/rule"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/subscriptionentitlementscreateddetail"

	"github.com/chargebee/chargebee-go/v3/models/subscriptionentitlementsupdateddetail"

	"github.com/chargebee/chargebee-go/v3/models/taxwithheld"

	"github.com/chargebee/chargebee-go/v3/models/token"

	"github.com/chargebee/chargebee-go/v3/models/transaction"

	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"

	"github.com/chargebee/chargebee-go/v3/models/usagefile"

	"github.com/chargebee/chargebee-go/v3/models/virtualbankaccount"
)

// WebhookContentType represents the type of webhook event content
type WebhookContentType string

const (
	WebhookContentTypeSubscriptionPauseScheduled WebhookContentType = "subscription_pause_scheduled"

	WebhookContentTypeCustomerBusinessEntityChanged WebhookContentType = "customer_business_entity_changed"

	WebhookContentTypeSubscriptionAdvanceInvoiceScheduleAdded WebhookContentType = "subscription_advance_invoice_schedule_added"

	WebhookContentTypeGiftExpired WebhookContentType = "gift_expired"

	WebhookContentTypeTaxWithheldDeleted WebhookContentType = "tax_withheld_deleted"

	WebhookContentTypeUnbilledChargesDeleted WebhookContentType = "unbilled_charges_deleted"

	WebhookContentTypeCouponUpdated WebhookContentType = "coupon_updated"

	WebhookContentTypeOmnichannelSubscriptionItemReactivated WebhookContentType = "omnichannel_subscription_item_reactivated"

	WebhookContentTypeOmnichannelSubscriptionItemRenewed WebhookContentType = "omnichannel_subscription_item_renewed"

	WebhookContentTypeUnbilledChargesCreated WebhookContentType = "unbilled_charges_created"

	WebhookContentTypeSubscriptionResumed WebhookContentType = "subscription_resumed"

	WebhookContentTypeOmnichannelOneTimeOrderItemCancelled WebhookContentType = "omnichannel_one_time_order_item_cancelled"

	WebhookContentTypeSubscriptionCancelled WebhookContentType = "subscription_cancelled"

	WebhookContentTypeItemEntitlementsRemoved WebhookContentType = "item_entitlements_removed"

	WebhookContentTypeBusinessEntityCreated WebhookContentType = "business_entity_created"

	WebhookContentTypeCouponSetUpdated WebhookContentType = "coupon_set_updated"

	WebhookContentTypeDifferentialPriceUpdated WebhookContentType = "differential_price_updated"

	WebhookContentTypeOmnichannelSubscriptionItemPaused WebhookContentType = "omnichannel_subscription_item_paused"

	WebhookContentTypeEntitlementOverridesRemoved WebhookContentType = "entitlement_overrides_removed"

	WebhookContentTypeSubscriptionActivatedWithBackdating WebhookContentType = "subscription_activated_with_backdating"

	WebhookContentTypeSubscriptionTrialEndReminder WebhookContentType = "subscription_trial_end_reminder"

	WebhookContentTypeSubscriptionShippingAddressUpdated WebhookContentType = "subscription_shipping_address_updated"

	WebhookContentTypeVoucherCreateFailed WebhookContentType = "voucher_create_failed"

	WebhookContentTypeGiftClaimed WebhookContentType = "gift_claimed"

	WebhookContentTypeCustomerDeleted WebhookContentType = "customer_deleted"

	WebhookContentTypeRefundInitiated WebhookContentType = "refund_initiated"

	WebhookContentTypeInvoiceGeneratedWithBackdating WebhookContentType = "invoice_generated_with_backdating"

	WebhookContentTypeOmnichannelTransactionCreated WebhookContentType = "omnichannel_transaction_created"

	WebhookContentTypeAddUsagesReminder WebhookContentType = "add_usages_reminder"

	WebhookContentTypeVoucherCreated WebhookContentType = "voucher_created"

	WebhookContentTypeRuleUpdated WebhookContentType = "rule_updated"

	WebhookContentTypePaymentSchedulesCreated WebhookContentType = "payment_schedules_created"

	WebhookContentTypeFeatureActivated WebhookContentType = "feature_activated"

	WebhookContentTypePaymentSourceLocallyDeleted WebhookContentType = "payment_source_locally_deleted"

	WebhookContentTypeInvoiceGenerated WebhookContentType = "invoice_generated"

	WebhookContentTypeVoucherExpired WebhookContentType = "voucher_expired"

	WebhookContentTypeAuthorizationSucceeded WebhookContentType = "authorization_succeeded"

	WebhookContentTypeGiftScheduled WebhookContentType = "gift_scheduled"

	WebhookContentTypeSubscriptionChangesScheduled WebhookContentType = "subscription_changes_scheduled"

	WebhookContentTypeSubscriptionChangedWithBackdating WebhookContentType = "subscription_changed_with_backdating"

	WebhookContentTypeOmnichannelSubscriptionItemChanged WebhookContentType = "omnichannel_subscription_item_changed"

	WebhookContentTypeGiftUnclaimed WebhookContentType = "gift_unclaimed"

	WebhookContentTypeVirtualBankAccountAdded WebhookContentType = "virtual_bank_account_added"

	WebhookContentTypePaymentIntentCreated WebhookContentType = "payment_intent_created"

	WebhookContentTypeCreditNoteCreatedWithBackdating WebhookContentType = "credit_note_created_with_backdating"

	WebhookContentTypeContractTermTerminated WebhookContentType = "contract_term_terminated"

	WebhookContentTypeItemFamilyUpdated WebhookContentType = "item_family_updated"

	WebhookContentTypeOrderCreated WebhookContentType = "order_created"

	WebhookContentTypePriceVariantDeleted WebhookContentType = "price_variant_deleted"

	WebhookContentTypeSubscriptionMovementFailed WebhookContentType = "subscription_movement_failed"

	WebhookContentTypeCustomerMovedIn WebhookContentType = "customer_moved_in"

	WebhookContentTypeSubscriptionAdvanceInvoiceScheduleUpdated WebhookContentType = "subscription_advance_invoice_schedule_updated"

	WebhookContentTypeItemDeleted WebhookContentType = "item_deleted"

	WebhookContentTypeSubscriptionRampDrafted WebhookContentType = "subscription_ramp_drafted"

	WebhookContentTypeDunningUpdated WebhookContentType = "dunning_updated"

	WebhookContentTypeItemEntitlementsUpdated WebhookContentType = "item_entitlements_updated"

	WebhookContentTypeTokenConsumed WebhookContentType = "token_consumed"

	WebhookContentTypeHierarchyDeleted WebhookContentType = "hierarchy_deleted"

	WebhookContentTypeSubscriptionCancellationScheduled WebhookContentType = "subscription_cancellation_scheduled"

	WebhookContentTypeSubscriptionRenewed WebhookContentType = "subscription_renewed"

	WebhookContentTypeFeatureUpdated WebhookContentType = "feature_updated"

	WebhookContentTypeFeatureDeleted WebhookContentType = "feature_deleted"

	WebhookContentTypeItemFamilyCreated WebhookContentType = "item_family_created"

	WebhookContentTypeOmnichannelSubscriptionItemScheduledChangeRemoved WebhookContentType = "omnichannel_subscription_item_scheduled_change_removed"

	WebhookContentTypeOmnichannelSubscriptionItemResumed WebhookContentType = "omnichannel_subscription_item_resumed"

	WebhookContentTypePurchaseCreated WebhookContentType = "purchase_created"

	WebhookContentTypeEntitlementOverridesUpdated WebhookContentType = "entitlement_overrides_updated"

	WebhookContentTypeItemFamilyDeleted WebhookContentType = "item_family_deleted"

	WebhookContentTypeSubscriptionResumptionScheduled WebhookContentType = "subscription_resumption_scheduled"

	WebhookContentTypeFeatureReactivated WebhookContentType = "feature_reactivated"

	WebhookContentTypeCouponCodesDeleted WebhookContentType = "coupon_codes_deleted"

	WebhookContentTypeCardExpired WebhookContentType = "card_expired"

	WebhookContentTypeCreditNoteUpdated WebhookContentType = "credit_note_updated"

	WebhookContentTypeOmnichannelSubscriptionItemDowngraded WebhookContentType = "omnichannel_subscription_item_downgraded"

	WebhookContentTypePriceVariantUpdated WebhookContentType = "price_variant_updated"

	WebhookContentTypePromotionalCreditsDeducted WebhookContentType = "promotional_credits_deducted"

	WebhookContentTypeSubscriptionRampApplied WebhookContentType = "subscription_ramp_applied"

	WebhookContentTypeSubscriptionPaused WebhookContentType = "subscription_paused"

	WebhookContentTypeOrderReadyToProcess WebhookContentType = "order_ready_to_process"

	WebhookContentTypeFeatureCreated WebhookContentType = "feature_created"

	WebhookContentTypeTransactionDeleted WebhookContentType = "transaction_deleted"

	WebhookContentTypeCreditNoteCreated WebhookContentType = "credit_note_created"

	WebhookContentTypeOmnichannelSubscriptionItemResubscribed WebhookContentType = "omnichannel_subscription_item_resubscribed"

	WebhookContentTypeRecordPurchaseFailed WebhookContentType = "record_purchase_failed"

	WebhookContentTypeItemCreated WebhookContentType = "item_created"

	WebhookContentTypeTransactionUpdated WebhookContentType = "transaction_updated"

	WebhookContentTypeMrrUpdated WebhookContentType = "mrr_updated"

	WebhookContentTypeUnbilledChargesInvoiced WebhookContentType = "unbilled_charges_invoiced"

	WebhookContentTypeItemPriceUpdated WebhookContentType = "item_price_updated"

	WebhookContentTypeCouponCodesUpdated WebhookContentType = "coupon_codes_updated"

	WebhookContentTypeVirtualBankAccountUpdated WebhookContentType = "virtual_bank_account_updated"

	WebhookContentTypeContractTermCreated WebhookContentType = "contract_term_created"

	WebhookContentTypeSubscriptionChanged WebhookContentType = "subscription_changed"

	WebhookContentTypePaymentFailed WebhookContentType = "payment_failed"

	WebhookContentTypeCreditNoteDeleted WebhookContentType = "credit_note_deleted"

	WebhookContentTypeTaxWithheldRefunded WebhookContentType = "tax_withheld_refunded"

	WebhookContentTypeContractTermCompleted WebhookContentType = "contract_term_completed"

	WebhookContentTypePaymentSchedulesUpdated WebhookContentType = "payment_schedules_updated"

	WebhookContentTypeOmnichannelSubscriptionItemExpired WebhookContentType = "omnichannel_subscription_item_expired"

	WebhookContentTypeCardUpdated WebhookContentType = "card_updated"

	WebhookContentTypeCustomerCreated WebhookContentType = "customer_created"

	WebhookContentTypeSubscriptionRenewalReminder WebhookContentType = "subscription_renewal_reminder"

	WebhookContentTypeOrderDelivered WebhookContentType = "order_delivered"

	WebhookContentTypeOmnichannelSubscriptionItemCancellationScheduled WebhookContentType = "omnichannel_subscription_item_cancellation_scheduled"

	WebhookContentTypeOmnichannelSubscriptionItemGracePeriodExpired WebhookContentType = "omnichannel_subscription_item_grace_period_expired"

	WebhookContentTypeCouponCodesAdded WebhookContentType = "coupon_codes_added"

	WebhookContentTypeGiftCancelled WebhookContentType = "gift_cancelled"

	WebhookContentTypeOrderCancelled WebhookContentType = "order_cancelled"

	WebhookContentTypeCouponDeleted WebhookContentType = "coupon_deleted"

	WebhookContentTypeSubscriptionScheduledChangesRemoved WebhookContentType = "subscription_scheduled_changes_removed"

	WebhookContentTypePendingInvoiceCreated WebhookContentType = "pending_invoice_created"

	WebhookContentTypeEntitlementOverridesAutoRemoved WebhookContentType = "entitlement_overrides_auto_removed"

	WebhookContentTypeOmnichannelSubscriptionItemUpgraded WebhookContentType = "omnichannel_subscription_item_upgraded"

	WebhookContentTypeSubscriptionBusinessEntityChanged WebhookContentType = "subscription_business_entity_changed"

	WebhookContentTypeOmnichannelOneTimeOrderCreated WebhookContentType = "omnichannel_one_time_order_created"

	WebhookContentTypePaymentSourceDeleted WebhookContentType = "payment_source_deleted"

	WebhookContentTypeOmnichannelSubscriptionItemCancelled WebhookContentType = "omnichannel_subscription_item_cancelled"

	WebhookContentTypeQuoteDeleted WebhookContentType = "quote_deleted"

	WebhookContentTypeInvoiceUpdated WebhookContentType = "invoice_updated"

	WebhookContentTypeSubscriptionAdvanceInvoiceScheduleRemoved WebhookContentType = "subscription_advance_invoice_schedule_removed"

	WebhookContentTypeCardDeleted WebhookContentType = "card_deleted"

	WebhookContentTypeOrderReadyToShip WebhookContentType = "order_ready_to_ship"

	WebhookContentTypeSubscriptionMovedOut WebhookContentType = "subscription_moved_out"

	WebhookContentTypePaymentScheduleSchemeCreated WebhookContentType = "payment_schedule_scheme_created"

	WebhookContentTypeBusinessEntityUpdated WebhookContentType = "business_entity_updated"

	WebhookContentTypeSubscriptionScheduledResumptionRemoved WebhookContentType = "subscription_scheduled_resumption_removed"

	WebhookContentTypePaymentInitiated WebhookContentType = "payment_initiated"

	WebhookContentTypeFeatureArchived WebhookContentType = "feature_archived"

	WebhookContentTypeSubscriptionReactivatedWithBackdating WebhookContentType = "subscription_reactivated_with_backdating"

	WebhookContentTypeOmnichannelSubscriptionImported WebhookContentType = "omnichannel_subscription_imported"

	WebhookContentTypeTokenExpired WebhookContentType = "token_expired"

	WebhookContentTypeCardAdded WebhookContentType = "card_added"

	WebhookContentTypeCouponCreated WebhookContentType = "coupon_created"

	WebhookContentTypeRuleDeleted WebhookContentType = "rule_deleted"

	WebhookContentTypeItemPriceEntitlementsUpdated WebhookContentType = "item_price_entitlements_updated"

	WebhookContentTypeItemPriceDeleted WebhookContentType = "item_price_deleted"

	WebhookContentTypeVirtualBankAccountDeleted WebhookContentType = "virtual_bank_account_deleted"

	WebhookContentTypePaymentScheduleSchemeDeleted WebhookContentType = "payment_schedule_scheme_deleted"

	WebhookContentTypeSubscriptionCreated WebhookContentType = "subscription_created"

	WebhookContentTypeSubscriptionEntitlementsCreated WebhookContentType = "subscription_entitlements_created"

	WebhookContentTypeOrderReturned WebhookContentType = "order_returned"

	WebhookContentTypeSubscriptionDeleted WebhookContentType = "subscription_deleted"

	WebhookContentTypePaymentSourceAdded WebhookContentType = "payment_source_added"

	WebhookContentTypeSubscriptionMovedIn WebhookContentType = "subscription_moved_in"

	WebhookContentTypeItemPriceCreated WebhookContentType = "item_price_created"

	WebhookContentTypeSubscriptionScheduledCancellationRemoved WebhookContentType = "subscription_scheduled_cancellation_removed"

	WebhookContentTypePaymentRefunded WebhookContentType = "payment_refunded"

	WebhookContentTypeUsageFileIngested WebhookContentType = "usage_file_ingested"

	WebhookContentTypeOmnichannelSubscriptionMovedIn WebhookContentType = "omnichannel_subscription_moved_in"

	WebhookContentTypeDifferentialPriceCreated WebhookContentType = "differential_price_created"

	WebhookContentTypeTransactionCreated WebhookContentType = "transaction_created"

	WebhookContentTypePaymentSucceeded WebhookContentType = "payment_succeeded"

	WebhookContentTypeSubscriptionCanceledWithBackdating WebhookContentType = "subscription_canceled_with_backdating"

	WebhookContentTypeUnbilledChargesVoided WebhookContentType = "unbilled_charges_voided"

	WebhookContentTypeQuoteCreated WebhookContentType = "quote_created"

	WebhookContentTypeCouponSetDeleted WebhookContentType = "coupon_set_deleted"

	WebhookContentTypeAttachedItemCreated WebhookContentType = "attached_item_created"

	WebhookContentTypeSalesOrderCreated WebhookContentType = "sales_order_created"

	WebhookContentTypeCustomerChanged WebhookContentType = "customer_changed"

	WebhookContentTypeSubscriptionStarted WebhookContentType = "subscription_started"

	WebhookContentTypeSubscriptionActivated WebhookContentType = "subscription_activated"

	WebhookContentTypePaymentSourceExpiring WebhookContentType = "payment_source_expiring"

	WebhookContentTypeSubscriptionReactivated WebhookContentType = "subscription_reactivated"

	WebhookContentTypeOrderUpdated WebhookContentType = "order_updated"

	WebhookContentTypeSubscriptionScheduledPauseRemoved WebhookContentType = "subscription_scheduled_pause_removed"

	WebhookContentTypeSubscriptionCancellationReminder WebhookContentType = "subscription_cancellation_reminder"

	WebhookContentTypeSubscriptionCreatedWithBackdating WebhookContentType = "subscription_created_with_backdating"

	WebhookContentTypeSubscriptionRampCreated WebhookContentType = "subscription_ramp_created"

	WebhookContentTypeOrderDeleted WebhookContentType = "order_deleted"

	WebhookContentTypeOmnichannelSubscriptionItemPauseScheduled WebhookContentType = "omnichannel_subscription_item_pause_scheduled"

	WebhookContentTypeGiftUpdated WebhookContentType = "gift_updated"

	WebhookContentTypeSubscriptionTrialExtended WebhookContentType = "subscription_trial_extended"

	WebhookContentTypeOmnichannelSubscriptionItemGracePeriodStarted WebhookContentType = "omnichannel_subscription_item_grace_period_started"

	WebhookContentTypeCardExpiryReminder WebhookContentType = "card_expiry_reminder"

	WebhookContentTypeTokenCreated WebhookContentType = "token_created"

	WebhookContentTypePromotionalCreditsAdded WebhookContentType = "promotional_credits_added"

	WebhookContentTypeSubscriptionRampUpdated WebhookContentType = "subscription_ramp_updated"

	WebhookContentTypeCustomerEntitlementsUpdated WebhookContentType = "customer_entitlements_updated"

	WebhookContentTypePaymentSourceExpired WebhookContentType = "payment_source_expired"

	WebhookContentTypeCustomerMovedOut WebhookContentType = "customer_moved_out"

	WebhookContentTypeSubscriptionEntitlementsUpdated WebhookContentType = "subscription_entitlements_updated"

	WebhookContentTypeOmnichannelSubscriptionItemDunningExpired WebhookContentType = "omnichannel_subscription_item_dunning_expired"

	WebhookContentTypeHierarchyCreated WebhookContentType = "hierarchy_created"

	WebhookContentTypeAttachedItemDeleted WebhookContentType = "attached_item_deleted"

	WebhookContentTypeOmnichannelSubscriptionItemScheduledCancellationRemoved WebhookContentType = "omnichannel_subscription_item_scheduled_cancellation_removed"

	WebhookContentTypeItemUpdated WebhookContentType = "item_updated"

	WebhookContentTypeCouponSetCreated WebhookContentType = "coupon_set_created"

	WebhookContentTypePaymentIntentUpdated WebhookContentType = "payment_intent_updated"

	WebhookContentTypeOrderResent WebhookContentType = "order_resent"

	WebhookContentTypeOmnichannelSubscriptionCreated WebhookContentType = "omnichannel_subscription_created"

	WebhookContentTypeTaxWithheldRecorded WebhookContentType = "tax_withheld_recorded"

	WebhookContentTypePriceVariantCreated WebhookContentType = "price_variant_created"

	WebhookContentTypeDifferentialPriceDeleted WebhookContentType = "differential_price_deleted"

	WebhookContentTypeSubscriptionItemsRenewed WebhookContentType = "subscription_items_renewed"

	WebhookContentTypeRuleCreated WebhookContentType = "rule_created"

	WebhookContentTypeContractTermCancelled WebhookContentType = "contract_term_cancelled"

	WebhookContentTypeContractTermRenewed WebhookContentType = "contract_term_renewed"

	WebhookContentTypeInvoiceDeleted WebhookContentType = "invoice_deleted"

	WebhookContentTypeItemPriceEntitlementsRemoved WebhookContentType = "item_price_entitlements_removed"

	WebhookContentTypeSalesOrderUpdated WebhookContentType = "sales_order_updated"

	WebhookContentTypeOmnichannelSubscriptionItemDunningStarted WebhookContentType = "omnichannel_subscription_item_dunning_started"

	WebhookContentTypeOmnichannelSubscriptionItemChangeScheduled WebhookContentType = "omnichannel_subscription_item_change_scheduled"

	WebhookContentTypePendingInvoiceUpdated WebhookContentType = "pending_invoice_updated"

	WebhookContentTypeQuoteUpdated WebhookContentType = "quote_updated"

	WebhookContentTypeAttachedItemUpdated WebhookContentType = "attached_item_updated"

	WebhookContentTypePaymentSourceUpdated WebhookContentType = "payment_source_updated"

	WebhookContentTypeBusinessEntityDeleted WebhookContentType = "business_entity_deleted"

	WebhookContentTypeAuthorizationVoided WebhookContentType = "authorization_voided"

	WebhookContentTypeSubscriptionRampDeleted WebhookContentType = "subscription_ramp_deleted"
)

// Event content structures for each webhook type

type SubscriptionPauseScheduledContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type CustomerBusinessEntityChangedContent struct {
	BusinessEntityTransfer *businessentitytransfer.BusinessEntityTransfer `json:"business_entity_transfer,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type SubscriptionAdvanceInvoiceScheduleAddedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type GiftExpiredContent struct {
	Gift *gift.Gift `json:"gift,omitempty"`
}

type TaxWithheldDeletedContent struct {
	TaxWithheld *taxwithheld.TaxWithheld `json:"tax_withheld,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`
}

type UnbilledChargesDeletedContent struct {
	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type CouponUpdatedContent struct {
	Coupon *coupon.Coupon `json:"coupon,omitempty"`
}

type OmnichannelSubscriptionItemReactivatedContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type OmnichannelSubscriptionItemRenewedContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`

	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type UnbilledChargesCreatedContent struct {
	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type SubscriptionResumedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type OmnichannelOneTimeOrderItemCancelledContent struct {
	OmnichannelOneTimeOrder *omnichannelonetimeorder.OmnichannelOneTimeOrder `json:"omnichannel_one_time_order,omitempty"`

	OmnichannelOneTimeOrderItem *omnichannelonetimeorderitem.OmnichannelOneTimeOrderItem `json:"omnichannel_one_time_order_item,omitempty"`

	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type SubscriptionCancelledContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type ItemEntitlementsRemovedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`

	ImpactedItem *impacteditem.ImpactedItem `json:"impacted_item,omitempty"`

	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type BusinessEntityCreatedContent struct {
	BusinessEntity *businessentity.BusinessEntity `json:"business_entity,omitempty"`
}

type CouponSetUpdatedContent struct {
	Coupon *coupon.Coupon `json:"coupon,omitempty"`

	CouponSet *couponset.CouponSet `json:"coupon_set,omitempty"`
}

type DifferentialPriceUpdatedContent struct {
	DifferentialPrice *differentialprice.DifferentialPrice `json:"differential_price,omitempty"`
}

type OmnichannelSubscriptionItemPausedContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type EntitlementOverridesRemovedContent struct {
	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`
}

type SubscriptionActivatedWithBackdatingContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type SubscriptionTrialEndReminderContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type SubscriptionShippingAddressUpdatedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type VoucherCreateFailedContent struct {
	PaymentVoucher *paymentvoucher.PaymentVoucher `json:"payment_voucher,omitempty"`
}

type GiftClaimedContent struct {
	Gift *gift.Gift `json:"gift,omitempty"`
}

type CustomerDeletedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Subscription *subscription.Subscription `json:"subscription,omitempty"`
}

type RefundInitiatedContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type InvoiceGeneratedWithBackdatingContent struct {
	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type OmnichannelTransactionCreatedContent struct {
	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`
}

type AddUsagesReminderContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type VoucherCreatedContent struct {
	PaymentVoucher *paymentvoucher.PaymentVoucher `json:"payment_voucher,omitempty"`
}

type RuleUpdatedContent struct {
	Rule *rule.Rule `json:"rule,omitempty"`
}

type PaymentSchedulesCreatedContent struct {
	PaymentSchedule *paymentschedule.PaymentSchedule `json:"payment_schedule,omitempty"`
}

type FeatureActivatedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`

	ImpactedItem *impacteditem.ImpactedItem `json:"impacted_item,omitempty"`

	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type PaymentSourceLocallyDeletedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	PaymentSource *paymentsource.PaymentSource `json:"payment_source,omitempty"`
}

type InvoiceGeneratedContent struct {
	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type VoucherExpiredContent struct {
	PaymentVoucher *paymentvoucher.PaymentVoucher `json:"payment_voucher,omitempty"`
}

type AuthorizationSucceededContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`
}

type GiftScheduledContent struct {
	Gift *gift.Gift `json:"gift,omitempty"`
}

type SubscriptionChangesScheduledContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type SubscriptionChangedWithBackdatingContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type OmnichannelSubscriptionItemChangedContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`

	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type GiftUnclaimedContent struct {
	Gift *gift.Gift `json:"gift,omitempty"`
}

type VirtualBankAccountAddedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	VirtualBankAccount *virtualbankaccount.VirtualBankAccount `json:"virtual_bank_account,omitempty"`
}

type PaymentIntentCreatedContent struct {
	PaymentIntent *paymentintent.PaymentIntent `json:"payment_intent,omitempty"`
}

type CreditNoteCreatedWithBackdatingContent struct {
	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`
}

type ContractTermTerminatedContent struct {
	ContractTerm *contractterm.ContractTerm `json:"contract_term,omitempty"`
}

type ItemFamilyUpdatedContent struct {
	ItemFamily *itemfamily.ItemFamily `json:"item_family,omitempty"`
}

type OrderCreatedContent struct {
	Order *order.Order `json:"order,omitempty"`
}

type PriceVariantDeletedContent struct {
	PriceVariant *pricevariant.PriceVariant `json:"price_variant,omitempty"`

	Attribute *attribute.Attribute `json:"attribute,omitempty"`
}

type SubscriptionMovementFailedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`
}

type CustomerMovedInContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type SubscriptionAdvanceInvoiceScheduleUpdatedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type ItemDeletedContent struct {
	Item *item.Item `json:"item,omitempty"`
}

type SubscriptionRampDraftedContent struct {
	Ramp *ramp.Ramp `json:"ramp,omitempty"`
}

type DunningUpdatedContent struct {
	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type ItemEntitlementsUpdatedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`

	ImpactedItem *impacteditem.ImpactedItem `json:"impacted_item,omitempty"`

	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type TokenConsumedContent struct {
	Token *token.Token `json:"token,omitempty"`
}

type HierarchyDeletedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`
}

type SubscriptionCancellationScheduledContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type SubscriptionRenewedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type FeatureUpdatedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`
}

type FeatureDeletedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`

	ImpactedItem *impacteditem.ImpactedItem `json:"impacted_item,omitempty"`

	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type ItemFamilyCreatedContent struct {
	ItemFamily *itemfamily.ItemFamily `json:"item_family,omitempty"`
}

type OmnichannelSubscriptionItemScheduledChangeRemovedContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type OmnichannelSubscriptionItemResumedContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type PurchaseCreatedContent struct {
	Purchase *purchase.Purchase `json:"purchase,omitempty"`
}

type EntitlementOverridesUpdatedContent struct {
	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`
}

type ItemFamilyDeletedContent struct {
	ItemFamily *itemfamily.ItemFamily `json:"item_family,omitempty"`
}

type SubscriptionResumptionScheduledContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type FeatureReactivatedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`
}

type CouponCodesDeletedContent struct {
	Coupon *coupon.Coupon `json:"coupon,omitempty"`

	CouponSet *couponset.CouponSet `json:"coupon_set,omitempty"`

	CouponCode *couponcode.CouponCode `json:"coupon_code,omitempty"`
}

type CardExpiredContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type CreditNoteUpdatedContent struct {
	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`
}

type OmnichannelSubscriptionItemDowngradedContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type PriceVariantUpdatedContent struct {
	PriceVariant *pricevariant.PriceVariant `json:"price_variant,omitempty"`

	Attribute *attribute.Attribute `json:"attribute,omitempty"`
}

type PromotionalCreditsDeductedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	PromotionalCredit *promotionalcredit.PromotionalCredit `json:"promotional_credit,omitempty"`
}

type SubscriptionRampAppliedContent struct {
	Ramp *ramp.Ramp `json:"ramp,omitempty"`
}

type SubscriptionPausedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type OrderReadyToProcessContent struct {
	Order *order.Order `json:"order,omitempty"`
}

type FeatureCreatedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`

	ImpactedItem *impacteditem.ImpactedItem `json:"impacted_item,omitempty"`

	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type TransactionDeletedContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`
}

type CreditNoteCreatedContent struct {
	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`
}

type OmnichannelSubscriptionItemResubscribedContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`

	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type RecordPurchaseFailedContent struct {
	RecordedPurchase *recordedpurchase.RecordedPurchase `json:"recorded_purchase,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type ItemCreatedContent struct {
	Item *item.Item `json:"item,omitempty"`
}

type TransactionUpdatedContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`
}

type MrrUpdatedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`
}

type UnbilledChargesInvoicedContent struct {
	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type ItemPriceUpdatedContent struct {
	ItemPrice *itemprice.ItemPrice `json:"item_price,omitempty"`
}

type CouponCodesUpdatedContent struct {
	Coupon *coupon.Coupon `json:"coupon,omitempty"`

	CouponSet *couponset.CouponSet `json:"coupon_set,omitempty"`
}

type VirtualBankAccountUpdatedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	VirtualBankAccount *virtualbankaccount.VirtualBankAccount `json:"virtual_bank_account,omitempty"`
}

type ContractTermCreatedContent struct {
	ContractTerm *contractterm.ContractTerm `json:"contract_term,omitempty"`
}

type SubscriptionChangedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type PaymentFailedContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type CreditNoteDeletedContent struct {
	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`
}

type TaxWithheldRefundedContent struct {
	TaxWithheld *taxwithheld.TaxWithheld `json:"tax_withheld,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`
}

type ContractTermCompletedContent struct {
	ContractTerm *contractterm.ContractTerm `json:"contract_term,omitempty"`
}

type PaymentSchedulesUpdatedContent struct {
	PaymentSchedule *paymentschedule.PaymentSchedule `json:"payment_schedule,omitempty"`
}

type OmnichannelSubscriptionItemExpiredContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type CardUpdatedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type CustomerCreatedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type SubscriptionRenewalReminderContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type OrderDeliveredContent struct {
	Order *order.Order `json:"order,omitempty"`
}

type OmnichannelSubscriptionItemCancellationScheduledContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type OmnichannelSubscriptionItemGracePeriodExpiredContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type CouponCodesAddedContent struct {
	Coupon *coupon.Coupon `json:"coupon,omitempty"`

	CouponSet *couponset.CouponSet `json:"coupon_set,omitempty"`
}

type GiftCancelledContent struct {
	Gift *gift.Gift `json:"gift,omitempty"`
}

type OrderCancelledContent struct {
	Order *order.Order `json:"order,omitempty"`
}

type CouponDeletedContent struct {
	Coupon *coupon.Coupon `json:"coupon,omitempty"`
}

type SubscriptionScheduledChangesRemovedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type PendingInvoiceCreatedContent struct {
	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type EntitlementOverridesAutoRemovedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`

	ImpactedItem *impacteditem.ImpactedItem `json:"impacted_item,omitempty"`

	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type OmnichannelSubscriptionItemUpgradedContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`

	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type SubscriptionBusinessEntityChangedContent struct {
	BusinessEntityTransfer *businessentitytransfer.BusinessEntityTransfer `json:"business_entity_transfer,omitempty"`

	Subscription *subscription.Subscription `json:"subscription,omitempty"`
}

type OmnichannelOneTimeOrderCreatedContent struct {
	OmnichannelOneTimeOrder *omnichannelonetimeorder.OmnichannelOneTimeOrder `json:"omnichannel_one_time_order,omitempty"`

	OmnichannelOneTimeOrderItem *omnichannelonetimeorderitem.OmnichannelOneTimeOrderItem `json:"omnichannel_one_time_order_item,omitempty"`

	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type PaymentSourceDeletedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	PaymentSource *paymentsource.PaymentSource `json:"payment_source,omitempty"`
}

type OmnichannelSubscriptionItemCancelledContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type QuoteDeletedContent struct {
	Quote *quote.Quote `json:"quote,omitempty"`
}

type InvoiceUpdatedContent struct {
	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type SubscriptionAdvanceInvoiceScheduleRemovedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type CardDeletedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type OrderReadyToShipContent struct {
	Order *order.Order `json:"order,omitempty"`
}

type SubscriptionMovedOutContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`
}

type PaymentScheduleSchemeCreatedContent struct {
	PaymentScheduleScheme *paymentschedulescheme.PaymentScheduleScheme `json:"payment_schedule_scheme,omitempty"`
}

type BusinessEntityUpdatedContent struct {
	BusinessEntity *businessentity.BusinessEntity `json:"business_entity,omitempty"`
}

type SubscriptionScheduledResumptionRemovedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type PaymentInitiatedContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type FeatureArchivedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`
}

type SubscriptionReactivatedWithBackdatingContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type OmnichannelSubscriptionImportedContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`

	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type TokenExpiredContent struct {
	Token *token.Token `json:"token,omitempty"`
}

type CardAddedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type CouponCreatedContent struct {
	Coupon *coupon.Coupon `json:"coupon,omitempty"`
}

type RuleDeletedContent struct {
	Rule *rule.Rule `json:"rule,omitempty"`
}

type ItemPriceEntitlementsUpdatedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`

	ImpactedItemPrice *impacteditemprice.ImpactedItemPrice `json:"impacted_item_price,omitempty"`

	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type ItemPriceDeletedContent struct {
	ItemPrice *itemprice.ItemPrice `json:"item_price,omitempty"`
}

type VirtualBankAccountDeletedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	VirtualBankAccount *virtualbankaccount.VirtualBankAccount `json:"virtual_bank_account,omitempty"`
}

type PaymentScheduleSchemeDeletedContent struct {
	PaymentScheduleScheme *paymentschedulescheme.PaymentScheduleScheme `json:"payment_schedule_scheme,omitempty"`
}

type SubscriptionCreatedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type SubscriptionEntitlementsCreatedContent struct {
	SubscriptionEntitlementsCreatedDetail *subscriptionentitlementscreateddetail.SubscriptionEntitlementsCreatedDetail `json:"subscription_entitlements_created_detail,omitempty"`
}

type OrderReturnedContent struct {
	Order *order.Order `json:"order,omitempty"`
}

type SubscriptionDeletedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type PaymentSourceAddedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	PaymentSource *paymentsource.PaymentSource `json:"payment_source,omitempty"`
}

type SubscriptionMovedInContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`
}

type ItemPriceCreatedContent struct {
	ItemPrice *itemprice.ItemPrice `json:"item_price,omitempty"`
}

type SubscriptionScheduledCancellationRemovedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type PaymentRefundedContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type UsageFileIngestedContent struct {
	UsageFile *usagefile.UsageFile `json:"usage_file,omitempty"`
}

type OmnichannelSubscriptionMovedInContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type DifferentialPriceCreatedContent struct {
	DifferentialPrice *differentialprice.DifferentialPrice `json:"differential_price,omitempty"`
}

type TransactionCreatedContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`
}

type PaymentSucceededContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type SubscriptionCanceledWithBackdatingContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type UnbilledChargesVoidedContent struct {
	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type QuoteCreatedContent struct {
	Quote *quote.Quote `json:"quote,omitempty"`
}

type CouponSetDeletedContent struct {
	Coupon *coupon.Coupon `json:"coupon,omitempty"`

	CouponSet *couponset.CouponSet `json:"coupon_set,omitempty"`
}

type AttachedItemCreatedContent struct {
	AttachedItem *attacheditem.AttachedItem `json:"attached_item,omitempty"`
}

type SalesOrderCreatedContent struct {
}

type CustomerChangedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type SubscriptionStartedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type SubscriptionActivatedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type PaymentSourceExpiringContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	PaymentSource *paymentsource.PaymentSource `json:"payment_source,omitempty"`
}

type SubscriptionReactivatedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type OrderUpdatedContent struct {
	Order *order.Order `json:"order,omitempty"`
}

type SubscriptionScheduledPauseRemovedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type SubscriptionCancellationReminderContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type SubscriptionCreatedWithBackdatingContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type SubscriptionRampCreatedContent struct {
	Ramp *ramp.Ramp `json:"ramp,omitempty"`
}

type OrderDeletedContent struct {
	Order *order.Order `json:"order,omitempty"`
}

type OmnichannelSubscriptionItemPauseScheduledContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type GiftUpdatedContent struct {
	Gift *gift.Gift `json:"gift,omitempty"`
}

type SubscriptionTrialExtendedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type OmnichannelSubscriptionItemGracePeriodStartedContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type CardExpiryReminderContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type TokenCreatedContent struct {
	Token *token.Token `json:"token,omitempty"`
}

type PromotionalCreditsAddedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	PromotionalCredit *promotionalcredit.PromotionalCredit `json:"promotional_credit,omitempty"`
}

type SubscriptionRampUpdatedContent struct {
	Ramp *ramp.Ramp `json:"ramp,omitempty"`
}

type CustomerEntitlementsUpdatedContent struct {
	ImpactedCustomer *impactedcustomer.ImpactedCustomer `json:"impacted_customer,omitempty"`
}

type PaymentSourceExpiredContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	PaymentSource *paymentsource.PaymentSource `json:"payment_source,omitempty"`
}

type CustomerMovedOutContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type SubscriptionEntitlementsUpdatedContent struct {
	SubscriptionEntitlementsUpdatedDetail *subscriptionentitlementsupdateddetail.SubscriptionEntitlementsUpdatedDetail `json:"subscription_entitlements_updated_detail,omitempty"`
}

type OmnichannelSubscriptionItemDunningExpiredContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type HierarchyCreatedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`
}

type AttachedItemDeletedContent struct {
	AttachedItem *attacheditem.AttachedItem `json:"attached_item,omitempty"`
}

type OmnichannelSubscriptionItemScheduledCancellationRemovedContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type ItemUpdatedContent struct {
	Item *item.Item `json:"item,omitempty"`
}

type CouponSetCreatedContent struct {
	Coupon *coupon.Coupon `json:"coupon,omitempty"`

	CouponSet *couponset.CouponSet `json:"coupon_set,omitempty"`
}

type PaymentIntentUpdatedContent struct {
	PaymentIntent *paymentintent.PaymentIntent `json:"payment_intent,omitempty"`
}

type OrderResentContent struct {
	Order *order.Order `json:"order,omitempty"`
}

type OmnichannelSubscriptionCreatedContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`

	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type TaxWithheldRecordedContent struct {
	TaxWithheld *taxwithheld.TaxWithheld `json:"tax_withheld,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`
}

type PriceVariantCreatedContent struct {
	PriceVariant *pricevariant.PriceVariant `json:"price_variant,omitempty"`

	Attribute *attribute.Attribute `json:"attribute,omitempty"`
}

type DifferentialPriceDeletedContent struct {
	DifferentialPrice *differentialprice.DifferentialPrice `json:"differential_price,omitempty"`
}

type SubscriptionItemsRenewedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type RuleCreatedContent struct {
	Rule *rule.Rule `json:"rule,omitempty"`
}

type ContractTermCancelledContent struct {
	ContractTerm *contractterm.ContractTerm `json:"contract_term,omitempty"`
}

type ContractTermRenewedContent struct {
	ContractTerm *contractterm.ContractTerm `json:"contract_term,omitempty"`
}

type InvoiceDeletedContent struct {
	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type ItemPriceEntitlementsRemovedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`

	ImpactedItemPrice *impacteditemprice.ImpactedItemPrice `json:"impacted_item_price,omitempty"`

	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type SalesOrderUpdatedContent struct {
}

type OmnichannelSubscriptionItemDunningStartedContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type OmnichannelSubscriptionItemChangeScheduledContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type PendingInvoiceUpdatedContent struct {
	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type QuoteUpdatedContent struct {
	Quote *quote.Quote `json:"quote,omitempty"`
}

type AttachedItemUpdatedContent struct {
	AttachedItem *attacheditem.AttachedItem `json:"attached_item,omitempty"`
}

type PaymentSourceUpdatedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	PaymentSource *paymentsource.PaymentSource `json:"payment_source,omitempty"`
}

type BusinessEntityDeletedContent struct {
	BusinessEntity *businessentity.BusinessEntity `json:"business_entity,omitempty"`
}

type AuthorizationVoidedContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`
}

type SubscriptionRampDeletedContent struct {
	Ramp *ramp.Ramp `json:"ramp,omitempty"`
}

// WebhookEventInterface defines the common interface for all webhook events
type WebhookEventInterface interface {
	GetEventType() string
	GetEventId() string
	GetOccurredAt() int64
	GetOccurredAtTime() time.Time
}

// Generated event types for each webhook event

// SubscriptionPauseScheduledEvent represents a subscription_pause_scheduled webhook event
type SubscriptionPauseScheduledEvent struct {
	Id         string                             `json:"id"`
	OccurredAt int64                              `json:"occurred_at"`
	EventType  string                             `json:"event_type"`
	Content    *SubscriptionPauseScheduledContent `json:"content"`
}

func (e *SubscriptionPauseScheduledEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionPauseScheduledEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionPauseScheduledEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionPauseScheduledEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// CustomerBusinessEntityChangedEvent represents a customer_business_entity_changed webhook event
type CustomerBusinessEntityChangedEvent struct {
	Id         string                                `json:"id"`
	OccurredAt int64                                 `json:"occurred_at"`
	EventType  string                                `json:"event_type"`
	Content    *CustomerBusinessEntityChangedContent `json:"content"`
}

func (e *CustomerBusinessEntityChangedEvent) GetEventType() string { return e.EventType }
func (e *CustomerBusinessEntityChangedEvent) GetEventId() string   { return e.Id }
func (e *CustomerBusinessEntityChangedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *CustomerBusinessEntityChangedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// SubscriptionAdvanceInvoiceScheduleAddedEvent represents a subscription_advance_invoice_schedule_added webhook event
type SubscriptionAdvanceInvoiceScheduleAddedEvent struct {
	Id         string                                          `json:"id"`
	OccurredAt int64                                           `json:"occurred_at"`
	EventType  string                                          `json:"event_type"`
	Content    *SubscriptionAdvanceInvoiceScheduleAddedContent `json:"content"`
}

func (e *SubscriptionAdvanceInvoiceScheduleAddedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionAdvanceInvoiceScheduleAddedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionAdvanceInvoiceScheduleAddedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionAdvanceInvoiceScheduleAddedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// GiftExpiredEvent represents a gift_expired webhook event
type GiftExpiredEvent struct {
	Id         string              `json:"id"`
	OccurredAt int64               `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *GiftExpiredContent `json:"content"`
}

func (e *GiftExpiredEvent) GetEventType() string         { return e.EventType }
func (e *GiftExpiredEvent) GetEventId() string           { return e.Id }
func (e *GiftExpiredEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *GiftExpiredEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// TaxWithheldDeletedEvent represents a tax_withheld_deleted webhook event
type TaxWithheldDeletedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt int64                      `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *TaxWithheldDeletedContent `json:"content"`
}

func (e *TaxWithheldDeletedEvent) GetEventType() string         { return e.EventType }
func (e *TaxWithheldDeletedEvent) GetEventId() string           { return e.Id }
func (e *TaxWithheldDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *TaxWithheldDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// UnbilledChargesDeletedEvent represents a unbilled_charges_deleted webhook event
type UnbilledChargesDeletedEvent struct {
	Id         string                         `json:"id"`
	OccurredAt int64                          `json:"occurred_at"`
	EventType  string                         `json:"event_type"`
	Content    *UnbilledChargesDeletedContent `json:"content"`
}

func (e *UnbilledChargesDeletedEvent) GetEventType() string { return e.EventType }
func (e *UnbilledChargesDeletedEvent) GetEventId() string   { return e.Id }
func (e *UnbilledChargesDeletedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *UnbilledChargesDeletedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// CouponUpdatedEvent represents a coupon_updated webhook event
type CouponUpdatedEvent struct {
	Id         string                `json:"id"`
	OccurredAt int64                 `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *CouponUpdatedContent `json:"content"`
}

func (e *CouponUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *CouponUpdatedEvent) GetEventId() string           { return e.Id }
func (e *CouponUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CouponUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionItemReactivatedEvent represents a omnichannel_subscription_item_reactivated webhook event
type OmnichannelSubscriptionItemReactivatedEvent struct {
	Id         string                                         `json:"id"`
	OccurredAt int64                                          `json:"occurred_at"`
	EventType  string                                         `json:"event_type"`
	Content    *OmnichannelSubscriptionItemReactivatedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemReactivatedEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionItemReactivatedEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionItemReactivatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelSubscriptionItemReactivatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// OmnichannelSubscriptionItemRenewedEvent represents a omnichannel_subscription_item_renewed webhook event
type OmnichannelSubscriptionItemRenewedEvent struct {
	Id         string                                     `json:"id"`
	OccurredAt int64                                      `json:"occurred_at"`
	EventType  string                                     `json:"event_type"`
	Content    *OmnichannelSubscriptionItemRenewedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemRenewedEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionItemRenewedEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionItemRenewedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelSubscriptionItemRenewedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// UnbilledChargesCreatedEvent represents a unbilled_charges_created webhook event
type UnbilledChargesCreatedEvent struct {
	Id         string                         `json:"id"`
	OccurredAt int64                          `json:"occurred_at"`
	EventType  string                         `json:"event_type"`
	Content    *UnbilledChargesCreatedContent `json:"content"`
}

func (e *UnbilledChargesCreatedEvent) GetEventType() string { return e.EventType }
func (e *UnbilledChargesCreatedEvent) GetEventId() string   { return e.Id }
func (e *UnbilledChargesCreatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *UnbilledChargesCreatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// SubscriptionResumedEvent represents a subscription_resumed webhook event
type SubscriptionResumedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *SubscriptionResumedContent `json:"content"`
}

func (e *SubscriptionResumedEvent) GetEventType() string         { return e.EventType }
func (e *SubscriptionResumedEvent) GetEventId() string           { return e.Id }
func (e *SubscriptionResumedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *SubscriptionResumedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelOneTimeOrderItemCancelledEvent represents a omnichannel_one_time_order_item_cancelled webhook event
type OmnichannelOneTimeOrderItemCancelledEvent struct {
	Id         string                                       `json:"id"`
	OccurredAt int64                                        `json:"occurred_at"`
	EventType  string                                       `json:"event_type"`
	Content    *OmnichannelOneTimeOrderItemCancelledContent `json:"content"`
}

func (e *OmnichannelOneTimeOrderItemCancelledEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelOneTimeOrderItemCancelledEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelOneTimeOrderItemCancelledEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelOneTimeOrderItemCancelledEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// SubscriptionCancelledEvent represents a subscription_cancelled webhook event
type SubscriptionCancelledEvent struct {
	Id         string                        `json:"id"`
	OccurredAt int64                         `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *SubscriptionCancelledContent `json:"content"`
}

func (e *SubscriptionCancelledEvent) GetEventType() string         { return e.EventType }
func (e *SubscriptionCancelledEvent) GetEventId() string           { return e.Id }
func (e *SubscriptionCancelledEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *SubscriptionCancelledEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ItemEntitlementsRemovedEvent represents a item_entitlements_removed webhook event
type ItemEntitlementsRemovedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *ItemEntitlementsRemovedContent `json:"content"`
}

func (e *ItemEntitlementsRemovedEvent) GetEventType() string { return e.EventType }
func (e *ItemEntitlementsRemovedEvent) GetEventId() string   { return e.Id }
func (e *ItemEntitlementsRemovedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *ItemEntitlementsRemovedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// BusinessEntityCreatedEvent represents a business_entity_created webhook event
type BusinessEntityCreatedEvent struct {
	Id         string                        `json:"id"`
	OccurredAt int64                         `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *BusinessEntityCreatedContent `json:"content"`
}

func (e *BusinessEntityCreatedEvent) GetEventType() string         { return e.EventType }
func (e *BusinessEntityCreatedEvent) GetEventId() string           { return e.Id }
func (e *BusinessEntityCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *BusinessEntityCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CouponSetUpdatedEvent represents a coupon_set_updated webhook event
type CouponSetUpdatedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *CouponSetUpdatedContent `json:"content"`
}

func (e *CouponSetUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *CouponSetUpdatedEvent) GetEventId() string           { return e.Id }
func (e *CouponSetUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CouponSetUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// DifferentialPriceUpdatedEvent represents a differential_price_updated webhook event
type DifferentialPriceUpdatedEvent struct {
	Id         string                           `json:"id"`
	OccurredAt int64                            `json:"occurred_at"`
	EventType  string                           `json:"event_type"`
	Content    *DifferentialPriceUpdatedContent `json:"content"`
}

func (e *DifferentialPriceUpdatedEvent) GetEventType() string { return e.EventType }
func (e *DifferentialPriceUpdatedEvent) GetEventId() string   { return e.Id }
func (e *DifferentialPriceUpdatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *DifferentialPriceUpdatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// OmnichannelSubscriptionItemPausedEvent represents a omnichannel_subscription_item_paused webhook event
type OmnichannelSubscriptionItemPausedEvent struct {
	Id         string                                    `json:"id"`
	OccurredAt int64                                     `json:"occurred_at"`
	EventType  string                                    `json:"event_type"`
	Content    *OmnichannelSubscriptionItemPausedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemPausedEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionItemPausedEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionItemPausedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelSubscriptionItemPausedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// EntitlementOverridesRemovedEvent represents a entitlement_overrides_removed webhook event
type EntitlementOverridesRemovedEvent struct {
	Id         string                              `json:"id"`
	OccurredAt int64                               `json:"occurred_at"`
	EventType  string                              `json:"event_type"`
	Content    *EntitlementOverridesRemovedContent `json:"content"`
}

func (e *EntitlementOverridesRemovedEvent) GetEventType() string { return e.EventType }
func (e *EntitlementOverridesRemovedEvent) GetEventId() string   { return e.Id }
func (e *EntitlementOverridesRemovedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *EntitlementOverridesRemovedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// SubscriptionActivatedWithBackdatingEvent represents a subscription_activated_with_backdating webhook event
type SubscriptionActivatedWithBackdatingEvent struct {
	Id         string                                      `json:"id"`
	OccurredAt int64                                       `json:"occurred_at"`
	EventType  string                                      `json:"event_type"`
	Content    *SubscriptionActivatedWithBackdatingContent `json:"content"`
}

func (e *SubscriptionActivatedWithBackdatingEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionActivatedWithBackdatingEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionActivatedWithBackdatingEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionActivatedWithBackdatingEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// SubscriptionTrialEndReminderEvent represents a subscription_trial_end_reminder webhook event
type SubscriptionTrialEndReminderEvent struct {
	Id         string                               `json:"id"`
	OccurredAt int64                                `json:"occurred_at"`
	EventType  string                               `json:"event_type"`
	Content    *SubscriptionTrialEndReminderContent `json:"content"`
}

func (e *SubscriptionTrialEndReminderEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionTrialEndReminderEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionTrialEndReminderEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionTrialEndReminderEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// SubscriptionShippingAddressUpdatedEvent represents a subscription_shipping_address_updated webhook event
type SubscriptionShippingAddressUpdatedEvent struct {
	Id         string                                     `json:"id"`
	OccurredAt int64                                      `json:"occurred_at"`
	EventType  string                                     `json:"event_type"`
	Content    *SubscriptionShippingAddressUpdatedContent `json:"content"`
}

func (e *SubscriptionShippingAddressUpdatedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionShippingAddressUpdatedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionShippingAddressUpdatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionShippingAddressUpdatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// VoucherCreateFailedEvent represents a voucher_create_failed webhook event
type VoucherCreateFailedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *VoucherCreateFailedContent `json:"content"`
}

func (e *VoucherCreateFailedEvent) GetEventType() string         { return e.EventType }
func (e *VoucherCreateFailedEvent) GetEventId() string           { return e.Id }
func (e *VoucherCreateFailedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *VoucherCreateFailedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// GiftClaimedEvent represents a gift_claimed webhook event
type GiftClaimedEvent struct {
	Id         string              `json:"id"`
	OccurredAt int64               `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *GiftClaimedContent `json:"content"`
}

func (e *GiftClaimedEvent) GetEventType() string         { return e.EventType }
func (e *GiftClaimedEvent) GetEventId() string           { return e.Id }
func (e *GiftClaimedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *GiftClaimedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CustomerDeletedEvent represents a customer_deleted webhook event
type CustomerDeletedEvent struct {
	Id         string                  `json:"id"`
	OccurredAt int64                   `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *CustomerDeletedContent `json:"content"`
}

func (e *CustomerDeletedEvent) GetEventType() string         { return e.EventType }
func (e *CustomerDeletedEvent) GetEventId() string           { return e.Id }
func (e *CustomerDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CustomerDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// RefundInitiatedEvent represents a refund_initiated webhook event
type RefundInitiatedEvent struct {
	Id         string                  `json:"id"`
	OccurredAt int64                   `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *RefundInitiatedContent `json:"content"`
}

func (e *RefundInitiatedEvent) GetEventType() string         { return e.EventType }
func (e *RefundInitiatedEvent) GetEventId() string           { return e.Id }
func (e *RefundInitiatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *RefundInitiatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// InvoiceGeneratedWithBackdatingEvent represents a invoice_generated_with_backdating webhook event
type InvoiceGeneratedWithBackdatingEvent struct {
	Id         string                                 `json:"id"`
	OccurredAt int64                                  `json:"occurred_at"`
	EventType  string                                 `json:"event_type"`
	Content    *InvoiceGeneratedWithBackdatingContent `json:"content"`
}

func (e *InvoiceGeneratedWithBackdatingEvent) GetEventType() string { return e.EventType }
func (e *InvoiceGeneratedWithBackdatingEvent) GetEventId() string   { return e.Id }
func (e *InvoiceGeneratedWithBackdatingEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *InvoiceGeneratedWithBackdatingEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// OmnichannelTransactionCreatedEvent represents a omnichannel_transaction_created webhook event
type OmnichannelTransactionCreatedEvent struct {
	Id         string                                `json:"id"`
	OccurredAt int64                                 `json:"occurred_at"`
	EventType  string                                `json:"event_type"`
	Content    *OmnichannelTransactionCreatedContent `json:"content"`
}

func (e *OmnichannelTransactionCreatedEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelTransactionCreatedEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelTransactionCreatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelTransactionCreatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// AddUsagesReminderEvent represents a add_usages_reminder webhook event
type AddUsagesReminderEvent struct {
	Id         string                    `json:"id"`
	OccurredAt int64                     `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *AddUsagesReminderContent `json:"content"`
}

func (e *AddUsagesReminderEvent) GetEventType() string         { return e.EventType }
func (e *AddUsagesReminderEvent) GetEventId() string           { return e.Id }
func (e *AddUsagesReminderEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *AddUsagesReminderEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// VoucherCreatedEvent represents a voucher_created webhook event
type VoucherCreatedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt int64                  `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *VoucherCreatedContent `json:"content"`
}

func (e *VoucherCreatedEvent) GetEventType() string         { return e.EventType }
func (e *VoucherCreatedEvent) GetEventId() string           { return e.Id }
func (e *VoucherCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *VoucherCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// RuleUpdatedEvent represents a rule_updated webhook event
type RuleUpdatedEvent struct {
	Id         string              `json:"id"`
	OccurredAt int64               `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *RuleUpdatedContent `json:"content"`
}

func (e *RuleUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *RuleUpdatedEvent) GetEventId() string           { return e.Id }
func (e *RuleUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *RuleUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentSchedulesCreatedEvent represents a payment_schedules_created webhook event
type PaymentSchedulesCreatedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *PaymentSchedulesCreatedContent `json:"content"`
}

func (e *PaymentSchedulesCreatedEvent) GetEventType() string { return e.EventType }
func (e *PaymentSchedulesCreatedEvent) GetEventId() string   { return e.Id }
func (e *PaymentSchedulesCreatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *PaymentSchedulesCreatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// FeatureActivatedEvent represents a feature_activated webhook event
type FeatureActivatedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *FeatureActivatedContent `json:"content"`
}

func (e *FeatureActivatedEvent) GetEventType() string         { return e.EventType }
func (e *FeatureActivatedEvent) GetEventId() string           { return e.Id }
func (e *FeatureActivatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *FeatureActivatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentSourceLocallyDeletedEvent represents a payment_source_locally_deleted webhook event
type PaymentSourceLocallyDeletedEvent struct {
	Id         string                              `json:"id"`
	OccurredAt int64                               `json:"occurred_at"`
	EventType  string                              `json:"event_type"`
	Content    *PaymentSourceLocallyDeletedContent `json:"content"`
}

func (e *PaymentSourceLocallyDeletedEvent) GetEventType() string { return e.EventType }
func (e *PaymentSourceLocallyDeletedEvent) GetEventId() string   { return e.Id }
func (e *PaymentSourceLocallyDeletedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *PaymentSourceLocallyDeletedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// InvoiceGeneratedEvent represents a invoice_generated webhook event
type InvoiceGeneratedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *InvoiceGeneratedContent `json:"content"`
}

func (e *InvoiceGeneratedEvent) GetEventType() string         { return e.EventType }
func (e *InvoiceGeneratedEvent) GetEventId() string           { return e.Id }
func (e *InvoiceGeneratedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *InvoiceGeneratedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// VoucherExpiredEvent represents a voucher_expired webhook event
type VoucherExpiredEvent struct {
	Id         string                 `json:"id"`
	OccurredAt int64                  `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *VoucherExpiredContent `json:"content"`
}

func (e *VoucherExpiredEvent) GetEventType() string         { return e.EventType }
func (e *VoucherExpiredEvent) GetEventId() string           { return e.Id }
func (e *VoucherExpiredEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *VoucherExpiredEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// AuthorizationSucceededEvent represents a authorization_succeeded webhook event
type AuthorizationSucceededEvent struct {
	Id         string                         `json:"id"`
	OccurredAt int64                          `json:"occurred_at"`
	EventType  string                         `json:"event_type"`
	Content    *AuthorizationSucceededContent `json:"content"`
}

func (e *AuthorizationSucceededEvent) GetEventType() string { return e.EventType }
func (e *AuthorizationSucceededEvent) GetEventId() string   { return e.Id }
func (e *AuthorizationSucceededEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *AuthorizationSucceededEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// GiftScheduledEvent represents a gift_scheduled webhook event
type GiftScheduledEvent struct {
	Id         string                `json:"id"`
	OccurredAt int64                 `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *GiftScheduledContent `json:"content"`
}

func (e *GiftScheduledEvent) GetEventType() string         { return e.EventType }
func (e *GiftScheduledEvent) GetEventId() string           { return e.Id }
func (e *GiftScheduledEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *GiftScheduledEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionChangesScheduledEvent represents a subscription_changes_scheduled webhook event
type SubscriptionChangesScheduledEvent struct {
	Id         string                               `json:"id"`
	OccurredAt int64                                `json:"occurred_at"`
	EventType  string                               `json:"event_type"`
	Content    *SubscriptionChangesScheduledContent `json:"content"`
}

func (e *SubscriptionChangesScheduledEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionChangesScheduledEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionChangesScheduledEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionChangesScheduledEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// SubscriptionChangedWithBackdatingEvent represents a subscription_changed_with_backdating webhook event
type SubscriptionChangedWithBackdatingEvent struct {
	Id         string                                    `json:"id"`
	OccurredAt int64                                     `json:"occurred_at"`
	EventType  string                                    `json:"event_type"`
	Content    *SubscriptionChangedWithBackdatingContent `json:"content"`
}

func (e *SubscriptionChangedWithBackdatingEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionChangedWithBackdatingEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionChangedWithBackdatingEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionChangedWithBackdatingEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// OmnichannelSubscriptionItemChangedEvent represents a omnichannel_subscription_item_changed webhook event
type OmnichannelSubscriptionItemChangedEvent struct {
	Id         string                                     `json:"id"`
	OccurredAt int64                                      `json:"occurred_at"`
	EventType  string                                     `json:"event_type"`
	Content    *OmnichannelSubscriptionItemChangedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemChangedEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionItemChangedEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionItemChangedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelSubscriptionItemChangedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// GiftUnclaimedEvent represents a gift_unclaimed webhook event
type GiftUnclaimedEvent struct {
	Id         string                `json:"id"`
	OccurredAt int64                 `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *GiftUnclaimedContent `json:"content"`
}

func (e *GiftUnclaimedEvent) GetEventType() string         { return e.EventType }
func (e *GiftUnclaimedEvent) GetEventId() string           { return e.Id }
func (e *GiftUnclaimedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *GiftUnclaimedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// VirtualBankAccountAddedEvent represents a virtual_bank_account_added webhook event
type VirtualBankAccountAddedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *VirtualBankAccountAddedContent `json:"content"`
}

func (e *VirtualBankAccountAddedEvent) GetEventType() string { return e.EventType }
func (e *VirtualBankAccountAddedEvent) GetEventId() string   { return e.Id }
func (e *VirtualBankAccountAddedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *VirtualBankAccountAddedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// PaymentIntentCreatedEvent represents a payment_intent_created webhook event
type PaymentIntentCreatedEvent struct {
	Id         string                       `json:"id"`
	OccurredAt int64                        `json:"occurred_at"`
	EventType  string                       `json:"event_type"`
	Content    *PaymentIntentCreatedContent `json:"content"`
}

func (e *PaymentIntentCreatedEvent) GetEventType() string         { return e.EventType }
func (e *PaymentIntentCreatedEvent) GetEventId() string           { return e.Id }
func (e *PaymentIntentCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *PaymentIntentCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CreditNoteCreatedWithBackdatingEvent represents a credit_note_created_with_backdating webhook event
type CreditNoteCreatedWithBackdatingEvent struct {
	Id         string                                  `json:"id"`
	OccurredAt int64                                   `json:"occurred_at"`
	EventType  string                                  `json:"event_type"`
	Content    *CreditNoteCreatedWithBackdatingContent `json:"content"`
}

func (e *CreditNoteCreatedWithBackdatingEvent) GetEventType() string { return e.EventType }
func (e *CreditNoteCreatedWithBackdatingEvent) GetEventId() string   { return e.Id }
func (e *CreditNoteCreatedWithBackdatingEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *CreditNoteCreatedWithBackdatingEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// ContractTermTerminatedEvent represents a contract_term_terminated webhook event
type ContractTermTerminatedEvent struct {
	Id         string                         `json:"id"`
	OccurredAt int64                          `json:"occurred_at"`
	EventType  string                         `json:"event_type"`
	Content    *ContractTermTerminatedContent `json:"content"`
}

func (e *ContractTermTerminatedEvent) GetEventType() string { return e.EventType }
func (e *ContractTermTerminatedEvent) GetEventId() string   { return e.Id }
func (e *ContractTermTerminatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *ContractTermTerminatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// ItemFamilyUpdatedEvent represents a item_family_updated webhook event
type ItemFamilyUpdatedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt int64                     `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *ItemFamilyUpdatedContent `json:"content"`
}

func (e *ItemFamilyUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *ItemFamilyUpdatedEvent) GetEventId() string           { return e.Id }
func (e *ItemFamilyUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *ItemFamilyUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OrderCreatedEvent represents a order_created webhook event
type OrderCreatedEvent struct {
	Id         string               `json:"id"`
	OccurredAt int64                `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *OrderCreatedContent `json:"content"`
}

func (e *OrderCreatedEvent) GetEventType() string         { return e.EventType }
func (e *OrderCreatedEvent) GetEventId() string           { return e.Id }
func (e *OrderCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *OrderCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PriceVariantDeletedEvent represents a price_variant_deleted webhook event
type PriceVariantDeletedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *PriceVariantDeletedContent `json:"content"`
}

func (e *PriceVariantDeletedEvent) GetEventType() string         { return e.EventType }
func (e *PriceVariantDeletedEvent) GetEventId() string           { return e.Id }
func (e *PriceVariantDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *PriceVariantDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionMovementFailedEvent represents a subscription_movement_failed webhook event
type SubscriptionMovementFailedEvent struct {
	Id         string                             `json:"id"`
	OccurredAt int64                              `json:"occurred_at"`
	EventType  string                             `json:"event_type"`
	Content    *SubscriptionMovementFailedContent `json:"content"`
}

func (e *SubscriptionMovementFailedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionMovementFailedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionMovementFailedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionMovementFailedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// CustomerMovedInEvent represents a customer_moved_in webhook event
type CustomerMovedInEvent struct {
	Id         string                  `json:"id"`
	OccurredAt int64                   `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *CustomerMovedInContent `json:"content"`
}

func (e *CustomerMovedInEvent) GetEventType() string         { return e.EventType }
func (e *CustomerMovedInEvent) GetEventId() string           { return e.Id }
func (e *CustomerMovedInEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CustomerMovedInEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionAdvanceInvoiceScheduleUpdatedEvent represents a subscription_advance_invoice_schedule_updated webhook event
type SubscriptionAdvanceInvoiceScheduleUpdatedEvent struct {
	Id         string                                            `json:"id"`
	OccurredAt int64                                             `json:"occurred_at"`
	EventType  string                                            `json:"event_type"`
	Content    *SubscriptionAdvanceInvoiceScheduleUpdatedContent `json:"content"`
}

func (e *SubscriptionAdvanceInvoiceScheduleUpdatedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionAdvanceInvoiceScheduleUpdatedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionAdvanceInvoiceScheduleUpdatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionAdvanceInvoiceScheduleUpdatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// ItemDeletedEvent represents a item_deleted webhook event
type ItemDeletedEvent struct {
	Id         string              `json:"id"`
	OccurredAt int64               `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *ItemDeletedContent `json:"content"`
}

func (e *ItemDeletedEvent) GetEventType() string         { return e.EventType }
func (e *ItemDeletedEvent) GetEventId() string           { return e.Id }
func (e *ItemDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *ItemDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionRampDraftedEvent represents a subscription_ramp_drafted webhook event
type SubscriptionRampDraftedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *SubscriptionRampDraftedContent `json:"content"`
}

func (e *SubscriptionRampDraftedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionRampDraftedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionRampDraftedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionRampDraftedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// DunningUpdatedEvent represents a dunning_updated webhook event
type DunningUpdatedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt int64                  `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *DunningUpdatedContent `json:"content"`
}

func (e *DunningUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *DunningUpdatedEvent) GetEventId() string           { return e.Id }
func (e *DunningUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *DunningUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ItemEntitlementsUpdatedEvent represents a item_entitlements_updated webhook event
type ItemEntitlementsUpdatedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *ItemEntitlementsUpdatedContent `json:"content"`
}

func (e *ItemEntitlementsUpdatedEvent) GetEventType() string { return e.EventType }
func (e *ItemEntitlementsUpdatedEvent) GetEventId() string   { return e.Id }
func (e *ItemEntitlementsUpdatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *ItemEntitlementsUpdatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// TokenConsumedEvent represents a token_consumed webhook event
type TokenConsumedEvent struct {
	Id         string                `json:"id"`
	OccurredAt int64                 `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *TokenConsumedContent `json:"content"`
}

func (e *TokenConsumedEvent) GetEventType() string         { return e.EventType }
func (e *TokenConsumedEvent) GetEventId() string           { return e.Id }
func (e *TokenConsumedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *TokenConsumedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// HierarchyDeletedEvent represents a hierarchy_deleted webhook event
type HierarchyDeletedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *HierarchyDeletedContent `json:"content"`
}

func (e *HierarchyDeletedEvent) GetEventType() string         { return e.EventType }
func (e *HierarchyDeletedEvent) GetEventId() string           { return e.Id }
func (e *HierarchyDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *HierarchyDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionCancellationScheduledEvent represents a subscription_cancellation_scheduled webhook event
type SubscriptionCancellationScheduledEvent struct {
	Id         string                                    `json:"id"`
	OccurredAt int64                                     `json:"occurred_at"`
	EventType  string                                    `json:"event_type"`
	Content    *SubscriptionCancellationScheduledContent `json:"content"`
}

func (e *SubscriptionCancellationScheduledEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionCancellationScheduledEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionCancellationScheduledEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionCancellationScheduledEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// SubscriptionRenewedEvent represents a subscription_renewed webhook event
type SubscriptionRenewedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *SubscriptionRenewedContent `json:"content"`
}

func (e *SubscriptionRenewedEvent) GetEventType() string         { return e.EventType }
func (e *SubscriptionRenewedEvent) GetEventId() string           { return e.Id }
func (e *SubscriptionRenewedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *SubscriptionRenewedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// FeatureUpdatedEvent represents a feature_updated webhook event
type FeatureUpdatedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt int64                  `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *FeatureUpdatedContent `json:"content"`
}

func (e *FeatureUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *FeatureUpdatedEvent) GetEventId() string           { return e.Id }
func (e *FeatureUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *FeatureUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// FeatureDeletedEvent represents a feature_deleted webhook event
type FeatureDeletedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt int64                  `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *FeatureDeletedContent `json:"content"`
}

func (e *FeatureDeletedEvent) GetEventType() string         { return e.EventType }
func (e *FeatureDeletedEvent) GetEventId() string           { return e.Id }
func (e *FeatureDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *FeatureDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ItemFamilyCreatedEvent represents a item_family_created webhook event
type ItemFamilyCreatedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt int64                     `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *ItemFamilyCreatedContent `json:"content"`
}

func (e *ItemFamilyCreatedEvent) GetEventType() string         { return e.EventType }
func (e *ItemFamilyCreatedEvent) GetEventId() string           { return e.Id }
func (e *ItemFamilyCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *ItemFamilyCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionItemScheduledChangeRemovedEvent represents a omnichannel_subscription_item_scheduled_change_removed webhook event
type OmnichannelSubscriptionItemScheduledChangeRemovedEvent struct {
	Id         string                                                    `json:"id"`
	OccurredAt int64                                                     `json:"occurred_at"`
	EventType  string                                                    `json:"event_type"`
	Content    *OmnichannelSubscriptionItemScheduledChangeRemovedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemScheduledChangeRemovedEvent) GetEventType() string {
	return e.EventType
}
func (e *OmnichannelSubscriptionItemScheduledChangeRemovedEvent) GetEventId() string { return e.Id }
func (e *OmnichannelSubscriptionItemScheduledChangeRemovedEvent) GetOccurredAt() int64 {
	return e.OccurredAt
}
func (e *OmnichannelSubscriptionItemScheduledChangeRemovedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// OmnichannelSubscriptionItemResumedEvent represents a omnichannel_subscription_item_resumed webhook event
type OmnichannelSubscriptionItemResumedEvent struct {
	Id         string                                     `json:"id"`
	OccurredAt int64                                      `json:"occurred_at"`
	EventType  string                                     `json:"event_type"`
	Content    *OmnichannelSubscriptionItemResumedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemResumedEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionItemResumedEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionItemResumedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelSubscriptionItemResumedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// PurchaseCreatedEvent represents a purchase_created webhook event
type PurchaseCreatedEvent struct {
	Id         string                  `json:"id"`
	OccurredAt int64                   `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *PurchaseCreatedContent `json:"content"`
}

func (e *PurchaseCreatedEvent) GetEventType() string         { return e.EventType }
func (e *PurchaseCreatedEvent) GetEventId() string           { return e.Id }
func (e *PurchaseCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *PurchaseCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// EntitlementOverridesUpdatedEvent represents a entitlement_overrides_updated webhook event
type EntitlementOverridesUpdatedEvent struct {
	Id         string                              `json:"id"`
	OccurredAt int64                               `json:"occurred_at"`
	EventType  string                              `json:"event_type"`
	Content    *EntitlementOverridesUpdatedContent `json:"content"`
}

func (e *EntitlementOverridesUpdatedEvent) GetEventType() string { return e.EventType }
func (e *EntitlementOverridesUpdatedEvent) GetEventId() string   { return e.Id }
func (e *EntitlementOverridesUpdatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *EntitlementOverridesUpdatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// ItemFamilyDeletedEvent represents a item_family_deleted webhook event
type ItemFamilyDeletedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt int64                     `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *ItemFamilyDeletedContent `json:"content"`
}

func (e *ItemFamilyDeletedEvent) GetEventType() string         { return e.EventType }
func (e *ItemFamilyDeletedEvent) GetEventId() string           { return e.Id }
func (e *ItemFamilyDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *ItemFamilyDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionResumptionScheduledEvent represents a subscription_resumption_scheduled webhook event
type SubscriptionResumptionScheduledEvent struct {
	Id         string                                  `json:"id"`
	OccurredAt int64                                   `json:"occurred_at"`
	EventType  string                                  `json:"event_type"`
	Content    *SubscriptionResumptionScheduledContent `json:"content"`
}

func (e *SubscriptionResumptionScheduledEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionResumptionScheduledEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionResumptionScheduledEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionResumptionScheduledEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// FeatureReactivatedEvent represents a feature_reactivated webhook event
type FeatureReactivatedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt int64                      `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *FeatureReactivatedContent `json:"content"`
}

func (e *FeatureReactivatedEvent) GetEventType() string         { return e.EventType }
func (e *FeatureReactivatedEvent) GetEventId() string           { return e.Id }
func (e *FeatureReactivatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *FeatureReactivatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CouponCodesDeletedEvent represents a coupon_codes_deleted webhook event
type CouponCodesDeletedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt int64                      `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *CouponCodesDeletedContent `json:"content"`
}

func (e *CouponCodesDeletedEvent) GetEventType() string         { return e.EventType }
func (e *CouponCodesDeletedEvent) GetEventId() string           { return e.Id }
func (e *CouponCodesDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CouponCodesDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CardExpiredEvent represents a card_expired webhook event
type CardExpiredEvent struct {
	Id         string              `json:"id"`
	OccurredAt int64               `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *CardExpiredContent `json:"content"`
}

func (e *CardExpiredEvent) GetEventType() string         { return e.EventType }
func (e *CardExpiredEvent) GetEventId() string           { return e.Id }
func (e *CardExpiredEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CardExpiredEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CreditNoteUpdatedEvent represents a credit_note_updated webhook event
type CreditNoteUpdatedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt int64                     `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *CreditNoteUpdatedContent `json:"content"`
}

func (e *CreditNoteUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *CreditNoteUpdatedEvent) GetEventId() string           { return e.Id }
func (e *CreditNoteUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CreditNoteUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionItemDowngradedEvent represents a omnichannel_subscription_item_downgraded webhook event
type OmnichannelSubscriptionItemDowngradedEvent struct {
	Id         string                                        `json:"id"`
	OccurredAt int64                                         `json:"occurred_at"`
	EventType  string                                        `json:"event_type"`
	Content    *OmnichannelSubscriptionItemDowngradedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemDowngradedEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionItemDowngradedEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionItemDowngradedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelSubscriptionItemDowngradedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// PriceVariantUpdatedEvent represents a price_variant_updated webhook event
type PriceVariantUpdatedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *PriceVariantUpdatedContent `json:"content"`
}

func (e *PriceVariantUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *PriceVariantUpdatedEvent) GetEventId() string           { return e.Id }
func (e *PriceVariantUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *PriceVariantUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PromotionalCreditsDeductedEvent represents a promotional_credits_deducted webhook event
type PromotionalCreditsDeductedEvent struct {
	Id         string                             `json:"id"`
	OccurredAt int64                              `json:"occurred_at"`
	EventType  string                             `json:"event_type"`
	Content    *PromotionalCreditsDeductedContent `json:"content"`
}

func (e *PromotionalCreditsDeductedEvent) GetEventType() string { return e.EventType }
func (e *PromotionalCreditsDeductedEvent) GetEventId() string   { return e.Id }
func (e *PromotionalCreditsDeductedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *PromotionalCreditsDeductedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// SubscriptionRampAppliedEvent represents a subscription_ramp_applied webhook event
type SubscriptionRampAppliedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *SubscriptionRampAppliedContent `json:"content"`
}

func (e *SubscriptionRampAppliedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionRampAppliedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionRampAppliedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionRampAppliedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// SubscriptionPausedEvent represents a subscription_paused webhook event
type SubscriptionPausedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt int64                      `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *SubscriptionPausedContent `json:"content"`
}

func (e *SubscriptionPausedEvent) GetEventType() string         { return e.EventType }
func (e *SubscriptionPausedEvent) GetEventId() string           { return e.Id }
func (e *SubscriptionPausedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *SubscriptionPausedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OrderReadyToProcessEvent represents a order_ready_to_process webhook event
type OrderReadyToProcessEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *OrderReadyToProcessContent `json:"content"`
}

func (e *OrderReadyToProcessEvent) GetEventType() string         { return e.EventType }
func (e *OrderReadyToProcessEvent) GetEventId() string           { return e.Id }
func (e *OrderReadyToProcessEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *OrderReadyToProcessEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// FeatureCreatedEvent represents a feature_created webhook event
type FeatureCreatedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt int64                  `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *FeatureCreatedContent `json:"content"`
}

func (e *FeatureCreatedEvent) GetEventType() string         { return e.EventType }
func (e *FeatureCreatedEvent) GetEventId() string           { return e.Id }
func (e *FeatureCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *FeatureCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// TransactionDeletedEvent represents a transaction_deleted webhook event
type TransactionDeletedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt int64                      `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *TransactionDeletedContent `json:"content"`
}

func (e *TransactionDeletedEvent) GetEventType() string         { return e.EventType }
func (e *TransactionDeletedEvent) GetEventId() string           { return e.Id }
func (e *TransactionDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *TransactionDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CreditNoteCreatedEvent represents a credit_note_created webhook event
type CreditNoteCreatedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt int64                     `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *CreditNoteCreatedContent `json:"content"`
}

func (e *CreditNoteCreatedEvent) GetEventType() string         { return e.EventType }
func (e *CreditNoteCreatedEvent) GetEventId() string           { return e.Id }
func (e *CreditNoteCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CreditNoteCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionItemResubscribedEvent represents a omnichannel_subscription_item_resubscribed webhook event
type OmnichannelSubscriptionItemResubscribedEvent struct {
	Id         string                                          `json:"id"`
	OccurredAt int64                                           `json:"occurred_at"`
	EventType  string                                          `json:"event_type"`
	Content    *OmnichannelSubscriptionItemResubscribedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemResubscribedEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionItemResubscribedEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionItemResubscribedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelSubscriptionItemResubscribedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// RecordPurchaseFailedEvent represents a record_purchase_failed webhook event
type RecordPurchaseFailedEvent struct {
	Id         string                       `json:"id"`
	OccurredAt int64                        `json:"occurred_at"`
	EventType  string                       `json:"event_type"`
	Content    *RecordPurchaseFailedContent `json:"content"`
}

func (e *RecordPurchaseFailedEvent) GetEventType() string         { return e.EventType }
func (e *RecordPurchaseFailedEvent) GetEventId() string           { return e.Id }
func (e *RecordPurchaseFailedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *RecordPurchaseFailedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ItemCreatedEvent represents a item_created webhook event
type ItemCreatedEvent struct {
	Id         string              `json:"id"`
	OccurredAt int64               `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *ItemCreatedContent `json:"content"`
}

func (e *ItemCreatedEvent) GetEventType() string         { return e.EventType }
func (e *ItemCreatedEvent) GetEventId() string           { return e.Id }
func (e *ItemCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *ItemCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// TransactionUpdatedEvent represents a transaction_updated webhook event
type TransactionUpdatedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt int64                      `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *TransactionUpdatedContent `json:"content"`
}

func (e *TransactionUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *TransactionUpdatedEvent) GetEventId() string           { return e.Id }
func (e *TransactionUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *TransactionUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// MrrUpdatedEvent represents a mrr_updated webhook event
type MrrUpdatedEvent struct {
	Id         string             `json:"id"`
	OccurredAt int64              `json:"occurred_at"`
	EventType  string             `json:"event_type"`
	Content    *MrrUpdatedContent `json:"content"`
}

func (e *MrrUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *MrrUpdatedEvent) GetEventId() string           { return e.Id }
func (e *MrrUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *MrrUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// UnbilledChargesInvoicedEvent represents a unbilled_charges_invoiced webhook event
type UnbilledChargesInvoicedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *UnbilledChargesInvoicedContent `json:"content"`
}

func (e *UnbilledChargesInvoicedEvent) GetEventType() string { return e.EventType }
func (e *UnbilledChargesInvoicedEvent) GetEventId() string   { return e.Id }
func (e *UnbilledChargesInvoicedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *UnbilledChargesInvoicedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// ItemPriceUpdatedEvent represents a item_price_updated webhook event
type ItemPriceUpdatedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *ItemPriceUpdatedContent `json:"content"`
}

func (e *ItemPriceUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *ItemPriceUpdatedEvent) GetEventId() string           { return e.Id }
func (e *ItemPriceUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *ItemPriceUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CouponCodesUpdatedEvent represents a coupon_codes_updated webhook event
type CouponCodesUpdatedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt int64                      `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *CouponCodesUpdatedContent `json:"content"`
}

func (e *CouponCodesUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *CouponCodesUpdatedEvent) GetEventId() string           { return e.Id }
func (e *CouponCodesUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CouponCodesUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// VirtualBankAccountUpdatedEvent represents a virtual_bank_account_updated webhook event
type VirtualBankAccountUpdatedEvent struct {
	Id         string                            `json:"id"`
	OccurredAt int64                             `json:"occurred_at"`
	EventType  string                            `json:"event_type"`
	Content    *VirtualBankAccountUpdatedContent `json:"content"`
}

func (e *VirtualBankAccountUpdatedEvent) GetEventType() string { return e.EventType }
func (e *VirtualBankAccountUpdatedEvent) GetEventId() string   { return e.Id }
func (e *VirtualBankAccountUpdatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *VirtualBankAccountUpdatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// ContractTermCreatedEvent represents a contract_term_created webhook event
type ContractTermCreatedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *ContractTermCreatedContent `json:"content"`
}

func (e *ContractTermCreatedEvent) GetEventType() string         { return e.EventType }
func (e *ContractTermCreatedEvent) GetEventId() string           { return e.Id }
func (e *ContractTermCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *ContractTermCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionChangedEvent represents a subscription_changed webhook event
type SubscriptionChangedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *SubscriptionChangedContent `json:"content"`
}

func (e *SubscriptionChangedEvent) GetEventType() string         { return e.EventType }
func (e *SubscriptionChangedEvent) GetEventId() string           { return e.Id }
func (e *SubscriptionChangedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *SubscriptionChangedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentFailedEvent represents a payment_failed webhook event
type PaymentFailedEvent struct {
	Id         string                `json:"id"`
	OccurredAt int64                 `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *PaymentFailedContent `json:"content"`
}

func (e *PaymentFailedEvent) GetEventType() string         { return e.EventType }
func (e *PaymentFailedEvent) GetEventId() string           { return e.Id }
func (e *PaymentFailedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *PaymentFailedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CreditNoteDeletedEvent represents a credit_note_deleted webhook event
type CreditNoteDeletedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt int64                     `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *CreditNoteDeletedContent `json:"content"`
}

func (e *CreditNoteDeletedEvent) GetEventType() string         { return e.EventType }
func (e *CreditNoteDeletedEvent) GetEventId() string           { return e.Id }
func (e *CreditNoteDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CreditNoteDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// TaxWithheldRefundedEvent represents a tax_withheld_refunded webhook event
type TaxWithheldRefundedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *TaxWithheldRefundedContent `json:"content"`
}

func (e *TaxWithheldRefundedEvent) GetEventType() string         { return e.EventType }
func (e *TaxWithheldRefundedEvent) GetEventId() string           { return e.Id }
func (e *TaxWithheldRefundedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *TaxWithheldRefundedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ContractTermCompletedEvent represents a contract_term_completed webhook event
type ContractTermCompletedEvent struct {
	Id         string                        `json:"id"`
	OccurredAt int64                         `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *ContractTermCompletedContent `json:"content"`
}

func (e *ContractTermCompletedEvent) GetEventType() string         { return e.EventType }
func (e *ContractTermCompletedEvent) GetEventId() string           { return e.Id }
func (e *ContractTermCompletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *ContractTermCompletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentSchedulesUpdatedEvent represents a payment_schedules_updated webhook event
type PaymentSchedulesUpdatedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *PaymentSchedulesUpdatedContent `json:"content"`
}

func (e *PaymentSchedulesUpdatedEvent) GetEventType() string { return e.EventType }
func (e *PaymentSchedulesUpdatedEvent) GetEventId() string   { return e.Id }
func (e *PaymentSchedulesUpdatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *PaymentSchedulesUpdatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// OmnichannelSubscriptionItemExpiredEvent represents a omnichannel_subscription_item_expired webhook event
type OmnichannelSubscriptionItemExpiredEvent struct {
	Id         string                                     `json:"id"`
	OccurredAt int64                                      `json:"occurred_at"`
	EventType  string                                     `json:"event_type"`
	Content    *OmnichannelSubscriptionItemExpiredContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemExpiredEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionItemExpiredEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionItemExpiredEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelSubscriptionItemExpiredEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// CardUpdatedEvent represents a card_updated webhook event
type CardUpdatedEvent struct {
	Id         string              `json:"id"`
	OccurredAt int64               `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *CardUpdatedContent `json:"content"`
}

func (e *CardUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *CardUpdatedEvent) GetEventId() string           { return e.Id }
func (e *CardUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CardUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CustomerCreatedEvent represents a customer_created webhook event
type CustomerCreatedEvent struct {
	Id         string                  `json:"id"`
	OccurredAt int64                   `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *CustomerCreatedContent `json:"content"`
}

func (e *CustomerCreatedEvent) GetEventType() string         { return e.EventType }
func (e *CustomerCreatedEvent) GetEventId() string           { return e.Id }
func (e *CustomerCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CustomerCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionRenewalReminderEvent represents a subscription_renewal_reminder webhook event
type SubscriptionRenewalReminderEvent struct {
	Id         string                              `json:"id"`
	OccurredAt int64                               `json:"occurred_at"`
	EventType  string                              `json:"event_type"`
	Content    *SubscriptionRenewalReminderContent `json:"content"`
}

func (e *SubscriptionRenewalReminderEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionRenewalReminderEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionRenewalReminderEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionRenewalReminderEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// OrderDeliveredEvent represents a order_delivered webhook event
type OrderDeliveredEvent struct {
	Id         string                 `json:"id"`
	OccurredAt int64                  `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *OrderDeliveredContent `json:"content"`
}

func (e *OrderDeliveredEvent) GetEventType() string         { return e.EventType }
func (e *OrderDeliveredEvent) GetEventId() string           { return e.Id }
func (e *OrderDeliveredEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *OrderDeliveredEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionItemCancellationScheduledEvent represents a omnichannel_subscription_item_cancellation_scheduled webhook event
type OmnichannelSubscriptionItemCancellationScheduledEvent struct {
	Id         string                                                   `json:"id"`
	OccurredAt int64                                                    `json:"occurred_at"`
	EventType  string                                                   `json:"event_type"`
	Content    *OmnichannelSubscriptionItemCancellationScheduledContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemCancellationScheduledEvent) GetEventType() string {
	return e.EventType
}
func (e *OmnichannelSubscriptionItemCancellationScheduledEvent) GetEventId() string { return e.Id }
func (e *OmnichannelSubscriptionItemCancellationScheduledEvent) GetOccurredAt() int64 {
	return e.OccurredAt
}
func (e *OmnichannelSubscriptionItemCancellationScheduledEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// OmnichannelSubscriptionItemGracePeriodExpiredEvent represents a omnichannel_subscription_item_grace_period_expired webhook event
type OmnichannelSubscriptionItemGracePeriodExpiredEvent struct {
	Id         string                                                `json:"id"`
	OccurredAt int64                                                 `json:"occurred_at"`
	EventType  string                                                `json:"event_type"`
	Content    *OmnichannelSubscriptionItemGracePeriodExpiredContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemGracePeriodExpiredEvent) GetEventType() string {
	return e.EventType
}
func (e *OmnichannelSubscriptionItemGracePeriodExpiredEvent) GetEventId() string { return e.Id }
func (e *OmnichannelSubscriptionItemGracePeriodExpiredEvent) GetOccurredAt() int64 {
	return e.OccurredAt
}
func (e *OmnichannelSubscriptionItemGracePeriodExpiredEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// CouponCodesAddedEvent represents a coupon_codes_added webhook event
type CouponCodesAddedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *CouponCodesAddedContent `json:"content"`
}

func (e *CouponCodesAddedEvent) GetEventType() string         { return e.EventType }
func (e *CouponCodesAddedEvent) GetEventId() string           { return e.Id }
func (e *CouponCodesAddedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CouponCodesAddedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// GiftCancelledEvent represents a gift_cancelled webhook event
type GiftCancelledEvent struct {
	Id         string                `json:"id"`
	OccurredAt int64                 `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *GiftCancelledContent `json:"content"`
}

func (e *GiftCancelledEvent) GetEventType() string         { return e.EventType }
func (e *GiftCancelledEvent) GetEventId() string           { return e.Id }
func (e *GiftCancelledEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *GiftCancelledEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OrderCancelledEvent represents a order_cancelled webhook event
type OrderCancelledEvent struct {
	Id         string                 `json:"id"`
	OccurredAt int64                  `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *OrderCancelledContent `json:"content"`
}

func (e *OrderCancelledEvent) GetEventType() string         { return e.EventType }
func (e *OrderCancelledEvent) GetEventId() string           { return e.Id }
func (e *OrderCancelledEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *OrderCancelledEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CouponDeletedEvent represents a coupon_deleted webhook event
type CouponDeletedEvent struct {
	Id         string                `json:"id"`
	OccurredAt int64                 `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *CouponDeletedContent `json:"content"`
}

func (e *CouponDeletedEvent) GetEventType() string         { return e.EventType }
func (e *CouponDeletedEvent) GetEventId() string           { return e.Id }
func (e *CouponDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CouponDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionScheduledChangesRemovedEvent represents a subscription_scheduled_changes_removed webhook event
type SubscriptionScheduledChangesRemovedEvent struct {
	Id         string                                      `json:"id"`
	OccurredAt int64                                       `json:"occurred_at"`
	EventType  string                                      `json:"event_type"`
	Content    *SubscriptionScheduledChangesRemovedContent `json:"content"`
}

func (e *SubscriptionScheduledChangesRemovedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionScheduledChangesRemovedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionScheduledChangesRemovedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionScheduledChangesRemovedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// PendingInvoiceCreatedEvent represents a pending_invoice_created webhook event
type PendingInvoiceCreatedEvent struct {
	Id         string                        `json:"id"`
	OccurredAt int64                         `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *PendingInvoiceCreatedContent `json:"content"`
}

func (e *PendingInvoiceCreatedEvent) GetEventType() string         { return e.EventType }
func (e *PendingInvoiceCreatedEvent) GetEventId() string           { return e.Id }
func (e *PendingInvoiceCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *PendingInvoiceCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// EntitlementOverridesAutoRemovedEvent represents a entitlement_overrides_auto_removed webhook event
type EntitlementOverridesAutoRemovedEvent struct {
	Id         string                                  `json:"id"`
	OccurredAt int64                                   `json:"occurred_at"`
	EventType  string                                  `json:"event_type"`
	Content    *EntitlementOverridesAutoRemovedContent `json:"content"`
}

func (e *EntitlementOverridesAutoRemovedEvent) GetEventType() string { return e.EventType }
func (e *EntitlementOverridesAutoRemovedEvent) GetEventId() string   { return e.Id }
func (e *EntitlementOverridesAutoRemovedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *EntitlementOverridesAutoRemovedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// OmnichannelSubscriptionItemUpgradedEvent represents a omnichannel_subscription_item_upgraded webhook event
type OmnichannelSubscriptionItemUpgradedEvent struct {
	Id         string                                      `json:"id"`
	OccurredAt int64                                       `json:"occurred_at"`
	EventType  string                                      `json:"event_type"`
	Content    *OmnichannelSubscriptionItemUpgradedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemUpgradedEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionItemUpgradedEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionItemUpgradedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelSubscriptionItemUpgradedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// SubscriptionBusinessEntityChangedEvent represents a subscription_business_entity_changed webhook event
type SubscriptionBusinessEntityChangedEvent struct {
	Id         string                                    `json:"id"`
	OccurredAt int64                                     `json:"occurred_at"`
	EventType  string                                    `json:"event_type"`
	Content    *SubscriptionBusinessEntityChangedContent `json:"content"`
}

func (e *SubscriptionBusinessEntityChangedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionBusinessEntityChangedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionBusinessEntityChangedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionBusinessEntityChangedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// OmnichannelOneTimeOrderCreatedEvent represents a omnichannel_one_time_order_created webhook event
type OmnichannelOneTimeOrderCreatedEvent struct {
	Id         string                                 `json:"id"`
	OccurredAt int64                                  `json:"occurred_at"`
	EventType  string                                 `json:"event_type"`
	Content    *OmnichannelOneTimeOrderCreatedContent `json:"content"`
}

func (e *OmnichannelOneTimeOrderCreatedEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelOneTimeOrderCreatedEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelOneTimeOrderCreatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelOneTimeOrderCreatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// PaymentSourceDeletedEvent represents a payment_source_deleted webhook event
type PaymentSourceDeletedEvent struct {
	Id         string                       `json:"id"`
	OccurredAt int64                        `json:"occurred_at"`
	EventType  string                       `json:"event_type"`
	Content    *PaymentSourceDeletedContent `json:"content"`
}

func (e *PaymentSourceDeletedEvent) GetEventType() string         { return e.EventType }
func (e *PaymentSourceDeletedEvent) GetEventId() string           { return e.Id }
func (e *PaymentSourceDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *PaymentSourceDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionItemCancelledEvent represents a omnichannel_subscription_item_cancelled webhook event
type OmnichannelSubscriptionItemCancelledEvent struct {
	Id         string                                       `json:"id"`
	OccurredAt int64                                        `json:"occurred_at"`
	EventType  string                                       `json:"event_type"`
	Content    *OmnichannelSubscriptionItemCancelledContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemCancelledEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionItemCancelledEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionItemCancelledEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelSubscriptionItemCancelledEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// QuoteDeletedEvent represents a quote_deleted webhook event
type QuoteDeletedEvent struct {
	Id         string               `json:"id"`
	OccurredAt int64                `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *QuoteDeletedContent `json:"content"`
}

func (e *QuoteDeletedEvent) GetEventType() string         { return e.EventType }
func (e *QuoteDeletedEvent) GetEventId() string           { return e.Id }
func (e *QuoteDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *QuoteDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// InvoiceUpdatedEvent represents a invoice_updated webhook event
type InvoiceUpdatedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt int64                  `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *InvoiceUpdatedContent `json:"content"`
}

func (e *InvoiceUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *InvoiceUpdatedEvent) GetEventId() string           { return e.Id }
func (e *InvoiceUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *InvoiceUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionAdvanceInvoiceScheduleRemovedEvent represents a subscription_advance_invoice_schedule_removed webhook event
type SubscriptionAdvanceInvoiceScheduleRemovedEvent struct {
	Id         string                                            `json:"id"`
	OccurredAt int64                                             `json:"occurred_at"`
	EventType  string                                            `json:"event_type"`
	Content    *SubscriptionAdvanceInvoiceScheduleRemovedContent `json:"content"`
}

func (e *SubscriptionAdvanceInvoiceScheduleRemovedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionAdvanceInvoiceScheduleRemovedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionAdvanceInvoiceScheduleRemovedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionAdvanceInvoiceScheduleRemovedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// CardDeletedEvent represents a card_deleted webhook event
type CardDeletedEvent struct {
	Id         string              `json:"id"`
	OccurredAt int64               `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *CardDeletedContent `json:"content"`
}

func (e *CardDeletedEvent) GetEventType() string         { return e.EventType }
func (e *CardDeletedEvent) GetEventId() string           { return e.Id }
func (e *CardDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CardDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OrderReadyToShipEvent represents a order_ready_to_ship webhook event
type OrderReadyToShipEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *OrderReadyToShipContent `json:"content"`
}

func (e *OrderReadyToShipEvent) GetEventType() string         { return e.EventType }
func (e *OrderReadyToShipEvent) GetEventId() string           { return e.Id }
func (e *OrderReadyToShipEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *OrderReadyToShipEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionMovedOutEvent represents a subscription_moved_out webhook event
type SubscriptionMovedOutEvent struct {
	Id         string                       `json:"id"`
	OccurredAt int64                        `json:"occurred_at"`
	EventType  string                       `json:"event_type"`
	Content    *SubscriptionMovedOutContent `json:"content"`
}

func (e *SubscriptionMovedOutEvent) GetEventType() string         { return e.EventType }
func (e *SubscriptionMovedOutEvent) GetEventId() string           { return e.Id }
func (e *SubscriptionMovedOutEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *SubscriptionMovedOutEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentScheduleSchemeCreatedEvent represents a payment_schedule_scheme_created webhook event
type PaymentScheduleSchemeCreatedEvent struct {
	Id         string                               `json:"id"`
	OccurredAt int64                                `json:"occurred_at"`
	EventType  string                               `json:"event_type"`
	Content    *PaymentScheduleSchemeCreatedContent `json:"content"`
}

func (e *PaymentScheduleSchemeCreatedEvent) GetEventType() string { return e.EventType }
func (e *PaymentScheduleSchemeCreatedEvent) GetEventId() string   { return e.Id }
func (e *PaymentScheduleSchemeCreatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *PaymentScheduleSchemeCreatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// BusinessEntityUpdatedEvent represents a business_entity_updated webhook event
type BusinessEntityUpdatedEvent struct {
	Id         string                        `json:"id"`
	OccurredAt int64                         `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *BusinessEntityUpdatedContent `json:"content"`
}

func (e *BusinessEntityUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *BusinessEntityUpdatedEvent) GetEventId() string           { return e.Id }
func (e *BusinessEntityUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *BusinessEntityUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionScheduledResumptionRemovedEvent represents a subscription_scheduled_resumption_removed webhook event
type SubscriptionScheduledResumptionRemovedEvent struct {
	Id         string                                         `json:"id"`
	OccurredAt int64                                          `json:"occurred_at"`
	EventType  string                                         `json:"event_type"`
	Content    *SubscriptionScheduledResumptionRemovedContent `json:"content"`
}

func (e *SubscriptionScheduledResumptionRemovedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionScheduledResumptionRemovedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionScheduledResumptionRemovedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionScheduledResumptionRemovedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// PaymentInitiatedEvent represents a payment_initiated webhook event
type PaymentInitiatedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *PaymentInitiatedContent `json:"content"`
}

func (e *PaymentInitiatedEvent) GetEventType() string         { return e.EventType }
func (e *PaymentInitiatedEvent) GetEventId() string           { return e.Id }
func (e *PaymentInitiatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *PaymentInitiatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// FeatureArchivedEvent represents a feature_archived webhook event
type FeatureArchivedEvent struct {
	Id         string                  `json:"id"`
	OccurredAt int64                   `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *FeatureArchivedContent `json:"content"`
}

func (e *FeatureArchivedEvent) GetEventType() string         { return e.EventType }
func (e *FeatureArchivedEvent) GetEventId() string           { return e.Id }
func (e *FeatureArchivedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *FeatureArchivedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionReactivatedWithBackdatingEvent represents a subscription_reactivated_with_backdating webhook event
type SubscriptionReactivatedWithBackdatingEvent struct {
	Id         string                                        `json:"id"`
	OccurredAt int64                                         `json:"occurred_at"`
	EventType  string                                        `json:"event_type"`
	Content    *SubscriptionReactivatedWithBackdatingContent `json:"content"`
}

func (e *SubscriptionReactivatedWithBackdatingEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionReactivatedWithBackdatingEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionReactivatedWithBackdatingEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionReactivatedWithBackdatingEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// OmnichannelSubscriptionImportedEvent represents a omnichannel_subscription_imported webhook event
type OmnichannelSubscriptionImportedEvent struct {
	Id         string                                  `json:"id"`
	OccurredAt int64                                   `json:"occurred_at"`
	EventType  string                                  `json:"event_type"`
	Content    *OmnichannelSubscriptionImportedContent `json:"content"`
}

func (e *OmnichannelSubscriptionImportedEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionImportedEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionImportedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelSubscriptionImportedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// TokenExpiredEvent represents a token_expired webhook event
type TokenExpiredEvent struct {
	Id         string               `json:"id"`
	OccurredAt int64                `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *TokenExpiredContent `json:"content"`
}

func (e *TokenExpiredEvent) GetEventType() string         { return e.EventType }
func (e *TokenExpiredEvent) GetEventId() string           { return e.Id }
func (e *TokenExpiredEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *TokenExpiredEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CardAddedEvent represents a card_added webhook event
type CardAddedEvent struct {
	Id         string            `json:"id"`
	OccurredAt int64             `json:"occurred_at"`
	EventType  string            `json:"event_type"`
	Content    *CardAddedContent `json:"content"`
}

func (e *CardAddedEvent) GetEventType() string         { return e.EventType }
func (e *CardAddedEvent) GetEventId() string           { return e.Id }
func (e *CardAddedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CardAddedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CouponCreatedEvent represents a coupon_created webhook event
type CouponCreatedEvent struct {
	Id         string                `json:"id"`
	OccurredAt int64                 `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *CouponCreatedContent `json:"content"`
}

func (e *CouponCreatedEvent) GetEventType() string         { return e.EventType }
func (e *CouponCreatedEvent) GetEventId() string           { return e.Id }
func (e *CouponCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CouponCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// RuleDeletedEvent represents a rule_deleted webhook event
type RuleDeletedEvent struct {
	Id         string              `json:"id"`
	OccurredAt int64               `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *RuleDeletedContent `json:"content"`
}

func (e *RuleDeletedEvent) GetEventType() string         { return e.EventType }
func (e *RuleDeletedEvent) GetEventId() string           { return e.Id }
func (e *RuleDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *RuleDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ItemPriceEntitlementsUpdatedEvent represents a item_price_entitlements_updated webhook event
type ItemPriceEntitlementsUpdatedEvent struct {
	Id         string                               `json:"id"`
	OccurredAt int64                                `json:"occurred_at"`
	EventType  string                               `json:"event_type"`
	Content    *ItemPriceEntitlementsUpdatedContent `json:"content"`
}

func (e *ItemPriceEntitlementsUpdatedEvent) GetEventType() string { return e.EventType }
func (e *ItemPriceEntitlementsUpdatedEvent) GetEventId() string   { return e.Id }
func (e *ItemPriceEntitlementsUpdatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *ItemPriceEntitlementsUpdatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// ItemPriceDeletedEvent represents a item_price_deleted webhook event
type ItemPriceDeletedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *ItemPriceDeletedContent `json:"content"`
}

func (e *ItemPriceDeletedEvent) GetEventType() string         { return e.EventType }
func (e *ItemPriceDeletedEvent) GetEventId() string           { return e.Id }
func (e *ItemPriceDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *ItemPriceDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// VirtualBankAccountDeletedEvent represents a virtual_bank_account_deleted webhook event
type VirtualBankAccountDeletedEvent struct {
	Id         string                            `json:"id"`
	OccurredAt int64                             `json:"occurred_at"`
	EventType  string                            `json:"event_type"`
	Content    *VirtualBankAccountDeletedContent `json:"content"`
}

func (e *VirtualBankAccountDeletedEvent) GetEventType() string { return e.EventType }
func (e *VirtualBankAccountDeletedEvent) GetEventId() string   { return e.Id }
func (e *VirtualBankAccountDeletedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *VirtualBankAccountDeletedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// PaymentScheduleSchemeDeletedEvent represents a payment_schedule_scheme_deleted webhook event
type PaymentScheduleSchemeDeletedEvent struct {
	Id         string                               `json:"id"`
	OccurredAt int64                                `json:"occurred_at"`
	EventType  string                               `json:"event_type"`
	Content    *PaymentScheduleSchemeDeletedContent `json:"content"`
}

func (e *PaymentScheduleSchemeDeletedEvent) GetEventType() string { return e.EventType }
func (e *PaymentScheduleSchemeDeletedEvent) GetEventId() string   { return e.Id }
func (e *PaymentScheduleSchemeDeletedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *PaymentScheduleSchemeDeletedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// SubscriptionCreatedEvent represents a subscription_created webhook event
type SubscriptionCreatedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *SubscriptionCreatedContent `json:"content"`
}

func (e *SubscriptionCreatedEvent) GetEventType() string         { return e.EventType }
func (e *SubscriptionCreatedEvent) GetEventId() string           { return e.Id }
func (e *SubscriptionCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *SubscriptionCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionEntitlementsCreatedEvent represents a subscription_entitlements_created webhook event
type SubscriptionEntitlementsCreatedEvent struct {
	Id         string                                  `json:"id"`
	OccurredAt int64                                   `json:"occurred_at"`
	EventType  string                                  `json:"event_type"`
	Content    *SubscriptionEntitlementsCreatedContent `json:"content"`
}

func (e *SubscriptionEntitlementsCreatedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionEntitlementsCreatedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionEntitlementsCreatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionEntitlementsCreatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// OrderReturnedEvent represents a order_returned webhook event
type OrderReturnedEvent struct {
	Id         string                `json:"id"`
	OccurredAt int64                 `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *OrderReturnedContent `json:"content"`
}

func (e *OrderReturnedEvent) GetEventType() string         { return e.EventType }
func (e *OrderReturnedEvent) GetEventId() string           { return e.Id }
func (e *OrderReturnedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *OrderReturnedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionDeletedEvent represents a subscription_deleted webhook event
type SubscriptionDeletedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *SubscriptionDeletedContent `json:"content"`
}

func (e *SubscriptionDeletedEvent) GetEventType() string         { return e.EventType }
func (e *SubscriptionDeletedEvent) GetEventId() string           { return e.Id }
func (e *SubscriptionDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *SubscriptionDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentSourceAddedEvent represents a payment_source_added webhook event
type PaymentSourceAddedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt int64                      `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *PaymentSourceAddedContent `json:"content"`
}

func (e *PaymentSourceAddedEvent) GetEventType() string         { return e.EventType }
func (e *PaymentSourceAddedEvent) GetEventId() string           { return e.Id }
func (e *PaymentSourceAddedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *PaymentSourceAddedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionMovedInEvent represents a subscription_moved_in webhook event
type SubscriptionMovedInEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *SubscriptionMovedInContent `json:"content"`
}

func (e *SubscriptionMovedInEvent) GetEventType() string         { return e.EventType }
func (e *SubscriptionMovedInEvent) GetEventId() string           { return e.Id }
func (e *SubscriptionMovedInEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *SubscriptionMovedInEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ItemPriceCreatedEvent represents a item_price_created webhook event
type ItemPriceCreatedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *ItemPriceCreatedContent `json:"content"`
}

func (e *ItemPriceCreatedEvent) GetEventType() string         { return e.EventType }
func (e *ItemPriceCreatedEvent) GetEventId() string           { return e.Id }
func (e *ItemPriceCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *ItemPriceCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionScheduledCancellationRemovedEvent represents a subscription_scheduled_cancellation_removed webhook event
type SubscriptionScheduledCancellationRemovedEvent struct {
	Id         string                                           `json:"id"`
	OccurredAt int64                                            `json:"occurred_at"`
	EventType  string                                           `json:"event_type"`
	Content    *SubscriptionScheduledCancellationRemovedContent `json:"content"`
}

func (e *SubscriptionScheduledCancellationRemovedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionScheduledCancellationRemovedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionScheduledCancellationRemovedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionScheduledCancellationRemovedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// PaymentRefundedEvent represents a payment_refunded webhook event
type PaymentRefundedEvent struct {
	Id         string                  `json:"id"`
	OccurredAt int64                   `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *PaymentRefundedContent `json:"content"`
}

func (e *PaymentRefundedEvent) GetEventType() string         { return e.EventType }
func (e *PaymentRefundedEvent) GetEventId() string           { return e.Id }
func (e *PaymentRefundedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *PaymentRefundedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// UsageFileIngestedEvent represents a usage_file_ingested webhook event
type UsageFileIngestedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt int64                     `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *UsageFileIngestedContent `json:"content"`
}

func (e *UsageFileIngestedEvent) GetEventType() string         { return e.EventType }
func (e *UsageFileIngestedEvent) GetEventId() string           { return e.Id }
func (e *UsageFileIngestedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *UsageFileIngestedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionMovedInEvent represents a omnichannel_subscription_moved_in webhook event
type OmnichannelSubscriptionMovedInEvent struct {
	Id         string                                 `json:"id"`
	OccurredAt int64                                  `json:"occurred_at"`
	EventType  string                                 `json:"event_type"`
	Content    *OmnichannelSubscriptionMovedInContent `json:"content"`
}

func (e *OmnichannelSubscriptionMovedInEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionMovedInEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionMovedInEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelSubscriptionMovedInEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// DifferentialPriceCreatedEvent represents a differential_price_created webhook event
type DifferentialPriceCreatedEvent struct {
	Id         string                           `json:"id"`
	OccurredAt int64                            `json:"occurred_at"`
	EventType  string                           `json:"event_type"`
	Content    *DifferentialPriceCreatedContent `json:"content"`
}

func (e *DifferentialPriceCreatedEvent) GetEventType() string { return e.EventType }
func (e *DifferentialPriceCreatedEvent) GetEventId() string   { return e.Id }
func (e *DifferentialPriceCreatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *DifferentialPriceCreatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// TransactionCreatedEvent represents a transaction_created webhook event
type TransactionCreatedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt int64                      `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *TransactionCreatedContent `json:"content"`
}

func (e *TransactionCreatedEvent) GetEventType() string         { return e.EventType }
func (e *TransactionCreatedEvent) GetEventId() string           { return e.Id }
func (e *TransactionCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *TransactionCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentSucceededEvent represents a payment_succeeded webhook event
type PaymentSucceededEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *PaymentSucceededContent `json:"content"`
}

func (e *PaymentSucceededEvent) GetEventType() string         { return e.EventType }
func (e *PaymentSucceededEvent) GetEventId() string           { return e.Id }
func (e *PaymentSucceededEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *PaymentSucceededEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionCanceledWithBackdatingEvent represents a subscription_canceled_with_backdating webhook event
type SubscriptionCanceledWithBackdatingEvent struct {
	Id         string                                     `json:"id"`
	OccurredAt int64                                      `json:"occurred_at"`
	EventType  string                                     `json:"event_type"`
	Content    *SubscriptionCanceledWithBackdatingContent `json:"content"`
}

func (e *SubscriptionCanceledWithBackdatingEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionCanceledWithBackdatingEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionCanceledWithBackdatingEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionCanceledWithBackdatingEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// UnbilledChargesVoidedEvent represents a unbilled_charges_voided webhook event
type UnbilledChargesVoidedEvent struct {
	Id         string                        `json:"id"`
	OccurredAt int64                         `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *UnbilledChargesVoidedContent `json:"content"`
}

func (e *UnbilledChargesVoidedEvent) GetEventType() string         { return e.EventType }
func (e *UnbilledChargesVoidedEvent) GetEventId() string           { return e.Id }
func (e *UnbilledChargesVoidedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *UnbilledChargesVoidedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// QuoteCreatedEvent represents a quote_created webhook event
type QuoteCreatedEvent struct {
	Id         string               `json:"id"`
	OccurredAt int64                `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *QuoteCreatedContent `json:"content"`
}

func (e *QuoteCreatedEvent) GetEventType() string         { return e.EventType }
func (e *QuoteCreatedEvent) GetEventId() string           { return e.Id }
func (e *QuoteCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *QuoteCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CouponSetDeletedEvent represents a coupon_set_deleted webhook event
type CouponSetDeletedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *CouponSetDeletedContent `json:"content"`
}

func (e *CouponSetDeletedEvent) GetEventType() string         { return e.EventType }
func (e *CouponSetDeletedEvent) GetEventId() string           { return e.Id }
func (e *CouponSetDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CouponSetDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// AttachedItemCreatedEvent represents a attached_item_created webhook event
type AttachedItemCreatedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *AttachedItemCreatedContent `json:"content"`
}

func (e *AttachedItemCreatedEvent) GetEventType() string         { return e.EventType }
func (e *AttachedItemCreatedEvent) GetEventId() string           { return e.Id }
func (e *AttachedItemCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *AttachedItemCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SalesOrderCreatedEvent represents a sales_order_created webhook event
type SalesOrderCreatedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt int64                     `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *SalesOrderCreatedContent `json:"content"`
}

func (e *SalesOrderCreatedEvent) GetEventType() string         { return e.EventType }
func (e *SalesOrderCreatedEvent) GetEventId() string           { return e.Id }
func (e *SalesOrderCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *SalesOrderCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CustomerChangedEvent represents a customer_changed webhook event
type CustomerChangedEvent struct {
	Id         string                  `json:"id"`
	OccurredAt int64                   `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *CustomerChangedContent `json:"content"`
}

func (e *CustomerChangedEvent) GetEventType() string         { return e.EventType }
func (e *CustomerChangedEvent) GetEventId() string           { return e.Id }
func (e *CustomerChangedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CustomerChangedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionStartedEvent represents a subscription_started webhook event
type SubscriptionStartedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *SubscriptionStartedContent `json:"content"`
}

func (e *SubscriptionStartedEvent) GetEventType() string         { return e.EventType }
func (e *SubscriptionStartedEvent) GetEventId() string           { return e.Id }
func (e *SubscriptionStartedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *SubscriptionStartedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionActivatedEvent represents a subscription_activated webhook event
type SubscriptionActivatedEvent struct {
	Id         string                        `json:"id"`
	OccurredAt int64                         `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *SubscriptionActivatedContent `json:"content"`
}

func (e *SubscriptionActivatedEvent) GetEventType() string         { return e.EventType }
func (e *SubscriptionActivatedEvent) GetEventId() string           { return e.Id }
func (e *SubscriptionActivatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *SubscriptionActivatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentSourceExpiringEvent represents a payment_source_expiring webhook event
type PaymentSourceExpiringEvent struct {
	Id         string                        `json:"id"`
	OccurredAt int64                         `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *PaymentSourceExpiringContent `json:"content"`
}

func (e *PaymentSourceExpiringEvent) GetEventType() string         { return e.EventType }
func (e *PaymentSourceExpiringEvent) GetEventId() string           { return e.Id }
func (e *PaymentSourceExpiringEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *PaymentSourceExpiringEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionReactivatedEvent represents a subscription_reactivated webhook event
type SubscriptionReactivatedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *SubscriptionReactivatedContent `json:"content"`
}

func (e *SubscriptionReactivatedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionReactivatedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionReactivatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionReactivatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// OrderUpdatedEvent represents a order_updated webhook event
type OrderUpdatedEvent struct {
	Id         string               `json:"id"`
	OccurredAt int64                `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *OrderUpdatedContent `json:"content"`
}

func (e *OrderUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *OrderUpdatedEvent) GetEventId() string           { return e.Id }
func (e *OrderUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *OrderUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionScheduledPauseRemovedEvent represents a subscription_scheduled_pause_removed webhook event
type SubscriptionScheduledPauseRemovedEvent struct {
	Id         string                                    `json:"id"`
	OccurredAt int64                                     `json:"occurred_at"`
	EventType  string                                    `json:"event_type"`
	Content    *SubscriptionScheduledPauseRemovedContent `json:"content"`
}

func (e *SubscriptionScheduledPauseRemovedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionScheduledPauseRemovedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionScheduledPauseRemovedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionScheduledPauseRemovedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// SubscriptionCancellationReminderEvent represents a subscription_cancellation_reminder webhook event
type SubscriptionCancellationReminderEvent struct {
	Id         string                                   `json:"id"`
	OccurredAt int64                                    `json:"occurred_at"`
	EventType  string                                   `json:"event_type"`
	Content    *SubscriptionCancellationReminderContent `json:"content"`
}

func (e *SubscriptionCancellationReminderEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionCancellationReminderEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionCancellationReminderEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionCancellationReminderEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// SubscriptionCreatedWithBackdatingEvent represents a subscription_created_with_backdating webhook event
type SubscriptionCreatedWithBackdatingEvent struct {
	Id         string                                    `json:"id"`
	OccurredAt int64                                     `json:"occurred_at"`
	EventType  string                                    `json:"event_type"`
	Content    *SubscriptionCreatedWithBackdatingContent `json:"content"`
}

func (e *SubscriptionCreatedWithBackdatingEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionCreatedWithBackdatingEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionCreatedWithBackdatingEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionCreatedWithBackdatingEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// SubscriptionRampCreatedEvent represents a subscription_ramp_created webhook event
type SubscriptionRampCreatedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *SubscriptionRampCreatedContent `json:"content"`
}

func (e *SubscriptionRampCreatedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionRampCreatedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionRampCreatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionRampCreatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// OrderDeletedEvent represents a order_deleted webhook event
type OrderDeletedEvent struct {
	Id         string               `json:"id"`
	OccurredAt int64                `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *OrderDeletedContent `json:"content"`
}

func (e *OrderDeletedEvent) GetEventType() string         { return e.EventType }
func (e *OrderDeletedEvent) GetEventId() string           { return e.Id }
func (e *OrderDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *OrderDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionItemPauseScheduledEvent represents a omnichannel_subscription_item_pause_scheduled webhook event
type OmnichannelSubscriptionItemPauseScheduledEvent struct {
	Id         string                                            `json:"id"`
	OccurredAt int64                                             `json:"occurred_at"`
	EventType  string                                            `json:"event_type"`
	Content    *OmnichannelSubscriptionItemPauseScheduledContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemPauseScheduledEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionItemPauseScheduledEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionItemPauseScheduledEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelSubscriptionItemPauseScheduledEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// GiftUpdatedEvent represents a gift_updated webhook event
type GiftUpdatedEvent struct {
	Id         string              `json:"id"`
	OccurredAt int64               `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *GiftUpdatedContent `json:"content"`
}

func (e *GiftUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *GiftUpdatedEvent) GetEventId() string           { return e.Id }
func (e *GiftUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *GiftUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionTrialExtendedEvent represents a subscription_trial_extended webhook event
type SubscriptionTrialExtendedEvent struct {
	Id         string                            `json:"id"`
	OccurredAt int64                             `json:"occurred_at"`
	EventType  string                            `json:"event_type"`
	Content    *SubscriptionTrialExtendedContent `json:"content"`
}

func (e *SubscriptionTrialExtendedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionTrialExtendedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionTrialExtendedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionTrialExtendedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// OmnichannelSubscriptionItemGracePeriodStartedEvent represents a omnichannel_subscription_item_grace_period_started webhook event
type OmnichannelSubscriptionItemGracePeriodStartedEvent struct {
	Id         string                                                `json:"id"`
	OccurredAt int64                                                 `json:"occurred_at"`
	EventType  string                                                `json:"event_type"`
	Content    *OmnichannelSubscriptionItemGracePeriodStartedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemGracePeriodStartedEvent) GetEventType() string {
	return e.EventType
}
func (e *OmnichannelSubscriptionItemGracePeriodStartedEvent) GetEventId() string { return e.Id }
func (e *OmnichannelSubscriptionItemGracePeriodStartedEvent) GetOccurredAt() int64 {
	return e.OccurredAt
}
func (e *OmnichannelSubscriptionItemGracePeriodStartedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// CardExpiryReminderEvent represents a card_expiry_reminder webhook event
type CardExpiryReminderEvent struct {
	Id         string                     `json:"id"`
	OccurredAt int64                      `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *CardExpiryReminderContent `json:"content"`
}

func (e *CardExpiryReminderEvent) GetEventType() string         { return e.EventType }
func (e *CardExpiryReminderEvent) GetEventId() string           { return e.Id }
func (e *CardExpiryReminderEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CardExpiryReminderEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// TokenCreatedEvent represents a token_created webhook event
type TokenCreatedEvent struct {
	Id         string               `json:"id"`
	OccurredAt int64                `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *TokenCreatedContent `json:"content"`
}

func (e *TokenCreatedEvent) GetEventType() string         { return e.EventType }
func (e *TokenCreatedEvent) GetEventId() string           { return e.Id }
func (e *TokenCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *TokenCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PromotionalCreditsAddedEvent represents a promotional_credits_added webhook event
type PromotionalCreditsAddedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *PromotionalCreditsAddedContent `json:"content"`
}

func (e *PromotionalCreditsAddedEvent) GetEventType() string { return e.EventType }
func (e *PromotionalCreditsAddedEvent) GetEventId() string   { return e.Id }
func (e *PromotionalCreditsAddedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *PromotionalCreditsAddedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// SubscriptionRampUpdatedEvent represents a subscription_ramp_updated webhook event
type SubscriptionRampUpdatedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *SubscriptionRampUpdatedContent `json:"content"`
}

func (e *SubscriptionRampUpdatedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionRampUpdatedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionRampUpdatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionRampUpdatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// CustomerEntitlementsUpdatedEvent represents a customer_entitlements_updated webhook event
type CustomerEntitlementsUpdatedEvent struct {
	Id         string                              `json:"id"`
	OccurredAt int64                               `json:"occurred_at"`
	EventType  string                              `json:"event_type"`
	Content    *CustomerEntitlementsUpdatedContent `json:"content"`
}

func (e *CustomerEntitlementsUpdatedEvent) GetEventType() string { return e.EventType }
func (e *CustomerEntitlementsUpdatedEvent) GetEventId() string   { return e.Id }
func (e *CustomerEntitlementsUpdatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *CustomerEntitlementsUpdatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// PaymentSourceExpiredEvent represents a payment_source_expired webhook event
type PaymentSourceExpiredEvent struct {
	Id         string                       `json:"id"`
	OccurredAt int64                        `json:"occurred_at"`
	EventType  string                       `json:"event_type"`
	Content    *PaymentSourceExpiredContent `json:"content"`
}

func (e *PaymentSourceExpiredEvent) GetEventType() string         { return e.EventType }
func (e *PaymentSourceExpiredEvent) GetEventId() string           { return e.Id }
func (e *PaymentSourceExpiredEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *PaymentSourceExpiredEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CustomerMovedOutEvent represents a customer_moved_out webhook event
type CustomerMovedOutEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *CustomerMovedOutContent `json:"content"`
}

func (e *CustomerMovedOutEvent) GetEventType() string         { return e.EventType }
func (e *CustomerMovedOutEvent) GetEventId() string           { return e.Id }
func (e *CustomerMovedOutEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CustomerMovedOutEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionEntitlementsUpdatedEvent represents a subscription_entitlements_updated webhook event
type SubscriptionEntitlementsUpdatedEvent struct {
	Id         string                                  `json:"id"`
	OccurredAt int64                                   `json:"occurred_at"`
	EventType  string                                  `json:"event_type"`
	Content    *SubscriptionEntitlementsUpdatedContent `json:"content"`
}

func (e *SubscriptionEntitlementsUpdatedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionEntitlementsUpdatedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionEntitlementsUpdatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionEntitlementsUpdatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// OmnichannelSubscriptionItemDunningExpiredEvent represents a omnichannel_subscription_item_dunning_expired webhook event
type OmnichannelSubscriptionItemDunningExpiredEvent struct {
	Id         string                                            `json:"id"`
	OccurredAt int64                                             `json:"occurred_at"`
	EventType  string                                            `json:"event_type"`
	Content    *OmnichannelSubscriptionItemDunningExpiredContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemDunningExpiredEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionItemDunningExpiredEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionItemDunningExpiredEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelSubscriptionItemDunningExpiredEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// HierarchyCreatedEvent represents a hierarchy_created webhook event
type HierarchyCreatedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *HierarchyCreatedContent `json:"content"`
}

func (e *HierarchyCreatedEvent) GetEventType() string         { return e.EventType }
func (e *HierarchyCreatedEvent) GetEventId() string           { return e.Id }
func (e *HierarchyCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *HierarchyCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// AttachedItemDeletedEvent represents a attached_item_deleted webhook event
type AttachedItemDeletedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *AttachedItemDeletedContent `json:"content"`
}

func (e *AttachedItemDeletedEvent) GetEventType() string         { return e.EventType }
func (e *AttachedItemDeletedEvent) GetEventId() string           { return e.Id }
func (e *AttachedItemDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *AttachedItemDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionItemScheduledCancellationRemovedEvent represents a omnichannel_subscription_item_scheduled_cancellation_removed webhook event
type OmnichannelSubscriptionItemScheduledCancellationRemovedEvent struct {
	Id         string                                                          `json:"id"`
	OccurredAt int64                                                           `json:"occurred_at"`
	EventType  string                                                          `json:"event_type"`
	Content    *OmnichannelSubscriptionItemScheduledCancellationRemovedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemScheduledCancellationRemovedEvent) GetEventType() string {
	return e.EventType
}
func (e *OmnichannelSubscriptionItemScheduledCancellationRemovedEvent) GetEventId() string {
	return e.Id
}
func (e *OmnichannelSubscriptionItemScheduledCancellationRemovedEvent) GetOccurredAt() int64 {
	return e.OccurredAt
}
func (e *OmnichannelSubscriptionItemScheduledCancellationRemovedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// ItemUpdatedEvent represents a item_updated webhook event
type ItemUpdatedEvent struct {
	Id         string              `json:"id"`
	OccurredAt int64               `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *ItemUpdatedContent `json:"content"`
}

func (e *ItemUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *ItemUpdatedEvent) GetEventId() string           { return e.Id }
func (e *ItemUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *ItemUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CouponSetCreatedEvent represents a coupon_set_created webhook event
type CouponSetCreatedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *CouponSetCreatedContent `json:"content"`
}

func (e *CouponSetCreatedEvent) GetEventType() string         { return e.EventType }
func (e *CouponSetCreatedEvent) GetEventId() string           { return e.Id }
func (e *CouponSetCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *CouponSetCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentIntentUpdatedEvent represents a payment_intent_updated webhook event
type PaymentIntentUpdatedEvent struct {
	Id         string                       `json:"id"`
	OccurredAt int64                        `json:"occurred_at"`
	EventType  string                       `json:"event_type"`
	Content    *PaymentIntentUpdatedContent `json:"content"`
}

func (e *PaymentIntentUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *PaymentIntentUpdatedEvent) GetEventId() string           { return e.Id }
func (e *PaymentIntentUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *PaymentIntentUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OrderResentEvent represents a order_resent webhook event
type OrderResentEvent struct {
	Id         string              `json:"id"`
	OccurredAt int64               `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *OrderResentContent `json:"content"`
}

func (e *OrderResentEvent) GetEventType() string         { return e.EventType }
func (e *OrderResentEvent) GetEventId() string           { return e.Id }
func (e *OrderResentEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *OrderResentEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionCreatedEvent represents a omnichannel_subscription_created webhook event
type OmnichannelSubscriptionCreatedEvent struct {
	Id         string                                 `json:"id"`
	OccurredAt int64                                  `json:"occurred_at"`
	EventType  string                                 `json:"event_type"`
	Content    *OmnichannelSubscriptionCreatedContent `json:"content"`
}

func (e *OmnichannelSubscriptionCreatedEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionCreatedEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionCreatedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelSubscriptionCreatedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// TaxWithheldRecordedEvent represents a tax_withheld_recorded webhook event
type TaxWithheldRecordedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *TaxWithheldRecordedContent `json:"content"`
}

func (e *TaxWithheldRecordedEvent) GetEventType() string         { return e.EventType }
func (e *TaxWithheldRecordedEvent) GetEventId() string           { return e.Id }
func (e *TaxWithheldRecordedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *TaxWithheldRecordedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PriceVariantCreatedEvent represents a price_variant_created webhook event
type PriceVariantCreatedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *PriceVariantCreatedContent `json:"content"`
}

func (e *PriceVariantCreatedEvent) GetEventType() string         { return e.EventType }
func (e *PriceVariantCreatedEvent) GetEventId() string           { return e.Id }
func (e *PriceVariantCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *PriceVariantCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// DifferentialPriceDeletedEvent represents a differential_price_deleted webhook event
type DifferentialPriceDeletedEvent struct {
	Id         string                           `json:"id"`
	OccurredAt int64                            `json:"occurred_at"`
	EventType  string                           `json:"event_type"`
	Content    *DifferentialPriceDeletedContent `json:"content"`
}

func (e *DifferentialPriceDeletedEvent) GetEventType() string { return e.EventType }
func (e *DifferentialPriceDeletedEvent) GetEventId() string   { return e.Id }
func (e *DifferentialPriceDeletedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *DifferentialPriceDeletedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// SubscriptionItemsRenewedEvent represents a subscription_items_renewed webhook event
type SubscriptionItemsRenewedEvent struct {
	Id         string                           `json:"id"`
	OccurredAt int64                            `json:"occurred_at"`
	EventType  string                           `json:"event_type"`
	Content    *SubscriptionItemsRenewedContent `json:"content"`
}

func (e *SubscriptionItemsRenewedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionItemsRenewedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionItemsRenewedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionItemsRenewedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// RuleCreatedEvent represents a rule_created webhook event
type RuleCreatedEvent struct {
	Id         string              `json:"id"`
	OccurredAt int64               `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *RuleCreatedContent `json:"content"`
}

func (e *RuleCreatedEvent) GetEventType() string         { return e.EventType }
func (e *RuleCreatedEvent) GetEventId() string           { return e.Id }
func (e *RuleCreatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *RuleCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ContractTermCancelledEvent represents a contract_term_cancelled webhook event
type ContractTermCancelledEvent struct {
	Id         string                        `json:"id"`
	OccurredAt int64                         `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *ContractTermCancelledContent `json:"content"`
}

func (e *ContractTermCancelledEvent) GetEventType() string         { return e.EventType }
func (e *ContractTermCancelledEvent) GetEventId() string           { return e.Id }
func (e *ContractTermCancelledEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *ContractTermCancelledEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ContractTermRenewedEvent represents a contract_term_renewed webhook event
type ContractTermRenewedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *ContractTermRenewedContent `json:"content"`
}

func (e *ContractTermRenewedEvent) GetEventType() string         { return e.EventType }
func (e *ContractTermRenewedEvent) GetEventId() string           { return e.Id }
func (e *ContractTermRenewedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *ContractTermRenewedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// InvoiceDeletedEvent represents a invoice_deleted webhook event
type InvoiceDeletedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt int64                  `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *InvoiceDeletedContent `json:"content"`
}

func (e *InvoiceDeletedEvent) GetEventType() string         { return e.EventType }
func (e *InvoiceDeletedEvent) GetEventId() string           { return e.Id }
func (e *InvoiceDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *InvoiceDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ItemPriceEntitlementsRemovedEvent represents a item_price_entitlements_removed webhook event
type ItemPriceEntitlementsRemovedEvent struct {
	Id         string                               `json:"id"`
	OccurredAt int64                                `json:"occurred_at"`
	EventType  string                               `json:"event_type"`
	Content    *ItemPriceEntitlementsRemovedContent `json:"content"`
}

func (e *ItemPriceEntitlementsRemovedEvent) GetEventType() string { return e.EventType }
func (e *ItemPriceEntitlementsRemovedEvent) GetEventId() string   { return e.Id }
func (e *ItemPriceEntitlementsRemovedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *ItemPriceEntitlementsRemovedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// SalesOrderUpdatedEvent represents a sales_order_updated webhook event
type SalesOrderUpdatedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt int64                     `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *SalesOrderUpdatedContent `json:"content"`
}

func (e *SalesOrderUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *SalesOrderUpdatedEvent) GetEventId() string           { return e.Id }
func (e *SalesOrderUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *SalesOrderUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionItemDunningStartedEvent represents a omnichannel_subscription_item_dunning_started webhook event
type OmnichannelSubscriptionItemDunningStartedEvent struct {
	Id         string                                            `json:"id"`
	OccurredAt int64                                             `json:"occurred_at"`
	EventType  string                                            `json:"event_type"`
	Content    *OmnichannelSubscriptionItemDunningStartedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemDunningStartedEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionItemDunningStartedEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionItemDunningStartedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelSubscriptionItemDunningStartedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// OmnichannelSubscriptionItemChangeScheduledEvent represents a omnichannel_subscription_item_change_scheduled webhook event
type OmnichannelSubscriptionItemChangeScheduledEvent struct {
	Id         string                                             `json:"id"`
	OccurredAt int64                                              `json:"occurred_at"`
	EventType  string                                             `json:"event_type"`
	Content    *OmnichannelSubscriptionItemChangeScheduledContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemChangeScheduledEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionItemChangeScheduledEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionItemChangeScheduledEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *OmnichannelSubscriptionItemChangeScheduledEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// PendingInvoiceUpdatedEvent represents a pending_invoice_updated webhook event
type PendingInvoiceUpdatedEvent struct {
	Id         string                        `json:"id"`
	OccurredAt int64                         `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *PendingInvoiceUpdatedContent `json:"content"`
}

func (e *PendingInvoiceUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *PendingInvoiceUpdatedEvent) GetEventId() string           { return e.Id }
func (e *PendingInvoiceUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *PendingInvoiceUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// QuoteUpdatedEvent represents a quote_updated webhook event
type QuoteUpdatedEvent struct {
	Id         string               `json:"id"`
	OccurredAt int64                `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *QuoteUpdatedContent `json:"content"`
}

func (e *QuoteUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *QuoteUpdatedEvent) GetEventId() string           { return e.Id }
func (e *QuoteUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *QuoteUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// AttachedItemUpdatedEvent represents a attached_item_updated webhook event
type AttachedItemUpdatedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *AttachedItemUpdatedContent `json:"content"`
}

func (e *AttachedItemUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *AttachedItemUpdatedEvent) GetEventId() string           { return e.Id }
func (e *AttachedItemUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *AttachedItemUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentSourceUpdatedEvent represents a payment_source_updated webhook event
type PaymentSourceUpdatedEvent struct {
	Id         string                       `json:"id"`
	OccurredAt int64                        `json:"occurred_at"`
	EventType  string                       `json:"event_type"`
	Content    *PaymentSourceUpdatedContent `json:"content"`
}

func (e *PaymentSourceUpdatedEvent) GetEventType() string         { return e.EventType }
func (e *PaymentSourceUpdatedEvent) GetEventId() string           { return e.Id }
func (e *PaymentSourceUpdatedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *PaymentSourceUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// BusinessEntityDeletedEvent represents a business_entity_deleted webhook event
type BusinessEntityDeletedEvent struct {
	Id         string                        `json:"id"`
	OccurredAt int64                         `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *BusinessEntityDeletedContent `json:"content"`
}

func (e *BusinessEntityDeletedEvent) GetEventType() string         { return e.EventType }
func (e *BusinessEntityDeletedEvent) GetEventId() string           { return e.Id }
func (e *BusinessEntityDeletedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *BusinessEntityDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// AuthorizationVoidedEvent represents a authorization_voided webhook event
type AuthorizationVoidedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *AuthorizationVoidedContent `json:"content"`
}

func (e *AuthorizationVoidedEvent) GetEventType() string         { return e.EventType }
func (e *AuthorizationVoidedEvent) GetEventId() string           { return e.Id }
func (e *AuthorizationVoidedEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *AuthorizationVoidedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionRampDeletedEvent represents a subscription_ramp_deleted webhook event
type SubscriptionRampDeletedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *SubscriptionRampDeletedContent `json:"content"`
}

func (e *SubscriptionRampDeletedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionRampDeletedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionRampDeletedEvent) GetOccurredAt() int64 { return e.OccurredAt }
func (e *SubscriptionRampDeletedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// GenericWebhookEvent represents an unsupported webhook event
type GenericWebhookEvent struct {
	Id         string          `json:"id"`
	OccurredAt int64           `json:"occurred_at"`
	EventType  string          `json:"event_type"`
	RawContent json.RawMessage `json:"content"`
}

func (e *GenericWebhookEvent) GetEventType() string         { return e.EventType }
func (e *GenericWebhookEvent) GetEventId() string           { return e.Id }
func (e *GenericWebhookEvent) GetOccurredAt() int64         { return e.OccurredAt }
func (e *GenericWebhookEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ParseWebhook deserializes webhook body and returns the appropriate typed event
func ParseWebhook(body string) (WebhookEventInterface, error) {
	evt := Deserialize(body)

	// Parse content based on event type and return the specific typed event
	switch WebhookContentType(evt.EventType) {

	case WebhookContentTypeSubscriptionPauseScheduled:
		var content SubscriptionPauseScheduledContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionPauseScheduledEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCustomerBusinessEntityChanged:
		var content CustomerBusinessEntityChangedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CustomerBusinessEntityChangedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionAdvanceInvoiceScheduleAdded:
		var content SubscriptionAdvanceInvoiceScheduleAddedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionAdvanceInvoiceScheduleAddedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeGiftExpired:
		var content GiftExpiredContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &GiftExpiredEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeTaxWithheldDeleted:
		var content TaxWithheldDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &TaxWithheldDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeUnbilledChargesDeleted:
		var content UnbilledChargesDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &UnbilledChargesDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCouponUpdated:
		var content CouponUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CouponUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemReactivated:
		var content OmnichannelSubscriptionItemReactivatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemReactivatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemRenewed:
		var content OmnichannelSubscriptionItemRenewedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemRenewedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeUnbilledChargesCreated:
		var content UnbilledChargesCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &UnbilledChargesCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionResumed:
		var content SubscriptionResumedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionResumedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelOneTimeOrderItemCancelled:
		var content OmnichannelOneTimeOrderItemCancelledContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelOneTimeOrderItemCancelledEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionCancelled:
		var content SubscriptionCancelledContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionCancelledEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemEntitlementsRemoved:
		var content ItemEntitlementsRemovedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &ItemEntitlementsRemovedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeBusinessEntityCreated:
		var content BusinessEntityCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &BusinessEntityCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCouponSetUpdated:
		var content CouponSetUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CouponSetUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeDifferentialPriceUpdated:
		var content DifferentialPriceUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &DifferentialPriceUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemPaused:
		var content OmnichannelSubscriptionItemPausedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemPausedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeEntitlementOverridesRemoved:
		var content EntitlementOverridesRemovedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &EntitlementOverridesRemovedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionActivatedWithBackdating:
		var content SubscriptionActivatedWithBackdatingContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionActivatedWithBackdatingEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionTrialEndReminder:
		var content SubscriptionTrialEndReminderContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionTrialEndReminderEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionShippingAddressUpdated:
		var content SubscriptionShippingAddressUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionShippingAddressUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeVoucherCreateFailed:
		var content VoucherCreateFailedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &VoucherCreateFailedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeGiftClaimed:
		var content GiftClaimedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &GiftClaimedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCustomerDeleted:
		var content CustomerDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CustomerDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeRefundInitiated:
		var content RefundInitiatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &RefundInitiatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeInvoiceGeneratedWithBackdating:
		var content InvoiceGeneratedWithBackdatingContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &InvoiceGeneratedWithBackdatingEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelTransactionCreated:
		var content OmnichannelTransactionCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelTransactionCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeAddUsagesReminder:
		var content AddUsagesReminderContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &AddUsagesReminderEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeVoucherCreated:
		var content VoucherCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &VoucherCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeRuleUpdated:
		var content RuleUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &RuleUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentSchedulesCreated:
		var content PaymentSchedulesCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PaymentSchedulesCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeFeatureActivated:
		var content FeatureActivatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &FeatureActivatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentSourceLocallyDeleted:
		var content PaymentSourceLocallyDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PaymentSourceLocallyDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeInvoiceGenerated:
		var content InvoiceGeneratedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &InvoiceGeneratedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeVoucherExpired:
		var content VoucherExpiredContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &VoucherExpiredEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeAuthorizationSucceeded:
		var content AuthorizationSucceededContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &AuthorizationSucceededEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeGiftScheduled:
		var content GiftScheduledContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &GiftScheduledEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionChangesScheduled:
		var content SubscriptionChangesScheduledContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionChangesScheduledEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionChangedWithBackdating:
		var content SubscriptionChangedWithBackdatingContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionChangedWithBackdatingEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemChanged:
		var content OmnichannelSubscriptionItemChangedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemChangedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeGiftUnclaimed:
		var content GiftUnclaimedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &GiftUnclaimedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeVirtualBankAccountAdded:
		var content VirtualBankAccountAddedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &VirtualBankAccountAddedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentIntentCreated:
		var content PaymentIntentCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PaymentIntentCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCreditNoteCreatedWithBackdating:
		var content CreditNoteCreatedWithBackdatingContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CreditNoteCreatedWithBackdatingEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeContractTermTerminated:
		var content ContractTermTerminatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &ContractTermTerminatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemFamilyUpdated:
		var content ItemFamilyUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &ItemFamilyUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOrderCreated:
		var content OrderCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OrderCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePriceVariantDeleted:
		var content PriceVariantDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PriceVariantDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionMovementFailed:
		var content SubscriptionMovementFailedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionMovementFailedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCustomerMovedIn:
		var content CustomerMovedInContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CustomerMovedInEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionAdvanceInvoiceScheduleUpdated:
		var content SubscriptionAdvanceInvoiceScheduleUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionAdvanceInvoiceScheduleUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemDeleted:
		var content ItemDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &ItemDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionRampDrafted:
		var content SubscriptionRampDraftedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionRampDraftedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeDunningUpdated:
		var content DunningUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &DunningUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemEntitlementsUpdated:
		var content ItemEntitlementsUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &ItemEntitlementsUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeTokenConsumed:
		var content TokenConsumedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &TokenConsumedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeHierarchyDeleted:
		var content HierarchyDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &HierarchyDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionCancellationScheduled:
		var content SubscriptionCancellationScheduledContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionCancellationScheduledEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionRenewed:
		var content SubscriptionRenewedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionRenewedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeFeatureUpdated:
		var content FeatureUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &FeatureUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeFeatureDeleted:
		var content FeatureDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &FeatureDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemFamilyCreated:
		var content ItemFamilyCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &ItemFamilyCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemScheduledChangeRemoved:
		var content OmnichannelSubscriptionItemScheduledChangeRemovedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemScheduledChangeRemovedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemResumed:
		var content OmnichannelSubscriptionItemResumedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemResumedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePurchaseCreated:
		var content PurchaseCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PurchaseCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeEntitlementOverridesUpdated:
		var content EntitlementOverridesUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &EntitlementOverridesUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemFamilyDeleted:
		var content ItemFamilyDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &ItemFamilyDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionResumptionScheduled:
		var content SubscriptionResumptionScheduledContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionResumptionScheduledEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeFeatureReactivated:
		var content FeatureReactivatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &FeatureReactivatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCouponCodesDeleted:
		var content CouponCodesDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CouponCodesDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCardExpired:
		var content CardExpiredContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CardExpiredEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCreditNoteUpdated:
		var content CreditNoteUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CreditNoteUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemDowngraded:
		var content OmnichannelSubscriptionItemDowngradedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemDowngradedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePriceVariantUpdated:
		var content PriceVariantUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PriceVariantUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePromotionalCreditsDeducted:
		var content PromotionalCreditsDeductedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PromotionalCreditsDeductedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionRampApplied:
		var content SubscriptionRampAppliedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionRampAppliedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionPaused:
		var content SubscriptionPausedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionPausedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOrderReadyToProcess:
		var content OrderReadyToProcessContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OrderReadyToProcessEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeFeatureCreated:
		var content FeatureCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &FeatureCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeTransactionDeleted:
		var content TransactionDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &TransactionDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCreditNoteCreated:
		var content CreditNoteCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CreditNoteCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemResubscribed:
		var content OmnichannelSubscriptionItemResubscribedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemResubscribedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeRecordPurchaseFailed:
		var content RecordPurchaseFailedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &RecordPurchaseFailedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemCreated:
		var content ItemCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &ItemCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeTransactionUpdated:
		var content TransactionUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &TransactionUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeMrrUpdated:
		var content MrrUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &MrrUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeUnbilledChargesInvoiced:
		var content UnbilledChargesInvoicedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &UnbilledChargesInvoicedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemPriceUpdated:
		var content ItemPriceUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &ItemPriceUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCouponCodesUpdated:
		var content CouponCodesUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CouponCodesUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeVirtualBankAccountUpdated:
		var content VirtualBankAccountUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &VirtualBankAccountUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeContractTermCreated:
		var content ContractTermCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &ContractTermCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionChanged:
		var content SubscriptionChangedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionChangedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentFailed:
		var content PaymentFailedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PaymentFailedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCreditNoteDeleted:
		var content CreditNoteDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CreditNoteDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeTaxWithheldRefunded:
		var content TaxWithheldRefundedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &TaxWithheldRefundedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeContractTermCompleted:
		var content ContractTermCompletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &ContractTermCompletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentSchedulesUpdated:
		var content PaymentSchedulesUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PaymentSchedulesUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemExpired:
		var content OmnichannelSubscriptionItemExpiredContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemExpiredEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCardUpdated:
		var content CardUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CardUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCustomerCreated:
		var content CustomerCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CustomerCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionRenewalReminder:
		var content SubscriptionRenewalReminderContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionRenewalReminderEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOrderDelivered:
		var content OrderDeliveredContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OrderDeliveredEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemCancellationScheduled:
		var content OmnichannelSubscriptionItemCancellationScheduledContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemCancellationScheduledEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemGracePeriodExpired:
		var content OmnichannelSubscriptionItemGracePeriodExpiredContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemGracePeriodExpiredEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCouponCodesAdded:
		var content CouponCodesAddedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CouponCodesAddedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeGiftCancelled:
		var content GiftCancelledContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &GiftCancelledEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOrderCancelled:
		var content OrderCancelledContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OrderCancelledEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCouponDeleted:
		var content CouponDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CouponDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionScheduledChangesRemoved:
		var content SubscriptionScheduledChangesRemovedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionScheduledChangesRemovedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePendingInvoiceCreated:
		var content PendingInvoiceCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PendingInvoiceCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeEntitlementOverridesAutoRemoved:
		var content EntitlementOverridesAutoRemovedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &EntitlementOverridesAutoRemovedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemUpgraded:
		var content OmnichannelSubscriptionItemUpgradedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemUpgradedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionBusinessEntityChanged:
		var content SubscriptionBusinessEntityChangedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionBusinessEntityChangedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelOneTimeOrderCreated:
		var content OmnichannelOneTimeOrderCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelOneTimeOrderCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentSourceDeleted:
		var content PaymentSourceDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PaymentSourceDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemCancelled:
		var content OmnichannelSubscriptionItemCancelledContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemCancelledEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeQuoteDeleted:
		var content QuoteDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &QuoteDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeInvoiceUpdated:
		var content InvoiceUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &InvoiceUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionAdvanceInvoiceScheduleRemoved:
		var content SubscriptionAdvanceInvoiceScheduleRemovedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionAdvanceInvoiceScheduleRemovedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCardDeleted:
		var content CardDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CardDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOrderReadyToShip:
		var content OrderReadyToShipContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OrderReadyToShipEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionMovedOut:
		var content SubscriptionMovedOutContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionMovedOutEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentScheduleSchemeCreated:
		var content PaymentScheduleSchemeCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PaymentScheduleSchemeCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeBusinessEntityUpdated:
		var content BusinessEntityUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &BusinessEntityUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionScheduledResumptionRemoved:
		var content SubscriptionScheduledResumptionRemovedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionScheduledResumptionRemovedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentInitiated:
		var content PaymentInitiatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PaymentInitiatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeFeatureArchived:
		var content FeatureArchivedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &FeatureArchivedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionReactivatedWithBackdating:
		var content SubscriptionReactivatedWithBackdatingContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionReactivatedWithBackdatingEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionImported:
		var content OmnichannelSubscriptionImportedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionImportedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeTokenExpired:
		var content TokenExpiredContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &TokenExpiredEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCardAdded:
		var content CardAddedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CardAddedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCouponCreated:
		var content CouponCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CouponCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeRuleDeleted:
		var content RuleDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &RuleDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemPriceEntitlementsUpdated:
		var content ItemPriceEntitlementsUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &ItemPriceEntitlementsUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemPriceDeleted:
		var content ItemPriceDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &ItemPriceDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeVirtualBankAccountDeleted:
		var content VirtualBankAccountDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &VirtualBankAccountDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentScheduleSchemeDeleted:
		var content PaymentScheduleSchemeDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PaymentScheduleSchemeDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionCreated:
		var content SubscriptionCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionEntitlementsCreated:
		var content SubscriptionEntitlementsCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionEntitlementsCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOrderReturned:
		var content OrderReturnedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OrderReturnedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionDeleted:
		var content SubscriptionDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentSourceAdded:
		var content PaymentSourceAddedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PaymentSourceAddedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionMovedIn:
		var content SubscriptionMovedInContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionMovedInEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemPriceCreated:
		var content ItemPriceCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &ItemPriceCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionScheduledCancellationRemoved:
		var content SubscriptionScheduledCancellationRemovedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionScheduledCancellationRemovedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentRefunded:
		var content PaymentRefundedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PaymentRefundedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeUsageFileIngested:
		var content UsageFileIngestedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &UsageFileIngestedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionMovedIn:
		var content OmnichannelSubscriptionMovedInContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionMovedInEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeDifferentialPriceCreated:
		var content DifferentialPriceCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &DifferentialPriceCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeTransactionCreated:
		var content TransactionCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &TransactionCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentSucceeded:
		var content PaymentSucceededContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PaymentSucceededEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionCanceledWithBackdating:
		var content SubscriptionCanceledWithBackdatingContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionCanceledWithBackdatingEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeUnbilledChargesVoided:
		var content UnbilledChargesVoidedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &UnbilledChargesVoidedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeQuoteCreated:
		var content QuoteCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &QuoteCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCouponSetDeleted:
		var content CouponSetDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CouponSetDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeAttachedItemCreated:
		var content AttachedItemCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &AttachedItemCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSalesOrderCreated:
		var content SalesOrderCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SalesOrderCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCustomerChanged:
		var content CustomerChangedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CustomerChangedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionStarted:
		var content SubscriptionStartedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionStartedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionActivated:
		var content SubscriptionActivatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionActivatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentSourceExpiring:
		var content PaymentSourceExpiringContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PaymentSourceExpiringEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionReactivated:
		var content SubscriptionReactivatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionReactivatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOrderUpdated:
		var content OrderUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OrderUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionScheduledPauseRemoved:
		var content SubscriptionScheduledPauseRemovedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionScheduledPauseRemovedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionCancellationReminder:
		var content SubscriptionCancellationReminderContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionCancellationReminderEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionCreatedWithBackdating:
		var content SubscriptionCreatedWithBackdatingContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionCreatedWithBackdatingEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionRampCreated:
		var content SubscriptionRampCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionRampCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOrderDeleted:
		var content OrderDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OrderDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemPauseScheduled:
		var content OmnichannelSubscriptionItemPauseScheduledContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemPauseScheduledEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeGiftUpdated:
		var content GiftUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &GiftUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionTrialExtended:
		var content SubscriptionTrialExtendedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionTrialExtendedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemGracePeriodStarted:
		var content OmnichannelSubscriptionItemGracePeriodStartedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemGracePeriodStartedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCardExpiryReminder:
		var content CardExpiryReminderContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CardExpiryReminderEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeTokenCreated:
		var content TokenCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &TokenCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePromotionalCreditsAdded:
		var content PromotionalCreditsAddedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PromotionalCreditsAddedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionRampUpdated:
		var content SubscriptionRampUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionRampUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCustomerEntitlementsUpdated:
		var content CustomerEntitlementsUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CustomerEntitlementsUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentSourceExpired:
		var content PaymentSourceExpiredContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PaymentSourceExpiredEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCustomerMovedOut:
		var content CustomerMovedOutContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CustomerMovedOutEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionEntitlementsUpdated:
		var content SubscriptionEntitlementsUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionEntitlementsUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemDunningExpired:
		var content OmnichannelSubscriptionItemDunningExpiredContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemDunningExpiredEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeHierarchyCreated:
		var content HierarchyCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &HierarchyCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeAttachedItemDeleted:
		var content AttachedItemDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &AttachedItemDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemScheduledCancellationRemoved:
		var content OmnichannelSubscriptionItemScheduledCancellationRemovedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemScheduledCancellationRemovedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemUpdated:
		var content ItemUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &ItemUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeCouponSetCreated:
		var content CouponSetCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &CouponSetCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentIntentUpdated:
		var content PaymentIntentUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PaymentIntentUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOrderResent:
		var content OrderResentContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OrderResentEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionCreated:
		var content OmnichannelSubscriptionCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeTaxWithheldRecorded:
		var content TaxWithheldRecordedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &TaxWithheldRecordedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePriceVariantCreated:
		var content PriceVariantCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PriceVariantCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeDifferentialPriceDeleted:
		var content DifferentialPriceDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &DifferentialPriceDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionItemsRenewed:
		var content SubscriptionItemsRenewedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionItemsRenewedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeRuleCreated:
		var content RuleCreatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &RuleCreatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeContractTermCancelled:
		var content ContractTermCancelledContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &ContractTermCancelledEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeContractTermRenewed:
		var content ContractTermRenewedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &ContractTermRenewedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeInvoiceDeleted:
		var content InvoiceDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &InvoiceDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemPriceEntitlementsRemoved:
		var content ItemPriceEntitlementsRemovedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &ItemPriceEntitlementsRemovedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSalesOrderUpdated:
		var content SalesOrderUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SalesOrderUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemDunningStarted:
		var content OmnichannelSubscriptionItemDunningStartedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemDunningStartedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemChangeScheduled:
		var content OmnichannelSubscriptionItemChangeScheduledContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemChangeScheduledEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePendingInvoiceUpdated:
		var content PendingInvoiceUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PendingInvoiceUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeQuoteUpdated:
		var content QuoteUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &QuoteUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeAttachedItemUpdated:
		var content AttachedItemUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &AttachedItemUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentSourceUpdated:
		var content PaymentSourceUpdatedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &PaymentSourceUpdatedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeBusinessEntityDeleted:
		var content BusinessEntityDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &BusinessEntityDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeAuthorizationVoided:
		var content AuthorizationVoidedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &AuthorizationVoidedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionRampDeleted:
		var content SubscriptionRampDeletedContent
		if err := json.Unmarshal(evt.Content, &content); err != nil {
			return nil, err
		}
		return &SubscriptionRampDeletedEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			Content:    &content,
		}, nil

	default:
		// For unsupported event types, return generic event
		return &GenericWebhookEvent{
			Id:         evt.Id,
			OccurredAt: evt.OccurredAt,
			EventType:  string(evt.EventType),
			RawContent: evt.Content,
		}, nil
	}
}
