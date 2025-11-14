package webhook

import (
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

// Generated event types for each webhook event

// SubscriptionPauseScheduledEvent represents a subscription_pause_scheduled webhook event
type SubscriptionPauseScheduledEvent struct {
	Id         string                             `json:"id"`
	OccurredAt int64                              `json:"occurred_at"`
	EventType  string                             `json:"event_type"`
	Content    *SubscriptionPauseScheduledContent `json:"content"`
}

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

func (e *GiftExpiredEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// TaxWithheldDeletedEvent represents a tax_withheld_deleted webhook event
type TaxWithheldDeletedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt int64                      `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *TaxWithheldDeletedContent `json:"content"`
}

func (e *TaxWithheldDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// UnbilledChargesDeletedEvent represents a unbilled_charges_deleted webhook event
type UnbilledChargesDeletedEvent struct {
	Id         string                         `json:"id"`
	OccurredAt int64                          `json:"occurred_at"`
	EventType  string                         `json:"event_type"`
	Content    *UnbilledChargesDeletedContent `json:"content"`
}

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

func (e *CouponUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionItemReactivatedEvent represents a omnichannel_subscription_item_reactivated webhook event
type OmnichannelSubscriptionItemReactivatedEvent struct {
	Id         string                                         `json:"id"`
	OccurredAt int64                                          `json:"occurred_at"`
	EventType  string                                         `json:"event_type"`
	Content    *OmnichannelSubscriptionItemReactivatedContent `json:"content"`
}

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

func (e *SubscriptionResumedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelOneTimeOrderItemCancelledEvent represents a omnichannel_one_time_order_item_cancelled webhook event
type OmnichannelOneTimeOrderItemCancelledEvent struct {
	Id         string                                       `json:"id"`
	OccurredAt int64                                        `json:"occurred_at"`
	EventType  string                                       `json:"event_type"`
	Content    *OmnichannelOneTimeOrderItemCancelledContent `json:"content"`
}

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

func (e *SubscriptionCancelledEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ItemEntitlementsRemovedEvent represents a item_entitlements_removed webhook event
type ItemEntitlementsRemovedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *ItemEntitlementsRemovedContent `json:"content"`
}

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

func (e *BusinessEntityCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CouponSetUpdatedEvent represents a coupon_set_updated webhook event
type CouponSetUpdatedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *CouponSetUpdatedContent `json:"content"`
}

func (e *CouponSetUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// DifferentialPriceUpdatedEvent represents a differential_price_updated webhook event
type DifferentialPriceUpdatedEvent struct {
	Id         string                           `json:"id"`
	OccurredAt int64                            `json:"occurred_at"`
	EventType  string                           `json:"event_type"`
	Content    *DifferentialPriceUpdatedContent `json:"content"`
}

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

func (e *VoucherCreateFailedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// GiftClaimedEvent represents a gift_claimed webhook event
type GiftClaimedEvent struct {
	Id         string              `json:"id"`
	OccurredAt int64               `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *GiftClaimedContent `json:"content"`
}

func (e *GiftClaimedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CustomerDeletedEvent represents a customer_deleted webhook event
type CustomerDeletedEvent struct {
	Id         string                  `json:"id"`
	OccurredAt int64                   `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *CustomerDeletedContent `json:"content"`
}

func (e *CustomerDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// RefundInitiatedEvent represents a refund_initiated webhook event
type RefundInitiatedEvent struct {
	Id         string                  `json:"id"`
	OccurredAt int64                   `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *RefundInitiatedContent `json:"content"`
}

func (e *RefundInitiatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// InvoiceGeneratedWithBackdatingEvent represents a invoice_generated_with_backdating webhook event
type InvoiceGeneratedWithBackdatingEvent struct {
	Id         string                                 `json:"id"`
	OccurredAt int64                                  `json:"occurred_at"`
	EventType  string                                 `json:"event_type"`
	Content    *InvoiceGeneratedWithBackdatingContent `json:"content"`
}

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

func (e *AddUsagesReminderEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// VoucherCreatedEvent represents a voucher_created webhook event
type VoucherCreatedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt int64                  `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *VoucherCreatedContent `json:"content"`
}

func (e *VoucherCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// RuleUpdatedEvent represents a rule_updated webhook event
type RuleUpdatedEvent struct {
	Id         string              `json:"id"`
	OccurredAt int64               `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *RuleUpdatedContent `json:"content"`
}

func (e *RuleUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentSchedulesCreatedEvent represents a payment_schedules_created webhook event
type PaymentSchedulesCreatedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *PaymentSchedulesCreatedContent `json:"content"`
}

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

func (e *FeatureActivatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentSourceLocallyDeletedEvent represents a payment_source_locally_deleted webhook event
type PaymentSourceLocallyDeletedEvent struct {
	Id         string                              `json:"id"`
	OccurredAt int64                               `json:"occurred_at"`
	EventType  string                              `json:"event_type"`
	Content    *PaymentSourceLocallyDeletedContent `json:"content"`
}

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

func (e *InvoiceGeneratedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// VoucherExpiredEvent represents a voucher_expired webhook event
type VoucherExpiredEvent struct {
	Id         string                 `json:"id"`
	OccurredAt int64                  `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *VoucherExpiredContent `json:"content"`
}

func (e *VoucherExpiredEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// AuthorizationSucceededEvent represents a authorization_succeeded webhook event
type AuthorizationSucceededEvent struct {
	Id         string                         `json:"id"`
	OccurredAt int64                          `json:"occurred_at"`
	EventType  string                         `json:"event_type"`
	Content    *AuthorizationSucceededContent `json:"content"`
}

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

func (e *GiftScheduledEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionChangesScheduledEvent represents a subscription_changes_scheduled webhook event
type SubscriptionChangesScheduledEvent struct {
	Id         string                               `json:"id"`
	OccurredAt int64                                `json:"occurred_at"`
	EventType  string                               `json:"event_type"`
	Content    *SubscriptionChangesScheduledContent `json:"content"`
}

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

func (e *GiftUnclaimedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// VirtualBankAccountAddedEvent represents a virtual_bank_account_added webhook event
type VirtualBankAccountAddedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *VirtualBankAccountAddedContent `json:"content"`
}

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

func (e *PaymentIntentCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CreditNoteCreatedWithBackdatingEvent represents a credit_note_created_with_backdating webhook event
type CreditNoteCreatedWithBackdatingEvent struct {
	Id         string                                  `json:"id"`
	OccurredAt int64                                   `json:"occurred_at"`
	EventType  string                                  `json:"event_type"`
	Content    *CreditNoteCreatedWithBackdatingContent `json:"content"`
}

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

func (e *ItemFamilyUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OrderCreatedEvent represents a order_created webhook event
type OrderCreatedEvent struct {
	Id         string               `json:"id"`
	OccurredAt int64                `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *OrderCreatedContent `json:"content"`
}

func (e *OrderCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PriceVariantDeletedEvent represents a price_variant_deleted webhook event
type PriceVariantDeletedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *PriceVariantDeletedContent `json:"content"`
}

func (e *PriceVariantDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionMovementFailedEvent represents a subscription_movement_failed webhook event
type SubscriptionMovementFailedEvent struct {
	Id         string                             `json:"id"`
	OccurredAt int64                              `json:"occurred_at"`
	EventType  string                             `json:"event_type"`
	Content    *SubscriptionMovementFailedContent `json:"content"`
}

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

func (e *CustomerMovedInEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionAdvanceInvoiceScheduleUpdatedEvent represents a subscription_advance_invoice_schedule_updated webhook event
type SubscriptionAdvanceInvoiceScheduleUpdatedEvent struct {
	Id         string                                            `json:"id"`
	OccurredAt int64                                             `json:"occurred_at"`
	EventType  string                                            `json:"event_type"`
	Content    *SubscriptionAdvanceInvoiceScheduleUpdatedContent `json:"content"`
}

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

func (e *ItemDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionRampDraftedEvent represents a subscription_ramp_drafted webhook event
type SubscriptionRampDraftedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *SubscriptionRampDraftedContent `json:"content"`
}

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

func (e *DunningUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ItemEntitlementsUpdatedEvent represents a item_entitlements_updated webhook event
type ItemEntitlementsUpdatedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *ItemEntitlementsUpdatedContent `json:"content"`
}

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

func (e *TokenConsumedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// HierarchyDeletedEvent represents a hierarchy_deleted webhook event
type HierarchyDeletedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *HierarchyDeletedContent `json:"content"`
}

func (e *HierarchyDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionCancellationScheduledEvent represents a subscription_cancellation_scheduled webhook event
type SubscriptionCancellationScheduledEvent struct {
	Id         string                                    `json:"id"`
	OccurredAt int64                                     `json:"occurred_at"`
	EventType  string                                    `json:"event_type"`
	Content    *SubscriptionCancellationScheduledContent `json:"content"`
}

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

func (e *SubscriptionRenewedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// FeatureUpdatedEvent represents a feature_updated webhook event
type FeatureUpdatedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt int64                  `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *FeatureUpdatedContent `json:"content"`
}

func (e *FeatureUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// FeatureDeletedEvent represents a feature_deleted webhook event
type FeatureDeletedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt int64                  `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *FeatureDeletedContent `json:"content"`
}

func (e *FeatureDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ItemFamilyCreatedEvent represents a item_family_created webhook event
type ItemFamilyCreatedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt int64                     `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *ItemFamilyCreatedContent `json:"content"`
}

func (e *ItemFamilyCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionItemScheduledChangeRemovedEvent represents a omnichannel_subscription_item_scheduled_change_removed webhook event
type OmnichannelSubscriptionItemScheduledChangeRemovedEvent struct {
	Id         string                                                    `json:"id"`
	OccurredAt int64                                                     `json:"occurred_at"`
	EventType  string                                                    `json:"event_type"`
	Content    *OmnichannelSubscriptionItemScheduledChangeRemovedContent `json:"content"`
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

func (e *PurchaseCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// EntitlementOverridesUpdatedEvent represents a entitlement_overrides_updated webhook event
type EntitlementOverridesUpdatedEvent struct {
	Id         string                              `json:"id"`
	OccurredAt int64                               `json:"occurred_at"`
	EventType  string                              `json:"event_type"`
	Content    *EntitlementOverridesUpdatedContent `json:"content"`
}

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

func (e *ItemFamilyDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionResumptionScheduledEvent represents a subscription_resumption_scheduled webhook event
type SubscriptionResumptionScheduledEvent struct {
	Id         string                                  `json:"id"`
	OccurredAt int64                                   `json:"occurred_at"`
	EventType  string                                  `json:"event_type"`
	Content    *SubscriptionResumptionScheduledContent `json:"content"`
}

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

func (e *FeatureReactivatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CouponCodesDeletedEvent represents a coupon_codes_deleted webhook event
type CouponCodesDeletedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt int64                      `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *CouponCodesDeletedContent `json:"content"`
}

func (e *CouponCodesDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CardExpiredEvent represents a card_expired webhook event
type CardExpiredEvent struct {
	Id         string              `json:"id"`
	OccurredAt int64               `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *CardExpiredContent `json:"content"`
}

func (e *CardExpiredEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CreditNoteUpdatedEvent represents a credit_note_updated webhook event
type CreditNoteUpdatedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt int64                     `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *CreditNoteUpdatedContent `json:"content"`
}

func (e *CreditNoteUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionItemDowngradedEvent represents a omnichannel_subscription_item_downgraded webhook event
type OmnichannelSubscriptionItemDowngradedEvent struct {
	Id         string                                        `json:"id"`
	OccurredAt int64                                         `json:"occurred_at"`
	EventType  string                                        `json:"event_type"`
	Content    *OmnichannelSubscriptionItemDowngradedContent `json:"content"`
}

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

func (e *PriceVariantUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PromotionalCreditsDeductedEvent represents a promotional_credits_deducted webhook event
type PromotionalCreditsDeductedEvent struct {
	Id         string                             `json:"id"`
	OccurredAt int64                              `json:"occurred_at"`
	EventType  string                             `json:"event_type"`
	Content    *PromotionalCreditsDeductedContent `json:"content"`
}

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

func (e *SubscriptionPausedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OrderReadyToProcessEvent represents a order_ready_to_process webhook event
type OrderReadyToProcessEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *OrderReadyToProcessContent `json:"content"`
}

func (e *OrderReadyToProcessEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// FeatureCreatedEvent represents a feature_created webhook event
type FeatureCreatedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt int64                  `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *FeatureCreatedContent `json:"content"`
}

func (e *FeatureCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// TransactionDeletedEvent represents a transaction_deleted webhook event
type TransactionDeletedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt int64                      `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *TransactionDeletedContent `json:"content"`
}

func (e *TransactionDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CreditNoteCreatedEvent represents a credit_note_created webhook event
type CreditNoteCreatedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt int64                     `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *CreditNoteCreatedContent `json:"content"`
}

func (e *CreditNoteCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionItemResubscribedEvent represents a omnichannel_subscription_item_resubscribed webhook event
type OmnichannelSubscriptionItemResubscribedEvent struct {
	Id         string                                          `json:"id"`
	OccurredAt int64                                           `json:"occurred_at"`
	EventType  string                                          `json:"event_type"`
	Content    *OmnichannelSubscriptionItemResubscribedContent `json:"content"`
}

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

func (e *RecordPurchaseFailedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ItemCreatedEvent represents a item_created webhook event
type ItemCreatedEvent struct {
	Id         string              `json:"id"`
	OccurredAt int64               `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *ItemCreatedContent `json:"content"`
}

func (e *ItemCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// TransactionUpdatedEvent represents a transaction_updated webhook event
type TransactionUpdatedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt int64                      `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *TransactionUpdatedContent `json:"content"`
}

func (e *TransactionUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// MrrUpdatedEvent represents a mrr_updated webhook event
type MrrUpdatedEvent struct {
	Id         string             `json:"id"`
	OccurredAt int64              `json:"occurred_at"`
	EventType  string             `json:"event_type"`
	Content    *MrrUpdatedContent `json:"content"`
}

func (e *MrrUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// UnbilledChargesInvoicedEvent represents a unbilled_charges_invoiced webhook event
type UnbilledChargesInvoicedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *UnbilledChargesInvoicedContent `json:"content"`
}

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

func (e *ItemPriceUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CouponCodesUpdatedEvent represents a coupon_codes_updated webhook event
type CouponCodesUpdatedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt int64                      `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *CouponCodesUpdatedContent `json:"content"`
}

func (e *CouponCodesUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// VirtualBankAccountUpdatedEvent represents a virtual_bank_account_updated webhook event
type VirtualBankAccountUpdatedEvent struct {
	Id         string                            `json:"id"`
	OccurredAt int64                             `json:"occurred_at"`
	EventType  string                            `json:"event_type"`
	Content    *VirtualBankAccountUpdatedContent `json:"content"`
}

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

func (e *ContractTermCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionChangedEvent represents a subscription_changed webhook event
type SubscriptionChangedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *SubscriptionChangedContent `json:"content"`
}

func (e *SubscriptionChangedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentFailedEvent represents a payment_failed webhook event
type PaymentFailedEvent struct {
	Id         string                `json:"id"`
	OccurredAt int64                 `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *PaymentFailedContent `json:"content"`
}

func (e *PaymentFailedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CreditNoteDeletedEvent represents a credit_note_deleted webhook event
type CreditNoteDeletedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt int64                     `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *CreditNoteDeletedContent `json:"content"`
}

func (e *CreditNoteDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// TaxWithheldRefundedEvent represents a tax_withheld_refunded webhook event
type TaxWithheldRefundedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *TaxWithheldRefundedContent `json:"content"`
}

func (e *TaxWithheldRefundedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ContractTermCompletedEvent represents a contract_term_completed webhook event
type ContractTermCompletedEvent struct {
	Id         string                        `json:"id"`
	OccurredAt int64                         `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *ContractTermCompletedContent `json:"content"`
}

func (e *ContractTermCompletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentSchedulesUpdatedEvent represents a payment_schedules_updated webhook event
type PaymentSchedulesUpdatedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *PaymentSchedulesUpdatedContent `json:"content"`
}

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

func (e *CardUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CustomerCreatedEvent represents a customer_created webhook event
type CustomerCreatedEvent struct {
	Id         string                  `json:"id"`
	OccurredAt int64                   `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *CustomerCreatedContent `json:"content"`
}

func (e *CustomerCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionRenewalReminderEvent represents a subscription_renewal_reminder webhook event
type SubscriptionRenewalReminderEvent struct {
	Id         string                              `json:"id"`
	OccurredAt int64                               `json:"occurred_at"`
	EventType  string                              `json:"event_type"`
	Content    *SubscriptionRenewalReminderContent `json:"content"`
}

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

func (e *OrderDeliveredEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionItemCancellationScheduledEvent represents a omnichannel_subscription_item_cancellation_scheduled webhook event
type OmnichannelSubscriptionItemCancellationScheduledEvent struct {
	Id         string                                                   `json:"id"`
	OccurredAt int64                                                    `json:"occurred_at"`
	EventType  string                                                   `json:"event_type"`
	Content    *OmnichannelSubscriptionItemCancellationScheduledContent `json:"content"`
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

func (e *CouponCodesAddedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// GiftCancelledEvent represents a gift_cancelled webhook event
type GiftCancelledEvent struct {
	Id         string                `json:"id"`
	OccurredAt int64                 `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *GiftCancelledContent `json:"content"`
}

func (e *GiftCancelledEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OrderCancelledEvent represents a order_cancelled webhook event
type OrderCancelledEvent struct {
	Id         string                 `json:"id"`
	OccurredAt int64                  `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *OrderCancelledContent `json:"content"`
}

func (e *OrderCancelledEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CouponDeletedEvent represents a coupon_deleted webhook event
type CouponDeletedEvent struct {
	Id         string                `json:"id"`
	OccurredAt int64                 `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *CouponDeletedContent `json:"content"`
}

func (e *CouponDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionScheduledChangesRemovedEvent represents a subscription_scheduled_changes_removed webhook event
type SubscriptionScheduledChangesRemovedEvent struct {
	Id         string                                      `json:"id"`
	OccurredAt int64                                       `json:"occurred_at"`
	EventType  string                                      `json:"event_type"`
	Content    *SubscriptionScheduledChangesRemovedContent `json:"content"`
}

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

func (e *PendingInvoiceCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// EntitlementOverridesAutoRemovedEvent represents a entitlement_overrides_auto_removed webhook event
type EntitlementOverridesAutoRemovedEvent struct {
	Id         string                                  `json:"id"`
	OccurredAt int64                                   `json:"occurred_at"`
	EventType  string                                  `json:"event_type"`
	Content    *EntitlementOverridesAutoRemovedContent `json:"content"`
}

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

func (e *PaymentSourceDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionItemCancelledEvent represents a omnichannel_subscription_item_cancelled webhook event
type OmnichannelSubscriptionItemCancelledEvent struct {
	Id         string                                       `json:"id"`
	OccurredAt int64                                        `json:"occurred_at"`
	EventType  string                                       `json:"event_type"`
	Content    *OmnichannelSubscriptionItemCancelledContent `json:"content"`
}

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

func (e *QuoteDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// InvoiceUpdatedEvent represents a invoice_updated webhook event
type InvoiceUpdatedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt int64                  `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *InvoiceUpdatedContent `json:"content"`
}

func (e *InvoiceUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionAdvanceInvoiceScheduleRemovedEvent represents a subscription_advance_invoice_schedule_removed webhook event
type SubscriptionAdvanceInvoiceScheduleRemovedEvent struct {
	Id         string                                            `json:"id"`
	OccurredAt int64                                             `json:"occurred_at"`
	EventType  string                                            `json:"event_type"`
	Content    *SubscriptionAdvanceInvoiceScheduleRemovedContent `json:"content"`
}

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

func (e *CardDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OrderReadyToShipEvent represents a order_ready_to_ship webhook event
type OrderReadyToShipEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *OrderReadyToShipContent `json:"content"`
}

func (e *OrderReadyToShipEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionMovedOutEvent represents a subscription_moved_out webhook event
type SubscriptionMovedOutEvent struct {
	Id         string                       `json:"id"`
	OccurredAt int64                        `json:"occurred_at"`
	EventType  string                       `json:"event_type"`
	Content    *SubscriptionMovedOutContent `json:"content"`
}

func (e *SubscriptionMovedOutEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentScheduleSchemeCreatedEvent represents a payment_schedule_scheme_created webhook event
type PaymentScheduleSchemeCreatedEvent struct {
	Id         string                               `json:"id"`
	OccurredAt int64                                `json:"occurred_at"`
	EventType  string                               `json:"event_type"`
	Content    *PaymentScheduleSchemeCreatedContent `json:"content"`
}

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

func (e *BusinessEntityUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionScheduledResumptionRemovedEvent represents a subscription_scheduled_resumption_removed webhook event
type SubscriptionScheduledResumptionRemovedEvent struct {
	Id         string                                         `json:"id"`
	OccurredAt int64                                          `json:"occurred_at"`
	EventType  string                                         `json:"event_type"`
	Content    *SubscriptionScheduledResumptionRemovedContent `json:"content"`
}

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

func (e *PaymentInitiatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// FeatureArchivedEvent represents a feature_archived webhook event
type FeatureArchivedEvent struct {
	Id         string                  `json:"id"`
	OccurredAt int64                   `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *FeatureArchivedContent `json:"content"`
}

func (e *FeatureArchivedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionReactivatedWithBackdatingEvent represents a subscription_reactivated_with_backdating webhook event
type SubscriptionReactivatedWithBackdatingEvent struct {
	Id         string                                        `json:"id"`
	OccurredAt int64                                         `json:"occurred_at"`
	EventType  string                                        `json:"event_type"`
	Content    *SubscriptionReactivatedWithBackdatingContent `json:"content"`
}

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

func (e *TokenExpiredEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CardAddedEvent represents a card_added webhook event
type CardAddedEvent struct {
	Id         string            `json:"id"`
	OccurredAt int64             `json:"occurred_at"`
	EventType  string            `json:"event_type"`
	Content    *CardAddedContent `json:"content"`
}

func (e *CardAddedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CouponCreatedEvent represents a coupon_created webhook event
type CouponCreatedEvent struct {
	Id         string                `json:"id"`
	OccurredAt int64                 `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *CouponCreatedContent `json:"content"`
}

func (e *CouponCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// RuleDeletedEvent represents a rule_deleted webhook event
type RuleDeletedEvent struct {
	Id         string              `json:"id"`
	OccurredAt int64               `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *RuleDeletedContent `json:"content"`
}

func (e *RuleDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ItemPriceEntitlementsUpdatedEvent represents a item_price_entitlements_updated webhook event
type ItemPriceEntitlementsUpdatedEvent struct {
	Id         string                               `json:"id"`
	OccurredAt int64                                `json:"occurred_at"`
	EventType  string                               `json:"event_type"`
	Content    *ItemPriceEntitlementsUpdatedContent `json:"content"`
}

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

func (e *ItemPriceDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// VirtualBankAccountDeletedEvent represents a virtual_bank_account_deleted webhook event
type VirtualBankAccountDeletedEvent struct {
	Id         string                            `json:"id"`
	OccurredAt int64                             `json:"occurred_at"`
	EventType  string                            `json:"event_type"`
	Content    *VirtualBankAccountDeletedContent `json:"content"`
}

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

func (e *SubscriptionCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionEntitlementsCreatedEvent represents a subscription_entitlements_created webhook event
type SubscriptionEntitlementsCreatedEvent struct {
	Id         string                                  `json:"id"`
	OccurredAt int64                                   `json:"occurred_at"`
	EventType  string                                  `json:"event_type"`
	Content    *SubscriptionEntitlementsCreatedContent `json:"content"`
}

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

func (e *OrderReturnedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionDeletedEvent represents a subscription_deleted webhook event
type SubscriptionDeletedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *SubscriptionDeletedContent `json:"content"`
}

func (e *SubscriptionDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentSourceAddedEvent represents a payment_source_added webhook event
type PaymentSourceAddedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt int64                      `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *PaymentSourceAddedContent `json:"content"`
}

func (e *PaymentSourceAddedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionMovedInEvent represents a subscription_moved_in webhook event
type SubscriptionMovedInEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *SubscriptionMovedInContent `json:"content"`
}

func (e *SubscriptionMovedInEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ItemPriceCreatedEvent represents a item_price_created webhook event
type ItemPriceCreatedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *ItemPriceCreatedContent `json:"content"`
}

func (e *ItemPriceCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionScheduledCancellationRemovedEvent represents a subscription_scheduled_cancellation_removed webhook event
type SubscriptionScheduledCancellationRemovedEvent struct {
	Id         string                                           `json:"id"`
	OccurredAt int64                                            `json:"occurred_at"`
	EventType  string                                           `json:"event_type"`
	Content    *SubscriptionScheduledCancellationRemovedContent `json:"content"`
}

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

func (e *PaymentRefundedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// UsageFileIngestedEvent represents a usage_file_ingested webhook event
type UsageFileIngestedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt int64                     `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *UsageFileIngestedContent `json:"content"`
}

func (e *UsageFileIngestedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionMovedInEvent represents a omnichannel_subscription_moved_in webhook event
type OmnichannelSubscriptionMovedInEvent struct {
	Id         string                                 `json:"id"`
	OccurredAt int64                                  `json:"occurred_at"`
	EventType  string                                 `json:"event_type"`
	Content    *OmnichannelSubscriptionMovedInContent `json:"content"`
}

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

func (e *TransactionCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentSucceededEvent represents a payment_succeeded webhook event
type PaymentSucceededEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *PaymentSucceededContent `json:"content"`
}

func (e *PaymentSucceededEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionCanceledWithBackdatingEvent represents a subscription_canceled_with_backdating webhook event
type SubscriptionCanceledWithBackdatingEvent struct {
	Id         string                                     `json:"id"`
	OccurredAt int64                                      `json:"occurred_at"`
	EventType  string                                     `json:"event_type"`
	Content    *SubscriptionCanceledWithBackdatingContent `json:"content"`
}

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

func (e *UnbilledChargesVoidedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// QuoteCreatedEvent represents a quote_created webhook event
type QuoteCreatedEvent struct {
	Id         string               `json:"id"`
	OccurredAt int64                `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *QuoteCreatedContent `json:"content"`
}

func (e *QuoteCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CouponSetDeletedEvent represents a coupon_set_deleted webhook event
type CouponSetDeletedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *CouponSetDeletedContent `json:"content"`
}

func (e *CouponSetDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// AttachedItemCreatedEvent represents a attached_item_created webhook event
type AttachedItemCreatedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *AttachedItemCreatedContent `json:"content"`
}

func (e *AttachedItemCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SalesOrderCreatedEvent represents a sales_order_created webhook event
type SalesOrderCreatedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt int64                     `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *SalesOrderCreatedContent `json:"content"`
}

func (e *SalesOrderCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CustomerChangedEvent represents a customer_changed webhook event
type CustomerChangedEvent struct {
	Id         string                  `json:"id"`
	OccurredAt int64                   `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *CustomerChangedContent `json:"content"`
}

func (e *CustomerChangedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionStartedEvent represents a subscription_started webhook event
type SubscriptionStartedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *SubscriptionStartedContent `json:"content"`
}

func (e *SubscriptionStartedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionActivatedEvent represents a subscription_activated webhook event
type SubscriptionActivatedEvent struct {
	Id         string                        `json:"id"`
	OccurredAt int64                         `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *SubscriptionActivatedContent `json:"content"`
}

func (e *SubscriptionActivatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentSourceExpiringEvent represents a payment_source_expiring webhook event
type PaymentSourceExpiringEvent struct {
	Id         string                        `json:"id"`
	OccurredAt int64                         `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *PaymentSourceExpiringContent `json:"content"`
}

func (e *PaymentSourceExpiringEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionReactivatedEvent represents a subscription_reactivated webhook event
type SubscriptionReactivatedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *SubscriptionReactivatedContent `json:"content"`
}

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

func (e *OrderUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionScheduledPauseRemovedEvent represents a subscription_scheduled_pause_removed webhook event
type SubscriptionScheduledPauseRemovedEvent struct {
	Id         string                                    `json:"id"`
	OccurredAt int64                                     `json:"occurred_at"`
	EventType  string                                    `json:"event_type"`
	Content    *SubscriptionScheduledPauseRemovedContent `json:"content"`
}

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

func (e *OrderDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionItemPauseScheduledEvent represents a omnichannel_subscription_item_pause_scheduled webhook event
type OmnichannelSubscriptionItemPauseScheduledEvent struct {
	Id         string                                            `json:"id"`
	OccurredAt int64                                             `json:"occurred_at"`
	EventType  string                                            `json:"event_type"`
	Content    *OmnichannelSubscriptionItemPauseScheduledContent `json:"content"`
}

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

func (e *GiftUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionTrialExtendedEvent represents a subscription_trial_extended webhook event
type SubscriptionTrialExtendedEvent struct {
	Id         string                            `json:"id"`
	OccurredAt int64                             `json:"occurred_at"`
	EventType  string                            `json:"event_type"`
	Content    *SubscriptionTrialExtendedContent `json:"content"`
}

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

func (e *CardExpiryReminderEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// TokenCreatedEvent represents a token_created webhook event
type TokenCreatedEvent struct {
	Id         string               `json:"id"`
	OccurredAt int64                `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *TokenCreatedContent `json:"content"`
}

func (e *TokenCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PromotionalCreditsAddedEvent represents a promotional_credits_added webhook event
type PromotionalCreditsAddedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *PromotionalCreditsAddedContent `json:"content"`
}

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

func (e *PaymentSourceExpiredEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CustomerMovedOutEvent represents a customer_moved_out webhook event
type CustomerMovedOutEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *CustomerMovedOutContent `json:"content"`
}

func (e *CustomerMovedOutEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionEntitlementsUpdatedEvent represents a subscription_entitlements_updated webhook event
type SubscriptionEntitlementsUpdatedEvent struct {
	Id         string                                  `json:"id"`
	OccurredAt int64                                   `json:"occurred_at"`
	EventType  string                                  `json:"event_type"`
	Content    *SubscriptionEntitlementsUpdatedContent `json:"content"`
}

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

func (e *HierarchyCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// AttachedItemDeletedEvent represents a attached_item_deleted webhook event
type AttachedItemDeletedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *AttachedItemDeletedContent `json:"content"`
}

func (e *AttachedItemDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionItemScheduledCancellationRemovedEvent represents a omnichannel_subscription_item_scheduled_cancellation_removed webhook event
type OmnichannelSubscriptionItemScheduledCancellationRemovedEvent struct {
	Id         string                                                          `json:"id"`
	OccurredAt int64                                                           `json:"occurred_at"`
	EventType  string                                                          `json:"event_type"`
	Content    *OmnichannelSubscriptionItemScheduledCancellationRemovedContent `json:"content"`
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

func (e *ItemUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// CouponSetCreatedEvent represents a coupon_set_created webhook event
type CouponSetCreatedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt int64                    `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *CouponSetCreatedContent `json:"content"`
}

func (e *CouponSetCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentIntentUpdatedEvent represents a payment_intent_updated webhook event
type PaymentIntentUpdatedEvent struct {
	Id         string                       `json:"id"`
	OccurredAt int64                        `json:"occurred_at"`
	EventType  string                       `json:"event_type"`
	Content    *PaymentIntentUpdatedContent `json:"content"`
}

func (e *PaymentIntentUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OrderResentEvent represents a order_resent webhook event
type OrderResentEvent struct {
	Id         string              `json:"id"`
	OccurredAt int64               `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *OrderResentContent `json:"content"`
}

func (e *OrderResentEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionCreatedEvent represents a omnichannel_subscription_created webhook event
type OmnichannelSubscriptionCreatedEvent struct {
	Id         string                                 `json:"id"`
	OccurredAt int64                                  `json:"occurred_at"`
	EventType  string                                 `json:"event_type"`
	Content    *OmnichannelSubscriptionCreatedContent `json:"content"`
}

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

func (e *TaxWithheldRecordedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PriceVariantCreatedEvent represents a price_variant_created webhook event
type PriceVariantCreatedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *PriceVariantCreatedContent `json:"content"`
}

func (e *PriceVariantCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// DifferentialPriceDeletedEvent represents a differential_price_deleted webhook event
type DifferentialPriceDeletedEvent struct {
	Id         string                           `json:"id"`
	OccurredAt int64                            `json:"occurred_at"`
	EventType  string                           `json:"event_type"`
	Content    *DifferentialPriceDeletedContent `json:"content"`
}

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

func (e *RuleCreatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ContractTermCancelledEvent represents a contract_term_cancelled webhook event
type ContractTermCancelledEvent struct {
	Id         string                        `json:"id"`
	OccurredAt int64                         `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *ContractTermCancelledContent `json:"content"`
}

func (e *ContractTermCancelledEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ContractTermRenewedEvent represents a contract_term_renewed webhook event
type ContractTermRenewedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *ContractTermRenewedContent `json:"content"`
}

func (e *ContractTermRenewedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// InvoiceDeletedEvent represents a invoice_deleted webhook event
type InvoiceDeletedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt int64                  `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *InvoiceDeletedContent `json:"content"`
}

func (e *InvoiceDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// ItemPriceEntitlementsRemovedEvent represents a item_price_entitlements_removed webhook event
type ItemPriceEntitlementsRemovedEvent struct {
	Id         string                               `json:"id"`
	OccurredAt int64                                `json:"occurred_at"`
	EventType  string                               `json:"event_type"`
	Content    *ItemPriceEntitlementsRemovedContent `json:"content"`
}

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

func (e *SalesOrderUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// OmnichannelSubscriptionItemDunningStartedEvent represents a omnichannel_subscription_item_dunning_started webhook event
type OmnichannelSubscriptionItemDunningStartedEvent struct {
	Id         string                                            `json:"id"`
	OccurredAt int64                                             `json:"occurred_at"`
	EventType  string                                            `json:"event_type"`
	Content    *OmnichannelSubscriptionItemDunningStartedContent `json:"content"`
}

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

func (e *PendingInvoiceUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// QuoteUpdatedEvent represents a quote_updated webhook event
type QuoteUpdatedEvent struct {
	Id         string               `json:"id"`
	OccurredAt int64                `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *QuoteUpdatedContent `json:"content"`
}

func (e *QuoteUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// AttachedItemUpdatedEvent represents a attached_item_updated webhook event
type AttachedItemUpdatedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *AttachedItemUpdatedContent `json:"content"`
}

func (e *AttachedItemUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// PaymentSourceUpdatedEvent represents a payment_source_updated webhook event
type PaymentSourceUpdatedEvent struct {
	Id         string                       `json:"id"`
	OccurredAt int64                        `json:"occurred_at"`
	EventType  string                       `json:"event_type"`
	Content    *PaymentSourceUpdatedContent `json:"content"`
}

func (e *PaymentSourceUpdatedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// BusinessEntityDeletedEvent represents a business_entity_deleted webhook event
type BusinessEntityDeletedEvent struct {
	Id         string                        `json:"id"`
	OccurredAt int64                         `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *BusinessEntityDeletedContent `json:"content"`
}

func (e *BusinessEntityDeletedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// AuthorizationVoidedEvent represents a authorization_voided webhook event
type AuthorizationVoidedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt int64                       `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *AuthorizationVoidedContent `json:"content"`
}

func (e *AuthorizationVoidedEvent) GetOccurredAtTime() time.Time { return time.Unix(e.OccurredAt, 0) }

// SubscriptionRampDeletedEvent represents a subscription_ramp_deleted webhook event
type SubscriptionRampDeletedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt int64                           `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *SubscriptionRampDeletedContent `json:"content"`
}

func (e *SubscriptionRampDeletedEvent) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}
