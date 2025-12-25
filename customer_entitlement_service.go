package chargebee

import (
	"fmt"
	"net/url"
)

type CustomerEntitlementService struct {
	config *ClientConfig
}

func (s *CustomerEntitlementService) EntitlementsForCustomer(id string, req *CustomerEntitlementEntitlementsForCustomerRequest) (*CustomerEntitlementEntitlementsForCustomerResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/customers/%v/customer_entitlements", url.PathEscape(id))
	req.isListRequest = true
	return send[*CustomerEntitlementEntitlementsForCustomerResponse](req, s.config)
}
