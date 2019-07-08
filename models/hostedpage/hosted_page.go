package hostedpage

import (
	"encoding/json"
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	hostedPageEnum "github.com/chargebee/chargebee-go/models/hostedpage/enum"
)

type HostedPage struct {
	Id              string                       `json:"id"`
	Type            hostedPageEnum.Type          `json:"type"`
	Url             string                       `json:"url"`
	State           hostedPageEnum.State         `json:"state"`
	FailureReason   hostedPageEnum.FailureReason `json:"failure_reason"`
	PassThruContent string                       `json:"pass_thru_content"`
	Embed           bool                         `json:"embed"`
	CreatedAt       int64                        `json:"created_at"`
	ExpiresAt       int64                        `json:"expires_at"`
	Content         json.RawMessage              `json:"content"`
	UpdatedAt       int64                        `json:"updated_at"`
	ResourceVersion int64                        `json:"resource_version"`
	CheckoutInfo    json.RawMessage              `json:"checkout_info"`
	Object          string                       `json:"object"`
}
type CheckoutNewRequestParams struct {
	Subscription            *CheckoutNewSubscriptionParams      `json:"subscription,omitempty"`
	Customer                *CheckoutNewCustomerParams          `json:"customer,omitempty"`
	BillingCycles           *int32                              `json:"billing_cycles,omitempty"`
	Addons                  []*CheckoutNewAddonParams           `json:"addons,omitempty"`
	EventBasedAddons        []*CheckoutNewEventBasedAddonParams `json:"event_based_addons,omitempty"`
	TermsToCharge           *int32                              `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode           `json:"billing_alignment_mode,omitempty"`
	MandatoryAddonsToRemove []string                            `json:"mandatory_addons_to_remove,omitempty"`
	Card                    *CheckoutNewCardParams              `json:"card,omitempty"`
	RedirectUrl             string                              `json:"redirect_url,omitempty"`
	CancelUrl               string                              `json:"cancel_url,omitempty"`
	PassThruContent         string                              `json:"pass_thru_content,omitempty"`
	Embed                   *bool                               `json:"embed,omitempty"`
	IframeMessaging         *bool                               `json:"iframe_messaging,omitempty"`
	BillingAddress          *CheckoutNewBillingAddressParams    `json:"billing_address,omitempty"`
	ShippingAddress         *CheckoutNewShippingAddressParams   `json:"shipping_address,omitempty"`
}
type CheckoutNewSubscriptionParams struct {
	Id             string              `json:"id,omitempty"`
	PlanId         string              `json:"plan_id"`
	PlanQuantity   *int32              `json:"plan_quantity,omitempty"`
	PlanUnitPrice  *int32              `json:"plan_unit_price,omitempty"`
	SetupFee       *int32              `json:"setup_fee,omitempty"`
	StartDate      *int64              `json:"start_date,omitempty"`
	TrialEnd       *int64              `json:"trial_end,omitempty"`
	Coupon         string              `json:"coupon,omitempty"`
	AutoCollection enum.AutoCollection `json:"auto_collection,omitempty"`
	InvoiceNotes   string              `json:"invoice_notes,omitempty"`
	AffiliateToken string              `json:"affiliate_token,omitempty"`
}
type CheckoutNewCustomerParams struct {
	Id                    string          `json:"id,omitempty"`
	Email                 string          `json:"email,omitempty"`
	FirstName             string          `json:"first_name,omitempty"`
	LastName              string          `json:"last_name,omitempty"`
	Company               string          `json:"company,omitempty"`
	Taxability            enum.Taxability `json:"taxability,omitempty"`
	Locale                string          `json:"locale,omitempty"`
	Phone                 string          `json:"phone,omitempty"`
	VatNumber             string          `json:"vat_number,omitempty"`
	ConsolidatedInvoicing *bool           `json:"consolidated_invoicing,omitempty"`
}
type CheckoutNewAddonParams struct {
	Id            string `json:"id,omitempty"`
	Quantity      *int32 `json:"quantity,omitempty"`
	UnitPrice     *int32 `json:"unit_price,omitempty"`
	BillingCycles *int32 `json:"billing_cycles,omitempty"`
}
type CheckoutNewEventBasedAddonParams struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int32        `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
}
type CheckoutNewCardParams struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type CheckoutNewBillingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CheckoutNewShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CheckoutExistingRequestParams struct {
	Subscription            *CheckoutExistingSubscriptionParams      `json:"subscription,omitempty"`
	BillingCycles           *int32                                   `json:"billing_cycles,omitempty"`
	Addons                  []*CheckoutExistingAddonParams           `json:"addons,omitempty"`
	EventBasedAddons        []*CheckoutExistingEventBasedAddonParams `json:"event_based_addons,omitempty"`
	ReplaceAddonList        *bool                                    `json:"replace_addon_list,omitempty"`
	MandatoryAddonsToRemove []string                                 `json:"mandatory_addons_to_remove,omitempty"`
	TermsToCharge           *int32                                   `json:"terms_to_charge,omitempty"`
	ReactivateFrom          *int64                                   `json:"reactivate_from,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode                `json:"billing_alignment_mode,omitempty"`
	Reactivate              *bool                                    `json:"reactivate,omitempty"`
	ForceTermReset          *bool                                    `json:"force_term_reset,omitempty"`
	Customer                *CheckoutExistingCustomerParams          `json:"customer,omitempty"`
	Card                    *CheckoutExistingCardParams              `json:"card,omitempty"`
	RedirectUrl             string                                   `json:"redirect_url,omitempty"`
	CancelUrl               string                                   `json:"cancel_url,omitempty"`
	PassThruContent         string                                   `json:"pass_thru_content,omitempty"`
	Embed                   *bool                                    `json:"embed,omitempty"`
	IframeMessaging         *bool                                    `json:"iframe_messaging,omitempty"`
}
type CheckoutExistingSubscriptionParams struct {
	Id            string `json:"id"`
	PlanId        string `json:"plan_id,omitempty"`
	PlanQuantity  *int32 `json:"plan_quantity,omitempty"`
	PlanUnitPrice *int32 `json:"plan_unit_price,omitempty"`
	SetupFee      *int32 `json:"setup_fee,omitempty"`
	StartDate     *int64 `json:"start_date,omitempty"`
	TrialEnd      *int64 `json:"trial_end,omitempty"`
	Coupon        string `json:"coupon,omitempty"`
	InvoiceNotes  string `json:"invoice_notes,omitempty"`
}
type CheckoutExistingAddonParams struct {
	Id            string `json:"id,omitempty"`
	Quantity      *int32 `json:"quantity,omitempty"`
	UnitPrice     *int32 `json:"unit_price,omitempty"`
	BillingCycles *int32 `json:"billing_cycles,omitempty"`
}
type CheckoutExistingEventBasedAddonParams struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int32        `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
}
type CheckoutExistingCustomerParams struct {
	VatNumber string `json:"vat_number,omitempty"`
}
type CheckoutExistingCardParams struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type UpdateCardRequestParams struct {
	Customer        *UpdateCardCustomerParams `json:"customer,omitempty"`
	Card            *UpdateCardCardParams     `json:"card,omitempty"`
	RedirectUrl     string                    `json:"redirect_url,omitempty"`
	CancelUrl       string                    `json:"cancel_url,omitempty"`
	PassThruContent string                    `json:"pass_thru_content,omitempty"`
	Embed           *bool                     `json:"embed,omitempty"`
	IframeMessaging *bool                     `json:"iframe_messaging,omitempty"`
}
type UpdateCardCustomerParams struct {
	Id        string `json:"id"`
	VatNumber string `json:"vat_number,omitempty"`
}
type UpdateCardCardParams struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type UpdatePaymentMethodRequestParams struct {
	Customer        *UpdatePaymentMethodCustomerParams `json:"customer,omitempty"`
	Card            *UpdatePaymentMethodCardParams     `json:"card,omitempty"`
	RedirectUrl     string                             `json:"redirect_url,omitempty"`
	CancelUrl       string                             `json:"cancel_url,omitempty"`
	PassThruContent string                             `json:"pass_thru_content,omitempty"`
	Embed           *bool                              `json:"embed,omitempty"`
	IframeMessaging *bool                              `json:"iframe_messaging,omitempty"`
}
type UpdatePaymentMethodCustomerParams struct {
	Id        string `json:"id"`
	VatNumber string `json:"vat_number,omitempty"`
}
type UpdatePaymentMethodCardParams struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type ManagePaymentSourcesRequestParams struct {
	Customer    *ManagePaymentSourcesCustomerParams `json:"customer,omitempty"`
	RedirectUrl string                              `json:"redirect_url,omitempty"`
	Card        *ManagePaymentSourcesCardParams     `json:"card,omitempty"`
}
type ManagePaymentSourcesCustomerParams struct {
	Id string `json:"id"`
}
type ManagePaymentSourcesCardParams struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type CollectNowRequestParams struct {
	Customer     *CollectNowCustomerParams `json:"customer,omitempty"`
	RedirectUrl  string                    `json:"redirect_url,omitempty"`
	Card         *CollectNowCardParams     `json:"card,omitempty"`
	CurrencyCode string                    `json:"currency_code,omitempty"`
}
type CollectNowCustomerParams struct {
	Id string `json:"id"`
}
type CollectNowCardParams struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type AcceptQuoteRequestParams struct {
	Quote *AcceptQuoteQuoteParams `json:"quote,omitempty"`
}
type AcceptQuoteQuoteParams struct {
	Id string `json:"id"`
}
type ExtendSubscriptionRequestParams struct {
	Subscription *ExtendSubscriptionSubscriptionParams `json:"subscription,omitempty"`
	Expiry       *int32                                `json:"expiry,omitempty"`
	BillingCycle *int32                                `json:"billing_cycle,omitempty"`
}
type ExtendSubscriptionSubscriptionParams struct {
	Id string `json:"id"`
}
type CheckoutGiftRequestParams struct {
	Gifter       *CheckoutGiftGifterParams       `json:"gifter,omitempty"`
	RedirectUrl  string                          `json:"redirect_url,omitempty"`
	Subscription *CheckoutGiftSubscriptionParams `json:"subscription,omitempty"`
	Addons       []*CheckoutGiftAddonParams      `json:"addons,omitempty"`
}
type CheckoutGiftGifterParams struct {
	CustomerId string `json:"customer_id,omitempty"`
	Locale     string `json:"locale,omitempty"`
}
type CheckoutGiftSubscriptionParams struct {
	PlanId       string `json:"plan_id"`
	PlanQuantity *int32 `json:"plan_quantity,omitempty"`
	Coupon       string `json:"coupon,omitempty"`
}
type CheckoutGiftAddonParams struct {
	Id       string `json:"id,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
}
type ClaimGiftRequestParams struct {
	Gift        *ClaimGiftGiftParams     `json:"gift,omitempty"`
	RedirectUrl string                   `json:"redirect_url,omitempty"`
	Customer    *ClaimGiftCustomerParams `json:"customer,omitempty"`
}
type ClaimGiftGiftParams struct {
	Id string `json:"id"`
}
type ClaimGiftCustomerParams struct {
	Locale string `json:"locale,omitempty"`
}
type RetrieveAgreementPdfRequestParams struct {
	PaymentSourceId string `json:"payment_source_id"`
}
type ListRequestParams struct {
	Limit     *int32                  `json:"limit,omitempty"`
	Offset    string                  `json:"offset,omitempty"`
	Id        *filter.StringFilter    `json:"id,omitempty"`
	Type      *filter.EnumFilter      `json:"type,omitempty"`
	State     *filter.EnumFilter      `json:"state,omitempty"`
	UpdatedAt *filter.TimestampFilter `json:"updated_at,omitempty"`
}
