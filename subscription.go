package chargebee

type Status string

const (
	StatusFuture      Status = "future"
	StatusInTrial     Status = "in_trial"
	StatusActive      Status = "active"
	StatusNonRenewing Status = "non_renewing"
	StatusPaused      Status = "paused"
	StatusCancelled   Status = "cancelled"
	StatusTransferred Status = "transferred"
)

type CancelReason string

const (
	CancelReasonNotPaid                         CancelReason = "not_paid"
	CancelReasonNoCard                          CancelReason = "no_card"
	CancelReasonFraudReviewFailed               CancelReason = "fraud_review_failed"
	CancelReasonNonCompliantEuCustomer          CancelReason = "non_compliant_eu_customer"
	CancelReasonTaxCalculationFailed            CancelReason = "tax_calculation_failed"
	CancelReasonCurrencyIncompatibleWithGateway CancelReason = "currency_incompatible_with_gateway"
	CancelReasonNonCompliantCustomer            CancelReason = "non_compliant_customer"
)

type BillingPeriodUnit string

const (
	BillingPeriodUnitDay   BillingPeriodUnit = "day"
	BillingPeriodUnitWeek  BillingPeriodUnit = "week"
	BillingPeriodUnitMonth BillingPeriodUnit = "month"
	BillingPeriodUnitYear  BillingPeriodUnit = "year"
)

type ReferralInfoRewardStatus string

const (
	ReferralInfoRewardStatusPending ReferralInfoRewardStatus = "pending"
	ReferralInfoRewardStatusPaid    ReferralInfoRewardStatus = "paid"
	ReferralInfoRewardStatusInvalid ReferralInfoRewardStatus = "invalid"
)

type ContractTermStatus string

const (
	ContractTermStatusActive     ContractTermStatus = "active"
	ContractTermStatusCompleted  ContractTermStatus = "completed"
	ContractTermStatusCancelled  ContractTermStatus = "cancelled"
	ContractTermStatusTerminated ContractTermStatus = "terminated"
)

type ContractTermActionAtTermEnd string

const (
	ContractTermActionAtTermEndRenew     ContractTermActionAtTermEnd = "renew"
	ContractTermActionAtTermEndEvergreen ContractTermActionAtTermEnd = "evergreen"
	ContractTermActionAtTermEndCancel    ContractTermActionAtTermEnd = "cancel"
	ContractTermActionAtTermEndRenewOnce ContractTermActionAtTermEnd = "renew_once"
)

type DiscountType string

const (
	DiscountTypeFixedAmount   DiscountType = "fixed_amount"
	DiscountTypePercentage    DiscountType = "percentage"
	DiscountTypeOfferQuantity DiscountType = "offer_quantity"
)

type Subscription struct {
	Id                                string                    `json:"id"`
	CurrencyCode                      string                    `json:"currency_code"`
	PlanId                            string                    `json:"plan_id"`
	PlanQuantity                      int32                     `json:"plan_quantity"`
	PlanUnitPrice                     int64                     `json:"plan_unit_price"`
	SetupFee                          int64                     `json:"setup_fee"`
	BillingPeriod                     int32                     `json:"billing_period"`
	BillingPeriodUnit                 BillingPeriodUnit         `json:"billing_period_unit"`
	StartDate                         int64                     `json:"start_date"`
	TrialEnd                          int64                     `json:"trial_end"`
	RemainingBillingCycles            int32                     `json:"remaining_billing_cycles"`
	PoNumber                          string                    `json:"po_number"`
	AutoCollection                    enum.AutoCollection       `json:"auto_collection"`
	PlanQuantityInDecimal             string                    `json:"plan_quantity_in_decimal"`
	PlanUnitPriceInDecimal            string                    `json:"plan_unit_price_in_decimal"`
	CustomerId                        string                    `json:"customer_id"`
	PlanAmount                        int64                     `json:"plan_amount"`
	PlanFreeQuantity                  int32                     `json:"plan_free_quantity"`
	Status                            Status                    `json:"status"`
	TrialStart                        int64                     `json:"trial_start"`
	TrialEndAction                    enum.TrialEndAction       `json:"trial_end_action"`
	CurrentTermStart                  int64                     `json:"current_term_start"`
	CurrentTermEnd                    int64                     `json:"current_term_end"`
	NextBillingAt                     int64                     `json:"next_billing_at"`
	CreatedAt                         int64                     `json:"created_at"`
	StartedAt                         int64                     `json:"started_at"`
	ActivatedAt                       int64                     `json:"activated_at"`
	GiftId                            string                    `json:"gift_id"`
	ContractTermBillingCycleOnRenewal int32                     `json:"contract_term_billing_cycle_on_renewal"`
	OverrideRelationship              bool                      `json:"override_relationship"`
	PauseDate                         int64                     `json:"pause_date"`
	ResumeDate                        int64                     `json:"resume_date"`
	CancelledAt                       int64                     `json:"cancelled_at"`
	CancelReason                      CancelReason              `json:"cancel_reason"`
	AffiliateToken                    string                    `json:"affiliate_token"`
	CreatedFromIp                     string                    `json:"created_from_ip"`
	ResourceVersion                   int64                     `json:"resource_version"`
	UpdatedAt                         int64                     `json:"updated_at"`
	HasScheduledAdvanceInvoices       bool                      `json:"has_scheduled_advance_invoices"`
	HasScheduledChanges               bool                      `json:"has_scheduled_changes"`
	PaymentSourceId                   string                    `json:"payment_source_id"`
	PlanFreeQuantityInDecimal         string                    `json:"plan_free_quantity_in_decimal"`
	PlanAmountInDecimal               string                    `json:"plan_amount_in_decimal"`
	CancelScheduleCreatedAt           int64                     `json:"cancel_schedule_created_at"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method"`
	Channel                           enum.Channel              `json:"channel"`
	NetTermDays                       int32                     `json:"net_term_days"`
	ActiveId                          string                    `json:"active_id"`
	SubscriptionItems                 []*SubscriptionItem       `json:"subscription_items"`
	ItemTiers                         []*ItemTier               `json:"item_tiers"`
	ChargedItems                      []*ChargedItem            `json:"charged_items"`
	DueInvoicesCount                  int32                     `json:"due_invoices_count"`
	DueSince                          int64                     `json:"due_since"`
	TotalDues                         int64                     `json:"total_dues"`
	Mrr                               int64                     `json:"mrr"`
	Arr                               int64                     `json:"arr"`
	ExchangeRate                      float64                   `json:"exchange_rate"`
	BaseCurrencyCode                  string                    `json:"base_currency_code"`
	Addons                            []*Addon                  `json:"addons"`
	EventBasedAddons                  []*EventBasedAddon        `json:"event_based_addons"`
	ChargedEventBasedAddons           []*ChargedEventBasedAddon `json:"charged_event_based_addons"`
	Coupon                            string                    `json:"coupon"`
	Coupons                           []*Coupon                 `json:"coupons"`
	ShippingAddress                   *ShippingAddress          `json:"shipping_address"`
	ReferralInfo                      *ReferralInfo             `json:"referral_info"`
	BillingOverride                   *BillingOverride          `json:"billing_override"`
	InvoiceNotes                      string                    `json:"invoice_notes"`
	MetaData                          json.RawMessage           `json:"meta_data"`
	Deleted                           bool                      `json:"deleted"`
	ChangesScheduledAt                int64                     `json:"changes_scheduled_at"`
	ContractTerm                      *ContractTerm             `json:"contract_term"`
	CancelReasonCode                  string                    `json:"cancel_reason_code"`
	FreePeriod                        int32                     `json:"free_period"`
	FreePeriodUnit                    enum.FreePeriodUnit       `json:"free_period_unit"`
	CreatePendingInvoices             bool                      `json:"create_pending_invoices"`
	AutoCloseInvoices                 bool                      `json:"auto_close_invoices"`
	Discounts                         []*Discount               `json:"discounts"`
	BusinessEntityId                  string                    `json:"business_entity_id"`
	CustomField                       map[string]interface{}    `json:"custom_field"`
	Object                            string                    `json:"object"`

	// Deprecated: MetaData is deprecated please use MetaData instead.
	Metadata json.RawMessage `json:"metadata"`
}
type SubscriptionItem struct {
	ItemPriceId                     string                               `json:"item_price_id"`
	ItemType                        enum.ItemType                        `json:"item_type"`
	Quantity                        int32                                `json:"quantity"`
	QuantityInDecimal               string                               `json:"quantity_in_decimal"`
	MeteredQuantity                 string                               `json:"metered_quantity"`
	LastCalculatedAt                int64                                `json:"last_calculated_at"`
	UnitPrice                       int64                                `json:"unit_price"`
	UnitPriceInDecimal              string                               `json:"unit_price_in_decimal"`
	Amount                          int64                                `json:"amount"`
	CurrentTermStart                int64                                `json:"current_term_start"`
	CurrentTermEnd                  int64                                `json:"current_term_end"`
	NextBillingAt                   int64                                `json:"next_billing_at"`
	AmountInDecimal                 string                               `json:"amount_in_decimal"`
	BillingPeriod                   int32                                `json:"billing_period"`
	BillingPeriodUnit               subscriptionEnum.BillingPeriodUnit   `json:"billing_period_unit"`
	FreeQuantity                    int32                                `json:"free_quantity"`
	FreeQuantityInDecimal           string                               `json:"free_quantity_in_decimal"`
	TrialEnd                        int64                                `json:"trial_end"`
	BillingCycles                   int32                                `json:"billing_cycles"`
	ServicePeriodDays               int32                                `json:"service_period_days"`
	ChargeOnEvent                   enum.ChargeOnEvent                   `json:"charge_on_event"`
	ChargeOnce                      bool                                 `json:"charge_once"`
	ChargeOnOption                  enum.ChargeOnOption                  `json:"charge_on_option"`
	ProrationType                   enum.ProrationType                   `json:"proration_type"`
	UsageAccumulationResetFrequency enum.UsageAccumulationResetFrequency `json:"usage_accumulation_reset_frequency"`
	Object                          string                               `json:"object"`
}
type ItemTier struct {
	ItemPriceId           string           `json:"item_price_id"`
	StartingUnit          int32            `json:"starting_unit"`
	EndingUnit            int32            `json:"ending_unit"`
	Price                 int64            `json:"price"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal"`
	PriceInDecimal        string           `json:"price_in_decimal"`
	PricingType           enum.PricingType `json:"pricing_type"`
	PackageSize           int32            `json:"package_size"`
	Index                 int32            `json:"index"`
	Object                string           `json:"object"`
}
type ChargedItem struct {
	ItemPriceId   string `json:"item_price_id"`
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
	ReferralCode              string                    `json:"referral_code"`
	CouponCode                string                    `json:"coupon_code"`
	ReferrerId                string                    `json:"referrer_id"`
	ExternalReferenceId       string                    `json:"external_reference_id"`
	RewardStatus              ReferralInfoRewardStatus  `json:"reward_status"`
	ReferralSystem            enum.ReferralSystem       `json:"referral_system"`
	AccountId                 string                    `json:"account_id"`
	CampaignId                string                    `json:"campaign_id"`
	ExternalCampaignId        string                    `json:"external_campaign_id"`
	FriendOfferType           enum.FriendOfferType      `json:"friend_offer_type"`
	ReferrerRewardType        enum.ReferrerRewardType   `json:"referrer_reward_type"`
	NotifyReferralSystem      enum.NotifyReferralSystem `json:"notify_referral_system"`
	DestinationUrl            string                    `json:"destination_url"`
	PostPurchaseWidgetEnabled bool                      `json:"post_purchase_widget_enabled"`
	Object                    string                    `json:"object"`
}
type BillingOverride struct {
	MaxExcessPaymentUsage     int64  `json:"max_excess_payment_usage"`
	MaxRefundableCreditsUsage int64  `json:"max_refundable_credits_usage"`
	Object                    string `json:"object"`
}
type ContractTerm struct {
	Id                          string                      `json:"id"`
	Status                      ContractTermStatus          `json:"status"`
	ContractStart               int64                       `json:"contract_start"`
	ContractEnd                 int64                       `json:"contract_end"`
	BillingCycle                int32                       `json:"billing_cycle"`
	ActionAtTermEnd             ContractTermActionAtTermEnd `json:"action_at_term_end"`
	TotalContractValue          int64                       `json:"total_contract_value"`
	TotalContractValueBeforeTax int64                       `json:"total_contract_value_before_tax"`
	CancellationCutoffPeriod    int32                       `json:"cancellation_cutoff_period"`
	CreatedAt                   int64                       `json:"created_at"`
	SubscriptionId              string                      `json:"subscription_id"`
	RemainingBillingCycles      int32                       `json:"remaining_billing_cycles"`
	Object                      string                      `json:"object"`
}
type Discount struct {
	Id            string            `json:"id"`
	InvoiceName   string            `json:"invoice_name"`
	Type          DiscountType      `json:"type"`
	Percentage    float64           `json:"percentage"`
	Amount        int64             `json:"amount"`
	Quantity      int32             `json:"quantity"`
	CurrencyCode  string            `json:"currency_code"`
	DurationType  enum.DurationType `json:"duration_type"`
	Period        int32             `json:"period"`
	PeriodUnit    enum.PeriodUnit   `json:"period_unit"`
	IncludedInMrr bool              `json:"included_in_mrr"`
	ApplyOn       enum.ApplyOn      `json:"apply_on"`
	ItemPriceId   string            `json:"item_price_id"`
	CreatedAt     int64             `json:"created_at"`
	ApplyTill     int64             `json:"apply_till"`
	AppliedCount  int32             `json:"applied_count"`
	CouponId      string            `json:"coupon_id"`
	Index         int32             `json:"index"`
	Object        string            `json:"object"`
}
type Addon struct {
	Id                     string             `json:"id"`
	Quantity               int32              `json:"quantity"`
	UnitPrice              int64              `json:"unit_price"`
	Amount                 int64              `json:"amount"`
	TrialEnd               int64              `json:"trial_end"`
	RemainingBillingCycles int32              `json:"remaining_billing_cycles"`
	QuantityInDecimal      string             `json:"quantity_in_decimal"`
	UnitPriceInDecimal     string             `json:"unit_price_in_decimal"`
	AmountInDecimal        string             `json:"amount_in_decimal"`
	ProrationType          enum.ProrationType `json:"proration_type"`
	Object                 string             `json:"object"`
}
type ChargedEventBasedAddon struct {
	Id            string `json:"id"`
	LastChargedAt int64  `json:"last_charged_at"`
	Object        string `json:"object"`
}
type EventBasedAddon struct {
	Id                  string       `json:"id"`
	Quantity            int32        `json:"quantity"`
	UnitPrice           int64        `json:"unit_price"`
	ServicePeriodInDays int32        `json:"service_period_in_days"`
	OnEvent             enum.OnEvent `json:"on_event"`
	ChargeOnce          bool         `json:"charge_once"`
	QuantityInDecimal   string       `json:"quantity_in_decimal"`
	UnitPriceInDecimal  string       `json:"unit_price_in_decimal"`
	Object              string       `json:"object"`
}
type CreateRequest struct {
	Id                                string                     `json:"id,omitempty"`
	Customer                          *CreateCustomer            `json:"customer,omitempty"`
	EntityIdentifiers                 []*CreateEntityIdentifier  `json:"entity_identifiers,omitempty"`
	TaxProvidersFields                []*CreateTaxProvidersField `json:"tax_providers_fields,omitempty"`
	PlanId                            string                     `json:"plan_id"`
	PlanQuantity                      *int32                     `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                     `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int64                     `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                     `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int64                     `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                     `json:"trial_end,omitempty"`
	BillingCycles                     *int32                     `json:"billing_cycles,omitempty"`
	Addons                            []*CreateAddon             `json:"addons,omitempty"`
	EventBasedAddons                  []*CreateEventBasedAddon   `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove           []string                   `json:"mandatory_addons_to_remove,omitempty"`
	StartDate                         *int64                     `json:"start_date,omitempty"`
	Coupon                            string                     `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection        `json:"auto_collection,omitempty"`
	TermsToCharge                     *int32                     `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode              enum.BillingAlignmentMode  `json:"billing_alignment_mode,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod  `json:"offline_payment_method,omitempty"`
	PoNumber                          string                     `json:"po_number,omitempty"`
	CouponIds                         []string                   `json:"coupon_ids,omitempty"`
	Card                              *CreateCard                `json:"card,omitempty"`
	BankAccount                       *CreateBankAccount         `json:"bank_account,omitempty"`
	TokenId                           string                     `json:"token_id,omitempty"`
	PaymentMethod                     *CreatePaymentMethod       `json:"payment_method,omitempty"`
	PaymentIntent                     *CreatePaymentIntent       `json:"payment_intent,omitempty"`
	BillingAddress                    *CreateBillingAddress      `json:"billing_address,omitempty"`
	ShippingAddress                   *CreateShippingAddress     `json:"shipping_address,omitempty"`
	StatementDescriptor               *CreateStatementDescriptor `json:"statement_descriptor,omitempty"`
	AffiliateToken                    string                     `json:"affiliate_token,omitempty"`
	CreatedFromIp                     string                     `json:"created_from_ip,omitempty"`
	InvoiceNotes                      string                     `json:"invoice_notes,omitempty"`
	InvoiceDate                       *int64                     `json:"invoice_date,omitempty"`
	MetaData                          map[string]interface{}     `json:"meta_data,omitempty"`
	InvoiceImmediately                *bool                      `json:"invoice_immediately,omitempty"`
	FreePeriod                        *int32                     `json:"free_period,omitempty"`
	FreePeriodUnit                    enum.FreePeriodUnit        `json:"free_period_unit,omitempty"`
	ContractTerm                      *CreateContractTerm        `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                     `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	TrialEndAction                    enum.TrialEndAction        `json:"trial_end_action,omitempty"`
	ClientProfileId                   string                     `json:"client_profile_id,omitempty"`
	PaymentInitiator                  enum.PaymentInitiator      `json:"payment_initiator,omitempty"`
	Coupons                           []*CreateCoupon            `json:"coupons,omitempty"`
}
type CreateCustomer struct {
	Id                               string                       `json:"id,omitempty"`
	Email                            string                       `json:"email,omitempty"`
	FirstName                        string                       `json:"first_name,omitempty"`
	LastName                         string                       `json:"last_name,omitempty"`
	Company                          string                       `json:"company,omitempty"`
	Phone                            string                       `json:"phone,omitempty"`
	Locale                           string                       `json:"locale,omitempty"`
	Taxability                       enum.Taxability              `json:"taxability,omitempty"`
	EntityCode                       enum.EntityCode              `json:"entity_code,omitempty"`
	ExemptNumber                     string                       `json:"exempt_number,omitempty"`
	NetTermDays                      *int32                       `json:"net_term_days,omitempty"`
	TaxjarExemptionCategory          enum.TaxjarExemptionCategory `json:"taxjar_exemption_category,omitempty"`
	AutoCollection                   enum.AutoCollection          `json:"auto_collection,omitempty"`
	OfflinePaymentMethod             enum.OfflinePaymentMethod    `json:"offline_payment_method,omitempty"`
	AllowDirectDebit                 *bool                        `json:"allow_direct_debit,omitempty"`
	ConsolidatedInvoicing            *bool                        `json:"consolidated_invoicing,omitempty"`
	VatNumber                        string                       `json:"vat_number,omitempty"`
	VatNumberPrefix                  string                       `json:"vat_number_prefix,omitempty"`
	EntityIdentifierScheme           string                       `json:"entity_identifier_scheme,omitempty"`
	EntityIdentifierStandard         string                       `json:"entity_identifier_standard,omitempty"`
	IsEinvoiceEnabled                *bool                        `json:"is_einvoice_enabled,omitempty"`
	EinvoicingMethod                 enum.EinvoicingMethod        `json:"einvoicing_method,omitempty"`
	RegisteredForGst                 *bool                        `json:"registered_for_gst,omitempty"`
	BusinessCustomerWithoutVatNumber *bool                        `json:"business_customer_without_vat_number,omitempty"`
	ExemptionDetails                 []map[string]interface{}     `json:"exemption_details,omitempty"`
	CustomerType                     enum.CustomerType            `json:"customer_type,omitempty"`
}
type CreateEntityIdentifier struct {
	Id       string `json:"id,omitempty"`
	Scheme   string `json:"scheme,omitempty"`
	Value    string `json:"value,omitempty"`
	Standard string `json:"standard,omitempty"`
}
type CreateTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type CreateAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}
type CreateEventBasedAddon struct {
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
type CreateCard struct {
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	Number                string                 `json:"number,omitempty"`
	ExpiryMonth           *int32                 `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                 `json:"expiry_year,omitempty"`
	Cvv                   string                 `json:"cvv,omitempty"`
	PreferredScheme       card.PreferredScheme   `json:"preferred_scheme,omitempty"`
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
type CreateBankAccount struct {
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
type CreatePaymentMethod struct {
	Type                  enum.Type              `json:"type,omitempty"`
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	ReferenceId           string                 `json:"reference_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	IssuingCountry        string                 `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type CreatePaymentIntent struct {
	Id                    string                          `json:"id,omitempty"`
	GatewayAccountId      string                          `json:"gateway_account_id,omitempty"`
	GwToken               string                          `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntent.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                          `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                          `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}          `json:"additional_information,omitempty"`
}
type CreateBillingAddress struct {
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
type CreateShippingAddress struct {
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
type CreateStatementDescriptor struct {
	Descriptor string `json:"descriptor,omitempty"`
}
type CreateContractTerm struct {
	ActionAtTermEnd          subscription.ContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                                   `json:"cancellation_cutoff_period,omitempty"`
}
type CreateCoupon struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}
type CreateForCustomerRequest struct {
	Id                                string                                `json:"id,omitempty"`
	PlanId                            string                                `json:"plan_id"`
	PlanQuantity                      *int32                                `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                                `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int64                                `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                                `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int64                                `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                                `json:"trial_end,omitempty"`
	BillingCycles                     *int32                                `json:"billing_cycles,omitempty"`
	Addons                            []*CreateForCustomerAddon             `json:"addons,omitempty"`
	EventBasedAddons                  []*CreateForCustomerEventBasedAddon   `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove           []string                              `json:"mandatory_addons_to_remove,omitempty"`
	StartDate                         *int64                                `json:"start_date,omitempty"`
	Coupon                            string                                `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection                   `json:"auto_collection,omitempty"`
	TermsToCharge                     *int32                                `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode              enum.BillingAlignmentMode             `json:"billing_alignment_mode,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod             `json:"offline_payment_method,omitempty"`
	PoNumber                          string                                `json:"po_number,omitempty"`
	CouponIds                         []string                              `json:"coupon_ids,omitempty"`
	PaymentSourceId                   string                                `json:"payment_source_id,omitempty"`
	OverrideRelationship              *bool                                 `json:"override_relationship,omitempty"`
	ShippingAddress                   *CreateForCustomerShippingAddress     `json:"shipping_address,omitempty"`
	StatementDescriptor               *CreateForCustomerStatementDescriptor `json:"statement_descriptor,omitempty"`
	InvoiceNotes                      string                                `json:"invoice_notes,omitempty"`
	InvoiceDate                       *int64                                `json:"invoice_date,omitempty"`
	MetaData                          map[string]interface{}                `json:"meta_data,omitempty"`
	InvoiceImmediately                *bool                                 `json:"invoice_immediately,omitempty"`
	ReplacePrimaryPaymentSource       *bool                                 `json:"replace_primary_payment_source,omitempty"`
	PaymentIntent                     *CreateForCustomerPaymentIntent       `json:"payment_intent,omitempty"`
	FreePeriod                        *int32                                `json:"free_period,omitempty"`
	FreePeriodUnit                    enum.FreePeriodUnit                   `json:"free_period_unit,omitempty"`
	ContractTerm                      *CreateForCustomerContractTerm        `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	TrialEndAction                    enum.TrialEndAction                   `json:"trial_end_action,omitempty"`
	PaymentInitiator                  enum.PaymentInitiator                 `json:"payment_initiator,omitempty"`
	Coupons                           []*CreateForCustomerCoupon            `json:"coupons,omitempty"`
}
type CreateForCustomerAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}
type CreateForCustomerEventBasedAddon struct {
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
type CreateForCustomerShippingAddress struct {
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
type CreateForCustomerStatementDescriptor struct {
	Descriptor string `json:"descriptor,omitempty"`
}
type CreateForCustomerPaymentIntent struct {
	Id                    string                          `json:"id,omitempty"`
	GatewayAccountId      string                          `json:"gateway_account_id,omitempty"`
	GwToken               string                          `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntent.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                          `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                          `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}          `json:"additional_information,omitempty"`
}
type CreateForCustomerContractTerm struct {
	ActionAtTermEnd          subscription.ContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                                   `json:"cancellation_cutoff_period,omitempty"`
}
type CreateForCustomerCoupon struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}
type CreateWithItemsRequest struct {
	Id                                string                              `json:"id,omitempty"`
	BusinessEntityId                  string                              `json:"business_entity_id,omitempty"`
	TrialEnd                          *int64                              `json:"trial_end,omitempty"`
	BillingCycles                     *int32                              `json:"billing_cycles,omitempty"`
	SubscriptionItems                 []*CreateWithItemsSubscriptionItem  `json:"subscription_items,omitempty"`
	SetupFee                          *int64                              `json:"setup_fee,omitempty"`
	Discounts                         []*CreateWithItemsDiscount          `json:"discounts,omitempty"`
	MandatoryItemsToRemove            []string                            `json:"mandatory_items_to_remove,omitempty"`
	ItemTiers                         []*CreateWithItemsItemTier          `json:"item_tiers,omitempty"`
	NetTermDays                       *int32                              `json:"net_term_days,omitempty"`
	StartDate                         *int64                              `json:"start_date,omitempty"`
	Coupon                            string                              `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection                 `json:"auto_collection,omitempty"`
	TermsToCharge                     *int32                              `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode              enum.BillingAlignmentMode           `json:"billing_alignment_mode,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod           `json:"offline_payment_method,omitempty"`
	PoNumber                          string                              `json:"po_number,omitempty"`
	CouponIds                         []string                            `json:"coupon_ids,omitempty"`
	PaymentSourceId                   string                              `json:"payment_source_id,omitempty"`
	OverrideRelationship              *bool                               `json:"override_relationship,omitempty"`
	ShippingAddress                   *CreateWithItemsShippingAddress     `json:"shipping_address,omitempty"`
	StatementDescriptor               *CreateWithItemsStatementDescriptor `json:"statement_descriptor,omitempty"`
	InvoiceNotes                      string                              `json:"invoice_notes,omitempty"`
	InvoiceDate                       *int64                              `json:"invoice_date,omitempty"`
	MetaData                          map[string]interface{}              `json:"meta_data,omitempty"`
	InvoiceImmediately                *bool                               `json:"invoice_immediately,omitempty"`
	ReplacePrimaryPaymentSource       *bool                               `json:"replace_primary_payment_source,omitempty"`
	PaymentIntent                     *CreateWithItemsPaymentIntent       `json:"payment_intent,omitempty"`
	FreePeriod                        *int32                              `json:"free_period,omitempty"`
	FreePeriodUnit                    enum.FreePeriodUnit                 `json:"free_period_unit,omitempty"`
	ContractTerm                      *CreateWithItemsContractTerm        `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                              `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	CreatePendingInvoices             *bool                               `json:"create_pending_invoices,omitempty"`
	AutoCloseInvoices                 *bool                               `json:"auto_close_invoices,omitempty"`
	FirstInvoicePending               *bool                               `json:"first_invoice_pending,omitempty"`
	TrialEndAction                    enum.TrialEndAction                 `json:"trial_end_action,omitempty"`
	PaymentInitiator                  enum.PaymentInitiator               `json:"payment_initiator,omitempty"`
	Coupons                           []*CreateWithItemsCoupon            `json:"coupons,omitempty"`
	BillingOverride                   *CreateWithItemsBillingOverride     `json:"billing_override,omitempty"`
}
type CreateWithItemsSubscriptionItem struct {
	ItemPriceId                     string                               `json:"item_price_id"`
	Quantity                        *int32                               `json:"quantity,omitempty"`
	QuantityInDecimal               string                               `json:"quantity_in_decimal,omitempty"`
	UnitPrice                       *int64                               `json:"unit_price,omitempty"`
	UnitPriceInDecimal              string                               `json:"unit_price_in_decimal,omitempty"`
	BillingCycles                   *int32                               `json:"billing_cycles,omitempty"`
	TrialEnd                        *int64                               `json:"trial_end,omitempty"`
	ServicePeriodDays               *int32                               `json:"service_period_days,omitempty"`
	ChargeOnEvent                   enum.ChargeOnEvent                   `json:"charge_on_event,omitempty"`
	ChargeOnce                      *bool                                `json:"charge_once,omitempty"`
	ItemType                        enum.ItemType                        `json:"item_type,omitempty"`
	ChargeOnOption                  enum.ChargeOnOption                  `json:"charge_on_option,omitempty"`
	UsageAccumulationResetFrequency enum.UsageAccumulationResetFrequency `json:"usage_accumulation_reset_frequency,omitempty"`
}
type CreateWithItemsDiscount struct {
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
type CreateWithItemsItemTier struct {
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
type CreateWithItemsShippingAddress struct {
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
type CreateWithItemsStatementDescriptor struct {
	Descriptor string `json:"descriptor,omitempty"`
}
type CreateWithItemsPaymentIntent struct {
	Id                    string                          `json:"id,omitempty"`
	GatewayAccountId      string                          `json:"gateway_account_id,omitempty"`
	GwToken               string                          `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntent.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                          `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                          `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}          `json:"additional_information,omitempty"`
}
type CreateWithItemsContractTerm struct {
	ActionAtTermEnd          subscription.ContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	ContractStart            *int64                                   `json:"contract_start,omitempty"`
	CancellationCutoffPeriod *int32                                   `json:"cancellation_cutoff_period,omitempty"`
}
type CreateWithItemsCoupon struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}
type CreateWithItemsBillingOverride struct {
	MaxExcessPaymentUsage     *int64 `json:"max_excess_payment_usage,omitempty"`
	MaxRefundableCreditsUsage *int64 `json:"max_refundable_credits_usage,omitempty"`
}
type ListRequest struct {
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
	BusinessEntityId       *filter.StringFilter    `json:"business_entity_id,omitempty"`
	Channel                *filter.EnumFilter      `json:"channel,omitempty"`
}
type SubscriptionsForCustomerRequest struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
type ContractTermsForSubscriptionRequest struct {
	Limit  *int32             `json:"limit,omitempty"`
	Offset string             `json:"offset,omitempty"`
	SortBy *filter.SortFilter `json:"sort_by,omitempty"`
}
type ListDiscountsRequest struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
type RemoveScheduledCancellationRequest struct {
	BillingCycles                     *int32                                   `json:"billing_cycles,omitempty"`
	ContractTerm                      *RemoveScheduledCancellationContractTerm `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                   `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type RemoveScheduledCancellationContractTerm struct {
	ActionAtTermEnd          subscription.ContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                                   `json:"cancellation_cutoff_period,omitempty"`
}
type RemoveCouponsRequest struct {
	CouponIds []string `json:"coupon_ids,omitempty"`
}
type UpdateRequest struct {
	PlanId                            string                     `json:"plan_id,omitempty"`
	PlanQuantity                      *int32                     `json:"plan_quantity,omitempty"`
	PlanUnitPrice                     *int64                     `json:"plan_unit_price,omitempty"`
	SetupFee                          *int64                     `json:"setup_fee,omitempty"`
	Addons                            []*UpdateAddon             `json:"addons,omitempty"`
	EventBasedAddons                  []*UpdateEventBasedAddon   `json:"event_based_addons,omitempty"`
	ReplaceAddonList                  *bool                      `json:"replace_addon_list,omitempty"`
	MandatoryAddonsToRemove           []string                   `json:"mandatory_addons_to_remove,omitempty"`
	PlanQuantityInDecimal             string                     `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPriceInDecimal            string                     `json:"plan_unit_price_in_decimal,omitempty"`
	InvoiceDate                       *int64                     `json:"invoice_date,omitempty"`
	StartDate                         *int64                     `json:"start_date,omitempty"`
	TrialEnd                          *int64                     `json:"trial_end,omitempty"`
	BillingCycles                     *int32                     `json:"billing_cycles,omitempty"`
	Coupon                            string                     `json:"coupon,omitempty"`
	TermsToCharge                     *int32                     `json:"terms_to_charge,omitempty"`
	ReactivateFrom                    *int64                     `json:"reactivate_from,omitempty"`
	BillingAlignmentMode              enum.BillingAlignmentMode  `json:"billing_alignment_mode,omitempty"`
	AutoCollection                    enum.AutoCollection        `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod  `json:"offline_payment_method,omitempty"`
	PoNumber                          string                     `json:"po_number,omitempty"`
	CouponIds                         []string                   `json:"coupon_ids,omitempty"`
	ReplaceCouponList                 *bool                      `json:"replace_coupon_list,omitempty"`
	Prorate                           *bool                      `json:"prorate,omitempty"`
	EndOfTerm                         *bool                      `json:"end_of_term,omitempty"`
	ForceTermReset                    *bool                      `json:"force_term_reset,omitempty"`
	Reactivate                        *bool                      `json:"reactivate,omitempty"`
	Card                              *UpdateCard                `json:"card,omitempty"`
	TokenId                           string                     `json:"token_id,omitempty"`
	PaymentMethod                     *UpdatePaymentMethod       `json:"payment_method,omitempty"`
	PaymentIntent                     *UpdatePaymentIntent       `json:"payment_intent,omitempty"`
	BillingAddress                    *UpdateBillingAddress      `json:"billing_address,omitempty"`
	ShippingAddress                   *UpdateShippingAddress     `json:"shipping_address,omitempty"`
	StatementDescriptor               *UpdateStatementDescriptor `json:"statement_descriptor,omitempty"`
	Customer                          *UpdateCustomer            `json:"customer,omitempty"`
	InvoiceNotes                      string                     `json:"invoice_notes,omitempty"`
	MetaData                          map[string]interface{}     `json:"meta_data,omitempty"`
	InvoiceImmediately                *bool                      `json:"invoice_immediately,omitempty"`
	OverrideRelationship              *bool                      `json:"override_relationship,omitempty"`
	ChangesScheduledAt                *int64                     `json:"changes_scheduled_at,omitempty"`
	ChangeOption                      enum.ChangeOption          `json:"change_option,omitempty"`
	ContractTerm                      *UpdateContractTerm        `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                     `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	FreePeriod                        *int32                     `json:"free_period,omitempty"`
	FreePeriodUnit                    enum.FreePeriodUnit        `json:"free_period_unit,omitempty"`
	TrialEndAction                    enum.TrialEndAction        `json:"trial_end_action,omitempty"`
	Coupons                           []*UpdateCoupon            `json:"coupons,omitempty"`
}
type UpdateAddon struct {
	Id                 string             `json:"id,omitempty"`
	Quantity           *int32             `json:"quantity,omitempty"`
	UnitPrice          *int64             `json:"unit_price,omitempty"`
	BillingCycles      *int32             `json:"billing_cycles,omitempty"`
	QuantityInDecimal  string             `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string             `json:"unit_price_in_decimal,omitempty"`
	TrialEnd           *int64             `json:"trial_end,omitempty"`
	ProrationType      enum.ProrationType `json:"proration_type,omitempty"`
}
type UpdateEventBasedAddon struct {
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
type UpdateCard struct {
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	Number                string                 `json:"number,omitempty"`
	ExpiryMonth           *int32                 `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                 `json:"expiry_year,omitempty"`
	Cvv                   string                 `json:"cvv,omitempty"`
	PreferredScheme       card.PreferredScheme   `json:"preferred_scheme,omitempty"`
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
type UpdatePaymentMethod struct {
	Type                  enum.Type              `json:"type,omitempty"`
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	ReferenceId           string                 `json:"reference_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	IssuingCountry        string                 `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type UpdatePaymentIntent struct {
	Id                    string                          `json:"id,omitempty"`
	GatewayAccountId      string                          `json:"gateway_account_id,omitempty"`
	GwToken               string                          `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntent.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                          `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                          `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}          `json:"additional_information,omitempty"`
}
type UpdateBillingAddress struct {
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
type UpdateShippingAddress struct {
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
type UpdateStatementDescriptor struct {
	Descriptor string `json:"descriptor,omitempty"`
}
type UpdateCustomer struct {
	VatNumber                        string                `json:"vat_number,omitempty"`
	VatNumberPrefix                  string                `json:"vat_number_prefix,omitempty"`
	EntityIdentifierScheme           string                `json:"entity_identifier_scheme,omitempty"`
	IsEinvoiceEnabled                *bool                 `json:"is_einvoice_enabled,omitempty"`
	EinvoicingMethod                 enum.EinvoicingMethod `json:"einvoicing_method,omitempty"`
	EntityIdentifierStandard         string                `json:"entity_identifier_standard,omitempty"`
	BusinessCustomerWithoutVatNumber *bool                 `json:"business_customer_without_vat_number,omitempty"`
	RegisteredForGst                 *bool                 `json:"registered_for_gst,omitempty"`
}
type UpdateContractTerm struct {
	ActionAtTermEnd          subscription.ContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                                   `json:"cancellation_cutoff_period,omitempty"`
}
type UpdateCoupon struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}
type UpdateForItemsRequest struct {
	SubscriptionItems                 []*UpdateForItemsSubscriptionItem  `json:"subscription_items,omitempty"`
	MandatoryItemsToRemove            []string                           `json:"mandatory_items_to_remove,omitempty"`
	ReplaceItemsList                  *bool                              `json:"replace_items_list,omitempty"`
	SetupFee                          *int64                             `json:"setup_fee,omitempty"`
	Discounts                         []*UpdateForItemsDiscount          `json:"discounts,omitempty"`
	ItemTiers                         []*UpdateForItemsItemTier          `json:"item_tiers,omitempty"`
	NetTermDays                       *int32                             `json:"net_term_days,omitempty"`
	InvoiceDate                       *int64                             `json:"invoice_date,omitempty"`
	StartDate                         *int64                             `json:"start_date,omitempty"`
	TrialEnd                          *int64                             `json:"trial_end,omitempty"`
	BillingCycles                     *int32                             `json:"billing_cycles,omitempty"`
	Coupon                            string                             `json:"coupon,omitempty"`
	TermsToCharge                     *int32                             `json:"terms_to_charge,omitempty"`
	ReactivateFrom                    *int64                             `json:"reactivate_from,omitempty"`
	BillingAlignmentMode              enum.BillingAlignmentMode          `json:"billing_alignment_mode,omitempty"`
	AutoCollection                    enum.AutoCollection                `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod          `json:"offline_payment_method,omitempty"`
	PoNumber                          string                             `json:"po_number,omitempty"`
	CouponIds                         []string                           `json:"coupon_ids,omitempty"`
	ReplaceCouponList                 *bool                              `json:"replace_coupon_list,omitempty"`
	Prorate                           *bool                              `json:"prorate,omitempty"`
	EndOfTerm                         *bool                              `json:"end_of_term,omitempty"`
	ForceTermReset                    *bool                              `json:"force_term_reset,omitempty"`
	Reactivate                        *bool                              `json:"reactivate,omitempty"`
	Card                              *UpdateForItemsCard                `json:"card,omitempty"`
	TokenId                           string                             `json:"token_id,omitempty"`
	PaymentMethod                     *UpdateForItemsPaymentMethod       `json:"payment_method,omitempty"`
	PaymentIntent                     *UpdateForItemsPaymentIntent       `json:"payment_intent,omitempty"`
	BillingAddress                    *UpdateForItemsBillingAddress      `json:"billing_address,omitempty"`
	ShippingAddress                   *UpdateForItemsShippingAddress     `json:"shipping_address,omitempty"`
	StatementDescriptor               *UpdateForItemsStatementDescriptor `json:"statement_descriptor,omitempty"`
	Customer                          *UpdateForItemsCustomer            `json:"customer,omitempty"`
	InvoiceNotes                      string                             `json:"invoice_notes,omitempty"`
	MetaData                          map[string]interface{}             `json:"meta_data,omitempty"`
	InvoiceImmediately                *bool                              `json:"invoice_immediately,omitempty"`
	OverrideRelationship              *bool                              `json:"override_relationship,omitempty"`
	ChangesScheduledAt                *int64                             `json:"changes_scheduled_at,omitempty"`
	ChangeOption                      enum.ChangeOption                  `json:"change_option,omitempty"`
	ContractTerm                      *UpdateForItemsContractTerm        `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                             `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	FreePeriod                        *int32                             `json:"free_period,omitempty"`
	FreePeriodUnit                    enum.FreePeriodUnit                `json:"free_period_unit,omitempty"`
	CreatePendingInvoices             *bool                              `json:"create_pending_invoices,omitempty"`
	AutoCloseInvoices                 *bool                              `json:"auto_close_invoices,omitempty"`
	TrialEndAction                    enum.TrialEndAction                `json:"trial_end_action,omitempty"`
	PaymentInitiator                  enum.PaymentInitiator              `json:"payment_initiator,omitempty"`
	Coupons                           []*UpdateForItemsCoupon            `json:"coupons,omitempty"`
	InvoiceUsages                     *bool                              `json:"invoice_usages,omitempty"`
	BillingOverride                   *UpdateForItemsBillingOverride     `json:"billing_override,omitempty"`
}
type UpdateForItemsSubscriptionItem struct {
	ItemPriceId                     string                               `json:"item_price_id"`
	Quantity                        *int32                               `json:"quantity,omitempty"`
	QuantityInDecimal               string                               `json:"quantity_in_decimal,omitempty"`
	UnitPrice                       *int64                               `json:"unit_price,omitempty"`
	UnitPriceInDecimal              string                               `json:"unit_price_in_decimal,omitempty"`
	BillingCycles                   *int32                               `json:"billing_cycles,omitempty"`
	TrialEnd                        *int64                               `json:"trial_end,omitempty"`
	ServicePeriodDays               *int32                               `json:"service_period_days,omitempty"`
	ChargeOnEvent                   enum.ChargeOnEvent                   `json:"charge_on_event,omitempty"`
	ChargeOnce                      *bool                                `json:"charge_once,omitempty"`
	ChargeOnOption                  enum.ChargeOnOption                  `json:"charge_on_option,omitempty"`
	ItemType                        enum.ItemType                        `json:"item_type,omitempty"`
	ProrationType                   enum.ProrationType                   `json:"proration_type,omitempty"`
	UsageAccumulationResetFrequency enum.UsageAccumulationResetFrequency `json:"usage_accumulation_reset_frequency,omitempty"`
}
type UpdateForItemsDiscount struct {
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
type UpdateForItemsItemTier struct {
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
type UpdateForItemsCard struct {
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	Number                string                 `json:"number,omitempty"`
	ExpiryMonth           *int32                 `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                 `json:"expiry_year,omitempty"`
	Cvv                   string                 `json:"cvv,omitempty"`
	PreferredScheme       card.PreferredScheme   `json:"preferred_scheme,omitempty"`
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
type UpdateForItemsPaymentMethod struct {
	Type                  enum.Type              `json:"type,omitempty"`
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	ReferenceId           string                 `json:"reference_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	IssuingCountry        string                 `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type UpdateForItemsPaymentIntent struct {
	Id                    string                          `json:"id,omitempty"`
	GatewayAccountId      string                          `json:"gateway_account_id,omitempty"`
	GwToken               string                          `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntent.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                          `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                          `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}          `json:"additional_information,omitempty"`
}
type UpdateForItemsBillingAddress struct {
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
type UpdateForItemsShippingAddress struct {
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
type UpdateForItemsStatementDescriptor struct {
	Descriptor string `json:"descriptor,omitempty"`
}
type UpdateForItemsCustomer struct {
	VatNumber                        string                `json:"vat_number,omitempty"`
	VatNumberPrefix                  string                `json:"vat_number_prefix,omitempty"`
	EntityIdentifierScheme           string                `json:"entity_identifier_scheme,omitempty"`
	IsEinvoiceEnabled                *bool                 `json:"is_einvoice_enabled,omitempty"`
	EinvoicingMethod                 enum.EinvoicingMethod `json:"einvoicing_method,omitempty"`
	EntityIdentifierStandard         string                `json:"entity_identifier_standard,omitempty"`
	BusinessCustomerWithoutVatNumber *bool                 `json:"business_customer_without_vat_number,omitempty"`
	RegisteredForGst                 *bool                 `json:"registered_for_gst,omitempty"`
}
type UpdateForItemsContractTerm struct {
	ActionAtTermEnd          subscription.ContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                                   `json:"cancellation_cutoff_period,omitempty"`
	ContractStart            *int64                                   `json:"contract_start,omitempty"`
}
type UpdateForItemsCoupon struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}
type UpdateForItemsBillingOverride struct {
	MaxExcessPaymentUsage     *int64 `json:"max_excess_payment_usage,omitempty"`
	MaxRefundableCreditsUsage *int64 `json:"max_refundable_credits_usage,omitempty"`
}
type ChangeTermEndRequest struct {
	TermEndsAt         *int64 `json:"term_ends_at"`
	Prorate            *bool  `json:"prorate,omitempty"`
	InvoiceImmediately *bool  `json:"invoice_immediately,omitempty"`
}
type ReactivateRequest struct {
	TrialEnd                          *int64                         `json:"trial_end,omitempty"`
	BillingCycles                     *int32                         `json:"billing_cycles,omitempty"`
	TrialPeriodDays                   *int32                         `json:"trial_period_days,omitempty"`
	ReactivateFrom                    *int64                         `json:"reactivate_from,omitempty"`
	InvoiceImmediately                *bool                          `json:"invoice_immediately,omitempty"`
	BillingAlignmentMode              enum.BillingAlignmentMode      `json:"billing_alignment_mode,omitempty"`
	TermsToCharge                     *int32                         `json:"terms_to_charge,omitempty"`
	InvoiceDate                       *int64                         `json:"invoice_date,omitempty"`
	ContractTerm                      *ReactivateContractTerm        `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                         `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	PaymentInitiator                  enum.PaymentInitiator          `json:"payment_initiator,omitempty"`
	StatementDescriptor               *ReactivateStatementDescriptor `json:"statement_descriptor,omitempty"`
	PaymentIntent                     *ReactivatePaymentIntent       `json:"payment_intent,omitempty"`
}
type ReactivateContractTerm struct {
	ActionAtTermEnd          subscription.ContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                                   `json:"cancellation_cutoff_period,omitempty"`
}
type ReactivateStatementDescriptor struct {
	Descriptor string `json:"descriptor,omitempty"`
}
type ReactivatePaymentIntent struct {
	Id                    string                          `json:"id,omitempty"`
	GatewayAccountId      string                          `json:"gateway_account_id,omitempty"`
	GwToken               string                          `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntent.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                          `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                          `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}          `json:"additional_information,omitempty"`
}
type AddChargeAtTermEndRequest struct {
	Amount                 *int64               `json:"amount,omitempty"`
	Description            string               `json:"description"`
	AmountInDecimal        string               `json:"amount_in_decimal,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	DateFrom               *int64               `json:"date_from,omitempty"`
	DateTo                 *int64               `json:"date_to,omitempty"`
}
type ChargeAddonAtTermEndRequest struct {
	AddonId                 string `json:"addon_id"`
	AddonQuantity           *int32 `json:"addon_quantity,omitempty"`
	AddonUnitPrice          *int64 `json:"addon_unit_price,omitempty"`
	AddonQuantityInDecimal  string `json:"addon_quantity_in_decimal,omitempty"`
	AddonUnitPriceInDecimal string `json:"addon_unit_price_in_decimal,omitempty"`
	DateFrom                *int64 `json:"date_from,omitempty"`
	DateTo                  *int64 `json:"date_to,omitempty"`
}
type ChargeFutureRenewalsRequest struct {
	TermsToCharge         *int32                                       `json:"terms_to_charge,omitempty"`
	SpecificDatesSchedule []*ChargeFutureRenewalsSpecificDatesSchedule `json:"specific_dates_schedule,omitempty"`
	FixedIntervalSchedule *ChargeFutureRenewalsFixedIntervalSchedule   `json:"fixed_interval_schedule,omitempty"`
	InvoiceImmediately    *bool                                        `json:"invoice_immediately,omitempty"`
	ScheduleType          enum.ScheduleType                            `json:"schedule_type,omitempty"`
}
type ChargeFutureRenewalsSpecificDatesSchedule struct {
	TermsToCharge *int32 `json:"terms_to_charge,omitempty"`
	Date          *int64 `json:"date,omitempty"`
}
type ChargeFutureRenewalsFixedIntervalSchedule struct {
	NumberOfOccurrences *int32             `json:"number_of_occurrences,omitempty"`
	DaysBeforeRenewal   *int32             `json:"days_before_renewal,omitempty"`
	EndScheduleOn       enum.EndScheduleOn `json:"end_schedule_on,omitempty"`
	EndDate             *int64             `json:"end_date,omitempty"`
}
type EditAdvanceInvoiceScheduleRequest struct {
	TermsToCharge         *int32                                             `json:"terms_to_charge,omitempty"`
	ScheduleType          enum.ScheduleType                                  `json:"schedule_type,omitempty"`
	SpecificDatesSchedule []*EditAdvanceInvoiceScheduleSpecificDatesSchedule `json:"specific_dates_schedule,omitempty"`
	FixedIntervalSchedule *EditAdvanceInvoiceScheduleFixedIntervalSchedule   `json:"fixed_interval_schedule,omitempty"`
}
type EditAdvanceInvoiceScheduleSpecificDatesSchedule struct {
	Id            string `json:"id,omitempty"`
	TermsToCharge *int32 `json:"terms_to_charge,omitempty"`
	Date          *int64 `json:"date,omitempty"`
}
type EditAdvanceInvoiceScheduleFixedIntervalSchedule struct {
	NumberOfOccurrences *int32             `json:"number_of_occurrences,omitempty"`
	DaysBeforeRenewal   *int32             `json:"days_before_renewal,omitempty"`
	EndScheduleOn       enum.EndScheduleOn `json:"end_schedule_on,omitempty"`
	EndDate             *int64             `json:"end_date,omitempty"`
}
type RemoveAdvanceInvoiceScheduleRequest struct {
	SpecificDatesSchedule []*RemoveAdvanceInvoiceScheduleSpecificDatesSchedule `json:"specific_dates_schedule,omitempty"`
}
type RemoveAdvanceInvoiceScheduleSpecificDatesSchedule struct {
	Id string `json:"id,omitempty"`
}
type RegenerateInvoiceRequest struct {
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
	Prorate            *bool  `json:"prorate,omitempty"`
	InvoiceImmediately *bool  `json:"invoice_immediately,omitempty"`
}
type ImportSubscriptionRequest struct {
	Id                                string                                      `json:"id,omitempty"`
	Customer                          *ImportSubscriptionCustomer                 `json:"customer,omitempty"`
	ClientProfileId                   string                                      `json:"client_profile_id,omitempty"`
	PlanId                            string                                      `json:"plan_id"`
	PlanQuantity                      *int32                                      `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                                      `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int64                                      `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                                      `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int64                                      `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                                      `json:"trial_end,omitempty"`
	BillingCycles                     *int32                                      `json:"billing_cycles,omitempty"`
	Addons                            []*ImportSubscriptionAddon                  `json:"addons,omitempty"`
	EventBasedAddons                  []*ImportSubscriptionEventBasedAddon        `json:"event_based_addons,omitempty"`
	ChargedEventBasedAddons           []*ImportSubscriptionChargedEventBasedAddon `json:"charged_event_based_addons,omitempty"`
	StartDate                         *int64                                      `json:"start_date,omitempty"`
	AutoCollection                    enum.AutoCollection                         `json:"auto_collection,omitempty"`
	PoNumber                          string                                      `json:"po_number,omitempty"`
	CouponIds                         []string                                    `json:"coupon_ids,omitempty"`
	ContractTerm                      *ImportSubscriptionContractTerm             `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                      `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	Status                            Status                                      `json:"status"`
	CurrentTermEnd                    *int64                                      `json:"current_term_end,omitempty"`
	CurrentTermStart                  *int64                                      `json:"current_term_start,omitempty"`
	TrialStart                        *int64                                      `json:"trial_start,omitempty"`
	CancelledAt                       *int64                                      `json:"cancelled_at,omitempty"`
	StartedAt                         *int64                                      `json:"started_at,omitempty"`
	ActivatedAt                       *int64                                      `json:"activated_at,omitempty"`
	PauseDate                         *int64                                      `json:"pause_date,omitempty"`
	ResumeDate                        *int64                                      `json:"resume_date,omitempty"`
	CreateCurrentTermInvoice          *bool                                       `json:"create_current_term_invoice,omitempty"`
	Card                              *ImportSubscriptionCard                     `json:"card,omitempty"`
	PaymentMethod                     *ImportSubscriptionPaymentMethod            `json:"payment_method,omitempty"`
	BillingAddress                    *ImportSubscriptionBillingAddress           `json:"billing_address,omitempty"`
	ShippingAddress                   *ImportSubscriptionShippingAddress          `json:"shipping_address,omitempty"`
	AffiliateToken                    string                                      `json:"affiliate_token,omitempty"`
	InvoiceNotes                      string                                      `json:"invoice_notes,omitempty"`
	MetaData                          map[string]interface{}                      `json:"meta_data,omitempty"`
	Transaction                       *ImportSubscriptionTransaction              `json:"transaction,omitempty"`
	Coupons                           []*ImportSubscriptionCoupon                 `json:"coupons,omitempty"`
}
type ImportSubscriptionCustomer struct {
	Id                      string                       `json:"id,omitempty"`
	Email                   string                       `json:"email,omitempty"`
	FirstName               string                       `json:"first_name,omitempty"`
	LastName                string                       `json:"last_name,omitempty"`
	Company                 string                       `json:"company,omitempty"`
	Phone                   string                       `json:"phone,omitempty"`
	Locale                  string                       `json:"locale,omitempty"`
	Taxability              enum.Taxability              `json:"taxability,omitempty"`
	EntityCode              enum.EntityCode              `json:"entity_code,omitempty"`
	ExemptNumber            string                       `json:"exempt_number,omitempty"`
	NetTermDays             *int32                       `json:"net_term_days,omitempty"`
	TaxjarExemptionCategory enum.TaxjarExemptionCategory `json:"taxjar_exemption_category,omitempty"`
	CustomerType            enum.CustomerType            `json:"customer_type,omitempty"`
	AutoCollection          enum.AutoCollection          `json:"auto_collection,omitempty"`
	AllowDirectDebit        *bool                        `json:"allow_direct_debit,omitempty"`
	VatNumber               string                       `json:"vat_number,omitempty"`
	VatNumberPrefix         string                       `json:"vat_number_prefix,omitempty"`
}
type ImportSubscriptionAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
}
type ImportSubscriptionEventBasedAddon struct {
	Id                  string       `json:"id,omitempty"`
	Quantity            *int32       `json:"quantity,omitempty"`
	UnitPrice           *int64       `json:"unit_price,omitempty"`
	QuantityInDecimal   string       `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string       `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodInDays *int32       `json:"service_period_in_days,omitempty"`
	OnEvent             enum.OnEvent `json:"on_event,omitempty"`
	ChargeOnce          *bool        `json:"charge_once,omitempty"`
}
type ImportSubscriptionChargedEventBasedAddon struct {
	Id            string `json:"id,omitempty"`
	LastChargedAt *int64 `json:"last_charged_at,omitempty"`
}
type ImportSubscriptionContractTerm struct {
	Id                         string                                   `json:"id,omitempty"`
	CreatedAt                  *int64                                   `json:"created_at,omitempty"`
	ContractStart              *int64                                   `json:"contract_start,omitempty"`
	BillingCycle               *int32                                   `json:"billing_cycle,omitempty"`
	TotalAmountRaised          *int64                                   `json:"total_amount_raised,omitempty"`
	TotalAmountRaisedBeforeTax *int64                                   `json:"total_amount_raised_before_tax,omitempty"`
	ActionAtTermEnd            subscription.ContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod   *int32                                   `json:"cancellation_cutoff_period,omitempty"`
}
type ImportSubscriptionCard struct {
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	Number                string                 `json:"number,omitempty"`
	ExpiryMonth           *int32                 `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                 `json:"expiry_year,omitempty"`
	Cvv                   string                 `json:"cvv,omitempty"`
	PreferredScheme       card.PreferredScheme   `json:"preferred_scheme,omitempty"`
	BillingAddr1          string                 `json:"billing_addr1,omitempty"`
	BillingAddr2          string                 `json:"billing_addr2,omitempty"`
	BillingCity           string                 `json:"billing_city,omitempty"`
	BillingStateCode      string                 `json:"billing_state_code,omitempty"`
	BillingState          string                 `json:"billing_state,omitempty"`
	BillingZip            string                 `json:"billing_zip,omitempty"`
	BillingCountry        string                 `json:"billing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type ImportSubscriptionPaymentMethod struct {
	Type                  enum.Type              `json:"type,omitempty"`
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	ReferenceId           string                 `json:"reference_id,omitempty"`
	IssuingCountry        string                 `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type ImportSubscriptionBillingAddress struct {
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
type ImportSubscriptionShippingAddress struct {
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
type ImportSubscriptionTransaction struct {
	Amount          *int64             `json:"amount,omitempty"`
	PaymentMethod   enum.PaymentMethod `json:"payment_method,omitempty"`
	ReferenceNumber string             `json:"reference_number,omitempty"`
	Date            *int64             `json:"date,omitempty"`
}
type ImportSubscriptionCoupon struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}
type ImportForCustomerRequest struct {
	Id                                string                                     `json:"id,omitempty"`
	PlanId                            string                                     `json:"plan_id"`
	PlanQuantity                      *int32                                     `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                                     `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int64                                     `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                                     `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int64                                     `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                                     `json:"trial_end,omitempty"`
	BillingCycles                     *int32                                     `json:"billing_cycles,omitempty"`
	Addons                            []*ImportForCustomerAddon                  `json:"addons,omitempty"`
	EventBasedAddons                  []*ImportForCustomerEventBasedAddon        `json:"event_based_addons,omitempty"`
	ChargedEventBasedAddons           []*ImportForCustomerChargedEventBasedAddon `json:"charged_event_based_addons,omitempty"`
	StartDate                         *int64                                     `json:"start_date,omitempty"`
	AutoCollection                    enum.AutoCollection                        `json:"auto_collection,omitempty"`
	PoNumber                          string                                     `json:"po_number,omitempty"`
	CouponIds                         []string                                   `json:"coupon_ids,omitempty"`
	PaymentSourceId                   string                                     `json:"payment_source_id,omitempty"`
	Status                            Status                                     `json:"status"`
	CurrentTermEnd                    *int64                                     `json:"current_term_end,omitempty"`
	CurrentTermStart                  *int64                                     `json:"current_term_start,omitempty"`
	TrialStart                        *int64                                     `json:"trial_start,omitempty"`
	CancelledAt                       *int64                                     `json:"cancelled_at,omitempty"`
	StartedAt                         *int64                                     `json:"started_at,omitempty"`
	ActivatedAt                       *int64                                     `json:"activated_at,omitempty"`
	PauseDate                         *int64                                     `json:"pause_date,omitempty"`
	ResumeDate                        *int64                                     `json:"resume_date,omitempty"`
	ContractTerm                      *ImportForCustomerContractTerm             `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                                     `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	CreateCurrentTermInvoice          *bool                                      `json:"create_current_term_invoice,omitempty"`
	Transaction                       *ImportForCustomerTransaction              `json:"transaction,omitempty"`
	ShippingAddress                   *ImportForCustomerShippingAddress          `json:"shipping_address,omitempty"`
	InvoiceNotes                      string                                     `json:"invoice_notes,omitempty"`
	MetaData                          map[string]interface{}                     `json:"meta_data,omitempty"`
	Coupons                           []*ImportForCustomerCoupon                 `json:"coupons,omitempty"`
}
type ImportForCustomerAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
}
type ImportForCustomerEventBasedAddon struct {
	Id                  string       `json:"id,omitempty"`
	Quantity            *int32       `json:"quantity,omitempty"`
	UnitPrice           *int64       `json:"unit_price,omitempty"`
	QuantityInDecimal   string       `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string       `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodInDays *int32       `json:"service_period_in_days,omitempty"`
	OnEvent             enum.OnEvent `json:"on_event,omitempty"`
	ChargeOnce          *bool        `json:"charge_once,omitempty"`
}
type ImportForCustomerChargedEventBasedAddon struct {
	Id            string `json:"id,omitempty"`
	LastChargedAt *int64 `json:"last_charged_at,omitempty"`
}
type ImportForCustomerContractTerm struct {
	Id                         string                                   `json:"id,omitempty"`
	CreatedAt                  *int64                                   `json:"created_at,omitempty"`
	ContractStart              *int64                                   `json:"contract_start,omitempty"`
	BillingCycle               *int32                                   `json:"billing_cycle,omitempty"`
	TotalAmountRaised          *int64                                   `json:"total_amount_raised,omitempty"`
	TotalAmountRaisedBeforeTax *int64                                   `json:"total_amount_raised_before_tax,omitempty"`
	ActionAtTermEnd            subscription.ContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod   *int32                                   `json:"cancellation_cutoff_period,omitempty"`
}
type ImportForCustomerTransaction struct {
	Amount          *int64             `json:"amount,omitempty"`
	PaymentMethod   enum.PaymentMethod `json:"payment_method,omitempty"`
	ReferenceNumber string             `json:"reference_number,omitempty"`
	Date            *int64             `json:"date,omitempty"`
}
type ImportForCustomerShippingAddress struct {
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
type ImportForCustomerCoupon struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}
type ImportContractTermRequest struct {
	ContractTerm                      *ImportContractTermContractTerm `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                          `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type ImportContractTermContractTerm struct {
	Id                          string                                   `json:"id,omitempty"`
	CreatedAt                   *int64                                   `json:"created_at,omitempty"`
	ContractStart               *int64                                   `json:"contract_start,omitempty"`
	ContractEnd                 *int64                                   `json:"contract_end,omitempty"`
	Status                      subscription.ContractTermStatus          `json:"status,omitempty"`
	TotalAmountRaised           *int64                                   `json:"total_amount_raised,omitempty"`
	TotalAmountRaisedBeforeTax  *int64                                   `json:"total_amount_raised_before_tax,omitempty"`
	TotalContractValue          *int64                                   `json:"total_contract_value,omitempty"`
	TotalContractValueBeforeTax *int64                                   `json:"total_contract_value_before_tax,omitempty"`
	BillingCycle                *int32                                   `json:"billing_cycle,omitempty"`
	ActionAtTermEnd             subscription.ContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod    *int32                                   `json:"cancellation_cutoff_period,omitempty"`
}
type ImportUnbilledChargesRequest struct {
	UnbilledCharges []*ImportUnbilledChargesUnbilledCharge `json:"unbilled_charges,omitempty"`
	Discounts       []*ImportUnbilledChargesDiscount       `json:"discounts,omitempty"`
	Tiers           []*ImportUnbilledChargesTier           `json:"tiers,omitempty"`
}
type ImportUnbilledChargesUnbilledCharge struct {
	Id                  string                    `json:"id,omitempty"`
	DateFrom            *int64                    `json:"date_from"`
	DateTo              *int64                    `json:"date_to"`
	EntityType          unbilledCharge.EntityType `json:"entity_type"`
	EntityId            string                    `json:"entity_id,omitempty"`
	Description         string                    `json:"description,omitempty"`
	UnitAmount          *int64                    `json:"unit_amount,omitempty"`
	Quantity            *int32                    `json:"quantity,omitempty"`
	Amount              *int64                    `json:"amount,omitempty"`
	UnitAmountInDecimal string                    `json:"unit_amount_in_decimal,omitempty"`
	QuantityInDecimal   string                    `json:"quantity_in_decimal,omitempty"`
	AmountInDecimal     string                    `json:"amount_in_decimal,omitempty"`
	DiscountAmount      *int64                    `json:"discount_amount,omitempty"`
	UseForProration     *bool                     `json:"use_for_proration,omitempty"`
	IsAdvanceCharge     *bool                     `json:"is_advance_charge,omitempty"`
}
type ImportUnbilledChargesDiscount struct {
	UnbilledChargeId string              `json:"unbilled_charge_id,omitempty"`
	EntityType       discount.EntityType `json:"entity_type,omitempty"`
	EntityId         string              `json:"entity_id,omitempty"`
	Description      string              `json:"description,omitempty"`
	Amount           *int64              `json:"amount"`
}
type ImportUnbilledChargesTier struct {
	UnbilledChargeId      string `json:"unbilled_charge_id"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	QuantityUsed          *int32 `json:"quantity_used,omitempty"`
	UnitAmount            *int64 `json:"unit_amount,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	QuantityUsedInDecimal string `json:"quantity_used_in_decimal,omitempty"`
	UnitAmountInDecimal   string `json:"unit_amount_in_decimal,omitempty"`
}
type ImportForItemsRequest struct {
	ExhaustedCouponIds                []string                          `json:"exhausted_coupon_ids,omitempty"`
	Id                                string                            `json:"id,omitempty"`
	TrialEnd                          *int64                            `json:"trial_end,omitempty"`
	BillingCycles                     *int32                            `json:"billing_cycles,omitempty"`
	SubscriptionItems                 []*ImportForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	SetupFee                          *int64                            `json:"setup_fee,omitempty"`
	Discounts                         []*ImportForItemsDiscount         `json:"discounts,omitempty"`
	ChargedItems                      []*ImportForItemsChargedItem      `json:"charged_items,omitempty"`
	ItemTiers                         []*ImportForItemsItemTier         `json:"item_tiers,omitempty"`
	NetTermDays                       *int32                            `json:"net_term_days,omitempty"`
	StartDate                         *int64                            `json:"start_date,omitempty"`
	AutoCollection                    enum.AutoCollection               `json:"auto_collection,omitempty"`
	PoNumber                          string                            `json:"po_number,omitempty"`
	CouponIds                         []string                          `json:"coupon_ids,omitempty"`
	PaymentSourceId                   string                            `json:"payment_source_id,omitempty"`
	Status                            Status                            `json:"status"`
	CurrentTermEnd                    *int64                            `json:"current_term_end,omitempty"`
	CurrentTermStart                  *int64                            `json:"current_term_start,omitempty"`
	TrialStart                        *int64                            `json:"trial_start,omitempty"`
	CancelledAt                       *int64                            `json:"cancelled_at,omitempty"`
	StartedAt                         *int64                            `json:"started_at,omitempty"`
	ActivatedAt                       *int64                            `json:"activated_at,omitempty"`
	PauseDate                         *int64                            `json:"pause_date,omitempty"`
	ResumeDate                        *int64                            `json:"resume_date,omitempty"`
	ContractTerm                      *ImportForItemsContractTerm       `json:"contract_term,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                            `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	CreateCurrentTermInvoice          *bool                             `json:"create_current_term_invoice,omitempty"`
	Transaction                       *ImportForItemsTransaction        `json:"transaction,omitempty"`
	ShippingAddress                   *ImportForItemsShippingAddress    `json:"shipping_address,omitempty"`
	InvoiceNotes                      string                            `json:"invoice_notes,omitempty"`
	MetaData                          map[string]interface{}            `json:"meta_data,omitempty"`
	CancelReasonCode                  string                            `json:"cancel_reason_code,omitempty"`
	CreatePendingInvoices             *bool                             `json:"create_pending_invoices,omitempty"`
	AutoCloseInvoices                 *bool                             `json:"auto_close_invoices,omitempty"`
	Coupons                           []*ImportForItemsCoupon           `json:"coupons,omitempty"`
}
type ImportForItemsSubscriptionItem struct {
	ItemPriceId        string             `json:"item_price_id"`
	Quantity           *int32             `json:"quantity,omitempty"`
	QuantityInDecimal  string             `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64             `json:"unit_price,omitempty"`
	UnitPriceInDecimal string             `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32             `json:"billing_cycles,omitempty"`
	TrialEnd           *int64             `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32             `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool              `json:"charge_once,omitempty"`
	ItemType           enum.ItemType      `json:"item_type,omitempty"`
}
type ImportForItemsDiscount struct {
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
type ImportForItemsChargedItem struct {
	ItemPriceId   string `json:"item_price_id,omitempty"`
	LastChargedAt *int64 `json:"last_charged_at,omitempty"`
}
type ImportForItemsItemTier struct {
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
type ImportForItemsContractTerm struct {
	Id                         string                                   `json:"id,omitempty"`
	CreatedAt                  *int64                                   `json:"created_at,omitempty"`
	ContractStart              *int64                                   `json:"contract_start,omitempty"`
	BillingCycle               *int32                                   `json:"billing_cycle,omitempty"`
	TotalAmountRaised          *int64                                   `json:"total_amount_raised,omitempty"`
	TotalAmountRaisedBeforeTax *int64                                   `json:"total_amount_raised_before_tax,omitempty"`
	ActionAtTermEnd            subscription.ContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod   *int32                                   `json:"cancellation_cutoff_period,omitempty"`
}
type ImportForItemsTransaction struct {
	Amount          *int64             `json:"amount,omitempty"`
	PaymentMethod   enum.PaymentMethod `json:"payment_method,omitempty"`
	ReferenceNumber string             `json:"reference_number,omitempty"`
	Date            *int64             `json:"date,omitempty"`
}
type ImportForItemsShippingAddress struct {
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
type ImportForItemsCoupon struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}
type OverrideBillingProfileRequest struct {
	PaymentSourceId string              `json:"payment_source_id,omitempty"`
	AutoCollection  enum.AutoCollection `json:"auto_collection,omitempty"`
}
type PauseRequest struct {
	PauseOption             enum.PauseOption             `json:"pause_option,omitempty"`
	PauseDate               *int64                       `json:"pause_date,omitempty"`
	UnbilledChargesHandling enum.UnbilledChargesHandling `json:"unbilled_charges_handling,omitempty"`
	InvoiceDunningHandling  enum.InvoiceDunningHandling  `json:"invoice_dunning_handling,omitempty"`
	SkipBillingCycles       *int32                       `json:"skip_billing_cycles,omitempty"`
	ResumeDate              *int64                       `json:"resume_date,omitempty"`
}
type CancelRequest struct {
	CancelOption                      enum.CancelOption                      `json:"cancel_option,omitempty"`
	EndOfTerm                         *bool                                  `json:"end_of_term,omitempty"`
	CancelAt                          *int64                                 `json:"cancel_at,omitempty"`
	CreditOptionForCurrentTermCharges enum.CreditOptionForCurrentTermCharges `json:"credit_option_for_current_term_charges,omitempty"`
	UnbilledChargesOption             enum.UnbilledChargesOption             `json:"unbilled_charges_option,omitempty"`
	AccountReceivablesHandling        enum.AccountReceivablesHandling        `json:"account_receivables_handling,omitempty"`
	RefundableCreditsHandling         enum.RefundableCreditsHandling         `json:"refundable_credits_handling,omitempty"`
	ContractTermCancelOption          enum.ContractTermCancelOption          `json:"contract_term_cancel_option,omitempty"`
	InvoiceDate                       *int64                                 `json:"invoice_date,omitempty"`
	EventBasedAddons                  []*CancelEventBasedAddon               `json:"event_based_addons,omitempty"`
	CancelReasonCode                  string                                 `json:"cancel_reason_code,omitempty"`
}
type CancelEventBasedAddon struct {
	Id                  string `json:"id,omitempty"`
	Quantity            *int32 `json:"quantity,omitempty"`
	UnitPrice           *int64 `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32 `json:"service_period_in_days,omitempty"`
}
type CancelForItemsRequest struct {
	CancelOption                      enum.CancelOption                      `json:"cancel_option,omitempty"`
	EndOfTerm                         *bool                                  `json:"end_of_term,omitempty"`
	CancelAt                          *int64                                 `json:"cancel_at,omitempty"`
	CreditOptionForCurrentTermCharges enum.CreditOptionForCurrentTermCharges `json:"credit_option_for_current_term_charges,omitempty"`
	UnbilledChargesOption             enum.UnbilledChargesOption             `json:"unbilled_charges_option,omitempty"`
	AccountReceivablesHandling        enum.AccountReceivablesHandling        `json:"account_receivables_handling,omitempty"`
	RefundableCreditsHandling         enum.RefundableCreditsHandling         `json:"refundable_credits_handling,omitempty"`
	ContractTermCancelOption          enum.ContractTermCancelOption          `json:"contract_term_cancel_option,omitempty"`
	InvoiceDate                       *int64                                 `json:"invoice_date,omitempty"`
	SubscriptionItems                 []*CancelForItemsSubscriptionItem      `json:"subscription_items,omitempty"`
	CancelReasonCode                  string                                 `json:"cancel_reason_code,omitempty"`
}
type CancelForItemsSubscriptionItem struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodDays  *int32 `json:"service_period_days,omitempty"`
}
type ResumeRequest struct {
	ResumeOption           enum.ResumeOption           `json:"resume_option,omitempty"`
	ResumeDate             *int64                      `json:"resume_date,omitempty"`
	ChargesHandling        enum.ChargesHandling        `json:"charges_handling,omitempty"`
	UnpaidInvoicesHandling enum.UnpaidInvoicesHandling `json:"unpaid_invoices_handling,omitempty"`
	PaymentInitiator       enum.PaymentInitiator       `json:"payment_initiator,omitempty"`
	PaymentIntent          *ResumePaymentIntent        `json:"payment_intent,omitempty"`
}
type ResumePaymentIntent struct {
	Id                    string                          `json:"id,omitempty"`
	GatewayAccountId      string                          `json:"gateway_account_id,omitempty"`
	GwToken               string                          `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntent.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                          `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                          `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}          `json:"additional_information,omitempty"`
}
type MoveRequest struct {
	ToCustomerId      string `json:"to_customer_id"`
	CopyPaymentSource *bool  `json:"copy_payment_source,omitempty"`
}

type CreateResponse struct {
	Subscription    *Subscription                    `json:"subscription,omitempty"`
	Customer        *customer.Customer               `json:"customer,omitempty"`
	Card            *card.Card                       `json:"card,omitempty"`
	Invoice         *invoice.Invoice                 `json:"invoice,omitempty"`
	UnbilledCharges []*unbilledcharge.UnbilledCharge `json:"unbilled_charges,omitempty"`
}

type CreateForCustomerResponse struct {
	Subscription    *Subscription                    `json:"subscription,omitempty"`
	Customer        *customer.Customer               `json:"customer,omitempty"`
	Card            *card.Card                       `json:"card,omitempty"`
	Invoice         *invoice.Invoice                 `json:"invoice,omitempty"`
	UnbilledCharges []*unbilledcharge.UnbilledCharge `json:"unbilled_charges,omitempty"`
}

type CreateWithItemsResponse struct {
	Subscription    *Subscription                    `json:"subscription,omitempty"`
	Customer        *customer.Customer               `json:"customer,omitempty"`
	Card            *card.Card                       `json:"card,omitempty"`
	Invoice         *invoice.Invoice                 `json:"invoice,omitempty"`
	UnbilledCharges []*unbilledcharge.UnbilledCharge `json:"unbilled_charges,omitempty"`
}

type ListSubscriptionResponse struct {
	Subscription *Subscription      `json:"subscription,omitempty"`
	Customer     *customer.Customer `json:"customer,omitempty"`
	Card         *card.Card         `json:"card,omitempty"`
}

type ListResponse struct {
	List       []*ListSubscriptionResponse `json:"list,omitempty"`
	NextOffset string                      `json:"next_offset,omitempty"`
}

type SubscriptionsForCustomerSubscriptionResponse struct {
	Subscription *Subscription `json:"subscription,omitempty"`
}

type SubscriptionsForCustomerResponse struct {
	List       []*SubscriptionsForCustomerSubscriptionResponse `json:"list,omitempty"`
	NextOffset string                                          `json:"next_offset,omitempty"`
}

type ContractTermsForSubscriptionSubscriptionResponse struct {
	ContractTerm *contractterm.ContractTerm `json:"contract_term,omitempty"`
}

type ContractTermsForSubscriptionResponse struct {
	List       []*ContractTermsForSubscriptionSubscriptionResponse `json:"list,omitempty"`
	NextOffset string                                              `json:"next_offset,omitempty"`
}

type ListDiscountsSubscriptionResponse struct {
	Discount *discount.Discount `json:"discount,omitempty"`
}

type ListDiscountsResponse struct {
	List       []*ListDiscountsSubscriptionResponse `json:"list,omitempty"`
	NextOffset string                               `json:"next_offset,omitempty"`
}

type RetrieveResponse struct {
	Subscription *Subscription      `json:"subscription,omitempty"`
	Customer     *customer.Customer `json:"customer,omitempty"`
	Card         *card.Card         `json:"card,omitempty"`
}

type RetrieveWithScheduledChangesResponse struct {
	Subscription *Subscription      `json:"subscription,omitempty"`
	Customer     *customer.Customer `json:"customer,omitempty"`
	Card         *card.Card         `json:"card,omitempty"`
}

type RemoveScheduledChangesResponse struct {
	Subscription *Subscription            `json:"subscription,omitempty"`
	Customer     *customer.Customer       `json:"customer,omitempty"`
	Card         *card.Card               `json:"card,omitempty"`
	CreditNotes  []*creditnote.CreditNote `json:"credit_notes,omitempty"`
}

type RemoveScheduledCancellationResponse struct {
	Subscription *Subscription      `json:"subscription,omitempty"`
	Customer     *customer.Customer `json:"customer,omitempty"`
	Card         *card.Card         `json:"card,omitempty"`
}

type RemoveCouponsResponse struct {
	Subscription *Subscription      `json:"subscription,omitempty"`
	Customer     *customer.Customer `json:"customer,omitempty"`
	Card         *card.Card         `json:"card,omitempty"`
}

type UpdateResponse struct {
	Subscription    *Subscription                    `json:"subscription,omitempty"`
	Customer        *customer.Customer               `json:"customer,omitempty"`
	Card            *card.Card                       `json:"card,omitempty"`
	Invoice         *invoice.Invoice                 `json:"invoice,omitempty"`
	UnbilledCharges []*unbilledcharge.UnbilledCharge `json:"unbilled_charges,omitempty"`
	CreditNotes     []*creditnote.CreditNote         `json:"credit_notes,omitempty"`
}

type UpdateForItemsResponse struct {
	Subscription    *Subscription                    `json:"subscription,omitempty"`
	Customer        *customer.Customer               `json:"customer,omitempty"`
	Card            *card.Card                       `json:"card,omitempty"`
	Invoice         *invoice.Invoice                 `json:"invoice,omitempty"`
	UnbilledCharges []*unbilledcharge.UnbilledCharge `json:"unbilled_charges,omitempty"`
	CreditNotes     []*creditnote.CreditNote         `json:"credit_notes,omitempty"`
}

type ChangeTermEndResponse struct {
	Subscription    *Subscription                    `json:"subscription,omitempty"`
	Customer        *customer.Customer               `json:"customer,omitempty"`
	Card            *card.Card                       `json:"card,omitempty"`
	Invoice         *invoice.Invoice                 `json:"invoice,omitempty"`
	UnbilledCharges []*unbilledcharge.UnbilledCharge `json:"unbilled_charges,omitempty"`
	CreditNotes     []*creditnote.CreditNote         `json:"credit_notes,omitempty"`
}

type ReactivateResponse struct {
	Subscription    *Subscription                    `json:"subscription,omitempty"`
	Customer        *customer.Customer               `json:"customer,omitempty"`
	Card            *card.Card                       `json:"card,omitempty"`
	Invoice         *invoice.Invoice                 `json:"invoice,omitempty"`
	UnbilledCharges []*unbilledcharge.UnbilledCharge `json:"unbilled_charges,omitempty"`
}

type AddChargeAtTermEndResponse struct {
	Estimate *estimate.Estimate `json:"estimate,omitempty"`
}

type ChargeAddonAtTermEndResponse struct {
	Estimate *estimate.Estimate `json:"estimate,omitempty"`
}

type ChargeFutureRenewalsResponse struct {
	Subscription            *Subscription                                    `json:"subscription,omitempty"`
	Customer                *customer.Customer                               `json:"customer,omitempty"`
	Card                    *card.Card                                       `json:"card,omitempty"`
	Invoice                 *invoice.Invoice                                 `json:"invoice,omitempty"`
	AdvanceInvoiceSchedules []*advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedules,omitempty"`
}

type EditAdvanceInvoiceScheduleResponse struct {
	AdvanceInvoiceSchedules []*advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedules,omitempty"`
}

type RetrieveAdvanceInvoiceScheduleResponse struct {
	AdvanceInvoiceSchedules []*advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedules,omitempty"`
}

type RemoveAdvanceInvoiceScheduleResponse struct {
	Subscription            *Subscription                                    `json:"subscription,omitempty"`
	AdvanceInvoiceSchedules []*advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedules,omitempty"`
}

type RegenerateInvoiceResponse struct {
	Invoice         *invoice.Invoice                 `json:"invoice,omitempty"`
	UnbilledCharges []*unbilledcharge.UnbilledCharge `json:"unbilled_charges,omitempty"`
}

type ImportSubscriptionResponse struct {
	Subscription *Subscription      `json:"subscription,omitempty"`
	Customer     *customer.Customer `json:"customer,omitempty"`
	Card         *card.Card         `json:"card,omitempty"`
	Invoice      *invoice.Invoice   `json:"invoice,omitempty"`
}

type ImportForCustomerResponse struct {
	Subscription *Subscription      `json:"subscription,omitempty"`
	Customer     *customer.Customer `json:"customer,omitempty"`
	Card         *card.Card         `json:"card,omitempty"`
	Invoice      *invoice.Invoice   `json:"invoice,omitempty"`
}

type ImportContractTermResponse struct {
	ContractTerm *contractterm.ContractTerm `json:"contract_term,omitempty"`
}

type ImportUnbilledChargesResponse struct {
	UnbilledCharges []*unbilledcharge.UnbilledCharge `json:"unbilled_charges,omitempty"`
}

type ImportForItemsResponse struct {
	Subscription *Subscription      `json:"subscription,omitempty"`
	Customer     *customer.Customer `json:"customer,omitempty"`
	Card         *card.Card         `json:"card,omitempty"`
	Invoice      *invoice.Invoice   `json:"invoice,omitempty"`
}

type OverrideBillingProfileResponse struct {
	Subscription  *Subscription                `json:"subscription,omitempty"`
	PaymentSource *paymentsource.PaymentSource `json:"payment_source,omitempty"`
}

type DeleteResponse struct {
	Subscription *Subscription      `json:"subscription,omitempty"`
	Customer     *customer.Customer `json:"customer,omitempty"`
	Card         *card.Card         `json:"card,omitempty"`
}

type PauseResponse struct {
	Subscription    *Subscription                    `json:"subscription,omitempty"`
	Customer        *customer.Customer               `json:"customer,omitempty"`
	Card            *card.Card                       `json:"card,omitempty"`
	Invoice         *invoice.Invoice                 `json:"invoice,omitempty"`
	UnbilledCharges []*unbilledcharge.UnbilledCharge `json:"unbilled_charges,omitempty"`
	CreditNotes     []*creditnote.CreditNote         `json:"credit_notes,omitempty"`
}

type CancelResponse struct {
	Subscription    *Subscription                    `json:"subscription,omitempty"`
	Customer        *customer.Customer               `json:"customer,omitempty"`
	Card            *card.Card                       `json:"card,omitempty"`
	Invoice         *invoice.Invoice                 `json:"invoice,omitempty"`
	UnbilledCharges []*unbilledcharge.UnbilledCharge `json:"unbilled_charges,omitempty"`
	CreditNotes     []*creditnote.CreditNote         `json:"credit_notes,omitempty"`
}

type CancelForItemsResponse struct {
	Subscription    *Subscription                    `json:"subscription,omitempty"`
	Customer        *customer.Customer               `json:"customer,omitempty"`
	Card            *card.Card                       `json:"card,omitempty"`
	Invoice         *invoice.Invoice                 `json:"invoice,omitempty"`
	UnbilledCharges []*unbilledcharge.UnbilledCharge `json:"unbilled_charges,omitempty"`
	CreditNotes     []*creditnote.CreditNote         `json:"credit_notes,omitempty"`
}

type ResumeResponse struct {
	Subscription    *Subscription                    `json:"subscription,omitempty"`
	Customer        *customer.Customer               `json:"customer,omitempty"`
	Card            *card.Card                       `json:"card,omitempty"`
	Invoice         *invoice.Invoice                 `json:"invoice,omitempty"`
	UnbilledCharges []*unbilledcharge.UnbilledCharge `json:"unbilled_charges,omitempty"`
}

type RemoveScheduledPauseResponse struct {
	Subscription *Subscription      `json:"subscription,omitempty"`
	Customer     *customer.Customer `json:"customer,omitempty"`
	Card         *card.Card         `json:"card,omitempty"`
}

type RemoveScheduledResumptionResponse struct {
	Subscription *Subscription      `json:"subscription,omitempty"`
	Customer     *customer.Customer `json:"customer,omitempty"`
	Card         *card.Card         `json:"card,omitempty"`
}

type MoveResponse struct {
	Subscription *Subscription `json:"subscription,omitempty"`
}
