package chargebee

import (
	"encoding/json"
)

type HostedPageType string

const (
	HostedPageTypeCheckoutNew          HostedPageType = "checkout_new"
	HostedPageTypeCheckoutExisting     HostedPageType = "checkout_existing"
	HostedPageTypeUpdatePaymentMethod  HostedPageType = "update_payment_method"
	HostedPageTypeManagePaymentSources HostedPageType = "manage_payment_sources"
	HostedPageTypeCollectNow           HostedPageType = "collect_now"
	HostedPageTypeExtendSubscription   HostedPageType = "extend_subscription"
	HostedPageTypeCheckoutOneTime      HostedPageType = "checkout_one_time"
	HostedPageTypePreCancel            HostedPageType = "pre_cancel"
	HostedPageTypeViewVoucher          HostedPageType = "view_voucher"
	HostedPageTypeAcceptQuote          HostedPageType = "accept_quote"
	HostedPageTypeCheckoutGift         HostedPageType = "checkout_gift"
	HostedPageTypeClaimGift            HostedPageType = "claim_gift"
	HostedPageTypeUpdateCard           HostedPageType = "update_card"
)

type HostedPageState string

const (
	HostedPageStateCreated      HostedPageState = "created"
	HostedPageStateRequested    HostedPageState = "requested"
	HostedPageStateSucceeded    HostedPageState = "succeeded"
	HostedPageStateCancelled    HostedPageState = "cancelled"
	HostedPageStateAcknowledged HostedPageState = "acknowledged"
	HostedPageStateFailed       HostedPageState = "failed"
)

type HostedPageFailureReason string

const (
	HostedPageFailureReasonCardError   HostedPageFailureReason = "card_error"
	HostedPageFailureReasonServerError HostedPageFailureReason = "server_error"
)

type HostedPageBillingAlignmentMode string

const (
	HostedPageBillingAlignmentModeImmediate HostedPageBillingAlignmentMode = "immediate"
	HostedPageBillingAlignmentModeDelayed   HostedPageBillingAlignmentMode = "delayed"
)

type HostedPageLayout string

const (
	HostedPageLayoutInApp    HostedPageLayout = "in_app"
	HostedPageLayoutFullPage HostedPageLayout = "full_page"
)

type HostedPageChangeOption string

const (
	HostedPageChangeOptionImmediately  HostedPageChangeOption = "immediately"
	HostedPageChangeOptionEndOfTerm    HostedPageChangeOption = "end_of_term"
	HostedPageChangeOptionSpecificDate HostedPageChangeOption = "specific_date"
)

type HostedPageEventName string

const (
	HostedPageEventNameCancellationPageLoaded HostedPageEventName = "cancellation_page_loaded"
)

// just struct
type HostedPage struct {
	Id               string          `json:"id"`
	Type             HostedPageType  `json:"type"`
	Url              string          `json:"url"`
	State            HostedPageState `json:"state"`
	PassThruContent  string          `json:"pass_thru_content"`
	Embed            bool            `json:"embed"`
	CreatedAt        int64           `json:"created_at"`
	ExpiresAt        int64           `json:"expires_at"`
	Content          json.RawMessage `json:"content"`
	UpdatedAt        int64           `json:"updated_at"`
	ResourceVersion  int64           `json:"resource_version"`
	CheckoutInfo     json.RawMessage `json:"checkout_info"`
	BusinessEntityId string          `json:"business_entity_id"`
	Object           string          `json:"object"`
}

// sub resources
// operations
// input params
type HostedPageCheckoutNewRequest struct {
	Subscription               *HostedPageCheckoutNewSubscription      `json:"subscription,omitempty"`
	Customer                   *HostedPageCheckoutNewCustomer          `json:"customer,omitempty"`
	BillingCycles              *int32                                  `json:"billing_cycles,omitempty"`
	Addons                     []*HostedPageCheckoutNewAddon           `json:"addons,omitempty"`
	EventBasedAddons           []*HostedPageCheckoutNewEventBasedAddon `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove    []string                                `json:"mandatory_addons_to_remove,omitempty"`
	TermsToCharge              *int32                                  `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode       HostedPageBillingAlignmentMode          `json:"billing_alignment_mode,omitempty"`
	CouponIds                  []string                                `json:"coupon_ids,omitempty"`
	Card                       *HostedPageCheckoutNewCard              `json:"card,omitempty"`
	RedirectUrl                string                                  `json:"redirect_url,omitempty"`
	CancelUrl                  string                                  `json:"cancel_url,omitempty"`
	PassThruContent            string                                  `json:"pass_thru_content,omitempty"`
	Embed                      *bool                                   `json:"embed,omitempty"`
	IframeMessaging            *bool                                   `json:"iframe_messaging,omitempty"`
	AllowOfflinePaymentMethods *bool                                   `json:"allow_offline_payment_methods,omitempty"`
	BillingAddress             *HostedPageCheckoutNewBillingAddress    `json:"billing_address,omitempty"`
	ShippingAddress            *HostedPageCheckoutNewShippingAddress   `json:"shipping_address,omitempty"`
	ContractTerm               *HostedPageCheckoutNewContractTerm      `json:"contract_term,omitempty"`
	apiRequest                 `json:"-" form:"-"`
}

func (r *HostedPageCheckoutNewRequest) payload() any { return r }

// input sub resource params single
type HostedPageCheckoutNewSubscription struct {
	Id                                string                                     `json:"id,omitempty"`
	PlanId                            string                                     `json:"plan_id"`
	PlanQuantity                      *int32                                     `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                                     `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int64                                     `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                                     `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int64                                     `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                                     `json:"trial_end,omitempty"`
	StartDate                         *int64                                     `json:"start_date,omitempty"`
	Coupon                            string                                     `json:"coupon,omitempty"`
	AutoCollection                    HostedPageSubscriptionAutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              HostedPageSubscriptionOfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	InvoiceNotes                      string                                     `json:"invoice_notes,omitempty"`
	AffiliateToken                    string                                     `json:"affiliate_token,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                     `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutNewCustomer struct {
	Id                    string                       `json:"id,omitempty"`
	Email                 string                       `json:"email,omitempty"`
	FirstName             string                       `json:"first_name,omitempty"`
	LastName              string                       `json:"last_name,omitempty"`
	Company               string                       `json:"company,omitempty"`
	Phone                 string                       `json:"phone,omitempty"`
	Locale                string                       `json:"locale,omitempty"`
	Taxability            HostedPageCustomerTaxability `json:"taxability,omitempty"`
	VatNumber             string                       `json:"vat_number,omitempty"`
	VatNumberPrefix       string                       `json:"vat_number_prefix,omitempty"`
	ConsolidatedInvoicing *bool                        `json:"consolidated_invoicing,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutNewAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutNewEventBasedAddon struct {
	Id                  string                            `json:"id,omitempty"`
	Quantity            *int32                            `json:"quantity,omitempty"`
	UnitPrice           *int64                            `json:"unit_price,omitempty"`
	QuantityInDecimal   string                            `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string                            `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodInDays *int32                            `json:"service_period_in_days,omitempty"`
	OnEvent             HostedPageEventBasedAddonOnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool                             `json:"charge_once,omitempty"`
	ChargeOn            HostedPageEventBasedAddonChargeOn `json:"charge_on,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutNewCard struct {
	Gateway          HostedPageCardGateway `json:"gateway,omitempty"`
	GatewayAccountId string                `json:"gateway_account_id,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutNewBillingAddress struct {
	FirstName        string                                   `json:"first_name,omitempty"`
	LastName         string                                   `json:"last_name,omitempty"`
	Email            string                                   `json:"email,omitempty"`
	Company          string                                   `json:"company,omitempty"`
	Phone            string                                   `json:"phone,omitempty"`
	Line1            string                                   `json:"line1,omitempty"`
	Line2            string                                   `json:"line2,omitempty"`
	Line3            string                                   `json:"line3,omitempty"`
	City             string                                   `json:"city,omitempty"`
	StateCode        string                                   `json:"state_code,omitempty"`
	State            string                                   `json:"state,omitempty"`
	Zip              string                                   `json:"zip,omitempty"`
	Country          string                                   `json:"country,omitempty"`
	ValidationStatus HostedPageBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutNewShippingAddress struct {
	FirstName        string                                    `json:"first_name,omitempty"`
	LastName         string                                    `json:"last_name,omitempty"`
	Email            string                                    `json:"email,omitempty"`
	Company          string                                    `json:"company,omitempty"`
	Phone            string                                    `json:"phone,omitempty"`
	Line1            string                                    `json:"line1,omitempty"`
	Line2            string                                    `json:"line2,omitempty"`
	Line3            string                                    `json:"line3,omitempty"`
	City             string                                    `json:"city,omitempty"`
	StateCode        string                                    `json:"state_code,omitempty"`
	State            string                                    `json:"state,omitempty"`
	Zip              string                                    `json:"zip,omitempty"`
	Country          string                                    `json:"country,omitempty"`
	ValidationStatus HostedPageShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutNewContractTerm struct {
	ActionAtTermEnd          HostedPageContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                                `json:"cancellation_cutoff_period,omitempty"`
}

type HostedPageCheckoutOneTimeRequest struct {
	Customer        *HostedPageCheckoutOneTimeCustomer        `json:"customer,omitempty"`
	Addons          []*HostedPageCheckoutOneTimeAddon         `json:"addons,omitempty"`
	CurrencyCode    string                                    `json:"currency_code,omitempty"`
	Charges         []*HostedPageCheckoutOneTimeCharge        `json:"charges,omitempty"`
	InvoiceNote     string                                    `json:"invoice_note,omitempty"`
	Invoice         *HostedPageCheckoutOneTimeInvoice         `json:"invoice,omitempty"`
	CouponIds       []string                                  `json:"coupon_ids,omitempty"`
	Card            *HostedPageCheckoutOneTimeCard            `json:"card,omitempty"`
	RedirectUrl     string                                    `json:"redirect_url,omitempty"`
	CancelUrl       string                                    `json:"cancel_url,omitempty"`
	PassThruContent string                                    `json:"pass_thru_content,omitempty"`
	Embed           *bool                                     `json:"embed,omitempty"`
	IframeMessaging *bool                                     `json:"iframe_messaging,omitempty"`
	BillingAddress  *HostedPageCheckoutOneTimeBillingAddress  `json:"billing_address,omitempty"`
	ShippingAddress *HostedPageCheckoutOneTimeShippingAddress `json:"shipping_address,omitempty"`
	apiRequest      `json:"-" form:"-"`
}

func (r *HostedPageCheckoutOneTimeRequest) payload() any { return r }

// input sub resource params single
type HostedPageCheckoutOneTimeCustomer struct {
	Id                    string                       `json:"id,omitempty"`
	Email                 string                       `json:"email,omitempty"`
	FirstName             string                       `json:"first_name,omitempty"`
	LastName              string                       `json:"last_name,omitempty"`
	Company               string                       `json:"company,omitempty"`
	Phone                 string                       `json:"phone,omitempty"`
	Locale                string                       `json:"locale,omitempty"`
	Taxability            HostedPageCustomerTaxability `json:"taxability,omitempty"`
	VatNumber             string                       `json:"vat_number,omitempty"`
	VatNumberPrefix       string                       `json:"vat_number_prefix,omitempty"`
	ConsolidatedInvoicing *bool                        `json:"consolidated_invoicing,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutOneTimeAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutOneTimeCharge struct {
	Amount                 *int64                          `json:"amount,omitempty"`
	AmountInDecimal        string                          `json:"amount_in_decimal,omitempty"`
	Description            string                          `json:"description,omitempty"`
	Taxable                *bool                           `json:"taxable,omitempty"`
	TaxProfileId           string                          `json:"tax_profile_id,omitempty"`
	AvalaraTaxCode         string                          `json:"avalara_tax_code,omitempty"`
	HsnCode                string                          `json:"hsn_code,omitempty"`
	TaxjarProductCode      string                          `json:"taxjar_product_code,omitempty"`
	AvalaraSaleType        HostedPageChargeAvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32                          `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32                          `json:"avalara_service_type,omitempty"`
	DateFrom               *int64                          `json:"date_from,omitempty"`
	DateTo                 *int64                          `json:"date_to,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutOneTimeInvoice struct {
	PoNumber string `json:"po_number,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutOneTimeCard struct {
	Gateway          HostedPageCardGateway `json:"gateway,omitempty"`
	GatewayAccountId string                `json:"gateway_account_id,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutOneTimeBillingAddress struct {
	FirstName        string                                   `json:"first_name,omitempty"`
	LastName         string                                   `json:"last_name,omitempty"`
	Email            string                                   `json:"email,omitempty"`
	Company          string                                   `json:"company,omitempty"`
	Phone            string                                   `json:"phone,omitempty"`
	Line1            string                                   `json:"line1,omitempty"`
	Line2            string                                   `json:"line2,omitempty"`
	Line3            string                                   `json:"line3,omitempty"`
	City             string                                   `json:"city,omitempty"`
	StateCode        string                                   `json:"state_code,omitempty"`
	State            string                                   `json:"state,omitempty"`
	Zip              string                                   `json:"zip,omitempty"`
	Country          string                                   `json:"country,omitempty"`
	ValidationStatus HostedPageBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutOneTimeShippingAddress struct {
	FirstName        string                                    `json:"first_name,omitempty"`
	LastName         string                                    `json:"last_name,omitempty"`
	Email            string                                    `json:"email,omitempty"`
	Company          string                                    `json:"company,omitempty"`
	Phone            string                                    `json:"phone,omitempty"`
	Line1            string                                    `json:"line1,omitempty"`
	Line2            string                                    `json:"line2,omitempty"`
	Line3            string                                    `json:"line3,omitempty"`
	City             string                                    `json:"city,omitempty"`
	StateCode        string                                    `json:"state_code,omitempty"`
	State            string                                    `json:"state,omitempty"`
	Zip              string                                    `json:"zip,omitempty"`
	Country          string                                    `json:"country,omitempty"`
	ValidationStatus HostedPageShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

type HostedPageCheckoutOneTimeForItemsRequest struct {
	BusinessEntityId  string                                               `json:"business_entity_id,omitempty"`
	Layout            HostedPageLayout                                     `json:"layout,omitempty"`
	Customer          *HostedPageCheckoutOneTimeForItemsCustomer           `json:"customer,omitempty"`
	ItemPrices        []*HostedPageCheckoutOneTimeForItemsItemPrice        `json:"item_prices,omitempty"`
	ItemTiers         []*HostedPageCheckoutOneTimeForItemsItemTier         `json:"item_tiers,omitempty"`
	Charges           []*HostedPageCheckoutOneTimeForItemsCharge           `json:"charges,omitempty"`
	Discounts         []*HostedPageCheckoutOneTimeForItemsDiscount         `json:"discounts,omitempty"`
	InvoiceNote       string                                               `json:"invoice_note,omitempty"`
	Invoice           *HostedPageCheckoutOneTimeForItemsInvoice            `json:"invoice,omitempty"`
	CouponIds         []string                                             `json:"coupon_ids,omitempty"`
	CurrencyCode      string                                               `json:"currency_code,omitempty"`
	Card              *HostedPageCheckoutOneTimeForItemsCard               `json:"card,omitempty"`
	EntityIdentifiers []*HostedPageCheckoutOneTimeForItemsEntityIdentifier `json:"entity_identifiers,omitempty"`
	RedirectUrl       string                                               `json:"redirect_url,omitempty"`
	CancelUrl         string                                               `json:"cancel_url,omitempty"`
	PassThruContent   string                                               `json:"pass_thru_content,omitempty"`
	BillingAddress    *HostedPageCheckoutOneTimeForItemsBillingAddress     `json:"billing_address,omitempty"`
	ShippingAddress   *HostedPageCheckoutOneTimeForItemsShippingAddress    `json:"shipping_address,omitempty"`
	apiRequest        `json:"-" form:"-"`
}

func (r *HostedPageCheckoutOneTimeForItemsRequest) payload() any { return r }

// input sub resource params single
type HostedPageCheckoutOneTimeForItemsCustomer struct {
	Id                       string                             `json:"id,omitempty"`
	Email                    string                             `json:"email,omitempty"`
	FirstName                string                             `json:"first_name,omitempty"`
	LastName                 string                             `json:"last_name,omitempty"`
	Company                  string                             `json:"company,omitempty"`
	Phone                    string                             `json:"phone,omitempty"`
	Locale                   string                             `json:"locale,omitempty"`
	Taxability               HostedPageCustomerTaxability       `json:"taxability,omitempty"`
	VatNumber                string                             `json:"vat_number,omitempty"`
	VatNumberPrefix          string                             `json:"vat_number_prefix,omitempty"`
	EinvoicingMethod         HostedPageCustomerEinvoicingMethod `json:"einvoicing_method,omitempty"`
	IsEinvoiceEnabled        *bool                              `json:"is_einvoice_enabled,omitempty"`
	EntityIdentifierScheme   string                             `json:"entity_identifier_scheme,omitempty"`
	EntityIdentifierStandard string                             `json:"entity_identifier_standard,omitempty"`
	ConsolidatedInvoicing    *bool                              `json:"consolidated_invoicing,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutOneTimeForItemsItemPrice struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutOneTimeForItemsItemTier struct {
	ItemPriceId           string                        `json:"item_price_id,omitempty"`
	StartingUnit          *int32                        `json:"starting_unit,omitempty"`
	EndingUnit            *int32                        `json:"ending_unit,omitempty"`
	Price                 *int64                        `json:"price,omitempty"`
	StartingUnitInDecimal string                        `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string                        `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string                        `json:"price_in_decimal,omitempty"`
	PricingType           HostedPageItemTierPricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32                        `json:"package_size,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutOneTimeForItemsCharge struct {
	Amount                 *int64                          `json:"amount,omitempty"`
	AmountInDecimal        string                          `json:"amount_in_decimal,omitempty"`
	Description            string                          `json:"description,omitempty"`
	Taxable                *bool                           `json:"taxable,omitempty"`
	TaxProfileId           string                          `json:"tax_profile_id,omitempty"`
	AvalaraTaxCode         string                          `json:"avalara_tax_code,omitempty"`
	HsnCode                string                          `json:"hsn_code,omitempty"`
	TaxjarProductCode      string                          `json:"taxjar_product_code,omitempty"`
	AvalaraSaleType        HostedPageChargeAvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32                          `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32                          `json:"avalara_service_type,omitempty"`
	DateFrom               *int64                          `json:"date_from,omitempty"`
	DateTo                 *int64                          `json:"date_to,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutOneTimeForItemsDiscount struct {
	Percentage  *float64                  `json:"percentage,omitempty"`
	Amount      *int64                    `json:"amount,omitempty"`
	Quantity    *int32                    `json:"quantity,omitempty"`
	ApplyOn     HostedPageDiscountApplyOn `json:"apply_on"`
	ItemPriceId string                    `json:"item_price_id,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutOneTimeForItemsInvoice struct {
	PoNumber string `json:"po_number,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutOneTimeForItemsCard struct {
	Gateway          HostedPageCardGateway `json:"gateway,omitempty"`
	GatewayAccountId string                `json:"gateway_account_id,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutOneTimeForItemsEntityIdentifier struct {
	Id        string                              `json:"id,omitempty"`
	Scheme    string                              `json:"scheme,omitempty"`
	Value     string                              `json:"value,omitempty"`
	Operation HostedPageEntityIdentifierOperation `json:"operation,omitempty"`
	Standard  string                              `json:"standard,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutOneTimeForItemsBillingAddress struct {
	FirstName        string                                   `json:"first_name,omitempty"`
	LastName         string                                   `json:"last_name,omitempty"`
	Email            string                                   `json:"email,omitempty"`
	Company          string                                   `json:"company,omitempty"`
	Phone            string                                   `json:"phone,omitempty"`
	Line1            string                                   `json:"line1,omitempty"`
	Line2            string                                   `json:"line2,omitempty"`
	Line3            string                                   `json:"line3,omitempty"`
	City             string                                   `json:"city,omitempty"`
	StateCode        string                                   `json:"state_code,omitempty"`
	State            string                                   `json:"state,omitempty"`
	Zip              string                                   `json:"zip,omitempty"`
	Country          string                                   `json:"country,omitempty"`
	ValidationStatus HostedPageBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutOneTimeForItemsShippingAddress struct {
	FirstName        string                                    `json:"first_name,omitempty"`
	LastName         string                                    `json:"last_name,omitempty"`
	Email            string                                    `json:"email,omitempty"`
	Company          string                                    `json:"company,omitempty"`
	Phone            string                                    `json:"phone,omitempty"`
	Line1            string                                    `json:"line1,omitempty"`
	Line2            string                                    `json:"line2,omitempty"`
	Line3            string                                    `json:"line3,omitempty"`
	City             string                                    `json:"city,omitempty"`
	StateCode        string                                    `json:"state_code,omitempty"`
	State            string                                    `json:"state,omitempty"`
	Zip              string                                    `json:"zip,omitempty"`
	Country          string                                    `json:"country,omitempty"`
	ValidationStatus HostedPageShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

type HostedPageCheckoutNewForItemsRequest struct {
	Subscription               *HostedPageCheckoutNewForItemsSubscription       `json:"subscription,omitempty"`
	Layout                     HostedPageLayout                                 `json:"layout,omitempty"`
	Customer                   *HostedPageCheckoutNewForItemsCustomer           `json:"customer,omitempty"`
	BusinessEntityId           string                                           `json:"business_entity_id,omitempty"`
	BillingCycles              *int32                                           `json:"billing_cycles,omitempty"`
	SubscriptionItems          []*HostedPageCheckoutNewForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	Discounts                  []*HostedPageCheckoutNewForItemsDiscount         `json:"discounts,omitempty"`
	MandatoryItemsToRemove     []string                                         `json:"mandatory_items_to_remove,omitempty"`
	ItemTiers                  []*HostedPageCheckoutNewForItemsItemTier         `json:"item_tiers,omitempty"`
	TermsToCharge              *int32                                           `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode       HostedPageBillingAlignmentMode                   `json:"billing_alignment_mode,omitempty"`
	CouponIds                  []string                                         `json:"coupon_ids,omitempty"`
	Card                       *HostedPageCheckoutNewForItemsCard               `json:"card,omitempty"`
	EntityIdentifiers          []*HostedPageCheckoutNewForItemsEntityIdentifier `json:"entity_identifiers,omitempty"`
	RedirectUrl                string                                           `json:"redirect_url,omitempty"`
	CancelUrl                  string                                           `json:"cancel_url,omitempty"`
	PassThruContent            string                                           `json:"pass_thru_content,omitempty"`
	AllowOfflinePaymentMethods *bool                                            `json:"allow_offline_payment_methods,omitempty"`
	BillingAddress             *HostedPageCheckoutNewForItemsBillingAddress     `json:"billing_address,omitempty"`
	ShippingAddress            *HostedPageCheckoutNewForItemsShippingAddress    `json:"shipping_address,omitempty"`
	ContractTerm               *HostedPageCheckoutNewForItemsContractTerm       `json:"contract_term,omitempty"`
	apiRequest                 `json:"-" form:"-"`
}

func (r *HostedPageCheckoutNewForItemsRequest) payload() any { return r }

// input sub resource params single
type HostedPageCheckoutNewForItemsSubscription struct {
	Id                                string                                     `json:"id,omitempty"`
	TrialEnd                          *int64                                     `json:"trial_end,omitempty"`
	SetupFee                          *int64                                     `json:"setup_fee,omitempty"`
	StartDate                         *int64                                     `json:"start_date,omitempty"`
	Coupon                            string                                     `json:"coupon,omitempty"`
	AutoCollection                    HostedPageSubscriptionAutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              HostedPageSubscriptionOfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	InvoiceNotes                      string                                     `json:"invoice_notes,omitempty"`
	PoNumber                          string                                     `json:"po_number,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                     `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutNewForItemsCustomer struct {
	Id                       string                             `json:"id,omitempty"`
	Email                    string                             `json:"email,omitempty"`
	FirstName                string                             `json:"first_name,omitempty"`
	LastName                 string                             `json:"last_name,omitempty"`
	Company                  string                             `json:"company,omitempty"`
	Phone                    string                             `json:"phone,omitempty"`
	Locale                   string                             `json:"locale,omitempty"`
	Taxability               HostedPageCustomerTaxability       `json:"taxability,omitempty"`
	VatNumber                string                             `json:"vat_number,omitempty"`
	VatNumberPrefix          string                             `json:"vat_number_prefix,omitempty"`
	IsEinvoiceEnabled        *bool                              `json:"is_einvoice_enabled,omitempty"`
	EntityIdentifierScheme   string                             `json:"entity_identifier_scheme,omitempty"`
	EntityIdentifierStandard string                             `json:"entity_identifier_standard,omitempty"`
	EinvoicingMethod         HostedPageCustomerEinvoicingMethod `json:"einvoicing_method,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutNewForItemsSubscriptionItem struct {
	ItemPriceId        string                                   `json:"item_price_id"`
	Quantity           *int32                                   `json:"quantity,omitempty"`
	QuantityInDecimal  string                                   `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64                                   `json:"unit_price,omitempty"`
	UnitPriceInDecimal string                                   `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32                                   `json:"billing_cycles,omitempty"`
	TrialEnd           *int64                                   `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32                                   `json:"service_period_days,omitempty"`
	ChargeOnEvent      HostedPageSubscriptionItemChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool                                    `json:"charge_once,omitempty"`
	ItemType           HostedPageSubscriptionItemItemType       `json:"item_type,omitempty"`
	ChargeOnOption     HostedPageSubscriptionItemChargeOnOption `json:"charge_on_option,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutNewForItemsDiscount struct {
	ApplyOn       HostedPageDiscountApplyOn      `json:"apply_on,omitempty"`
	DurationType  HostedPageDiscountDurationType `json:"duration_type"`
	Percentage    *float64                       `json:"percentage,omitempty"`
	Amount        *int64                         `json:"amount,omitempty"`
	Period        *int32                         `json:"period,omitempty"`
	PeriodUnit    HostedPageDiscountPeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool                          `json:"included_in_mrr,omitempty"`
	ItemPriceId   string                         `json:"item_price_id,omitempty"`
	Quantity      *int32                         `json:"quantity,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutNewForItemsItemTier struct {
	ItemPriceId           string                        `json:"item_price_id,omitempty"`
	StartingUnit          *int32                        `json:"starting_unit,omitempty"`
	EndingUnit            *int32                        `json:"ending_unit,omitempty"`
	Price                 *int64                        `json:"price,omitempty"`
	StartingUnitInDecimal string                        `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string                        `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string                        `json:"price_in_decimal,omitempty"`
	PricingType           HostedPageItemTierPricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32                        `json:"package_size,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutNewForItemsCard struct {
	Gateway          HostedPageCardGateway `json:"gateway,omitempty"`
	GatewayAccountId string                `json:"gateway_account_id,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutNewForItemsEntityIdentifier struct {
	Id        string                              `json:"id,omitempty"`
	Scheme    string                              `json:"scheme,omitempty"`
	Value     string                              `json:"value,omitempty"`
	Operation HostedPageEntityIdentifierOperation `json:"operation,omitempty"`
	Standard  string                              `json:"standard,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutNewForItemsBillingAddress struct {
	FirstName        string                                   `json:"first_name,omitempty"`
	LastName         string                                   `json:"last_name,omitempty"`
	Email            string                                   `json:"email,omitempty"`
	Company          string                                   `json:"company,omitempty"`
	Phone            string                                   `json:"phone,omitempty"`
	Line1            string                                   `json:"line1,omitempty"`
	Line2            string                                   `json:"line2,omitempty"`
	Line3            string                                   `json:"line3,omitempty"`
	City             string                                   `json:"city,omitempty"`
	StateCode        string                                   `json:"state_code,omitempty"`
	State            string                                   `json:"state,omitempty"`
	Zip              string                                   `json:"zip,omitempty"`
	Country          string                                   `json:"country,omitempty"`
	ValidationStatus HostedPageBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutNewForItemsShippingAddress struct {
	FirstName        string                                    `json:"first_name,omitempty"`
	LastName         string                                    `json:"last_name,omitempty"`
	Email            string                                    `json:"email,omitempty"`
	Company          string                                    `json:"company,omitempty"`
	Phone            string                                    `json:"phone,omitempty"`
	Line1            string                                    `json:"line1,omitempty"`
	Line2            string                                    `json:"line2,omitempty"`
	Line3            string                                    `json:"line3,omitempty"`
	City             string                                    `json:"city,omitempty"`
	StateCode        string                                    `json:"state_code,omitempty"`
	State            string                                    `json:"state,omitempty"`
	Zip              string                                    `json:"zip,omitempty"`
	Country          string                                    `json:"country,omitempty"`
	ValidationStatus HostedPageShippingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutNewForItemsContractTerm struct {
	ActionAtTermEnd          HostedPageContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                                `json:"cancellation_cutoff_period,omitempty"`
}

type HostedPageCheckoutExistingRequest struct {
	Subscription               *HostedPageCheckoutExistingSubscription      `json:"subscription,omitempty"`
	Addons                     []*HostedPageCheckoutExistingAddon           `json:"addons,omitempty"`
	EventBasedAddons           []*HostedPageCheckoutExistingEventBasedAddon `json:"event_based_addons,omitempty"`
	ReplaceAddonList           *bool                                        `json:"replace_addon_list,omitempty"`
	MandatoryAddonsToRemove    []string                                     `json:"mandatory_addons_to_remove,omitempty"`
	InvoiceDate                *int64                                       `json:"invoice_date,omitempty"`
	BillingCycles              *int32                                       `json:"billing_cycles,omitempty"`
	TermsToCharge              *int32                                       `json:"terms_to_charge,omitempty"`
	ReactivateFrom             *int64                                       `json:"reactivate_from,omitempty"`
	BillingAlignmentMode       HostedPageBillingAlignmentMode               `json:"billing_alignment_mode,omitempty"`
	CouponIds                  []string                                     `json:"coupon_ids,omitempty"`
	ReplaceCouponList          *bool                                        `json:"replace_coupon_list,omitempty"`
	Reactivate                 *bool                                        `json:"reactivate,omitempty"`
	ForceTermReset             *bool                                        `json:"force_term_reset,omitempty"`
	Customer                   *HostedPageCheckoutExistingCustomer          `json:"customer,omitempty"`
	Card                       *HostedPageCheckoutExistingCard              `json:"card,omitempty"`
	RedirectUrl                string                                       `json:"redirect_url,omitempty"`
	CancelUrl                  string                                       `json:"cancel_url,omitempty"`
	PassThruContent            string                                       `json:"pass_thru_content,omitempty"`
	Embed                      *bool                                        `json:"embed,omitempty"`
	IframeMessaging            *bool                                        `json:"iframe_messaging,omitempty"`
	AllowOfflinePaymentMethods *bool                                        `json:"allow_offline_payment_methods,omitempty"`
	ContractTerm               *HostedPageCheckoutExistingContractTerm      `json:"contract_term,omitempty"`
	apiRequest                 `json:"-" form:"-"`
}

func (r *HostedPageCheckoutExistingRequest) payload() any { return r }

// input sub resource params single
type HostedPageCheckoutExistingSubscription struct {
	Id                                string                                     `json:"id"`
	PlanId                            string                                     `json:"plan_id,omitempty"`
	PlanQuantity                      *int32                                     `json:"plan_quantity,omitempty"`
	PlanUnitPrice                     *int64                                     `json:"plan_unit_price,omitempty"`
	SetupFee                          *int64                                     `json:"setup_fee,omitempty"`
	PlanQuantityInDecimal             string                                     `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPriceInDecimal            string                                     `json:"plan_unit_price_in_decimal,omitempty"`
	StartDate                         *int64                                     `json:"start_date,omitempty"`
	TrialEnd                          *int64                                     `json:"trial_end,omitempty"`
	Coupon                            string                                     `json:"coupon,omitempty"`
	AutoCollection                    HostedPageSubscriptionAutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              HostedPageSubscriptionOfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	InvoiceNotes                      string                                     `json:"invoice_notes,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                     `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutExistingAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutExistingEventBasedAddon struct {
	Id                  string                            `json:"id,omitempty"`
	Quantity            *int32                            `json:"quantity,omitempty"`
	UnitPrice           *int64                            `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32                            `json:"service_period_in_days,omitempty"`
	ChargeOn            HostedPageEventBasedAddonChargeOn `json:"charge_on,omitempty"`
	OnEvent             HostedPageEventBasedAddonOnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool                             `json:"charge_once,omitempty"`
	QuantityInDecimal   string                            `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string                            `json:"unit_price_in_decimal,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutExistingCustomer struct {
	VatNumber       string `json:"vat_number,omitempty"`
	VatNumberPrefix string `json:"vat_number_prefix,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutExistingCard struct {
	Gateway          HostedPageCardGateway `json:"gateway,omitempty"`
	GatewayAccountId string                `json:"gateway_account_id,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutExistingContractTerm struct {
	ActionAtTermEnd          HostedPageContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                                `json:"cancellation_cutoff_period,omitempty"`
}

type HostedPageCheckoutExistingForItemsRequest struct {
	Layout                     HostedPageLayout                                      `json:"layout,omitempty"`
	Subscription               *HostedPageCheckoutExistingForItemsSubscription       `json:"subscription,omitempty"`
	SubscriptionItems          []*HostedPageCheckoutExistingForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	MandatoryItemsToRemove     []string                                              `json:"mandatory_items_to_remove,omitempty"`
	ReplaceItemsList           *bool                                                 `json:"replace_items_list,omitempty"`
	Discounts                  []*HostedPageCheckoutExistingForItemsDiscount         `json:"discounts,omitempty"`
	ItemTiers                  []*HostedPageCheckoutExistingForItemsItemTier         `json:"item_tiers,omitempty"`
	InvoiceDate                *int64                                                `json:"invoice_date,omitempty"`
	BillingCycles              *int32                                                `json:"billing_cycles,omitempty"`
	TermsToCharge              *int32                                                `json:"terms_to_charge,omitempty"`
	ReactivateFrom             *int64                                                `json:"reactivate_from,omitempty"`
	BillingAlignmentMode       HostedPageBillingAlignmentMode                        `json:"billing_alignment_mode,omitempty"`
	CouponIds                  []string                                              `json:"coupon_ids,omitempty"`
	ReplaceCouponList          *bool                                                 `json:"replace_coupon_list,omitempty"`
	Reactivate                 *bool                                                 `json:"reactivate,omitempty"`
	ForceTermReset             *bool                                                 `json:"force_term_reset,omitempty"`
	ChangeOption               HostedPageChangeOption                                `json:"change_option,omitempty"`
	ChangesScheduledAt         *int64                                                `json:"changes_scheduled_at,omitempty"`
	Customer                   *HostedPageCheckoutExistingForItemsCustomer           `json:"customer,omitempty"`
	EntityIdentifiers          []*HostedPageCheckoutExistingForItemsEntityIdentifier `json:"entity_identifiers,omitempty"`
	Card                       *HostedPageCheckoutExistingForItemsCard               `json:"card,omitempty"`
	RedirectUrl                string                                                `json:"redirect_url,omitempty"`
	CancelUrl                  string                                                `json:"cancel_url,omitempty"`
	PassThruContent            string                                                `json:"pass_thru_content,omitempty"`
	AllowOfflinePaymentMethods *bool                                                 `json:"allow_offline_payment_methods,omitempty"`
	ContractTerm               *HostedPageCheckoutExistingForItemsContractTerm       `json:"contract_term,omitempty"`
	apiRequest                 `json:"-" form:"-"`
}

func (r *HostedPageCheckoutExistingForItemsRequest) payload() any { return r }

// input sub resource params single
type HostedPageCheckoutExistingForItemsSubscription struct {
	Id                                string                                     `json:"id"`
	SetupFee                          *int64                                     `json:"setup_fee,omitempty"`
	StartDate                         *int64                                     `json:"start_date,omitempty"`
	TrialEnd                          *int64                                     `json:"trial_end,omitempty"`
	Coupon                            string                                     `json:"coupon,omitempty"`
	AutoCollection                    HostedPageSubscriptionAutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              HostedPageSubscriptionOfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	InvoiceNotes                      string                                     `json:"invoice_notes,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                     `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutExistingForItemsSubscriptionItem struct {
	ItemPriceId        string                                   `json:"item_price_id"`
	Quantity           *int32                                   `json:"quantity,omitempty"`
	QuantityInDecimal  string                                   `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64                                   `json:"unit_price,omitempty"`
	UnitPriceInDecimal string                                   `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32                                   `json:"billing_cycles,omitempty"`
	TrialEnd           *int64                                   `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32                                   `json:"service_period_days,omitempty"`
	ChargeOnEvent      HostedPageSubscriptionItemChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool                                    `json:"charge_once,omitempty"`
	ChargeOnOption     HostedPageSubscriptionItemChargeOnOption `json:"charge_on_option,omitempty"`
	ItemType           HostedPageSubscriptionItemItemType       `json:"item_type,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutExistingForItemsDiscount struct {
	ApplyOn       HostedPageDiscountApplyOn       `json:"apply_on,omitempty"`
	DurationType  HostedPageDiscountDurationType  `json:"duration_type"`
	Percentage    *float64                        `json:"percentage,omitempty"`
	Amount        *int64                          `json:"amount,omitempty"`
	Period        *int32                          `json:"period,omitempty"`
	PeriodUnit    HostedPageDiscountPeriodUnit    `json:"period_unit,omitempty"`
	IncludedInMrr *bool                           `json:"included_in_mrr,omitempty"`
	ItemPriceId   string                          `json:"item_price_id,omitempty"`
	Quantity      *int32                          `json:"quantity,omitempty"`
	OperationType HostedPageDiscountOperationType `json:"operation_type"`
	Id            string                          `json:"id,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutExistingForItemsItemTier struct {
	ItemPriceId           string                        `json:"item_price_id,omitempty"`
	StartingUnit          *int32                        `json:"starting_unit,omitempty"`
	EndingUnit            *int32                        `json:"ending_unit,omitempty"`
	Price                 *int64                        `json:"price,omitempty"`
	StartingUnitInDecimal string                        `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string                        `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string                        `json:"price_in_decimal,omitempty"`
	PricingType           HostedPageItemTierPricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32                        `json:"package_size,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutExistingForItemsCustomer struct {
	VatNumber                string `json:"vat_number,omitempty"`
	VatNumberPrefix          string `json:"vat_number_prefix,omitempty"`
	IsEinvoiceEnabled        *bool  `json:"is_einvoice_enabled,omitempty"`
	EntityIdentifierScheme   string `json:"entity_identifier_scheme,omitempty"`
	EntityIdentifierStandard string `json:"entity_identifier_standard,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutExistingForItemsEntityIdentifier struct {
	Id        string                              `json:"id,omitempty"`
	Scheme    string                              `json:"scheme,omitempty"`
	Value     string                              `json:"value,omitempty"`
	Operation HostedPageEntityIdentifierOperation `json:"operation,omitempty"`
	Standard  string                              `json:"standard,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutExistingForItemsCard struct {
	Gateway          HostedPageCardGateway `json:"gateway,omitempty"`
	GatewayAccountId string                `json:"gateway_account_id,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutExistingForItemsContractTerm struct {
	ActionAtTermEnd          HostedPageContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                                `json:"cancellation_cutoff_period,omitempty"`
}

type HostedPageUpdateCardRequest struct {
	Customer        *HostedPageUpdateCardCustomer `json:"customer,omitempty"`
	Card            *HostedPageUpdateCardCard     `json:"card,omitempty"`
	RedirectUrl     string                        `json:"redirect_url,omitempty"`
	CancelUrl       string                        `json:"cancel_url,omitempty"`
	PassThruContent string                        `json:"pass_thru_content,omitempty"`
	Embed           *bool                         `json:"embed,omitempty"`
	IframeMessaging *bool                         `json:"iframe_messaging,omitempty"`
	apiRequest      `json:"-" form:"-"`
}

func (r *HostedPageUpdateCardRequest) payload() any { return r }

// input sub resource params single
type HostedPageUpdateCardCustomer struct {
	Id              string `json:"id"`
	VatNumber       string `json:"vat_number,omitempty"`
	VatNumberPrefix string `json:"vat_number_prefix,omitempty"`
}

// input sub resource params single
type HostedPageUpdateCardCard struct {
	Gateway          HostedPageCardGateway `json:"gateway,omitempty"`
	GatewayAccountId string                `json:"gateway_account_id,omitempty"`
}

type HostedPageUpdatePaymentMethodRequest struct {
	Customer        *HostedPageUpdatePaymentMethodCustomer `json:"customer,omitempty"`
	Card            *HostedPageUpdatePaymentMethodCard     `json:"card,omitempty"`
	RedirectUrl     string                                 `json:"redirect_url,omitempty"`
	CancelUrl       string                                 `json:"cancel_url,omitempty"`
	PassThruContent string                                 `json:"pass_thru_content,omitempty"`
	Embed           *bool                                  `json:"embed,omitempty"`
	IframeMessaging *bool                                  `json:"iframe_messaging,omitempty"`
	apiRequest      `json:"-" form:"-"`
}

func (r *HostedPageUpdatePaymentMethodRequest) payload() any { return r }

// input sub resource params single
type HostedPageUpdatePaymentMethodCustomer struct {
	Id              string `json:"id"`
	VatNumber       string `json:"vat_number,omitempty"`
	VatNumberPrefix string `json:"vat_number_prefix,omitempty"`
}

// input sub resource params single
type HostedPageUpdatePaymentMethodCard struct {
	Gateway          HostedPageCardGateway `json:"gateway,omitempty"`
	GatewayAccountId string                `json:"gateway_account_id,omitempty"`
}

type HostedPageManagePaymentSourcesRequest struct {
	Customer    *HostedPageManagePaymentSourcesCustomer `json:"customer,omitempty"`
	RedirectUrl string                                  `json:"redirect_url,omitempty"`
	Card        *HostedPageManagePaymentSourcesCard     `json:"card,omitempty"`
	apiRequest  `json:"-" form:"-"`
}

func (r *HostedPageManagePaymentSourcesRequest) payload() any { return r }

// input sub resource params single
type HostedPageManagePaymentSourcesCustomer struct {
	Id string `json:"id"`
}

// input sub resource params single
type HostedPageManagePaymentSourcesCard struct {
	Gateway          HostedPageCardGateway `json:"gateway,omitempty"`
	GatewayAccountId string                `json:"gateway_account_id,omitempty"`
}

type HostedPageCollectNowRequest struct {
	Customer     *HostedPageCollectNowCustomer `json:"customer,omitempty"`
	RedirectUrl  string                        `json:"redirect_url,omitempty"`
	Card         *HostedPageCollectNowCard     `json:"card,omitempty"`
	CurrencyCode string                        `json:"currency_code,omitempty"`
	apiRequest   `json:"-" form:"-"`
}

func (r *HostedPageCollectNowRequest) payload() any { return r }

// input sub resource params single
type HostedPageCollectNowCustomer struct {
	Id string `json:"id"`
}

// input sub resource params single
type HostedPageCollectNowCard struct {
	Gateway          HostedPageCardGateway `json:"gateway,omitempty"`
	GatewayAccountId string                `json:"gateway_account_id,omitempty"`
}

type HostedPageAcceptQuoteRequest struct {
	Quote       *HostedPageAcceptQuoteQuote `json:"quote,omitempty"`
	RedirectUrl string                      `json:"redirect_url,omitempty"`
	Layout      HostedPageLayout            `json:"layout,omitempty"`
	apiRequest  `json:"-" form:"-"`
}

func (r *HostedPageAcceptQuoteRequest) payload() any { return r }

// input sub resource params single
type HostedPageAcceptQuoteQuote struct {
	Id string `json:"id"`
}

type HostedPageExtendSubscriptionRequest struct {
	Subscription *HostedPageExtendSubscriptionSubscription `json:"subscription,omitempty"`
	Expiry       *int32                                    `json:"expiry,omitempty"`
	BillingCycle *int32                                    `json:"billing_cycle,omitempty"`
	apiRequest   `json:"-" form:"-"`
}

func (r *HostedPageExtendSubscriptionRequest) payload() any { return r }

// input sub resource params single
type HostedPageExtendSubscriptionSubscription struct {
	Id string `json:"id"`
}

type HostedPageCheckoutGiftRequest struct {
	Gifter       *HostedPageCheckoutGiftGifter       `json:"gifter,omitempty"`
	RedirectUrl  string                              `json:"redirect_url,omitempty"`
	Subscription *HostedPageCheckoutGiftSubscription `json:"subscription,omitempty"`
	Addons       []*HostedPageCheckoutGiftAddon      `json:"addons,omitempty"`
	CouponIds    []string                            `json:"coupon_ids,omitempty"`
	apiRequest   `json:"-" form:"-"`
}

func (r *HostedPageCheckoutGiftRequest) payload() any { return r }

// input sub resource params single
type HostedPageCheckoutGiftGifter struct {
	CustomerId string `json:"customer_id,omitempty"`
}

// input sub resource params single
type HostedPageCheckoutGiftSubscription struct {
	PlanId                string `json:"plan_id"`
	PlanQuantity          *int32 `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal string `json:"plan_quantity_in_decimal,omitempty"`
	Coupon                string `json:"coupon,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutGiftAddon struct {
	Id                string `json:"id,omitempty"`
	Quantity          *int32 `json:"quantity,omitempty"`
	QuantityInDecimal string `json:"quantity_in_decimal,omitempty"`
}

type HostedPageCheckoutGiftForItemsRequest struct {
	BusinessEntityId  string                                            `json:"business_entity_id,omitempty"`
	Gifter            *HostedPageCheckoutGiftForItemsGifter             `json:"gifter,omitempty"`
	RedirectUrl       string                                            `json:"redirect_url,omitempty"`
	SubscriptionItems []*HostedPageCheckoutGiftForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	ItemTiers         []*HostedPageCheckoutGiftForItemsItemTier         `json:"item_tiers,omitempty"`
	CouponIds         []string                                          `json:"coupon_ids,omitempty"`
	apiRequest        `json:"-" form:"-"`
}

func (r *HostedPageCheckoutGiftForItemsRequest) payload() any { return r }

// input sub resource params single
type HostedPageCheckoutGiftForItemsGifter struct {
	CustomerId string `json:"customer_id,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutGiftForItemsSubscriptionItem struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
}

// input sub resource params multi
type HostedPageCheckoutGiftForItemsItemTier struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int64 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}

type HostedPageClaimGiftRequest struct {
	Gift        *HostedPageClaimGiftGift     `json:"gift,omitempty"`
	RedirectUrl string                       `json:"redirect_url,omitempty"`
	Customer    *HostedPageClaimGiftCustomer `json:"customer,omitempty"`
	apiRequest  `json:"-" form:"-"`
}

func (r *HostedPageClaimGiftRequest) payload() any { return r }

// input sub resource params single
type HostedPageClaimGiftGift struct {
	Id string `json:"id"`
}

// input sub resource params single
type HostedPageClaimGiftCustomer struct {
	Locale string `json:"locale,omitempty"`
}

type HostedPageRetrieveAgreementPdfRequest struct {
	PaymentSourceId string `json:"payment_source_id"`
	apiRequest      `json:"-" form:"-"`
}

func (r *HostedPageRetrieveAgreementPdfRequest) payload() any { return r }

type HostedPageListRequest struct {
	Limit      *int32           `json:"limit,omitempty"`
	Offset     string           `json:"offset,omitempty"`
	Id         *StringFilter    `json:"id,omitempty"`
	Type       *EnumFilter      `json:"type,omitempty"`
	State      *EnumFilter      `json:"state,omitempty"`
	UpdatedAt  *TimestampFilter `json:"updated_at,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *HostedPageListRequest) payload() any { return r }

type HostedPagePreCancelRequest struct {
	Subscription    *HostedPagePreCancelSubscription `json:"subscription,omitempty"`
	PassThruContent string                           `json:"pass_thru_content,omitempty"`
	CancelUrl       string                           `json:"cancel_url,omitempty"`
	RedirectUrl     string                           `json:"redirect_url,omitempty"`
	apiRequest      `json:"-" form:"-"`
}

func (r *HostedPagePreCancelRequest) payload() any { return r }

// input sub resource params single
type HostedPagePreCancelSubscription struct {
	Id string `json:"id"`
}

type HostedPageEventsRequest struct {
	EventName  HostedPageEventName    `json:"event_name"`
	OccurredAt *int64                 `json:"occurred_at,omitempty"`
	EventData  map[string]interface{} `json:"event_data"`
	apiRequest `json:"-" form:"-"`
}

func (r *HostedPageEventsRequest) payload() any { return r }

type HostedPageViewVoucherRequest struct {
	PaymentVoucher *HostedPageViewVoucherPaymentVoucher `json:"payment_voucher,omitempty"`
	Customer       *HostedPageViewVoucherCustomer       `json:"customer,omitempty"`
	apiRequest     `json:"-" form:"-"`
}

func (r *HostedPageViewVoucherRequest) payload() any { return r }

// input sub resource params single
type HostedPageViewVoucherPaymentVoucher struct {
	Id string `json:"id"`
}

// input sub resource params single
type HostedPageViewVoucherCustomer struct {
	Locale string `json:"locale,omitempty"`
}

// operation response
type HostedPageCheckoutNewResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}

// operation response
type HostedPageCheckoutOneTimeResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}

// operation response
type HostedPageCheckoutOneTimeForItemsResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}

// operation response
type HostedPageCheckoutNewForItemsResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}

// operation response
type HostedPageCheckoutExistingResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}

// operation response
type HostedPageCheckoutExistingForItemsResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}

// operation response
type HostedPageUpdateCardResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}

// operation response
type HostedPageUpdatePaymentMethodResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}

// operation response
type HostedPageManagePaymentSourcesResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}

// operation response
type HostedPageCollectNowResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}

// operation response
type HostedPageAcceptQuoteResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}

// operation response
type HostedPageExtendSubscriptionResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}

// operation response
type HostedPageCheckoutGiftResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}

// operation response
type HostedPageCheckoutGiftForItemsResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}

// operation response
type HostedPageClaimGiftResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}

// operation response
type HostedPageRetrieveAgreementPdfResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}

// operation response
type HostedPageAcknowledgeResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}

// operation response
type HostedPageRetrieveResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}

// operation sub response
type HostedPageListHostedPageResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

// operation response
type HostedPageListResponse struct {
	List       []*HostedPageListHostedPageResponse `json:"list,omitempty"`
	NextOffset string                              `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type HostedPagePreCancelResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}

// operation response
type HostedPageEventsResponse struct {
	Success bool `json:"success,omitempty"`
	apiResponse
}

// operation response
type HostedPageViewVoucherResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
	apiResponse
}
