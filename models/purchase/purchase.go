package purchase

import (
	"github.com/chargebee/chargebee-go/v3/enum"
	contractTermEnum "github.com/chargebee/chargebee-go/v3/models/contractterm/enum"
	paymentIntentEnum "github.com/chargebee/chargebee-go/v3/models/paymentintent/enum"
)

type Purchase struct {
	Id              string   `json:"id"`
	CustomerId      string   `json:"customer_id"`
	CreatedAt       int64    `json:"created_at"`
	ModifiedAt      int64    `json:"modified_at"`
	SubscriptionIds []string `json:"subscription_ids"`
	InvoiceIds      []string `json:"invoice_ids"`
	Object          string   `json:"object"`
}
type CreateRequestParams struct {
	PurchaseItems       []*CreatePurchaseItemParams      `json:"purchase_items,omitempty"`
	ItemTiers           []*CreateItemTierParams          `json:"item_tiers,omitempty"`
	ShippingAddresses   []*CreateShippingAddressParams   `json:"shipping_addresses,omitempty"`
	Discounts           []*CreateDiscountParams          `json:"discounts,omitempty"`
	SubscriptionInfo    []*CreateSubscriptionInfoParams  `json:"subscription_info,omitempty"`
	ContractTerms       []*CreateContractTermParams      `json:"contract_terms,omitempty"`
	InvoiceInfo         *CreateInvoiceInfoParams         `json:"invoice_info,omitempty"`
	PaymentSchedule     *CreatePaymentScheduleParams     `json:"payment_schedule,omitempty"`
	StatementDescriptor *CreateStatementDescriptorParams `json:"statement_descriptor,omitempty"`
	CustomerId          string                           `json:"customer_id"`
	PaymentSourceId     string                           `json:"payment_source_id,omitempty"`
	PaymentIntent       *CreatePaymentIntentParams       `json:"payment_intent,omitempty"`
}
type CreatePurchaseItemParams struct {
	Index               *int32 `json:"index"`
	ItemPriceId         string `json:"item_price_id"`
	Quantity            *int32 `json:"quantity,omitempty"`
	UnitAmount          *int64 `json:"unit_amount,omitempty"`
	UnitAmountInDecimal string `json:"unit_amount_in_decimal,omitempty"`
	QuantityInDecimal   string `json:"quantity_in_decimal,omitempty"`
}
type CreateItemTierParams struct {
	Index                 *int32 `json:"index"`
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int64 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
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
	State            string                `json:"state,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Country          string                `json:"country,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateDiscountParams struct {
	Index         *int32   `json:"index,omitempty"`
	CouponId      string   `json:"coupon_id,omitempty"`
	Percentage    *float64 `json:"percentage,omitempty"`
	Amount        *int64   `json:"amount,omitempty"`
	IncludedInMrr *bool    `json:"included_in_mrr,omitempty"`
}
type CreateSubscriptionInfoParams struct {
	Index                             *int32                 `json:"index"`
	SubscriptionId                    string                 `json:"subscription_id,omitempty"`
	BillingCycles                     *int32                 `json:"billing_cycles,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                 `json:"contract_term_billing_cycle_on_renewal,omitempty"`
	MetaData                          map[string]interface{} `json:"meta_data,omitempty"`
}
type CreateContractTermParams struct {
	Index                    *int32                           `json:"index"`
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type CreateInvoiceInfoParams struct {
	PoNumber string `json:"po_number,omitempty"`
	Notes    string `json:"notes,omitempty"`
}
type CreatePaymentScheduleParams struct {
	SchemeId string `json:"scheme_id,omitempty"`
	Amount   *int64 `json:"amount,omitempty"`
}
type CreateStatementDescriptorParams struct {
	Descriptor string `json:"descriptor,omitempty"`
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
type EstimateRequestParams struct {
	PurchaseItems     []*EstimatePurchaseItemParams     `json:"purchase_items,omitempty"`
	ItemTiers         []*EstimateItemTierParams         `json:"item_tiers,omitempty"`
	ShippingAddresses []*EstimateShippingAddressParams  `json:"shipping_addresses,omitempty"`
	Discounts         []*EstimateDiscountParams         `json:"discounts,omitempty"`
	SubscriptionInfo  []*EstimateSubscriptionInfoParams `json:"subscription_info,omitempty"`
	ContractTerms     []*EstimateContractTermParams     `json:"contract_terms,omitempty"`
	Customer          *EstimateCustomerParams           `json:"customer,omitempty"`
	BillingAddress    *EstimateBillingAddressParams     `json:"billing_address,omitempty"`
	ClientProfileId   string                            `json:"client_profile_id,omitempty"`
	CustomerId        string                            `json:"customer_id,omitempty"`
}
type EstimatePurchaseItemParams struct {
	Index               *int32 `json:"index"`
	ItemPriceId         string `json:"item_price_id"`
	Quantity            *int32 `json:"quantity,omitempty"`
	UnitAmount          *int64 `json:"unit_amount,omitempty"`
	UnitAmountInDecimal string `json:"unit_amount_in_decimal,omitempty"`
	QuantityInDecimal   string `json:"quantity_in_decimal,omitempty"`
}
type EstimateItemTierParams struct {
	Index                 *int32 `json:"index"`
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int64 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type EstimateShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	State            string                `json:"state,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Country          string                `json:"country,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type EstimateDiscountParams struct {
	Index         *int32   `json:"index,omitempty"`
	CouponId      string   `json:"coupon_id,omitempty"`
	Percentage    *float64 `json:"percentage,omitempty"`
	Amount        *int64   `json:"amount,omitempty"`
	IncludedInMrr *bool    `json:"included_in_mrr,omitempty"`
}
type EstimateSubscriptionInfoParams struct {
	Index                             *int32 `json:"index"`
	SubscriptionId                    string `json:"subscription_id,omitempty"`
	BillingCycles                     *int32 `json:"billing_cycles,omitempty"`
	ContractTermBillingCycleOnRenewal *int32 `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type EstimateContractTermParams struct {
	Index                    *int32                           `json:"index"`
	ActionAtTermEnd          contractTermEnum.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
}
type EstimateCustomerParams struct {
	VatNumber        string                   `json:"vat_number,omitempty"`
	VatNumberPrefix  string                   `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool                    `json:"registered_for_gst,omitempty"`
	Taxability       enum.Taxability          `json:"taxability,omitempty"`
	EntityCode       enum.EntityCode          `json:"entity_code,omitempty"`
	ExemptNumber     string                   `json:"exempt_number,omitempty"`
	ExemptionDetails []map[string]interface{} `json:"exemption_details,omitempty"`
	CustomerType     enum.CustomerType        `json:"customer_type,omitempty"`
}
type EstimateBillingAddressParams struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
