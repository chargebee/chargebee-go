package subscriptionestimate

import (
	"github.com/chargebee/chargebee-go/enum"
	subscriptionEstimateEnum "github.com/chargebee/chargebee-go/models/subscriptionestimate/enum"
)

type SubscriptionEstimate struct {
	Id              string                          `json:"id"`
	CurrencyCode    string                          `json:"currency_code"`
	Status          subscriptionEstimateEnum.Status `json:"status"`
	NextBillingAt   int64                           `json:"next_billing_at"`
	PauseDate       int64                           `json:"pause_date"`
	ResumeDate      int64                           `json:"resume_date"`
	ShippingAddress *ShippingAddress                `json:"shipping_address"`
	Object          string                          `json:"object"`
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
