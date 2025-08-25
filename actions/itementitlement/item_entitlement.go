package itementitlement

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/itementitlement"
	"net/url"
)

func ItemEntitlementsForItem(id string, params *itementitlement.ItemEntitlementsForItemRequestParams) chargebee.ListRequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/items/%v/item_entitlements", url.PathEscape(id)), params)
}
func ItemEntitlementsForFeature(id string, params *itementitlement.ItemEntitlementsForFeatureRequestParams) chargebee.ListRequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/features/%v/item_entitlements", url.PathEscape(id)), params)
}
func AddItemEntitlements(id string, params *itementitlement.AddItemEntitlementsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/features/%v/item_entitlements", url.PathEscape(id)), params).SetIdempotency(true)
}
func UpsertOrRemoveItemEntitlementsForItem(id string, params *itementitlement.UpsertOrRemoveItemEntitlementsForItemRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/items/%v/item_entitlements", url.PathEscape(id)), params).SetIdempotency(true)
}
