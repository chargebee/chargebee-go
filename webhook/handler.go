package webhook

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/chargebee/chargebee-go/v4"
)

type WebhookHandler struct {
	RequestValidator func(*http.Request) error
	OnError          func(http.ResponseWriter, *http.Request, error)
	OnUnhandledEvent func(chargebee.EventType, []byte) error

	OnAddUsagesReminder func(AddUsagesReminderEvent) error

	OnAddonCreated func(AddonCreatedEvent) error

	OnAddonDeleted func(AddonDeletedEvent) error

	OnAddonUpdated func(AddonUpdatedEvent) error

	OnAttachedItemCreated func(AttachedItemCreatedEvent) error

	OnAttachedItemDeleted func(AttachedItemDeletedEvent) error

	OnAttachedItemUpdated func(AttachedItemUpdatedEvent) error

	OnAuthorizationSucceeded func(AuthorizationSucceededEvent) error

	OnAuthorizationVoided func(AuthorizationVoidedEvent) error

	OnBusinessEntityCreated func(BusinessEntityCreatedEvent) error

	OnBusinessEntityDeleted func(BusinessEntityDeletedEvent) error

	OnBusinessEntityUpdated func(BusinessEntityUpdatedEvent) error

	OnCardAdded func(CardAddedEvent) error

	OnCardDeleted func(CardDeletedEvent) error

	OnCardExpired func(CardExpiredEvent) error

	OnCardExpiryReminder func(CardExpiryReminderEvent) error

	OnCardUpdated func(CardUpdatedEvent) error

	OnContractTermCancelled func(ContractTermCancelledEvent) error

	OnContractTermCompleted func(ContractTermCompletedEvent) error

	OnContractTermCreated func(ContractTermCreatedEvent) error

	OnContractTermRenewed func(ContractTermRenewedEvent) error

	OnContractTermTerminated func(ContractTermTerminatedEvent) error

	OnCouponCodesAdded func(CouponCodesAddedEvent) error

	OnCouponCodesDeleted func(CouponCodesDeletedEvent) error

	OnCouponCodesUpdated func(CouponCodesUpdatedEvent) error

	OnCouponCreated func(CouponCreatedEvent) error

	OnCouponDeleted func(CouponDeletedEvent) error

	OnCouponSetCreated func(CouponSetCreatedEvent) error

	OnCouponSetDeleted func(CouponSetDeletedEvent) error

	OnCouponSetUpdated func(CouponSetUpdatedEvent) error

	OnCouponUpdated func(CouponUpdatedEvent) error

	OnCreditNoteCreated func(CreditNoteCreatedEvent) error

	OnCreditNoteCreatedWithBackdating func(CreditNoteCreatedWithBackdatingEvent) error

	OnCreditNoteDeleted func(CreditNoteDeletedEvent) error

	OnCreditNoteUpdated func(CreditNoteUpdatedEvent) error

	OnCustomerBusinessEntityChanged func(CustomerBusinessEntityChangedEvent) error

	OnCustomerChanged func(CustomerChangedEvent) error

	OnCustomerCreated func(CustomerCreatedEvent) error

	OnCustomerDeleted func(CustomerDeletedEvent) error

	OnCustomerEntitlementsUpdated func(CustomerEntitlementsUpdatedEvent) error

	OnCustomerMovedIn func(CustomerMovedInEvent) error

	OnCustomerMovedOut func(CustomerMovedOutEvent) error

	OnDifferentialPriceCreated func(DifferentialPriceCreatedEvent) error

	OnDifferentialPriceDeleted func(DifferentialPriceDeletedEvent) error

	OnDifferentialPriceUpdated func(DifferentialPriceUpdatedEvent) error

	OnDunningUpdated func(DunningUpdatedEvent) error

	OnEntitlementOverridesAutoRemoved func(EntitlementOverridesAutoRemovedEvent) error

	OnEntitlementOverridesRemoved func(EntitlementOverridesRemovedEvent) error

	OnEntitlementOverridesUpdated func(EntitlementOverridesUpdatedEvent) error

	OnFeatureActivated func(FeatureActivatedEvent) error

	OnFeatureArchived func(FeatureArchivedEvent) error

	OnFeatureCreated func(FeatureCreatedEvent) error

	OnFeatureDeleted func(FeatureDeletedEvent) error

	OnFeatureReactivated func(FeatureReactivatedEvent) error

	OnFeatureUpdated func(FeatureUpdatedEvent) error

	OnGiftCancelled func(GiftCancelledEvent) error

	OnGiftClaimed func(GiftClaimedEvent) error

	OnGiftExpired func(GiftExpiredEvent) error

	OnGiftScheduled func(GiftScheduledEvent) error

	OnGiftUnclaimed func(GiftUnclaimedEvent) error

	OnGiftUpdated func(GiftUpdatedEvent) error

	OnHierarchyCreated func(HierarchyCreatedEvent) error

	OnHierarchyDeleted func(HierarchyDeletedEvent) error

	OnInvoiceDeleted func(InvoiceDeletedEvent) error

	OnInvoiceGenerated func(InvoiceGeneratedEvent) error

	OnInvoiceGeneratedWithBackdating func(InvoiceGeneratedWithBackdatingEvent) error

	OnInvoiceUpdated func(InvoiceUpdatedEvent) error

	OnItemCreated func(ItemCreatedEvent) error

	OnItemDeleted func(ItemDeletedEvent) error

	OnItemEntitlementsRemoved func(ItemEntitlementsRemovedEvent) error

	OnItemEntitlementsUpdated func(ItemEntitlementsUpdatedEvent) error

	OnItemFamilyCreated func(ItemFamilyCreatedEvent) error

	OnItemFamilyDeleted func(ItemFamilyDeletedEvent) error

	OnItemFamilyUpdated func(ItemFamilyUpdatedEvent) error

	OnItemPriceCreated func(ItemPriceCreatedEvent) error

	OnItemPriceDeleted func(ItemPriceDeletedEvent) error

	OnItemPriceEntitlementsRemoved func(ItemPriceEntitlementsRemovedEvent) error

	OnItemPriceEntitlementsUpdated func(ItemPriceEntitlementsUpdatedEvent) error

	OnItemPriceUpdated func(ItemPriceUpdatedEvent) error

	OnItemUpdated func(ItemUpdatedEvent) error

	OnMrrUpdated func(MrrUpdatedEvent) error

	OnNetdPaymentDueReminder func(NetdPaymentDueReminderEvent) error

	OnOmnichannelOneTimeOrderCreated func(OmnichannelOneTimeOrderCreatedEvent) error

	OnOmnichannelOneTimeOrderItemCancelled func(OmnichannelOneTimeOrderItemCancelledEvent) error

	OnOmnichannelSubscriptionCreated func(OmnichannelSubscriptionCreatedEvent) error

	OnOmnichannelSubscriptionImported func(OmnichannelSubscriptionImportedEvent) error

	OnOmnichannelSubscriptionItemCancellationScheduled func(OmnichannelSubscriptionItemCancellationScheduledEvent) error

	OnOmnichannelSubscriptionItemCancelled func(OmnichannelSubscriptionItemCancelledEvent) error

	OnOmnichannelSubscriptionItemChangeScheduled func(OmnichannelSubscriptionItemChangeScheduledEvent) error

	OnOmnichannelSubscriptionItemChanged func(OmnichannelSubscriptionItemChangedEvent) error

	OnOmnichannelSubscriptionItemDowngradeScheduled func(OmnichannelSubscriptionItemDowngradeScheduledEvent) error

	OnOmnichannelSubscriptionItemDowngraded func(OmnichannelSubscriptionItemDowngradedEvent) error

	OnOmnichannelSubscriptionItemDunningExpired func(OmnichannelSubscriptionItemDunningExpiredEvent) error

	OnOmnichannelSubscriptionItemDunningStarted func(OmnichannelSubscriptionItemDunningStartedEvent) error

	OnOmnichannelSubscriptionItemExpired func(OmnichannelSubscriptionItemExpiredEvent) error

	OnOmnichannelSubscriptionItemGracePeriodExpired func(OmnichannelSubscriptionItemGracePeriodExpiredEvent) error

	OnOmnichannelSubscriptionItemGracePeriodStarted func(OmnichannelSubscriptionItemGracePeriodStartedEvent) error

	OnOmnichannelSubscriptionItemPauseScheduled func(OmnichannelSubscriptionItemPauseScheduledEvent) error

	OnOmnichannelSubscriptionItemPaused func(OmnichannelSubscriptionItemPausedEvent) error

	OnOmnichannelSubscriptionItemReactivated func(OmnichannelSubscriptionItemReactivatedEvent) error

	OnOmnichannelSubscriptionItemRenewed func(OmnichannelSubscriptionItemRenewedEvent) error

	OnOmnichannelSubscriptionItemResubscribed func(OmnichannelSubscriptionItemResubscribedEvent) error

	OnOmnichannelSubscriptionItemResumed func(OmnichannelSubscriptionItemResumedEvent) error

	OnOmnichannelSubscriptionItemScheduledCancellationRemoved func(OmnichannelSubscriptionItemScheduledCancellationRemovedEvent) error

	OnOmnichannelSubscriptionItemScheduledChangeRemoved func(OmnichannelSubscriptionItemScheduledChangeRemovedEvent) error

	OnOmnichannelSubscriptionItemScheduledDowngradeRemoved func(OmnichannelSubscriptionItemScheduledDowngradeRemovedEvent) error

	OnOmnichannelSubscriptionItemUpgraded func(OmnichannelSubscriptionItemUpgradedEvent) error

	OnOmnichannelSubscriptionMovedIn func(OmnichannelSubscriptionMovedInEvent) error

	OnOmnichannelTransactionCreated func(OmnichannelTransactionCreatedEvent) error

	OnOrderCancelled func(OrderCancelledEvent) error

	OnOrderCreated func(OrderCreatedEvent) error

	OnOrderDeleted func(OrderDeletedEvent) error

	OnOrderDelivered func(OrderDeliveredEvent) error

	OnOrderReadyToProcess func(OrderReadyToProcessEvent) error

	OnOrderReadyToShip func(OrderReadyToShipEvent) error

	OnOrderResent func(OrderResentEvent) error

	OnOrderReturned func(OrderReturnedEvent) error

	OnOrderUpdated func(OrderUpdatedEvent) error

	OnPaymentFailed func(PaymentFailedEvent) error

	OnPaymentInitiated func(PaymentInitiatedEvent) error

	OnPaymentIntentCreated func(PaymentIntentCreatedEvent) error

	OnPaymentIntentUpdated func(PaymentIntentUpdatedEvent) error

	OnPaymentRefunded func(PaymentRefundedEvent) error

	OnPaymentScheduleSchemeCreated func(PaymentScheduleSchemeCreatedEvent) error

	OnPaymentScheduleSchemeDeleted func(PaymentScheduleSchemeDeletedEvent) error

	OnPaymentSchedulesCreated func(PaymentSchedulesCreatedEvent) error

	OnPaymentSchedulesUpdated func(PaymentSchedulesUpdatedEvent) error

	OnPaymentSourceAdded func(PaymentSourceAddedEvent) error

	OnPaymentSourceDeleted func(PaymentSourceDeletedEvent) error

	OnPaymentSourceExpired func(PaymentSourceExpiredEvent) error

	OnPaymentSourceExpiring func(PaymentSourceExpiringEvent) error

	OnPaymentSourceLocallyDeleted func(PaymentSourceLocallyDeletedEvent) error

	OnPaymentSourceUpdated func(PaymentSourceUpdatedEvent) error

	OnPaymentSucceeded func(PaymentSucceededEvent) error

	OnPendingInvoiceCreated func(PendingInvoiceCreatedEvent) error

	OnPendingInvoiceUpdated func(PendingInvoiceUpdatedEvent) error

	OnPlanCreated func(PlanCreatedEvent) error

	OnPlanDeleted func(PlanDeletedEvent) error

	OnPlanUpdated func(PlanUpdatedEvent) error

	OnPriceVariantCreated func(PriceVariantCreatedEvent) error

	OnPriceVariantDeleted func(PriceVariantDeletedEvent) error

	OnPriceVariantUpdated func(PriceVariantUpdatedEvent) error

	OnProductCreated func(ProductCreatedEvent) error

	OnProductDeleted func(ProductDeletedEvent) error

	OnProductUpdated func(ProductUpdatedEvent) error

	OnPromotionalCreditsAdded func(PromotionalCreditsAddedEvent) error

	OnPromotionalCreditsDeducted func(PromotionalCreditsDeductedEvent) error

	OnPurchaseCreated func(PurchaseCreatedEvent) error

	OnQuoteCreated func(QuoteCreatedEvent) error

	OnQuoteDeleted func(QuoteDeletedEvent) error

	OnQuoteUpdated func(QuoteUpdatedEvent) error

	OnRecordPurchaseFailed func(RecordPurchaseFailedEvent) error

	OnRefundInitiated func(RefundInitiatedEvent) error

	OnRuleCreated func(RuleCreatedEvent) error

	OnRuleDeleted func(RuleDeletedEvent) error

	OnRuleUpdated func(RuleUpdatedEvent) error

	OnSalesOrderCreated func(SalesOrderCreatedEvent) error

	OnSalesOrderUpdated func(SalesOrderUpdatedEvent) error

	OnSubscriptionActivated func(SubscriptionActivatedEvent) error

	OnSubscriptionActivatedWithBackdating func(SubscriptionActivatedWithBackdatingEvent) error

	OnSubscriptionAdvanceInvoiceScheduleAdded func(SubscriptionAdvanceInvoiceScheduleAddedEvent) error

	OnSubscriptionAdvanceInvoiceScheduleRemoved func(SubscriptionAdvanceInvoiceScheduleRemovedEvent) error

	OnSubscriptionAdvanceInvoiceScheduleUpdated func(SubscriptionAdvanceInvoiceScheduleUpdatedEvent) error

	OnSubscriptionBusinessEntityChanged func(SubscriptionBusinessEntityChangedEvent) error

	OnSubscriptionCanceledWithBackdating func(SubscriptionCanceledWithBackdatingEvent) error

	OnSubscriptionCancellationReminder func(SubscriptionCancellationReminderEvent) error

	OnSubscriptionCancellationScheduled func(SubscriptionCancellationScheduledEvent) error

	OnSubscriptionCancelled func(SubscriptionCancelledEvent) error

	OnSubscriptionChanged func(SubscriptionChangedEvent) error

	OnSubscriptionChangedWithBackdating func(SubscriptionChangedWithBackdatingEvent) error

	OnSubscriptionChangesScheduled func(SubscriptionChangesScheduledEvent) error

	OnSubscriptionCreated func(SubscriptionCreatedEvent) error

	OnSubscriptionCreatedWithBackdating func(SubscriptionCreatedWithBackdatingEvent) error

	OnSubscriptionDeleted func(SubscriptionDeletedEvent) error

	OnSubscriptionEntitlementsCreated func(SubscriptionEntitlementsCreatedEvent) error

	OnSubscriptionEntitlementsUpdated func(SubscriptionEntitlementsUpdatedEvent) error

	OnSubscriptionItemsRenewed func(SubscriptionItemsRenewedEvent) error

	OnSubscriptionMovedIn func(SubscriptionMovedInEvent) error

	OnSubscriptionMovedOut func(SubscriptionMovedOutEvent) error

	OnSubscriptionMovementFailed func(SubscriptionMovementFailedEvent) error

	OnSubscriptionPauseScheduled func(SubscriptionPauseScheduledEvent) error

	OnSubscriptionPaused func(SubscriptionPausedEvent) error

	OnSubscriptionRampApplied func(SubscriptionRampAppliedEvent) error

	OnSubscriptionRampCreated func(SubscriptionRampCreatedEvent) error

	OnSubscriptionRampDeleted func(SubscriptionRampDeletedEvent) error

	OnSubscriptionRampDrafted func(SubscriptionRampDraftedEvent) error

	OnSubscriptionRampUpdated func(SubscriptionRampUpdatedEvent) error

	OnSubscriptionReactivated func(SubscriptionReactivatedEvent) error

	OnSubscriptionReactivatedWithBackdating func(SubscriptionReactivatedWithBackdatingEvent) error

	OnSubscriptionRenewalReminder func(SubscriptionRenewalReminderEvent) error

	OnSubscriptionRenewed func(SubscriptionRenewedEvent) error

	OnSubscriptionResumed func(SubscriptionResumedEvent) error

	OnSubscriptionResumptionScheduled func(SubscriptionResumptionScheduledEvent) error

	OnSubscriptionScheduledCancellationRemoved func(SubscriptionScheduledCancellationRemovedEvent) error

	OnSubscriptionScheduledChangesRemoved func(SubscriptionScheduledChangesRemovedEvent) error

	OnSubscriptionScheduledPauseRemoved func(SubscriptionScheduledPauseRemovedEvent) error

	OnSubscriptionScheduledResumptionRemoved func(SubscriptionScheduledResumptionRemovedEvent) error

	OnSubscriptionShippingAddressUpdated func(SubscriptionShippingAddressUpdatedEvent) error

	OnSubscriptionStarted func(SubscriptionStartedEvent) error

	OnSubscriptionTrialEndReminder func(SubscriptionTrialEndReminderEvent) error

	OnSubscriptionTrialExtended func(SubscriptionTrialExtendedEvent) error

	OnTaxWithheldDeleted func(TaxWithheldDeletedEvent) error

	OnTaxWithheldRecorded func(TaxWithheldRecordedEvent) error

	OnTaxWithheldRefunded func(TaxWithheldRefundedEvent) error

	OnTokenConsumed func(TokenConsumedEvent) error

	OnTokenCreated func(TokenCreatedEvent) error

	OnTokenExpired func(TokenExpiredEvent) error

	OnTransactionCreated func(TransactionCreatedEvent) error

	OnTransactionDeleted func(TransactionDeletedEvent) error

	OnTransactionUpdated func(TransactionUpdatedEvent) error

	OnUnbilledChargesCreated func(UnbilledChargesCreatedEvent) error

	OnUnbilledChargesDeleted func(UnbilledChargesDeletedEvent) error

	OnUnbilledChargesInvoiced func(UnbilledChargesInvoicedEvent) error

	OnUnbilledChargesVoided func(UnbilledChargesVoidedEvent) error

	OnUsageFileIngested func(UsageFileIngestedEvent) error

	OnVariantCreated func(VariantCreatedEvent) error

	OnVariantDeleted func(VariantDeletedEvent) error

	OnVariantUpdated func(VariantUpdatedEvent) error

	OnVirtualBankAccountAdded func(VirtualBankAccountAddedEvent) error

	OnVirtualBankAccountDeleted func(VirtualBankAccountDeletedEvent) error

	OnVirtualBankAccountUpdated func(VirtualBankAccountUpdatedEvent) error

	OnVoucherCreateFailed func(VoucherCreateFailedEvent) error

	OnVoucherCreated func(VoucherCreatedEvent) error

	OnVoucherExpired func(VoucherExpiredEvent) error
}

func (h *WebhookHandler) ParseAndDispatch(body []byte) error {
	if h == nil {
		return errors.New("webhook handler not configured")
	}
	eventType, err := ParseEventType(body)
	if err != nil {
		return err
	}
	switch eventType {

	case chargebee.EventTypeAddUsagesReminder:
		if h.OnAddUsagesReminder != nil {
			var e AddUsagesReminderEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnAddUsagesReminder(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeAddonCreated:
		if h.OnAddonCreated != nil {
			var e AddonCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnAddonCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeAddonDeleted:
		if h.OnAddonDeleted != nil {
			var e AddonDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnAddonDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeAddonUpdated:
		if h.OnAddonUpdated != nil {
			var e AddonUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnAddonUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeAttachedItemCreated:
		if h.OnAttachedItemCreated != nil {
			var e AttachedItemCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnAttachedItemCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeAttachedItemDeleted:
		if h.OnAttachedItemDeleted != nil {
			var e AttachedItemDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnAttachedItemDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeAttachedItemUpdated:
		if h.OnAttachedItemUpdated != nil {
			var e AttachedItemUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnAttachedItemUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeAuthorizationSucceeded:
		if h.OnAuthorizationSucceeded != nil {
			var e AuthorizationSucceededEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnAuthorizationSucceeded(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeAuthorizationVoided:
		if h.OnAuthorizationVoided != nil {
			var e AuthorizationVoidedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnAuthorizationVoided(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeBusinessEntityCreated:
		if h.OnBusinessEntityCreated != nil {
			var e BusinessEntityCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnBusinessEntityCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeBusinessEntityDeleted:
		if h.OnBusinessEntityDeleted != nil {
			var e BusinessEntityDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnBusinessEntityDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeBusinessEntityUpdated:
		if h.OnBusinessEntityUpdated != nil {
			var e BusinessEntityUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnBusinessEntityUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCardAdded:
		if h.OnCardAdded != nil {
			var e CardAddedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCardAdded(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCardDeleted:
		if h.OnCardDeleted != nil {
			var e CardDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCardDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCardExpired:
		if h.OnCardExpired != nil {
			var e CardExpiredEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCardExpired(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCardExpiryReminder:
		if h.OnCardExpiryReminder != nil {
			var e CardExpiryReminderEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCardExpiryReminder(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCardUpdated:
		if h.OnCardUpdated != nil {
			var e CardUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCardUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeContractTermCancelled:
		if h.OnContractTermCancelled != nil {
			var e ContractTermCancelledEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnContractTermCancelled(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeContractTermCompleted:
		if h.OnContractTermCompleted != nil {
			var e ContractTermCompletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnContractTermCompleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeContractTermCreated:
		if h.OnContractTermCreated != nil {
			var e ContractTermCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnContractTermCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeContractTermRenewed:
		if h.OnContractTermRenewed != nil {
			var e ContractTermRenewedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnContractTermRenewed(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeContractTermTerminated:
		if h.OnContractTermTerminated != nil {
			var e ContractTermTerminatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnContractTermTerminated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCouponCodesAdded:
		if h.OnCouponCodesAdded != nil {
			var e CouponCodesAddedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCouponCodesAdded(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCouponCodesDeleted:
		if h.OnCouponCodesDeleted != nil {
			var e CouponCodesDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCouponCodesDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCouponCodesUpdated:
		if h.OnCouponCodesUpdated != nil {
			var e CouponCodesUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCouponCodesUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCouponCreated:
		if h.OnCouponCreated != nil {
			var e CouponCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCouponCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCouponDeleted:
		if h.OnCouponDeleted != nil {
			var e CouponDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCouponDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCouponSetCreated:
		if h.OnCouponSetCreated != nil {
			var e CouponSetCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCouponSetCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCouponSetDeleted:
		if h.OnCouponSetDeleted != nil {
			var e CouponSetDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCouponSetDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCouponSetUpdated:
		if h.OnCouponSetUpdated != nil {
			var e CouponSetUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCouponSetUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCouponUpdated:
		if h.OnCouponUpdated != nil {
			var e CouponUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCouponUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCreditNoteCreated:
		if h.OnCreditNoteCreated != nil {
			var e CreditNoteCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCreditNoteCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCreditNoteCreatedWithBackdating:
		if h.OnCreditNoteCreatedWithBackdating != nil {
			var e CreditNoteCreatedWithBackdatingEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCreditNoteCreatedWithBackdating(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCreditNoteDeleted:
		if h.OnCreditNoteDeleted != nil {
			var e CreditNoteDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCreditNoteDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCreditNoteUpdated:
		if h.OnCreditNoteUpdated != nil {
			var e CreditNoteUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCreditNoteUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCustomerBusinessEntityChanged:
		if h.OnCustomerBusinessEntityChanged != nil {
			var e CustomerBusinessEntityChangedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCustomerBusinessEntityChanged(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCustomerChanged:
		if h.OnCustomerChanged != nil {
			var e CustomerChangedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCustomerChanged(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCustomerCreated:
		if h.OnCustomerCreated != nil {
			var e CustomerCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCustomerCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCustomerDeleted:
		if h.OnCustomerDeleted != nil {
			var e CustomerDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCustomerDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCustomerEntitlementsUpdated:
		if h.OnCustomerEntitlementsUpdated != nil {
			var e CustomerEntitlementsUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCustomerEntitlementsUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCustomerMovedIn:
		if h.OnCustomerMovedIn != nil {
			var e CustomerMovedInEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCustomerMovedIn(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeCustomerMovedOut:
		if h.OnCustomerMovedOut != nil {
			var e CustomerMovedOutEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnCustomerMovedOut(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeDifferentialPriceCreated:
		if h.OnDifferentialPriceCreated != nil {
			var e DifferentialPriceCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnDifferentialPriceCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeDifferentialPriceDeleted:
		if h.OnDifferentialPriceDeleted != nil {
			var e DifferentialPriceDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnDifferentialPriceDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeDifferentialPriceUpdated:
		if h.OnDifferentialPriceUpdated != nil {
			var e DifferentialPriceUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnDifferentialPriceUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeDunningUpdated:
		if h.OnDunningUpdated != nil {
			var e DunningUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnDunningUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeEntitlementOverridesAutoRemoved:
		if h.OnEntitlementOverridesAutoRemoved != nil {
			var e EntitlementOverridesAutoRemovedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnEntitlementOverridesAutoRemoved(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeEntitlementOverridesRemoved:
		if h.OnEntitlementOverridesRemoved != nil {
			var e EntitlementOverridesRemovedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnEntitlementOverridesRemoved(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeEntitlementOverridesUpdated:
		if h.OnEntitlementOverridesUpdated != nil {
			var e EntitlementOverridesUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnEntitlementOverridesUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeFeatureActivated:
		if h.OnFeatureActivated != nil {
			var e FeatureActivatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnFeatureActivated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeFeatureArchived:
		if h.OnFeatureArchived != nil {
			var e FeatureArchivedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnFeatureArchived(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeFeatureCreated:
		if h.OnFeatureCreated != nil {
			var e FeatureCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnFeatureCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeFeatureDeleted:
		if h.OnFeatureDeleted != nil {
			var e FeatureDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnFeatureDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeFeatureReactivated:
		if h.OnFeatureReactivated != nil {
			var e FeatureReactivatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnFeatureReactivated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeFeatureUpdated:
		if h.OnFeatureUpdated != nil {
			var e FeatureUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnFeatureUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeGiftCancelled:
		if h.OnGiftCancelled != nil {
			var e GiftCancelledEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnGiftCancelled(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeGiftClaimed:
		if h.OnGiftClaimed != nil {
			var e GiftClaimedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnGiftClaimed(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeGiftExpired:
		if h.OnGiftExpired != nil {
			var e GiftExpiredEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnGiftExpired(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeGiftScheduled:
		if h.OnGiftScheduled != nil {
			var e GiftScheduledEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnGiftScheduled(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeGiftUnclaimed:
		if h.OnGiftUnclaimed != nil {
			var e GiftUnclaimedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnGiftUnclaimed(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeGiftUpdated:
		if h.OnGiftUpdated != nil {
			var e GiftUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnGiftUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeHierarchyCreated:
		if h.OnHierarchyCreated != nil {
			var e HierarchyCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnHierarchyCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeHierarchyDeleted:
		if h.OnHierarchyDeleted != nil {
			var e HierarchyDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnHierarchyDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeInvoiceDeleted:
		if h.OnInvoiceDeleted != nil {
			var e InvoiceDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnInvoiceDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeInvoiceGenerated:
		if h.OnInvoiceGenerated != nil {
			var e InvoiceGeneratedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnInvoiceGenerated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeInvoiceGeneratedWithBackdating:
		if h.OnInvoiceGeneratedWithBackdating != nil {
			var e InvoiceGeneratedWithBackdatingEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnInvoiceGeneratedWithBackdating(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeInvoiceUpdated:
		if h.OnInvoiceUpdated != nil {
			var e InvoiceUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnInvoiceUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeItemCreated:
		if h.OnItemCreated != nil {
			var e ItemCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnItemCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeItemDeleted:
		if h.OnItemDeleted != nil {
			var e ItemDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnItemDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeItemEntitlementsRemoved:
		if h.OnItemEntitlementsRemoved != nil {
			var e ItemEntitlementsRemovedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnItemEntitlementsRemoved(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeItemEntitlementsUpdated:
		if h.OnItemEntitlementsUpdated != nil {
			var e ItemEntitlementsUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnItemEntitlementsUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeItemFamilyCreated:
		if h.OnItemFamilyCreated != nil {
			var e ItemFamilyCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnItemFamilyCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeItemFamilyDeleted:
		if h.OnItemFamilyDeleted != nil {
			var e ItemFamilyDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnItemFamilyDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeItemFamilyUpdated:
		if h.OnItemFamilyUpdated != nil {
			var e ItemFamilyUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnItemFamilyUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeItemPriceCreated:
		if h.OnItemPriceCreated != nil {
			var e ItemPriceCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnItemPriceCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeItemPriceDeleted:
		if h.OnItemPriceDeleted != nil {
			var e ItemPriceDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnItemPriceDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeItemPriceEntitlementsRemoved:
		if h.OnItemPriceEntitlementsRemoved != nil {
			var e ItemPriceEntitlementsRemovedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnItemPriceEntitlementsRemoved(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeItemPriceEntitlementsUpdated:
		if h.OnItemPriceEntitlementsUpdated != nil {
			var e ItemPriceEntitlementsUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnItemPriceEntitlementsUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeItemPriceUpdated:
		if h.OnItemPriceUpdated != nil {
			var e ItemPriceUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnItemPriceUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeItemUpdated:
		if h.OnItemUpdated != nil {
			var e ItemUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnItemUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeMrrUpdated:
		if h.OnMrrUpdated != nil {
			var e MrrUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnMrrUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelOneTimeOrderCreated:
		if h.OnOmnichannelOneTimeOrderCreated != nil {
			var e OmnichannelOneTimeOrderCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelOneTimeOrderCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelOneTimeOrderItemCancelled:
		if h.OnOmnichannelOneTimeOrderItemCancelled != nil {
			var e OmnichannelOneTimeOrderItemCancelledEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelOneTimeOrderItemCancelled(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionCreated:
		if h.OnOmnichannelSubscriptionCreated != nil {
			var e OmnichannelSubscriptionCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionImported:
		if h.OnOmnichannelSubscriptionImported != nil {
			var e OmnichannelSubscriptionImportedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionImported(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionItemCancellationScheduled:
		if h.OnOmnichannelSubscriptionItemCancellationScheduled != nil {
			var e OmnichannelSubscriptionItemCancellationScheduledEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionItemCancellationScheduled(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionItemCancelled:
		if h.OnOmnichannelSubscriptionItemCancelled != nil {
			var e OmnichannelSubscriptionItemCancelledEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionItemCancelled(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionItemChangeScheduled:
		if h.OnOmnichannelSubscriptionItemChangeScheduled != nil {
			var e OmnichannelSubscriptionItemChangeScheduledEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionItemChangeScheduled(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionItemChanged:
		if h.OnOmnichannelSubscriptionItemChanged != nil {
			var e OmnichannelSubscriptionItemChangedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionItemChanged(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionItemDowngraded:
		if h.OnOmnichannelSubscriptionItemDowngraded != nil {
			var e OmnichannelSubscriptionItemDowngradedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionItemDowngraded(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionItemDunningExpired:
		if h.OnOmnichannelSubscriptionItemDunningExpired != nil {
			var e OmnichannelSubscriptionItemDunningExpiredEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionItemDunningExpired(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionItemDunningStarted:
		if h.OnOmnichannelSubscriptionItemDunningStarted != nil {
			var e OmnichannelSubscriptionItemDunningStartedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionItemDunningStarted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionItemExpired:
		if h.OnOmnichannelSubscriptionItemExpired != nil {
			var e OmnichannelSubscriptionItemExpiredEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionItemExpired(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionItemGracePeriodExpired:
		if h.OnOmnichannelSubscriptionItemGracePeriodExpired != nil {
			var e OmnichannelSubscriptionItemGracePeriodExpiredEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionItemGracePeriodExpired(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionItemGracePeriodStarted:
		if h.OnOmnichannelSubscriptionItemGracePeriodStarted != nil {
			var e OmnichannelSubscriptionItemGracePeriodStartedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionItemGracePeriodStarted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionItemPauseScheduled:
		if h.OnOmnichannelSubscriptionItemPauseScheduled != nil {
			var e OmnichannelSubscriptionItemPauseScheduledEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionItemPauseScheduled(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionItemPaused:
		if h.OnOmnichannelSubscriptionItemPaused != nil {
			var e OmnichannelSubscriptionItemPausedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionItemPaused(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionItemReactivated:
		if h.OnOmnichannelSubscriptionItemReactivated != nil {
			var e OmnichannelSubscriptionItemReactivatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionItemReactivated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionItemRenewed:
		if h.OnOmnichannelSubscriptionItemRenewed != nil {
			var e OmnichannelSubscriptionItemRenewedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionItemRenewed(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionItemResubscribed:
		if h.OnOmnichannelSubscriptionItemResubscribed != nil {
			var e OmnichannelSubscriptionItemResubscribedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionItemResubscribed(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionItemResumed:
		if h.OnOmnichannelSubscriptionItemResumed != nil {
			var e OmnichannelSubscriptionItemResumedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionItemResumed(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionItemScheduledCancellationRemoved:
		if h.OnOmnichannelSubscriptionItemScheduledCancellationRemoved != nil {
			var e OmnichannelSubscriptionItemScheduledCancellationRemovedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionItemScheduledCancellationRemoved(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionItemScheduledChangeRemoved:
		if h.OnOmnichannelSubscriptionItemScheduledChangeRemoved != nil {
			var e OmnichannelSubscriptionItemScheduledChangeRemovedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionItemScheduledChangeRemoved(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionItemUpgraded:
		if h.OnOmnichannelSubscriptionItemUpgraded != nil {
			var e OmnichannelSubscriptionItemUpgradedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionItemUpgraded(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelSubscriptionMovedIn:
		if h.OnOmnichannelSubscriptionMovedIn != nil {
			var e OmnichannelSubscriptionMovedInEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelSubscriptionMovedIn(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOmnichannelTransactionCreated:
		if h.OnOmnichannelTransactionCreated != nil {
			var e OmnichannelTransactionCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOmnichannelTransactionCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOrderCancelled:
		if h.OnOrderCancelled != nil {
			var e OrderCancelledEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOrderCancelled(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOrderCreated:
		if h.OnOrderCreated != nil {
			var e OrderCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOrderCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOrderDeleted:
		if h.OnOrderDeleted != nil {
			var e OrderDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOrderDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOrderDelivered:
		if h.OnOrderDelivered != nil {
			var e OrderDeliveredEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOrderDelivered(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOrderReadyToProcess:
		if h.OnOrderReadyToProcess != nil {
			var e OrderReadyToProcessEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOrderReadyToProcess(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOrderReadyToShip:
		if h.OnOrderReadyToShip != nil {
			var e OrderReadyToShipEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOrderReadyToShip(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOrderResent:
		if h.OnOrderResent != nil {
			var e OrderResentEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOrderResent(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOrderReturned:
		if h.OnOrderReturned != nil {
			var e OrderReturnedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOrderReturned(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeOrderUpdated:
		if h.OnOrderUpdated != nil {
			var e OrderUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnOrderUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePaymentFailed:
		if h.OnPaymentFailed != nil {
			var e PaymentFailedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPaymentFailed(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePaymentInitiated:
		if h.OnPaymentInitiated != nil {
			var e PaymentInitiatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPaymentInitiated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePaymentIntentCreated:
		if h.OnPaymentIntentCreated != nil {
			var e PaymentIntentCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPaymentIntentCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePaymentIntentUpdated:
		if h.OnPaymentIntentUpdated != nil {
			var e PaymentIntentUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPaymentIntentUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePaymentRefunded:
		if h.OnPaymentRefunded != nil {
			var e PaymentRefundedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPaymentRefunded(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePaymentScheduleSchemeCreated:
		if h.OnPaymentScheduleSchemeCreated != nil {
			var e PaymentScheduleSchemeCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPaymentScheduleSchemeCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePaymentScheduleSchemeDeleted:
		if h.OnPaymentScheduleSchemeDeleted != nil {
			var e PaymentScheduleSchemeDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPaymentScheduleSchemeDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePaymentSchedulesCreated:
		if h.OnPaymentSchedulesCreated != nil {
			var e PaymentSchedulesCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPaymentSchedulesCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePaymentSchedulesUpdated:
		if h.OnPaymentSchedulesUpdated != nil {
			var e PaymentSchedulesUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPaymentSchedulesUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePaymentSourceAdded:
		if h.OnPaymentSourceAdded != nil {
			var e PaymentSourceAddedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPaymentSourceAdded(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePaymentSourceDeleted:
		if h.OnPaymentSourceDeleted != nil {
			var e PaymentSourceDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPaymentSourceDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePaymentSourceExpired:
		if h.OnPaymentSourceExpired != nil {
			var e PaymentSourceExpiredEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPaymentSourceExpired(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePaymentSourceExpiring:
		if h.OnPaymentSourceExpiring != nil {
			var e PaymentSourceExpiringEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPaymentSourceExpiring(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePaymentSourceLocallyDeleted:
		if h.OnPaymentSourceLocallyDeleted != nil {
			var e PaymentSourceLocallyDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPaymentSourceLocallyDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePaymentSourceUpdated:
		if h.OnPaymentSourceUpdated != nil {
			var e PaymentSourceUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPaymentSourceUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePaymentSucceeded:
		if h.OnPaymentSucceeded != nil {
			var e PaymentSucceededEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPaymentSucceeded(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePendingInvoiceCreated:
		if h.OnPendingInvoiceCreated != nil {
			var e PendingInvoiceCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPendingInvoiceCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePendingInvoiceUpdated:
		if h.OnPendingInvoiceUpdated != nil {
			var e PendingInvoiceUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPendingInvoiceUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePlanCreated:
		if h.OnPlanCreated != nil {
			var e PlanCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPlanCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePlanDeleted:
		if h.OnPlanDeleted != nil {
			var e PlanDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPlanDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePlanUpdated:
		if h.OnPlanUpdated != nil {
			var e PlanUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPlanUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePriceVariantCreated:
		if h.OnPriceVariantCreated != nil {
			var e PriceVariantCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPriceVariantCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePriceVariantDeleted:
		if h.OnPriceVariantDeleted != nil {
			var e PriceVariantDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPriceVariantDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePriceVariantUpdated:
		if h.OnPriceVariantUpdated != nil {
			var e PriceVariantUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPriceVariantUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePromotionalCreditsAdded:
		if h.OnPromotionalCreditsAdded != nil {
			var e PromotionalCreditsAddedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPromotionalCreditsAdded(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePromotionalCreditsDeducted:
		if h.OnPromotionalCreditsDeducted != nil {
			var e PromotionalCreditsDeductedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPromotionalCreditsDeducted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypePurchaseCreated:
		if h.OnPurchaseCreated != nil {
			var e PurchaseCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnPurchaseCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeQuoteCreated:
		if h.OnQuoteCreated != nil {
			var e QuoteCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnQuoteCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeQuoteDeleted:
		if h.OnQuoteDeleted != nil {
			var e QuoteDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnQuoteDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeQuoteUpdated:
		if h.OnQuoteUpdated != nil {
			var e QuoteUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnQuoteUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeRecordPurchaseFailed:
		if h.OnRecordPurchaseFailed != nil {
			var e RecordPurchaseFailedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnRecordPurchaseFailed(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeRefundInitiated:
		if h.OnRefundInitiated != nil {
			var e RefundInitiatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnRefundInitiated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeRuleCreated:
		if h.OnRuleCreated != nil {
			var e RuleCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnRuleCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeRuleDeleted:
		if h.OnRuleDeleted != nil {
			var e RuleDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnRuleDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeRuleUpdated:
		if h.OnRuleUpdated != nil {
			var e RuleUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnRuleUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSalesOrderCreated:
		if h.OnSalesOrderCreated != nil {
			var e SalesOrderCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSalesOrderCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSalesOrderUpdated:
		if h.OnSalesOrderUpdated != nil {
			var e SalesOrderUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSalesOrderUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionActivated:
		if h.OnSubscriptionActivated != nil {
			var e SubscriptionActivatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionActivated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionActivatedWithBackdating:
		if h.OnSubscriptionActivatedWithBackdating != nil {
			var e SubscriptionActivatedWithBackdatingEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionActivatedWithBackdating(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionAdvanceInvoiceScheduleAdded:
		if h.OnSubscriptionAdvanceInvoiceScheduleAdded != nil {
			var e SubscriptionAdvanceInvoiceScheduleAddedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionAdvanceInvoiceScheduleAdded(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionAdvanceInvoiceScheduleRemoved:
		if h.OnSubscriptionAdvanceInvoiceScheduleRemoved != nil {
			var e SubscriptionAdvanceInvoiceScheduleRemovedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionAdvanceInvoiceScheduleRemoved(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionAdvanceInvoiceScheduleUpdated:
		if h.OnSubscriptionAdvanceInvoiceScheduleUpdated != nil {
			var e SubscriptionAdvanceInvoiceScheduleUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionAdvanceInvoiceScheduleUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionBusinessEntityChanged:
		if h.OnSubscriptionBusinessEntityChanged != nil {
			var e SubscriptionBusinessEntityChangedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionBusinessEntityChanged(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionCanceledWithBackdating:
		if h.OnSubscriptionCanceledWithBackdating != nil {
			var e SubscriptionCanceledWithBackdatingEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionCanceledWithBackdating(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionCancellationReminder:
		if h.OnSubscriptionCancellationReminder != nil {
			var e SubscriptionCancellationReminderEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionCancellationReminder(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionCancellationScheduled:
		if h.OnSubscriptionCancellationScheduled != nil {
			var e SubscriptionCancellationScheduledEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionCancellationScheduled(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionCancelled:
		if h.OnSubscriptionCancelled != nil {
			var e SubscriptionCancelledEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionCancelled(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionChanged:
		if h.OnSubscriptionChanged != nil {
			var e SubscriptionChangedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionChanged(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionChangedWithBackdating:
		if h.OnSubscriptionChangedWithBackdating != nil {
			var e SubscriptionChangedWithBackdatingEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionChangedWithBackdating(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionChangesScheduled:
		if h.OnSubscriptionChangesScheduled != nil {
			var e SubscriptionChangesScheduledEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionChangesScheduled(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionCreated:
		if h.OnSubscriptionCreated != nil {
			var e SubscriptionCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionCreatedWithBackdating:
		if h.OnSubscriptionCreatedWithBackdating != nil {
			var e SubscriptionCreatedWithBackdatingEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionCreatedWithBackdating(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionDeleted:
		if h.OnSubscriptionDeleted != nil {
			var e SubscriptionDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionEntitlementsCreated:
		if h.OnSubscriptionEntitlementsCreated != nil {
			var e SubscriptionEntitlementsCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionEntitlementsCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionEntitlementsUpdated:
		if h.OnSubscriptionEntitlementsUpdated != nil {
			var e SubscriptionEntitlementsUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionEntitlementsUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionItemsRenewed:
		if h.OnSubscriptionItemsRenewed != nil {
			var e SubscriptionItemsRenewedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionItemsRenewed(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionMovedIn:
		if h.OnSubscriptionMovedIn != nil {
			var e SubscriptionMovedInEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionMovedIn(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionMovedOut:
		if h.OnSubscriptionMovedOut != nil {
			var e SubscriptionMovedOutEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionMovedOut(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionMovementFailed:
		if h.OnSubscriptionMovementFailed != nil {
			var e SubscriptionMovementFailedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionMovementFailed(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionPauseScheduled:
		if h.OnSubscriptionPauseScheduled != nil {
			var e SubscriptionPauseScheduledEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionPauseScheduled(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionPaused:
		if h.OnSubscriptionPaused != nil {
			var e SubscriptionPausedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionPaused(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionRampApplied:
		if h.OnSubscriptionRampApplied != nil {
			var e SubscriptionRampAppliedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionRampApplied(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionRampCreated:
		if h.OnSubscriptionRampCreated != nil {
			var e SubscriptionRampCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionRampCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionRampDeleted:
		if h.OnSubscriptionRampDeleted != nil {
			var e SubscriptionRampDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionRampDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionRampDrafted:
		if h.OnSubscriptionRampDrafted != nil {
			var e SubscriptionRampDraftedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionRampDrafted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionRampUpdated:
		if h.OnSubscriptionRampUpdated != nil {
			var e SubscriptionRampUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionRampUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionReactivated:
		if h.OnSubscriptionReactivated != nil {
			var e SubscriptionReactivatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionReactivated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionReactivatedWithBackdating:
		if h.OnSubscriptionReactivatedWithBackdating != nil {
			var e SubscriptionReactivatedWithBackdatingEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionReactivatedWithBackdating(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionRenewalReminder:
		if h.OnSubscriptionRenewalReminder != nil {
			var e SubscriptionRenewalReminderEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionRenewalReminder(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionRenewed:
		if h.OnSubscriptionRenewed != nil {
			var e SubscriptionRenewedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionRenewed(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionResumed:
		if h.OnSubscriptionResumed != nil {
			var e SubscriptionResumedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionResumed(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionResumptionScheduled:
		if h.OnSubscriptionResumptionScheduled != nil {
			var e SubscriptionResumptionScheduledEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionResumptionScheduled(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionScheduledCancellationRemoved:
		if h.OnSubscriptionScheduledCancellationRemoved != nil {
			var e SubscriptionScheduledCancellationRemovedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionScheduledCancellationRemoved(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionScheduledChangesRemoved:
		if h.OnSubscriptionScheduledChangesRemoved != nil {
			var e SubscriptionScheduledChangesRemovedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionScheduledChangesRemoved(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionScheduledPauseRemoved:
		if h.OnSubscriptionScheduledPauseRemoved != nil {
			var e SubscriptionScheduledPauseRemovedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionScheduledPauseRemoved(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionScheduledResumptionRemoved:
		if h.OnSubscriptionScheduledResumptionRemoved != nil {
			var e SubscriptionScheduledResumptionRemovedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionScheduledResumptionRemoved(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionShippingAddressUpdated:
		if h.OnSubscriptionShippingAddressUpdated != nil {
			var e SubscriptionShippingAddressUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionShippingAddressUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionStarted:
		if h.OnSubscriptionStarted != nil {
			var e SubscriptionStartedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionStarted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionTrialEndReminder:
		if h.OnSubscriptionTrialEndReminder != nil {
			var e SubscriptionTrialEndReminderEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionTrialEndReminder(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeSubscriptionTrialExtended:
		if h.OnSubscriptionTrialExtended != nil {
			var e SubscriptionTrialExtendedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnSubscriptionTrialExtended(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeTaxWithheldDeleted:
		if h.OnTaxWithheldDeleted != nil {
			var e TaxWithheldDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnTaxWithheldDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeTaxWithheldRecorded:
		if h.OnTaxWithheldRecorded != nil {
			var e TaxWithheldRecordedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnTaxWithheldRecorded(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeTaxWithheldRefunded:
		if h.OnTaxWithheldRefunded != nil {
			var e TaxWithheldRefundedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnTaxWithheldRefunded(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeTokenConsumed:
		if h.OnTokenConsumed != nil {
			var e TokenConsumedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnTokenConsumed(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeTokenCreated:
		if h.OnTokenCreated != nil {
			var e TokenCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnTokenCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeTokenExpired:
		if h.OnTokenExpired != nil {
			var e TokenExpiredEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnTokenExpired(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeTransactionCreated:
		if h.OnTransactionCreated != nil {
			var e TransactionCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnTransactionCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeTransactionDeleted:
		if h.OnTransactionDeleted != nil {
			var e TransactionDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnTransactionDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeTransactionUpdated:
		if h.OnTransactionUpdated != nil {
			var e TransactionUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnTransactionUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeUnbilledChargesCreated:
		if h.OnUnbilledChargesCreated != nil {
			var e UnbilledChargesCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnUnbilledChargesCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeUnbilledChargesDeleted:
		if h.OnUnbilledChargesDeleted != nil {
			var e UnbilledChargesDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnUnbilledChargesDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeUnbilledChargesInvoiced:
		if h.OnUnbilledChargesInvoiced != nil {
			var e UnbilledChargesInvoicedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnUnbilledChargesInvoiced(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeUnbilledChargesVoided:
		if h.OnUnbilledChargesVoided != nil {
			var e UnbilledChargesVoidedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnUnbilledChargesVoided(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeUsageFileIngested:
		if h.OnUsageFileIngested != nil {
			var e UsageFileIngestedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnUsageFileIngested(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeVirtualBankAccountAdded:
		if h.OnVirtualBankAccountAdded != nil {
			var e VirtualBankAccountAddedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnVirtualBankAccountAdded(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeVirtualBankAccountDeleted:
		if h.OnVirtualBankAccountDeleted != nil {
			var e VirtualBankAccountDeletedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnVirtualBankAccountDeleted(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeVirtualBankAccountUpdated:
		if h.OnVirtualBankAccountUpdated != nil {
			var e VirtualBankAccountUpdatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnVirtualBankAccountUpdated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeVoucherCreateFailed:
		if h.OnVoucherCreateFailed != nil {
			var e VoucherCreateFailedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnVoucherCreateFailed(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeVoucherCreated:
		if h.OnVoucherCreated != nil {
			var e VoucherCreatedEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnVoucherCreated(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	case chargebee.EventTypeVoucherExpired:
		if h.OnVoucherExpired != nil {
			var e VoucherExpiredEvent
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			return h.OnVoucherExpired(e)
		} else {
			return h.handleUnhandledEvent(eventType, body)
		}

	default:
		return h.handleUnhandledEvent(eventType, body)
	}
}

func (h *WebhookHandler) HTTPHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if h == nil {
			http.Error(w, "webhook handler not configured", http.StatusInternalServerError)
			return
		}
		if h.RequestValidator != nil {
			if err := h.RequestValidator(r); err != nil {
				h.handleError(w, r, err)
				return
			}
		}
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			h.handleError(w, r, err)
			return
		}
		if err := h.ParseAndDispatch(body); err != nil {
			h.handleError(w, r, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

func (h *WebhookHandler) handleError(w http.ResponseWriter, r *http.Request, err error) {
	if h.OnError != nil {
		h.OnError(w, r, err)
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func (h *WebhookHandler) handleUnhandledEvent(eventType chargebee.EventType, body []byte) error {
	if h.OnUnhandledEvent != nil {
		return h.OnUnhandledEvent(eventType, body)
	}
	return fmt.Errorf("unhandled event type: %s", eventType)
}
