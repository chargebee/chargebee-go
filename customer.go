package chargebee

import (
	"encoding/json"
)

type CustomerAutoCollection string

const (
	CustomerAutoCollectionOn  CustomerAutoCollection = "on"
	CustomerAutoCollectionOff CustomerAutoCollection = "off"
)

type CustomerOfflinePaymentMethod string

const (
	CustomerOfflinePaymentMethodNoPreference            CustomerOfflinePaymentMethod = "no_preference"
	CustomerOfflinePaymentMethodCash                    CustomerOfflinePaymentMethod = "cash"
	CustomerOfflinePaymentMethodCheck                   CustomerOfflinePaymentMethod = "check"
	CustomerOfflinePaymentMethodBankTransfer            CustomerOfflinePaymentMethod = "bank_transfer"
	CustomerOfflinePaymentMethodAchCredit               CustomerOfflinePaymentMethod = "ach_credit"
	CustomerOfflinePaymentMethodSepaCredit              CustomerOfflinePaymentMethod = "sepa_credit"
	CustomerOfflinePaymentMethodBoleto                  CustomerOfflinePaymentMethod = "boleto"
	CustomerOfflinePaymentMethodUsAutomatedBankTransfer CustomerOfflinePaymentMethod = "us_automated_bank_transfer"
	CustomerOfflinePaymentMethodEuAutomatedBankTransfer CustomerOfflinePaymentMethod = "eu_automated_bank_transfer"
	CustomerOfflinePaymentMethodUkAutomatedBankTransfer CustomerOfflinePaymentMethod = "uk_automated_bank_transfer"
	CustomerOfflinePaymentMethodJpAutomatedBankTransfer CustomerOfflinePaymentMethod = "jp_automated_bank_transfer"
	CustomerOfflinePaymentMethodMxAutomatedBankTransfer CustomerOfflinePaymentMethod = "mx_automated_bank_transfer"
	CustomerOfflinePaymentMethodCustom                  CustomerOfflinePaymentMethod = "custom"
)

type CustomerVatNumberStatus string

const (
	CustomerVatNumberStatusValid        CustomerVatNumberStatus = "valid"
	CustomerVatNumberStatusInvalid      CustomerVatNumberStatus = "invalid"
	CustomerVatNumberStatusNotValidated CustomerVatNumberStatus = "not_validated"
	CustomerVatNumberStatusUndetermined CustomerVatNumberStatus = "undetermined"
)

type CustomerTaxability string

const (
	CustomerTaxabilityTaxable CustomerTaxability = "taxable"
	CustomerTaxabilityExempt  CustomerTaxability = "exempt"
)

type CustomerEntityCode string

const (
	CustomerEntityCodeA    CustomerEntityCode = "a"
	CustomerEntityCodeB    CustomerEntityCode = "b"
	CustomerEntityCodeC    CustomerEntityCode = "c"
	CustomerEntityCodeD    CustomerEntityCode = "d"
	CustomerEntityCodeE    CustomerEntityCode = "e"
	CustomerEntityCodeF    CustomerEntityCode = "f"
	CustomerEntityCodeG    CustomerEntityCode = "g"
	CustomerEntityCodeH    CustomerEntityCode = "h"
	CustomerEntityCodeI    CustomerEntityCode = "i"
	CustomerEntityCodeJ    CustomerEntityCode = "j"
	CustomerEntityCodeK    CustomerEntityCode = "k"
	CustomerEntityCodeL    CustomerEntityCode = "l"
	CustomerEntityCodeM    CustomerEntityCode = "m"
	CustomerEntityCodeN    CustomerEntityCode = "n"
	CustomerEntityCodeP    CustomerEntityCode = "p"
	CustomerEntityCodeQ    CustomerEntityCode = "q"
	CustomerEntityCodeR    CustomerEntityCode = "r"
	CustomerEntityCodeMed1 CustomerEntityCode = "med1"
	CustomerEntityCodeMed2 CustomerEntityCode = "med2"
)

type CustomerBillingDateMode string

const (
	CustomerBillingDateModeUsingDefaults CustomerBillingDateMode = "using_defaults"
	CustomerBillingDateModeManuallySet   CustomerBillingDateMode = "manually_set"
)

type CustomerBillingDayOfWeek string

const (
	CustomerBillingDayOfWeekSunday    CustomerBillingDayOfWeek = "sunday"
	CustomerBillingDayOfWeekMonday    CustomerBillingDayOfWeek = "monday"
	CustomerBillingDayOfWeekTuesday   CustomerBillingDayOfWeek = "tuesday"
	CustomerBillingDayOfWeekWednesday CustomerBillingDayOfWeek = "wednesday"
	CustomerBillingDayOfWeekThursday  CustomerBillingDayOfWeek = "thursday"
	CustomerBillingDayOfWeekFriday    CustomerBillingDayOfWeek = "friday"
	CustomerBillingDayOfWeekSaturday  CustomerBillingDayOfWeek = "saturday"
)

type CustomerBillingDayOfWeekMode string

const (
	CustomerBillingDayOfWeekModeUsingDefaults CustomerBillingDayOfWeekMode = "using_defaults"
	CustomerBillingDayOfWeekModeManuallySet   CustomerBillingDayOfWeekMode = "manually_set"
)

type CustomerPiiCleared string

const (
	CustomerPiiClearedActive            CustomerPiiCleared = "active"
	CustomerPiiClearedScheduledForClear CustomerPiiCleared = "scheduled_for_clear"
	CustomerPiiClearedCleared           CustomerPiiCleared = "cleared"
)

type CustomerChannel string

const (
	CustomerChannelWeb       CustomerChannel = "web"
	CustomerChannelAppStore  CustomerChannel = "app_store"
	CustomerChannelPlayStore CustomerChannel = "play_store"
)

type CustomerCardStatus string

const (
	CustomerCardStatusNoCard              CustomerCardStatus = "no_card"
	CustomerCardStatusValid               CustomerCardStatus = "valid"
	CustomerCardStatusExpiring            CustomerCardStatus = "expiring"
	CustomerCardStatusExpired             CustomerCardStatus = "expired"
	CustomerCardStatusPendingVerification CustomerCardStatus = "pending_verification"
	CustomerCardStatusInvalid             CustomerCardStatus = "invalid"
)

type CustomerFraudFlag string

const (
	CustomerFraudFlagSafe       CustomerFraudFlag = "safe"
	CustomerFraudFlagSuspicious CustomerFraudFlag = "suspicious"
	CustomerFraudFlagFraudulent CustomerFraudFlag = "fraudulent"
)

type CustomerEinvoicingMethod string

const (
	CustomerEinvoicingMethodAutomatic   CustomerEinvoicingMethod = "automatic"
	CustomerEinvoicingMethodManual      CustomerEinvoicingMethod = "manual"
	CustomerEinvoicingMethodSiteDefault CustomerEinvoicingMethod = "site_default"
)

type CustomerCustomerType string

const (
	CustomerCustomerTypeResidential   CustomerCustomerType = "residential"
	CustomerCustomerTypeBusiness      CustomerCustomerType = "business"
	CustomerCustomerTypeSeniorCitizen CustomerCustomerType = "senior_citizen"
	CustomerCustomerTypeIndustrial    CustomerCustomerType = "industrial"
)

type CustomerBillingAddressValidationStatus string

const (
	CustomerBillingAddressValidationStatusNotValidated   CustomerBillingAddressValidationStatus = "not_validated"
	CustomerBillingAddressValidationStatusValid          CustomerBillingAddressValidationStatus = "valid"
	CustomerBillingAddressValidationStatusPartiallyValid CustomerBillingAddressValidationStatus = "partially_valid"
	CustomerBillingAddressValidationStatusInvalid        CustomerBillingAddressValidationStatus = "invalid"
)

type CustomerReferralUrlReferralSystem string

const (
	CustomerReferralUrlReferralSystemReferralCandy      CustomerReferralUrlReferralSystem = "referral_candy"
	CustomerReferralUrlReferralSystemReferralSaasquatch CustomerReferralUrlReferralSystem = "referral_saasquatch"
	CustomerReferralUrlReferralSystemFriendbuy          CustomerReferralUrlReferralSystem = "friendbuy"
)

type CustomerPaymentMethodType string

const (
	CustomerPaymentMethodTypeCard                  CustomerPaymentMethodType = "card"
	CustomerPaymentMethodTypePaypalExpressCheckout CustomerPaymentMethodType = "paypal_express_checkout"
	CustomerPaymentMethodTypeAmazonPayments        CustomerPaymentMethodType = "amazon_payments"
	CustomerPaymentMethodTypeDirectDebit           CustomerPaymentMethodType = "direct_debit"
	CustomerPaymentMethodTypeGeneric               CustomerPaymentMethodType = "generic"
	CustomerPaymentMethodTypeAlipay                CustomerPaymentMethodType = "alipay"
	CustomerPaymentMethodTypeUnionpay              CustomerPaymentMethodType = "unionpay"
	CustomerPaymentMethodTypeApplePay              CustomerPaymentMethodType = "apple_pay"
	CustomerPaymentMethodTypeWechatPay             CustomerPaymentMethodType = "wechat_pay"
	CustomerPaymentMethodTypeIdeal                 CustomerPaymentMethodType = "ideal"
	CustomerPaymentMethodTypeGooglePay             CustomerPaymentMethodType = "google_pay"
	CustomerPaymentMethodTypeSofort                CustomerPaymentMethodType = "sofort"
	CustomerPaymentMethodTypeBancontact            CustomerPaymentMethodType = "bancontact"
	CustomerPaymentMethodTypeGiropay               CustomerPaymentMethodType = "giropay"
	CustomerPaymentMethodTypeDotpay                CustomerPaymentMethodType = "dotpay"
	CustomerPaymentMethodTypeUpi                   CustomerPaymentMethodType = "upi"
	CustomerPaymentMethodTypeNetbankingEmandates   CustomerPaymentMethodType = "netbanking_emandates"
	CustomerPaymentMethodTypeVenmo                 CustomerPaymentMethodType = "venmo"
	CustomerPaymentMethodTypePayTo                 CustomerPaymentMethodType = "pay_to"
	CustomerPaymentMethodTypeFasterPayments        CustomerPaymentMethodType = "faster_payments"
	CustomerPaymentMethodTypeSepaInstantTransfer   CustomerPaymentMethodType = "sepa_instant_transfer"
	CustomerPaymentMethodTypeAutomatedBankTransfer CustomerPaymentMethodType = "automated_bank_transfer"
	CustomerPaymentMethodTypeKlarnaPayNow          CustomerPaymentMethodType = "klarna_pay_now"
	CustomerPaymentMethodTypeOnlineBankingPoland   CustomerPaymentMethodType = "online_banking_poland"
	CustomerPaymentMethodTypePayconiqByBancontact  CustomerPaymentMethodType = "payconiq_by_bancontact"
)

type CustomerPaymentMethodGateway string

const (
	CustomerPaymentMethodGatewayChargebee             CustomerPaymentMethodGateway = "chargebee"
	CustomerPaymentMethodGatewayChargebeePayments     CustomerPaymentMethodGateway = "chargebee_payments"
	CustomerPaymentMethodGatewayAdyen                 CustomerPaymentMethodGateway = "adyen"
	CustomerPaymentMethodGatewayStripe                CustomerPaymentMethodGateway = "stripe"
	CustomerPaymentMethodGatewayWepay                 CustomerPaymentMethodGateway = "wepay"
	CustomerPaymentMethodGatewayBraintree             CustomerPaymentMethodGateway = "braintree"
	CustomerPaymentMethodGatewayAuthorizeNet          CustomerPaymentMethodGateway = "authorize_net"
	CustomerPaymentMethodGatewayPaypalPro             CustomerPaymentMethodGateway = "paypal_pro"
	CustomerPaymentMethodGatewayPin                   CustomerPaymentMethodGateway = "pin"
	CustomerPaymentMethodGatewayEway                  CustomerPaymentMethodGateway = "eway"
	CustomerPaymentMethodGatewayEwayRapid             CustomerPaymentMethodGateway = "eway_rapid"
	CustomerPaymentMethodGatewayWorldpay              CustomerPaymentMethodGateway = "worldpay"
	CustomerPaymentMethodGatewayBalancedPayments      CustomerPaymentMethodGateway = "balanced_payments"
	CustomerPaymentMethodGatewayBeanstream            CustomerPaymentMethodGateway = "beanstream"
	CustomerPaymentMethodGatewayBluepay               CustomerPaymentMethodGateway = "bluepay"
	CustomerPaymentMethodGatewayElavon                CustomerPaymentMethodGateway = "elavon"
	CustomerPaymentMethodGatewayFirstDataGlobal       CustomerPaymentMethodGateway = "first_data_global"
	CustomerPaymentMethodGatewayHdfc                  CustomerPaymentMethodGateway = "hdfc"
	CustomerPaymentMethodGatewayMigs                  CustomerPaymentMethodGateway = "migs"
	CustomerPaymentMethodGatewayNmi                   CustomerPaymentMethodGateway = "nmi"
	CustomerPaymentMethodGatewayOgone                 CustomerPaymentMethodGateway = "ogone"
	CustomerPaymentMethodGatewayPaymill               CustomerPaymentMethodGateway = "paymill"
	CustomerPaymentMethodGatewayPaypalPayflowPro      CustomerPaymentMethodGateway = "paypal_payflow_pro"
	CustomerPaymentMethodGatewaySagePay               CustomerPaymentMethodGateway = "sage_pay"
	CustomerPaymentMethodGatewayTco                   CustomerPaymentMethodGateway = "tco"
	CustomerPaymentMethodGatewayWirecard              CustomerPaymentMethodGateway = "wirecard"
	CustomerPaymentMethodGatewayAmazonPayments        CustomerPaymentMethodGateway = "amazon_payments"
	CustomerPaymentMethodGatewayPaypalExpressCheckout CustomerPaymentMethodGateway = "paypal_express_checkout"
	CustomerPaymentMethodGatewayGocardless            CustomerPaymentMethodGateway = "gocardless"
	CustomerPaymentMethodGatewayOrbital               CustomerPaymentMethodGateway = "orbital"
	CustomerPaymentMethodGatewayMonerisUs             CustomerPaymentMethodGateway = "moneris_us"
	CustomerPaymentMethodGatewayMoneris               CustomerPaymentMethodGateway = "moneris"
	CustomerPaymentMethodGatewayBluesnap              CustomerPaymentMethodGateway = "bluesnap"
	CustomerPaymentMethodGatewayCybersource           CustomerPaymentMethodGateway = "cybersource"
	CustomerPaymentMethodGatewayVantiv                CustomerPaymentMethodGateway = "vantiv"
	CustomerPaymentMethodGatewayCheckoutCom           CustomerPaymentMethodGateway = "checkout_com"
	CustomerPaymentMethodGatewayPaypal                CustomerPaymentMethodGateway = "paypal"
	CustomerPaymentMethodGatewayIngenicoDirect        CustomerPaymentMethodGateway = "ingenico_direct"
	CustomerPaymentMethodGatewayExact                 CustomerPaymentMethodGateway = "exact"
	CustomerPaymentMethodGatewayMollie                CustomerPaymentMethodGateway = "mollie"
	CustomerPaymentMethodGatewayQuickbooks            CustomerPaymentMethodGateway = "quickbooks"
	CustomerPaymentMethodGatewayRazorpay              CustomerPaymentMethodGateway = "razorpay"
	CustomerPaymentMethodGatewayGlobalPayments        CustomerPaymentMethodGateway = "global_payments"
	CustomerPaymentMethodGatewayBankOfAmerica         CustomerPaymentMethodGateway = "bank_of_america"
	CustomerPaymentMethodGatewayEcentric              CustomerPaymentMethodGateway = "ecentric"
	CustomerPaymentMethodGatewayMetricsGlobal         CustomerPaymentMethodGateway = "metrics_global"
	CustomerPaymentMethodGatewayWindcave              CustomerPaymentMethodGateway = "windcave"
	CustomerPaymentMethodGatewayPayCom                CustomerPaymentMethodGateway = "pay_com"
	CustomerPaymentMethodGatewayEbanx                 CustomerPaymentMethodGateway = "ebanx"
	CustomerPaymentMethodGatewayDlocal                CustomerPaymentMethodGateway = "dlocal"
	CustomerPaymentMethodGatewayNuvei                 CustomerPaymentMethodGateway = "nuvei"
	CustomerPaymentMethodGatewaySolidgate             CustomerPaymentMethodGateway = "solidgate"
	CustomerPaymentMethodGatewayPaystack              CustomerPaymentMethodGateway = "paystack"
	CustomerPaymentMethodGatewayJpMorgan              CustomerPaymentMethodGateway = "jp_morgan"
	CustomerPaymentMethodGatewayDeutscheBank          CustomerPaymentMethodGateway = "deutsche_bank"
	CustomerPaymentMethodGatewayEzidebit              CustomerPaymentMethodGateway = "ezidebit"
	CustomerPaymentMethodGatewayNotApplicable         CustomerPaymentMethodGateway = "not_applicable"
)

type CustomerPaymentMethodStatus string

const (
	CustomerPaymentMethodStatusValid               CustomerPaymentMethodStatus = "valid"
	CustomerPaymentMethodStatusExpiring            CustomerPaymentMethodStatus = "expiring"
	CustomerPaymentMethodStatusExpired             CustomerPaymentMethodStatus = "expired"
	CustomerPaymentMethodStatusInvalid             CustomerPaymentMethodStatus = "invalid"
	CustomerPaymentMethodStatusPendingVerification CustomerPaymentMethodStatus = "pending_verification"
)

type CustomerParentAccountAccessPortalEditChildSubscriptions string

const (
	CustomerParentAccountAccessPortalEditChildSubscriptionsYes      CustomerParentAccountAccessPortalEditChildSubscriptions = "yes"
	CustomerParentAccountAccessPortalEditChildSubscriptionsViewOnly CustomerParentAccountAccessPortalEditChildSubscriptions = "view_only"
	CustomerParentAccountAccessPortalEditChildSubscriptionsNo       CustomerParentAccountAccessPortalEditChildSubscriptions = "no"
)

type CustomerParentAccountAccessPortalDownloadChildInvoices string

const (
	CustomerParentAccountAccessPortalDownloadChildInvoicesYes      CustomerParentAccountAccessPortalDownloadChildInvoices = "yes"
	CustomerParentAccountAccessPortalDownloadChildInvoicesViewOnly CustomerParentAccountAccessPortalDownloadChildInvoices = "view_only"
	CustomerParentAccountAccessPortalDownloadChildInvoicesNo       CustomerParentAccountAccessPortalDownloadChildInvoices = "no"
)

type CustomerChildAccountAccessPortalEditSubscriptions string

const (
	CustomerChildAccountAccessPortalEditSubscriptionsYes      CustomerChildAccountAccessPortalEditSubscriptions = "yes"
	CustomerChildAccountAccessPortalEditSubscriptionsViewOnly CustomerChildAccountAccessPortalEditSubscriptions = "view_only"
)

type CustomerChildAccountAccessPortalDownloadInvoices string

const (
	CustomerChildAccountAccessPortalDownloadInvoicesYes      CustomerChildAccountAccessPortalDownloadInvoices = "yes"
	CustomerChildAccountAccessPortalDownloadInvoicesViewOnly CustomerChildAccountAccessPortalDownloadInvoices = "view_only"
	CustomerChildAccountAccessPortalDownloadInvoicesNo       CustomerChildAccountAccessPortalDownloadInvoices = "no"
)

type CustomerTaxjarExemptionCategory string

const (
	CustomerTaxjarExemptionCategoryWholesale  CustomerTaxjarExemptionCategory = "wholesale"
	CustomerTaxjarExemptionCategoryGovernment CustomerTaxjarExemptionCategory = "government"
	CustomerTaxjarExemptionCategoryOther      CustomerTaxjarExemptionCategory = "other"
)

type CustomerRole string

const (
	CustomerRolePrimary CustomerRole = "primary"
	CustomerRoleBackup  CustomerRole = "backup"
	CustomerRoleNone    CustomerRole = "none"
)

type CustomerCreditType string

const (
	CustomerCreditTypeLoyaltyCredits  CustomerCreditType = "loyalty_credits"
	CustomerCreditTypeReferralRewards CustomerCreditType = "referral_rewards"
	CustomerCreditTypeGeneral         CustomerCreditType = "general"
)

type CustomerPaymentInitiator string

const (
	CustomerPaymentInitiatorCustomer CustomerPaymentInitiator = "customer"
	CustomerPaymentInitiatorMerchant CustomerPaymentInitiator = "merchant"
)

// just struct
type Customer struct {
	Id                               string                       `json:"id"`
	FirstName                        string                       `json:"first_name"`
	LastName                         string                       `json:"last_name"`
	Email                            string                       `json:"email"`
	Phone                            string                       `json:"phone"`
	Company                          string                       `json:"company"`
	VatNumber                        string                       `json:"vat_number"`
	AutoCollection                   CustomerAutoCollection       `json:"auto_collection"`
	OfflinePaymentMethod             CustomerOfflinePaymentMethod `json:"offline_payment_method"`
	NetTermDays                      int32                        `json:"net_term_days"`
	VatNumberValidatedTime           int64                        `json:"vat_number_validated_time"`
	VatNumberStatus                  CustomerVatNumberStatus      `json:"vat_number_status"`
	AllowDirectDebit                 bool                         `json:"allow_direct_debit"`
	IsLocationValid                  bool                         `json:"is_location_valid"`
	CreatedAt                        int64                        `json:"created_at"`
	CreatedFromIp                    string                       `json:"created_from_ip"`
	ExemptionDetails                 json.RawMessage              `json:"exemption_details"`
	Taxability                       CustomerTaxability           `json:"taxability"`
	EntityCode                       CustomerEntityCode           `json:"entity_code"`
	ExemptNumber                     string                       `json:"exempt_number"`
	ResourceVersion                  int64                        `json:"resource_version"`
	UpdatedAt                        int64                        `json:"updated_at"`
	Locale                           string                       `json:"locale"`
	BillingDate                      int32                        `json:"billing_date"`
	BillingMonth                     int32                        `json:"billing_month"`
	BillingDateMode                  CustomerBillingDateMode      `json:"billing_date_mode"`
	BillingDayOfWeek                 CustomerBillingDayOfWeek     `json:"billing_day_of_week"`
	BillingDayOfWeekMode             CustomerBillingDayOfWeekMode `json:"billing_day_of_week_mode"`
	PiiCleared                       CustomerPiiCleared           `json:"pii_cleared"`
	AutoCloseInvoices                bool                         `json:"auto_close_invoices"`
	Channel                          CustomerChannel              `json:"channel"`
	ActiveId                         string                       `json:"active_id"`
	FraudFlag                        CustomerFraudFlag            `json:"fraud_flag"`
	PrimaryPaymentSourceId           string                       `json:"primary_payment_source_id"`
	BackupPaymentSourceId            string                       `json:"backup_payment_source_id"`
	BillingAddress                   *CustomerBillingAddress      `json:"billing_address"`
	ReferralUrls                     []*CustomerReferralUrl       `json:"referral_urls"`
	Contacts                         []*CustomerContact           `json:"contacts"`
	PaymentMethod                    *CustomerPaymentMethod       `json:"payment_method"`
	InvoiceNotes                     string                       `json:"invoice_notes"`
	BusinessEntityId                 string                       `json:"business_entity_id"`
	PreferredCurrencyCode            string                       `json:"preferred_currency_code"`
	PromotionalCredits               int64                        `json:"promotional_credits"`
	UnbilledCharges                  int64                        `json:"unbilled_charges"`
	RefundableCredits                int64                        `json:"refundable_credits"`
	ExcessPayments                   int64                        `json:"excess_payments"`
	Balances                         []*CustomerBalance           `json:"balances"`
	EntityIdentifiers                []*CustomerEntityIdentifier  `json:"entity_identifiers"`
	TaxProvidersFields               []*CustomerTaxProvidersField `json:"tax_providers_fields"`
	IsEinvoiceEnabled                bool                         `json:"is_einvoice_enabled"`
	EinvoicingMethod                 CustomerEinvoicingMethod     `json:"einvoicing_method"`
	MetaData                         json.RawMessage              `json:"meta_data"`
	Deleted                          bool                         `json:"deleted"`
	RegisteredForGst                 bool                         `json:"registered_for_gst"`
	ConsolidatedInvoicing            bool                         `json:"consolidated_invoicing"`
	CustomerType                     CustomerCustomerType         `json:"customer_type"`
	BusinessCustomerWithoutVatNumber bool                         `json:"business_customer_without_vat_number"`
	ClientProfileId                  string                       `json:"client_profile_id"`
	Relationship                     *CustomerRelationship        `json:"relationship"`
	UseDefaultHierarchySettings      bool                         `json:"use_default_hierarchy_settings"`
	ParentAccountAccess              *CustomerParentAccountAccess `json:"parent_account_access"`
	ChildAccountAccess               *CustomerChildAccountAccess  `json:"child_account_access"`
	VatNumberPrefix                  string                       `json:"vat_number_prefix"`
	EntityIdentifierScheme           string                       `json:"entity_identifier_scheme"`
	EntityIdentifierStandard         string                       `json:"entity_identifier_standard"`
	CustomField                      CustomField                  `json:"custom_field"`
	Consents                         map[string]interface{}       `json:"consents"`
	Object                           string                       `json:"object"`
}

// sub resources
type CustomerBillingAddress struct {
	FirstName        string                                 `json:"first_name"`
	LastName         string                                 `json:"last_name"`
	Email            string                                 `json:"email"`
	Company          string                                 `json:"company"`
	Phone            string                                 `json:"phone"`
	Line1            string                                 `json:"line1"`
	Line2            string                                 `json:"line2"`
	Line3            string                                 `json:"line3"`
	City             string                                 `json:"city"`
	StateCode        string                                 `json:"state_code"`
	State            string                                 `json:"state"`
	Country          string                                 `json:"country"`
	Zip              string                                 `json:"zip"`
	ValidationStatus CustomerBillingAddressValidationStatus `json:"validation_status"`
	Object           string                                 `json:"object"`
}

type CustomerReferralUrl struct {
	ExternalCustomerId         string                            `json:"external_customer_id"`
	ReferralSharingUrl         string                            `json:"referral_sharing_url"`
	CreatedAt                  int64                             `json:"created_at"`
	UpdatedAt                  int64                             `json:"updated_at"`
	ReferralCampaignId         string                            `json:"referral_campaign_id"`
	ReferralAccountId          string                            `json:"referral_account_id"`
	ReferralExternalCampaignId string                            `json:"referral_external_campaign_id"`
	ReferralSystem             CustomerReferralUrlReferralSystem `json:"referral_system"`
	Object                     string                            `json:"object"`
}

type CustomerContact struct {
	Id               string `json:"id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Label            string `json:"label"`
	Enabled          bool   `json:"enabled"`
	SendAccountEmail bool   `json:"send_account_email"`
	SendBillingEmail bool   `json:"send_billing_email"`
	Object           string `json:"object"`
}

type CustomerPaymentMethod struct {
	Type             CustomerPaymentMethodType    `json:"type"`
	Gateway          CustomerPaymentMethodGateway `json:"gateway"`
	GatewayAccountId string                       `json:"gateway_account_id"`
	Status           CustomerPaymentMethodStatus  `json:"status"`
	ReferenceId      string                       `json:"reference_id"`
	Object           string                       `json:"object"`
}

type CustomerBalance struct {
	PromotionalCredits  int64  `json:"promotional_credits"`
	ExcessPayments      int64  `json:"excess_payments"`
	RefundableCredits   int64  `json:"refundable_credits"`
	UnbilledCharges     int64  `json:"unbilled_charges"`
	CurrencyCode        string `json:"currency_code"`
	BalanceCurrencyCode string `json:"balance_currency_code"`
	BusinessEntityId    string `json:"business_entity_id"`
	Object              string `json:"object"`
}

type CustomerEntityIdentifier struct {
	Id       string `json:"id"`
	Value    string `json:"value"`
	Scheme   string `json:"scheme"`
	Standard string `json:"standard"`
	Object   string `json:"object"`
}

type CustomerTaxProvidersField struct {
	ProviderName string `json:"provider_name"`
	FieldId      string `json:"field_id"`
	FieldValue   string `json:"field_value"`
	Object       string `json:"object"`
}

type CustomerRelationship struct {
	ParentId       string `json:"parent_id"`
	PaymentOwnerId string `json:"payment_owner_id"`
	InvoiceOwnerId string `json:"invoice_owner_id"`
	Object         string `json:"object"`
}

type CustomerParentAccountAccess struct {
	PortalEditChildSubscriptions CustomerParentAccountAccessPortalEditChildSubscriptions `json:"portal_edit_child_subscriptions"`
	PortalDownloadChildInvoices  CustomerParentAccountAccessPortalDownloadChildInvoices  `json:"portal_download_child_invoices"`
	SendSubscriptionEmails       bool                                                    `json:"send_subscription_emails"`
	SendInvoiceEmails            bool                                                    `json:"send_invoice_emails"`
	SendPaymentEmails            bool                                                    `json:"send_payment_emails"`
	Object                       string                                                  `json:"object"`
}

type CustomerChildAccountAccess struct {
	PortalEditSubscriptions CustomerChildAccountAccessPortalEditSubscriptions `json:"portal_edit_subscriptions"`
	PortalDownloadInvoices  CustomerChildAccountAccessPortalDownloadInvoices  `json:"portal_download_invoices"`
	SendSubscriptionEmails  bool                                              `json:"send_subscription_emails"`
	SendInvoiceEmails       bool                                              `json:"send_invoice_emails"`
	SendPaymentEmails       bool                                              `json:"send_payment_emails"`
	Object                  string                                            `json:"object"`
}

// operations
// input params
type CustomerCreateRequest struct {
	Id                               string                             `json:"id,omitempty"`
	FirstName                        string                             `json:"first_name,omitempty"`
	LastName                         string                             `json:"last_name,omitempty"`
	Email                            string                             `json:"email,omitempty"`
	PreferredCurrencyCode            string                             `json:"preferred_currency_code,omitempty"`
	Phone                            string                             `json:"phone,omitempty"`
	Company                          string                             `json:"company,omitempty"`
	AutoCollection                   CustomerAutoCollection             `json:"auto_collection,omitempty"`
	NetTermDays                      *int32                             `json:"net_term_days,omitempty"`
	AllowDirectDebit                 *bool                              `json:"allow_direct_debit,omitempty"`
	VatNumber                        string                             `json:"vat_number,omitempty"`
	VatNumberPrefix                  string                             `json:"vat_number_prefix,omitempty"`
	EntityIdentifierScheme           string                             `json:"entity_identifier_scheme,omitempty"`
	EntityIdentifierStandard         string                             `json:"entity_identifier_standard,omitempty"`
	RegisteredForGst                 *bool                              `json:"registered_for_gst,omitempty"`
	IsEinvoiceEnabled                *bool                              `json:"is_einvoice_enabled,omitempty"`
	EinvoicingMethod                 CustomerEinvoicingMethod           `json:"einvoicing_method,omitempty"`
	Taxability                       CustomerTaxability                 `json:"taxability,omitempty"`
	ExemptionDetails                 []map[string]interface{}           `json:"exemption_details,omitempty"`
	CustomerType                     CustomerCustomerType               `json:"customer_type,omitempty"`
	ClientProfileId                  string                             `json:"client_profile_id,omitempty"`
	TaxjarExemptionCategory          CustomerTaxjarExemptionCategory    `json:"taxjar_exemption_category,omitempty"`
	BusinessCustomerWithoutVatNumber *bool                              `json:"business_customer_without_vat_number,omitempty"`
	Locale                           string                             `json:"locale,omitempty"`
	EntityCode                       CustomerEntityCode                 `json:"entity_code,omitempty"`
	ExemptNumber                     string                             `json:"exempt_number,omitempty"`
	MetaData                         map[string]interface{}             `json:"meta_data,omitempty"`
	OfflinePaymentMethod             CustomerOfflinePaymentMethod       `json:"offline_payment_method,omitempty"`
	AutoCloseInvoices                *bool                              `json:"auto_close_invoices,omitempty"`
	ConsolidatedInvoicing            *bool                              `json:"consolidated_invoicing,omitempty"`
	Card                             *CustomerCreateCard                `json:"card,omitempty"`
	BankAccount                      *CustomerCreateBankAccount         `json:"bank_account,omitempty"`
	TokenId                          string                             `json:"token_id,omitempty"`
	PaymentMethod                    *CustomerCreatePaymentMethod       `json:"payment_method,omitempty"`
	PaymentIntent                    *CustomerCreatePaymentIntent       `json:"payment_intent,omitempty"`
	BillingAddress                   *CustomerCreateBillingAddress      `json:"billing_address,omitempty"`
	EntityIdentifiers                []*CustomerCreateEntityIdentifier  `json:"entity_identifiers,omitempty"`
	BusinessEntityId                 string                             `json:"business_entity_id,omitempty"`
	TaxProvidersFields               []*CustomerCreateTaxProvidersField `json:"tax_providers_fields,omitempty"`
	InvoiceNotes                     string                             `json:"invoice_notes,omitempty"`
	apiRequest                       `json:"-" form:"-"`
}

func (r *CustomerCreateRequest) payload() any { return r }

// input sub resource params single
type CustomerCreateCard struct {
	Gateway               CustomerCardGateway         `json:"gateway,omitempty"`
	GatewayAccountId      string                      `json:"gateway_account_id,omitempty"`
	TmpToken              string                      `json:"tmp_token,omitempty"`
	FirstName             string                      `json:"first_name,omitempty"`
	LastName              string                      `json:"last_name,omitempty"`
	Number                string                      `json:"number,omitempty"`
	ExpiryMonth           *int32                      `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                      `json:"expiry_year,omitempty"`
	Cvv                   string                      `json:"cvv,omitempty"`
	PreferredScheme       CustomerCardPreferredScheme `json:"preferred_scheme,omitempty"`
	BillingAddr1          string                      `json:"billing_addr1,omitempty"`
	BillingAddr2          string                      `json:"billing_addr2,omitempty"`
	BillingCity           string                      `json:"billing_city,omitempty"`
	BillingStateCode      string                      `json:"billing_state_code,omitempty"`
	BillingState          string                      `json:"billing_state,omitempty"`
	BillingZip            string                      `json:"billing_zip,omitempty"`
	BillingCountry        string                      `json:"billing_country,omitempty"`
	IpAddress             string                      `json:"ip_address,omitempty"`
	AdditionalInformation map[string]interface{}      `json:"additional_information,omitempty"`
}

// input sub resource params single
type CustomerCreateBankAccount struct {
	GatewayAccountId      string                               `json:"gateway_account_id,omitempty"`
	Iban                  string                               `json:"iban,omitempty"`
	FirstName             string                               `json:"first_name,omitempty"`
	LastName              string                               `json:"last_name,omitempty"`
	Company               string                               `json:"company,omitempty"`
	Email                 string                               `json:"email,omitempty"`
	Phone                 string                               `json:"phone,omitempty"`
	BankName              string                               `json:"bank_name,omitempty"`
	AccountNumber         string                               `json:"account_number,omitempty"`
	RoutingNumber         string                               `json:"routing_number,omitempty"`
	BankCode              string                               `json:"bank_code,omitempty"`
	AccountType           CustomerBankAccountAccountType       `json:"account_type,omitempty"`
	AccountHolderType     CustomerBankAccountAccountHolderType `json:"account_holder_type,omitempty"`
	EcheckType            CustomerBankAccountEcheckType        `json:"echeck_type,omitempty"`
	IssuingCountry        string                               `json:"issuing_country,omitempty"`
	SwedishIdentityNumber string                               `json:"swedish_identity_number,omitempty"`
	BillingAddress        map[string]interface{}               `json:"billing_address,omitempty"`
}

// input sub resource params single
type CustomerCreatePaymentMethod struct {
	Type                  CustomerPaymentMethodType    `json:"type,omitempty"`
	Gateway               CustomerPaymentMethodGateway `json:"gateway,omitempty"`
	GatewayAccountId      string                       `json:"gateway_account_id,omitempty"`
	ReferenceId           string                       `json:"reference_id,omitempty"`
	TmpToken              string                       `json:"tmp_token,omitempty"`
	IssuingCountry        string                       `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{}       `json:"additional_information,omitempty"`
}

// input sub resource params single
type CustomerCreatePaymentIntent struct {
	Id                    string                                 `json:"id,omitempty"`
	GatewayAccountId      string                                 `json:"gateway_account_id,omitempty"`
	GwToken               string                                 `json:"gw_token,omitempty"`
	PaymentMethodType     CustomerPaymentIntentPaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                                 `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                                 `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}                 `json:"additional_information,omitempty"`
}

// input sub resource params single
type CustomerCreateBillingAddress struct {
	FirstName        string                                 `json:"first_name,omitempty"`
	LastName         string                                 `json:"last_name,omitempty"`
	Email            string                                 `json:"email,omitempty"`
	Company          string                                 `json:"company,omitempty"`
	Phone            string                                 `json:"phone,omitempty"`
	Line1            string                                 `json:"line1,omitempty"`
	Line2            string                                 `json:"line2,omitempty"`
	Line3            string                                 `json:"line3,omitempty"`
	City             string                                 `json:"city,omitempty"`
	StateCode        string                                 `json:"state_code,omitempty"`
	State            string                                 `json:"state,omitempty"`
	Zip              string                                 `json:"zip,omitempty"`
	Country          string                                 `json:"country,omitempty"`
	ValidationStatus CustomerBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params multi
type CustomerCreateEntityIdentifier struct {
	Id       string `json:"id,omitempty"`
	Scheme   string `json:"scheme,omitempty"`
	Value    string `json:"value,omitempty"`
	Standard string `json:"standard,omitempty"`
}

// input sub resource params multi
type CustomerCreateTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}

type CustomerListRequest struct {
	Limit                *int32            `json:"limit,omitempty"`
	Offset               string            `json:"offset,omitempty"`
	Relationship         *ListRelationship `json:"relationship,omitempty"`
	IncludeDeleted       *bool             `json:"include_deleted,omitempty"`
	Id                   *StringFilter     `json:"id,omitempty"`
	FirstName            *StringFilter     `json:"first_name,omitempty"`
	LastName             *StringFilter     `json:"last_name,omitempty"`
	Email                *StringFilter     `json:"email,omitempty"`
	Company              *StringFilter     `json:"company,omitempty"`
	Phone                *StringFilter     `json:"phone,omitempty"`
	AutoCollection       *EnumFilter       `json:"auto_collection,omitempty"`
	Taxability           *EnumFilter       `json:"taxability,omitempty"`
	CreatedAt            *TimestampFilter  `json:"created_at,omitempty"`
	UpdatedAt            *TimestampFilter  `json:"updated_at,omitempty"`
	OfflinePaymentMethod *EnumFilter       `json:"offline_payment_method,omitempty"`
	AutoCloseInvoices    *BooleanFilter    `json:"auto_close_invoices,omitempty"`
	Channel              *EnumFilter       `json:"channel,omitempty"`
	BusinessEntityId     *StringFilter     `json:"business_entity_id,omitempty"`
	SortBy               *SortFilter       `json:"sort_by,omitempty"`
	apiRequest           `json:"-" form:"-"`
}

func (r *CustomerListRequest) payload() any { return r }

// input sub resource params single
type CustomerListRelationship struct {
	ParentId       *StringFilter `json:"parent_id,omitempty"`
	PaymentOwnerId *StringFilter `json:"payment_owner_id,omitempty"`
	InvoiceOwnerId *StringFilter `json:"invoice_owner_id,omitempty"`
}

type CustomerUpdateRequest struct {
	FirstName               string                             `json:"first_name,omitempty"`
	LastName                string                             `json:"last_name,omitempty"`
	Email                   string                             `json:"email,omitempty"`
	PreferredCurrencyCode   string                             `json:"preferred_currency_code,omitempty"`
	Phone                   string                             `json:"phone,omitempty"`
	Company                 string                             `json:"company,omitempty"`
	AutoCollection          CustomerAutoCollection             `json:"auto_collection,omitempty"`
	AllowDirectDebit        *bool                              `json:"allow_direct_debit,omitempty"`
	NetTermDays             *int32                             `json:"net_term_days,omitempty"`
	Taxability              CustomerTaxability                 `json:"taxability,omitempty"`
	ExemptionDetails        []map[string]interface{}           `json:"exemption_details,omitempty"`
	CustomerType            CustomerCustomerType               `json:"customer_type,omitempty"`
	ClientProfileId         string                             `json:"client_profile_id,omitempty"`
	TaxjarExemptionCategory CustomerTaxjarExemptionCategory    `json:"taxjar_exemption_category,omitempty"`
	Locale                  string                             `json:"locale,omitempty"`
	EntityCode              CustomerEntityCode                 `json:"entity_code,omitempty"`
	ExemptNumber            string                             `json:"exempt_number,omitempty"`
	OfflinePaymentMethod    CustomerOfflinePaymentMethod       `json:"offline_payment_method,omitempty"`
	InvoiceNotes            string                             `json:"invoice_notes,omitempty"`
	AutoCloseInvoices       *bool                              `json:"auto_close_invoices,omitempty"`
	MetaData                map[string]interface{}             `json:"meta_data,omitempty"`
	FraudFlag               CustomerFraudFlag                  `json:"fraud_flag,omitempty"`
	ConsolidatedInvoicing   *bool                              `json:"consolidated_invoicing,omitempty"`
	TaxProvidersFields      []*CustomerUpdateTaxProvidersField `json:"tax_providers_fields,omitempty"`
	apiRequest              `json:"-" form:"-"`
}

func (r *CustomerUpdateRequest) payload() any { return r }

// input sub resource params multi
type CustomerUpdateTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}

type CustomerUpdatePaymentMethodRequest struct {
	PaymentMethod *CustomerUpdatePaymentMethodPaymentMethod `json:"payment_method,omitempty"`
	apiRequest    `json:"-" form:"-"`
}

func (r *CustomerUpdatePaymentMethodRequest) payload() any { return r }

// input sub resource params single
type CustomerUpdatePaymentMethodPaymentMethod struct {
	Type                  CustomerPaymentMethodType    `json:"type"`
	Gateway               CustomerPaymentMethodGateway `json:"gateway,omitempty"`
	GatewayAccountId      string                       `json:"gateway_account_id,omitempty"`
	ReferenceId           string                       `json:"reference_id,omitempty"`
	TmpToken              string                       `json:"tmp_token,omitempty"`
	IssuingCountry        string                       `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{}       `json:"additional_information,omitempty"`
}

type CustomerUpdateBillingInfoRequest struct {
	BillingAddress                   *CustomerUpdateBillingInfoBillingAddress      `json:"billing_address,omitempty"`
	EntityIdentifiers                []*CustomerUpdateBillingInfoEntityIdentifier  `json:"entity_identifiers,omitempty"`
	VatNumber                        string                                        `json:"vat_number,omitempty"`
	VatNumberPrefix                  string                                        `json:"vat_number_prefix,omitempty"`
	EntityIdentifierScheme           string                                        `json:"entity_identifier_scheme,omitempty"`
	EntityIdentifierStandard         string                                        `json:"entity_identifier_standard,omitempty"`
	RegisteredForGst                 *bool                                         `json:"registered_for_gst,omitempty"`
	BusinessCustomerWithoutVatNumber *bool                                         `json:"business_customer_without_vat_number,omitempty"`
	IsEinvoiceEnabled                *bool                                         `json:"is_einvoice_enabled,omitempty"`
	EinvoicingMethod                 CustomerEinvoicingMethod                      `json:"einvoicing_method,omitempty"`
	TaxProvidersFields               []*CustomerUpdateBillingInfoTaxProvidersField `json:"tax_providers_fields,omitempty"`
	apiRequest                       `json:"-" form:"-"`
}

func (r *CustomerUpdateBillingInfoRequest) payload() any { return r }

// input sub resource params single
type CustomerUpdateBillingInfoBillingAddress struct {
	FirstName        string                                 `json:"first_name,omitempty"`
	LastName         string                                 `json:"last_name,omitempty"`
	Email            string                                 `json:"email,omitempty"`
	Company          string                                 `json:"company,omitempty"`
	Phone            string                                 `json:"phone,omitempty"`
	Line1            string                                 `json:"line1,omitempty"`
	Line2            string                                 `json:"line2,omitempty"`
	Line3            string                                 `json:"line3,omitempty"`
	City             string                                 `json:"city,omitempty"`
	StateCode        string                                 `json:"state_code,omitempty"`
	State            string                                 `json:"state,omitempty"`
	Zip              string                                 `json:"zip,omitempty"`
	Country          string                                 `json:"country,omitempty"`
	ValidationStatus CustomerBillingAddressValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params multi
type CustomerUpdateBillingInfoEntityIdentifier struct {
	Id        string                            `json:"id,omitempty"`
	Scheme    string                            `json:"scheme,omitempty"`
	Value     string                            `json:"value,omitempty"`
	Operation CustomerEntityIdentifierOperation `json:"operation,omitempty"`
	Standard  string                            `json:"standard,omitempty"`
}

// input sub resource params multi
type CustomerUpdateBillingInfoTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}

type CustomerContactsForCustomerRequest struct {
	Limit      *int32 `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *CustomerContactsForCustomerRequest) payload() any { return r }

type CustomerAssignPaymentRoleRequest struct {
	PaymentSourceId string       `json:"payment_source_id"`
	Role            CustomerRole `json:"role"`
	apiRequest      `json:"-" form:"-"`
}

func (r *CustomerAssignPaymentRoleRequest) payload() any { return r }

type CustomerAddContactRequest struct {
	Contact    *CustomerAddContactContact `json:"contact,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *CustomerAddContactRequest) payload() any { return r }

// input sub resource params single
type CustomerAddContactContact struct {
	Id               string `json:"id,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Email            string `json:"email"`
	Phone            string `json:"phone,omitempty"`
	Label            string `json:"label,omitempty"`
	Enabled          *bool  `json:"enabled,omitempty"`
	SendBillingEmail *bool  `json:"send_billing_email,omitempty"`
	SendAccountEmail *bool  `json:"send_account_email,omitempty"`
}

type CustomerUpdateContactRequest struct {
	Contact    *CustomerUpdateContactContact `json:"contact,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *CustomerUpdateContactRequest) payload() any { return r }

// input sub resource params single
type CustomerUpdateContactContact struct {
	Id               string `json:"id"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Email            string `json:"email,omitempty"`
	Phone            string `json:"phone,omitempty"`
	Label            string `json:"label,omitempty"`
	Enabled          *bool  `json:"enabled,omitempty"`
	SendBillingEmail *bool  `json:"send_billing_email,omitempty"`
	SendAccountEmail *bool  `json:"send_account_email,omitempty"`
}

type CustomerDeleteContactRequest struct {
	Contact    *CustomerDeleteContactContact `json:"contact,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *CustomerDeleteContactRequest) payload() any { return r }

// input sub resource params single
type CustomerDeleteContactContact struct {
	Id string `json:"id"`
}

type CustomerAddPromotionalCreditsRequest struct {
	Amount       *int64             `json:"amount"`
	CurrencyCode string             `json:"currency_code,omitempty"`
	Description  string             `json:"description"`
	CreditType   CustomerCreditType `json:"credit_type,omitempty"`
	Reference    string             `json:"reference,omitempty"`
	apiRequest   `json:"-" form:"-"`
}

func (r *CustomerAddPromotionalCreditsRequest) payload() any { return r }

type CustomerDeductPromotionalCreditsRequest struct {
	Amount       *int64             `json:"amount"`
	CurrencyCode string             `json:"currency_code,omitempty"`
	Description  string             `json:"description"`
	CreditType   CustomerCreditType `json:"credit_type,omitempty"`
	Reference    string             `json:"reference,omitempty"`
	apiRequest   `json:"-" form:"-"`
}

func (r *CustomerDeductPromotionalCreditsRequest) payload() any { return r }

type CustomerSetPromotionalCreditsRequest struct {
	Amount       *int64             `json:"amount"`
	CurrencyCode string             `json:"currency_code,omitempty"`
	Description  string             `json:"description"`
	CreditType   CustomerCreditType `json:"credit_type,omitempty"`
	Reference    string             `json:"reference,omitempty"`
	apiRequest   `json:"-" form:"-"`
}

func (r *CustomerSetPromotionalCreditsRequest) payload() any { return r }

type CustomerRecordExcessPaymentRequest struct {
	Transaction *CustomerRecordExcessPaymentTransaction `json:"transaction,omitempty"`
	Comment     string                                  `json:"comment,omitempty"`
	apiRequest  `json:"-" form:"-"`
}

func (r *CustomerRecordExcessPaymentRequest) payload() any { return r }

// input sub resource params single
type CustomerRecordExcessPaymentTransaction struct {
	Id                    string                           `json:"id,omitempty"`
	Amount                *int64                           `json:"amount"`
	CurrencyCode          string                           `json:"currency_code,omitempty"`
	Date                  *int64                           `json:"date"`
	PaymentMethod         CustomerTransactionPaymentMethod `json:"payment_method"`
	ReferenceNumber       string                           `json:"reference_number,omitempty"`
	CustomPaymentMethodId string                           `json:"custom_payment_method_id,omitempty"`
}

type CustomerCollectPaymentRequest struct {
	Amount                      *int64                                     `json:"amount,omitempty"`
	InvoiceAllocations          []*CustomerCollectPaymentInvoiceAllocation `json:"invoice_allocations,omitempty"`
	PaymentSourceId             string                                     `json:"payment_source_id,omitempty"`
	TokenId                     string                                     `json:"token_id,omitempty"`
	PaymentMethod               *CustomerCollectPaymentPaymentMethod       `json:"payment_method,omitempty"`
	Card                        *CustomerCollectPaymentCard                `json:"card,omitempty"`
	PaymentIntent               *CustomerCollectPaymentPaymentIntent       `json:"payment_intent,omitempty"`
	ReplacePrimaryPaymentSource *bool                                      `json:"replace_primary_payment_source,omitempty"`
	RetainPaymentSource         *bool                                      `json:"retain_payment_source,omitempty"`
	PaymentInitiator            CustomerPaymentInitiator                   `json:"payment_initiator,omitempty"`
	apiRequest                  `json:"-" form:"-"`
}

func (r *CustomerCollectPaymentRequest) payload() any { return r }

// input sub resource params multi
type CustomerCollectPaymentInvoiceAllocation struct {
	InvoiceId        string `json:"invoice_id"`
	AllocationAmount *int64 `json:"allocation_amount,omitempty"`
}

// input sub resource params single
type CustomerCollectPaymentPaymentMethod struct {
	Type                  CustomerPaymentMethodType `json:"type,omitempty"`
	GatewayAccountId      string                    `json:"gateway_account_id,omitempty"`
	ReferenceId           string                    `json:"reference_id,omitempty"`
	TmpToken              string                    `json:"tmp_token,omitempty"`
	AdditionalInformation map[string]interface{}    `json:"additional_information,omitempty"`
}

// input sub resource params single
type CustomerCollectPaymentCard struct {
	GatewayAccountId      string                      `json:"gateway_account_id,omitempty"`
	FirstName             string                      `json:"first_name,omitempty"`
	LastName              string                      `json:"last_name,omitempty"`
	Number                string                      `json:"number,omitempty"`
	ExpiryMonth           *int32                      `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                      `json:"expiry_year,omitempty"`
	Cvv                   string                      `json:"cvv,omitempty"`
	PreferredScheme       CustomerCardPreferredScheme `json:"preferred_scheme,omitempty"`
	BillingAddr1          string                      `json:"billing_addr1,omitempty"`
	BillingAddr2          string                      `json:"billing_addr2,omitempty"`
	BillingCity           string                      `json:"billing_city,omitempty"`
	BillingStateCode      string                      `json:"billing_state_code,omitempty"`
	BillingState          string                      `json:"billing_state,omitempty"`
	BillingZip            string                      `json:"billing_zip,omitempty"`
	BillingCountry        string                      `json:"billing_country,omitempty"`
	AdditionalInformation map[string]interface{}      `json:"additional_information,omitempty"`
}

// input sub resource params single
type CustomerCollectPaymentPaymentIntent struct {
	Id                    string                                 `json:"id,omitempty"`
	GatewayAccountId      string                                 `json:"gateway_account_id,omitempty"`
	GwToken               string                                 `json:"gw_token,omitempty"`
	PaymentMethodType     CustomerPaymentIntentPaymentMethodType `json:"payment_method_type,omitempty"`
	GwPaymentMethodId     string                                 `json:"gw_payment_method_id,omitempty"`
	ReferenceId           string                                 `json:"reference_id,omitempty"`
	AdditionalInformation map[string]interface{}                 `json:"additional_information,omitempty"`
}

type CustomerDeleteRequest struct {
	DeletePaymentMethod *bool `json:"delete_payment_method,omitempty"`
	apiRequest          `json:"-" form:"-"`
}

func (r *CustomerDeleteRequest) payload() any { return r }

type CustomerMoveRequest struct {
	IdAtFromSite string `json:"id_at_from_site"`
	FromSite     string `json:"from_site"`
	apiRequest   `json:"-" form:"-"`
}

func (r *CustomerMoveRequest) payload() any { return r }

type CustomerChangeBillingDateRequest struct {
	BillingDate          *int32                       `json:"billing_date,omitempty"`
	BillingMonth         *int32                       `json:"billing_month,omitempty"`
	BillingDateMode      CustomerBillingDateMode      `json:"billing_date_mode,omitempty"`
	BillingDayOfWeek     CustomerBillingDayOfWeek     `json:"billing_day_of_week,omitempty"`
	BillingDayOfWeekMode CustomerBillingDayOfWeekMode `json:"billing_day_of_week_mode,omitempty"`
	apiRequest           `json:"-" form:"-"`
}

func (r *CustomerChangeBillingDateRequest) payload() any { return r }

type CustomerMergeRequest struct {
	FromCustomerId string `json:"from_customer_id"`
	ToCustomerId   string `json:"to_customer_id"`
	apiRequest     `json:"-" form:"-"`
}

func (r *CustomerMergeRequest) payload() any { return r }

type CustomerRelationshipsRequest struct {
	ParentId                    string                                    `json:"parent_id,omitempty"`
	PaymentOwnerId              string                                    `json:"payment_owner_id,omitempty"`
	InvoiceOwnerId              string                                    `json:"invoice_owner_id,omitempty"`
	UseDefaultHierarchySettings *bool                                     `json:"use_default_hierarchy_settings,omitempty"`
	ParentAccountAccess         *CustomerRelationshipsParentAccountAccess `json:"parent_account_access,omitempty"`
	ChildAccountAccess          *CustomerRelationshipsChildAccountAccess  `json:"child_account_access,omitempty"`
	apiRequest                  `json:"-" form:"-"`
}

func (r *CustomerRelationshipsRequest) payload() any { return r }

// input sub resource params single
type CustomerRelationshipsParentAccountAccess struct {
	PortalEditChildSubscriptions CustomerParentAccountAccessPortalEditChildSubscriptions `json:"portal_edit_child_subscriptions,omitempty"`
	PortalDownloadChildInvoices  CustomerParentAccountAccessPortalDownloadChildInvoices  `json:"portal_download_child_invoices,omitempty"`
	SendSubscriptionEmails       *bool                                                   `json:"send_subscription_emails,omitempty"`
	SendPaymentEmails            *bool                                                   `json:"send_payment_emails,omitempty"`
	SendInvoiceEmails            *bool                                                   `json:"send_invoice_emails,omitempty"`
}

// input sub resource params single
type CustomerRelationshipsChildAccountAccess struct {
	PortalEditSubscriptions CustomerChildAccountAccessPortalEditSubscriptions `json:"portal_edit_subscriptions,omitempty"`
	PortalDownloadInvoices  CustomerChildAccountAccessPortalDownloadInvoices  `json:"portal_download_invoices,omitempty"`
	SendSubscriptionEmails  *bool                                             `json:"send_subscription_emails,omitempty"`
	SendPaymentEmails       *bool                                             `json:"send_payment_emails,omitempty"`
	SendInvoiceEmails       *bool                                             `json:"send_invoice_emails,omitempty"`
}

type CustomerHierarchyRequest struct {
	HierarchyOperationType CustomerHierarchyOperationType `json:"hierarchy_operation_type"`
	apiRequest             `json:"-" form:"-"`
}

func (r *CustomerHierarchyRequest) payload() any { return r }

type CustomerListHierarchyDetailRequest struct {
	Limit                  *int32                         `json:"limit,omitempty"`
	Offset                 string                         `json:"offset,omitempty"`
	HierarchyOperationType CustomerHierarchyOperationType `json:"hierarchy_operation_type"`
	apiRequest             `json:"-" form:"-"`
}

func (r *CustomerListHierarchyDetailRequest) payload() any { return r }

type CustomerUpdateHierarchySettingsRequest struct {
	UseDefaultHierarchySettings *bool                                               `json:"use_default_hierarchy_settings,omitempty"`
	ParentAccountAccess         *CustomerUpdateHierarchySettingsParentAccountAccess `json:"parent_account_access,omitempty"`
	ChildAccountAccess          *CustomerUpdateHierarchySettingsChildAccountAccess  `json:"child_account_access,omitempty"`
	apiRequest                  `json:"-" form:"-"`
}

func (r *CustomerUpdateHierarchySettingsRequest) payload() any { return r }

// input sub resource params single
type CustomerUpdateHierarchySettingsParentAccountAccess struct {
	PortalEditChildSubscriptions CustomerParentAccountAccessPortalEditChildSubscriptions `json:"portal_edit_child_subscriptions,omitempty"`
	PortalDownloadChildInvoices  CustomerParentAccountAccessPortalDownloadChildInvoices  `json:"portal_download_child_invoices,omitempty"`
	SendSubscriptionEmails       *bool                                                   `json:"send_subscription_emails,omitempty"`
	SendPaymentEmails            *bool                                                   `json:"send_payment_emails,omitempty"`
	SendInvoiceEmails            *bool                                                   `json:"send_invoice_emails,omitempty"`
}

// input sub resource params single
type CustomerUpdateHierarchySettingsChildAccountAccess struct {
	PortalEditSubscriptions CustomerChildAccountAccessPortalEditSubscriptions `json:"portal_edit_subscriptions,omitempty"`
	PortalDownloadInvoices  CustomerChildAccountAccessPortalDownloadInvoices  `json:"portal_download_invoices,omitempty"`
	SendSubscriptionEmails  *bool                                             `json:"send_subscription_emails,omitempty"`
	SendPaymentEmails       *bool                                             `json:"send_payment_emails,omitempty"`
	SendInvoiceEmails       *bool                                             `json:"send_invoice_emails,omitempty"`
}

// operation response
type CustomerCreateResponse struct {
	Customer *Customer `json:"customer,omitempty"`
	Card     Card      `json:"card,omitempty"`
	apiResponse
}

// operation sub response
type CustomerListCustomerResponse struct {
	Customer *Customer `json:"customer,omitempty"`
	Card     Card      `json:"card,omitempty"`
}

// operation response
type CustomerListResponse struct {
	List       []*CustomerListCustomerResponse `json:"list,omitempty"`
	NextOffset string                          `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type CustomerRetrieveResponse struct {
	Customer *Customer `json:"customer,omitempty"`
	Card     Card      `json:"card,omitempty"`
	apiResponse
}

// operation response
type CustomerUpdateResponse struct {
	Customer *Customer `json:"customer,omitempty"`
	Card     Card      `json:"card,omitempty"`
	apiResponse
}

// operation response
type CustomerUpdatePaymentMethodResponse struct {
	Customer *Customer `json:"customer,omitempty"`
	Card     Card      `json:"card,omitempty"`
	apiResponse
}

// operation response
type CustomerUpdateBillingInfoResponse struct {
	Customer *Customer `json:"customer,omitempty"`
	Card     Card      `json:"card,omitempty"`
	apiResponse
}

// operation sub response
type CustomerContactsForCustomerCustomerResponse struct {
	Contact Contact `json:"contact,omitempty"`
}

// operation response
type CustomerContactsForCustomerResponse struct {
	List       []*CustomerContactsForCustomerCustomerResponse `json:"list,omitempty"`
	NextOffset string                                         `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type CustomerAssignPaymentRoleResponse struct {
	Customer      *Customer     `json:"customer,omitempty"`
	PaymentSource PaymentSource `json:"payment_source,omitempty"`
	apiResponse
}

// operation response
type CustomerAddContactResponse struct {
	Customer *Customer `json:"customer,omitempty"`
	Card     Card      `json:"card,omitempty"`
	apiResponse
}

// operation response
type CustomerUpdateContactResponse struct {
	Customer *Customer `json:"customer,omitempty"`
	Card     Card      `json:"card,omitempty"`
	apiResponse
}

// operation response
type CustomerDeleteContactResponse struct {
	Customer *Customer `json:"customer,omitempty"`
	Card     Card      `json:"card,omitempty"`
	apiResponse
}

// operation response
type CustomerAddPromotionalCreditsResponse struct {
	Customer *Customer `json:"customer,omitempty"`
	apiResponse
}

// operation response
type CustomerDeductPromotionalCreditsResponse struct {
	Customer *Customer `json:"customer,omitempty"`
	apiResponse
}

// operation response
type CustomerSetPromotionalCreditsResponse struct {
	Customer *Customer `json:"customer,omitempty"`
	apiResponse
}

// operation response
type CustomerRecordExcessPaymentResponse struct {
	Customer    *Customer   `json:"customer,omitempty"`
	Transaction Transaction `json:"transaction,omitempty"`
	apiResponse
}

// operation response
type CustomerCollectPaymentResponse struct {
	Customer    *Customer   `json:"customer,omitempty"`
	Transaction Transaction `json:"transaction,omitempty"`
	apiResponse
}

// operation response
type CustomerDeleteResponse struct {
	Customer *Customer `json:"customer,omitempty"`
	Card     Card      `json:"card,omitempty"`
	apiResponse
}

// operation response
type CustomerMoveResponse struct {
	ResourceMigration ResourceMigration `json:"resource_migration,omitempty"`
	apiResponse
}

// operation response
type CustomerChangeBillingDateResponse struct {
	Customer *Customer `json:"customer,omitempty"`
	apiResponse
}

// operation response
type CustomerMergeResponse struct {
	Customer *Customer `json:"customer,omitempty"`
	apiResponse
}

// operation response
type CustomerClearPersonalDataResponse struct {
	Customer *Customer `json:"customer,omitempty"`
	apiResponse
}

// operation response
type CustomerRelationshipsResponse struct {
	Customer *Customer `json:"customer,omitempty"`
	apiResponse
}

// operation response
type CustomerDeleteRelationshipResponse struct {
	Customer *Customer `json:"customer,omitempty"`
	apiResponse
}

// operation response
type CustomerHierarchyResponse struct {
	Hierarchies []Hierarchy `json:"hierarchies,omitempty"`
	apiResponse
}

// operation sub response
type CustomerListHierarchyDetailCustomerResponse struct {
	Hierarchy Hierarchy `json:"hierarchy,omitempty"`
}

// operation response
type CustomerListHierarchyDetailResponse struct {
	List       []*CustomerListHierarchyDetailCustomerResponse `json:"list,omitempty"`
	NextOffset string                                         `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type CustomerUpdateHierarchySettingsResponse struct {
	Customer *Customer `json:"customer,omitempty"`
	apiResponse
}
