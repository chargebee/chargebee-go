package chargebee

type Type string

const (
	TypeCheckoutNew          Type = "checkout_new"
	TypeCheckoutExisting     Type = "checkout_existing"
	TypeUpdatePaymentMethod  Type = "update_payment_method"
	TypeManagePaymentSources Type = "manage_payment_sources"
	TypeCollectNow           Type = "collect_now"
	TypeExtendSubscription   Type = "extend_subscription"
	TypeCheckoutOneTime      Type = "checkout_one_time"
	TypePreCancel            Type = "pre_cancel"
	TypeViewVoucher          Type = "view_voucher"
	TypeAcceptQuote          Type = "accept_quote"
	TypeCheckoutGift         Type = "checkout_gift"
	TypeClaimGift            Type = "claim_gift"
	TypeUpdateCard           Type = "update_card"
)

type State string

const (
	StateCreated      State = "created"
	StateRequested    State = "requested"
	StateSucceeded    State = "succeeded"
	StateCancelled    State = "cancelled"
	StateAcknowledged State = "acknowledged"
	StateFailed       State = "failed"
)

type FailureReason string

const (
	FailureReasonCardError   FailureReason = "card_error"
	FailureReasonServerError FailureReason = "server_error"
)

type HostedPage struct {
	Id               string          `json:"id"`
	Type             Type            `json:"type"`
	Url              string          `json:"url"`
	State            State           `json:"state"`
	FailureReason    FailureReason   `json:"failure_reason"`
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
type CheckoutNewRequest struct {
	Subscription               *CheckoutNewSubscription      `json:"subscription,omitempty"`
	Customer                   *CheckoutNewCustomer          `json:"customer,omitempty"`
	BillingCycles              *int32                        `json:"billing_cycles,omitempty"`
	Addons                     []*CheckoutNewAddon           `json:"addons,omitempty"`
	EventBasedAddons           []*CheckoutNewEventBasedAddon `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove    []string                      `json:"mandatory_addons_to_remove,omitempty"`
	TermsToCharge              *int32                        `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode       enum.BillingAlignmentMode     `json:"billing_alignment_mode,omitempty"`
	CouponIds                  []string                      `json:"coupon_ids,omitempty"`
	Card                       *CheckoutNewCard              `json:"card,omitempty"`
	RedirectUrl                string                        `json:"redirect_url,omitempty"`
	CancelUrl                  string                        `json:"cancel_url,omitempty"`
	PassThruContent            string                        `json:"pass_thru_content,omitempty"`
	Embed                      *bool                         `json:"embed,omitempty"`
	IframeMessaging            *bool                         `json:"iframe_messaging,omitempty"`
	AllowOfflinePaymentMethods *bool                         `json:"allow_offline_payment_methods,omitempty"`
	BillingAddress             *CheckoutNewBillingAddress    `json:"billing_address,omitempty"`
	ShippingAddress            *CheckoutNewShippingAddress   `json:"shipping_address,omitempty"`
	ContractTerm               *CheckoutNewContractTerm      `json:"contract_term,omitempty"`
}
type CheckoutNewSubscription struct {
	Id                                string                    `json:"id,omitempty"`
	PlanId                            string                    `json:"plan_id"`
	PlanQuantity                      *int32                    `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                    `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int64                    `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                    `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int64                    `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	Coupon                            string                    `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	InvoiceNotes                      string                    `json:"invoice_notes,omitempty"`
	AffiliateToken                    string                    `json:"affiliate_token,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type CheckoutNewCustomer struct {
	Id                    string          `json:"id,omitempty"`
	Email                 string          `json:"email,omitempty"`
	FirstName             string          `json:"first_name,omitempty"`
	LastName              string          `json:"last_name,omitempty"`
	Company               string          `json:"company,omitempty"`
	Phone                 string          `json:"phone,omitempty"`
	Locale                string          `json:"locale,omitempty"`
	Taxability            enum.Taxability `json:"taxability,omitempty"`
	VatNumber             string          `json:"vat_number,omitempty"`
	VatNumberPrefix       string          `json:"vat_number_prefix,omitempty"`
	ConsolidatedInvoicing *bool           `json:"consolidated_invoicing,omitempty"`
}
type CheckoutNewAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
}
type CheckoutNewEventBasedAddon struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int64        `json:"unit_price,omitempty"`
	QuantityInDecimal   string        `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string        `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
}
type CheckoutNewCard struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type CheckoutNewBillingAddress struct {
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
type CheckoutNewShippingAddress struct {
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
type CheckoutNewContractTerm struct {
	ActionAtTermEnd          contractTerm.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                       `json:"cancellation_cutoff_period,omitempty"`
}
type CheckoutOneTimeRequest struct {
	Customer        *CheckoutOneTimeCustomer        `json:"customer,omitempty"`
	Addons          []*CheckoutOneTimeAddon         `json:"addons,omitempty"`
	CurrencyCode    string                          `json:"currency_code,omitempty"`
	Charges         []*CheckoutOneTimeCharge        `json:"charges,omitempty"`
	InvoiceNote     string                          `json:"invoice_note,omitempty"`
	Invoice         *CheckoutOneTimeInvoice         `json:"invoice,omitempty"`
	Coupon          string                          `json:"coupon,omitempty"`
	CouponIds       []string                        `json:"coupon_ids,omitempty"`
	Card            *CheckoutOneTimeCard            `json:"card,omitempty"`
	RedirectUrl     string                          `json:"redirect_url,omitempty"`
	CancelUrl       string                          `json:"cancel_url,omitempty"`
	PassThruContent string                          `json:"pass_thru_content,omitempty"`
	Embed           *bool                           `json:"embed,omitempty"`
	IframeMessaging *bool                           `json:"iframe_messaging,omitempty"`
	BillingAddress  *CheckoutOneTimeBillingAddress  `json:"billing_address,omitempty"`
	ShippingAddress *CheckoutOneTimeShippingAddress `json:"shipping_address,omitempty"`
}
type CheckoutOneTimeCustomer struct {
	Id                    string          `json:"id,omitempty"`
	Email                 string          `json:"email,omitempty"`
	FirstName             string          `json:"first_name,omitempty"`
	LastName              string          `json:"last_name,omitempty"`
	Company               string          `json:"company,omitempty"`
	Phone                 string          `json:"phone,omitempty"`
	Locale                string          `json:"locale,omitempty"`
	Taxability            enum.Taxability `json:"taxability,omitempty"`
	VatNumber             string          `json:"vat_number,omitempty"`
	VatNumberPrefix       string          `json:"vat_number_prefix,omitempty"`
	ConsolidatedInvoicing *bool           `json:"consolidated_invoicing,omitempty"`
}
type CheckoutOneTimeAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type CheckoutOneTimeCharge struct {
	Amount                 *int64               `json:"amount,omitempty"`
	AmountInDecimal        string               `json:"amount_in_decimal,omitempty"`
	Description            string               `json:"description,omitempty"`
	Taxable                *bool                `json:"taxable,omitempty"`
	TaxProfileId           string               `json:"tax_profile_id,omitempty"`
	AvalaraTaxCode         string               `json:"avalara_tax_code,omitempty"`
	HsnCode                string               `json:"hsn_code,omitempty"`
	TaxjarProductCode      string               `json:"taxjar_product_code,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	DateFrom               *int64               `json:"date_from,omitempty"`
	DateTo                 *int64               `json:"date_to,omitempty"`
}
type CheckoutOneTimeInvoice struct {
	PoNumber string `json:"po_number,omitempty"`
}
type CheckoutOneTimeCard struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type CheckoutOneTimeBillingAddress struct {
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
type CheckoutOneTimeShippingAddress struct {
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
type CheckoutOneTimeForItemsRequest struct {
	BusinessEntityId  string                                     `json:"business_entity_id,omitempty"`
	Layout            enum.Layout                                `json:"layout,omitempty"`
	Customer          *CheckoutOneTimeForItemsCustomer           `json:"customer,omitempty"`
	ItemPrices        []*CheckoutOneTimeForItemsItemPrice        `json:"item_prices,omitempty"`
	ItemTiers         []*CheckoutOneTimeForItemsItemTier         `json:"item_tiers,omitempty"`
	Charges           []*CheckoutOneTimeForItemsCharge           `json:"charges,omitempty"`
	Discounts         []*CheckoutOneTimeForItemsDiscount         `json:"discounts,omitempty"`
	InvoiceNote       string                                     `json:"invoice_note,omitempty"`
	Invoice           *CheckoutOneTimeForItemsInvoice            `json:"invoice,omitempty"`
	Coupon            string                                     `json:"coupon,omitempty"`
	CouponIds         []string                                   `json:"coupon_ids,omitempty"`
	CurrencyCode      string                                     `json:"currency_code,omitempty"`
	Card              *CheckoutOneTimeForItemsCard               `json:"card,omitempty"`
	EntityIdentifiers []*CheckoutOneTimeForItemsEntityIdentifier `json:"entity_identifiers,omitempty"`
	RedirectUrl       string                                     `json:"redirect_url,omitempty"`
	CancelUrl         string                                     `json:"cancel_url,omitempty"`
	PassThruContent   string                                     `json:"pass_thru_content,omitempty"`
	BillingAddress    *CheckoutOneTimeForItemsBillingAddress     `json:"billing_address,omitempty"`
	ShippingAddress   *CheckoutOneTimeForItemsShippingAddress    `json:"shipping_address,omitempty"`
}
type CheckoutOneTimeForItemsCustomer struct {
	Id                       string                `json:"id,omitempty"`
	Email                    string                `json:"email,omitempty"`
	FirstName                string                `json:"first_name,omitempty"`
	LastName                 string                `json:"last_name,omitempty"`
	Company                  string                `json:"company,omitempty"`
	Phone                    string                `json:"phone,omitempty"`
	Locale                   string                `json:"locale,omitempty"`
	Taxability               enum.Taxability       `json:"taxability,omitempty"`
	VatNumber                string                `json:"vat_number,omitempty"`
	VatNumberPrefix          string                `json:"vat_number_prefix,omitempty"`
	EinvoicingMethod         enum.EinvoicingMethod `json:"einvoicing_method,omitempty"`
	IsEinvoiceEnabled        *bool                 `json:"is_einvoice_enabled,omitempty"`
	EntityIdentifierScheme   string                `json:"entity_identifier_scheme,omitempty"`
	EntityIdentifierStandard string                `json:"entity_identifier_standard,omitempty"`
	ConsolidatedInvoicing    *bool                 `json:"consolidated_invoicing,omitempty"`
}
type CheckoutOneTimeForItemsItemPrice struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type CheckoutOneTimeForItemsItemTier struct {
	ItemPriceId           string           `json:"item_price_id,omitempty"`
	StartingUnit          *int32           `json:"starting_unit,omitempty"`
	EndingUnit            *int32           `json:"ending_unit,omitempty"`
	Price                 *int64           `json:"price,omitempty"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string           `json:"price_in_decimal,omitempty"`
	PricingType           enum.PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32           `json:"package_size,omitempty"`
}
type CheckoutOneTimeForItemsCharge struct {
	Amount                 *int64               `json:"amount,omitempty"`
	AmountInDecimal        string               `json:"amount_in_decimal,omitempty"`
	Description            string               `json:"description,omitempty"`
	Taxable                *bool                `json:"taxable,omitempty"`
	TaxProfileId           string               `json:"tax_profile_id,omitempty"`
	AvalaraTaxCode         string               `json:"avalara_tax_code,omitempty"`
	HsnCode                string               `json:"hsn_code,omitempty"`
	TaxjarProductCode      string               `json:"taxjar_product_code,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	DateFrom               *int64               `json:"date_from,omitempty"`
	DateTo                 *int64               `json:"date_to,omitempty"`
}
type CheckoutOneTimeForItemsDiscount struct {
	Percentage  *float64     `json:"percentage,omitempty"`
	Amount      *int64       `json:"amount,omitempty"`
	Quantity    *int32       `json:"quantity,omitempty"`
	ApplyOn     enum.ApplyOn `json:"apply_on"`
	ItemPriceId string       `json:"item_price_id,omitempty"`
}
type CheckoutOneTimeForItemsInvoice struct {
	PoNumber string `json:"po_number,omitempty"`
}
type CheckoutOneTimeForItemsCard struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type CheckoutOneTimeForItemsEntityIdentifier struct {
	Id        string         `json:"id,omitempty"`
	Scheme    string         `json:"scheme,omitempty"`
	Value     string         `json:"value,omitempty"`
	Operation enum.Operation `json:"operation,omitempty"`
	Standard  string         `json:"standard,omitempty"`
}
type CheckoutOneTimeForItemsBillingAddress struct {
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
type CheckoutOneTimeForItemsShippingAddress struct {
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
type CheckoutNewForItemsRequest struct {
	Subscription               *CheckoutNewForItemsSubscription       `json:"subscription,omitempty"`
	Layout                     enum.Layout                            `json:"layout,omitempty"`
	Customer                   *CheckoutNewForItemsCustomer           `json:"customer,omitempty"`
	BusinessEntityId           string                                 `json:"business_entity_id,omitempty"`
	BillingCycles              *int32                                 `json:"billing_cycles,omitempty"`
	SubscriptionItems          []*CheckoutNewForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	Discounts                  []*CheckoutNewForItemsDiscount         `json:"discounts,omitempty"`
	MandatoryItemsToRemove     []string                               `json:"mandatory_items_to_remove,omitempty"`
	ItemTiers                  []*CheckoutNewForItemsItemTier         `json:"item_tiers,omitempty"`
	TermsToCharge              *int32                                 `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode       enum.BillingAlignmentMode              `json:"billing_alignment_mode,omitempty"`
	CouponIds                  []string                               `json:"coupon_ids,omitempty"`
	Card                       *CheckoutNewForItemsCard               `json:"card,omitempty"`
	EntityIdentifiers          []*CheckoutNewForItemsEntityIdentifier `json:"entity_identifiers,omitempty"`
	RedirectUrl                string                                 `json:"redirect_url,omitempty"`
	CancelUrl                  string                                 `json:"cancel_url,omitempty"`
	PassThruContent            string                                 `json:"pass_thru_content,omitempty"`
	AllowOfflinePaymentMethods *bool                                  `json:"allow_offline_payment_methods,omitempty"`
	BillingAddress             *CheckoutNewForItemsBillingAddress     `json:"billing_address,omitempty"`
	ShippingAddress            *CheckoutNewForItemsShippingAddress    `json:"shipping_address,omitempty"`
	ContractTerm               *CheckoutNewForItemsContractTerm       `json:"contract_term,omitempty"`
}
type CheckoutNewForItemsSubscription struct {
	Id                                string                    `json:"id,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	SetupFee                          *int64                    `json:"setup_fee,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	Coupon                            string                    `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	InvoiceNotes                      string                    `json:"invoice_notes,omitempty"`
	PoNumber                          string                    `json:"po_number,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type CheckoutNewForItemsCustomer struct {
	Id                       string                `json:"id,omitempty"`
	Email                    string                `json:"email,omitempty"`
	FirstName                string                `json:"first_name,omitempty"`
	LastName                 string                `json:"last_name,omitempty"`
	Company                  string                `json:"company,omitempty"`
	Phone                    string                `json:"phone,omitempty"`
	Locale                   string                `json:"locale,omitempty"`
	Taxability               enum.Taxability       `json:"taxability,omitempty"`
	VatNumber                string                `json:"vat_number,omitempty"`
	VatNumberPrefix          string                `json:"vat_number_prefix,omitempty"`
	IsEinvoiceEnabled        *bool                 `json:"is_einvoice_enabled,omitempty"`
	EntityIdentifierScheme   string                `json:"entity_identifier_scheme,omitempty"`
	EntityIdentifierStandard string                `json:"entity_identifier_standard,omitempty"`
	EinvoicingMethod         enum.EinvoicingMethod `json:"einvoicing_method,omitempty"`
}
type CheckoutNewForItemsSubscriptionItem struct {
	ItemPriceId        string              `json:"item_price_id"`
	Quantity           *int32              `json:"quantity,omitempty"`
	QuantityInDecimal  string              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32              `json:"billing_cycles,omitempty"`
	TrialEnd           *int64              `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32              `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool               `json:"charge_once,omitempty"`
	ItemType           enum.ItemType       `json:"item_type,omitempty"`
	ChargeOnOption     enum.ChargeOnOption `json:"charge_on_option,omitempty"`
}
type CheckoutNewForItemsDiscount struct {
	ApplyOn       enum.ApplyOn      `json:"apply_on,omitempty"`
	DurationType  enum.DurationType `json:"duration_type"`
	Percentage    *float64          `json:"percentage,omitempty"`
	Amount        *int64            `json:"amount,omitempty"`
	Period        *int32            `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool             `json:"included_in_mrr,omitempty"`
	ItemPriceId   string            `json:"item_price_id,omitempty"`
	Quantity      *int32            `json:"quantity,omitempty"`
}
type CheckoutNewForItemsItemTier struct {
	ItemPriceId           string           `json:"item_price_id,omitempty"`
	StartingUnit          *int32           `json:"starting_unit,omitempty"`
	EndingUnit            *int32           `json:"ending_unit,omitempty"`
	Price                 *int64           `json:"price,omitempty"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string           `json:"price_in_decimal,omitempty"`
	PricingType           enum.PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32           `json:"package_size,omitempty"`
}
type CheckoutNewForItemsCard struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type CheckoutNewForItemsEntityIdentifier struct {
	Id        string         `json:"id,omitempty"`
	Scheme    string         `json:"scheme,omitempty"`
	Value     string         `json:"value,omitempty"`
	Operation enum.Operation `json:"operation,omitempty"`
	Standard  string         `json:"standard,omitempty"`
}
type CheckoutNewForItemsBillingAddress struct {
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
type CheckoutNewForItemsShippingAddress struct {
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
type CheckoutNewForItemsContractTerm struct {
	ActionAtTermEnd          contractTerm.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                       `json:"cancellation_cutoff_period,omitempty"`
}
type CheckoutExistingRequest struct {
	Subscription               *CheckoutExistingSubscription      `json:"subscription,omitempty"`
	Addons                     []*CheckoutExistingAddon           `json:"addons,omitempty"`
	EventBasedAddons           []*CheckoutExistingEventBasedAddon `json:"event_based_addons,omitempty"`
	ReplaceAddonList           *bool                              `json:"replace_addon_list,omitempty"`
	MandatoryAddonsToRemove    []string                           `json:"mandatory_addons_to_remove,omitempty"`
	InvoiceDate                *int64                             `json:"invoice_date,omitempty"`
	BillingCycles              *int32                             `json:"billing_cycles,omitempty"`
	TermsToCharge              *int32                             `json:"terms_to_charge,omitempty"`
	ReactivateFrom             *int64                             `json:"reactivate_from,omitempty"`
	BillingAlignmentMode       enum.BillingAlignmentMode          `json:"billing_alignment_mode,omitempty"`
	CouponIds                  []string                           `json:"coupon_ids,omitempty"`
	ReplaceCouponList          *bool                              `json:"replace_coupon_list,omitempty"`
	Reactivate                 *bool                              `json:"reactivate,omitempty"`
	ForceTermReset             *bool                              `json:"force_term_reset,omitempty"`
	Customer                   *CheckoutExistingCustomer          `json:"customer,omitempty"`
	Card                       *CheckoutExistingCard              `json:"card,omitempty"`
	RedirectUrl                string                             `json:"redirect_url,omitempty"`
	CancelUrl                  string                             `json:"cancel_url,omitempty"`
	PassThruContent            string                             `json:"pass_thru_content,omitempty"`
	Embed                      *bool                              `json:"embed,omitempty"`
	IframeMessaging            *bool                              `json:"iframe_messaging,omitempty"`
	AllowOfflinePaymentMethods *bool                              `json:"allow_offline_payment_methods,omitempty"`
	ContractTerm               *CheckoutExistingContractTerm      `json:"contract_term,omitempty"`
}
type CheckoutExistingSubscription struct {
	Id                                string                    `json:"id"`
	PlanId                            string                    `json:"plan_id,omitempty"`
	PlanQuantity                      *int32                    `json:"plan_quantity,omitempty"`
	PlanUnitPrice                     *int64                    `json:"plan_unit_price,omitempty"`
	SetupFee                          *int64                    `json:"setup_fee,omitempty"`
	PlanQuantityInDecimal             string                    `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPriceInDecimal            string                    `json:"plan_unit_price_in_decimal,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	Coupon                            string                    `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	InvoiceNotes                      string                    `json:"invoice_notes,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type CheckoutExistingAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
}
type CheckoutExistingEventBasedAddon struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int64        `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
	QuantityInDecimal   string        `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string        `json:"unit_price_in_decimal,omitempty"`
}
type CheckoutExistingCustomer struct {
	VatNumber       string `json:"vat_number,omitempty"`
	VatNumberPrefix string `json:"vat_number_prefix,omitempty"`
}
type CheckoutExistingCard struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type CheckoutExistingContractTerm struct {
	ActionAtTermEnd          contractTerm.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                       `json:"cancellation_cutoff_period,omitempty"`
}
type CheckoutExistingForItemsRequest struct {
	Layout                     enum.Layout                                 `json:"layout,omitempty"`
	Subscription               *CheckoutExistingForItemsSubscription       `json:"subscription,omitempty"`
	SubscriptionItems          []*CheckoutExistingForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	MandatoryItemsToRemove     []string                                    `json:"mandatory_items_to_remove,omitempty"`
	ReplaceItemsList           *bool                                       `json:"replace_items_list,omitempty"`
	Discounts                  []*CheckoutExistingForItemsDiscount         `json:"discounts,omitempty"`
	ItemTiers                  []*CheckoutExistingForItemsItemTier         `json:"item_tiers,omitempty"`
	InvoiceDate                *int64                                      `json:"invoice_date,omitempty"`
	BillingCycles              *int32                                      `json:"billing_cycles,omitempty"`
	TermsToCharge              *int32                                      `json:"terms_to_charge,omitempty"`
	ReactivateFrom             *int64                                      `json:"reactivate_from,omitempty"`
	BillingAlignmentMode       enum.BillingAlignmentMode                   `json:"billing_alignment_mode,omitempty"`
	CouponIds                  []string                                    `json:"coupon_ids,omitempty"`
	ReplaceCouponList          *bool                                       `json:"replace_coupon_list,omitempty"`
	Reactivate                 *bool                                       `json:"reactivate,omitempty"`
	ForceTermReset             *bool                                       `json:"force_term_reset,omitempty"`
	ChangeOption               enum.ChangeOption                           `json:"change_option,omitempty"`
	ChangesScheduledAt         *int64                                      `json:"changes_scheduled_at,omitempty"`
	Customer                   *CheckoutExistingForItemsCustomer           `json:"customer,omitempty"`
	EntityIdentifiers          []*CheckoutExistingForItemsEntityIdentifier `json:"entity_identifiers,omitempty"`
	Card                       *CheckoutExistingForItemsCard               `json:"card,omitempty"`
	RedirectUrl                string                                      `json:"redirect_url,omitempty"`
	CancelUrl                  string                                      `json:"cancel_url,omitempty"`
	PassThruContent            string                                      `json:"pass_thru_content,omitempty"`
	AllowOfflinePaymentMethods *bool                                       `json:"allow_offline_payment_methods,omitempty"`
	ContractTerm               *CheckoutExistingForItemsContractTerm       `json:"contract_term,omitempty"`
}
type CheckoutExistingForItemsSubscription struct {
	Id                                string                    `json:"id"`
	SetupFee                          *int64                    `json:"setup_fee,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	Coupon                            string                    `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	InvoiceNotes                      string                    `json:"invoice_notes,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type CheckoutExistingForItemsSubscriptionItem struct {
	ItemPriceId        string              `json:"item_price_id"`
	Quantity           *int32              `json:"quantity,omitempty"`
	QuantityInDecimal  string              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32              `json:"billing_cycles,omitempty"`
	TrialEnd           *int64              `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32              `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool               `json:"charge_once,omitempty"`
	ChargeOnOption     enum.ChargeOnOption `json:"charge_on_option,omitempty"`
	ItemType           enum.ItemType       `json:"item_type,omitempty"`
}
type CheckoutExistingForItemsDiscount struct {
	ApplyOn       enum.ApplyOn       `json:"apply_on,omitempty"`
	DurationType  enum.DurationType  `json:"duration_type"`
	Percentage    *float64           `json:"percentage,omitempty"`
	Amount        *int64             `json:"amount,omitempty"`
	Period        *int32             `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit    `json:"period_unit,omitempty"`
	IncludedInMrr *bool              `json:"included_in_mrr,omitempty"`
	ItemPriceId   string             `json:"item_price_id,omitempty"`
	Quantity      *int32             `json:"quantity,omitempty"`
	OperationType enum.OperationType `json:"operation_type"`
	Id            string             `json:"id,omitempty"`
}
type CheckoutExistingForItemsItemTier struct {
	ItemPriceId           string           `json:"item_price_id,omitempty"`
	StartingUnit          *int32           `json:"starting_unit,omitempty"`
	EndingUnit            *int32           `json:"ending_unit,omitempty"`
	Price                 *int64           `json:"price,omitempty"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string           `json:"price_in_decimal,omitempty"`
	PricingType           enum.PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32           `json:"package_size,omitempty"`
}
type CheckoutExistingForItemsCustomer struct {
	VatNumber                string `json:"vat_number,omitempty"`
	VatNumberPrefix          string `json:"vat_number_prefix,omitempty"`
	IsEinvoiceEnabled        *bool  `json:"is_einvoice_enabled,omitempty"`
	EntityIdentifierScheme   string `json:"entity_identifier_scheme,omitempty"`
	EntityIdentifierStandard string `json:"entity_identifier_standard,omitempty"`
}
type CheckoutExistingForItemsEntityIdentifier struct {
	Id        string         `json:"id,omitempty"`
	Scheme    string         `json:"scheme,omitempty"`
	Value     string         `json:"value,omitempty"`
	Operation enum.Operation `json:"operation,omitempty"`
	Standard  string         `json:"standard,omitempty"`
}
type CheckoutExistingForItemsCard struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type CheckoutExistingForItemsContractTerm struct {
	ActionAtTermEnd          contractTerm.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                       `json:"cancellation_cutoff_period,omitempty"`
}
type UpdateCardRequest struct {
	Customer        *UpdateCardCustomer `json:"customer,omitempty"`
	Card            *UpdateCardCard     `json:"card,omitempty"`
	RedirectUrl     string              `json:"redirect_url,omitempty"`
	CancelUrl       string              `json:"cancel_url,omitempty"`
	PassThruContent string              `json:"pass_thru_content,omitempty"`
	Embed           *bool               `json:"embed,omitempty"`
	IframeMessaging *bool               `json:"iframe_messaging,omitempty"`
}
type UpdateCardCustomer struct {
	Id              string `json:"id"`
	VatNumber       string `json:"vat_number,omitempty"`
	VatNumberPrefix string `json:"vat_number_prefix,omitempty"`
}
type UpdateCardCard struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type UpdatePaymentMethodRequest struct {
	Customer        *UpdatePaymentMethodCustomer `json:"customer,omitempty"`
	Card            *UpdatePaymentMethodCard     `json:"card,omitempty"`
	RedirectUrl     string                       `json:"redirect_url,omitempty"`
	CancelUrl       string                       `json:"cancel_url,omitempty"`
	PassThruContent string                       `json:"pass_thru_content,omitempty"`
	Embed           *bool                        `json:"embed,omitempty"`
	IframeMessaging *bool                        `json:"iframe_messaging,omitempty"`
}
type UpdatePaymentMethodCustomer struct {
	Id              string `json:"id"`
	VatNumber       string `json:"vat_number,omitempty"`
	VatNumberPrefix string `json:"vat_number_prefix,omitempty"`
}
type UpdatePaymentMethodCard struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type ManagePaymentSourcesRequest struct {
	Customer    *ManagePaymentSourcesCustomer `json:"customer,omitempty"`
	RedirectUrl string                        `json:"redirect_url,omitempty"`
	Card        *ManagePaymentSourcesCard     `json:"card,omitempty"`
}
type ManagePaymentSourcesCustomer struct {
	Id string `json:"id"`
}
type ManagePaymentSourcesCard struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type CollectNowRequest struct {
	Customer     *CollectNowCustomer `json:"customer,omitempty"`
	RedirectUrl  string              `json:"redirect_url,omitempty"`
	Card         *CollectNowCard     `json:"card,omitempty"`
	CurrencyCode string              `json:"currency_code,omitempty"`
}
type CollectNowCustomer struct {
	Id string `json:"id"`
}
type CollectNowCard struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type AcceptQuoteRequest struct {
	Quote       *AcceptQuoteQuote `json:"quote,omitempty"`
	RedirectUrl string            `json:"redirect_url,omitempty"`
	Layout      enum.Layout       `json:"layout,omitempty"`
}
type AcceptQuoteQuote struct {
	Id string `json:"id"`
}
type ExtendSubscriptionRequest struct {
	Subscription *ExtendSubscriptionSubscription `json:"subscription,omitempty"`
	Expiry       *int32                          `json:"expiry,omitempty"`
	BillingCycle *int32                          `json:"billing_cycle,omitempty"`
}
type ExtendSubscriptionSubscription struct {
	Id string `json:"id"`
}
type CheckoutGiftRequest struct {
	Gifter       *CheckoutGiftGifter       `json:"gifter,omitempty"`
	RedirectUrl  string                    `json:"redirect_url,omitempty"`
	Subscription *CheckoutGiftSubscription `json:"subscription,omitempty"`
	Addons       []*CheckoutGiftAddon      `json:"addons,omitempty"`
	CouponIds    []string                  `json:"coupon_ids,omitempty"`
}
type CheckoutGiftGifter struct {
	CustomerId string `json:"customer_id,omitempty"`
}
type CheckoutGiftSubscription struct {
	PlanId                string `json:"plan_id"`
	PlanQuantity          *int32 `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal string `json:"plan_quantity_in_decimal,omitempty"`
	Coupon                string `json:"coupon,omitempty"`
}
type CheckoutGiftAddon struct {
	Id                string `json:"id,omitempty"`
	Quantity          *int32 `json:"quantity,omitempty"`
	QuantityInDecimal string `json:"quantity_in_decimal,omitempty"`
}
type CheckoutGiftForItemsRequest struct {
	BusinessEntityId  string                                  `json:"business_entity_id,omitempty"`
	Gifter            *CheckoutGiftForItemsGifter             `json:"gifter,omitempty"`
	RedirectUrl       string                                  `json:"redirect_url,omitempty"`
	SubscriptionItems []*CheckoutGiftForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	ItemTiers         []*CheckoutGiftForItemsItemTier         `json:"item_tiers,omitempty"`
	CouponIds         []string                                `json:"coupon_ids,omitempty"`
}
type CheckoutGiftForItemsGifter struct {
	CustomerId string `json:"customer_id,omitempty"`
}
type CheckoutGiftForItemsSubscriptionItem struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
}
type CheckoutGiftForItemsItemTier struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int64 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type ClaimGiftRequest struct {
	Gift        *ClaimGiftGift     `json:"gift,omitempty"`
	RedirectUrl string             `json:"redirect_url,omitempty"`
	Customer    *ClaimGiftCustomer `json:"customer,omitempty"`
}
type ClaimGiftGift struct {
	Id string `json:"id"`
}
type ClaimGiftCustomer struct {
	Locale string `json:"locale,omitempty"`
}
type RetrieveAgreementPdfRequest struct {
	PaymentSourceId string `json:"payment_source_id"`
}
type ListRequest struct {
	Limit     *int32                  `json:"limit,omitempty"`
	Offset    string                  `json:"offset,omitempty"`
	Id        *filter.StringFilter    `json:"id,omitempty"`
	Type      *filter.EnumFilter      `json:"type,omitempty"`
	State     *filter.EnumFilter      `json:"state,omitempty"`
	UpdatedAt *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type PreCancelRequest struct {
	Subscription    *PreCancelSubscription `json:"subscription,omitempty"`
	PassThruContent string                 `json:"pass_thru_content,omitempty"`
	CancelUrl       string                 `json:"cancel_url,omitempty"`
	RedirectUrl     string                 `json:"redirect_url,omitempty"`
}
type PreCancelSubscription struct {
	Id string `json:"id"`
}
type EventsRequest struct {
	EventName  enum.EventName         `json:"event_name"`
	OccurredAt *int64                 `json:"occurred_at,omitempty"`
	EventData  map[string]interface{} `json:"event_data"`
}
type ViewVoucherRequest struct {
	PaymentVoucher *ViewVoucherPaymentVoucher `json:"payment_voucher,omitempty"`
	Customer       *ViewVoucherCustomer       `json:"customer,omitempty"`
}
type ViewVoucherPaymentVoucher struct {
	Id string `json:"id"`
}
type ViewVoucherCustomer struct {
	Locale string `json:"locale,omitempty"`
}

type CheckoutNewResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type CheckoutOneTimeResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type CheckoutOneTimeForItemsResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type CheckoutNewForItemsResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type CheckoutExistingResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type CheckoutExistingForItemsResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type UpdateCardResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type UpdatePaymentMethodResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type ManagePaymentSourcesResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type CollectNowResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type AcceptQuoteResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type ExtendSubscriptionResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type CheckoutGiftResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type CheckoutGiftForItemsResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type ClaimGiftResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type RetrieveAgreementPdfResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type AcknowledgeResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type RetrieveResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type ListHostedPageResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type ListResponse struct {
	List       []*ListHostedPageResponse `json:"list,omitempty"`
	NextOffset string                    `json:"next_offset,omitempty"`
}

type PreCancelResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}

type EventsResponse struct {
	Success bool `json:"success,omitempty"`
}

type ViewVoucherResponse struct {
	HostedPage *HostedPage `json:"hosted_page,omitempty"`
}
