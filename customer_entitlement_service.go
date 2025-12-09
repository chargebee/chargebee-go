package chargebee

import (
	"fmt"
	"net/url"
)

type CustomerEntitlementService struct {
	transport *Transport
}

func (s *CustomerEntitlementService) EntitlementsForCustomer(id string, req *EntitlementsForCustomerRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/customers/%v/customer_entitlements", url.PathEscape(id)), req)
}
