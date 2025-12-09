package chargebee

import (
	"fmt"
	"net/url"
)

type ItemEntitlementService struct {
	transport *Transport
}

func (s *ItemEntitlementService) ItemEntitlementsForItem(id string, req *ItemEntitlementsForItemRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/items/%v/item_entitlements", url.PathEscape(id)), req)
}

func (s *ItemEntitlementService) ItemEntitlementsForFeature(id string, req *ItemEntitlementsForFeatureRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/features/%v/item_entitlements", url.PathEscape(id)), req)
}

func (s *ItemEntitlementService) AddItemEntitlements(id string, req *AddItemEntitlementsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/features/%v/item_entitlements", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *ItemEntitlementService) UpsertOrRemoveItemEntitlementsForItem(id string, req *UpsertOrRemoveItemEntitlementsForItemRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/items/%v/item_entitlements", url.PathEscape(id)), req).SetIdempotency(true)
}
