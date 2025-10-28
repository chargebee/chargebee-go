package chargebee

import (
	"github.com/chargebee/chargebee-go/v3/models/addon"
	"github.com/chargebee/chargebee-go/v3/models/address"
	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"
	"github.com/chargebee/chargebee-go/v3/models/attacheditem"
	"github.com/chargebee/chargebee-go/v3/models/attribute"
	"github.com/chargebee/chargebee-go/v3/models/billingconfiguration"
	"github.com/chargebee/chargebee-go/v3/models/brand"
	"github.com/chargebee/chargebee-go/v3/models/businessentity"
	"github.com/chargebee/chargebee-go/v3/models/businessentitytransfer"
	"github.com/chargebee/chargebee-go/v3/models/card"
	"github.com/chargebee/chargebee-go/v3/models/comment"
	"github.com/chargebee/chargebee-go/v3/models/configuration"
	"github.com/chargebee/chargebee-go/v3/models/contact"
	"github.com/chargebee/chargebee-go/v3/models/contractterm"
	"github.com/chargebee/chargebee-go/v3/models/coupon"
	"github.com/chargebee/chargebee-go/v3/models/couponcode"
	"github.com/chargebee/chargebee-go/v3/models/couponset"
	"github.com/chargebee/chargebee-go/v3/models/creditnote"
	"github.com/chargebee/chargebee-go/v3/models/currency"
	"github.com/chargebee/chargebee-go/v3/models/customer"
	"github.com/chargebee/chargebee-go/v3/models/customerentitlement"
	"github.com/chargebee/chargebee-go/v3/models/differentialprice"
	"github.com/chargebee/chargebee-go/v3/models/discount"
	"github.com/chargebee/chargebee-go/v3/models/download"
	"github.com/chargebee/chargebee-go/v3/models/entitlement"
	"github.com/chargebee/chargebee-go/v3/models/entitlementoverride"
	"github.com/chargebee/chargebee-go/v3/models/estimate"
	"github.com/chargebee/chargebee-go/v3/models/event"
	"github.com/chargebee/chargebee-go/v3/models/export"
	"github.com/chargebee/chargebee-go/v3/models/feature"
	"github.com/chargebee/chargebee-go/v3/models/gatewayerrordetail"
	"github.com/chargebee/chargebee-go/v3/models/gift"
	"github.com/chargebee/chargebee-go/v3/models/hierarchy"
	"github.com/chargebee/chargebee-go/v3/models/hostedpage"
	"github.com/chargebee/chargebee-go/v3/models/impactedcustomer"
	"github.com/chargebee/chargebee-go/v3/models/impacteditem"
	"github.com/chargebee/chargebee-go/v3/models/impacteditemprice"
	"github.com/chargebee/chargebee-go/v3/models/impactedsubscription"
	"github.com/chargebee/chargebee-go/v3/models/inappsubscription"
	"github.com/chargebee/chargebee-go/v3/models/invoice"
	"github.com/chargebee/chargebee-go/v3/models/item"
	"github.com/chargebee/chargebee-go/v3/models/itementitlement"
	"github.com/chargebee/chargebee-go/v3/models/itemfamily"
	"github.com/chargebee/chargebee-go/v3/models/itemprice"
	"github.com/chargebee/chargebee-go/v3/models/metadata"
	"github.com/chargebee/chargebee-go/v3/models/offerevent"
	"github.com/chargebee/chargebee-go/v3/models/offerfulfillment"
	"github.com/chargebee/chargebee-go/v3/models/omnichannelonetimeorder"
	"github.com/chargebee/chargebee-go/v3/models/omnichannelonetimeorderitem"
	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"
	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"
	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitemscheduledchange"
	"github.com/chargebee/chargebee-go/v3/models/omnichanneltransaction"
	"github.com/chargebee/chargebee-go/v3/models/order"
	"github.com/chargebee/chargebee-go/v3/models/paymentintent"
	"github.com/chargebee/chargebee-go/v3/models/paymentreferencenumber"
	"github.com/chargebee/chargebee-go/v3/models/paymentschedule"
	"github.com/chargebee/chargebee-go/v3/models/paymentschedulescheme"
	"github.com/chargebee/chargebee-go/v3/models/paymentsource"
	"github.com/chargebee/chargebee-go/v3/models/paymentvoucher"
	"github.com/chargebee/chargebee-go/v3/models/personalizedoffer"
	"github.com/chargebee/chargebee-go/v3/models/plan"
	"github.com/chargebee/chargebee-go/v3/models/portalsession"
	"github.com/chargebee/chargebee-go/v3/models/pricevariant"
	"github.com/chargebee/chargebee-go/v3/models/pricingpagesession"
	"github.com/chargebee/chargebee-go/v3/models/promotionalcredit"
	"github.com/chargebee/chargebee-go/v3/models/purchase"
	"github.com/chargebee/chargebee-go/v3/models/quote"
	"github.com/chargebee/chargebee-go/v3/models/quotedcharge"
	"github.com/chargebee/chargebee-go/v3/models/quotedramp"
	"github.com/chargebee/chargebee-go/v3/models/quotedsubscription"
	"github.com/chargebee/chargebee-go/v3/models/quotelinegroup"
	"github.com/chargebee/chargebee-go/v3/models/ramp"
	"github.com/chargebee/chargebee-go/v3/models/recordedpurchase"
	"github.com/chargebee/chargebee-go/v3/models/resourcemigration"
	"github.com/chargebee/chargebee-go/v3/models/rule"
	"github.com/chargebee/chargebee-go/v3/models/sitemigrationdetail"
	"github.com/chargebee/chargebee-go/v3/models/subscription"
	"github.com/chargebee/chargebee-go/v3/models/subscriptionentitlement"
	"github.com/chargebee/chargebee-go/v3/models/subscriptionentitlementscreateddetail"
	"github.com/chargebee/chargebee-go/v3/models/subscriptionentitlementsupdateddetail"
	"github.com/chargebee/chargebee-go/v3/models/taxwithheld"
	"github.com/chargebee/chargebee-go/v3/models/thirdpartypaymentmethod"
	"github.com/chargebee/chargebee-go/v3/models/timemachine"
	"github.com/chargebee/chargebee-go/v3/models/token"
	"github.com/chargebee/chargebee-go/v3/models/transaction"
	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"
	"github.com/chargebee/chargebee-go/v3/models/usage"
	"github.com/chargebee/chargebee-go/v3/models/usageevent"
	"github.com/chargebee/chargebee-go/v3/models/usagefile"
	"github.com/chargebee/chargebee-go/v3/models/virtualbankaccount"
	"github.com/chargebee/chargebee-go/v3/models/webhookendpoint"
	"net/http"
	"strconv"
)

type ResultList struct {
	List            []*Result `json:"list"`
	NextOffset      string    `json:"next_offset"`
	responseHeaders http.Header
	httpStatusCode  int
}
type Result struct {
	Subscription                               *subscription.Subscription                                                             `json:"subscription,omitempty"`
	ContractTerm                               *contractterm.ContractTerm                                                             `json:"contract_term,omitempty"`
	Discount                                   *discount.Discount                                                                     `json:"discount,omitempty"`
	AdvanceInvoiceSchedule                     *advanceinvoiceschedule.AdvanceInvoiceSchedule                                         `json:"advance_invoice_schedule,omitempty"`
	Customer                                   *customer.Customer                                                                     `json:"customer,omitempty"`
	Hierarchy                                  *hierarchy.Hierarchy                                                                   `json:"hierarchy,omitempty"`
	Contact                                    *contact.Contact                                                                       `json:"contact,omitempty"`
	BusinessEntityTransfer                     *businessentitytransfer.BusinessEntityTransfer                                         `json:"business_entity_transfer,omitempty"`
	Token                                      *token.Token                                                                           `json:"token,omitempty"`
	PaymentSource                              *paymentsource.PaymentSource                                                           `json:"payment_source,omitempty"`
	ThirdPartyPaymentMethod                    *thirdpartypaymentmethod.ThirdPartyPaymentMethod                                       `json:"third_party_payment_method,omitempty"`
	VirtualBankAccount                         *virtualbankaccount.VirtualBankAccount                                                 `json:"virtual_bank_account,omitempty"`
	Card                                       *card.Card                                                                             `json:"card,omitempty"`
	PromotionalCredit                          *promotionalcredit.PromotionalCredit                                                   `json:"promotional_credit,omitempty"`
	Invoice                                    *invoice.Invoice                                                                       `json:"invoice,omitempty"`
	PaymentReferenceNumber                     *paymentreferencenumber.PaymentReferenceNumber                                         `json:"payment_reference_number,omitempty"`
	PaymentSchedule                            *paymentschedule.PaymentSchedule                                                       `json:"payment_schedule,omitempty"`
	TaxWithheld                                *taxwithheld.TaxWithheld                                                               `json:"tax_withheld,omitempty"`
	CreditNote                                 *creditnote.CreditNote                                                                 `json:"credit_note,omitempty"`
	UnbilledCharge                             *unbilledcharge.UnbilledCharge                                                         `json:"unbilled_charge,omitempty"`
	Order                                      *order.Order                                                                           `json:"order,omitempty"`
	Gift                                       *gift.Gift                                                                             `json:"gift,omitempty"`
	Transaction                                *transaction.Transaction                                                               `json:"transaction,omitempty"`
	HostedPage                                 *hostedpage.HostedPage                                                                 `json:"hosted_page,omitempty"`
	Estimate                                   *estimate.Estimate                                                                     `json:"estimate,omitempty"`
	Quote                                      *quote.Quote                                                                           `json:"quote,omitempty"`
	QuotedSubscription                         *quotedsubscription.QuotedSubscription                                                 `json:"quoted_subscription,omitempty"`
	QuotedCharge                               *quotedcharge.QuotedCharge                                                             `json:"quoted_charge,omitempty"`
	QuotedRamp                                 *quotedramp.QuotedRamp                                                                 `json:"quoted_ramp,omitempty"`
	BillingConfiguration                       *billingconfiguration.BillingConfiguration                                             `json:"billing_configuration,omitempty"`
	QuoteLineGroup                             *quotelinegroup.QuoteLineGroup                                                         `json:"quote_line_group,omitempty"`
	Plan                                       *plan.Plan                                                                             `json:"plan,omitempty"`
	Addon                                      *addon.Addon                                                                           `json:"addon,omitempty"`
	Coupon                                     *coupon.Coupon                                                                         `json:"coupon,omitempty"`
	CouponSet                                  *couponset.CouponSet                                                                   `json:"coupon_set,omitempty"`
	CouponCode                                 *couponcode.CouponCode                                                                 `json:"coupon_code,omitempty"`
	Address                                    *address.Address                                                                       `json:"address,omitempty"`
	Usage                                      *usage.Usage                                                                           `json:"usage,omitempty"`
	Event                                      *event.Event                                                                           `json:"event,omitempty"`
	Comment                                    *comment.Comment                                                                       `json:"comment,omitempty"`
	Download                                   *download.Download                                                                     `json:"download,omitempty"`
	PortalSession                              *portalsession.PortalSession                                                           `json:"portal_session,omitempty"`
	SiteMigrationDetail                        *sitemigrationdetail.SiteMigrationDetail                                               `json:"site_migration_detail,omitempty"`
	ResourceMigration                          *resourcemigration.ResourceMigration                                                   `json:"resource_migration,omitempty"`
	TimeMachine                                *timemachine.TimeMachine                                                               `json:"time_machine,omitempty"`
	Export                                     *export.Export                                                                         `json:"export,omitempty"`
	PaymentIntent                              *paymentintent.PaymentIntent                                                           `json:"payment_intent,omitempty"`
	GatewayErrorDetail                         *gatewayerrordetail.GatewayErrorDetail                                                 `json:"gateway_error_detail,omitempty"`
	ItemFamily                                 *itemfamily.ItemFamily                                                                 `json:"item_family,omitempty"`
	Item                                       *item.Item                                                                             `json:"item,omitempty"`
	PriceVariant                               *pricevariant.PriceVariant                                                             `json:"price_variant,omitempty"`
	Attribute                                  *attribute.Attribute                                                                   `json:"attribute,omitempty"`
	ItemPrice                                  *itemprice.ItemPrice                                                                   `json:"item_price,omitempty"`
	AttachedItem                               *attacheditem.AttachedItem                                                             `json:"attached_item,omitempty"`
	DifferentialPrice                          *differentialprice.DifferentialPrice                                                   `json:"differential_price,omitempty"`
	Configuration                              *configuration.Configuration                                                           `json:"configuration,omitempty"`
	Feature                                    *feature.Feature                                                                       `json:"feature,omitempty"`
	ImpactedSubscription                       *impactedsubscription.ImpactedSubscription                                             `json:"impacted_subscription,omitempty"`
	ImpactedItem                               *impacteditem.ImpactedItem                                                             `json:"impacted_item,omitempty"`
	ImpactedItemPrice                          *impacteditemprice.ImpactedItemPrice                                                   `json:"impacted_item_price,omitempty"`
	Metadata                                   *metadata.Metadata                                                                     `json:"metadata,omitempty"`
	SubscriptionEntitlement                    *subscriptionentitlement.SubscriptionEntitlement                                       `json:"subscription_entitlement,omitempty"`
	CustomerEntitlement                        *customerentitlement.CustomerEntitlement                                               `json:"customer_entitlement,omitempty"`
	ItemEntitlement                            *itementitlement.ItemEntitlement                                                       `json:"item_entitlement,omitempty"`
	Entitlement                                *entitlement.Entitlement                                                               `json:"entitlement,omitempty"`
	InAppSubscription                          *inappsubscription.InAppSubscription                                                   `json:"in_app_subscription,omitempty"`
	EntitlementOverride                        *entitlementoverride.EntitlementOverride                                               `json:"entitlement_override,omitempty"`
	BusinessEntity                             *businessentity.BusinessEntity                                                         `json:"business_entity,omitempty"`
	Purchase                                   *purchase.Purchase                                                                     `json:"purchase,omitempty"`
	PaymentVoucher                             *paymentvoucher.PaymentVoucher                                                         `json:"payment_voucher,omitempty"`
	Currency                                   *currency.Currency                                                                     `json:"currency,omitempty"`
	Ramp                                       *ramp.Ramp                                                                             `json:"ramp,omitempty"`
	PaymentScheduleScheme                      *paymentschedulescheme.PaymentScheduleScheme                                           `json:"payment_schedule_scheme,omitempty"`
	PricingPageSession                         *pricingpagesession.PricingPageSession                                                 `json:"pricing_page_session,omitempty"`
	OmnichannelSubscription                    *omnichannelsubscription.OmnichannelSubscription                                       `json:"omnichannel_subscription,omitempty"`
	OmnichannelTransaction                     *omnichanneltransaction.OmnichannelTransaction                                         `json:"omnichannel_transaction,omitempty"`
	OmnichannelSubscriptionItem                *omnichannelsubscriptionitem.OmnichannelSubscriptionItem                               `json:"omnichannel_subscription_item,omitempty"`
	RecordedPurchase                           *recordedpurchase.RecordedPurchase                                                     `json:"recorded_purchase,omitempty"`
	OmnichannelOneTimeOrder                    *omnichannelonetimeorder.OmnichannelOneTimeOrder                                       `json:"omnichannel_one_time_order,omitempty"`
	OmnichannelOneTimeOrderItem                *omnichannelonetimeorderitem.OmnichannelOneTimeOrderItem                               `json:"omnichannel_one_time_order_item,omitempty"`
	Rule                                       *rule.Rule                                                                             `json:"rule,omitempty"`
	UsageEvent                                 *usageevent.UsageEvent                                                                 `json:"usage_event,omitempty"`
	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`
	UsageFile                                  *usagefile.UsageFile                                                                   `json:"usage_file,omitempty"`
	PersonalizedOffer                          *personalizedoffer.PersonalizedOffer                                                   `json:"personalized_offer,omitempty"`
	Brand                                      *brand.Brand                                                                           `json:"brand,omitempty"`
	OfferFulfillment                           *offerfulfillment.OfferFulfillment                                                     `json:"offer_fulfillment,omitempty"`
	OfferEvent                                 *offerevent.OfferEvent                                                                 `json:"offer_event,omitempty"`
	WebhookEndpoint                            *webhookendpoint.WebhookEndpoint                                                       `json:"webhook_endpoint,omitempty"`
	ImpactedCustomer                           *impactedcustomer.ImpactedCustomer                                                     `json:"impacted_customer,omitempty"`
	SubscriptionEntitlementsUpdatedDetail      *subscriptionentitlementsupdateddetail.SubscriptionEntitlementsUpdatedDetail           `json:"subscription_entitlements_updated_detail,omitempty"`
	SubscriptionEntitlementsCreatedDetail      *subscriptionentitlementscreateddetail.SubscriptionEntitlementsCreatedDetail           `json:"subscription_entitlements_created_detail,omitempty"`
	AdvanceInvoiceSchedules                    []*advanceinvoiceschedule.AdvanceInvoiceSchedule                                       `json:"advance_invoice_schedules,omitempty"`
	Hierarchies                                []*hierarchy.Hierarchy                                                                 `json:"hierarchies,omitempty"`
	Invoices                                   []*invoice.Invoice                                                                     `json:"invoices,omitempty"`
	PaymentSchedules                           []*paymentschedule.PaymentSchedule                                                     `json:"payment_schedules,omitempty"`
	CreditNotes                                []*creditnote.CreditNote                                                               `json:"credit_notes,omitempty"`
	UnbilledCharges                            []*unbilledcharge.UnbilledCharge                                                       `json:"unbilled_charges,omitempty"`
	Downloads                                  []*download.Download                                                                   `json:"downloads,omitempty"`
	Configurations                             []*configuration.Configuration                                                         `json:"configurations,omitempty"`
	InAppSubscriptions                         []*inappsubscription.InAppSubscription                                                 `json:"in_app_subscriptions,omitempty"`
	PersonalizedOffers                         []*personalizedoffer.PersonalizedOffer                                                 `json:"personalized_offers,omitempty"`
	FailedEvents                               interface{}                                                                            `json:"failed_events,omitempty"`
	ExpiresAt                                  interface{}                                                                            `json:"expires_at,omitempty"`
	BatchId                                    interface{}                                                                            `json:"batch_id,omitempty"`
	Success                                    interface{}                                                                            `json:"success,omitempty"`
	ScheduledAt                                interface{}                                                                            `json:"scheduled_at,omitempty"`
	List                                       interface{}                                                                            `json:"list,omitempty"`
	responseHeaders                            http.Header
	httpStatusCode                             int
}

func (rl *ResultList) GetResponseHeaders() http.Header {
	return rl.responseHeaders
}
func (rl *ResultList) GetHttpStatusCode() int {
	return rl.httpStatusCode
}
func (r *Result) GetResponseHeaders() http.Header {
	return r.responseHeaders
}
func (r *Result) GetHttpStatusCode() int {
	return r.httpStatusCode
}

func (r *Result) IsIdempotencyReplayed() bool {
	value := r.responseHeaders.Get(IdempotencyReplayHeader)
	replayed, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return replayed
}

type ResponseMeta struct {
	Headers    http.Header
	Status     string
	StatusCode int
}

type CBResponse struct {
	Body []byte
	ResponseMeta
}
