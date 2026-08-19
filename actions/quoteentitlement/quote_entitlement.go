package quoteentitlement

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/quoteentitlement"
	"net/url"
)

func ListQuoteEntitlements(id string, params *quoteentitlement.ListQuoteEntitlementsRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/quotes/%v/quote_entitlements", url.PathEscape(id)), params)
}
