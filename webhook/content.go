package webhook

import (
	"time"

	"github.com/chargebee/chargebee-go/v4"
)

type BaseEvent struct {
	Id            string `json:"id"`
	OccurredAt    int64  `json:"occurred_at"`
	Source        string `json:"source"`
	Object        string `json:"object"`
	ApiVersion    string `json:"api_version"`
	EventType     string `json:"event_type"`
	WebhookStatus string `json:"webhook_status"`
}

func (e *BaseEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// Event content structures for each webhook type

type SubscriptionPauseScheduledContent struct {
	Subscription           *chargebee.Subscription           `json:"subscription,omitempty"`
	Customer               *chargebee.Customer               `json:"customer,omitempty"`
	Card                   *chargebee.Card                   `json:"card,omitempty"`
	AdvanceInvoiceSchedule *chargebee.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type CustomerBusinessEntityChangedContent struct {
	BusinessEntityTransfer *chargebee.BusinessEntityTransfer `json:"business_entity_transfer,omitempty"`
	Customer               *chargebee.Customer               `json:"customer,omitempty"`
}

type SubscriptionAdvanceInvoiceScheduleAddedContent struct {
	Subscription           *chargebee.Subscription           `json:"subscription,omitempty"`
	Customer               *chargebee.Customer               `json:"customer,omitempty"`
	AdvanceInvoiceSchedule *chargebee.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type GiftExpiredContent struct {
	Gift *chargebee.Gift `json:"gift,omitempty"`
}

type TaxWithheldDeletedContent struct {
	TaxWithheld *chargebee.TaxWithheld `json:"tax_withheld,omitempty"`
	Invoice     *chargebee.Invoice     `json:"invoice,omitempty"`
	CreditNote  *chargebee.CreditNote  `json:"credit_note,omitempty"`
}

type UnbilledChargesDeletedContent struct {
	UnbilledCharge *chargebee.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type CouponUpdatedContent struct {
	Coupon *chargebee.Coupon `json:"coupon,omitempty"`
}

type ProductUpdatedContent struct {
}

type OmnichannelSubscriptionItemReactivatedContent struct {
	OmnichannelSubscription     *chargebee.OmnichannelSubscription     `json:"omnichannel_subscription,omitempty"`
	OmnichannelSubscriptionItem *chargebee.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`
	Customer                    *chargebee.Customer                    `json:"customer,omitempty"`
}

type OmnichannelSubscriptionItemRenewedContent struct {
	OmnichannelSubscriptionItem                *chargebee.OmnichannelSubscriptionItem                `json:"omnichannel_subscription_item,omitempty"`
	OmnichannelSubscription                    *chargebee.OmnichannelSubscription                    `json:"omnichannel_subscription,omitempty"`
	OmnichannelTransaction                     *chargebee.OmnichannelTransaction                     `json:"omnichannel_transaction,omitempty"`
	OmnichannelSubscriptionItemScheduledChange *chargebee.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`
	Customer                                   *chargebee.Customer                                   `json:"customer,omitempty"`
}

type UnbilledChargesCreatedContent struct {
	UnbilledCharge *chargebee.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type SubscriptionResumedContent struct {
	Subscription   *chargebee.Subscription   `json:"subscription,omitempty"`
	Customer       *chargebee.Customer       `json:"customer,omitempty"`
	Card           *chargebee.Card           `json:"card,omitempty"`
	Invoice        *chargebee.Invoice        `json:"invoice,omitempty"`
	UnbilledCharge *chargebee.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type OmnichannelOneTimeOrderItemCancelledContent struct {
	OmnichannelOneTimeOrder     *chargebee.OmnichannelOneTimeOrder     `json:"omnichannel_one_time_order,omitempty"`
	OmnichannelOneTimeOrderItem *chargebee.OmnichannelOneTimeOrderItem `json:"omnichannel_one_time_order_item,omitempty"`
	OmnichannelTransaction      *chargebee.OmnichannelTransaction      `json:"omnichannel_transaction,omitempty"`
	Customer                    *chargebee.Customer                    `json:"customer,omitempty"`
}

type SubscriptionCancelledContent struct {
	Subscription   *chargebee.Subscription   `json:"subscription,omitempty"`
	Customer       *chargebee.Customer       `json:"customer,omitempty"`
	Card           *chargebee.Card           `json:"card,omitempty"`
	Invoice        *chargebee.Invoice        `json:"invoice,omitempty"`
	CreditNote     *chargebee.CreditNote     `json:"credit_note,omitempty"`
	UnbilledCharge *chargebee.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type ItemEntitlementsRemovedContent struct {
	Feature              *chargebee.Feature              `json:"feature,omitempty"`
	Metadata             *chargebee.Metadata             `json:"metadata,omitempty"`
	ImpactedItem         *chargebee.ImpactedItem         `json:"impacted_item,omitempty"`
	ImpactedSubscription *chargebee.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type BusinessEntityCreatedContent struct {
	BusinessEntity *chargebee.BusinessEntity `json:"business_entity,omitempty"`
}

type CouponSetUpdatedContent struct {
	Coupon    *chargebee.Coupon    `json:"coupon,omitempty"`
	CouponSet *chargebee.CouponSet `json:"coupon_set,omitempty"`
}

type DifferentialPriceUpdatedContent struct {
	DifferentialPrice *chargebee.DifferentialPrice `json:"differential_price,omitempty"`
}

type OmnichannelSubscriptionItemPausedContent struct {
	OmnichannelSubscription     *chargebee.OmnichannelSubscription     `json:"omnichannel_subscription,omitempty"`
	OmnichannelSubscriptionItem *chargebee.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`
	Customer                    *chargebee.Customer                    `json:"customer,omitempty"`
}

type EntitlementOverridesRemovedContent struct {
	ImpactedSubscription *chargebee.ImpactedSubscription `json:"impacted_subscription,omitempty"`
	Metadata             *chargebee.Metadata             `json:"metadata,omitempty"`
}

type SubscriptionActivatedWithBackdatingContent struct {
	Subscription   *chargebee.Subscription   `json:"subscription,omitempty"`
	Customer       *chargebee.Customer       `json:"customer,omitempty"`
	Card           *chargebee.Card           `json:"card,omitempty"`
	Invoice        *chargebee.Invoice        `json:"invoice,omitempty"`
	UnbilledCharge *chargebee.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type SubscriptionTrialEndReminderContent struct {
	Subscription           *chargebee.Subscription           `json:"subscription,omitempty"`
	Customer               *chargebee.Customer               `json:"customer,omitempty"`
	Card                   *chargebee.Card                   `json:"card,omitempty"`
	AdvanceInvoiceSchedule *chargebee.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type SubscriptionShippingAddressUpdatedContent struct {
	Subscription           *chargebee.Subscription           `json:"subscription,omitempty"`
	Customer               *chargebee.Customer               `json:"customer,omitempty"`
	Card                   *chargebee.Card                   `json:"card,omitempty"`
	AdvanceInvoiceSchedule *chargebee.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type VoucherCreateFailedContent struct {
	PaymentVoucher *chargebee.PaymentVoucher `json:"payment_voucher,omitempty"`
}

type GiftClaimedContent struct {
	Gift *chargebee.Gift `json:"gift,omitempty"`
}

type CustomerDeletedContent struct {
	Customer     *chargebee.Customer     `json:"customer,omitempty"`
	Card         *chargebee.Card         `json:"card,omitempty"`
	Subscription *chargebee.Subscription `json:"subscription,omitempty"`
}

type RefundInitiatedContent struct {
	Transaction  *chargebee.Transaction  `json:"transaction,omitempty"`
	Invoice      *chargebee.Invoice      `json:"invoice,omitempty"`
	CreditNote   *chargebee.CreditNote   `json:"credit_note,omitempty"`
	Customer     *chargebee.Customer     `json:"customer,omitempty"`
	Subscription *chargebee.Subscription `json:"subscription,omitempty"`
	Card         *chargebee.Card         `json:"card,omitempty"`
}

type InvoiceGeneratedWithBackdatingContent struct {
	Invoice *chargebee.Invoice `json:"invoice,omitempty"`
}

type OmnichannelTransactionCreatedContent struct {
	OmnichannelTransaction *chargebee.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`
}

type AddUsagesReminderContent struct {
	Subscription *chargebee.Subscription `json:"subscription,omitempty"`
	Customer     *chargebee.Customer     `json:"customer,omitempty"`
}

type VoucherCreatedContent struct {
	PaymentVoucher *chargebee.PaymentVoucher `json:"payment_voucher,omitempty"`
}

type RuleUpdatedContent struct {
	Rule *chargebee.Rule `json:"rule,omitempty"`
}

type PaymentSchedulesCreatedContent struct {
	PaymentSchedule *chargebee.PaymentSchedule `json:"payment_schedule,omitempty"`
}

type FeatureActivatedContent struct {
	Feature              *chargebee.Feature              `json:"feature,omitempty"`
	Metadata             *chargebee.Metadata             `json:"metadata,omitempty"`
	ImpactedItem         *chargebee.ImpactedItem         `json:"impacted_item,omitempty"`
	ImpactedSubscription *chargebee.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type PaymentSourceLocallyDeletedContent struct {
	Customer      *chargebee.Customer      `json:"customer,omitempty"`
	PaymentSource *chargebee.PaymentSource `json:"payment_source,omitempty"`
}

type InvoiceGeneratedContent struct {
	Invoice *chargebee.Invoice `json:"invoice,omitempty"`
}

type VoucherExpiredContent struct {
	PaymentVoucher *chargebee.PaymentVoucher `json:"payment_voucher,omitempty"`
}

type AuthorizationSucceededContent struct {
	Transaction *chargebee.Transaction `json:"transaction,omitempty"`
}

type GiftScheduledContent struct {
	Gift *chargebee.Gift `json:"gift,omitempty"`
}

type SubscriptionChangesScheduledContent struct {
	Subscription           *chargebee.Subscription           `json:"subscription,omitempty"`
	Customer               *chargebee.Customer               `json:"customer,omitempty"`
	Card                   *chargebee.Card                   `json:"card,omitempty"`
	AdvanceInvoiceSchedule *chargebee.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type SubscriptionChangedWithBackdatingContent struct {
	Subscription   *chargebee.Subscription   `json:"subscription,omitempty"`
	Customer       *chargebee.Customer       `json:"customer,omitempty"`
	Card           *chargebee.Card           `json:"card,omitempty"`
	Invoice        *chargebee.Invoice        `json:"invoice,omitempty"`
	CreditNote     *chargebee.CreditNote     `json:"credit_note,omitempty"`
	UnbilledCharge *chargebee.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type VariantCreatedContent struct {
}

type OmnichannelSubscriptionItemChangedContent struct {
	OmnichannelSubscriptionItem                *chargebee.OmnichannelSubscriptionItem                `json:"omnichannel_subscription_item,omitempty"`
	OmnichannelSubscription                    *chargebee.OmnichannelSubscription                    `json:"omnichannel_subscription,omitempty"`
	OmnichannelTransaction                     *chargebee.OmnichannelTransaction                     `json:"omnichannel_transaction,omitempty"`
	OmnichannelSubscriptionItemScheduledChange *chargebee.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`
	Customer                                   *chargebee.Customer                                   `json:"customer,omitempty"`
}

type GiftUnclaimedContent struct {
	Gift *chargebee.Gift `json:"gift,omitempty"`
}

type VirtualBankAccountAddedContent struct {
	Customer           *chargebee.Customer           `json:"customer,omitempty"`
	VirtualBankAccount *chargebee.VirtualBankAccount `json:"virtual_bank_account,omitempty"`
}

type PaymentIntentCreatedContent struct {
	PaymentIntent *chargebee.PaymentIntent `json:"payment_intent,omitempty"`
}

type CreditNoteCreatedWithBackdatingContent struct {
	CreditNote *chargebee.CreditNote `json:"credit_note,omitempty"`
}

type ContractTermTerminatedContent struct {
	ContractTerm *chargebee.ContractTerm `json:"contract_term,omitempty"`
}

type ItemFamilyUpdatedContent struct {
	ItemFamily *chargebee.ItemFamily `json:"item_family,omitempty"`
}

type OrderCreatedContent struct {
	Order *chargebee.Order `json:"order,omitempty"`
}

type PriceVariantDeletedContent struct {
	PriceVariant *chargebee.PriceVariant `json:"price_variant,omitempty"`
	Attribute    *chargebee.Attribute    `json:"attribute,omitempty"`
}

type SubscriptionMovementFailedContent struct {
	Subscription *chargebee.Subscription `json:"subscription,omitempty"`
}

type CustomerMovedInContent struct {
	Customer *chargebee.Customer `json:"customer,omitempty"`
	Card     *chargebee.Card     `json:"card,omitempty"`
}

type SubscriptionAdvanceInvoiceScheduleUpdatedContent struct {
	Subscription           *chargebee.Subscription           `json:"subscription,omitempty"`
	Customer               *chargebee.Customer               `json:"customer,omitempty"`
	AdvanceInvoiceSchedule *chargebee.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type ItemDeletedContent struct {
	Item *chargebee.Item `json:"item,omitempty"`
}

type SubscriptionRampDraftedContent struct {
	Ramp *chargebee.Ramp `json:"ramp,omitempty"`
}

type DunningUpdatedContent struct {
	Invoice *chargebee.Invoice `json:"invoice,omitempty"`
}

type ItemEntitlementsUpdatedContent struct {
	Feature              *chargebee.Feature              `json:"feature,omitempty"`
	Metadata             *chargebee.Metadata             `json:"metadata,omitempty"`
	ImpactedItem         *chargebee.ImpactedItem         `json:"impacted_item,omitempty"`
	ImpactedSubscription *chargebee.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type TokenConsumedContent struct {
	Token *chargebee.Token `json:"token,omitempty"`
}

type HierarchyDeletedContent struct {
	Customer *chargebee.Customer `json:"customer,omitempty"`
}

type SubscriptionCancellationScheduledContent struct {
	Subscription           *chargebee.Subscription           `json:"subscription,omitempty"`
	Customer               *chargebee.Customer               `json:"customer,omitempty"`
	Card                   *chargebee.Card                   `json:"card,omitempty"`
	AdvanceInvoiceSchedule *chargebee.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type SubscriptionRenewedContent struct {
	Subscription   *chargebee.Subscription   `json:"subscription,omitempty"`
	Customer       *chargebee.Customer       `json:"customer,omitempty"`
	Card           *chargebee.Card           `json:"card,omitempty"`
	Invoice        *chargebee.Invoice        `json:"invoice,omitempty"`
	UnbilledCharge *chargebee.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type FeatureUpdatedContent struct {
	Feature  *chargebee.Feature  `json:"feature,omitempty"`
	Metadata *chargebee.Metadata `json:"metadata,omitempty"`
}

type FeatureDeletedContent struct {
	Feature              *chargebee.Feature              `json:"feature,omitempty"`
	Metadata             *chargebee.Metadata             `json:"metadata,omitempty"`
	ImpactedItem         *chargebee.ImpactedItem         `json:"impacted_item,omitempty"`
	ImpactedSubscription *chargebee.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type ItemFamilyCreatedContent struct {
	ItemFamily *chargebee.ItemFamily `json:"item_family,omitempty"`
}

type OmnichannelSubscriptionItemScheduledChangeRemovedContent struct {
	OmnichannelSubscriptionItem                *chargebee.OmnichannelSubscriptionItem                `json:"omnichannel_subscription_item,omitempty"`
	OmnichannelSubscriptionItemScheduledChange *chargebee.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`
	Customer                                   *chargebee.Customer                                   `json:"customer,omitempty"`
}

type OmnichannelSubscriptionItemResumedContent struct {
	OmnichannelSubscription     *chargebee.OmnichannelSubscription     `json:"omnichannel_subscription,omitempty"`
	OmnichannelSubscriptionItem *chargebee.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`
	Customer                    *chargebee.Customer                    `json:"customer,omitempty"`
}

type PurchaseCreatedContent struct {
	Purchase *chargebee.Purchase `json:"purchase,omitempty"`
}

type EntitlementOverridesUpdatedContent struct {
	ImpactedSubscription *chargebee.ImpactedSubscription `json:"impacted_subscription,omitempty"`
	Metadata             *chargebee.Metadata             `json:"metadata,omitempty"`
}

type ItemFamilyDeletedContent struct {
	ItemFamily *chargebee.ItemFamily `json:"item_family,omitempty"`
}

type SubscriptionResumptionScheduledContent struct {
	Subscription           *chargebee.Subscription           `json:"subscription,omitempty"`
	Customer               *chargebee.Customer               `json:"customer,omitempty"`
	Card                   *chargebee.Card                   `json:"card,omitempty"`
	AdvanceInvoiceSchedule *chargebee.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type FeatureReactivatedContent struct {
	Feature  *chargebee.Feature  `json:"feature,omitempty"`
	Metadata *chargebee.Metadata `json:"metadata,omitempty"`
}

type CouponCodesDeletedContent struct {
	Coupon     *chargebee.Coupon     `json:"coupon,omitempty"`
	CouponSet  *chargebee.CouponSet  `json:"coupon_set,omitempty"`
	CouponCode *chargebee.CouponCode `json:"coupon_code,omitempty"`
}

type CardExpiredContent struct {
	Customer *chargebee.Customer `json:"customer,omitempty"`
	Card     *chargebee.Card     `json:"card,omitempty"`
}

type CreditNoteUpdatedContent struct {
	CreditNote *chargebee.CreditNote `json:"credit_note,omitempty"`
}

type OmnichannelSubscriptionItemDowngradedContent struct {
	OmnichannelSubscriptionItem *chargebee.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`
	OmnichannelSubscription     *chargebee.OmnichannelSubscription     `json:"omnichannel_subscription,omitempty"`
	OmnichannelTransaction      *chargebee.OmnichannelTransaction      `json:"omnichannel_transaction,omitempty"`
	Customer                    *chargebee.Customer                    `json:"customer,omitempty"`
}

type PriceVariantUpdatedContent struct {
	PriceVariant *chargebee.PriceVariant `json:"price_variant,omitempty"`
	Attribute    *chargebee.Attribute    `json:"attribute,omitempty"`
}

type PromotionalCreditsDeductedContent struct {
	Customer          *chargebee.Customer          `json:"customer,omitempty"`
	PromotionalCredit *chargebee.PromotionalCredit `json:"promotional_credit,omitempty"`
}

type SubscriptionRampAppliedContent struct {
	Ramp *chargebee.Ramp `json:"ramp,omitempty"`
}

type SubscriptionPausedContent struct {
	Subscription   *chargebee.Subscription   `json:"subscription,omitempty"`
	Customer       *chargebee.Customer       `json:"customer,omitempty"`
	Card           *chargebee.Card           `json:"card,omitempty"`
	Invoice        *chargebee.Invoice        `json:"invoice,omitempty"`
	CreditNote     *chargebee.CreditNote     `json:"credit_note,omitempty"`
	UnbilledCharge *chargebee.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type OrderReadyToProcessContent struct {
	Order *chargebee.Order `json:"order,omitempty"`
}

type FeatureCreatedContent struct {
	Feature              *chargebee.Feature              `json:"feature,omitempty"`
	Metadata             *chargebee.Metadata             `json:"metadata,omitempty"`
	ImpactedItem         *chargebee.ImpactedItem         `json:"impacted_item,omitempty"`
	ImpactedSubscription *chargebee.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type TransactionDeletedContent struct {
	Transaction *chargebee.Transaction `json:"transaction,omitempty"`
}

type CreditNoteCreatedContent struct {
	CreditNote *chargebee.CreditNote `json:"credit_note,omitempty"`
}

type OmnichannelSubscriptionItemResubscribedContent struct {
	OmnichannelSubscriptionItem                *chargebee.OmnichannelSubscriptionItem                `json:"omnichannel_subscription_item,omitempty"`
	OmnichannelSubscription                    *chargebee.OmnichannelSubscription                    `json:"omnichannel_subscription,omitempty"`
	OmnichannelTransaction                     *chargebee.OmnichannelTransaction                     `json:"omnichannel_transaction,omitempty"`
	OmnichannelSubscriptionItemScheduledChange *chargebee.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`
	Customer                                   *chargebee.Customer                                   `json:"customer,omitempty"`
}

type RecordPurchaseFailedContent struct {
	RecordedPurchase *chargebee.RecordedPurchase `json:"recorded_purchase,omitempty"`
	Customer         *chargebee.Customer         `json:"customer,omitempty"`
}

type ItemCreatedContent struct {
	Item *chargebee.Item `json:"item,omitempty"`
}

type TransactionUpdatedContent struct {
	Transaction *chargebee.Transaction `json:"transaction,omitempty"`
}

type VariantDeletedContent struct {
}

type MrrUpdatedContent struct {
	Subscription *chargebee.Subscription `json:"subscription,omitempty"`
}

type UnbilledChargesInvoicedContent struct {
	UnbilledCharge *chargebee.UnbilledCharge `json:"unbilled_charge,omitempty"`
	Invoice        *chargebee.Invoice        `json:"invoice,omitempty"`
}

type ItemPriceUpdatedContent struct {
	ItemPrice *chargebee.ItemPrice `json:"item_price,omitempty"`
}

type CouponCodesUpdatedContent struct {
	Coupon    *chargebee.Coupon    `json:"coupon,omitempty"`
	CouponSet *chargebee.CouponSet `json:"coupon_set,omitempty"`
}

type VirtualBankAccountUpdatedContent struct {
	Customer           *chargebee.Customer           `json:"customer,omitempty"`
	VirtualBankAccount *chargebee.VirtualBankAccount `json:"virtual_bank_account,omitempty"`
}

type ContractTermCreatedContent struct {
	ContractTerm *chargebee.ContractTerm `json:"contract_term,omitempty"`
}

type SubscriptionChangedContent struct {
	Subscription   *chargebee.Subscription   `json:"subscription,omitempty"`
	Customer       *chargebee.Customer       `json:"customer,omitempty"`
	Card           *chargebee.Card           `json:"card,omitempty"`
	Invoice        *chargebee.Invoice        `json:"invoice,omitempty"`
	CreditNote     *chargebee.CreditNote     `json:"credit_note,omitempty"`
	UnbilledCharge *chargebee.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type PaymentFailedContent struct {
	Transaction  *chargebee.Transaction  `json:"transaction,omitempty"`
	Invoice      *chargebee.Invoice      `json:"invoice,omitempty"`
	Customer     *chargebee.Customer     `json:"customer,omitempty"`
	Subscription *chargebee.Subscription `json:"subscription,omitempty"`
	Card         *chargebee.Card         `json:"card,omitempty"`
}

type CreditNoteDeletedContent struct {
	CreditNote *chargebee.CreditNote `json:"credit_note,omitempty"`
}

type TaxWithheldRefundedContent struct {
	TaxWithheld *chargebee.TaxWithheld `json:"tax_withheld,omitempty"`
	Invoice     *chargebee.Invoice     `json:"invoice,omitempty"`
	CreditNote  *chargebee.CreditNote  `json:"credit_note,omitempty"`
}

type ContractTermCompletedContent struct {
	ContractTerm *chargebee.ContractTerm `json:"contract_term,omitempty"`
}

type PaymentSchedulesUpdatedContent struct {
	PaymentSchedule *chargebee.PaymentSchedule `json:"payment_schedule,omitempty"`
}

type OmnichannelSubscriptionItemExpiredContent struct {
	OmnichannelSubscription     *chargebee.OmnichannelSubscription     `json:"omnichannel_subscription,omitempty"`
	OmnichannelSubscriptionItem *chargebee.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`
	Customer                    *chargebee.Customer                    `json:"customer,omitempty"`
}

type CardUpdatedContent struct {
	Customer *chargebee.Customer `json:"customer,omitempty"`
	Card     *chargebee.Card     `json:"card,omitempty"`
}

type CustomerCreatedContent struct {
	Customer *chargebee.Customer `json:"customer,omitempty"`
	Card     *chargebee.Card     `json:"card,omitempty"`
}

type SubscriptionRenewalReminderContent struct {
	Subscription           *chargebee.Subscription           `json:"subscription,omitempty"`
	Customer               *chargebee.Customer               `json:"customer,omitempty"`
	Card                   *chargebee.Card                   `json:"card,omitempty"`
	AdvanceInvoiceSchedule *chargebee.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type NetdPaymentDueReminderContent struct {
	Invoice *chargebee.Invoice `json:"invoice,omitempty"`
}

type OrderDeliveredContent struct {
	Order *chargebee.Order `json:"order,omitempty"`
}

type OmnichannelSubscriptionItemCancellationScheduledContent struct {
	OmnichannelSubscription     *chargebee.OmnichannelSubscription     `json:"omnichannel_subscription,omitempty"`
	OmnichannelSubscriptionItem *chargebee.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`
	Customer                    *chargebee.Customer                    `json:"customer,omitempty"`
}

type OmnichannelSubscriptionItemGracePeriodExpiredContent struct {
	OmnichannelSubscription     *chargebee.OmnichannelSubscription     `json:"omnichannel_subscription,omitempty"`
	OmnichannelSubscriptionItem *chargebee.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`
	Customer                    *chargebee.Customer                    `json:"customer,omitempty"`
}

type CouponCodesAddedContent struct {
	Coupon    *chargebee.Coupon    `json:"coupon,omitempty"`
	CouponSet *chargebee.CouponSet `json:"coupon_set,omitempty"`
}

type GiftCancelledContent struct {
	Gift *chargebee.Gift `json:"gift,omitempty"`
}

type OrderCancelledContent struct {
	Order *chargebee.Order `json:"order,omitempty"`
}

type CouponDeletedContent struct {
	Coupon *chargebee.Coupon `json:"coupon,omitempty"`
}

type SubscriptionScheduledChangesRemovedContent struct {
	Subscription           *chargebee.Subscription           `json:"subscription,omitempty"`
	Customer               *chargebee.Customer               `json:"customer,omitempty"`
	Card                   *chargebee.Card                   `json:"card,omitempty"`
	AdvanceInvoiceSchedule *chargebee.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type PendingInvoiceCreatedContent struct {
	Invoice *chargebee.Invoice `json:"invoice,omitempty"`
}

type ProductDeletedContent struct {
}

type EntitlementOverridesAutoRemovedContent struct {
	Feature              *chargebee.Feature              `json:"feature,omitempty"`
	Metadata             *chargebee.Metadata             `json:"metadata,omitempty"`
	ImpactedItem         *chargebee.ImpactedItem         `json:"impacted_item,omitempty"`
	ImpactedSubscription *chargebee.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type OmnichannelSubscriptionItemUpgradedContent struct {
	OmnichannelSubscriptionItem                *chargebee.OmnichannelSubscriptionItem                `json:"omnichannel_subscription_item,omitempty"`
	OmnichannelSubscription                    *chargebee.OmnichannelSubscription                    `json:"omnichannel_subscription,omitempty"`
	OmnichannelTransaction                     *chargebee.OmnichannelTransaction                     `json:"omnichannel_transaction,omitempty"`
	OmnichannelSubscriptionItemScheduledChange *chargebee.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`
	Customer                                   *chargebee.Customer                                   `json:"customer,omitempty"`
}

type SubscriptionBusinessEntityChangedContent struct {
	BusinessEntityTransfer *chargebee.BusinessEntityTransfer `json:"business_entity_transfer,omitempty"`
	Subscription           *chargebee.Subscription           `json:"subscription,omitempty"`
}

type OmnichannelOneTimeOrderCreatedContent struct {
	OmnichannelOneTimeOrder     *chargebee.OmnichannelOneTimeOrder     `json:"omnichannel_one_time_order,omitempty"`
	OmnichannelOneTimeOrderItem *chargebee.OmnichannelOneTimeOrderItem `json:"omnichannel_one_time_order_item,omitempty"`
	OmnichannelTransaction      *chargebee.OmnichannelTransaction      `json:"omnichannel_transaction,omitempty"`
	Customer                    *chargebee.Customer                    `json:"customer,omitempty"`
}

type PaymentSourceDeletedContent struct {
	Customer      *chargebee.Customer      `json:"customer,omitempty"`
	PaymentSource *chargebee.PaymentSource `json:"payment_source,omitempty"`
}

type OmnichannelSubscriptionItemCancelledContent struct {
	OmnichannelSubscription     *chargebee.OmnichannelSubscription     `json:"omnichannel_subscription,omitempty"`
	OmnichannelSubscriptionItem *chargebee.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`
	Customer                    *chargebee.Customer                    `json:"customer,omitempty"`
}

type QuoteDeletedContent struct {
	Quote *chargebee.Quote `json:"quote,omitempty"`
}

type InvoiceUpdatedContent struct {
	Invoice *chargebee.Invoice `json:"invoice,omitempty"`
}

type SubscriptionAdvanceInvoiceScheduleRemovedContent struct {
	Subscription           *chargebee.Subscription           `json:"subscription,omitempty"`
	Customer               *chargebee.Customer               `json:"customer,omitempty"`
	AdvanceInvoiceSchedule *chargebee.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type CardDeletedContent struct {
	Customer *chargebee.Customer `json:"customer,omitempty"`
	Card     *chargebee.Card     `json:"card,omitempty"`
}

type OrderReadyToShipContent struct {
	Order *chargebee.Order `json:"order,omitempty"`
}

type VariantUpdatedContent struct {
}

type SubscriptionMovedOutContent struct {
	Subscription *chargebee.Subscription `json:"subscription,omitempty"`
}

type PaymentScheduleSchemeCreatedContent struct {
	PaymentScheduleScheme *chargebee.PaymentScheduleScheme `json:"payment_schedule_scheme,omitempty"`
}

type BusinessEntityUpdatedContent struct {
	BusinessEntity *chargebee.BusinessEntity `json:"business_entity,omitempty"`
}

type SubscriptionScheduledResumptionRemovedContent struct {
	Subscription           *chargebee.Subscription           `json:"subscription,omitempty"`
	Customer               *chargebee.Customer               `json:"customer,omitempty"`
	Card                   *chargebee.Card                   `json:"card,omitempty"`
	AdvanceInvoiceSchedule *chargebee.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type PaymentInitiatedContent struct {
	Transaction  *chargebee.Transaction  `json:"transaction,omitempty"`
	Invoice      *chargebee.Invoice      `json:"invoice,omitempty"`
	Customer     *chargebee.Customer     `json:"customer,omitempty"`
	Subscription *chargebee.Subscription `json:"subscription,omitempty"`
	Card         *chargebee.Card         `json:"card,omitempty"`
}

type FeatureArchivedContent struct {
	Feature  *chargebee.Feature  `json:"feature,omitempty"`
	Metadata *chargebee.Metadata `json:"metadata,omitempty"`
}

type SubscriptionReactivatedWithBackdatingContent struct {
	Subscription   *chargebee.Subscription   `json:"subscription,omitempty"`
	Customer       *chargebee.Customer       `json:"customer,omitempty"`
	Card           *chargebee.Card           `json:"card,omitempty"`
	Invoice        *chargebee.Invoice        `json:"invoice,omitempty"`
	UnbilledCharge *chargebee.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type OmnichannelSubscriptionImportedContent struct {
	OmnichannelSubscriptionItem                *chargebee.OmnichannelSubscriptionItem                `json:"omnichannel_subscription_item,omitempty"`
	OmnichannelSubscription                    *chargebee.OmnichannelSubscription                    `json:"omnichannel_subscription,omitempty"`
	OmnichannelTransaction                     *chargebee.OmnichannelTransaction                     `json:"omnichannel_transaction,omitempty"`
	OmnichannelSubscriptionItemScheduledChange *chargebee.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`
	Customer                                   *chargebee.Customer                                   `json:"customer,omitempty"`
}

type TokenExpiredContent struct {
	Token *chargebee.Token `json:"token,omitempty"`
}

type CardAddedContent struct {
	Customer *chargebee.Customer `json:"customer,omitempty"`
	Card     *chargebee.Card     `json:"card,omitempty"`
}

type CouponCreatedContent struct {
	Coupon *chargebee.Coupon `json:"coupon,omitempty"`
}

type RuleDeletedContent struct {
	Rule *chargebee.Rule `json:"rule,omitempty"`
}

type ItemPriceEntitlementsUpdatedContent struct {
	Feature              *chargebee.Feature              `json:"feature,omitempty"`
	Metadata             *chargebee.Metadata             `json:"metadata,omitempty"`
	ImpactedItemPrice    *chargebee.ImpactedItemPrice    `json:"impacted_item_price,omitempty"`
	ImpactedSubscription *chargebee.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type ItemPriceDeletedContent struct {
	ItemPrice *chargebee.ItemPrice `json:"item_price,omitempty"`
}

type VirtualBankAccountDeletedContent struct {
	Customer           *chargebee.Customer           `json:"customer,omitempty"`
	VirtualBankAccount *chargebee.VirtualBankAccount `json:"virtual_bank_account,omitempty"`
}

type PaymentScheduleSchemeDeletedContent struct {
	PaymentScheduleScheme *chargebee.PaymentScheduleScheme `json:"payment_schedule_scheme,omitempty"`
}

type SubscriptionCreatedContent struct {
	Subscription   *chargebee.Subscription   `json:"subscription,omitempty"`
	Customer       *chargebee.Customer       `json:"customer,omitempty"`
	Card           *chargebee.Card           `json:"card,omitempty"`
	Invoice        *chargebee.Invoice        `json:"invoice,omitempty"`
	UnbilledCharge *chargebee.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type SubscriptionEntitlementsCreatedContent struct {
	SubscriptionEntitlementsCreatedDetail *chargebee.SubscriptionEntitlementsCreatedDetail `json:"subscription_entitlements_created_detail,omitempty"`
}

type OrderReturnedContent struct {
	Order *chargebee.Order `json:"order,omitempty"`
}

type SubscriptionDeletedContent struct {
	Subscription           *chargebee.Subscription           `json:"subscription,omitempty"`
	Customer               *chargebee.Customer               `json:"customer,omitempty"`
	Card                   *chargebee.Card                   `json:"card,omitempty"`
	AdvanceInvoiceSchedule *chargebee.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type PaymentSourceAddedContent struct {
	Customer      *chargebee.Customer      `json:"customer,omitempty"`
	PaymentSource *chargebee.PaymentSource `json:"payment_source,omitempty"`
}

type SubscriptionMovedInContent struct {
	Subscription *chargebee.Subscription `json:"subscription,omitempty"`
}

type ItemPriceCreatedContent struct {
	ItemPrice *chargebee.ItemPrice `json:"item_price,omitempty"`
}

type SubscriptionScheduledCancellationRemovedContent struct {
	Subscription           *chargebee.Subscription           `json:"subscription,omitempty"`
	Customer               *chargebee.Customer               `json:"customer,omitempty"`
	Card                   *chargebee.Card                   `json:"card,omitempty"`
	AdvanceInvoiceSchedule *chargebee.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type PaymentRefundedContent struct {
	Transaction  *chargebee.Transaction  `json:"transaction,omitempty"`
	Invoice      *chargebee.Invoice      `json:"invoice,omitempty"`
	CreditNote   *chargebee.CreditNote   `json:"credit_note,omitempty"`
	Customer     *chargebee.Customer     `json:"customer,omitempty"`
	Subscription *chargebee.Subscription `json:"subscription,omitempty"`
	Card         *chargebee.Card         `json:"card,omitempty"`
}

type UsageFileIngestedContent struct {
	UsageFile *chargebee.UsageFile `json:"usage_file,omitempty"`
}

type ProductCreatedContent struct {
}

type OmnichannelSubscriptionMovedInContent struct {
	OmnichannelSubscription *chargebee.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`
	Customer                *chargebee.Customer                `json:"customer,omitempty"`
}

type DifferentialPriceCreatedContent struct {
	DifferentialPrice *chargebee.DifferentialPrice `json:"differential_price,omitempty"`
}

type TransactionCreatedContent struct {
	Transaction *chargebee.Transaction `json:"transaction,omitempty"`
}

type OmnichannelSubscriptionItemDowngradeScheduledContent struct {
	OmnichannelSubscription     *chargebee.OmnichannelSubscription     `json:"omnichannel_subscription,omitempty"`
	OmnichannelSubscriptionItem *chargebee.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`
	Customer                    *chargebee.Customer                    `json:"customer,omitempty"`
}

type PaymentSucceededContent struct {
	Transaction  *chargebee.Transaction  `json:"transaction,omitempty"`
	Invoice      *chargebee.Invoice      `json:"invoice,omitempty"`
	Customer     *chargebee.Customer     `json:"customer,omitempty"`
	Subscription *chargebee.Subscription `json:"subscription,omitempty"`
	Card         *chargebee.Card         `json:"card,omitempty"`
}

type SubscriptionCanceledWithBackdatingContent struct {
	Subscription   *chargebee.Subscription   `json:"subscription,omitempty"`
	Customer       *chargebee.Customer       `json:"customer,omitempty"`
	Card           *chargebee.Card           `json:"card,omitempty"`
	Invoice        *chargebee.Invoice        `json:"invoice,omitempty"`
	CreditNote     *chargebee.CreditNote     `json:"credit_note,omitempty"`
	UnbilledCharge *chargebee.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type UnbilledChargesVoidedContent struct {
	UnbilledCharge *chargebee.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type QuoteCreatedContent struct {
	Quote *chargebee.Quote `json:"quote,omitempty"`
}

type CouponSetDeletedContent struct {
	Coupon    *chargebee.Coupon    `json:"coupon,omitempty"`
	CouponSet *chargebee.CouponSet `json:"coupon_set,omitempty"`
}

type AttachedItemCreatedContent struct {
	AttachedItem *chargebee.AttachedItem `json:"attached_item,omitempty"`
}

type SalesOrderCreatedContent struct {
}

type CustomerChangedContent struct {
	Customer *chargebee.Customer `json:"customer,omitempty"`
	Card     *chargebee.Card     `json:"card,omitempty"`
}

type SubscriptionStartedContent struct {
	Subscription *chargebee.Subscription `json:"subscription,omitempty"`
	Customer     *chargebee.Customer     `json:"customer,omitempty"`
	Card         *chargebee.Card         `json:"card,omitempty"`
	Invoice      *chargebee.Invoice      `json:"invoice,omitempty"`
}

type SubscriptionActivatedContent struct {
	Subscription *chargebee.Subscription `json:"subscription,omitempty"`
	Customer     *chargebee.Customer     `json:"customer,omitempty"`
	Card         *chargebee.Card         `json:"card,omitempty"`
	Invoice      *chargebee.Invoice      `json:"invoice,omitempty"`
}

type PaymentSourceExpiringContent struct {
	Customer      *chargebee.Customer      `json:"customer,omitempty"`
	PaymentSource *chargebee.PaymentSource `json:"payment_source,omitempty"`
}

type SubscriptionReactivatedContent struct {
	Subscription   *chargebee.Subscription   `json:"subscription,omitempty"`
	Customer       *chargebee.Customer       `json:"customer,omitempty"`
	Card           *chargebee.Card           `json:"card,omitempty"`
	Invoice        *chargebee.Invoice        `json:"invoice,omitempty"`
	UnbilledCharge *chargebee.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type OrderUpdatedContent struct {
	Order *chargebee.Order `json:"order,omitempty"`
}

type SubscriptionScheduledPauseRemovedContent struct {
	Subscription           *chargebee.Subscription           `json:"subscription,omitempty"`
	Customer               *chargebee.Customer               `json:"customer,omitempty"`
	Card                   *chargebee.Card                   `json:"card,omitempty"`
	AdvanceInvoiceSchedule *chargebee.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type SubscriptionCancellationReminderContent struct {
	Subscription           *chargebee.Subscription           `json:"subscription,omitempty"`
	Customer               *chargebee.Customer               `json:"customer,omitempty"`
	Card                   *chargebee.Card                   `json:"card,omitempty"`
	AdvanceInvoiceSchedule *chargebee.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type SubscriptionCreatedWithBackdatingContent struct {
	Subscription   *chargebee.Subscription   `json:"subscription,omitempty"`
	Customer       *chargebee.Customer       `json:"customer,omitempty"`
	Card           *chargebee.Card           `json:"card,omitempty"`
	Invoice        *chargebee.Invoice        `json:"invoice,omitempty"`
	UnbilledCharge *chargebee.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type SubscriptionRampCreatedContent struct {
	Ramp *chargebee.Ramp `json:"ramp,omitempty"`
}

type OrderDeletedContent struct {
	Order *chargebee.Order `json:"order,omitempty"`
}

type OmnichannelSubscriptionItemPauseScheduledContent struct {
	OmnichannelSubscriptionItem                *chargebee.OmnichannelSubscriptionItem                `json:"omnichannel_subscription_item,omitempty"`
	OmnichannelSubscriptionItemScheduledChange *chargebee.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`
	Customer                                   *chargebee.Customer                                   `json:"customer,omitempty"`
}

type GiftUpdatedContent struct {
	Gift *chargebee.Gift `json:"gift,omitempty"`
}

type SubscriptionTrialExtendedContent struct {
	Subscription           *chargebee.Subscription           `json:"subscription,omitempty"`
	Customer               *chargebee.Customer               `json:"customer,omitempty"`
	Card                   *chargebee.Card                   `json:"card,omitempty"`
	AdvanceInvoiceSchedule *chargebee.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type OmnichannelSubscriptionItemGracePeriodStartedContent struct {
	OmnichannelSubscription     *chargebee.OmnichannelSubscription     `json:"omnichannel_subscription,omitempty"`
	OmnichannelSubscriptionItem *chargebee.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`
	Customer                    *chargebee.Customer                    `json:"customer,omitempty"`
}

type CardExpiryReminderContent struct {
	Customer *chargebee.Customer `json:"customer,omitempty"`
	Card     *chargebee.Card     `json:"card,omitempty"`
}

type TokenCreatedContent struct {
	Token *chargebee.Token `json:"token,omitempty"`
}

type PromotionalCreditsAddedContent struct {
	Customer          *chargebee.Customer          `json:"customer,omitempty"`
	PromotionalCredit *chargebee.PromotionalCredit `json:"promotional_credit,omitempty"`
}

type SubscriptionRampUpdatedContent struct {
	Ramp *chargebee.Ramp `json:"ramp,omitempty"`
}

type CustomerEntitlementsUpdatedContent struct {
	ImpactedCustomer *chargebee.ImpactedCustomer `json:"impacted_customer,omitempty"`
}

type PaymentSourceExpiredContent struct {
	Customer      *chargebee.Customer      `json:"customer,omitempty"`
	PaymentSource *chargebee.PaymentSource `json:"payment_source,omitempty"`
}

type CustomerMovedOutContent struct {
	Customer *chargebee.Customer `json:"customer,omitempty"`
	Card     *chargebee.Card     `json:"card,omitempty"`
}

type SubscriptionEntitlementsUpdatedContent struct {
	SubscriptionEntitlementsUpdatedDetail *chargebee.SubscriptionEntitlementsUpdatedDetail `json:"subscription_entitlements_updated_detail,omitempty"`
}

type OmnichannelSubscriptionItemDunningExpiredContent struct {
	OmnichannelSubscription     *chargebee.OmnichannelSubscription     `json:"omnichannel_subscription,omitempty"`
	OmnichannelSubscriptionItem *chargebee.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`
	Customer                    *chargebee.Customer                    `json:"customer,omitempty"`
}

type HierarchyCreatedContent struct {
	Customer *chargebee.Customer `json:"customer,omitempty"`
}

type AttachedItemDeletedContent struct {
	AttachedItem *chargebee.AttachedItem `json:"attached_item,omitempty"`
}

type OmnichannelSubscriptionItemScheduledCancellationRemovedContent struct {
	OmnichannelSubscription     *chargebee.OmnichannelSubscription     `json:"omnichannel_subscription,omitempty"`
	OmnichannelSubscriptionItem *chargebee.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`
	Customer                    *chargebee.Customer                    `json:"customer,omitempty"`
}

type ItemUpdatedContent struct {
	Item *chargebee.Item `json:"item,omitempty"`
}

type CouponSetCreatedContent struct {
	Coupon    *chargebee.Coupon    `json:"coupon,omitempty"`
	CouponSet *chargebee.CouponSet `json:"coupon_set,omitempty"`
}

type PaymentIntentUpdatedContent struct {
	PaymentIntent *chargebee.PaymentIntent `json:"payment_intent,omitempty"`
}

type OrderResentContent struct {
	Order *chargebee.Order `json:"order,omitempty"`
}

type OmnichannelSubscriptionItemScheduledDowngradeRemovedContent struct {
	OmnichannelSubscription     *chargebee.OmnichannelSubscription     `json:"omnichannel_subscription,omitempty"`
	OmnichannelSubscriptionItem *chargebee.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`
	Customer                    *chargebee.Customer                    `json:"customer,omitempty"`
}

type OmnichannelSubscriptionCreatedContent struct {
	OmnichannelSubscriptionItem                *chargebee.OmnichannelSubscriptionItem                `json:"omnichannel_subscription_item,omitempty"`
	OmnichannelSubscription                    *chargebee.OmnichannelSubscription                    `json:"omnichannel_subscription,omitempty"`
	OmnichannelTransaction                     *chargebee.OmnichannelTransaction                     `json:"omnichannel_transaction,omitempty"`
	OmnichannelSubscriptionItemScheduledChange *chargebee.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`
	Customer                                   *chargebee.Customer                                   `json:"customer,omitempty"`
}

type TaxWithheldRecordedContent struct {
	TaxWithheld *chargebee.TaxWithheld `json:"tax_withheld,omitempty"`
	Invoice     *chargebee.Invoice     `json:"invoice,omitempty"`
	CreditNote  *chargebee.CreditNote  `json:"credit_note,omitempty"`
}

type PriceVariantCreatedContent struct {
	PriceVariant *chargebee.PriceVariant `json:"price_variant,omitempty"`
	Attribute    *chargebee.Attribute    `json:"attribute,omitempty"`
}

type DifferentialPriceDeletedContent struct {
	DifferentialPrice *chargebee.DifferentialPrice `json:"differential_price,omitempty"`
}

type SubscriptionItemsRenewedContent struct {
	Subscription   *chargebee.Subscription   `json:"subscription,omitempty"`
	Customer       *chargebee.Customer       `json:"customer,omitempty"`
	Card           *chargebee.Card           `json:"card,omitempty"`
	Invoice        *chargebee.Invoice        `json:"invoice,omitempty"`
	CreditNote     *chargebee.CreditNote     `json:"credit_note,omitempty"`
	UnbilledCharge *chargebee.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type RuleCreatedContent struct {
	Rule *chargebee.Rule `json:"rule,omitempty"`
}

type ContractTermCancelledContent struct {
	ContractTerm *chargebee.ContractTerm `json:"contract_term,omitempty"`
}

type ContractTermRenewedContent struct {
	ContractTerm *chargebee.ContractTerm `json:"contract_term,omitempty"`
}

type InvoiceDeletedContent struct {
	Invoice *chargebee.Invoice `json:"invoice,omitempty"`
}

type ItemPriceEntitlementsRemovedContent struct {
	Feature              *chargebee.Feature              `json:"feature,omitempty"`
	Metadata             *chargebee.Metadata             `json:"metadata,omitempty"`
	ImpactedItemPrice    *chargebee.ImpactedItemPrice    `json:"impacted_item_price,omitempty"`
	ImpactedSubscription *chargebee.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type SalesOrderUpdatedContent struct {
}

type OmnichannelSubscriptionItemDunningStartedContent struct {
	OmnichannelSubscription     *chargebee.OmnichannelSubscription     `json:"omnichannel_subscription,omitempty"`
	OmnichannelSubscriptionItem *chargebee.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`
	Customer                    *chargebee.Customer                    `json:"customer,omitempty"`
}

type OmnichannelSubscriptionItemChangeScheduledContent struct {
	OmnichannelSubscriptionItem                *chargebee.OmnichannelSubscriptionItem                `json:"omnichannel_subscription_item,omitempty"`
	OmnichannelSubscriptionItemScheduledChange *chargebee.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`
	Customer                                   *chargebee.Customer                                   `json:"customer,omitempty"`
}

type PendingInvoiceUpdatedContent struct {
	Invoice *chargebee.Invoice `json:"invoice,omitempty"`
}

type QuoteUpdatedContent struct {
	Quote *chargebee.Quote `json:"quote,omitempty"`
}

type AttachedItemUpdatedContent struct {
	AttachedItem *chargebee.AttachedItem `json:"attached_item,omitempty"`
}

type PaymentSourceUpdatedContent struct {
	Customer      *chargebee.Customer      `json:"customer,omitempty"`
	PaymentSource *chargebee.PaymentSource `json:"payment_source,omitempty"`
}

type BusinessEntityDeletedContent struct {
	BusinessEntity *chargebee.BusinessEntity `json:"business_entity,omitempty"`
}

type AuthorizationVoidedContent struct {
	Transaction *chargebee.Transaction `json:"transaction,omitempty"`
}

type SubscriptionRampDeletedContent struct {
	Ramp *chargebee.Ramp `json:"ramp,omitempty"`
}

type PlanDeletedContent struct {
	Plan *chargebee.Plan `json:"plan,omitempty"`
}

type AddonDeletedContent struct {
	Addon *chargebee.Addon `json:"addon,omitempty"`
}

type AddonUpdatedContent struct {
	Addon *chargebee.Addon `json:"addon,omitempty"`
}

type AddonCreatedContent struct {
	Addon *chargebee.Addon `json:"addon,omitempty"`
}

type PlanCreatedContent struct {
	Plan *chargebee.Plan `json:"plan,omitempty"`
}

type PlanUpdatedContent struct {
	Plan *chargebee.Plan `json:"plan,omitempty"`
}

// SubscriptionPauseScheduledEvent represents a subscription_pause_scheduled webhook event
type SubscriptionPauseScheduledEvent struct {
	BaseEvent
	Content *SubscriptionPauseScheduledContent `json:"content"`
}

// CustomerBusinessEntityChangedEvent represents a customer_business_entity_changed webhook event
type CustomerBusinessEntityChangedEvent struct {
	BaseEvent
	Content *CustomerBusinessEntityChangedContent `json:"content"`
}

// SubscriptionAdvanceInvoiceScheduleAddedEvent represents a subscription_advance_invoice_schedule_added webhook event
type SubscriptionAdvanceInvoiceScheduleAddedEvent struct {
	BaseEvent
	Content *SubscriptionAdvanceInvoiceScheduleAddedContent `json:"content"`
}

// GiftExpiredEvent represents a gift_expired webhook event
type GiftExpiredEvent struct {
	BaseEvent
	Content *GiftExpiredContent `json:"content"`
}

// TaxWithheldDeletedEvent represents a tax_withheld_deleted webhook event
type TaxWithheldDeletedEvent struct {
	BaseEvent
	Content *TaxWithheldDeletedContent `json:"content"`
}

// UnbilledChargesDeletedEvent represents a unbilled_charges_deleted webhook event
type UnbilledChargesDeletedEvent struct {
	BaseEvent
	Content *UnbilledChargesDeletedContent `json:"content"`
}

// CouponUpdatedEvent represents a coupon_updated webhook event
type CouponUpdatedEvent struct {
	BaseEvent
	Content *CouponUpdatedContent `json:"content"`
}

// ProductUpdatedEvent represents a product_updated webhook event
type ProductUpdatedEvent struct {
	BaseEvent
	Content *ProductUpdatedContent `json:"content"`
}

// OmnichannelSubscriptionItemReactivatedEvent represents a omnichannel_subscription_item_reactivated webhook event
type OmnichannelSubscriptionItemReactivatedEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemReactivatedContent `json:"content"`
}

// OmnichannelSubscriptionItemRenewedEvent represents a omnichannel_subscription_item_renewed webhook event
type OmnichannelSubscriptionItemRenewedEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemRenewedContent `json:"content"`
}

// UnbilledChargesCreatedEvent represents a unbilled_charges_created webhook event
type UnbilledChargesCreatedEvent struct {
	BaseEvent
	Content *UnbilledChargesCreatedContent `json:"content"`
}

// SubscriptionResumedEvent represents a subscription_resumed webhook event
type SubscriptionResumedEvent struct {
	BaseEvent
	Content *SubscriptionResumedContent `json:"content"`
}

// OmnichannelOneTimeOrderItemCancelledEvent represents a omnichannel_one_time_order_item_cancelled webhook event
type OmnichannelOneTimeOrderItemCancelledEvent struct {
	BaseEvent
	Content *OmnichannelOneTimeOrderItemCancelledContent `json:"content"`
}

// SubscriptionCancelledEvent represents a subscription_cancelled webhook event
type SubscriptionCancelledEvent struct {
	BaseEvent
	Content *SubscriptionCancelledContent `json:"content"`
}

// ItemEntitlementsRemovedEvent represents a item_entitlements_removed webhook event
type ItemEntitlementsRemovedEvent struct {
	BaseEvent
	Content *ItemEntitlementsRemovedContent `json:"content"`
}

// BusinessEntityCreatedEvent represents a business_entity_created webhook event
type BusinessEntityCreatedEvent struct {
	BaseEvent
	Content *BusinessEntityCreatedContent `json:"content"`
}

// CouponSetUpdatedEvent represents a coupon_set_updated webhook event
type CouponSetUpdatedEvent struct {
	BaseEvent
	Content *CouponSetUpdatedContent `json:"content"`
}

// DifferentialPriceUpdatedEvent represents a differential_price_updated webhook event
type DifferentialPriceUpdatedEvent struct {
	BaseEvent
	Content *DifferentialPriceUpdatedContent `json:"content"`
}

// OmnichannelSubscriptionItemPausedEvent represents a omnichannel_subscription_item_paused webhook event
type OmnichannelSubscriptionItemPausedEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemPausedContent `json:"content"`
}

// EntitlementOverridesRemovedEvent represents a entitlement_overrides_removed webhook event
type EntitlementOverridesRemovedEvent struct {
	BaseEvent
	Content *EntitlementOverridesRemovedContent `json:"content"`
}

// SubscriptionActivatedWithBackdatingEvent represents a subscription_activated_with_backdating webhook event
type SubscriptionActivatedWithBackdatingEvent struct {
	BaseEvent
	Content *SubscriptionActivatedWithBackdatingContent `json:"content"`
}

// SubscriptionTrialEndReminderEvent represents a subscription_trial_end_reminder webhook event
type SubscriptionTrialEndReminderEvent struct {
	BaseEvent
	Content *SubscriptionTrialEndReminderContent `json:"content"`
}

// SubscriptionShippingAddressUpdatedEvent represents a subscription_shipping_address_updated webhook event
type SubscriptionShippingAddressUpdatedEvent struct {
	BaseEvent
	Content *SubscriptionShippingAddressUpdatedContent `json:"content"`
}

// VoucherCreateFailedEvent represents a voucher_create_failed webhook event
type VoucherCreateFailedEvent struct {
	BaseEvent
	Content *VoucherCreateFailedContent `json:"content"`
}

// GiftClaimedEvent represents a gift_claimed webhook event
type GiftClaimedEvent struct {
	BaseEvent
	Content *GiftClaimedContent `json:"content"`
}

// CustomerDeletedEvent represents a customer_deleted webhook event
type CustomerDeletedEvent struct {
	BaseEvent
	Content *CustomerDeletedContent `json:"content"`
}

// RefundInitiatedEvent represents a refund_initiated webhook event
type RefundInitiatedEvent struct {
	BaseEvent
	Content *RefundInitiatedContent `json:"content"`
}

// InvoiceGeneratedWithBackdatingEvent represents a invoice_generated_with_backdating webhook event
type InvoiceGeneratedWithBackdatingEvent struct {
	BaseEvent
	Content *InvoiceGeneratedWithBackdatingContent `json:"content"`
}

// OmnichannelTransactionCreatedEvent represents a omnichannel_transaction_created webhook event
type OmnichannelTransactionCreatedEvent struct {
	BaseEvent
	Content *OmnichannelTransactionCreatedContent `json:"content"`
}

// AddUsagesReminderEvent represents a add_usages_reminder webhook event
type AddUsagesReminderEvent struct {
	BaseEvent
	Content *AddUsagesReminderContent `json:"content"`
}

// VoucherCreatedEvent represents a voucher_created webhook event
type VoucherCreatedEvent struct {
	BaseEvent
	Content *VoucherCreatedContent `json:"content"`
}

// RuleUpdatedEvent represents a rule_updated webhook event
type RuleUpdatedEvent struct {
	BaseEvent
	Content *RuleUpdatedContent `json:"content"`
}

// PaymentSchedulesCreatedEvent represents a payment_schedules_created webhook event
type PaymentSchedulesCreatedEvent struct {
	BaseEvent
	Content *PaymentSchedulesCreatedContent `json:"content"`
}

// FeatureActivatedEvent represents a feature_activated webhook event
type FeatureActivatedEvent struct {
	BaseEvent
	Content *FeatureActivatedContent `json:"content"`
}

// PaymentSourceLocallyDeletedEvent represents a payment_source_locally_deleted webhook event
type PaymentSourceLocallyDeletedEvent struct {
	BaseEvent
	Content *PaymentSourceLocallyDeletedContent `json:"content"`
}

// InvoiceGeneratedEvent represents a invoice_generated webhook event
type InvoiceGeneratedEvent struct {
	BaseEvent
	Content *InvoiceGeneratedContent `json:"content"`
}

// VoucherExpiredEvent represents a voucher_expired webhook event
type VoucherExpiredEvent struct {
	BaseEvent
	Content *VoucherExpiredContent `json:"content"`
}

// AuthorizationSucceededEvent represents a authorization_succeeded webhook event
type AuthorizationSucceededEvent struct {
	BaseEvent
	Content *AuthorizationSucceededContent `json:"content"`
}

// GiftScheduledEvent represents a gift_scheduled webhook event
type GiftScheduledEvent struct {
	BaseEvent
	Content *GiftScheduledContent `json:"content"`
}

// SubscriptionChangesScheduledEvent represents a subscription_changes_scheduled webhook event
type SubscriptionChangesScheduledEvent struct {
	BaseEvent
	Content *SubscriptionChangesScheduledContent `json:"content"`
}

// SubscriptionChangedWithBackdatingEvent represents a subscription_changed_with_backdating webhook event
type SubscriptionChangedWithBackdatingEvent struct {
	BaseEvent
	Content *SubscriptionChangedWithBackdatingContent `json:"content"`
}

// VariantCreatedEvent represents a variant_created webhook event
type VariantCreatedEvent struct {
	BaseEvent
	Content *VariantCreatedContent `json:"content"`
}

// OmnichannelSubscriptionItemChangedEvent represents a omnichannel_subscription_item_changed webhook event
type OmnichannelSubscriptionItemChangedEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemChangedContent `json:"content"`
}

// GiftUnclaimedEvent represents a gift_unclaimed webhook event
type GiftUnclaimedEvent struct {
	BaseEvent
	Content *GiftUnclaimedContent `json:"content"`
}

// VirtualBankAccountAddedEvent represents a virtual_bank_account_added webhook event
type VirtualBankAccountAddedEvent struct {
	BaseEvent
	Content *VirtualBankAccountAddedContent `json:"content"`
}

// PaymentIntentCreatedEvent represents a payment_intent_created webhook event
type PaymentIntentCreatedEvent struct {
	BaseEvent
	Content *PaymentIntentCreatedContent `json:"content"`
}

// CreditNoteCreatedWithBackdatingEvent represents a credit_note_created_with_backdating webhook event
type CreditNoteCreatedWithBackdatingEvent struct {
	BaseEvent
	Content *CreditNoteCreatedWithBackdatingContent `json:"content"`
}

// ContractTermTerminatedEvent represents a contract_term_terminated webhook event
type ContractTermTerminatedEvent struct {
	BaseEvent
	Content *ContractTermTerminatedContent `json:"content"`
}

// ItemFamilyUpdatedEvent represents a item_family_updated webhook event
type ItemFamilyUpdatedEvent struct {
	BaseEvent
	Content *ItemFamilyUpdatedContent `json:"content"`
}

// OrderCreatedEvent represents a order_created webhook event
type OrderCreatedEvent struct {
	BaseEvent
	Content *OrderCreatedContent `json:"content"`
}

// PriceVariantDeletedEvent represents a price_variant_deleted webhook event
type PriceVariantDeletedEvent struct {
	BaseEvent
	Content *PriceVariantDeletedContent `json:"content"`
}

// SubscriptionMovementFailedEvent represents a subscription_movement_failed webhook event
type SubscriptionMovementFailedEvent struct {
	BaseEvent
	Content *SubscriptionMovementFailedContent `json:"content"`
}

// CustomerMovedInEvent represents a customer_moved_in webhook event
type CustomerMovedInEvent struct {
	BaseEvent
	Content *CustomerMovedInContent `json:"content"`
}

// SubscriptionAdvanceInvoiceScheduleUpdatedEvent represents a subscription_advance_invoice_schedule_updated webhook event
type SubscriptionAdvanceInvoiceScheduleUpdatedEvent struct {
	BaseEvent
	Content *SubscriptionAdvanceInvoiceScheduleUpdatedContent `json:"content"`
}

// ItemDeletedEvent represents a item_deleted webhook event
type ItemDeletedEvent struct {
	BaseEvent
	Content *ItemDeletedContent `json:"content"`
}

// SubscriptionRampDraftedEvent represents a subscription_ramp_drafted webhook event
type SubscriptionRampDraftedEvent struct {
	BaseEvent
	Content *SubscriptionRampDraftedContent `json:"content"`
}

// DunningUpdatedEvent represents a dunning_updated webhook event
type DunningUpdatedEvent struct {
	BaseEvent
	Content *DunningUpdatedContent `json:"content"`
}

// ItemEntitlementsUpdatedEvent represents a item_entitlements_updated webhook event
type ItemEntitlementsUpdatedEvent struct {
	BaseEvent
	Content *ItemEntitlementsUpdatedContent `json:"content"`
}

// TokenConsumedEvent represents a token_consumed webhook event
type TokenConsumedEvent struct {
	BaseEvent
	Content *TokenConsumedContent `json:"content"`
}

// HierarchyDeletedEvent represents a hierarchy_deleted webhook event
type HierarchyDeletedEvent struct {
	BaseEvent
	Content *HierarchyDeletedContent `json:"content"`
}

// SubscriptionCancellationScheduledEvent represents a subscription_cancellation_scheduled webhook event
type SubscriptionCancellationScheduledEvent struct {
	BaseEvent
	Content *SubscriptionCancellationScheduledContent `json:"content"`
}

// SubscriptionRenewedEvent represents a subscription_renewed webhook event
type SubscriptionRenewedEvent struct {
	BaseEvent
	Content *SubscriptionRenewedContent `json:"content"`
}

// FeatureUpdatedEvent represents a feature_updated webhook event
type FeatureUpdatedEvent struct {
	BaseEvent
	Content *FeatureUpdatedContent `json:"content"`
}

// FeatureDeletedEvent represents a feature_deleted webhook event
type FeatureDeletedEvent struct {
	BaseEvent
	Content *FeatureDeletedContent `json:"content"`
}

// ItemFamilyCreatedEvent represents a item_family_created webhook event
type ItemFamilyCreatedEvent struct {
	BaseEvent
	Content *ItemFamilyCreatedContent `json:"content"`
}

// OmnichannelSubscriptionItemScheduledChangeRemovedEvent represents a omnichannel_subscription_item_scheduled_change_removed webhook event
type OmnichannelSubscriptionItemScheduledChangeRemovedEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemScheduledChangeRemovedContent `json:"content"`
}

// OmnichannelSubscriptionItemResumedEvent represents a omnichannel_subscription_item_resumed webhook event
type OmnichannelSubscriptionItemResumedEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemResumedContent `json:"content"`
}

// PurchaseCreatedEvent represents a purchase_created webhook event
type PurchaseCreatedEvent struct {
	BaseEvent
	Content *PurchaseCreatedContent `json:"content"`
}

// EntitlementOverridesUpdatedEvent represents a entitlement_overrides_updated webhook event
type EntitlementOverridesUpdatedEvent struct {
	BaseEvent
	Content *EntitlementOverridesUpdatedContent `json:"content"`
}

// ItemFamilyDeletedEvent represents a item_family_deleted webhook event
type ItemFamilyDeletedEvent struct {
	BaseEvent
	Content *ItemFamilyDeletedContent `json:"content"`
}

// SubscriptionResumptionScheduledEvent represents a subscription_resumption_scheduled webhook event
type SubscriptionResumptionScheduledEvent struct {
	BaseEvent
	Content *SubscriptionResumptionScheduledContent `json:"content"`
}

// FeatureReactivatedEvent represents a feature_reactivated webhook event
type FeatureReactivatedEvent struct {
	BaseEvent
	Content *FeatureReactivatedContent `json:"content"`
}

// CouponCodesDeletedEvent represents a coupon_codes_deleted webhook event
type CouponCodesDeletedEvent struct {
	BaseEvent
	Content *CouponCodesDeletedContent `json:"content"`
}

// CardExpiredEvent represents a card_expired webhook event
type CardExpiredEvent struct {
	BaseEvent
	Content *CardExpiredContent `json:"content"`
}

// CreditNoteUpdatedEvent represents a credit_note_updated webhook event
type CreditNoteUpdatedEvent struct {
	BaseEvent
	Content *CreditNoteUpdatedContent `json:"content"`
}

// OmnichannelSubscriptionItemDowngradedEvent represents a omnichannel_subscription_item_downgraded webhook event
type OmnichannelSubscriptionItemDowngradedEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemDowngradedContent `json:"content"`
}

// PriceVariantUpdatedEvent represents a price_variant_updated webhook event
type PriceVariantUpdatedEvent struct {
	BaseEvent
	Content *PriceVariantUpdatedContent `json:"content"`
}

// PromotionalCreditsDeductedEvent represents a promotional_credits_deducted webhook event
type PromotionalCreditsDeductedEvent struct {
	BaseEvent
	Content *PromotionalCreditsDeductedContent `json:"content"`
}

// SubscriptionRampAppliedEvent represents a subscription_ramp_applied webhook event
type SubscriptionRampAppliedEvent struct {
	BaseEvent
	Content *SubscriptionRampAppliedContent `json:"content"`
}

// SubscriptionPausedEvent represents a subscription_paused webhook event
type SubscriptionPausedEvent struct {
	BaseEvent
	Content *SubscriptionPausedContent `json:"content"`
}

// OrderReadyToProcessEvent represents a order_ready_to_process webhook event
type OrderReadyToProcessEvent struct {
	BaseEvent
	Content *OrderReadyToProcessContent `json:"content"`
}

// FeatureCreatedEvent represents a feature_created webhook event
type FeatureCreatedEvent struct {
	BaseEvent
	Content *FeatureCreatedContent `json:"content"`
}

// TransactionDeletedEvent represents a transaction_deleted webhook event
type TransactionDeletedEvent struct {
	BaseEvent
	Content *TransactionDeletedContent `json:"content"`
}

// CreditNoteCreatedEvent represents a credit_note_created webhook event
type CreditNoteCreatedEvent struct {
	BaseEvent
	Content *CreditNoteCreatedContent `json:"content"`
}

// OmnichannelSubscriptionItemResubscribedEvent represents a omnichannel_subscription_item_resubscribed webhook event
type OmnichannelSubscriptionItemResubscribedEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemResubscribedContent `json:"content"`
}

// RecordPurchaseFailedEvent represents a record_purchase_failed webhook event
type RecordPurchaseFailedEvent struct {
	BaseEvent
	Content *RecordPurchaseFailedContent `json:"content"`
}

// ItemCreatedEvent represents a item_created webhook event
type ItemCreatedEvent struct {
	BaseEvent
	Content *ItemCreatedContent `json:"content"`
}

// TransactionUpdatedEvent represents a transaction_updated webhook event
type TransactionUpdatedEvent struct {
	BaseEvent
	Content *TransactionUpdatedContent `json:"content"`
}

// VariantDeletedEvent represents a variant_deleted webhook event
type VariantDeletedEvent struct {
	BaseEvent
	Content *VariantDeletedContent `json:"content"`
}

// MrrUpdatedEvent represents a mrr_updated webhook event
type MrrUpdatedEvent struct {
	BaseEvent
	Content *MrrUpdatedContent `json:"content"`
}

// UnbilledChargesInvoicedEvent represents a unbilled_charges_invoiced webhook event
type UnbilledChargesInvoicedEvent struct {
	BaseEvent
	Content *UnbilledChargesInvoicedContent `json:"content"`
}

// ItemPriceUpdatedEvent represents a item_price_updated webhook event
type ItemPriceUpdatedEvent struct {
	BaseEvent
	Content *ItemPriceUpdatedContent `json:"content"`
}

// CouponCodesUpdatedEvent represents a coupon_codes_updated webhook event
type CouponCodesUpdatedEvent struct {
	BaseEvent
	Content *CouponCodesUpdatedContent `json:"content"`
}

// VirtualBankAccountUpdatedEvent represents a virtual_bank_account_updated webhook event
type VirtualBankAccountUpdatedEvent struct {
	BaseEvent
	Content *VirtualBankAccountUpdatedContent `json:"content"`
}

// ContractTermCreatedEvent represents a contract_term_created webhook event
type ContractTermCreatedEvent struct {
	BaseEvent
	Content *ContractTermCreatedContent `json:"content"`
}

// SubscriptionChangedEvent represents a subscription_changed webhook event
type SubscriptionChangedEvent struct {
	BaseEvent
	Content *SubscriptionChangedContent `json:"content"`
}

// PaymentFailedEvent represents a payment_failed webhook event
type PaymentFailedEvent struct {
	BaseEvent
	Content *PaymentFailedContent `json:"content"`
}

// CreditNoteDeletedEvent represents a credit_note_deleted webhook event
type CreditNoteDeletedEvent struct {
	BaseEvent
	Content *CreditNoteDeletedContent `json:"content"`
}

// TaxWithheldRefundedEvent represents a tax_withheld_refunded webhook event
type TaxWithheldRefundedEvent struct {
	BaseEvent
	Content *TaxWithheldRefundedContent `json:"content"`
}

// ContractTermCompletedEvent represents a contract_term_completed webhook event
type ContractTermCompletedEvent struct {
	BaseEvent
	Content *ContractTermCompletedContent `json:"content"`
}

// PaymentSchedulesUpdatedEvent represents a payment_schedules_updated webhook event
type PaymentSchedulesUpdatedEvent struct {
	BaseEvent
	Content *PaymentSchedulesUpdatedContent `json:"content"`
}

// OmnichannelSubscriptionItemExpiredEvent represents a omnichannel_subscription_item_expired webhook event
type OmnichannelSubscriptionItemExpiredEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemExpiredContent `json:"content"`
}

// CardUpdatedEvent represents a card_updated webhook event
type CardUpdatedEvent struct {
	BaseEvent
	Content *CardUpdatedContent `json:"content"`
}

// CustomerCreatedEvent represents a customer_created webhook event
type CustomerCreatedEvent struct {
	BaseEvent
	Content *CustomerCreatedContent `json:"content"`
}

// SubscriptionRenewalReminderEvent represents a subscription_renewal_reminder webhook event
type SubscriptionRenewalReminderEvent struct {
	BaseEvent
	Content *SubscriptionRenewalReminderContent `json:"content"`
}

// NetdPaymentDueReminderEvent represents a netd_payment_due_reminder webhook event
type NetdPaymentDueReminderEvent struct {
	BaseEvent
	Content *NetdPaymentDueReminderContent `json:"content"`
}

// OrderDeliveredEvent represents a order_delivered webhook event
type OrderDeliveredEvent struct {
	BaseEvent
	Content *OrderDeliveredContent `json:"content"`
}

// OmnichannelSubscriptionItemCancellationScheduledEvent represents a omnichannel_subscription_item_cancellation_scheduled webhook event
type OmnichannelSubscriptionItemCancellationScheduledEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemCancellationScheduledContent `json:"content"`
}

// OmnichannelSubscriptionItemGracePeriodExpiredEvent represents a omnichannel_subscription_item_grace_period_expired webhook event
type OmnichannelSubscriptionItemGracePeriodExpiredEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemGracePeriodExpiredContent `json:"content"`
}

// CouponCodesAddedEvent represents a coupon_codes_added webhook event
type CouponCodesAddedEvent struct {
	BaseEvent
	Content *CouponCodesAddedContent `json:"content"`
}

// GiftCancelledEvent represents a gift_cancelled webhook event
type GiftCancelledEvent struct {
	BaseEvent
	Content *GiftCancelledContent `json:"content"`
}

// OrderCancelledEvent represents a order_cancelled webhook event
type OrderCancelledEvent struct {
	BaseEvent
	Content *OrderCancelledContent `json:"content"`
}

// CouponDeletedEvent represents a coupon_deleted webhook event
type CouponDeletedEvent struct {
	BaseEvent
	Content *CouponDeletedContent `json:"content"`
}

// SubscriptionScheduledChangesRemovedEvent represents a subscription_scheduled_changes_removed webhook event
type SubscriptionScheduledChangesRemovedEvent struct {
	BaseEvent
	Content *SubscriptionScheduledChangesRemovedContent `json:"content"`
}

// PendingInvoiceCreatedEvent represents a pending_invoice_created webhook event
type PendingInvoiceCreatedEvent struct {
	BaseEvent
	Content *PendingInvoiceCreatedContent `json:"content"`
}

// ProductDeletedEvent represents a product_deleted webhook event
type ProductDeletedEvent struct {
	BaseEvent
	Content *ProductDeletedContent `json:"content"`
}

// EntitlementOverridesAutoRemovedEvent represents a entitlement_overrides_auto_removed webhook event
type EntitlementOverridesAutoRemovedEvent struct {
	BaseEvent
	Content *EntitlementOverridesAutoRemovedContent `json:"content"`
}

// OmnichannelSubscriptionItemUpgradedEvent represents a omnichannel_subscription_item_upgraded webhook event
type OmnichannelSubscriptionItemUpgradedEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemUpgradedContent `json:"content"`
}

// SubscriptionBusinessEntityChangedEvent represents a subscription_business_entity_changed webhook event
type SubscriptionBusinessEntityChangedEvent struct {
	BaseEvent
	Content *SubscriptionBusinessEntityChangedContent `json:"content"`
}

// OmnichannelOneTimeOrderCreatedEvent represents a omnichannel_one_time_order_created webhook event
type OmnichannelOneTimeOrderCreatedEvent struct {
	BaseEvent
	Content *OmnichannelOneTimeOrderCreatedContent `json:"content"`
}

// PaymentSourceDeletedEvent represents a payment_source_deleted webhook event
type PaymentSourceDeletedEvent struct {
	BaseEvent
	Content *PaymentSourceDeletedContent `json:"content"`
}

// OmnichannelSubscriptionItemCancelledEvent represents a omnichannel_subscription_item_cancelled webhook event
type OmnichannelSubscriptionItemCancelledEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemCancelledContent `json:"content"`
}

// QuoteDeletedEvent represents a quote_deleted webhook event
type QuoteDeletedEvent struct {
	BaseEvent
	Content *QuoteDeletedContent `json:"content"`
}

// InvoiceUpdatedEvent represents a invoice_updated webhook event
type InvoiceUpdatedEvent struct {
	BaseEvent
	Content *InvoiceUpdatedContent `json:"content"`
}

// SubscriptionAdvanceInvoiceScheduleRemovedEvent represents a subscription_advance_invoice_schedule_removed webhook event
type SubscriptionAdvanceInvoiceScheduleRemovedEvent struct {
	BaseEvent
	Content *SubscriptionAdvanceInvoiceScheduleRemovedContent `json:"content"`
}

// CardDeletedEvent represents a card_deleted webhook event
type CardDeletedEvent struct {
	BaseEvent
	Content *CardDeletedContent `json:"content"`
}

// OrderReadyToShipEvent represents a order_ready_to_ship webhook event
type OrderReadyToShipEvent struct {
	BaseEvent
	Content *OrderReadyToShipContent `json:"content"`
}

// VariantUpdatedEvent represents a variant_updated webhook event
type VariantUpdatedEvent struct {
	BaseEvent
	Content *VariantUpdatedContent `json:"content"`
}

// SubscriptionMovedOutEvent represents a subscription_moved_out webhook event
type SubscriptionMovedOutEvent struct {
	BaseEvent
	Content *SubscriptionMovedOutContent `json:"content"`
}

// PaymentScheduleSchemeCreatedEvent represents a payment_schedule_scheme_created webhook event
type PaymentScheduleSchemeCreatedEvent struct {
	BaseEvent
	Content *PaymentScheduleSchemeCreatedContent `json:"content"`
}

// BusinessEntityUpdatedEvent represents a business_entity_updated webhook event
type BusinessEntityUpdatedEvent struct {
	BaseEvent
	Content *BusinessEntityUpdatedContent `json:"content"`
}

// SubscriptionScheduledResumptionRemovedEvent represents a subscription_scheduled_resumption_removed webhook event
type SubscriptionScheduledResumptionRemovedEvent struct {
	BaseEvent
	Content *SubscriptionScheduledResumptionRemovedContent `json:"content"`
}

// PaymentInitiatedEvent represents a payment_initiated webhook event
type PaymentInitiatedEvent struct {
	BaseEvent
	Content *PaymentInitiatedContent `json:"content"`
}

// FeatureArchivedEvent represents a feature_archived webhook event
type FeatureArchivedEvent struct {
	BaseEvent
	Content *FeatureArchivedContent `json:"content"`
}

// SubscriptionReactivatedWithBackdatingEvent represents a subscription_reactivated_with_backdating webhook event
type SubscriptionReactivatedWithBackdatingEvent struct {
	BaseEvent
	Content *SubscriptionReactivatedWithBackdatingContent `json:"content"`
}

// OmnichannelSubscriptionImportedEvent represents a omnichannel_subscription_imported webhook event
type OmnichannelSubscriptionImportedEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionImportedContent `json:"content"`
}

// TokenExpiredEvent represents a token_expired webhook event
type TokenExpiredEvent struct {
	BaseEvent
	Content *TokenExpiredContent `json:"content"`
}

// CardAddedEvent represents a card_added webhook event
type CardAddedEvent struct {
	BaseEvent
	Content *CardAddedContent `json:"content"`
}

// CouponCreatedEvent represents a coupon_created webhook event
type CouponCreatedEvent struct {
	BaseEvent
	Content *CouponCreatedContent `json:"content"`
}

// RuleDeletedEvent represents a rule_deleted webhook event
type RuleDeletedEvent struct {
	BaseEvent
	Content *RuleDeletedContent `json:"content"`
}

// ItemPriceEntitlementsUpdatedEvent represents a item_price_entitlements_updated webhook event
type ItemPriceEntitlementsUpdatedEvent struct {
	BaseEvent
	Content *ItemPriceEntitlementsUpdatedContent `json:"content"`
}

// ItemPriceDeletedEvent represents a item_price_deleted webhook event
type ItemPriceDeletedEvent struct {
	BaseEvent
	Content *ItemPriceDeletedContent `json:"content"`
}

// VirtualBankAccountDeletedEvent represents a virtual_bank_account_deleted webhook event
type VirtualBankAccountDeletedEvent struct {
	BaseEvent
	Content *VirtualBankAccountDeletedContent `json:"content"`
}

// PaymentScheduleSchemeDeletedEvent represents a payment_schedule_scheme_deleted webhook event
type PaymentScheduleSchemeDeletedEvent struct {
	BaseEvent
	Content *PaymentScheduleSchemeDeletedContent `json:"content"`
}

// SubscriptionCreatedEvent represents a subscription_created webhook event
type SubscriptionCreatedEvent struct {
	BaseEvent
	Content *SubscriptionCreatedContent `json:"content"`
}

// SubscriptionEntitlementsCreatedEvent represents a subscription_entitlements_created webhook event
type SubscriptionEntitlementsCreatedEvent struct {
	BaseEvent
	Content *SubscriptionEntitlementsCreatedContent `json:"content"`
}

// OrderReturnedEvent represents a order_returned webhook event
type OrderReturnedEvent struct {
	BaseEvent
	Content *OrderReturnedContent `json:"content"`
}

// SubscriptionDeletedEvent represents a subscription_deleted webhook event
type SubscriptionDeletedEvent struct {
	BaseEvent
	Content *SubscriptionDeletedContent `json:"content"`
}

// PaymentSourceAddedEvent represents a payment_source_added webhook event
type PaymentSourceAddedEvent struct {
	BaseEvent
	Content *PaymentSourceAddedContent `json:"content"`
}

// SubscriptionMovedInEvent represents a subscription_moved_in webhook event
type SubscriptionMovedInEvent struct {
	BaseEvent
	Content *SubscriptionMovedInContent `json:"content"`
}

// ItemPriceCreatedEvent represents a item_price_created webhook event
type ItemPriceCreatedEvent struct {
	BaseEvent
	Content *ItemPriceCreatedContent `json:"content"`
}

// SubscriptionScheduledCancellationRemovedEvent represents a subscription_scheduled_cancellation_removed webhook event
type SubscriptionScheduledCancellationRemovedEvent struct {
	BaseEvent
	Content *SubscriptionScheduledCancellationRemovedContent `json:"content"`
}

// PaymentRefundedEvent represents a payment_refunded webhook event
type PaymentRefundedEvent struct {
	BaseEvent
	Content *PaymentRefundedContent `json:"content"`
}

// UsageFileIngestedEvent represents a usage_file_ingested webhook event
type UsageFileIngestedEvent struct {
	BaseEvent
	Content *UsageFileIngestedContent `json:"content"`
}

// ProductCreatedEvent represents a product_created webhook event
type ProductCreatedEvent struct {
	BaseEvent
	Content *ProductCreatedContent `json:"content"`
}

// OmnichannelSubscriptionMovedInEvent represents a omnichannel_subscription_moved_in webhook event
type OmnichannelSubscriptionMovedInEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionMovedInContent `json:"content"`
}

// DifferentialPriceCreatedEvent represents a differential_price_created webhook event
type DifferentialPriceCreatedEvent struct {
	BaseEvent
	Content *DifferentialPriceCreatedContent `json:"content"`
}

// TransactionCreatedEvent represents a transaction_created webhook event
type TransactionCreatedEvent struct {
	BaseEvent
	Content *TransactionCreatedContent `json:"content"`
}

// OmnichannelSubscriptionItemDowngradeScheduledEvent represents a omnichannel_subscription_item_downgrade_scheduled webhook event
type OmnichannelSubscriptionItemDowngradeScheduledEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemDowngradeScheduledContent `json:"content"`
}

// PaymentSucceededEvent represents a payment_succeeded webhook event
type PaymentSucceededEvent struct {
	BaseEvent
	Content *PaymentSucceededContent `json:"content"`
}

// SubscriptionCanceledWithBackdatingEvent represents a subscription_canceled_with_backdating webhook event
type SubscriptionCanceledWithBackdatingEvent struct {
	BaseEvent
	Content *SubscriptionCanceledWithBackdatingContent `json:"content"`
}

// UnbilledChargesVoidedEvent represents a unbilled_charges_voided webhook event
type UnbilledChargesVoidedEvent struct {
	BaseEvent
	Content *UnbilledChargesVoidedContent `json:"content"`
}

// QuoteCreatedEvent represents a quote_created webhook event
type QuoteCreatedEvent struct {
	BaseEvent
	Content *QuoteCreatedContent `json:"content"`
}

// CouponSetDeletedEvent represents a coupon_set_deleted webhook event
type CouponSetDeletedEvent struct {
	BaseEvent
	Content *CouponSetDeletedContent `json:"content"`
}

// AttachedItemCreatedEvent represents a attached_item_created webhook event
type AttachedItemCreatedEvent struct {
	BaseEvent
	Content *AttachedItemCreatedContent `json:"content"`
}

// SalesOrderCreatedEvent represents a sales_order_created webhook event
type SalesOrderCreatedEvent struct {
	BaseEvent
	Content *SalesOrderCreatedContent `json:"content"`
}

// CustomerChangedEvent represents a customer_changed webhook event
type CustomerChangedEvent struct {
	BaseEvent
	Content *CustomerChangedContent `json:"content"`
}

// SubscriptionStartedEvent represents a subscription_started webhook event
type SubscriptionStartedEvent struct {
	BaseEvent
	Content *SubscriptionStartedContent `json:"content"`
}

// SubscriptionActivatedEvent represents a subscription_activated webhook event
type SubscriptionActivatedEvent struct {
	BaseEvent
	Content *SubscriptionActivatedContent `json:"content"`
}

// PaymentSourceExpiringEvent represents a payment_source_expiring webhook event
type PaymentSourceExpiringEvent struct {
	BaseEvent
	Content *PaymentSourceExpiringContent `json:"content"`
}

// SubscriptionReactivatedEvent represents a subscription_reactivated webhook event
type SubscriptionReactivatedEvent struct {
	BaseEvent
	Content *SubscriptionReactivatedContent `json:"content"`
}

// OrderUpdatedEvent represents a order_updated webhook event
type OrderUpdatedEvent struct {
	BaseEvent
	Content *OrderUpdatedContent `json:"content"`
}

// SubscriptionScheduledPauseRemovedEvent represents a subscription_scheduled_pause_removed webhook event
type SubscriptionScheduledPauseRemovedEvent struct {
	BaseEvent
	Content *SubscriptionScheduledPauseRemovedContent `json:"content"`
}

// SubscriptionCancellationReminderEvent represents a subscription_cancellation_reminder webhook event
type SubscriptionCancellationReminderEvent struct {
	BaseEvent
	Content *SubscriptionCancellationReminderContent `json:"content"`
}

// SubscriptionCreatedWithBackdatingEvent represents a subscription_created_with_backdating webhook event
type SubscriptionCreatedWithBackdatingEvent struct {
	BaseEvent
	Content *SubscriptionCreatedWithBackdatingContent `json:"content"`
}

// SubscriptionRampCreatedEvent represents a subscription_ramp_created webhook event
type SubscriptionRampCreatedEvent struct {
	BaseEvent
	Content *SubscriptionRampCreatedContent `json:"content"`
}

// OrderDeletedEvent represents a order_deleted webhook event
type OrderDeletedEvent struct {
	BaseEvent
	Content *OrderDeletedContent `json:"content"`
}

// OmnichannelSubscriptionItemPauseScheduledEvent represents a omnichannel_subscription_item_pause_scheduled webhook event
type OmnichannelSubscriptionItemPauseScheduledEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemPauseScheduledContent `json:"content"`
}

// GiftUpdatedEvent represents a gift_updated webhook event
type GiftUpdatedEvent struct {
	BaseEvent
	Content *GiftUpdatedContent `json:"content"`
}

// SubscriptionTrialExtendedEvent represents a subscription_trial_extended webhook event
type SubscriptionTrialExtendedEvent struct {
	BaseEvent
	Content *SubscriptionTrialExtendedContent `json:"content"`
}

// OmnichannelSubscriptionItemGracePeriodStartedEvent represents a omnichannel_subscription_item_grace_period_started webhook event
type OmnichannelSubscriptionItemGracePeriodStartedEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemGracePeriodStartedContent `json:"content"`
}

// CardExpiryReminderEvent represents a card_expiry_reminder webhook event
type CardExpiryReminderEvent struct {
	BaseEvent
	Content *CardExpiryReminderContent `json:"content"`
}

// TokenCreatedEvent represents a token_created webhook event
type TokenCreatedEvent struct {
	BaseEvent
	Content *TokenCreatedContent `json:"content"`
}

// PromotionalCreditsAddedEvent represents a promotional_credits_added webhook event
type PromotionalCreditsAddedEvent struct {
	BaseEvent
	Content *PromotionalCreditsAddedContent `json:"content"`
}

// SubscriptionRampUpdatedEvent represents a subscription_ramp_updated webhook event
type SubscriptionRampUpdatedEvent struct {
	BaseEvent
	Content *SubscriptionRampUpdatedContent `json:"content"`
}

// CustomerEntitlementsUpdatedEvent represents a customer_entitlements_updated webhook event
type CustomerEntitlementsUpdatedEvent struct {
	BaseEvent
	Content *CustomerEntitlementsUpdatedContent `json:"content"`
}

// PaymentSourceExpiredEvent represents a payment_source_expired webhook event
type PaymentSourceExpiredEvent struct {
	BaseEvent
	Content *PaymentSourceExpiredContent `json:"content"`
}

// CustomerMovedOutEvent represents a customer_moved_out webhook event
type CustomerMovedOutEvent struct {
	BaseEvent
	Content *CustomerMovedOutContent `json:"content"`
}

// SubscriptionEntitlementsUpdatedEvent represents a subscription_entitlements_updated webhook event
type SubscriptionEntitlementsUpdatedEvent struct {
	BaseEvent
	Content *SubscriptionEntitlementsUpdatedContent `json:"content"`
}

// OmnichannelSubscriptionItemDunningExpiredEvent represents a omnichannel_subscription_item_dunning_expired webhook event
type OmnichannelSubscriptionItemDunningExpiredEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemDunningExpiredContent `json:"content"`
}

// HierarchyCreatedEvent represents a hierarchy_created webhook event
type HierarchyCreatedEvent struct {
	BaseEvent
	Content *HierarchyCreatedContent `json:"content"`
}

// AttachedItemDeletedEvent represents a attached_item_deleted webhook event
type AttachedItemDeletedEvent struct {
	BaseEvent
	Content *AttachedItemDeletedContent `json:"content"`
}

// OmnichannelSubscriptionItemScheduledCancellationRemovedEvent represents a omnichannel_subscription_item_scheduled_cancellation_removed webhook event
type OmnichannelSubscriptionItemScheduledCancellationRemovedEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemScheduledCancellationRemovedContent `json:"content"`
}

// ItemUpdatedEvent represents a item_updated webhook event
type ItemUpdatedEvent struct {
	BaseEvent
	Content *ItemUpdatedContent `json:"content"`
}

// CouponSetCreatedEvent represents a coupon_set_created webhook event
type CouponSetCreatedEvent struct {
	BaseEvent
	Content *CouponSetCreatedContent `json:"content"`
}

// PaymentIntentUpdatedEvent represents a payment_intent_updated webhook event
type PaymentIntentUpdatedEvent struct {
	BaseEvent
	Content *PaymentIntentUpdatedContent `json:"content"`
}

// OrderResentEvent represents a order_resent webhook event
type OrderResentEvent struct {
	BaseEvent
	Content *OrderResentContent `json:"content"`
}

// OmnichannelSubscriptionItemScheduledDowngradeRemovedEvent represents a omnichannel_subscription_item_scheduled_downgrade_removed webhook event
type OmnichannelSubscriptionItemScheduledDowngradeRemovedEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemScheduledDowngradeRemovedContent `json:"content"`
}

// OmnichannelSubscriptionCreatedEvent represents a omnichannel_subscription_created webhook event
type OmnichannelSubscriptionCreatedEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionCreatedContent `json:"content"`
}

// TaxWithheldRecordedEvent represents a tax_withheld_recorded webhook event
type TaxWithheldRecordedEvent struct {
	BaseEvent
	Content *TaxWithheldRecordedContent `json:"content"`
}

// PriceVariantCreatedEvent represents a price_variant_created webhook event
type PriceVariantCreatedEvent struct {
	BaseEvent
	Content *PriceVariantCreatedContent `json:"content"`
}

// DifferentialPriceDeletedEvent represents a differential_price_deleted webhook event
type DifferentialPriceDeletedEvent struct {
	BaseEvent
	Content *DifferentialPriceDeletedContent `json:"content"`
}

// SubscriptionItemsRenewedEvent represents a subscription_items_renewed webhook event
type SubscriptionItemsRenewedEvent struct {
	BaseEvent
	Content *SubscriptionItemsRenewedContent `json:"content"`
}

// RuleCreatedEvent represents a rule_created webhook event
type RuleCreatedEvent struct {
	BaseEvent
	Content *RuleCreatedContent `json:"content"`
}

// ContractTermCancelledEvent represents a contract_term_cancelled webhook event
type ContractTermCancelledEvent struct {
	BaseEvent
	Content *ContractTermCancelledContent `json:"content"`
}

// ContractTermRenewedEvent represents a contract_term_renewed webhook event
type ContractTermRenewedEvent struct {
	BaseEvent
	Content *ContractTermRenewedContent `json:"content"`
}

// InvoiceDeletedEvent represents a invoice_deleted webhook event
type InvoiceDeletedEvent struct {
	BaseEvent
	Content *InvoiceDeletedContent `json:"content"`
}

// ItemPriceEntitlementsRemovedEvent represents a item_price_entitlements_removed webhook event
type ItemPriceEntitlementsRemovedEvent struct {
	BaseEvent
	Content *ItemPriceEntitlementsRemovedContent `json:"content"`
}

// SalesOrderUpdatedEvent represents a sales_order_updated webhook event
type SalesOrderUpdatedEvent struct {
	BaseEvent
	Content *SalesOrderUpdatedContent `json:"content"`
}

// OmnichannelSubscriptionItemDunningStartedEvent represents a omnichannel_subscription_item_dunning_started webhook event
type OmnichannelSubscriptionItemDunningStartedEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemDunningStartedContent `json:"content"`
}

// OmnichannelSubscriptionItemChangeScheduledEvent represents a omnichannel_subscription_item_change_scheduled webhook event
type OmnichannelSubscriptionItemChangeScheduledEvent struct {
	BaseEvent
	Content *OmnichannelSubscriptionItemChangeScheduledContent `json:"content"`
}

// PendingInvoiceUpdatedEvent represents a pending_invoice_updated webhook event
type PendingInvoiceUpdatedEvent struct {
	BaseEvent
	Content *PendingInvoiceUpdatedContent `json:"content"`
}

// QuoteUpdatedEvent represents a quote_updated webhook event
type QuoteUpdatedEvent struct {
	BaseEvent
	Content *QuoteUpdatedContent `json:"content"`
}

// AttachedItemUpdatedEvent represents a attached_item_updated webhook event
type AttachedItemUpdatedEvent struct {
	BaseEvent
	Content *AttachedItemUpdatedContent `json:"content"`
}

// PaymentSourceUpdatedEvent represents a payment_source_updated webhook event
type PaymentSourceUpdatedEvent struct {
	BaseEvent
	Content *PaymentSourceUpdatedContent `json:"content"`
}

// BusinessEntityDeletedEvent represents a business_entity_deleted webhook event
type BusinessEntityDeletedEvent struct {
	BaseEvent
	Content *BusinessEntityDeletedContent `json:"content"`
}

// AuthorizationVoidedEvent represents a authorization_voided webhook event
type AuthorizationVoidedEvent struct {
	BaseEvent
	Content *AuthorizationVoidedContent `json:"content"`
}

// SubscriptionRampDeletedEvent represents a subscription_ramp_deleted webhook event
type SubscriptionRampDeletedEvent struct {
	BaseEvent
	Content *SubscriptionRampDeletedContent `json:"content"`
}

// PlanDeletedEvent represents a plan_deleted webhook event
type PlanDeletedEvent struct {
	BaseEvent
	Content *PlanDeletedContent `json:"content"`
}

// AddonDeletedEvent represents a addon_deleted webhook event
type AddonDeletedEvent struct {
	BaseEvent
	Content *AddonDeletedContent `json:"content"`
}

// AddonUpdatedEvent represents a addon_updated webhook event
type AddonUpdatedEvent struct {
	BaseEvent
	Content *AddonUpdatedContent `json:"content"`
}

// AddonCreatedEvent represents a addon_created webhook event
type AddonCreatedEvent struct {
	BaseEvent
	Content *AddonCreatedContent `json:"content"`
}

// PlanCreatedEvent represents a plan_created webhook event
type PlanCreatedEvent struct {
	BaseEvent
	Content *PlanCreatedContent `json:"content"`
}

// PlanUpdatedEvent represents a plan_updated webhook event
type PlanUpdatedEvent struct {
	BaseEvent
	Content *PlanUpdatedContent `json:"content"`
}
