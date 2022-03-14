package subscription

import (
	"encoding/json"

	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	contractTermEnum "github.com/chargebee/chargebee-go/models/contractterm/enum"
	paymentIntentEnum "github.com/chargebee/chargebee-go/models/paymentintent/enum"
	subscriptionEnum "github.com/chargebee/chargebee-go/models/subscription/enum"
)

type Subscription struct {
	Id                                string                             `json:"id"`
	CurrencyCode                      string                             `json:"currency_code"`
	PlanId                            string                             `json:"plan_id"`
	PlanQuantity                      int32                              `json:"plan_quantity"`
	PlanUnitPrice                     int32                              `json:"plan_unit_price"`
	SetupFee                          int32                              `json:"setup_fee"`
	BillingPeriod                     int32                              `json:"billing_period"`
	BillingPeriodUnit                 subscriptionEnum.BillingPeriodUnit `json:"billing_period_unit"`
	StartDate                         int64                              `json:"start_date"`
	TrialEnd                          int64                              `json:"trial_end"`
	RemainingBillingCycles            int32                              `json:"remaining_billing_cycles"`
	PoNumber                          string                             `json:"po_number"`
	AutoCollection                    enum.AutoCollection                `json:"auto_collection"`
	PlanQuantityInDecimal             string                             `json:"plan_quantity_in_decimal"`
	PlanUnitPriceInDecimal            string                             `json:"plan_unit_price_in_decimal"`
	CustomerId                        string                             `json:"customer_id"`
	PlanAmount                        int32                              `json:"plan_amount"`
	PlanFreeQuantity                  int32                              `json:"plan_free_quantity"`
	Status                            subscriptionEnum.Status            `json:"status"`
	TrialStart                        int64                              `json:"trial_start"`
	TrialEndAction                    enum.TrialEndAction                `json:"trial_end_action"`
	CurrentTermStart                  int64                              `json:"current_term_start"`
	CurrentTermEnd                    int64                              `json:"current_term_end"`
	NextBillingAt                     int64                              `json:"next_billing_at"`
	CreatedAt                         int64                              `json:"created_at"`
	StartedAt                         int64                              `json:"started_at"`
	ActivatedAt                       int64                              `json:"activated_at"`
	GiftId                            string                             `json:"gift_id"`
	ContractTermBillingCycleOnRenewal int32                              `json:"contract_term_billing_cycle_on_renewal"`
	OverrideRelationship              bool                               `json:"override_relationship"`
	PauseDate                         int64                              `json:"pause_date"`
	ResumeDate                        int64                              `json:"resume_date"`
	CancelledAt                       int64                              `json:"cancelled_at"`
	CancelReason                      subscriptionEnum.CancelReason      `json:"cancel_reason"`
	AffiliateToken                    string                             `json:"affiliate_token"`
	CreatedFromIp                     string                             `json:"created_from_ip"`
	ResourceVersion                   int64                              `json:"resource_version"`
	UpdatedAt                         int64                              `json:"updated_at"`
	HasScheduledAdvanceInvoices       bool                               `json:"has_scheduled_advance_invoices"`
	HasScheduledChanges               bool                               `json:"has_scheduled_changes"`
	PaymentSourceId                   string                             `json:"payment_source_id"`
	PlanFreeQuantityInDecimal         string                             `json:"plan_free_quantity_in_decimal"`
	PlanAmountInDecimal               string                             `json:"plan_amount_in_decimal"`
	CancelScheduleCreatedAt           int64                              `json:"cancel_schedule_created_at"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod          `json:"offline_payment_method"`
	NetTermDays                       int32                              `json:"net_term_days"`
	SubscriptionItems                 []*SubscriptionItem                `json:"subscription_items"`
	ItemTiers                         []*ItemTier                        `json:"item_tiers"`
	ChargedItems                      []*ChargedItem                     `json:"charged_items"`
	DueInvoicesCount                  int32                              `json:"due_invoices_count"`
	DueSince                          int64                              `json:"due_since"`
	TotalDues                         int32                              `json:"total_dues"`
	Mrr                               int32                              `json:"mrr"`
	ExchangeRate                      float64                            `json:"exchange_rate"`
	BaseCurrencyCode                  string                             `json:"base_currency_code"`
	Addons                            []*Addon                           `json:"addons"`
	EventBasedAddons                  []*EventBasedAddon                 `json:"event_based_addons"`
	ChargedEventBasedAddons           []*ChargedEventBasedAddon          `json:"charged_event_based_addons"`
	Coupon                            string                             `json:"coupon"`
	Coupons                           []*Coupon                          `json:"coupons"`
	ShippingAddress                   *ShippingAddress                   `json:"shipping_address"`
	ReferralInfo                      *ReferralInfo                      `json:"referral_info"`
	InvoiceNotes                      string                             `json:"invoice_notes"`
	MetaData                          json.RawMessage                    `json:"meta_data"`
	Metadata                          json.RawMessage                    `json:"metadata"`
	Deleted                           bool                               `json:"deleted"`
	ChangesScheduledAt                int64                              `json:"changes_scheduled_at"`
	ContractTerm                      *ContractTerm                      `json:"contract_term"`
	CancelReasonCode                  string                             `json:"cancel_reason_code"`
	FreePeriod                        int32                              `json:"free_period"`
	FreePeriodUnit                    enum.FreePeriodUnit                `json:"free_period_unit"`
	CreatePendingInvoices             bool                               `json:"create_pending_invoices"`
	AutoCloseInvoices                 bool                               `json:"auto_close_invoices"`
	CustomField                       map[string]interface{}             `json:"custom_field"`
	Object                            string                             `json:"object"`
}
type SubscriptionItem struct {
	ItemPriceId           string              `json:"item_price_id"`
	ItemType              enum.ItemType       `json:"item_type"`
	Quantity              int32               `json:"quantity"`
	QuantityInDecimal     string              `json:"quantity_in_decimal"`
	MeteredQuantity       string              `json:"metered_quantity"`
	LastCalculatedAt      int64               `json:"last_calculated_at"`
	UnitPrice             int32               `json:"unit_price"`
	UnitPriceInDecimal    string              `json:"unit_price_in_decimal"`
	Amount                int32               `json:"amount"`
	AmountInDecimal       string              `json:"amount_in_decimal"`
	FreeQuantity          int32               `json:"free_quantity"`
	FreeQuantityInDecimal string              `json:"free_quantity_in_decimal"`
	TrialEnd              int64               `json:"trial_end"`
	BillingCycles         int32               `json:"billing_cycles"`
	ServicePeriodDays     int32               `json:"service_period_days"`
	ChargeOnEvent         enum.ChargeOnEvent  `json:"charge_on_event"`
	ChargeOnce            bool                `json:"charge_once"`
	ChargeOnOption        enum.ChargeOnOption `json:"charge_on_option"`
	Object                string              `json:"object"`
}
type ItemTier struct {
	ItemPriceId           string `json:"item_price_id"`
	StartingUnit          int32  `json:"starting_unit"`
	EndingUnit            int32  `json:"ending_unit"`
	Price                 int32  `json:"price"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal"`
	PriceInDecimal        string `json:"price_in_decimal"`
	Object                string `json:"object"`
}
type ChargedItem struct {
	ItemPriceId   string `json:"item_price_id"`
	LastChargedAt int64  `json:"last_charged_at"`
	Object        string `json:"object"`
}
type Addon struct {
	Id                     string `json:"id"`
	Quantity               int32  `json:"quantity"`
	UnitPrice              int32  `json:"unit_price"`
	Amount                 int32  `json:"amount"`
	TrialEnd               int64  `json:"trial_end"`
	RemainingBillingCycles int32  `json:"remaining_billing_cycles"`
	QuantityInDecimal      string `json:"quantity_in_decimal"`
	UnitPriceInDecimal     string `json:"unit_price_in_decimal"`
	AmountInDecimal        string `json:"amount_in_decimal"`
	Object                 string `json:"object"`
}
type EventBasedAddon struct {
	Id                  string       `json:"id"`
	Quantity            int32        `json:"quantity"`
	UnitPrice           int32        `json:"unit_price"`
	ServicePeriodInDays int32        `json:"service_period_in_days"`
	OnEvent             enum.OnEvent `json:"on_event"`
	ChargeOnce          bool         `json:"charge_once"`
	QuantityInDecimal   string       `json:"quantity_in_decimal"`
	UnitPriceInDecimal  string       `json:"unit_price_in_decimal"`
	Object              string       `json:"object"`
}
type ChargedEventBasedAddon struct {
	Id            string `json:"id"`
	LastChargedAt int64  `json:"last_charged_at"`
	Object        string `json:"object"`
}
type Coupon struct {
	CouponId     string `json:"coupon_id"`
	ApplyTill    int64  `json:"apply_till"`
	AppliedCount int32  `json:"applied_count"`
	CouponCode   string `json:"coupon_code"`
	Object       string `json:"object"`
}
type ShippingAddress struct {
	FirstName        string                `json:"first_name"`
	LastName         string                `json:"last_name"`
	Email            string                `json:"email"`
	Company          string                `json:"company"`
	Phone            string                `json:"phone"`
	Line1            string                `json:"line1"`
	Line2            string                `json:"line2"`
	Line3            string                `json:"line3"`
	City             string                `json:"city"`
	StateCode        string                `json:"state_code"`
	State            string                `json:"state"`
	Country          string                `json:"country"`
	Zip              string                `json:"zip"`
	ValidationStatus enum.ValidationStatus `json:"validation_status"`
	Object           string                `json:"object"`
}
type ReferralInfo struct {
	ReferralCode              string                                    `json:"referral_code"`
	CouponCode                string                                    `json:"coupon_code"`
	ReferrerId                string                                    `json:"referrer_id"`
	ExternalReferenceId       string                                    `json:"external_reference_id"`
	RewardStatus              subscriptionEnum.ReferralInfoRewardStatus `json:"reward_status"`
	ReferralSystem            enum.ReferralSystem                       `json:"referral_system"`
	AccountId                 string                                    `json:"account_id"`
	CampaignId                string                                    `json:"campaign_id"`
	ExternalCampaignId        string                                    `json:"external_campaign_id"`
	FriendOfferType           enum.FriendOfferType                      `json:"friend_offer_type"`
	ReferrerRewardType        enum.ReferrerRewardType                   `json:"referrer_reward_type"`
	NotifyReferralSystem      enum.NotifyReferralSystem                 `json:"notify_referral_system"`
	DestinationUrl            string                                    `json:"destination_url"`
	PostPurchaseWidgetEnabled bool                                      `json:"post_purchase_widget_enabled"`
	Object                    string                                    `json:"object"`
}
type ContractTerm struct {
	Id                       string                                       `json:"id"`
	Status                   subscriptionEnum.ContractTermStatus          `json:"status"`
	ContractStart            int64                                        `json:"contract_start"`
	ContractEnd              int64                                        `json:"contract_end"`
	BillingCycle             int32                                        `json:"billing_cycle"`
	ActionAtTermEnd          subscriptionEnum.ContractTermActionAtTermEnd `json:"action_at_term_end"`
	TotalContractValue       int64                                        `json:"total_contract_value"`
	CancellationCutoffPeriod int32                                        `json:"cancellation_cutoff_period"`
	CreatedAt                int64                                        `json:"created_at"`
	SubscriptionId           string                                       `json:"subscription_id"`
	RemainingBillingCycles   int32                                        `json:"remaining_billing_cycles"`
	Object                   string                                       `json:"object"`
}
type CreateRequestParams struct {
	Id                                string                          `json:"id,omitempty"`
	Customer                          *CreateCustomerParams           `json:"customer,omitempty"`
	EntityIdentifiers                 []*CreateEntityIdentifierParams `json:"entity_identifiers,omitempty"`
	PlanId                            string                          `json:"plan_id"`
	PlanQuantity                      *int32                          `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                          `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int32                          `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                          `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int32                          `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                          `json:"trial_end,omitempty"`
	BillingCycles                     *int32                          `json:"billing_cycles,omitempty"`
	Addons                            []*CreateAddonParams            `json:"addons,omitempty"`
	EventBasedAddons                  []*CreateEventBasedAddonParams  `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove           []string                        `json:"mandatory_addons_to_remove,omitempty"`
	StartDate                         *int64                          `json:"start_date,omitempty"`
	Coupon                            string                          `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection             `json:"auto_collection,omitempty"`
	TermsToCharge                     *int32                          `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode              enum.BillingAlignmentMode       `json:"billing_alignment_mode,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod       `json:"offline_payment_method,omitempty"`
	PoNumber                          string                          `json:"po_number,omitempty"`
	CouponIds                         []string                        `json:"coupon_ids,omitempty"`
	Card                              *CreateCardParams               `json:"card,omitempty"`
	BankAccount                       *CreateBankAccountParams        `json:"bank_account,omitempty"`
	TokenId                           string                          `json:"token_id,omitempty"`
	PaymentMethod                     *CreatePaymentMethodParams      `json:"payment_method,omitempty"`
	PaymentIntent                     *CreatePaymentIntentParams      `json:"payment_intent,omitempty"`
	BillingAddress                    *CreateBillingAddressParams     `json:"billing_address,omitempty"`
	ShippingAddress                   *CreateShippingAddressParams    `json:"shipping_address,omitempty"`
	AffiliateToken                    string                          `json:"affiliate_token,omitempty"`
	CreatedFromIp                     string                          `json:"created_from_ip,omitempty"`
	InvoiceNotes                      string                          `json:"invoice_notes,omitempty"`
	InvoiceDate                       *int64                          `json:"invoice_date,omitempty"`
	MetaData                          map[string]interface{}          `json:"meta_data,omitempty"`
	InvoiceImmediately                *bool                           `json:"invoice_immediately,omitempty"`
	FreePeriod                        *int32                          `json:"free_period,omitempty"`
	FreePeriodUnit                    enum.FreePeriodUnit             `json:"free_period_unit,omitempty"`
	ContractTerm                      *CreateContractTermParams       `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                          `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	TrialEndAction                    enum.TrialEndAction             `json:"trial_end_action,omitempty"`
	ClientProfileId                   string                          `json:"client_profile_id,omitempty"`
}
type CreateCustomerParams struct {
	Id                               string                       `json:"id,omitempty"`
	Email                            string                       `json:"email,omitempty"`
	FirstName                        string                       `json:"first_name,omitempty"`
	LastName                         string                       `json:"last_name,omitempty"`
	Company                          string                       `json:"company,omitempty"`
	Taxability                       enum.Taxability              `json:"taxability,omitempty"`
	Locale                           string                       `json:"locale,omitempty"`
	EntityCode                       enum.EntityCode              `json:"entity_code,omitempty"`
	ExemptNumber                     string                       `json:"exempt_number,omitempty"`
	NetTermDays                      *int32                       `json:"net_term_days,omitempty"`
	TaxjarExemptionCategory          enum.TaxjarExemptionCategory `json:"taxjar_exemption_category,omitempty"`
	Phone                            string                       `json:"phone,omitempty"`
	AutoCollection                   enum.AutoCollection          `json:"auto_collection,omitempty"`
	OfflinePaymentMethod             enum.OfflinePaymentMethod    `json:"offline_payment_method,omitempty"`
	AllowDirectDebit                 *bool                        `json:"allow_direct_debit,omitempty"`
	ConsolidatedInvoicing            *bool                        `json:"consolidated_invoicing,omitempty"`
	VatNumber                        string                       `json:"vat_number,omitempty"`
	VatNumberPrefix                  string                       `json:"vat_number_prefix,omitempty"`
	EntityIdentifierScheme           string                       `json:"entity_identifier_scheme,omitempty"`
	EntityIdentifierStandard         string                       `json:"entity_identifier_standard,omitempty"`
	IsEinvoiceEnabled                *bool                        `json:"is_einvoice_enabled,omitempty"`
	RegisteredForGst                 *bool                        `json:"registered_for_gst,omitempty"`
	BusinessCustomerWithoutVatNumber *bool                        `json:"business_customer_without_vat_number,omitempty"`
	ExemptionDetails                 []map[string]interface{}     `json:"exemption_details,omitempty"`
	CustomerType                     enum.CustomerType            `json:"customer_type,omitempty"`
}
type CreateEntityIdentifierParams struct {
	Id       string `json:"id,omitempty"`
	Scheme   string `json:"scheme,omitempty"`
	Value    string `json:"value,omitempty"`
	Standard string `json:"standard,omitempty"`
}
type CreateAddonParams struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}
type CreateEventBasedAddonParams struct {
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
type CreateCardParams struct {
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	Number                string                 `json:"number,omitempty"`
	ExpiryMonth           *int32                 `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                 `json:"expiry_year,omitempty"`
	Cvv                   string                 `json:"cvv,omitempty"`
	BillingAddr1          string                 `json:"billing_addr1,omitempty"`
	BillingAddr2          string                 `json:"billing_addr2,omitempty"`
	BillingCity           string                 `json:"billing_city,omitempty"`
	BillingStateCode      string                 `json:"billing_state_code,omitempty"`
	BillingState          string                 `json:"billing_state,omitempty"`
	BillingZip            string                 `json:"billing_zip,omitempty"`
	BillingCountry        string                 `json:"billing_country,omitempty"`
	IpAddress             string                 `json:"ip_address,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type CreateBankAccountParams struct {
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	Iban                  string                 `json:"iban,omitempty"`
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	Company               string                 `json:"company,omitempty"`
	Email                 string                 `json:"email,omitempty"`
	Phone                 string                 `json:"phone,omitempty"`
	BankName              string                 `json:"bank_name,omitempty"`
	AccountNumber         string                 `json:"account_number,omitempty"`
	RoutingNumber         string                 `json:"routing_number,omitempty"`
	BankCode              string                 `json:"bank_code,omitempty"`
	AccountType           enum.AccountType       `json:"account_type,omitempty"`
	AccountHolderType     enum.AccountHolderType `json:"account_holder_type,omitempty"`
	EcheckType            enum.EcheckType        `json:"echeck_type,omitempty"`
	IssuingCountry        string                 `json:"issuing_country,omitempty"`
	SwedishIdentityNumber string                 `json:"swedish_identity_number,omitempty"`
	BillingAddress        map[string]interface{} `json:"billing_address,omitempty"`
}
type CreatePaymentMethodParams struct {
	Type                  enum.Type              `json:"type,omitempty"`
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	ReferenceId           string                 `json:"reference_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	IssuingCountry        string                 `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type CreatePaymentIntentParams struct {
	Id                    string                              `json:"id,omitempty"`
	GatewayAccountId      string                              `json:"gateway_account_id,omitempty"`
	GwToken               string                              `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntentEnum.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                              `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                              `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}              `json:"additional_information,omitempty"`
}
type CreateBillingAddressParams struct {
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
type CreateShippingAddressParams struct {
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
type CreateContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type CreateForCustomerRequestParams struct {
	Id                                string                                    `json:"id,omitempty"`
	PlanId                            string                                    `json:"plan_id"`
	PlanQuantity                      *int32                                    `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                                    `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int32                                    `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                                    `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int32                                    `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                                    `json:"trial_end,omitempty"`
	BillingCycles                     *int32                                    `json:"billing_cycles,omitempty"`
	Addons                            []*CreateForCustomerAddonParams           `json:"addons,omitempty"`
	EventBasedAddons                  []*CreateForCustomerEventBasedAddonParams `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove           []string                                  `json:"mandatory_addons_to_remove,omitempty"`
	StartDate                         *int64                                    `json:"start_date,omitempty"`
	Coupon                            string                                    `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection                       `json:"auto_collection,omitempty"`
	TermsToCharge                     *int32                                    `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode              enum.BillingAlignmentMode                 `json:"billing_alignment_mode,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod                 `json:"offline_payment_method,omitempty"`
	PoNumber                          string                                    `json:"po_number,omitempty"`
	CouponIds                         []string                                  `json:"coupon_ids,omitempty"`
	PaymentSourceId                   string                                    `json:"payment_source_id,omitempty"`
	OverrideRelationship              *bool                                     `json:"override_relationship,omitempty"`
	ShippingAddress                   *CreateForCustomerShippingAddressParams   `json:"shipping_address,omitempty"`
	InvoiceNotes                      string                                    `json:"invoice_notes,omitempty"`
	InvoiceDate                       *int64                                    `json:"invoice_date,omitempty"`
	MetaData                          map[string]interface{}                    `json:"meta_data,omitempty"`
	InvoiceImmediately                *bool                                     `json:"invoice_immediately,omitempty"`
	ReplacePrimaryPaymentSource       *bool                                     `json:"replace_primary_payment_source,omitempty"`
	PaymentIntent                     *CreateForCustomerPaymentIntentParams     `json:"payment_intent,omitempty"`
	FreePeriod                        *int32                                    `json:"free_period,omitempty"`
	FreePeriodUnit                    enum.FreePeriodUnit                       `json:"free_period_unit,omitempty"`
	ContractTerm                      *CreateForCustomerContractTermParams      `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	TrialEndAction                    enum.TrialEndAction                       `json:"trial_end_action,omitempty"`
}
type CreateForCustomerAddonParams struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}
type CreateForCustomerEventBasedAddonParams struct {
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
type CreateForCustomerShippingAddressParams struct {
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
type CreateForCustomerPaymentIntentParams struct {
	Id                    string                              `json:"id,omitempty"`
	GatewayAccountId      string                              `json:"gateway_account_id,omitempty"`
	GwToken               string                              `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntentEnum.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                              `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                              `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}              `json:"additional_information,omitempty"`
}
type CreateForCustomerContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type CreateWithItemsRequestParams struct {
	Id                                string                                   `json:"id,omitempty"`
	TrialEnd                          *int64                                   `json:"trial_end,omitempty"`
	BillingCycles                     *int32                                   `json:"billing_cycles,omitempty"`
	SubscriptionItems                 []*CreateWithItemsSubscriptionItemParams `json:"subscription_items,omitempty"`
	SetupFee                          *int32                                   `json:"setup_fee,omitempty"`
	MandatoryItemsToRemove            []string                                 `json:"mandatory_items_to_remove,omitempty"`
	ItemTiers                         []*CreateWithItemsItemTierParams         `json:"item_tiers,omitempty"`
	NetTermDays                       *int32                                   `json:"net_term_days,omitempty"`
	StartDate                         *int64                                   `json:"start_date,omitempty"`
	Coupon                            string                                   `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection                      `json:"auto_collection,omitempty"`
	TermsToCharge                     *int32                                   `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode              enum.BillingAlignmentMode                `json:"billing_alignment_mode,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod                `json:"offline_payment_method,omitempty"`
	PoNumber                          string                                   `json:"po_number,omitempty"`
	CouponIds                         []string                                 `json:"coupon_ids,omitempty"`
	PaymentSourceId                   string                                   `json:"payment_source_id,omitempty"`
	OverrideRelationship              *bool                                    `json:"override_relationship,omitempty"`
	ShippingAddress                   *CreateWithItemsShippingAddressParams    `json:"shipping_address,omitempty"`
	InvoiceNotes                      string                                   `json:"invoice_notes,omitempty"`
	InvoiceDate                       *int64                                   `json:"invoice_date,omitempty"`
	MetaData                          map[string]interface{}                   `json:"meta_data,omitempty"`
	InvoiceImmediately                *bool                                    `json:"invoice_immediately,omitempty"`
	ReplacePrimaryPaymentSource       *bool                                    `json:"replace_primary_payment_source,omitempty"`
	PaymentIntent                     *CreateWithItemsPaymentIntentParams      `json:"payment_intent,omitempty"`
	FreePeriod                        *int32                                   `json:"free_period,omitempty"`
	FreePeriodUnit                    enum.FreePeriodUnit                      `json:"free_period_unit,omitempty"`
	ContractTerm                      *CreateWithItemsContractTermParams       `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                   `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	CreatePendingInvoices             *bool                                    `json:"create_pending_invoices,omitempty"`
	AutoCloseInvoices                 *bool                                    `json:"auto_close_invoices,omitempty"`
	FirstInvoicePending               *bool                                    `json:"first_invoice_pending,omitempty"`
	TrialEndAction                    enum.TrialEndAction                      `json:"trial_end_action,omitempty"`
}
type CreateWithItemsSubscriptionItemParams struct {
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
type CreateWithItemsItemTierParams struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type CreateWithItemsShippingAddressParams struct {
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
type CreateWithItemsPaymentIntentParams struct {
	Id                    string                              `json:"id,omitempty"`
	GatewayAccountId      string                              `json:"gateway_account_id,omitempty"`
	GwToken               string                              `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntentEnum.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                              `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                              `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}              `json:"additional_information,omitempty"`
}
type CreateWithItemsContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type ListRequestParams struct {
	Limit                  *int32                  `json:"limit,omitempty"`
	Offset                 string                  `json:"offset,omitempty"`
	IncludeDeleted         *bool                   `json:"include_deleted,omitempty"`
	Id                     *filter.StringFilter    `json:"id,omitempty"`
	CustomerId             *filter.StringFilter    `json:"customer_id,omitempty"`
	PlanId                 *filter.StringFilter    `json:"plan_id,omitempty"`
	ItemId                 *filter.StringFilter    `json:"item_id,omitempty"`
	ItemPriceId            *filter.StringFilter    `json:"item_price_id,omitempty"`
	Status                 *filter.EnumFilter      `json:"status,omitempty"`
	CancelReason           *filter.EnumFilter      `json:"cancel_reason,omitempty"`
	CancelReasonCode       *filter.StringFilter    `json:"cancel_reason_code,omitempty"`
	RemainingBillingCycles *filter.NumberFilter    `json:"remaining_billing_cycles,omitempty"`
	CreatedAt              *filter.TimestampFilter `json:"created_at,omitempty"`
	ActivatedAt            *filter.TimestampFilter `json:"activated_at,omitempty"`
	NextBillingAt          *filter.TimestampFilter `json:"next_billing_at,omitempty"`
	CancelledAt            *filter.TimestampFilter `json:"cancelled_at,omitempty"`
	HasScheduledChanges    *filter.BooleanFilter   `json:"has_scheduled_changes,omitempty"`
	UpdatedAt              *filter.TimestampFilter `json:"updated_at,omitempty"`
	OfflinePaymentMethod   *filter.EnumFilter      `json:"offline_payment_method,omitempty"`
	AutoCloseInvoices      *filter.BooleanFilter   `json:"auto_close_invoices,omitempty"`
	OverrideRelationship   *filter.BooleanFilter   `json:"override_relationship,omitempty"`
	SortBy                 *filter.SortFilter      `json:"sort_by,omitempty"`
}
type SubscriptionsForCustomerRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
type ContractTermsForSubscriptionRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
type RetrieveRequestParams struct {
}
type RetrieveWithScheduledChangesRequestParams struct {
}
type RemoveScheduledCancellationRequestParams struct {
	BillingCycles *int32 `json:"billing_cycles,omitempty"`
}
type RemoveCouponsRequestParams struct {
	CouponIds []string `json:"coupon_ids,omitempty"`
}
type UpdateRequestParams struct {
	PlanId                            string                         `json:"plan_id,omitempty"`
	PlanQuantity                      *int32                         `json:"plan_quantity,omitempty"`
	PlanUnitPrice                     *int32                         `json:"plan_unit_price,omitempty"`
	SetupFee                          *int32                         `json:"setup_fee,omitempty"`
	Addons                            []*UpdateAddonParams           `json:"addons,omitempty"`
	EventBasedAddons                  []*UpdateEventBasedAddonParams `json:"event_based_addons,omitempty"`
	ReplaceAddonList                  *bool                          `json:"replace_addon_list,omitempty"`
	MandatoryAddonsToRemove           []string                       `json:"mandatory_addons_to_remove,omitempty"`
	PlanQuantityInDecimal             string                         `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPriceInDecimal            string                         `json:"plan_unit_price_in_decimal,omitempty"`
	InvoiceDate                       *int64                         `json:"invoice_date,omitempty"`
	StartDate                         *int64                         `json:"start_date,omitempty"`
	TrialEnd                          *int64                         `json:"trial_end,omitempty"`
	BillingCycles                     *int32                         `json:"billing_cycles,omitempty"`
	Coupon                            string                         `json:"coupon,omitempty"`
	TermsToCharge                     *int32                         `json:"terms_to_charge,omitempty"`
	ReactivateFrom                    *int64                         `json:"reactivate_from,omitempty"`
	BillingAlignmentMode              enum.BillingAlignmentMode      `json:"billing_alignment_mode,omitempty"`
	AutoCollection                    enum.AutoCollection            `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod      `json:"offline_payment_method,omitempty"`
	PoNumber                          string                         `json:"po_number,omitempty"`
	CouponIds                         []string                       `json:"coupon_ids,omitempty"`
	ReplaceCouponList                 *bool                          `json:"replace_coupon_list,omitempty"`
	Prorate                           *bool                          `json:"prorate,omitempty"`
	EndOfTerm                         *bool                          `json:"end_of_term,omitempty"`
	ForceTermReset                    *bool                          `json:"force_term_reset,omitempty"`
	Reactivate                        *bool                          `json:"reactivate,omitempty"`
	Card                              *UpdateCardParams              `json:"card,omitempty"`
	TokenId                           string                         `json:"token_id,omitempty"`
	PaymentMethod                     *UpdatePaymentMethodParams     `json:"payment_method,omitempty"`
	PaymentIntent                     *UpdatePaymentIntentParams     `json:"payment_intent,omitempty"`
	BillingAddress                    *UpdateBillingAddressParams    `json:"billing_address,omitempty"`
	ShippingAddress                   *UpdateShippingAddressParams   `json:"shipping_address,omitempty"`
	Customer                          *UpdateCustomerParams          `json:"customer,omitempty"`
	InvoiceNotes                      string                         `json:"invoice_notes,omitempty"`
	MetaData                          map[string]interface{}         `json:"meta_data,omitempty"`
	InvoiceImmediately                *bool                          `json:"invoice_immediately,omitempty"`
	OverrideRelationship              *bool                          `json:"override_relationship,omitempty"`
	ChangesScheduledAt                *int64                         `json:"changes_scheduled_at,omitempty"`
	ChangeOption                      enum.ChangeOption              `json:"change_option,omitempty"`
	ContractTerm                      *UpdateContractTermParams      `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                         `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	FreePeriod                        *int32                         `json:"free_period,omitempty"`
	FreePeriodUnit                    enum.FreePeriodUnit            `json:"free_period_unit,omitempty"`
	TrialEndAction                    enum.TrialEndAction            `json:"trial_end_action,omitempty"`
}
type UpdateAddonParams struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}
type UpdateEventBasedAddonParams struct {
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
type UpdateCardParams struct {
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	Number                string                 `json:"number,omitempty"`
	ExpiryMonth           *int32                 `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                 `json:"expiry_year,omitempty"`
	Cvv                   string                 `json:"cvv,omitempty"`
	BillingAddr1          string                 `json:"billing_addr1,omitempty"`
	BillingAddr2          string                 `json:"billing_addr2,omitempty"`
	BillingCity           string                 `json:"billing_city,omitempty"`
	BillingStateCode      string                 `json:"billing_state_code,omitempty"`
	BillingState          string                 `json:"billing_state,omitempty"`
	BillingZip            string                 `json:"billing_zip,omitempty"`
	BillingCountry        string                 `json:"billing_country,omitempty"`
	IpAddress             string                 `json:"ip_address,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type UpdatePaymentMethodParams struct {
	Type                  enum.Type              `json:"type,omitempty"`
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	ReferenceId           string                 `json:"reference_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	IssuingCountry        string                 `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type UpdatePaymentIntentParams struct {
	Id                    string                              `json:"id,omitempty"`
	GatewayAccountId      string                              `json:"gateway_account_id,omitempty"`
	GwToken               string                              `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntentEnum.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                              `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                              `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}              `json:"additional_information,omitempty"`
}
type UpdateBillingAddressParams struct {
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
type UpdateShippingAddressParams struct {
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
type UpdateCustomerParams struct {
	VatNumber                        string `json:"vat_number,omitempty"`
	VatNumberPrefix                  string `json:"vat_number_prefix,omitempty"`
	EntityIdentifierScheme           string `json:"entity_identifier_scheme,omitempty"`
	IsEinvoiceEnabled                *bool  `json:"is_einvoice_enabled,omitempty"`
	EntityIdentifierStandard         string `json:"entity_identifier_standard,omitempty"`
	BusinessCustomerWithoutVatNumber *bool  `json:"business_customer_without_vat_number,omitempty"`
	RegisteredForGst                 *bool  `json:"registered_for_gst,omitempty"`
}
type UpdateContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type UpdateForItemsRequestParams struct {
	SubscriptionItems                 []*UpdateForItemsSubscriptionItemParams `json:"subscription_items,omitempty"`
	MandatoryItemsToRemove            []string                                `json:"mandatory_items_to_remove,omitempty"`
	ReplaceItemsList                  *bool                                   `json:"replace_items_list,omitempty"`
	SetupFee                          *int32                                  `json:"setup_fee,omitempty"`
	ItemTiers                         []*UpdateForItemsItemTierParams         `json:"item_tiers,omitempty"`
	NetTermDays                       *int32                                  `json:"net_term_days,omitempty"`
	InvoiceDate                       *int64                                  `json:"invoice_date,omitempty"`
	StartDate                         *int64                                  `json:"start_date,omitempty"`
	TrialEnd                          *int64                                  `json:"trial_end,omitempty"`
	BillingCycles                     *int32                                  `json:"billing_cycles,omitempty"`
	Coupon                            string                                  `json:"coupon,omitempty"`
	TermsToCharge                     *int32                                  `json:"terms_to_charge,omitempty"`
	ReactivateFrom                    *int64                                  `json:"reactivate_from,omitempty"`
	BillingAlignmentMode              enum.BillingAlignmentMode               `json:"billing_alignment_mode,omitempty"`
	AutoCollection                    enum.AutoCollection                     `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod               `json:"offline_payment_method,omitempty"`
	PoNumber                          string                                  `json:"po_number,omitempty"`
	CouponIds                         []string                                `json:"coupon_ids,omitempty"`
	ReplaceCouponList                 *bool                                   `json:"replace_coupon_list,omitempty"`
	Prorate                           *bool                                   `json:"prorate,omitempty"`
	EndOfTerm                         *bool                                   `json:"end_of_term,omitempty"`
	ForceTermReset                    *bool                                   `json:"force_term_reset,omitempty"`
	Reactivate                        *bool                                   `json:"reactivate,omitempty"`
	Card                              *UpdateForItemsCardParams               `json:"card,omitempty"`
	TokenId                           string                                  `json:"token_id,omitempty"`
	PaymentMethod                     *UpdateForItemsPaymentMethodParams      `json:"payment_method,omitempty"`
	PaymentIntent                     *UpdateForItemsPaymentIntentParams      `json:"payment_intent,omitempty"`
	BillingAddress                    *UpdateForItemsBillingAddressParams     `json:"billing_address,omitempty"`
	ShippingAddress                   *UpdateForItemsShippingAddressParams    `json:"shipping_address,omitempty"`
	Customer                          *UpdateForItemsCustomerParams           `json:"customer,omitempty"`
	InvoiceNotes                      string                                  `json:"invoice_notes,omitempty"`
	MetaData                          map[string]interface{}                  `json:"meta_data,omitempty"`
	InvoiceImmediately                *bool                                   `json:"invoice_immediately,omitempty"`
	OverrideRelationship              *bool                                   `json:"override_relationship,omitempty"`
	ChangesScheduledAt                *int64                                  `json:"changes_scheduled_at,omitempty"`
	ChangeOption                      enum.ChangeOption                       `json:"change_option,omitempty"`
	ContractTerm                      *UpdateForItemsContractTermParams       `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                  `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	FreePeriod                        *int32                                  `json:"free_period,omitempty"`
	FreePeriodUnit                    enum.FreePeriodUnit                     `json:"free_period_unit,omitempty"`
	CreatePendingInvoices             *bool                                   `json:"create_pending_invoices,omitempty"`
	AutoCloseInvoices                 *bool                                   `json:"auto_close_invoices,omitempty"`
	TrialEndAction                    enum.TrialEndAction                     `json:"trial_end_action,omitempty"`
}
type UpdateForItemsSubscriptionItemParams struct {
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
type UpdateForItemsItemTierParams struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type UpdateForItemsCardParams struct {
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	Number                string                 `json:"number,omitempty"`
	ExpiryMonth           *int32                 `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                 `json:"expiry_year,omitempty"`
	Cvv                   string                 `json:"cvv,omitempty"`
	BillingAddr1          string                 `json:"billing_addr1,omitempty"`
	BillingAddr2          string                 `json:"billing_addr2,omitempty"`
	BillingCity           string                 `json:"billing_city,omitempty"`
	BillingStateCode      string                 `json:"billing_state_code,omitempty"`
	BillingState          string                 `json:"billing_state,omitempty"`
	BillingZip            string                 `json:"billing_zip,omitempty"`
	BillingCountry        string                 `json:"billing_country,omitempty"`
	IpAddress             string                 `json:"ip_address,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type UpdateForItemsPaymentMethodParams struct {
	Type                  enum.Type              `json:"type,omitempty"`
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	ReferenceId           string                 `json:"reference_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	IssuingCountry        string                 `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type UpdateForItemsPaymentIntentParams struct {
	Id                    string                              `json:"id,omitempty"`
	GatewayAccountId      string                              `json:"gateway_account_id,omitempty"`
	GwToken               string                              `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntentEnum.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                              `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                              `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}              `json:"additional_information,omitempty"`
}
type UpdateForItemsBillingAddressParams struct {
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
type UpdateForItemsShippingAddressParams struct {
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
type UpdateForItemsCustomerParams struct {
	VatNumber                        string `json:"vat_number,omitempty"`
	VatNumberPrefix                  string `json:"vat_number_prefix,omitempty"`
	EntityIdentifierScheme           string `json:"entity_identifier_scheme,omitempty"`
	IsEinvoiceEnabled                *bool  `json:"is_einvoice_enabled,omitempty"`
	EntityIdentifierStandard         string `json:"entity_identifier_standard,omitempty"`
	BusinessCustomerWithoutVatNumber *bool  `json:"business_customer_without_vat_number,omitempty"`
	RegisteredForGst                 *bool  `json:"registered_for_gst,omitempty"`
}
type UpdateForItemsContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type ChangeTermEndRequestParams struct {
	TermEndsAt         *int64 `json:"term_ends_at"`
	Prorate            *bool  `json:"prorate,omitempty"`
	InvoiceImmediately *bool  `json:"invoice_immediately,omitempty"`
}
type ReactivateRequestParams struct {
	TrialEnd                          *int64                         `json:"trial_end,omitempty"`
	BillingCycles                     *int32                         `json:"billing_cycles,omitempty"`
	TrialPeriodDays                   *int32                         `json:"trial_period_days,omitempty"`
	ReactivateFrom                    *int64                         `json:"reactivate_from,omitempty"`
	InvoiceImmediately                *bool                          `json:"invoice_immediately,omitempty"`
	BillingAlignmentMode              enum.BillingAlignmentMode      `json:"billing_alignment_mode,omitempty"`
	TermsToCharge                     *int32                         `json:"terms_to_charge,omitempty"`
	InvoiceDate                       *int64                         `json:"invoice_date,omitempty"`
	ContractTerm                      *ReactivateContractTermParams  `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                         `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	PaymentIntent                     *ReactivatePaymentIntentParams `json:"payment_intent,omitempty"`
}
type ReactivateContractTermParams struct {
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type ReactivatePaymentIntentParams struct {
	Id                    string                              `json:"id,omitempty"`
	GatewayAccountId      string                              `json:"gateway_account_id,omitempty"`
	GwToken               string                              `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntentEnum.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                              `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                              `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}              `json:"additional_information,omitempty"`
}
type AddChargeAtTermEndRequestParams struct {
	Amount                 *int32               `json:"amount,omitempty"`
	Description            string               `json:"description"`
	AmountInDecimal        string               `json:"amount_in_decimal,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	DateFrom               *int64               `json:"date_from,omitempty"`
	DateTo                 *int64               `json:"date_to,omitempty"`
}
type ChargeAddonAtTermEndRequestParams struct {
	AddonId                 string `json:"addon_id"`
	AddonQuantity           *int32 `json:"addon_quantity,omitempty"`
	AddonUnitPrice          *int32 `json:"addon_unit_price,omitempty"`
	AddonQuantityInDecimal  string `json:"addon_quantity_in_decimal,omitempty"`
	AddonUnitPriceInDecimal string `json:"addon_unit_price_in_decimal,omitempty"`
	DateFrom                *int64 `json:"date_from,omitempty"`
	DateTo                  *int64 `json:"date_to,omitempty"`
}
type ChargeFutureRenewalsRequestParams struct {
	TermsToCharge         *int32                                             `json:"terms_to_charge,omitempty"`
	SpecificDatesSchedule []*ChargeFutureRenewalsSpecificDatesScheduleParams `json:"specific_dates_schedule,omitempty"`
	FixedIntervalSchedule *ChargeFutureRenewalsFixedIntervalScheduleParams   `json:"fixed_interval_schedule,omitempty"`
	InvoiceImmediately    *bool                                              `json:"invoice_immediately,omitempty"`
	ScheduleType          enum.ScheduleType                                  `json:"schedule_type,omitempty"`
}
type ChargeFutureRenewalsSpecificDatesScheduleParams struct {
	TermsToCharge *int32 `json:"terms_to_charge,omitempty"`
	Date          *int64 `json:"date,omitempty"`
}
type ChargeFutureRenewalsFixedIntervalScheduleParams struct {
	NumberOfOccurrences *int32             `json:"number_of_occurrences,omitempty"`
	DaysBeforeRenewal   *int32             `json:"days_before_renewal,omitempty"`
	EndScheduleOn       enum.EndScheduleOn `json:"end_schedule_on,omitempty"`
	EndDate             *int64             `json:"end_date,omitempty"`
}
type EditAdvanceInvoiceScheduleRequestParams struct {
	TermsToCharge         *int32                                                   `json:"terms_to_charge,omitempty"`
	ScheduleType          enum.ScheduleType                                        `json:"schedule_type,omitempty"`
	SpecificDatesSchedule []*EditAdvanceInvoiceScheduleSpecificDatesScheduleParams `json:"specific_dates_schedule,omitempty"`
	FixedIntervalSchedule *EditAdvanceInvoiceScheduleFixedIntervalScheduleParams   `json:"fixed_interval_schedule,omitempty"`
}
type EditAdvanceInvoiceScheduleSpecificDatesScheduleParams struct {
	Id            string `json:"id,omitempty"`
	TermsToCharge *int32 `json:"terms_to_charge,omitempty"`
	Date          *int64 `json:"date,omitempty"`
}
type EditAdvanceInvoiceScheduleFixedIntervalScheduleParams struct {
	NumberOfOccurrences *int32             `json:"number_of_occurrences,omitempty"`
	DaysBeforeRenewal   *int32             `json:"days_before_renewal,omitempty"`
	EndScheduleOn       enum.EndScheduleOn `json:"end_schedule_on,omitempty"`
	EndDate             *int64             `json:"end_date,omitempty"`
}
type RemoveAdvanceInvoiceScheduleRequestParams struct {
	SpecificDatesSchedule []*RemoveAdvanceInvoiceScheduleSpecificDatesScheduleParams `json:"specific_dates_schedule,omitempty"`
}
type RemoveAdvanceInvoiceScheduleSpecificDatesScheduleParams struct {
	Id string `json:"id,omitempty"`
}
type RegenerateInvoiceRequestParams struct {
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
	Prorate            *bool  `json:"prorate,omitempty"`
	InvoiceImmediately *bool  `json:"invoice_immediately,omitempty"`
}
type ImportSubscriptionRequestParams struct {
	Id                                string                                            `json:"id,omitempty"`
	Customer                          *ImportSubscriptionCustomerParams                 `json:"customer,omitempty"`
	ClientProfileId                   string                                            `json:"client_profile_id,omitempty"`
	PlanId                            string                                            `json:"plan_id"`
	PlanQuantity                      *int32                                            `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                                            `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int32                                            `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                                            `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int32                                            `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                                            `json:"trial_end,omitempty"`
	BillingCycles                     *int32                                            `json:"billing_cycles,omitempty"`
	Addons                            []*ImportSubscriptionAddonParams                  `json:"addons,omitempty"`
	EventBasedAddons                  []*ImportSubscriptionEventBasedAddonParams        `json:"event_based_addons,omitempty"`
	ChargedEventBasedAddons           []*ImportSubscriptionChargedEventBasedAddonParams `json:"charged_event_based_addons,omitempty"`
	StartDate                         *int64                                            `json:"start_date,omitempty"`
	AutoCollection                    enum.AutoCollection                               `json:"auto_collection,omitempty"`
	PoNumber                          string                                            `json:"po_number,omitempty"`
	CouponIds                         []string                                          `json:"coupon_ids,omitempty"`
	ContractTerm                      *ImportSubscriptionContractTermParams             `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                            `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	Status                            subscriptionEnum.Status                           `json:"status"`
	CurrentTermEnd                    *int64                                            `json:"current_term_end,omitempty"`
	CurrentTermStart                  *int64                                            `json:"current_term_start,omitempty"`
	TrialStart                        *int64                                            `json:"trial_start,omitempty"`
	CancelledAt                       *int64                                            `json:"cancelled_at,omitempty"`
	StartedAt                         *int64                                            `json:"started_at,omitempty"`
	ActivatedAt                       *int64                                            `json:"activated_at,omitempty"`
	PauseDate                         *int64                                            `json:"pause_date,omitempty"`
	ResumeDate                        *int64                                            `json:"resume_date,omitempty"`
	CreateCurrentTermInvoice          *bool                                             `json:"create_current_term_invoice,omitempty"`
	Card                              *ImportSubscriptionCardParams                     `json:"card,omitempty"`
	PaymentMethod                     *ImportSubscriptionPaymentMethodParams            `json:"payment_method,omitempty"`
	BillingAddress                    *ImportSubscriptionBillingAddressParams           `json:"billing_address,omitempty"`
	ShippingAddress                   *ImportSubscriptionShippingAddressParams          `json:"shipping_address,omitempty"`
	AffiliateToken                    string                                            `json:"affiliate_token,omitempty"`
	InvoiceNotes                      string                                            `json:"invoice_notes,omitempty"`
	MetaData                          map[string]interface{}                            `json:"meta_data,omitempty"`
	Transaction                       *ImportSubscriptionTransactionParams              `json:"transaction,omitempty"`
}
type ImportSubscriptionCustomerParams struct {
	Id                      string                       `json:"id,omitempty"`
	Email                   string                       `json:"email,omitempty"`
	FirstName               string                       `json:"first_name,omitempty"`
	LastName                string                       `json:"last_name,omitempty"`
	Company                 string                       `json:"company,omitempty"`
	Taxability              enum.Taxability              `json:"taxability,omitempty"`
	Locale                  string                       `json:"locale,omitempty"`
	EntityCode              enum.EntityCode              `json:"entity_code,omitempty"`
	ExemptNumber            string                       `json:"exempt_number,omitempty"`
	NetTermDays             *int32                       `json:"net_term_days,omitempty"`
	TaxjarExemptionCategory enum.TaxjarExemptionCategory `json:"taxjar_exemption_category,omitempty"`
	Phone                   string                       `json:"phone,omitempty"`
	CustomerType            enum.CustomerType            `json:"customer_type,omitempty"`
	AutoCollection          enum.AutoCollection          `json:"auto_collection,omitempty"`
	AllowDirectDebit        *bool                        `json:"allow_direct_debit,omitempty"`
	VatNumber               string                       `json:"vat_number,omitempty"`
	VatNumberPrefix         string                       `json:"vat_number_prefix,omitempty"`
}
type ImportSubscriptionAddonParams struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
}
type ImportSubscriptionEventBasedAddonParams struct {
	Id                  string       `json:"id,omitempty"`
	Quantity            *int32       `json:"quantity,omitempty"`
	UnitPrice           *int32       `json:"unit_price,omitempty"`
	QuantityInDecimal   string       `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string       `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodInDays *int32       `json:"service_period_in_days,omitempty"`
	OnEvent             enum.OnEvent `json:"on_event,omitempty"`
	ChargeOnce          *bool        `json:"charge_once,omitempty"`
}
type ImportSubscriptionChargedEventBasedAddonParams struct {
	Id            string `json:"id,omitempty"`
	LastChargedAt *int64 `json:"last_charged_at,omitempty"`
}
type ImportSubscriptionContractTermParams struct {
	Id                       string                           `json:"id,omitempty"`
	CreatedAt                *int64                           `json:"created_at,omitempty"`
	ContractStart            *int64                           `json:"contract_start,omitempty"`
	BillingCycle             *int32                           `json:"billing_cycle,omitempty"`
	TotalAmountRaised        *int64                           `json:"total_amount_raised,omitempty"`
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type ImportSubscriptionCardParams struct {
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	Number                string                 `json:"number,omitempty"`
	ExpiryMonth           *int32                 `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                 `json:"expiry_year,omitempty"`
	Cvv                   string                 `json:"cvv,omitempty"`
	BillingAddr1          string                 `json:"billing_addr1,omitempty"`
	BillingAddr2          string                 `json:"billing_addr2,omitempty"`
	BillingCity           string                 `json:"billing_city,omitempty"`
	BillingStateCode      string                 `json:"billing_state_code,omitempty"`
	BillingState          string                 `json:"billing_state,omitempty"`
	BillingZip            string                 `json:"billing_zip,omitempty"`
	BillingCountry        string                 `json:"billing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type ImportSubscriptionPaymentMethodParams struct {
	Type                  enum.Type              `json:"type,omitempty"`
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	ReferenceId           string                 `json:"reference_id,omitempty"`
	IssuingCountry        string                 `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type ImportSubscriptionBillingAddressParams struct {
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
type ImportSubscriptionShippingAddressParams struct {
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
type ImportSubscriptionTransactionParams struct {
	Amount          *int32             `json:"amount,omitempty"`
	PaymentMethod   enum.PaymentMethod `json:"payment_method,omitempty"`
	ReferenceNumber string             `json:"reference_number,omitempty"`
	Date            *int64             `json:"date,omitempty"`
}
type ImportForCustomerRequestParams struct {
	Id                                string                                           `json:"id,omitempty"`
	PlanId                            string                                           `json:"plan_id"`
	PlanQuantity                      *int32                                           `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                                           `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int32                                           `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                                           `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int32                                           `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                                           `json:"trial_end,omitempty"`
	BillingCycles                     *int32                                           `json:"billing_cycles,omitempty"`
	Addons                            []*ImportForCustomerAddonParams                  `json:"addons,omitempty"`
	EventBasedAddons                  []*ImportForCustomerEventBasedAddonParams        `json:"event_based_addons,omitempty"`
	ChargedEventBasedAddons           []*ImportForCustomerChargedEventBasedAddonParams `json:"charged_event_based_addons,omitempty"`
	StartDate                         *int64                                           `json:"start_date,omitempty"`
	AutoCollection                    enum.AutoCollection                              `json:"auto_collection,omitempty"`
	PoNumber                          string                                           `json:"po_number,omitempty"`
	CouponIds                         []string                                         `json:"coupon_ids,omitempty"`
	PaymentSourceId                   string                                           `json:"payment_source_id,omitempty"`
	Status                            subscriptionEnum.Status                          `json:"status"`
	CurrentTermEnd                    *int64                                           `json:"current_term_end,omitempty"`
	CurrentTermStart                  *int64                                           `json:"current_term_start,omitempty"`
	TrialStart                        *int64                                           `json:"trial_start,omitempty"`
	CancelledAt                       *int64                                           `json:"cancelled_at,omitempty"`
	StartedAt                         *int64                                           `json:"started_at,omitempty"`
	ActivatedAt                       *int64                                           `json:"activated_at,omitempty"`
	PauseDate                         *int64                                           `json:"pause_date,omitempty"`
	ResumeDate                        *int64                                           `json:"resume_date,omitempty"`
	ContractTerm                      *ImportForCustomerContractTermParams             `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                           `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	CreateCurrentTermInvoice          *bool                                            `json:"create_current_term_invoice,omitempty"`
	Transaction                       *ImportForCustomerTransactionParams              `json:"transaction,omitempty"`
	ShippingAddress                   *ImportForCustomerShippingAddressParams          `json:"shipping_address,omitempty"`
	InvoiceNotes                      string                                           `json:"invoice_notes,omitempty"`
	MetaData                          map[string]interface{}                           `json:"meta_data,omitempty"`
}
type ImportForCustomerAddonParams struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
}
type ImportForCustomerEventBasedAddonParams struct {
	Id                  string       `json:"id,omitempty"`
	Quantity            *int32       `json:"quantity,omitempty"`
	UnitPrice           *int32       `json:"unit_price,omitempty"`
	QuantityInDecimal   string       `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string       `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodInDays *int32       `json:"service_period_in_days,omitempty"`
	OnEvent             enum.OnEvent `json:"on_event,omitempty"`
	ChargeOnce          *bool        `json:"charge_once,omitempty"`
}
type ImportForCustomerChargedEventBasedAddonParams struct {
	Id            string `json:"id,omitempty"`
	LastChargedAt *int64 `json:"last_charged_at,omitempty"`
}
type ImportForCustomerContractTermParams struct {
	Id                       string                           `json:"id,omitempty"`
	CreatedAt                *int64                           `json:"created_at,omitempty"`
	ContractStart            *int64                           `json:"contract_start,omitempty"`
	BillingCycle             *int32                           `json:"billing_cycle,omitempty"`
	TotalAmountRaised        *int64                           `json:"total_amount_raised,omitempty"`
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type ImportForCustomerTransactionParams struct {
	Amount          *int32             `json:"amount,omitempty"`
	PaymentMethod   enum.PaymentMethod `json:"payment_method,omitempty"`
	ReferenceNumber string             `json:"reference_number,omitempty"`
	Date            *int64             `json:"date,omitempty"`
}
type ImportForCustomerShippingAddressParams struct {
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
type ImportContractTermRequestParams struct {
	ContractTerm                      *ImportContractTermContractTermParams `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type ImportContractTermContractTermParams struct {
	Id                       string                           `json:"id,omitempty"`
	CreatedAt                *int64                           `json:"created_at,omitempty"`
	ContractStart            *int64                           `json:"contract_start,omitempty"`
	ContractEnd              *int64                           `json:"contract_end,omitempty"`
	Status                   contractTermEnum.Status          `json:"status,omitempty"`
	TotalAmountRaised        *int64                           `json:"total_amount_raised,omitempty"`
	TotalContractValue       *int64                           `json:"total_contract_value,omitempty"`
	BillingCycle             *int32                           `json:"billing_cycle,omitempty"`
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type ImportForItemsRequestParams struct {
	Id                                string                                  `json:"id,omitempty"`
	TrialEnd                          *int64                                  `json:"trial_end,omitempty"`
	BillingCycles                     *int32                                  `json:"billing_cycles,omitempty"`
	SubscriptionItems                 []*ImportForItemsSubscriptionItemParams `json:"subscription_items,omitempty"`
	SetupFee                          *int32                                  `json:"setup_fee,omitempty"`
	ChargedItems                      []*ImportForItemsChargedItemParams      `json:"charged_items,omitempty"`
	ItemTiers                         []*ImportForItemsItemTierParams         `json:"item_tiers,omitempty"`
	NetTermDays                       *int32                                  `json:"net_term_days,omitempty"`
	StartDate                         *int64                                  `json:"start_date,omitempty"`
	AutoCollection                    enum.AutoCollection                     `json:"auto_collection,omitempty"`
	PoNumber                          string                                  `json:"po_number,omitempty"`
	CouponIds                         []string                                `json:"coupon_ids,omitempty"`
	PaymentSourceId                   string                                  `json:"payment_source_id,omitempty"`
	Status                            subscriptionEnum.Status                 `json:"status"`
	CurrentTermEnd                    *int64                                  `json:"current_term_end,omitempty"`
	CurrentTermStart                  *int64                                  `json:"current_term_start,omitempty"`
	TrialStart                        *int64                                  `json:"trial_start,omitempty"`
	CancelledAt                       *int64                                  `json:"cancelled_at,omitempty"`
	StartedAt                         *int64                                  `json:"started_at,omitempty"`
	ActivatedAt                       *int64                                  `json:"activated_at,omitempty"`
	PauseDate                         *int64                                  `json:"pause_date,omitempty"`
	ResumeDate                        *int64                                  `json:"resume_date,omitempty"`
	ContractTerm                      *ImportForItemsContractTermParams       `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                  `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	CreateCurrentTermInvoice          *bool                                   `json:"create_current_term_invoice,omitempty"`
	Transaction                       *ImportForItemsTransactionParams        `json:"transaction,omitempty"`
	ShippingAddress                   *ImportForItemsShippingAddressParams    `json:"shipping_address,omitempty"`
	InvoiceNotes                      string                                  `json:"invoice_notes,omitempty"`
	MetaData                          map[string]interface{}                  `json:"meta_data,omitempty"`
	CreatePendingInvoices             *bool                                   `json:"create_pending_invoices,omitempty"`
	AutoCloseInvoices                 *bool                                   `json:"auto_close_invoices,omitempty"`
}
type ImportForItemsSubscriptionItemParams struct {
	ItemPriceId        string             `json:"item_price_id"`
	Quantity           *int32             `json:"quantity,omitempty"`
	QuantityInDecimal  string             `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32             `json:"unit_price,omitempty"`
	UnitPriceInDecimal string             `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32             `json:"billing_cycles,omitempty"`
	TrialEnd           *int64             `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32             `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool              `json:"charge_once,omitempty"`
	ItemType           enum.ItemType      `json:"item_type,omitempty"`
}
type ImportForItemsChargedItemParams struct {
	ItemPriceId   string `json:"item_price_id,omitempty"`
	LastChargedAt *int64 `json:"last_charged_at,omitempty"`
}
type ImportForItemsItemTierParams struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type ImportForItemsContractTermParams struct {
	Id                       string                           `json:"id,omitempty"`
	CreatedAt                *int64                           `json:"created_at,omitempty"`
	ContractStart            *int64                           `json:"contract_start,omitempty"`
	BillingCycle             *int32                           `json:"billing_cycle,omitempty"`
	TotalAmountRaised        *int64                           `json:"total_amount_raised,omitempty"`
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type ImportForItemsTransactionParams struct {
	Amount          *int32             `json:"amount,omitempty"`
	PaymentMethod   enum.PaymentMethod `json:"payment_method,omitempty"`
	ReferenceNumber string             `json:"reference_number,omitempty"`
	Date            *int64             `json:"date,omitempty"`
}
type ImportForItemsShippingAddressParams struct {
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
type OverrideBillingProfileRequestParams struct {
	PaymentSourceId string              `json:"payment_source_id,omitempty"`
	AutoCollection  enum.AutoCollection `json:"auto_collection,omitempty"`
}
type PauseRequestParams struct {
	PauseOption             enum.PauseOption             `json:"pause_option,omitempty"`
	PauseDate               *int64                       `json:"pause_date,omitempty"`
	UnbilledChargesHandling enum.UnbilledChargesHandling `json:"unbilled_charges_handling,omitempty"`
	InvoiceDunningHandling  enum.InvoiceDunningHandling  `json:"invoice_dunning_handling,omitempty"`
	ResumeDate              *int64                       `json:"resume_date,omitempty"`
}
type CancelRequestParams struct {
	EndOfTerm                         *bool                                  `json:"end_of_term,omitempty"`
	CancelAt                          *int64                                 `json:"cancel_at,omitempty"`
	CreditOptionForCurrentTermCharges enum.CreditOptionForCurrentTermCharges `json:"credit_option_for_current_term_charges,omitempty"`
	UnbilledChargesOption             enum.UnbilledChargesOption             `json:"unbilled_charges_option,omitempty"`
	AccountReceivablesHandling        enum.AccountReceivablesHandling        `json:"account_receivables_handling,omitempty"`
	RefundableCreditsHandling         enum.RefundableCreditsHandling         `json:"refundable_credits_handling,omitempty"`
	ContractTermCancelOption          enum.ContractTermCancelOption          `json:"contract_term_cancel_option,omitempty"`
	InvoiceDate                       *int64                                 `json:"invoice_date,omitempty"`
	EventBasedAddons                  []*CancelEventBasedAddonParams         `json:"event_based_addons,omitempty"`
	CancelReasonCode                  string                                 `json:"cancel_reason_code,omitempty"`
}
type CancelEventBasedAddonParams struct {
	Id                  string `json:"id,omitempty"`
	Quantity            *int32 `json:"quantity,omitempty"`
	UnitPrice           *int32 `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32 `json:"service_period_in_days,omitempty"`
}
type CancelForItemsRequestParams struct {
	EndOfTerm                         *bool                                   `json:"end_of_term,omitempty"`
	CancelAt                          *int64                                  `json:"cancel_at,omitempty"`
	CreditOptionForCurrentTermCharges enum.CreditOptionForCurrentTermCharges  `json:"credit_option_for_current_term_charges,omitempty"`
	UnbilledChargesOption             enum.UnbilledChargesOption              `json:"unbilled_charges_option,omitempty"`
	AccountReceivablesHandling        enum.AccountReceivablesHandling         `json:"account_receivables_handling,omitempty"`
	RefundableCreditsHandling         enum.RefundableCreditsHandling          `json:"refundable_credits_handling,omitempty"`
	ContractTermCancelOption          enum.ContractTermCancelOption           `json:"contract_term_cancel_option,omitempty"`
	InvoiceDate                       *int64                                  `json:"invoice_date,omitempty"`
	SubscriptionItems                 []*CancelForItemsSubscriptionItemParams `json:"subscription_items,omitempty"`
	CancelReasonCode                  string                                  `json:"cancel_reason_code,omitempty"`
}
type CancelForItemsSubscriptionItemParams struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodDays  *int32 `json:"service_period_days,omitempty"`
}
type ResumeRequestParams struct {
	ResumeOption           enum.ResumeOption           `json:"resume_option,omitempty"`
	ResumeDate             *int64                      `json:"resume_date,omitempty"`
	ChargesHandling        enum.ChargesHandling        `json:"charges_handling,omitempty"`
	UnpaidInvoicesHandling enum.UnpaidInvoicesHandling `json:"unpaid_invoices_handling,omitempty"`
	PaymentIntent          *ResumePaymentIntentParams  `json:"payment_intent,omitempty"`
}
type ResumePaymentIntentParams struct {
	Id                    string                              `json:"id,omitempty"`
	GatewayAccountId      string                              `json:"gateway_account_id,omitempty"`
	GwToken               string                              `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntentEnum.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                              `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                              `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}              `json:"additional_information,omitempty"`
}
