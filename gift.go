package chargebee

type Status string

const (
	StatusScheduled Status = "scheduled"
	StatusUnclaimed Status = "unclaimed"
	StatusClaimed   Status = "claimed"
	StatusCancelled Status = "cancelled"
	StatusExpired   Status = "expired"
)

type Gift struct {
	Id              string          `json:"id"`
	Status          Status          `json:"status"`
	ScheduledAt     int64           `json:"scheduled_at"`
	AutoClaim       bool            `json:"auto_claim"`
	NoExpiry        bool            `json:"no_expiry"`
	ClaimExpiryDate int64           `json:"claim_expiry_date"`
	ResourceVersion int64           `json:"resource_version"`
	UpdatedAt       int64           `json:"updated_at"`
	Gifter          *Gifter         `json:"gifter"`
	GiftReceiver    *GiftReceiver   `json:"gift_receiver"`
	GiftTimelines   []*GiftTimeline `json:"gift_timelines"`
	Object          string          `json:"object"`
}
type Gifter struct {
	CustomerId string `json:"customer_id"`
	InvoiceId  string `json:"invoice_id"`
	Signature  string `json:"signature"`
	Note       string `json:"note"`
	Object     string `json:"object"`
}
type GiftReceiver struct {
	CustomerId     string `json:"customer_id"`
	SubscriptionId string `json:"subscription_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Object         string `json:"object"`
}
type GiftTimeline struct {
	Status     giftEnum.Status `json:"status"`
	OccurredAt int64           `json:"occurred_at"`
	Object     string          `json:"object"`
}
type CreateRequest struct {
	ScheduledAt     *int64                 `json:"scheduled_at,omitempty"`
	AutoClaim       *bool                  `json:"auto_claim,omitempty"`
	NoExpiry        *bool                  `json:"no_expiry,omitempty"`
	ClaimExpiryDate *int64                 `json:"claim_expiry_date,omitempty"`
	Gifter          *CreateGifter          `json:"gifter,omitempty"`
	GiftReceiver    *CreateGiftReceiver    `json:"gift_receiver,omitempty"`
	CouponIds       []string               `json:"coupon_ids,omitempty"`
	PaymentIntent   *CreatePaymentIntent   `json:"payment_intent,omitempty"`
	ShippingAddress *CreateShippingAddress `json:"shipping_address,omitempty"`
	Subscription    *CreateSubscription    `json:"subscription,omitempty"`
	Addons          []*CreateAddon         `json:"addons,omitempty"`
}
type CreateGifter struct {
	CustomerId   string `json:"customer_id"`
	Signature    string `json:"signature"`
	Note         string `json:"note,omitempty"`
	PaymentSrcId string `json:"payment_src_id,omitempty"`
}
type CreateGiftReceiver struct {
	CustomerId string `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
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
type CreateSubscription struct {
	PlanId                string `json:"plan_id"`
	PlanQuantity          *int32 `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal string `json:"plan_quantity_in_decimal,omitempty"`
}
type CreateAddon struct {
	Id                string `json:"id,omitempty"`
	Quantity          *int32 `json:"quantity,omitempty"`
	QuantityInDecimal string `json:"quantity_in_decimal,omitempty"`
}
type CreateForItemsRequest struct {
	ScheduledAt       *int64                            `json:"scheduled_at,omitempty"`
	AutoClaim         *bool                             `json:"auto_claim,omitempty"`
	NoExpiry          *bool                             `json:"no_expiry,omitempty"`
	ClaimExpiryDate   *int64                            `json:"claim_expiry_date,omitempty"`
	Gifter            *CreateForItemsGifter             `json:"gifter,omitempty"`
	GiftReceiver      *CreateForItemsGiftReceiver       `json:"gift_receiver,omitempty"`
	CouponIds         []string                          `json:"coupon_ids,omitempty"`
	PaymentIntent     *CreateForItemsPaymentIntent      `json:"payment_intent,omitempty"`
	ShippingAddress   *CreateForItemsShippingAddress    `json:"shipping_address,omitempty"`
	SubscriptionItems []*CreateForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	MetaData          map[string]interface{}            `json:"meta_data,omitempty"`
	ItemTiers         []*CreateForItemsItemTier         `json:"item_tiers,omitempty"`
}
type CreateForItemsGifter struct {
	CustomerId   string `json:"customer_id"`
	Signature    string `json:"signature"`
	Note         string `json:"note,omitempty"`
	PaymentSrcId string `json:"payment_src_id,omitempty"`
}
type CreateForItemsGiftReceiver struct {
	CustomerId string `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
}
type CreateForItemsPaymentIntent struct {
	Id                    string                          `json:"id,omitempty"`
	GatewayAccountId      string                          `json:"gateway_account_id,omitempty"`
	GwToken               string                          `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntent.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                          `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                          `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}          `json:"additional_information,omitempty"`
}
type CreateForItemsShippingAddress struct {
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
type CreateForItemsSubscriptionItem struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
}
type CreateForItemsItemTier struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int64 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type ListRequest struct {
	Limit        *int32             `json:"limit,omitempty"`
	Offset       string             `json:"offset,omitempty"`
	GiftReceiver *ListGiftReceiver  `json:"gift_receiver,omitempty"`
	Gifter       *ListGifter        `json:"gifter,omitempty"`
	Status       *filter.EnumFilter `json:"status,omitempty"`
}
type ListGiftReceiver struct {
	Email      *filter.StringFilter `json:"email,omitempty"`
	CustomerId *filter.StringFilter `json:"customer_id,omitempty"`
}
type ListGifter struct {
	CustomerId *filter.StringFilter `json:"customer_id,omitempty"`
}
type UpdateGiftRequest struct {
	ScheduledAt *int64 `json:"scheduled_at"`
	Comment     string `json:"comment,omitempty"`
}

type CreateResponse struct {
	Gift         *Gift                      `json:"gift,omitempty"`
	Subscription *subscription.Subscription `json:"subscription,omitempty"`
	Invoice      *invoice.Invoice           `json:"invoice,omitempty"`
}

type CreateForItemsResponse struct {
	Gift         *Gift                      `json:"gift,omitempty"`
	Subscription *subscription.Subscription `json:"subscription,omitempty"`
	Invoice      *invoice.Invoice           `json:"invoice,omitempty"`
}

type RetrieveResponse struct {
	Gift         *Gift                      `json:"gift,omitempty"`
	Subscription *subscription.Subscription `json:"subscription,omitempty"`
}

type ListGiftResponse struct {
	Gift         *Gift                      `json:"gift,omitempty"`
	Subscription *subscription.Subscription `json:"subscription,omitempty"`
}

type ListResponse struct {
	List       []*ListGiftResponse `json:"list,omitempty"`
	NextOffset string              `json:"next_offset,omitempty"`
}

type ClaimResponse struct {
	Gift         *Gift                      `json:"gift,omitempty"`
	Subscription *subscription.Subscription `json:"subscription,omitempty"`
}

type CancelResponse struct {
	Gift         *Gift                      `json:"gift,omitempty"`
	Subscription *subscription.Subscription `json:"subscription,omitempty"`
}

type UpdateGiftResponse struct {
	Gift         *Gift                      `json:"gift,omitempty"`
	Subscription *subscription.Subscription `json:"subscription,omitempty"`
}
