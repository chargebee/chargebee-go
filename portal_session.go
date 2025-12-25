package chargebee

type PortalSessionStatus string

const (
	PortalSessionStatusCreated         PortalSessionStatus = "created"
	PortalSessionStatusLoggedIn        PortalSessionStatus = "logged_in"
	PortalSessionStatusLoggedOut       PortalSessionStatus = "logged_out"
	PortalSessionStatusNotYetActivated PortalSessionStatus = "not_yet_activated"
	PortalSessionStatusActivated       PortalSessionStatus = "activated"
)

// just struct
type PortalSession struct {
	Id              string                         `json:"id"`
	Token           string                         `json:"token"`
	AccessUrl       string                         `json:"access_url"`
	RedirectUrl     string                         `json:"redirect_url"`
	Status          PortalSessionStatus            `json:"status"`
	CreatedAt       int64                          `json:"created_at"`
	ExpiresAt       int64                          `json:"expires_at"`
	CustomerId      string                         `json:"customer_id"`
	LoginAt         int64                          `json:"login_at"`
	LogoutAt        int64                          `json:"logout_at"`
	LoginIpaddress  string                         `json:"login_ipaddress"`
	LogoutIpaddress string                         `json:"logout_ipaddress"`
	LinkedCustomers []*PortalSessionLinkedCustomer `json:"linked_customers"`
	Object          string                         `json:"object"`
}

// sub resources
type PortalSessionLinkedCustomer struct {
	CustomerId            string `json:"customer_id"`
	Email                 string `json:"email"`
	HasBillingAddress     bool   `json:"has_billing_address"`
	HasPaymentMethod      bool   `json:"has_payment_method"`
	HasActiveSubscription bool   `json:"has_active_subscription"`
	Object                string `json:"object"`
}

// operations
// input params
type PortalSessionCreateRequest struct {
	Customer    *PortalSessionCreateCustomer `json:"customer,omitempty"`
	RedirectUrl string                       `json:"redirect_url,omitempty"`
	ForwardUrl  string                       `json:"forward_url,omitempty"`
	apiRequest  `json:"-" form:"-"`
}

func (r *PortalSessionCreateRequest) payload() any { return r }

// input sub resource params single
type PortalSessionCreateCustomer struct {
	Id string `json:"id"`
}

type PortalSessionActivateRequest struct {
	Token      string `json:"token"`
	apiRequest `json:"-" form:"-"`
}

func (r *PortalSessionActivateRequest) payload() any { return r }

// operation response
type PortalSessionCreateResponse struct {
	PortalSession *PortalSession `json:"portal_session,omitempty"`
	apiResponse
}

// operation response
type PortalSessionRetrieveResponse struct {
	PortalSession *PortalSession `json:"portal_session,omitempty"`
	apiResponse
}

// operation response
type PortalSessionLogoutResponse struct {
	PortalSession *PortalSession `json:"portal_session,omitempty"`
	apiResponse
}

// operation response
type PortalSessionActivateResponse struct {
	PortalSession *PortalSession `json:"portal_session,omitempty"`
	apiResponse
}
