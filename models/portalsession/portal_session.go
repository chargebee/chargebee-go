package portalsession

import (
	portalSessionEnum "github.com/chargebee/chargebee-go/models/portalsession/enum"
)

type PortalSession struct {
	Id              string                   `json:"id"`
	Token           string                   `json:"token"`
	AccessUrl       string                   `json:"access_url"`
	RedirectUrl     string                   `json:"redirect_url"`
	Status          portalSessionEnum.Status `json:"status"`
	CreatedAt       int64                    `json:"created_at"`
	ExpiresAt       int64                    `json:"expires_at"`
	CustomerId      string                   `json:"customer_id"`
	LoginAt         int64                    `json:"login_at"`
	LogoutAt        int64                    `json:"logout_at"`
	LoginIpaddress  string                   `json:"login_ipaddress"`
	LogoutIpaddress string                   `json:"logout_ipaddress"`
	LinkedCustomers []*LinkedCustomer        `json:"linked_customers"`
	Object          string                   `json:"object"`
}
type LinkedCustomer struct {
	CustomerId            string `json:"customer_id"`
	Email                 string `json:"email"`
	HasBillingAddress     bool   `json:"has_billing_address"`
	HasPaymentMethod      bool   `json:"has_payment_method"`
	HasActiveSubscription bool   `json:"has_active_subscription"`
	Object                string `json:"object"`
}
type CreateRequestParams struct {
	Customer    *CreateCustomerParams `json:"customer,omitempty"`
	RedirectUrl string                `json:"redirect_url,omitempty"`
	ForwardUrl  string                `json:"forward_url,omitempty"`
}
type CreateCustomerParams struct {
	Id string `json:"id"`
}
type ActivateRequestParams struct {
	Token string `json:"token"`
}
