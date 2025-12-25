package chargebee

type OrderStatus string

const (
	OrderStatusNew                OrderStatus = "new"
	OrderStatusProcessing         OrderStatus = "processing"
	OrderStatusComplete           OrderStatus = "complete"
	OrderStatusCancelled          OrderStatus = "cancelled"
	OrderStatusVoided             OrderStatus = "voided"
	OrderStatusQueued             OrderStatus = "queued"
	OrderStatusAwaitingShipment   OrderStatus = "awaiting_shipment"
	OrderStatusOnHold             OrderStatus = "on_hold"
	OrderStatusDelivered          OrderStatus = "delivered"
	OrderStatusShipped            OrderStatus = "shipped"
	OrderStatusPartiallyDelivered OrderStatus = "partially_delivered"
	OrderStatusReturned           OrderStatus = "returned"
)

type OrderCancellationReason string

const (
	OrderCancellationReasonShippingCutOffPassed   OrderCancellationReason = "shipping_cut_off_passed"
	OrderCancellationReasonProductUnsatisfactory  OrderCancellationReason = "product_unsatisfactory"
	OrderCancellationReasonThirdPartyCancellation OrderCancellationReason = "third_party_cancellation"
	OrderCancellationReasonProductNotRequired     OrderCancellationReason = "product_not_required"
	OrderCancellationReasonDeliveryDateMissed     OrderCancellationReason = "delivery_date_missed"
	OrderCancellationReasonAlternativeFound       OrderCancellationReason = "alternative_found"
	OrderCancellationReasonInvoiceWrittenOff      OrderCancellationReason = "invoice_written_off"
	OrderCancellationReasonInvoiceVoided          OrderCancellationReason = "invoice_voided"
	OrderCancellationReasonFraudulentTransaction  OrderCancellationReason = "fraudulent_transaction"
	OrderCancellationReasonPaymentDeclined        OrderCancellationReason = "payment_declined"
	OrderCancellationReasonSubscriptionCancelled  OrderCancellationReason = "subscription_cancelled"
	OrderCancellationReasonProductNotAvailable    OrderCancellationReason = "product_not_available"
	OrderCancellationReasonOthers                 OrderCancellationReason = "others"
	OrderCancellationReasonOrderResent            OrderCancellationReason = "order_resent"
)

type OrderPaymentStatus string

const (
	OrderPaymentStatusNotPaid OrderPaymentStatus = "not_paid"
	OrderPaymentStatusPaid    OrderPaymentStatus = "paid"
)

type OrderOrderType string

const (
	OrderOrderTypeManual          OrderOrderType = "manual"
	OrderOrderTypeSystemGenerated OrderOrderType = "system_generated"
)

type OrderPriceType string

const (
	OrderPriceTypeTaxExclusive OrderPriceType = "tax_exclusive"
	OrderPriceTypeTaxInclusive OrderPriceType = "tax_inclusive"
)

type OrderResentStatus string

const (
	OrderResentStatusFullyResent     OrderResentStatus = "fully_resent"
	OrderResentStatusPartiallyResent OrderResentStatus = "partially_resent"
)

type OrderOrderLineItemStatus string

const (
	OrderOrderLineItemStatusQueued             OrderOrderLineItemStatus = "queued"
	OrderOrderLineItemStatusAwaitingShipment   OrderOrderLineItemStatus = "awaiting_shipment"
	OrderOrderLineItemStatusOnHold             OrderOrderLineItemStatus = "on_hold"
	OrderOrderLineItemStatusDelivered          OrderOrderLineItemStatus = "delivered"
	OrderOrderLineItemStatusShipped            OrderOrderLineItemStatus = "shipped"
	OrderOrderLineItemStatusPartiallyDelivered OrderOrderLineItemStatus = "partially_delivered"
	OrderOrderLineItemStatusReturned           OrderOrderLineItemStatus = "returned"
	OrderOrderLineItemStatusCancelled          OrderOrderLineItemStatus = "cancelled"
)

type OrderOrderLineItemEntityType string

const (
	OrderOrderLineItemEntityTypeAdhoc           OrderOrderLineItemEntityType = "adhoc"
	OrderOrderLineItemEntityTypePlanItemPrice   OrderOrderLineItemEntityType = "plan_item_price"
	OrderOrderLineItemEntityTypeAddonItemPrice  OrderOrderLineItemEntityType = "addon_item_price"
	OrderOrderLineItemEntityTypeChargeItemPrice OrderOrderLineItemEntityType = "charge_item_price"
	OrderOrderLineItemEntityTypePlanSetup       OrderOrderLineItemEntityType = "plan_setup"
	OrderOrderLineItemEntityTypePlan            OrderOrderLineItemEntityType = "plan"
	OrderOrderLineItemEntityTypeAddon           OrderOrderLineItemEntityType = "addon"
)

type OrderShippingAddressValidationStatus string

const (
	OrderShippingAddressValidationStatusNotValidated   OrderShippingAddressValidationStatus = "not_validated"
	OrderShippingAddressValidationStatusValid          OrderShippingAddressValidationStatus = "valid"
	OrderShippingAddressValidationStatusPartiallyValid OrderShippingAddressValidationStatus = "partially_valid"
	OrderShippingAddressValidationStatusInvalid        OrderShippingAddressValidationStatus = "invalid"
)

type OrderBillingAddressValidationStatus string

const (
	OrderBillingAddressValidationStatusNotValidated   OrderBillingAddressValidationStatus = "not_validated"
	OrderBillingAddressValidationStatusValid          OrderBillingAddressValidationStatus = "valid"
	OrderBillingAddressValidationStatusPartiallyValid OrderBillingAddressValidationStatus = "partially_valid"
	OrderBillingAddressValidationStatusInvalid        OrderBillingAddressValidationStatus = "invalid"
)

type OrderLineItemTaxTaxJurisType string

const (
	OrderLineItemTaxTaxJurisTypeCountry        OrderLineItemTaxTaxJurisType = "country"
	OrderLineItemTaxTaxJurisTypeFederal        OrderLineItemTaxTaxJurisType = "federal"
	OrderLineItemTaxTaxJurisTypeState          OrderLineItemTaxTaxJurisType = "state"
	OrderLineItemTaxTaxJurisTypeCounty         OrderLineItemTaxTaxJurisType = "county"
	OrderLineItemTaxTaxJurisTypeCity           OrderLineItemTaxTaxJurisType = "city"
	OrderLineItemTaxTaxJurisTypeSpecial        OrderLineItemTaxTaxJurisType = "special"
	OrderLineItemTaxTaxJurisTypeUnincorporated OrderLineItemTaxTaxJurisType = "unincorporated"
	OrderLineItemTaxTaxJurisTypeOther          OrderLineItemTaxTaxJurisType = "other"
)

type OrderLineItemDiscountDiscountType string

const (
	OrderLineItemDiscountDiscountTypeItemLevelCoupon       OrderLineItemDiscountDiscountType = "item_level_coupon"
	OrderLineItemDiscountDiscountTypeDocumentLevelCoupon   OrderLineItemDiscountDiscountType = "document_level_coupon"
	OrderLineItemDiscountDiscountTypePromotionalCredits    OrderLineItemDiscountDiscountType = "promotional_credits"
	OrderLineItemDiscountDiscountTypeProratedCredits       OrderLineItemDiscountDiscountType = "prorated_credits"
	OrderLineItemDiscountDiscountTypeCustomDiscount        OrderLineItemDiscountDiscountType = "custom_discount"
	OrderLineItemDiscountDiscountTypeItemLevelDiscount     OrderLineItemDiscountDiscountType = "item_level_discount"
	OrderLineItemDiscountDiscountTypeDocumentLevelDiscount OrderLineItemDiscountDiscountType = "document_level_discount"
)

type OrderOrderLineItemLinkedCreditType string

const (
	OrderOrderLineItemLinkedCreditTypeAdjustment OrderOrderLineItemLinkedCreditType = "adjustment"
	OrderOrderLineItemLinkedCreditTypeRefundable OrderOrderLineItemLinkedCreditType = "refundable"
	OrderOrderLineItemLinkedCreditTypeStore      OrderOrderLineItemLinkedCreditType = "store"
)

type OrderOrderLineItemLinkedCreditStatus string

const (
	OrderOrderLineItemLinkedCreditStatusAdjusted  OrderOrderLineItemLinkedCreditStatus = "adjusted"
	OrderOrderLineItemLinkedCreditStatusRefunded  OrderOrderLineItemLinkedCreditStatus = "refunded"
	OrderOrderLineItemLinkedCreditStatusRefundDue OrderOrderLineItemLinkedCreditStatus = "refund_due"
	OrderOrderLineItemLinkedCreditStatusVoided    OrderOrderLineItemLinkedCreditStatus = "voided"
)

// just struct
type Order struct {
	Id                      string                   `json:"id"`
	DocumentNumber          string                   `json:"document_number"`
	InvoiceId               string                   `json:"invoice_id"`
	SubscriptionId          string                   `json:"subscription_id"`
	CustomerId              string                   `json:"customer_id"`
	Status                  OrderStatus              `json:"status"`
	CancellationReason      OrderCancellationReason  `json:"cancellation_reason"`
	PaymentStatus           OrderPaymentStatus       `json:"payment_status"`
	OrderType               OrderOrderType           `json:"order_type"`
	PriceType               OrderPriceType           `json:"price_type"`
	ReferenceId             string                   `json:"reference_id"`
	FulfillmentStatus       string                   `json:"fulfillment_status"`
	OrderDate               int64                    `json:"order_date"`
	ShippingDate            int64                    `json:"shipping_date"`
	Note                    string                   `json:"note"`
	TrackingId              string                   `json:"tracking_id"`
	TrackingUrl             string                   `json:"tracking_url"`
	BatchId                 string                   `json:"batch_id"`
	CreatedBy               string                   `json:"created_by"`
	ShipmentCarrier         string                   `json:"shipment_carrier"`
	InvoiceRoundOffAmount   int64                    `json:"invoice_round_off_amount"`
	Tax                     int64                    `json:"tax"`
	AmountPaid              int64                    `json:"amount_paid"`
	AmountAdjusted          int64                    `json:"amount_adjusted"`
	RefundableCreditsIssued int64                    `json:"refundable_credits_issued"`
	RefundableCredits       int64                    `json:"refundable_credits"`
	RoundingAdjustement     int64                    `json:"rounding_adjustement"`
	PaidOn                  int64                    `json:"paid_on"`
	ShippingCutOffDate      int64                    `json:"shipping_cut_off_date"`
	CreatedAt               int64                    `json:"created_at"`
	StatusUpdateAt          int64                    `json:"status_update_at"`
	DeliveredAt             int64                    `json:"delivered_at"`
	ShippedAt               int64                    `json:"shipped_at"`
	ResourceVersion         int64                    `json:"resource_version"`
	UpdatedAt               int64                    `json:"updated_at"`
	CancelledAt             int64                    `json:"cancelled_at"`
	ResentStatus            OrderResentStatus        `json:"resent_status"`
	IsResent                bool                     `json:"is_resent"`
	OriginalOrderId         string                   `json:"original_order_id"`
	OrderLineItems          []*OrderOrderLineItem    `json:"order_line_items"`
	ShippingAddress         *OrderShippingAddress    `json:"shipping_address"`
	BillingAddress          *OrderBillingAddress     `json:"billing_address"`
	Discount                int64                    `json:"discount"`
	SubTotal                int64                    `json:"sub_total"`
	Total                   int64                    `json:"total"`
	LineItemTaxes           []*OrderLineItemTax      `json:"line_item_taxes"`
	LineItemDiscounts       []*OrderLineItemDiscount `json:"line_item_discounts"`
	LinkedCreditNotes       []*OrderLinkedCreditNote `json:"linked_credit_notes"`
	Deleted                 bool                     `json:"deleted"`
	CurrencyCode            string                   `json:"currency_code"`
	IsGifted                bool                     `json:"is_gifted"`
	GiftNote                string                   `json:"gift_note"`
	GiftId                  string                   `json:"gift_id"`
	ResendReason            string                   `json:"resend_reason"`
	ResentOrders            []*OrderResentOrder      `json:"resent_orders"`
	BusinessEntityId        string                   `json:"business_entity_id"`
	Object                  string                   `json:"object"`
}

// sub resources
type OrderOrderLineItem struct {
	Id                      string                  `json:"id"`
	InvoiceId               string                  `json:"invoice_id"`
	InvoiceLineItemId       string                  `json:"invoice_line_item_id"`
	UnitPrice               int64                   `json:"unit_price"`
	Description             string                  `json:"description"`
	Amount                  int64                   `json:"amount"`
	FulfillmentQuantity     int32                   `json:"fulfillment_quantity"`
	FulfillmentAmount       int64                   `json:"fulfillment_amount"`
	TaxAmount               int64                   `json:"tax_amount"`
	AmountPaid              int64                   `json:"amount_paid"`
	AmountAdjusted          int64                   `json:"amount_adjusted"`
	RefundableCreditsIssued int64                   `json:"refundable_credits_issued"`
	RefundableCredits       int64                   `json:"refundable_credits"`
	IsShippable             bool                    `json:"is_shippable"`
	Sku                     string                  `json:"sku"`
	Status                  OrderLineItemStatus     `json:"status"`
	EntityType              OrderLineItemEntityType `json:"entity_type"`
	ItemLevelDiscountAmount int64                   `json:"item_level_discount_amount"`
	DiscountAmount          int64                   `json:"discount_amount"`
	EntityId                string                  `json:"entity_id"`
	Object                  string                  `json:"object"`
}
type OrderShippingAddress struct {
	FirstName        string                               `json:"first_name"`
	LastName         string                               `json:"last_name"`
	Email            string                               `json:"email"`
	Company          string                               `json:"company"`
	Phone            string                               `json:"phone"`
	Line1            string                               `json:"line1"`
	Line2            string                               `json:"line2"`
	Line3            string                               `json:"line3"`
	City             string                               `json:"city"`
	StateCode        string                               `json:"state_code"`
	State            string                               `json:"state"`
	Country          string                               `json:"country"`
	Zip              string                               `json:"zip"`
	ValidationStatus OrderShippingAddressValidationStatus `json:"validation_status"`
	Object           string                               `json:"object"`
}
type OrderBillingAddress struct {
	FirstName        string                              `json:"first_name"`
	LastName         string                              `json:"last_name"`
	Email            string                              `json:"email"`
	Company          string                              `json:"company"`
	Phone            string                              `json:"phone"`
	Line1            string                              `json:"line1"`
	Line2            string                              `json:"line2"`
	Line3            string                              `json:"line3"`
	City             string                              `json:"city"`
	StateCode        string                              `json:"state_code"`
	State            string                              `json:"state"`
	Country          string                              `json:"country"`
	Zip              string                              `json:"zip"`
	ValidationStatus OrderBillingAddressValidationStatus `json:"validation_status"`
	Object           string                              `json:"object"`
}
type OrderLineItemTax struct {
	LineItemId               string                       `json:"line_item_id"`
	TaxName                  string                       `json:"tax_name"`
	TaxRate                  float64                      `json:"tax_rate"`
	DateTo                   int64                        `json:"date_to"`
	DateFrom                 int64                        `json:"date_from"`
	ProratedTaxableAmount    float64                      `json:"prorated_taxable_amount"`
	IsPartialTaxApplied      bool                         `json:"is_partial_tax_applied"`
	IsNonComplianceTax       bool                         `json:"is_non_compliance_tax"`
	TaxableAmount            int64                        `json:"taxable_amount"`
	TaxAmount                int64                        `json:"tax_amount"`
	TaxJurisType             OrderLineItemTaxTaxJurisType `json:"tax_juris_type"`
	TaxJurisName             string                       `json:"tax_juris_name"`
	TaxJurisCode             string                       `json:"tax_juris_code"`
	TaxAmountInLocalCurrency int64                        `json:"tax_amount_in_local_currency"`
	LocalCurrencyCode        string                       `json:"local_currency_code"`
	Object                   string                       `json:"object"`
}
type OrderLineItemDiscount struct {
	LineItemId     string                       `json:"line_item_id"`
	DiscountType   LineItemDiscountDiscountType `json:"discount_type"`
	CouponId       string                       `json:"coupon_id"`
	EntityId       string                       `json:"entity_id"`
	DiscountAmount int64                        `json:"discount_amount"`
	Object         string                       `json:"object"`
}
type OrderLinkedCreditNote struct {
	Amount         int64                  `json:"amount"`
	Type           LinkedCreditNoteType   `json:"type"`
	Id             string                 `json:"id"`
	Status         LinkedCreditNoteStatus `json:"status"`
	AmountAdjusted int64                  `json:"amount_adjusted"`
	AmountRefunded int64                  `json:"amount_refunded"`
	Object         string                 `json:"object"`
}
type OrderResentOrder struct {
	OrderId string `json:"order_id"`
	Reason  string `json:"reason"`
	Amount  int64  `json:"amount"`
	Object  string `json:"object"`
}

// operations
// input params
type OrderCreateRequest struct {
	Id                string      `json:"id,omitempty"`
	InvoiceId         string      `json:"invoice_id"`
	Status            OrderStatus `json:"status,omitempty"`
	ReferenceId       string      `json:"reference_id,omitempty"`
	FulfillmentStatus string      `json:"fulfillment_status,omitempty"`
	Note              string      `json:"note,omitempty"`
	TrackingId        string      `json:"tracking_id,omitempty"`
	TrackingUrl       string      `json:"tracking_url,omitempty"`
	BatchId           string      `json:"batch_id,omitempty"`
	apiRequest        `json:"-" form:"-"`
}

func (r *OrderCreateRequest) payload() any { return r }

type OrderUpdateRequest struct {
	ReferenceId        string                      `json:"reference_id,omitempty"`
	BatchId            string                      `json:"batch_id,omitempty"`
	Note               string                      `json:"note,omitempty"`
	ShippingDate       *int64                      `json:"shipping_date,omitempty"`
	OrderDate          *int64                      `json:"order_date,omitempty"`
	CancelledAt        *int64                      `json:"cancelled_at,omitempty"`
	CancellationReason OrderCancellationReason     `json:"cancellation_reason,omitempty"`
	ShippedAt          *int64                      `json:"shipped_at,omitempty"`
	DeliveredAt        *int64                      `json:"delivered_at,omitempty"`
	OrderLineItems     []*OrderUpdateOrderLineItem `json:"order_line_items,omitempty"`
	TrackingUrl        string                      `json:"tracking_url,omitempty"`
	TrackingId         string                      `json:"tracking_id,omitempty"`
	ShipmentCarrier    string                      `json:"shipment_carrier,omitempty"`
	FulfillmentStatus  string                      `json:"fulfillment_status,omitempty"`
	Status             OrderStatus                 `json:"status,omitempty"`
	ShippingAddress    *OrderUpdateShippingAddress `json:"shipping_address,omitempty"`
	apiRequest         `json:"-" form:"-"`
}

func (r *OrderUpdateRequest) payload() any { return r }

// input sub resource params multi
type OrderUpdateOrderLineItem struct {
	Id     string              `json:"id,omitempty"`
	Status OrderLineItemStatus `json:"status,omitempty"`
	Sku    string              `json:"sku,omitempty"`
}

// input sub resource params single
type OrderUpdateShippingAddress struct {
	FirstName        string           `json:"first_name,omitempty"`
	LastName         string           `json:"last_name,omitempty"`
	Email            string           `json:"email,omitempty"`
	Company          string           `json:"company,omitempty"`
	Phone            string           `json:"phone,omitempty"`
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	State            string           `json:"state,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}
type OrderImportOrderRequest struct {
	Id                      string                           `json:"id,omitempty"`
	DocumentNumber          string                           `json:"document_number,omitempty"`
	InvoiceId               string                           `json:"invoice_id"`
	Status                  OrderStatus                      `json:"status"`
	SubscriptionId          string                           `json:"subscription_id,omitempty"`
	CustomerId              string                           `json:"customer_id,omitempty"`
	CreatedAt               *int64                           `json:"created_at"`
	OrderDate               *int64                           `json:"order_date"`
	ShippingDate            *int64                           `json:"shipping_date"`
	ReferenceId             string                           `json:"reference_id,omitempty"`
	FulfillmentStatus       string                           `json:"fulfillment_status,omitempty"`
	Note                    string                           `json:"note,omitempty"`
	TrackingId              string                           `json:"tracking_id,omitempty"`
	TrackingUrl             string                           `json:"tracking_url,omitempty"`
	BatchId                 string                           `json:"batch_id,omitempty"`
	ShipmentCarrier         string                           `json:"shipment_carrier,omitempty"`
	ShippingCutOffDate      *int64                           `json:"shipping_cut_off_date,omitempty"`
	DeliveredAt             *int64                           `json:"delivered_at,omitempty"`
	ShippedAt               *int64                           `json:"shipped_at,omitempty"`
	CancelledAt             *int64                           `json:"cancelled_at,omitempty"`
	CancellationReason      OrderCancellationReason          `json:"cancellation_reason,omitempty"`
	RefundableCreditsIssued *int64                           `json:"refundable_credits_issued,omitempty"`
	ShippingAddress         *OrderImportOrderShippingAddress `json:"shipping_address,omitempty"`
	BillingAddress          *OrderImportOrderBillingAddress  `json:"billing_address,omitempty"`
	apiRequest              `json:"-" form:"-"`
}

func (r *OrderImportOrderRequest) payload() any { return r }

// input sub resource params single
type OrderImportOrderShippingAddress struct {
	FirstName        string           `json:"first_name,omitempty"`
	LastName         string           `json:"last_name,omitempty"`
	Email            string           `json:"email,omitempty"`
	Company          string           `json:"company,omitempty"`
	Phone            string           `json:"phone,omitempty"`
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	State            string           `json:"state,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type OrderImportOrderBillingAddress struct {
	FirstName        string           `json:"first_name,omitempty"`
	LastName         string           `json:"last_name,omitempty"`
	Email            string           `json:"email,omitempty"`
	Company          string           `json:"company,omitempty"`
	Phone            string           `json:"phone,omitempty"`
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	State            string           `json:"state,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}
type OrderCancelRequest struct {
	CancellationReason OrderCancellationReason `json:"cancellation_reason"`
	CreditNote         *OrderCancelCreditNote  `json:"credit_note,omitempty"`
	CustomerNotes      string                  `json:"customer_notes,omitempty"`
	Comment            string                  `json:"comment,omitempty"`
	CancelledAt        *int64                  `json:"cancelled_at,omitempty"`
	apiRequest         `json:"-" form:"-"`
}

func (r *OrderCancelRequest) payload() any { return r }

// input sub resource params single
type OrderCancelCreditNote struct {
	Total *int64 `json:"total,omitempty"`
}
type OrderCreateRefundableCreditNoteRequest struct {
	CreditNote    *OrderCreateRefundableCreditNoteCreditNote `json:"credit_note,omitempty"`
	CustomerNotes string                                     `json:"customer_notes,omitempty"`
	Comment       string                                     `json:"comment,omitempty"`
	apiRequest    `json:"-" form:"-"`
}

func (r *OrderCreateRefundableCreditNoteRequest) payload() any { return r }

// input sub resource params single
type OrderCreateRefundableCreditNoteCreditNote struct {
	ReasonCode ReasonCode `json:"reason_code"`
	Total      *int64     `json:"total"`
}
type OrderReopenRequest struct {
	VoidCancellationCreditNotes *bool `json:"void_cancellation_credit_notes,omitempty"`
	apiRequest                  `json:"-" form:"-"`
}

func (r *OrderReopenRequest) payload() any { return r }

type OrderListRequest struct {
	Limit                     *int32           `json:"limit,omitempty"`
	Offset                    string           `json:"offset,omitempty"`
	IncludeDeleted            *bool            `json:"include_deleted,omitempty"`
	ExcludeDeletedCreditNotes *bool            `json:"exclude_deleted_credit_notes,omitempty"`
	Id                        *StringFilter    `json:"id,omitempty"`
	InvoiceId                 *StringFilter    `json:"invoice_id,omitempty"`
	SubscriptionId            *StringFilter    `json:"subscription_id,omitempty"`
	Status                    *EnumFilter      `json:"status,omitempty"`
	ShippingDate              *TimestampFilter `json:"shipping_date,omitempty"`
	ShippedAt                 *TimestampFilter `json:"shipped_at,omitempty"`
	OrderType                 *EnumFilter      `json:"order_type,omitempty"`
	OrderDate                 *TimestampFilter `json:"order_date,omitempty"`
	PaidOn                    *TimestampFilter `json:"paid_on,omitempty"`
	UpdatedAt                 *TimestampFilter `json:"updated_at,omitempty"`
	CreatedAt                 *TimestampFilter `json:"created_at,omitempty"`
	ResentStatus              *EnumFilter      `json:"resent_status,omitempty"`
	IsResent                  *BooleanFilter   `json:"is_resent,omitempty"`
	OriginalOrderId           *StringFilter    `json:"original_order_id,omitempty"`
	SortBy                    *SortFilter      `json:"sort_by,omitempty"`
	apiRequest                `json:"-" form:"-"`
}

func (r *OrderListRequest) payload() any { return r }

type OrderOrdersForInvoiceRequest struct {
	Limit      *int32 `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *OrderOrdersForInvoiceRequest) payload() any { return r }

type OrderResendRequest struct {
	ShippingDate   *int64                      `json:"shipping_date,omitempty"`
	ResendReason   string                      `json:"resend_reason,omitempty"`
	OrderLineItems []*OrderResendOrderLineItem `json:"order_line_items,omitempty"`
	apiRequest     `json:"-" form:"-"`
}

func (r *OrderResendRequest) payload() any { return r }

// input sub resource params multi
type OrderResendOrderLineItem struct {
	Id                  string `json:"id,omitempty"`
	FulfillmentQuantity *int32 `json:"fulfillment_quantity,omitempty"`
}

// operation response
type OrderCreateResponse struct {
	Order *Order `json:"order,omitempty"`
	apiResponse
}

// operation response
type OrderUpdateResponse struct {
	Order *Order `json:"order,omitempty"`
	apiResponse
}

// operation response
type OrderImportOrderResponse struct {
	Order *Order `json:"order,omitempty"`
	apiResponse
}

// operation response
type OrderAssignOrderNumberResponse struct {
	Order *Order `json:"order,omitempty"`
	apiResponse
}

// operation response
type OrderCancelResponse struct {
	Order *Order `json:"order,omitempty"`
	apiResponse
}

// operation response
type OrderCreateRefundableCreditNoteResponse struct {
	Order *Order `json:"order,omitempty"`
	apiResponse
}

// operation response
type OrderReopenResponse struct {
	Order *Order `json:"order,omitempty"`
	apiResponse
}

// operation response
type OrderRetrieveResponse struct {
	Order *Order `json:"order,omitempty"`
	apiResponse
}

// operation response
type OrderDeleteResponse struct {
	Order *Order `json:"order,omitempty"`
	apiResponse
}

// operation sub response
type OrderListOrderResponse struct {
	Order *Order `json:"order,omitempty"`
}

// operation response
type OrderListResponse struct {
	List       []*OrderListOrderResponse `json:"list,omitempty"`
	NextOffset string                    `json:"next_offset,omitempty"`
	apiResponse
}

// operation sub response
type OrderOrdersForInvoiceOrderResponse struct {
	Order *Order `json:"order,omitempty"`
}

// operation response
type OrderOrdersForInvoiceResponse struct {
	List       []*OrderOrdersForInvoiceOrderResponse `json:"list,omitempty"`
	NextOffset string                                `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type OrderResendResponse struct {
	Order *Order `json:"order,omitempty"`
	apiResponse
}
