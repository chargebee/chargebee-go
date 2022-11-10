package order

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	creditNoteEnum "github.com/chargebee/chargebee-go/models/creditnote/enum"
	orderEnum "github.com/chargebee/chargebee-go/models/order/enum"
)

type Order struct {
	Id                      string                       `json:"id"`
	DocumentNumber          string                       `json:"document_number"`
	InvoiceId               string                       `json:"invoice_id"`
	SubscriptionId          string                       `json:"subscription_id"`
	CustomerId              string                       `json:"customer_id"`
	Status                  orderEnum.Status             `json:"status"`
	CancellationReason      orderEnum.CancellationReason `json:"cancellation_reason"`
	PaymentStatus           orderEnum.PaymentStatus      `json:"payment_status"`
	OrderType               orderEnum.OrderType          `json:"order_type"`
	PriceType               enum.PriceType               `json:"price_type"`
	ReferenceId             string                       `json:"reference_id"`
	FulfillmentStatus       string                       `json:"fulfillment_status"`
	OrderDate               int64                        `json:"order_date"`
	ShippingDate            int64                        `json:"shipping_date"`
	Note                    string                       `json:"note"`
	TrackingId              string                       `json:"tracking_id"`
	TrackingUrl             string                       `json:"tracking_url"`
	BatchId                 string                       `json:"batch_id"`
	CreatedBy               string                       `json:"created_by"`
	ShipmentCarrier         string                       `json:"shipment_carrier"`
	InvoiceRoundOffAmount   int32                        `json:"invoice_round_off_amount"`
	Tax                     int32                        `json:"tax"`
	AmountPaid              int32                        `json:"amount_paid"`
	AmountAdjusted          int32                        `json:"amount_adjusted"`
	RefundableCreditsIssued int32                        `json:"refundable_credits_issued"`
	RefundableCredits       int32                        `json:"refundable_credits"`
	RoundingAdjustement     int32                        `json:"rounding_adjustement"`
	PaidOn                  int64                        `json:"paid_on"`
	ShippingCutOffDate      int64                        `json:"shipping_cut_off_date"`
	CreatedAt               int64                        `json:"created_at"`
	StatusUpdateAt          int64                        `json:"status_update_at"`
	DeliveredAt             int64                        `json:"delivered_at"`
	ShippedAt               int64                        `json:"shipped_at"`
	ResourceVersion         int64                        `json:"resource_version"`
	UpdatedAt               int64                        `json:"updated_at"`
	CancelledAt             int64                        `json:"cancelled_at"`
	ResentStatus            orderEnum.ResentStatus       `json:"resent_status"`
	IsResent                bool                         `json:"is_resent"`
	OriginalOrderId         string                       `json:"original_order_id"`
	OrderLineItems          []*OrderLineItem             `json:"order_line_items"`
	ShippingAddress         *ShippingAddress             `json:"shipping_address"`
	BillingAddress          *BillingAddress              `json:"billing_address"`
	Discount                int32                        `json:"discount"`
	SubTotal                int32                        `json:"sub_total"`
	Total                   int32                        `json:"total"`
	LineItemTaxes           []*LineItemTax               `json:"line_item_taxes"`
	LineItemDiscounts       []*LineItemDiscount          `json:"line_item_discounts"`
	LinkedCreditNotes       []*LinkedCreditNote          `json:"linked_credit_notes"`
	Deleted                 bool                         `json:"deleted"`
	CurrencyCode            string                       `json:"currency_code"`
	IsGifted                bool                         `json:"is_gifted"`
	GiftNote                string                       `json:"gift_note"`
	GiftId                  string                       `json:"gift_id"`
	ResendReason            string                       `json:"resend_reason"`
	ResentOrders            []*ResentOrder               `json:"resent_orders"`
	BusinessEntityId        string                       `json:"business_entity_id"`
	Object                  string                       `json:"object"`
}
type OrderLineItem struct {
	Id                      string                            `json:"id"`
	InvoiceId               string                            `json:"invoice_id"`
	InvoiceLineItemId       string                            `json:"invoice_line_item_id"`
	UnitPrice               int32                             `json:"unit_price"`
	Description             string                            `json:"description"`
	Amount                  int32                             `json:"amount"`
	FulfillmentQuantity     int32                             `json:"fulfillment_quantity"`
	FulfillmentAmount       int32                             `json:"fulfillment_amount"`
	TaxAmount               int32                             `json:"tax_amount"`
	AmountPaid              int32                             `json:"amount_paid"`
	AmountAdjusted          int32                             `json:"amount_adjusted"`
	RefundableCreditsIssued int32                             `json:"refundable_credits_issued"`
	RefundableCredits       int32                             `json:"refundable_credits"`
	IsShippable             bool                              `json:"is_shippable"`
	Sku                     string                            `json:"sku"`
	Status                  orderEnum.OrderLineItemStatus     `json:"status"`
	EntityType              orderEnum.OrderLineItemEntityType `json:"entity_type"`
	ItemLevelDiscountAmount int32                             `json:"item_level_discount_amount"`
	DiscountAmount          int32                             `json:"discount_amount"`
	EntityId                string                            `json:"entity_id"`
	Object                  string                            `json:"object"`
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
	Index            int32                 `json:"index"`
	Object           string                `json:"object"`
}
type BillingAddress struct {
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
type LineItemTax struct {
	LineItemId               string            `json:"line_item_id"`
	TaxName                  string            `json:"tax_name"`
	TaxRate                  float64           `json:"tax_rate"`
	IsPartialTaxApplied      bool              `json:"is_partial_tax_applied"`
	IsNonComplianceTax       bool              `json:"is_non_compliance_tax"`
	TaxableAmount            int32             `json:"taxable_amount"`
	TaxAmount                int32             `json:"tax_amount"`
	TaxJurisType             enum.TaxJurisType `json:"tax_juris_type"`
	TaxJurisName             string            `json:"tax_juris_name"`
	TaxJurisCode             string            `json:"tax_juris_code"`
	TaxAmountInLocalCurrency int32             `json:"tax_amount_in_local_currency"`
	LocalCurrencyCode        string            `json:"local_currency_code"`
	Object                   string            `json:"object"`
}
type LineItemDiscount struct {
	LineItemId     string                                 `json:"line_item_id"`
	DiscountType   orderEnum.LineItemDiscountDiscountType `json:"discount_type"`
	CouponId       string                                 `json:"coupon_id"`
	EntityId       string                                 `json:"entity_id"`
	DiscountAmount int32                                  `json:"discount_amount"`
	Object         string                                 `json:"object"`
}
type LinkedCreditNote struct {
	Amount         int32                                     `json:"amount"`
	Type           orderEnum.OrderLineItemLinkedCreditType   `json:"type"`
	Id             string                                    `json:"id"`
	Status         orderEnum.OrderLineItemLinkedCreditStatus `json:"status"`
	AmountAdjusted int32                                     `json:"amount_adjusted"`
	AmountRefunded int32                                     `json:"amount_refunded"`
	Object         string                                    `json:"object"`
}
type ResentOrder struct {
	OrderId string `json:"order_id"`
	Reason  string `json:"reason"`
	Amount  int32  `json:"amount"`
	Object  string `json:"object"`
}
type CreateRequestParams struct {
	Id                string           `json:"id,omitempty"`
	InvoiceId         string           `json:"invoice_id"`
	Status            orderEnum.Status `json:"status,omitempty"`
	ReferenceId       string           `json:"reference_id,omitempty"`
	FulfillmentStatus string           `json:"fulfillment_status,omitempty"`
	Note              string           `json:"note,omitempty"`
	TrackingId        string           `json:"tracking_id,omitempty"`
	TrackingUrl       string           `json:"tracking_url,omitempty"`
	BatchId           string           `json:"batch_id,omitempty"`
}
type UpdateRequestParams struct {
	ReferenceId        string                       `json:"reference_id,omitempty"`
	BatchId            string                       `json:"batch_id,omitempty"`
	Note               string                       `json:"note,omitempty"`
	ShippingDate       *int64                       `json:"shipping_date,omitempty"`
	OrderDate          *int64                       `json:"order_date,omitempty"`
	CancelledAt        *int64                       `json:"cancelled_at,omitempty"`
	CancellationReason orderEnum.CancellationReason `json:"cancellation_reason,omitempty"`
	ShippedAt          *int64                       `json:"shipped_at,omitempty"`
	DeliveredAt        *int64                       `json:"delivered_at,omitempty"`
	OrderLineItems     []*UpdateOrderLineItemParams `json:"order_line_items,omitempty"`
	TrackingUrl        string                       `json:"tracking_url,omitempty"`
	TrackingId         string                       `json:"tracking_id,omitempty"`
	ShipmentCarrier    string                       `json:"shipment_carrier,omitempty"`
	FulfillmentStatus  string                       `json:"fulfillment_status,omitempty"`
	Status             orderEnum.Status             `json:"status,omitempty"`
	ShippingAddress    *UpdateShippingAddressParams `json:"shipping_address,omitempty"`
}
type UpdateOrderLineItemParams struct {
	Id     string                        `json:"id,omitempty"`
	Status orderEnum.OrderLineItemStatus `json:"status,omitempty"`
	Sku    string                        `json:"sku,omitempty"`
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
type ImportOrderRequestParams struct {
	Id                      string                            `json:"id,omitempty"`
	DocumentNumber          string                            `json:"document_number,omitempty"`
	InvoiceId               string                            `json:"invoice_id"`
	Status                  orderEnum.Status                  `json:"status"`
	SubscriptionId          string                            `json:"subscription_id,omitempty"`
	CustomerId              string                            `json:"customer_id,omitempty"`
	CreatedAt               *int64                            `json:"created_at"`
	OrderDate               *int64                            `json:"order_date"`
	ShippingDate            *int64                            `json:"shipping_date"`
	ReferenceId             string                            `json:"reference_id,omitempty"`
	FulfillmentStatus       string                            `json:"fulfillment_status,omitempty"`
	Note                    string                            `json:"note,omitempty"`
	TrackingId              string                            `json:"tracking_id,omitempty"`
	TrackingUrl             string                            `json:"tracking_url,omitempty"`
	BatchId                 string                            `json:"batch_id,omitempty"`
	ShipmentCarrier         string                            `json:"shipment_carrier,omitempty"`
	ShippingCutOffDate      *int64                            `json:"shipping_cut_off_date,omitempty"`
	DeliveredAt             *int64                            `json:"delivered_at,omitempty"`
	ShippedAt               *int64                            `json:"shipped_at,omitempty"`
	CancelledAt             *int64                            `json:"cancelled_at,omitempty"`
	CancellationReason      orderEnum.CancellationReason      `json:"cancellation_reason,omitempty"`
	RefundableCreditsIssued *int32                            `json:"refundable_credits_issued,omitempty"`
	ShippingAddress         *ImportOrderShippingAddressParams `json:"shipping_address,omitempty"`
	BillingAddress          *ImportOrderBillingAddressParams  `json:"billing_address,omitempty"`
}
type ImportOrderShippingAddressParams struct {
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
type ImportOrderBillingAddressParams struct {
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
type CancelRequestParams struct {
	CancellationReason orderEnum.CancellationReason `json:"cancellation_reason"`
	CreditNote         *CancelCreditNoteParams      `json:"credit_note,omitempty"`
	CustomerNotes      string                       `json:"customer_notes,omitempty"`
	Comment            string                       `json:"comment,omitempty"`
	CancelledAt        *int64                       `json:"cancelled_at,omitempty"`
}
type CancelCreditNoteParams struct {
	Total *int32 `json:"total,omitempty"`
}
type CreateRefundableCreditNoteRequestParams struct {
	CreditNote    *CreateRefundableCreditNoteCreditNoteParams `json:"credit_note,omitempty"`
	CustomerNotes string                                      `json:"customer_notes,omitempty"`
	Comment       string                                      `json:"comment,omitempty"`
}
type CreateRefundableCreditNoteCreditNoteParams struct {
	ReasonCode creditNoteEnum.ReasonCode `json:"reason_code"`
	Total      *int32                    `json:"total"`
}
type ReopenRequestParams struct {
	VoidCancellationCreditNotes *bool `json:"void_cancellation_credit_notes,omitempty"`
}
type ListRequestParams struct {
	Limit                     *int32                  `json:"limit,omitempty"`
	Offset                    string                  `json:"offset,omitempty"`
	IncludeDeleted            *bool                   `json:"include_deleted,omitempty"`
	ExcludeDeletedCreditNotes *bool                   `json:"exclude_deleted_credit_notes,omitempty"`
	Id                        *filter.StringFilter    `json:"id,omitempty"`
	InvoiceId                 *filter.StringFilter    `json:"invoice_id,omitempty"`
	SubscriptionId            *filter.StringFilter    `json:"subscription_id,omitempty"`
	Status                    *filter.EnumFilter      `json:"status,omitempty"`
	ShippingDate              *filter.TimestampFilter `json:"shipping_date,omitempty"`
	ShippedAt                 *filter.TimestampFilter `json:"shipped_at,omitempty"`
	OrderType                 *filter.EnumFilter      `json:"order_type,omitempty"`
	OrderDate                 *filter.TimestampFilter `json:"order_date,omitempty"`
	PaidOn                    *filter.TimestampFilter `json:"paid_on,omitempty"`
	UpdatedAt                 *filter.TimestampFilter `json:"updated_at,omitempty"`
	CreatedAt                 *filter.TimestampFilter `json:"created_at,omitempty"`
	ResentStatus              *filter.EnumFilter      `json:"resent_status,omitempty"`
	IsResent                  *filter.BooleanFilter   `json:"is_resent,omitempty"`
	OriginalOrderId           *filter.StringFilter    `json:"original_order_id,omitempty"`
	SortBy                    *filter.SortFilter      `json:"sort_by,omitempty"`
}
type OrdersForInvoiceRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
type ResendRequestParams struct {
	ShippingDate   *int64                       `json:"shipping_date,omitempty"`
	ResendReason   string                       `json:"resend_reason,omitempty"`
	OrderLineItems []*ResendOrderLineItemParams `json:"order_line_items,omitempty"`
}
type ResendOrderLineItemParams struct {
	Id                  string `json:"id,omitempty"`
	FulfillmentQuantity *int32 `json:"fulfillment_quantity,omitempty"`
}
