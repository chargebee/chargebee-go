package webhook

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/chargebee/chargebee-go/v3/enum"
)

type WebhookHandler struct {
	RequestValidator func(*http.Request) error
	OnError          func(http.ResponseWriter, *http.Request, error)

	OnAddUsagesReminder func(AddUsagesReminderEvent) error

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

	OnOmnichannelOneTimeOrderCreated func(OmnichannelOneTimeOrderCreatedEvent) error

	OnOmnichannelOneTimeOrderItemCancelled func(OmnichannelOneTimeOrderItemCancelledEvent) error

	OnOmnichannelSubscriptionCreated func(OmnichannelSubscriptionCreatedEvent) error

	OnOmnichannelSubscriptionImported func(OmnichannelSubscriptionImportedEvent) error

	OnOmnichannelSubscriptionItemCancellationScheduled func(OmnichannelSubscriptionItemCancellationScheduledEvent) error

	OnOmnichannelSubscriptionItemCancelled func(OmnichannelSubscriptionItemCancelledEvent) error

	OnOmnichannelSubscriptionItemChangeScheduled func(OmnichannelSubscriptionItemChangeScheduledEvent) error

	OnOmnichannelSubscriptionItemChanged func(OmnichannelSubscriptionItemChangedEvent) error

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

	OnPriceVariantCreated func(PriceVariantCreatedEvent) error

	OnPriceVariantDeleted func(PriceVariantDeletedEvent) error

	OnPriceVariantUpdated func(PriceVariantUpdatedEvent) error

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

	OnVirtualBankAccountAdded func(VirtualBankAccountAddedEvent) error

	OnVirtualBankAccountDeleted func(VirtualBankAccountDeletedEvent) error

	OnVirtualBankAccountUpdated func(VirtualBankAccountUpdatedEvent) error

	OnVoucherCreateFailed func(VoucherCreateFailedEvent) error

	OnVoucherCreated func(VoucherCreatedEvent) error

	OnVoucherExpired func(VoucherExpiredEvent) error
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
		eventType, err := ParseEventType(body)
		if err != nil {
			h.handleError(w, r, err)
			return
		}
		switch eventType {

		case enum.EventTypeAddUsagesReminder:
			if h.OnAddUsagesReminder != nil {
				var e AddUsagesReminderEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnAddUsagesReminder(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeAttachedItemCreated:
			if h.OnAttachedItemCreated != nil {
				var e AttachedItemCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnAttachedItemCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeAttachedItemDeleted:
			if h.OnAttachedItemDeleted != nil {
				var e AttachedItemDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnAttachedItemDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeAttachedItemUpdated:
			if h.OnAttachedItemUpdated != nil {
				var e AttachedItemUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnAttachedItemUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeAuthorizationSucceeded:
			if h.OnAuthorizationSucceeded != nil {
				var e AuthorizationSucceededEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnAuthorizationSucceeded(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeAuthorizationVoided:
			if h.OnAuthorizationVoided != nil {
				var e AuthorizationVoidedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnAuthorizationVoided(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeBusinessEntityCreated:
			if h.OnBusinessEntityCreated != nil {
				var e BusinessEntityCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnBusinessEntityCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeBusinessEntityDeleted:
			if h.OnBusinessEntityDeleted != nil {
				var e BusinessEntityDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnBusinessEntityDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeBusinessEntityUpdated:
			if h.OnBusinessEntityUpdated != nil {
				var e BusinessEntityUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnBusinessEntityUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCardAdded:
			if h.OnCardAdded != nil {
				var e CardAddedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCardAdded(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCardDeleted:
			if h.OnCardDeleted != nil {
				var e CardDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCardDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCardExpired:
			if h.OnCardExpired != nil {
				var e CardExpiredEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCardExpired(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCardExpiryReminder:
			if h.OnCardExpiryReminder != nil {
				var e CardExpiryReminderEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCardExpiryReminder(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCardUpdated:
			if h.OnCardUpdated != nil {
				var e CardUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCardUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeContractTermCancelled:
			if h.OnContractTermCancelled != nil {
				var e ContractTermCancelledEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnContractTermCancelled(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeContractTermCompleted:
			if h.OnContractTermCompleted != nil {
				var e ContractTermCompletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnContractTermCompleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeContractTermCreated:
			if h.OnContractTermCreated != nil {
				var e ContractTermCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnContractTermCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeContractTermRenewed:
			if h.OnContractTermRenewed != nil {
				var e ContractTermRenewedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnContractTermRenewed(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeContractTermTerminated:
			if h.OnContractTermTerminated != nil {
				var e ContractTermTerminatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnContractTermTerminated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCouponCodesAdded:
			if h.OnCouponCodesAdded != nil {
				var e CouponCodesAddedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCouponCodesAdded(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCouponCodesDeleted:
			if h.OnCouponCodesDeleted != nil {
				var e CouponCodesDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCouponCodesDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCouponCodesUpdated:
			if h.OnCouponCodesUpdated != nil {
				var e CouponCodesUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCouponCodesUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCouponCreated:
			if h.OnCouponCreated != nil {
				var e CouponCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCouponCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCouponDeleted:
			if h.OnCouponDeleted != nil {
				var e CouponDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCouponDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCouponSetCreated:
			if h.OnCouponSetCreated != nil {
				var e CouponSetCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCouponSetCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCouponSetDeleted:
			if h.OnCouponSetDeleted != nil {
				var e CouponSetDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCouponSetDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCouponSetUpdated:
			if h.OnCouponSetUpdated != nil {
				var e CouponSetUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCouponSetUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCouponUpdated:
			if h.OnCouponUpdated != nil {
				var e CouponUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCouponUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCreditNoteCreated:
			if h.OnCreditNoteCreated != nil {
				var e CreditNoteCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCreditNoteCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCreditNoteCreatedWithBackdating:
			if h.OnCreditNoteCreatedWithBackdating != nil {
				var e CreditNoteCreatedWithBackdatingEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCreditNoteCreatedWithBackdating(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCreditNoteDeleted:
			if h.OnCreditNoteDeleted != nil {
				var e CreditNoteDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCreditNoteDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCreditNoteUpdated:
			if h.OnCreditNoteUpdated != nil {
				var e CreditNoteUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCreditNoteUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCustomerBusinessEntityChanged:
			if h.OnCustomerBusinessEntityChanged != nil {
				var e CustomerBusinessEntityChangedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCustomerBusinessEntityChanged(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCustomerChanged:
			if h.OnCustomerChanged != nil {
				var e CustomerChangedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCustomerChanged(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCustomerCreated:
			if h.OnCustomerCreated != nil {
				var e CustomerCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCustomerCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCustomerDeleted:
			if h.OnCustomerDeleted != nil {
				var e CustomerDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCustomerDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCustomerEntitlementsUpdated:
			if h.OnCustomerEntitlementsUpdated != nil {
				var e CustomerEntitlementsUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCustomerEntitlementsUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCustomerMovedIn:
			if h.OnCustomerMovedIn != nil {
				var e CustomerMovedInEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCustomerMovedIn(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeCustomerMovedOut:
			if h.OnCustomerMovedOut != nil {
				var e CustomerMovedOutEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnCustomerMovedOut(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeDifferentialPriceCreated:
			if h.OnDifferentialPriceCreated != nil {
				var e DifferentialPriceCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnDifferentialPriceCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeDifferentialPriceDeleted:
			if h.OnDifferentialPriceDeleted != nil {
				var e DifferentialPriceDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnDifferentialPriceDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeDifferentialPriceUpdated:
			if h.OnDifferentialPriceUpdated != nil {
				var e DifferentialPriceUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnDifferentialPriceUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeDunningUpdated:
			if h.OnDunningUpdated != nil {
				var e DunningUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnDunningUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeEntitlementOverridesAutoRemoved:
			if h.OnEntitlementOverridesAutoRemoved != nil {
				var e EntitlementOverridesAutoRemovedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnEntitlementOverridesAutoRemoved(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeEntitlementOverridesRemoved:
			if h.OnEntitlementOverridesRemoved != nil {
				var e EntitlementOverridesRemovedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnEntitlementOverridesRemoved(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeEntitlementOverridesUpdated:
			if h.OnEntitlementOverridesUpdated != nil {
				var e EntitlementOverridesUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnEntitlementOverridesUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeFeatureActivated:
			if h.OnFeatureActivated != nil {
				var e FeatureActivatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnFeatureActivated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeFeatureArchived:
			if h.OnFeatureArchived != nil {
				var e FeatureArchivedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnFeatureArchived(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeFeatureCreated:
			if h.OnFeatureCreated != nil {
				var e FeatureCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnFeatureCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeFeatureDeleted:
			if h.OnFeatureDeleted != nil {
				var e FeatureDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnFeatureDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeFeatureReactivated:
			if h.OnFeatureReactivated != nil {
				var e FeatureReactivatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnFeatureReactivated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeFeatureUpdated:
			if h.OnFeatureUpdated != nil {
				var e FeatureUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnFeatureUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeGiftCancelled:
			if h.OnGiftCancelled != nil {
				var e GiftCancelledEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnGiftCancelled(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeGiftClaimed:
			if h.OnGiftClaimed != nil {
				var e GiftClaimedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnGiftClaimed(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeGiftExpired:
			if h.OnGiftExpired != nil {
				var e GiftExpiredEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnGiftExpired(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeGiftScheduled:
			if h.OnGiftScheduled != nil {
				var e GiftScheduledEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnGiftScheduled(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeGiftUnclaimed:
			if h.OnGiftUnclaimed != nil {
				var e GiftUnclaimedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnGiftUnclaimed(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeGiftUpdated:
			if h.OnGiftUpdated != nil {
				var e GiftUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnGiftUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeHierarchyCreated:
			if h.OnHierarchyCreated != nil {
				var e HierarchyCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnHierarchyCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeHierarchyDeleted:
			if h.OnHierarchyDeleted != nil {
				var e HierarchyDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnHierarchyDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeInvoiceDeleted:
			if h.OnInvoiceDeleted != nil {
				var e InvoiceDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnInvoiceDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeInvoiceGenerated:
			if h.OnInvoiceGenerated != nil {
				var e InvoiceGeneratedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnInvoiceGenerated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeInvoiceGeneratedWithBackdating:
			if h.OnInvoiceGeneratedWithBackdating != nil {
				var e InvoiceGeneratedWithBackdatingEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnInvoiceGeneratedWithBackdating(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeInvoiceUpdated:
			if h.OnInvoiceUpdated != nil {
				var e InvoiceUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnInvoiceUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeItemCreated:
			if h.OnItemCreated != nil {
				var e ItemCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnItemCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeItemDeleted:
			if h.OnItemDeleted != nil {
				var e ItemDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnItemDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeItemEntitlementsRemoved:
			if h.OnItemEntitlementsRemoved != nil {
				var e ItemEntitlementsRemovedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnItemEntitlementsRemoved(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeItemEntitlementsUpdated:
			if h.OnItemEntitlementsUpdated != nil {
				var e ItemEntitlementsUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnItemEntitlementsUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeItemFamilyCreated:
			if h.OnItemFamilyCreated != nil {
				var e ItemFamilyCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnItemFamilyCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeItemFamilyDeleted:
			if h.OnItemFamilyDeleted != nil {
				var e ItemFamilyDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnItemFamilyDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeItemFamilyUpdated:
			if h.OnItemFamilyUpdated != nil {
				var e ItemFamilyUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnItemFamilyUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeItemPriceCreated:
			if h.OnItemPriceCreated != nil {
				var e ItemPriceCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnItemPriceCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeItemPriceDeleted:
			if h.OnItemPriceDeleted != nil {
				var e ItemPriceDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnItemPriceDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeItemPriceEntitlementsRemoved:
			if h.OnItemPriceEntitlementsRemoved != nil {
				var e ItemPriceEntitlementsRemovedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnItemPriceEntitlementsRemoved(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeItemPriceEntitlementsUpdated:
			if h.OnItemPriceEntitlementsUpdated != nil {
				var e ItemPriceEntitlementsUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnItemPriceEntitlementsUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeItemPriceUpdated:
			if h.OnItemPriceUpdated != nil {
				var e ItemPriceUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnItemPriceUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeItemUpdated:
			if h.OnItemUpdated != nil {
				var e ItemUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnItemUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeMrrUpdated:
			if h.OnMrrUpdated != nil {
				var e MrrUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnMrrUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelOneTimeOrderCreated:
			if h.OnOmnichannelOneTimeOrderCreated != nil {
				var e OmnichannelOneTimeOrderCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelOneTimeOrderCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelOneTimeOrderItemCancelled:
			if h.OnOmnichannelOneTimeOrderItemCancelled != nil {
				var e OmnichannelOneTimeOrderItemCancelledEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelOneTimeOrderItemCancelled(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionCreated:
			if h.OnOmnichannelSubscriptionCreated != nil {
				var e OmnichannelSubscriptionCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionImported:
			if h.OnOmnichannelSubscriptionImported != nil {
				var e OmnichannelSubscriptionImportedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionImported(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionItemCancellationScheduled:
			if h.OnOmnichannelSubscriptionItemCancellationScheduled != nil {
				var e OmnichannelSubscriptionItemCancellationScheduledEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionItemCancellationScheduled(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionItemCancelled:
			if h.OnOmnichannelSubscriptionItemCancelled != nil {
				var e OmnichannelSubscriptionItemCancelledEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionItemCancelled(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionItemChangeScheduled:
			if h.OnOmnichannelSubscriptionItemChangeScheduled != nil {
				var e OmnichannelSubscriptionItemChangeScheduledEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionItemChangeScheduled(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionItemChanged:
			if h.OnOmnichannelSubscriptionItemChanged != nil {
				var e OmnichannelSubscriptionItemChangedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionItemChanged(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionItemDowngraded:
			if h.OnOmnichannelSubscriptionItemDowngraded != nil {
				var e OmnichannelSubscriptionItemDowngradedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionItemDowngraded(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionItemDunningExpired:
			if h.OnOmnichannelSubscriptionItemDunningExpired != nil {
				var e OmnichannelSubscriptionItemDunningExpiredEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionItemDunningExpired(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionItemDunningStarted:
			if h.OnOmnichannelSubscriptionItemDunningStarted != nil {
				var e OmnichannelSubscriptionItemDunningStartedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionItemDunningStarted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionItemExpired:
			if h.OnOmnichannelSubscriptionItemExpired != nil {
				var e OmnichannelSubscriptionItemExpiredEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionItemExpired(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionItemGracePeriodExpired:
			if h.OnOmnichannelSubscriptionItemGracePeriodExpired != nil {
				var e OmnichannelSubscriptionItemGracePeriodExpiredEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionItemGracePeriodExpired(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionItemGracePeriodStarted:
			if h.OnOmnichannelSubscriptionItemGracePeriodStarted != nil {
				var e OmnichannelSubscriptionItemGracePeriodStartedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionItemGracePeriodStarted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionItemPauseScheduled:
			if h.OnOmnichannelSubscriptionItemPauseScheduled != nil {
				var e OmnichannelSubscriptionItemPauseScheduledEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionItemPauseScheduled(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionItemPaused:
			if h.OnOmnichannelSubscriptionItemPaused != nil {
				var e OmnichannelSubscriptionItemPausedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionItemPaused(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionItemReactivated:
			if h.OnOmnichannelSubscriptionItemReactivated != nil {
				var e OmnichannelSubscriptionItemReactivatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionItemReactivated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionItemRenewed:
			if h.OnOmnichannelSubscriptionItemRenewed != nil {
				var e OmnichannelSubscriptionItemRenewedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionItemRenewed(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionItemResubscribed:
			if h.OnOmnichannelSubscriptionItemResubscribed != nil {
				var e OmnichannelSubscriptionItemResubscribedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionItemResubscribed(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionItemResumed:
			if h.OnOmnichannelSubscriptionItemResumed != nil {
				var e OmnichannelSubscriptionItemResumedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionItemResumed(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionItemScheduledCancellationRemoved:
			if h.OnOmnichannelSubscriptionItemScheduledCancellationRemoved != nil {
				var e OmnichannelSubscriptionItemScheduledCancellationRemovedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionItemScheduledCancellationRemoved(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionItemScheduledChangeRemoved:
			if h.OnOmnichannelSubscriptionItemScheduledChangeRemoved != nil {
				var e OmnichannelSubscriptionItemScheduledChangeRemovedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionItemScheduledChangeRemoved(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionItemUpgraded:
			if h.OnOmnichannelSubscriptionItemUpgraded != nil {
				var e OmnichannelSubscriptionItemUpgradedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionItemUpgraded(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelSubscriptionMovedIn:
			if h.OnOmnichannelSubscriptionMovedIn != nil {
				var e OmnichannelSubscriptionMovedInEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelSubscriptionMovedIn(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOmnichannelTransactionCreated:
			if h.OnOmnichannelTransactionCreated != nil {
				var e OmnichannelTransactionCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOmnichannelTransactionCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOrderCancelled:
			if h.OnOrderCancelled != nil {
				var e OrderCancelledEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOrderCancelled(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOrderCreated:
			if h.OnOrderCreated != nil {
				var e OrderCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOrderCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOrderDeleted:
			if h.OnOrderDeleted != nil {
				var e OrderDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOrderDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOrderDelivered:
			if h.OnOrderDelivered != nil {
				var e OrderDeliveredEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOrderDelivered(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOrderReadyToProcess:
			if h.OnOrderReadyToProcess != nil {
				var e OrderReadyToProcessEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOrderReadyToProcess(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOrderReadyToShip:
			if h.OnOrderReadyToShip != nil {
				var e OrderReadyToShipEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOrderReadyToShip(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOrderResent:
			if h.OnOrderResent != nil {
				var e OrderResentEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOrderResent(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOrderReturned:
			if h.OnOrderReturned != nil {
				var e OrderReturnedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOrderReturned(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeOrderUpdated:
			if h.OnOrderUpdated != nil {
				var e OrderUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnOrderUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePaymentFailed:
			if h.OnPaymentFailed != nil {
				var e PaymentFailedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPaymentFailed(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePaymentInitiated:
			if h.OnPaymentInitiated != nil {
				var e PaymentInitiatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPaymentInitiated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePaymentIntentCreated:
			if h.OnPaymentIntentCreated != nil {
				var e PaymentIntentCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPaymentIntentCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePaymentIntentUpdated:
			if h.OnPaymentIntentUpdated != nil {
				var e PaymentIntentUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPaymentIntentUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePaymentRefunded:
			if h.OnPaymentRefunded != nil {
				var e PaymentRefundedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPaymentRefunded(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePaymentScheduleSchemeCreated:
			if h.OnPaymentScheduleSchemeCreated != nil {
				var e PaymentScheduleSchemeCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPaymentScheduleSchemeCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePaymentScheduleSchemeDeleted:
			if h.OnPaymentScheduleSchemeDeleted != nil {
				var e PaymentScheduleSchemeDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPaymentScheduleSchemeDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePaymentSchedulesCreated:
			if h.OnPaymentSchedulesCreated != nil {
				var e PaymentSchedulesCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPaymentSchedulesCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePaymentSchedulesUpdated:
			if h.OnPaymentSchedulesUpdated != nil {
				var e PaymentSchedulesUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPaymentSchedulesUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePaymentSourceAdded:
			if h.OnPaymentSourceAdded != nil {
				var e PaymentSourceAddedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPaymentSourceAdded(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePaymentSourceDeleted:
			if h.OnPaymentSourceDeleted != nil {
				var e PaymentSourceDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPaymentSourceDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePaymentSourceExpired:
			if h.OnPaymentSourceExpired != nil {
				var e PaymentSourceExpiredEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPaymentSourceExpired(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePaymentSourceExpiring:
			if h.OnPaymentSourceExpiring != nil {
				var e PaymentSourceExpiringEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPaymentSourceExpiring(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePaymentSourceLocallyDeleted:
			if h.OnPaymentSourceLocallyDeleted != nil {
				var e PaymentSourceLocallyDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPaymentSourceLocallyDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePaymentSourceUpdated:
			if h.OnPaymentSourceUpdated != nil {
				var e PaymentSourceUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPaymentSourceUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePaymentSucceeded:
			if h.OnPaymentSucceeded != nil {
				var e PaymentSucceededEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPaymentSucceeded(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePendingInvoiceCreated:
			if h.OnPendingInvoiceCreated != nil {
				var e PendingInvoiceCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPendingInvoiceCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePendingInvoiceUpdated:
			if h.OnPendingInvoiceUpdated != nil {
				var e PendingInvoiceUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPendingInvoiceUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePriceVariantCreated:
			if h.OnPriceVariantCreated != nil {
				var e PriceVariantCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPriceVariantCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePriceVariantDeleted:
			if h.OnPriceVariantDeleted != nil {
				var e PriceVariantDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPriceVariantDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePriceVariantUpdated:
			if h.OnPriceVariantUpdated != nil {
				var e PriceVariantUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPriceVariantUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePromotionalCreditsAdded:
			if h.OnPromotionalCreditsAdded != nil {
				var e PromotionalCreditsAddedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPromotionalCreditsAdded(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePromotionalCreditsDeducted:
			if h.OnPromotionalCreditsDeducted != nil {
				var e PromotionalCreditsDeductedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPromotionalCreditsDeducted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypePurchaseCreated:
			if h.OnPurchaseCreated != nil {
				var e PurchaseCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnPurchaseCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeQuoteCreated:
			if h.OnQuoteCreated != nil {
				var e QuoteCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnQuoteCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeQuoteDeleted:
			if h.OnQuoteDeleted != nil {
				var e QuoteDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnQuoteDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeQuoteUpdated:
			if h.OnQuoteUpdated != nil {
				var e QuoteUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnQuoteUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeRecordPurchaseFailed:
			if h.OnRecordPurchaseFailed != nil {
				var e RecordPurchaseFailedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnRecordPurchaseFailed(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeRefundInitiated:
			if h.OnRefundInitiated != nil {
				var e RefundInitiatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnRefundInitiated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeRuleCreated:
			if h.OnRuleCreated != nil {
				var e RuleCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnRuleCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeRuleDeleted:
			if h.OnRuleDeleted != nil {
				var e RuleDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnRuleDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeRuleUpdated:
			if h.OnRuleUpdated != nil {
				var e RuleUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnRuleUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSalesOrderCreated:
			if h.OnSalesOrderCreated != nil {
				var e SalesOrderCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSalesOrderCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSalesOrderUpdated:
			if h.OnSalesOrderUpdated != nil {
				var e SalesOrderUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSalesOrderUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionActivated:
			if h.OnSubscriptionActivated != nil {
				var e SubscriptionActivatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionActivated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionActivatedWithBackdating:
			if h.OnSubscriptionActivatedWithBackdating != nil {
				var e SubscriptionActivatedWithBackdatingEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionActivatedWithBackdating(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionAdvanceInvoiceScheduleAdded:
			if h.OnSubscriptionAdvanceInvoiceScheduleAdded != nil {
				var e SubscriptionAdvanceInvoiceScheduleAddedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionAdvanceInvoiceScheduleAdded(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionAdvanceInvoiceScheduleRemoved:
			if h.OnSubscriptionAdvanceInvoiceScheduleRemoved != nil {
				var e SubscriptionAdvanceInvoiceScheduleRemovedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionAdvanceInvoiceScheduleRemoved(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionAdvanceInvoiceScheduleUpdated:
			if h.OnSubscriptionAdvanceInvoiceScheduleUpdated != nil {
				var e SubscriptionAdvanceInvoiceScheduleUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionAdvanceInvoiceScheduleUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionBusinessEntityChanged:
			if h.OnSubscriptionBusinessEntityChanged != nil {
				var e SubscriptionBusinessEntityChangedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionBusinessEntityChanged(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionCanceledWithBackdating:
			if h.OnSubscriptionCanceledWithBackdating != nil {
				var e SubscriptionCanceledWithBackdatingEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionCanceledWithBackdating(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionCancellationReminder:
			if h.OnSubscriptionCancellationReminder != nil {
				var e SubscriptionCancellationReminderEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionCancellationReminder(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionCancellationScheduled:
			if h.OnSubscriptionCancellationScheduled != nil {
				var e SubscriptionCancellationScheduledEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionCancellationScheduled(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionCancelled:
			if h.OnSubscriptionCancelled != nil {
				var e SubscriptionCancelledEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionCancelled(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionChanged:
			if h.OnSubscriptionChanged != nil {
				var e SubscriptionChangedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionChanged(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionChangedWithBackdating:
			if h.OnSubscriptionChangedWithBackdating != nil {
				var e SubscriptionChangedWithBackdatingEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionChangedWithBackdating(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionChangesScheduled:
			if h.OnSubscriptionChangesScheduled != nil {
				var e SubscriptionChangesScheduledEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionChangesScheduled(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionCreated:
			if h.OnSubscriptionCreated != nil {
				var e SubscriptionCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionCreatedWithBackdating:
			if h.OnSubscriptionCreatedWithBackdating != nil {
				var e SubscriptionCreatedWithBackdatingEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionCreatedWithBackdating(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionDeleted:
			if h.OnSubscriptionDeleted != nil {
				var e SubscriptionDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionEntitlementsCreated:
			if h.OnSubscriptionEntitlementsCreated != nil {
				var e SubscriptionEntitlementsCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionEntitlementsCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionEntitlementsUpdated:
			if h.OnSubscriptionEntitlementsUpdated != nil {
				var e SubscriptionEntitlementsUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionEntitlementsUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionItemsRenewed:
			if h.OnSubscriptionItemsRenewed != nil {
				var e SubscriptionItemsRenewedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionItemsRenewed(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionMovedIn:
			if h.OnSubscriptionMovedIn != nil {
				var e SubscriptionMovedInEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionMovedIn(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionMovedOut:
			if h.OnSubscriptionMovedOut != nil {
				var e SubscriptionMovedOutEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionMovedOut(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionMovementFailed:
			if h.OnSubscriptionMovementFailed != nil {
				var e SubscriptionMovementFailedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionMovementFailed(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionPauseScheduled:
			if h.OnSubscriptionPauseScheduled != nil {
				var e SubscriptionPauseScheduledEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionPauseScheduled(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionPaused:
			if h.OnSubscriptionPaused != nil {
				var e SubscriptionPausedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionPaused(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionRampApplied:
			if h.OnSubscriptionRampApplied != nil {
				var e SubscriptionRampAppliedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionRampApplied(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionRampCreated:
			if h.OnSubscriptionRampCreated != nil {
				var e SubscriptionRampCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionRampCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionRampDeleted:
			if h.OnSubscriptionRampDeleted != nil {
				var e SubscriptionRampDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionRampDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionRampDrafted:
			if h.OnSubscriptionRampDrafted != nil {
				var e SubscriptionRampDraftedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionRampDrafted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionRampUpdated:
			if h.OnSubscriptionRampUpdated != nil {
				var e SubscriptionRampUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionRampUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionReactivated:
			if h.OnSubscriptionReactivated != nil {
				var e SubscriptionReactivatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionReactivated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionReactivatedWithBackdating:
			if h.OnSubscriptionReactivatedWithBackdating != nil {
				var e SubscriptionReactivatedWithBackdatingEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionReactivatedWithBackdating(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionRenewalReminder:
			if h.OnSubscriptionRenewalReminder != nil {
				var e SubscriptionRenewalReminderEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionRenewalReminder(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionRenewed:
			if h.OnSubscriptionRenewed != nil {
				var e SubscriptionRenewedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionRenewed(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionResumed:
			if h.OnSubscriptionResumed != nil {
				var e SubscriptionResumedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionResumed(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionResumptionScheduled:
			if h.OnSubscriptionResumptionScheduled != nil {
				var e SubscriptionResumptionScheduledEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionResumptionScheduled(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionScheduledCancellationRemoved:
			if h.OnSubscriptionScheduledCancellationRemoved != nil {
				var e SubscriptionScheduledCancellationRemovedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionScheduledCancellationRemoved(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionScheduledChangesRemoved:
			if h.OnSubscriptionScheduledChangesRemoved != nil {
				var e SubscriptionScheduledChangesRemovedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionScheduledChangesRemoved(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionScheduledPauseRemoved:
			if h.OnSubscriptionScheduledPauseRemoved != nil {
				var e SubscriptionScheduledPauseRemovedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionScheduledPauseRemoved(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionScheduledResumptionRemoved:
			if h.OnSubscriptionScheduledResumptionRemoved != nil {
				var e SubscriptionScheduledResumptionRemovedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionScheduledResumptionRemoved(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionShippingAddressUpdated:
			if h.OnSubscriptionShippingAddressUpdated != nil {
				var e SubscriptionShippingAddressUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionShippingAddressUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionStarted:
			if h.OnSubscriptionStarted != nil {
				var e SubscriptionStartedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionStarted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionTrialEndReminder:
			if h.OnSubscriptionTrialEndReminder != nil {
				var e SubscriptionTrialEndReminderEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionTrialEndReminder(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeSubscriptionTrialExtended:
			if h.OnSubscriptionTrialExtended != nil {
				var e SubscriptionTrialExtendedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnSubscriptionTrialExtended(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeTaxWithheldDeleted:
			if h.OnTaxWithheldDeleted != nil {
				var e TaxWithheldDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnTaxWithheldDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeTaxWithheldRecorded:
			if h.OnTaxWithheldRecorded != nil {
				var e TaxWithheldRecordedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnTaxWithheldRecorded(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeTaxWithheldRefunded:
			if h.OnTaxWithheldRefunded != nil {
				var e TaxWithheldRefundedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnTaxWithheldRefunded(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeTokenConsumed:
			if h.OnTokenConsumed != nil {
				var e TokenConsumedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnTokenConsumed(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeTokenCreated:
			if h.OnTokenCreated != nil {
				var e TokenCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnTokenCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeTokenExpired:
			if h.OnTokenExpired != nil {
				var e TokenExpiredEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnTokenExpired(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeTransactionCreated:
			if h.OnTransactionCreated != nil {
				var e TransactionCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnTransactionCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeTransactionDeleted:
			if h.OnTransactionDeleted != nil {
				var e TransactionDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnTransactionDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeTransactionUpdated:
			if h.OnTransactionUpdated != nil {
				var e TransactionUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnTransactionUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeUnbilledChargesCreated:
			if h.OnUnbilledChargesCreated != nil {
				var e UnbilledChargesCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnUnbilledChargesCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeUnbilledChargesDeleted:
			if h.OnUnbilledChargesDeleted != nil {
				var e UnbilledChargesDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnUnbilledChargesDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeUnbilledChargesInvoiced:
			if h.OnUnbilledChargesInvoiced != nil {
				var e UnbilledChargesInvoicedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnUnbilledChargesInvoiced(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeUnbilledChargesVoided:
			if h.OnUnbilledChargesVoided != nil {
				var e UnbilledChargesVoidedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnUnbilledChargesVoided(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeUsageFileIngested:
			if h.OnUsageFileIngested != nil {
				var e UsageFileIngestedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnUsageFileIngested(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeVirtualBankAccountAdded:
			if h.OnVirtualBankAccountAdded != nil {
				var e VirtualBankAccountAddedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnVirtualBankAccountAdded(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeVirtualBankAccountDeleted:
			if h.OnVirtualBankAccountDeleted != nil {
				var e VirtualBankAccountDeletedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnVirtualBankAccountDeleted(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeVirtualBankAccountUpdated:
			if h.OnVirtualBankAccountUpdated != nil {
				var e VirtualBankAccountUpdatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnVirtualBankAccountUpdated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeVoucherCreateFailed:
			if h.OnVoucherCreateFailed != nil {
				var e VoucherCreateFailedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnVoucherCreateFailed(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeVoucherCreated:
			if h.OnVoucherCreated != nil {
				var e VoucherCreatedEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnVoucherCreated(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		case enum.EventTypeVoucherExpired:
			if h.OnVoucherExpired != nil {
				var e VoucherExpiredEvent
				if err := json.Unmarshal(body, &e); err != nil {
					h.handleError(w, r, err)
					return
				}
				if err := h.OnVoucherExpired(e); err != nil {
					h.handleError(w, r, err)
					return
				}
			}

		default:
			// Unknown or unhandled event type; treat as success
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
