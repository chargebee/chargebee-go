package customerentitlement

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/customerentitlement"
	"net/url"
)

func EntitlementsForCustomer(id string, params *customerentitlement.EntitlementsForCustomerRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/customer_entitlements", url.PathEscape(id)), params)
}
