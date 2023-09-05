package chargebee

import (
	"github.com/chargebee/chargebee-go/v3/models/addon"
	"github.com/chargebee/chargebee-go/v3/models/address"
	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"
	"github.com/chargebee/chargebee-go/v3/models/attacheditem"
	"github.com/chargebee/chargebee-go/v3/models/card"
	"github.com/chargebee/chargebee-go/v3/models/comment"
	"github.com/chargebee/chargebee-go/v3/models/contact"
	"github.com/chargebee/chargebee-go/v3/models/contractterm"
	"github.com/chargebee/chargebee-go/v3/models/coupon"
	"github.com/chargebee/chargebee-go/v3/models/couponcode"
	"github.com/chargebee/chargebee-go/v3/models/couponset"
	"github.com/chargebee/chargebee-go/v3/models/creditnote"
	"github.com/chargebee/chargebee-go/v3/models/customer"
	"github.com/chargebee/chargebee-go/v3/models/differentialprice"
	"github.com/chargebee/chargebee-go/v3/models/discount"
	"github.com/chargebee/chargebee-go/v3/models/download"
	"github.com/chargebee/chargebee-go/v3/models/entitlementoverride"
	"github.com/chargebee/chargebee-go/v3/models/estimate"
	"github.com/chargebee/chargebee-go/v3/models/event"
	"github.com/chargebee/chargebee-go/v3/models/export"
	"github.com/chargebee/chargebee-go/v3/models/feature"
	"github.com/chargebee/chargebee-go/v3/models/gift"
	"github.com/chargebee/chargebee-go/v3/models/hierarchy"
	"github.com/chargebee/chargebee-go/v3/models/hostedpage"
	"github.com/chargebee/chargebee-go/v3/models/impacteditem"
	"github.com/chargebee/chargebee-go/v3/models/impactedsubscription"
	"github.com/chargebee/chargebee-go/v3/models/inappsubscription"
	"github.com/chargebee/chargebee-go/v3/models/invoice"
	"github.com/chargebee/chargebee-go/v3/models/item"
	"github.com/chargebee/chargebee-go/v3/models/itementitlement"
	"github.com/chargebee/chargebee-go/v3/models/itemfamily"
	"github.com/chargebee/chargebee-go/v3/models/itemprice"
	"github.com/chargebee/chargebee-go/v3/models/order"
	"github.com/chargebee/chargebee-go/v3/models/paymentintent"
	"github.com/chargebee/chargebee-go/v3/models/paymentsource"
	"github.com/chargebee/chargebee-go/v3/models/paymentreferencenumber"
	"github.com/chargebee/chargebee-go/v3/models/plan"
	"github.com/chargebee/chargebee-go/v3/models/portalsession"
	"github.com/chargebee/chargebee-go/v3/models/promotionalcredit"
	"github.com/chargebee/chargebee-go/v3/models/purchase"
	"github.com/chargebee/chargebee-go/v3/models/quote"
	"github.com/chargebee/chargebee-go/v3/models/quotedcharge"
	"github.com/chargebee/chargebee-go/v3/models/quotedsubscription"
	"github.com/chargebee/chargebee-go/v3/models/quotelinegroup"
	"github.com/chargebee/chargebee-go/v3/models/resourcemigration"
	"github.com/chargebee/chargebee-go/v3/models/sitemigrationdetail"
	"github.com/chargebee/chargebee-go/v3/models/subscription"
	"github.com/chargebee/chargebee-go/v3/models/subscriptionentitlement"
	"github.com/chargebee/chargebee-go/v3/models/taxwithheld"
	"github.com/chargebee/chargebee-go/v3/models/thirdpartypaymentmethod"
	"github.com/chargebee/chargebee-go/v3/models/timemachine"
	"github.com/chargebee/chargebee-go/v3/models/token"
	"github.com/chargebee/chargebee-go/v3/models/transaction"
	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"
	"github.com/chargebee/chargebee-go/v3/models/usage"
	"github.com/chargebee/chargebee-go/v3/models/virtualbankaccount"
	"github.com/chargebee/chargebee-go/v3/models/paymentvoucher"
	"github.com/chargebee/chargebee-go/v3/models/csvtaxrule"
	"net/http"
	"strconv"
)

type ResultList struct {
	List            []*Result `json:"list"`
	NextOffset      string    `json:"next_offset"`
	responseHeaders http.Header
}
type Result struct {
	Subscription            *subscription.Subscription                       `json:"subscription,omitempty"`
	ContractTerm            *contractterm.ContractTerm                       `json:"contract_term,omitempty"`
	Discount                *discount.Discount                               `json:"discount,omitempty"`
	AdvanceInvoiceSchedule  *advanceinvoiceschedule.AdvanceInvoiceSchedule   `json:"advance_invoice_schedule,omitempty"`
	Customer                *customer.Customer                               `json:"customer,omitempty"`
	Hierarchy               *hierarchy.Hierarchy                             `json:"hierarchy,omitempty"`
	Contact                 *contact.Contact                                 `json:"contact,omitempty"`
	Token                   *token.Token                                     `json:"token,omitempty"`
	PaymentSource           *paymentsource.PaymentSource                     `json:"payment_source,omitempty"`
	ThirdPartyPaymentMethod *thirdpartypaymentmethod.ThirdPartyPaymentMethod `json:"third_party_payment_method,omitempty"`
	VirtualBankAccount      *virtualbankaccount.VirtualBankAccount           `json:"virtual_bank_account,omitempty"`
	Card                    *card.Card                                       `json:"card,omitempty"`
	PromotionalCredit       *promotionalcredit.PromotionalCredit             `json:"promotional_credit,omitempty"`
	Invoice                 *invoice.Invoice                                 `json:"invoice,omitempty"`
	PaymentReferenceNumber  *paymentreferencenumber.PaymentReferenceNumber   `json:"payment_reference_number,omitempty"`
	TaxWithheld             *taxwithheld.TaxWithheld                         `json:"tax_withheld,omitempty"`
	CreditNote              *creditnote.CreditNote                           `json:"credit_note,omitempty"`
	UnbilledCharge          *unbilledcharge.UnbilledCharge                   `json:"unbilled_charge,omitempty"`
	Order                   *order.Order                                     `json:"order,omitempty"`
	Gift                    *gift.Gift                                       `json:"gift,omitempty"`
	Transaction             *transaction.Transaction                         `json:"transaction,omitempty"`
	HostedPage              *hostedpage.HostedPage                           `json:"hosted_page,omitempty"`
	Estimate                *estimate.Estimate                               `json:"estimate,omitempty"`
	Quote                   *quote.Quote                                     `json:"quote,omitempty"`
	QuotedSubscription      *quotedsubscription.QuotedSubscription           `json:"quoted_subscription,omitempty"`
	QuotedCharge            *quotedcharge.QuotedCharge                       `json:"quoted_charge,omitempty"`
	QuoteLineGroup          *quotelinegroup.QuoteLineGroup                   `json:"quote_line_group,omitempty"`
	Plan                    *plan.Plan                                       `json:"plan,omitempty"`
	Addon                   *addon.Addon                                     `json:"addon,omitempty"`
	Coupon                  *coupon.Coupon                                   `json:"coupon,omitempty"`
	CouponSet               *couponset.CouponSet                             `json:"coupon_set,omitempty"`
	CouponCode              *couponcode.CouponCode                           `json:"coupon_code,omitempty"`
	Address                 *address.Address                                 `json:"address,omitempty"`
	Usage                   *usage.Usage                                     `json:"usage,omitempty"`
	Event                   *event.Event                                     `json:"event,omitempty"`
	Comment                 *comment.Comment                                 `json:"comment,omitempty"`
	Download                *download.Download                               `json:"download,omitempty"`
	PortalSession           *portalsession.PortalSession                     `json:"portal_session,omitempty"`
	SiteMigrationDetail     *sitemigrationdetail.SiteMigrationDetail         `json:"site_migration_detail,omitempty"`
	ResourceMigration       *resourcemigration.ResourceMigration             `json:"resource_migration,omitempty"`
	TimeMachine             *timemachine.TimeMachine                         `json:"time_machine,omitempty"`
	Export                  *export.Export                                   `json:"export,omitempty"`
	PaymentIntent           *paymentintent.PaymentIntent                     `json:"payment_intent,omitempty"`
	ItemFamily              *itemfamily.ItemFamily                           `json:"item_family,omitempty"`
	Item                    *item.Item                                       `json:"item,omitempty"`
	ItemPrice               *itemprice.ItemPrice                             `json:"item_price,omitempty"`
	AttachedItem            *attacheditem.AttachedItem                       `json:"attached_item,omitempty"`
	DifferentialPrice       *differentialprice.DifferentialPrice             `json:"differential_price,omitempty"`
	Feature                 *feature.Feature                                 `json:"feature,omitempty"`
	ImpactedSubscription    *impactedsubscription.ImpactedSubscription       `json:"impacted_subscription,omitempty"`
	ImpactedItem            *impacteditem.ImpactedItem                       `json:"impacted_item,omitempty"`
	SubscriptionEntitlement *subscriptionentitlement.SubscriptionEntitlement `json:"subscription_entitlement,omitempty"`
	ItemEntitlement         *itementitlement.ItemEntitlement                 `json:"item_entitlement,omitempty"`
	InAppSubscription       *inappsubscription.InAppSubscription             `json:"in_app_subscription,omitempty"`
	EntitlementOverride     *entitlementoverride.EntitlementOverride         `json:"entitlement_override,omitempty"`
	Purchase                *purchase.Purchase                               `json:"purchase,omitempty"`
	PaymentVoucher          *paymentvoucher.PaymentVoucher                   `json:"payment_voucher,omitempty"`
	CsvTaxRule              *csvtaxrule.CsvTaxRule                           `json:"csv_tax_rule,omitempty"`
	UnbilledCharges         []*unbilledcharge.UnbilledCharge                 `json:"unbilled_charges,omitempty"`
	CreditNotes             []*creditnote.CreditNote                         `json:"credit_notes,omitempty"`
	AdvanceInvoiceSchedules []*advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedules,omitempty"`
	Hierarchies             []*hierarchy.Hierarchy                           `json:"hierarchies,omitempty"`
	Downloads               []*download.Download                             `json:"downloads,omitempty"`
	Invoices                []*invoice.Invoice                               `json:"invoices,omitempty"`
	DifferentialPrices      []*differentialprice.DifferentialPrice           `json:"differential_prices,omitempty"`
	InAppSubscriptions      []*inappsubscription.InAppSubscription           `json:"in_app_subscriptions,omitempty"`
	responseHeaders         http.Header
}

func (rl *ResultList) GetResponseHeaders() http.Header {
	return rl.responseHeaders
}

func (r *Result) GetResponseHeaders() http.Header {
	return r.responseHeaders
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