package hostedpage

import (
	"encoding/json"

	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	contractTermEnum "github.com/chargebee/chargebee-go/models/contractterm/enum"
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
	Subscription               *CheckoutNewSubscriptionParams      `json:"subscription,omitempty"`
	Customer                   *CheckoutNewCustomerParams          `json:"customer,omitempty"`
	BillingCycles              *int32                              `json:"billing_cycles,omitempty"`
	Addons                     []*CheckoutNewAddonParams           `json:"addons,omitempty"`
	EventBasedAddons           []*CheckoutNewEventBasedAddonParams `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove    []string                            `json:"mandatory_addons_to_remove,omitempty"`
	TermsToCharge              *int32                              `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode       enum.BillingAlignmentMode           `json:"billing_alignment_mode,omitempty"`
	CouponIds                  []string                            `json:"coupon_ids,omitempty"`
	Card                       *CheckoutNewCardParams              `json:"card,omitempty"`
	RedirectUrl                string                              `json:"redirect_url,omitempty"`
	CancelUrl                  string                              `json:"cancel_url,omitempty"`
	PassThruContent            string                              `json:"pass_thru_content,omitempty"`
	Embed                      *bool                               `json:"embed,omitempty"`
	IframeMessaging            *bool                               `json:"iframe_messaging,omitempty"`
	AllowOfflinePaymentMethods *bool                               `json:"allow_offline_payment_methods,omitempty"`
	BillingAddress             *CheckoutNewBillingAddressParams    `json:"billing_address,omitempty"`
	ShippingAddress            *CheckoutNewShippingAddressParams   `json:"shipping_address,omitempty"`
	ContractTerm               *CheckoutNewContractTermParams      `json:"contract_term,omitempty"`
}
type CheckoutNewSubscriptionParams struct {
	Id                                string                    `json:"id,omitempty"`
	PlanId                            string                    `json:"plan_id"`
	PlanQuantity                      *int32                    `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                    `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int32                    `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                    `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int32                    `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	Coupon                            string                    `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	InvoiceNotes                      string                    `json:"invoice_notes,omitempty"`
	AffiliateToken                    string                    `json:"affiliate_token,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
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
	VatNumberPrefix       string          `json:"vat_number_prefix,omitempty"`
	ConsolidatedInvoicing *bool           `json:"consolidated_invoicing,omitempty"`
}
type CheckoutNewAddonParams struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
}
type CheckoutNewEventBasedAddonParams struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int32        `json:"unit_price,omitempty"`
	QuantityInDecimal   string        `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string        `json:"unit_price_in_decimal,omitempty"`
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
type CheckoutNewContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type CheckoutOneTimeRequestParams struct {
	Customer        *CheckoutOneTimeCustomerParams        `json:"customer,omitempty"`
	Addons          []*CheckoutOneTimeAddonParams         `json:"addons,omitempty"`
	CurrencyCode    string                                `json:"currency_code,omitempty"`
	Charges         []*CheckoutOneTimeChargeParams        `json:"charges,omitempty"`
	InvoiceNote     string                                `json:"invoice_note,omitempty"`
	Invoice         *CheckoutOneTimeInvoiceParams         `json:"invoice,omitempty"`
	Coupon          string                                `json:"coupon,omitempty"`
	CouponIds       []string                              `json:"coupon_ids,omitempty"`
	Card            *CheckoutOneTimeCardParams            `json:"card,omitempty"`
	RedirectUrl     string                                `json:"redirect_url,omitempty"`
	CancelUrl       string                                `json:"cancel_url,omitempty"`
	PassThruContent string                                `json:"pass_thru_content,omitempty"`
	Embed           *bool                                 `json:"embed,omitempty"`
	IframeMessaging *bool                                 `json:"iframe_messaging,omitempty"`
	BillingAddress  *CheckoutOneTimeBillingAddressParams  `json:"billing_address,omitempty"`
	ShippingAddress *CheckoutOneTimeShippingAddressParams `json:"shipping_address,omitempty"`
}
type CheckoutOneTimeCustomerParams struct {
	Id                    string          `json:"id,omitempty"`
	Email                 string          `json:"email,omitempty"`
	FirstName             string          `json:"first_name,omitempty"`
	LastName              string          `json:"last_name,omitempty"`
	Company               string          `json:"company,omitempty"`
	Taxability            enum.Taxability `json:"taxability,omitempty"`
	Locale                string          `json:"locale,omitempty"`
	Phone                 string          `json:"phone,omitempty"`
	VatNumber             string          `json:"vat_number,omitempty"`
	VatNumberPrefix       string          `json:"vat_number_prefix,omitempty"`
	ConsolidatedInvoicing *bool           `json:"consolidated_invoicing,omitempty"`
}
type CheckoutOneTimeAddonParams struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type CheckoutOneTimeChargeParams struct {
	Amount                 *int32               `json:"amount,omitempty"`
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
type CheckoutOneTimeInvoiceParams struct {
	PoNumber string `json:"po_number,omitempty"`
}
type CheckoutOneTimeCardParams struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type CheckoutOneTimeBillingAddressParams struct {
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
type CheckoutOneTimeShippingAddressParams struct {
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
type CheckoutOneTimeForItemsRequestParams struct {
	Customer          *CheckoutOneTimeForItemsCustomerParams           `json:"customer,omitempty"`
	ItemPrices        []*CheckoutOneTimeForItemsItemPriceParams        `json:"item_prices,omitempty"`
	ItemTiers         []*CheckoutOneTimeForItemsItemTierParams         `json:"item_tiers,omitempty"`
	Charges           []*CheckoutOneTimeForItemsChargeParams           `json:"charges,omitempty"`
	InvoiceNote       string                                           `json:"invoice_note,omitempty"`
	Invoice           *CheckoutOneTimeForItemsInvoiceParams            `json:"invoice,omitempty"`
	Coupon            string                                           `json:"coupon,omitempty"`
	CouponIds         []string                                         `json:"coupon_ids,omitempty"`
	CurrencyCode      string                                           `json:"currency_code,omitempty"`
	Card              *CheckoutOneTimeForItemsCardParams               `json:"card,omitempty"`
	EntityIdentifiers []*CheckoutOneTimeForItemsEntityIdentifierParams `json:"entity_identifiers,omitempty"`
	RedirectUrl       string                                           `json:"redirect_url,omitempty"`
	CancelUrl         string                                           `json:"cancel_url,omitempty"`
	PassThruContent   string                                           `json:"pass_thru_content,omitempty"`
	BillingAddress    *CheckoutOneTimeForItemsBillingAddressParams     `json:"billing_address,omitempty"`
	ShippingAddress   *CheckoutOneTimeForItemsShippingAddressParams    `json:"shipping_address,omitempty"`
}
type CheckoutOneTimeForItemsCustomerParams struct {
	Id                       string          `json:"id,omitempty"`
	Email                    string          `json:"email,omitempty"`
	FirstName                string          `json:"first_name,omitempty"`
	LastName                 string          `json:"last_name,omitempty"`
	Company                  string          `json:"company,omitempty"`
	Taxability               enum.Taxability `json:"taxability,omitempty"`
	Locale                   string          `json:"locale,omitempty"`
	Phone                    string          `json:"phone,omitempty"`
	VatNumber                string          `json:"vat_number,omitempty"`
	VatNumberPrefix          string          `json:"vat_number_prefix,omitempty"`
	IsEinvoiceEnabled        *bool           `json:"is_einvoice_enabled,omitempty"`
	EntityIdentifierScheme   string          `json:"entity_identifier_scheme,omitempty"`
	EntityIdentifierStandard string          `json:"entity_identifier_standard,omitempty"`
	ConsolidatedInvoicing    *bool           `json:"consolidated_invoicing,omitempty"`
}
type CheckoutOneTimeForItemsItemPriceParams struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type CheckoutOneTimeForItemsItemTierParams struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type CheckoutOneTimeForItemsChargeParams struct {
	Amount                 *int32               `json:"amount,omitempty"`
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
type CheckoutOneTimeForItemsInvoiceParams struct {
	PoNumber string `json:"po_number,omitempty"`
}
type CheckoutOneTimeForItemsCardParams struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type CheckoutOneTimeForItemsEntityIdentifierParams struct {
	Id        string         `json:"id,omitempty"`
	Scheme    string         `json:"scheme,omitempty"`
	Value     string         `json:"value,omitempty"`
	Operation enum.Operation `json:"operation,omitempty"`
	Standard  string         `json:"standard,omitempty"`
}
type CheckoutOneTimeForItemsBillingAddressParams struct {
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
type CheckoutOneTimeForItemsShippingAddressParams struct {
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
type CheckoutNewForItemsRequestParams struct {
	Subscription               *CheckoutNewForItemsSubscriptionParams       `json:"subscription,omitempty"`
	Customer                   *CheckoutNewForItemsCustomerParams           `json:"customer,omitempty"`
	BillingCycles              *int32                                       `json:"billing_cycles,omitempty"`
	SubscriptionItems          []*CheckoutNewForItemsSubscriptionItemParams `json:"subscription_items,omitempty"`
	MandatoryItemsToRemove     []string                                     `json:"mandatory_items_to_remove,omitempty"`
	ItemTiers                  []*CheckoutNewForItemsItemTierParams         `json:"item_tiers,omitempty"`
	TermsToCharge              *int32                                       `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode       enum.BillingAlignmentMode                    `json:"billing_alignment_mode,omitempty"`
	CouponIds                  []string                                     `json:"coupon_ids,omitempty"`
	Card                       *CheckoutNewForItemsCardParams               `json:"card,omitempty"`
	EntityIdentifiers          []*CheckoutNewForItemsEntityIdentifierParams `json:"entity_identifiers,omitempty"`
	RedirectUrl                string                                       `json:"redirect_url,omitempty"`
	CancelUrl                  string                                       `json:"cancel_url,omitempty"`
	PassThruContent            string                                       `json:"pass_thru_content,omitempty"`
	AllowOfflinePaymentMethods *bool                                        `json:"allow_offline_payment_methods,omitempty"`
	BillingAddress             *CheckoutNewForItemsBillingAddressParams     `json:"billing_address,omitempty"`
	ShippingAddress            *CheckoutNewForItemsShippingAddressParams    `json:"shipping_address,omitempty"`
	ContractTerm               *CheckoutNewForItemsContractTermParams       `json:"contract_term,omitempty"`
}
type CheckoutNewForItemsSubscriptionParams struct {
	Id                                string                    `json:"id,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	SetupFee                          *int32                    `json:"setup_fee,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	Coupon                            string                    `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	InvoiceNotes                      string                    `json:"invoice_notes,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	AffiliateToken                    string                    `json:"affiliate_token,omitempty"`
}
type CheckoutNewForItemsCustomerParams struct {
	Id                       string          `json:"id,omitempty"`
	Email                    string          `json:"email,omitempty"`
	FirstName                string          `json:"first_name,omitempty"`
	LastName                 string          `json:"last_name,omitempty"`
	Company                  string          `json:"company,omitempty"`
	Taxability               enum.Taxability `json:"taxability,omitempty"`
	Locale                   string          `json:"locale,omitempty"`
	Phone                    string          `json:"phone,omitempty"`
	VatNumber                string          `json:"vat_number,omitempty"`
	VatNumberPrefix          string          `json:"vat_number_prefix,omitempty"`
	IsEinvoiceEnabled        *bool           `json:"is_einvoice_enabled,omitempty"`
	EntityIdentifierScheme   string          `json:"entity_identifier_scheme,omitempty"`
	EntityIdentifierStandard string          `json:"entity_identifier_standard,omitempty"`
}
type CheckoutNewForItemsSubscriptionItemParams struct {
	ItemPriceId        string              `json:"item_price_id"`
	Quantity           *int32              `json:"quantity,omitempty"`
	QuantityInDecimal  string              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32              `json:"billing_cycles,omitempty"`
	TrialEnd           *int64              `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32              `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool               `json:"charge_once,omitempty"`
	ItemType           enum.ItemType       `json:"item_type,omitempty"`
	ChargeOnOption     enum.ChargeOnOption `json:"charge_on_option,omitempty"`
}
type CheckoutNewForItemsItemTierParams struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type CheckoutNewForItemsCardParams struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type CheckoutNewForItemsEntityIdentifierParams struct {
	Id        string         `json:"id,omitempty"`
	Scheme    string         `json:"scheme,omitempty"`
	Value     string         `json:"value,omitempty"`
	Operation enum.Operation `json:"operation,omitempty"`
	Standard  string         `json:"standard,omitempty"`
}
type CheckoutNewForItemsBillingAddressParams struct {
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
type CheckoutNewForItemsShippingAddressParams struct {
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
type CheckoutNewForItemsContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type CheckoutExistingRequestParams struct {
	Subscription               *CheckoutExistingSubscriptionParams      `json:"subscription,omitempty"`
	Addons                     []*CheckoutExistingAddonParams           `json:"addons,omitempty"`
	EventBasedAddons           []*CheckoutExistingEventBasedAddonParams `json:"event_based_addons,omitempty"`
	ReplaceAddonList           *bool                                    `json:"replace_addon_list,omitempty"`
	MandatoryAddonsToRemove    []string                                 `json:"mandatory_addons_to_remove,omitempty"`
	InvoiceDate                *int64                                   `json:"invoice_date,omitempty"`
	BillingCycles              *int32                                   `json:"billing_cycles,omitempty"`
	TermsToCharge              *int32                                   `json:"terms_to_charge,omitempty"`
	ReactivateFrom             *int64                                   `json:"reactivate_from,omitempty"`
	BillingAlignmentMode       enum.BillingAlignmentMode                `json:"billing_alignment_mode,omitempty"`
	CouponIds                  []string                                 `json:"coupon_ids,omitempty"`
	Reactivate                 *bool                                    `json:"reactivate,omitempty"`
	ForceTermReset             *bool                                    `json:"force_term_reset,omitempty"`
	Customer                   *CheckoutExistingCustomerParams          `json:"customer,omitempty"`
	Card                       *CheckoutExistingCardParams              `json:"card,omitempty"`
	RedirectUrl                string                                   `json:"redirect_url,omitempty"`
	CancelUrl                  string                                   `json:"cancel_url,omitempty"`
	PassThruContent            string                                   `json:"pass_thru_content,omitempty"`
	Embed                      *bool                                    `json:"embed,omitempty"`
	IframeMessaging            *bool                                    `json:"iframe_messaging,omitempty"`
	AllowOfflinePaymentMethods *bool                                    `json:"allow_offline_payment_methods,omitempty"`
	ContractTerm               *CheckoutExistingContractTermParams      `json:"contract_term,omitempty"`
}
type CheckoutExistingSubscriptionParams struct {
	Id                                string                    `json:"id"`
	PlanId                            string                    `json:"plan_id,omitempty"`
	PlanQuantity                      *int32                    `json:"plan_quantity,omitempty"`
	PlanUnitPrice                     *int32                    `json:"plan_unit_price,omitempty"`
	SetupFee                          *int32                    `json:"setup_fee,omitempty"`
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
type CheckoutExistingAddonParams struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
}
type CheckoutExistingEventBasedAddonParams struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int32        `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
	QuantityInDecimal   string        `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string        `json:"unit_price_in_decimal,omitempty"`
}
type CheckoutExistingCustomerParams struct {
	VatNumber       string `json:"vat_number,omitempty"`
	VatNumberPrefix string `json:"vat_number_prefix,omitempty"`
}
type CheckoutExistingCardParams struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type CheckoutExistingContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type CheckoutExistingForItemsRequestParams struct {
	Subscription               *CheckoutExistingForItemsSubscriptionParams       `json:"subscription,omitempty"`
	SubscriptionItems          []*CheckoutExistingForItemsSubscriptionItemParams `json:"subscription_items,omitempty"`
	MandatoryItemsToRemove     []string                                          `json:"mandatory_items_to_remove,omitempty"`
	ReplaceItemsList           *bool                                             `json:"replace_items_list,omitempty"`
	ItemTiers                  []*CheckoutExistingForItemsItemTierParams         `json:"item_tiers,omitempty"`
	InvoiceDate                *int64                                            `json:"invoice_date,omitempty"`
	BillingCycles              *int32                                            `json:"billing_cycles,omitempty"`
	TermsToCharge              *int32                                            `json:"terms_to_charge,omitempty"`
	ReactivateFrom             *int64                                            `json:"reactivate_from,omitempty"`
	BillingAlignmentMode       enum.BillingAlignmentMode                         `json:"billing_alignment_mode,omitempty"`
	CouponIds                  []string                                          `json:"coupon_ids,omitempty"`
	Reactivate                 *bool                                             `json:"reactivate,omitempty"`
	ForceTermReset             *bool                                             `json:"force_term_reset,omitempty"`
	Customer                   *CheckoutExistingForItemsCustomerParams           `json:"customer,omitempty"`
	EntityIdentifiers          []*CheckoutExistingForItemsEntityIdentifierParams `json:"entity_identifiers,omitempty"`
	Card                       *CheckoutExistingForItemsCardParams               `json:"card,omitempty"`
	RedirectUrl                string                                            `json:"redirect_url,omitempty"`
	CancelUrl                  string                                            `json:"cancel_url,omitempty"`
	PassThruContent            string                                            `json:"pass_thru_content,omitempty"`
	AllowOfflinePaymentMethods *bool                                             `json:"allow_offline_payment_methods,omitempty"`
	ContractTerm               *CheckoutExistingForItemsContractTermParams       `json:"contract_term,omitempty"`
}
type CheckoutExistingForItemsSubscriptionParams struct {
	Id                                string                    `json:"id"`
	SetupFee                          *int32                    `json:"setup_fee,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	Coupon                            string                    `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	InvoiceNotes                      string                    `json:"invoice_notes,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type CheckoutExistingForItemsSubscriptionItemParams struct {
	ItemPriceId        string              `json:"item_price_id"`
	Quantity           *int32              `json:"quantity,omitempty"`
	QuantityInDecimal  string              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32              `json:"billing_cycles,omitempty"`
	TrialEnd           *int64              `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32              `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool               `json:"charge_once,omitempty"`
	ChargeOnOption     enum.ChargeOnOption `json:"charge_on_option,omitempty"`
	ItemType           enum.ItemType       `json:"item_type,omitempty"`
}
type CheckoutExistingForItemsItemTierParams struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type CheckoutExistingForItemsCustomerParams struct {
	VatNumber                string `json:"vat_number,omitempty"`
	VatNumberPrefix          string `json:"vat_number_prefix,omitempty"`
	IsEinvoiceEnabled        *bool  `json:"is_einvoice_enabled,omitempty"`
	EntityIdentifierScheme   string `json:"entity_identifier_scheme,omitempty"`
	EntityIdentifierStandard string `json:"entity_identifier_standard,omitempty"`
}
type CheckoutExistingForItemsEntityIdentifierParams struct {
	Id        string         `json:"id,omitempty"`
	Scheme    string         `json:"scheme,omitempty"`
	Value     string         `json:"value,omitempty"`
	Operation enum.Operation `json:"operation,omitempty"`
	Standard  string         `json:"standard,omitempty"`
}
type CheckoutExistingForItemsCardParams struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
}
type CheckoutExistingForItemsContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
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
	Id              string `json:"id"`
	VatNumber       string `json:"vat_number,omitempty"`
	VatNumberPrefix string `json:"vat_number_prefix,omitempty"`
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
	Id              string `json:"id"`
	VatNumber       string `json:"vat_number,omitempty"`
	VatNumberPrefix string `json:"vat_number_prefix,omitempty"`
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
	Quote       *AcceptQuoteQuoteParams `json:"quote,omitempty"`
	RedirectUrl string                  `json:"redirect_url,omitempty"`
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
	PlanId                string `json:"plan_id"`
	PlanQuantity          *int32 `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal string `json:"plan_quantity_in_decimal,omitempty"`
	Coupon                string `json:"coupon,omitempty"`
}
type CheckoutGiftAddonParams struct {
	Id                string `json:"id,omitempty"`
	Quantity          *int32 `json:"quantity,omitempty"`
	QuantityInDecimal string `json:"quantity_in_decimal,omitempty"`
}
type CheckoutGiftForItemsRequestParams struct {
	Gifter            *CheckoutGiftForItemsGifterParams             `json:"gifter,omitempty"`
	RedirectUrl       string                                        `json:"redirect_url,omitempty"`
	SubscriptionItems []*CheckoutGiftForItemsSubscriptionItemParams `json:"subscription_items,omitempty"`
	CouponIds         []string                                      `json:"coupon_ids,omitempty"`
}
type CheckoutGiftForItemsGifterParams struct {
	CustomerId string `json:"customer_id,omitempty"`
	Locale     string `json:"locale,omitempty"`
}
type CheckoutGiftForItemsSubscriptionItemParams struct {
	ItemPriceId       string `json:"item_price_id,omitempty"`
	Quantity          *int32 `json:"quantity,omitempty"`
	QuantityInDecimal string `json:"quantity_in_decimal,omitempty"`
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
