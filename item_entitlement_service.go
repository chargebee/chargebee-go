package chargebee

import (
	"fmt"
	"net/url"
)

type ItemEntitlementService struct {
	config *ClientConfig
}

func (s *ItemEntitlementService) ItemEntitlementsForItem(id string, req *ItemEntitlementItemEntitlementsForItemRequest) (*ItemEntitlementItemEntitlementsForItemResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/items/%v/item_entitlements", url.PathEscape(id))
	req.isListRequest = true
	return send[*ItemEntitlementItemEntitlementsForItemResponse](req, s.config)
}

func (s *ItemEntitlementService) ItemEntitlementsForFeature(id string, req *ItemEntitlementItemEntitlementsForFeatureRequest) (*ItemEntitlementItemEntitlementsForFeatureResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/features/%v/item_entitlements", url.PathEscape(id))
	req.isListRequest = true
	return send[*ItemEntitlementItemEntitlementsForFeatureResponse](req, s.config)
}

func (s *ItemEntitlementService) AddItemEntitlements(id string, req *ItemEntitlementAddItemEntitlementsRequest) (*ItemEntitlementAddItemEntitlementsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/features/%v/item_entitlements", url.PathEscape(id))
	req.isIdempotent = true
	return send[*ItemEntitlementAddItemEntitlementsResponse](req, s.config)
}

func (s *ItemEntitlementService) UpsertOrRemoveItemEntitlementsForItem(id string, req *ItemEntitlementUpsertOrRemoveItemEntitlementsForItemRequest) (*ItemEntitlementUpsertOrRemoveItemEntitlementsForItemResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/items/%v/item_entitlements", url.PathEscape(id))
	req.isIdempotent = true
	return send[*ItemEntitlementUpsertOrRemoveItemEntitlementsForItemResponse](req, s.config)
}
