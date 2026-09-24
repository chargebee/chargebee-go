package einvoice

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/einvoice"
	"net/url"
)

func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/einvoices/%v", url.PathEscape(id)), nil)
}
func ListEinvoices(params *einvoice.ListEinvoicesRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/einvoices"), params)
}
