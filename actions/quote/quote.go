package quote

import (
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/quote"
)

func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/quotes/%v", id), nil)
}
func CreateSubForCustomerQuote(id string, params *quote.CreateSubForCustomerQuoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/create_subscription_quote", id), params)
}
func UpdateSubscriptionQuote(params *quote.UpdateSubscriptionQuoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/update_subscription_quote"), params)
}
func CreateForOnetimeCharges(params *quote.CreateForOnetimeChargesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/create_for_onetime_charges"), params)
}
func List(params *quote.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/quotes"), params)
}
func QuoteLineGroupsForQuote(id string, params *quote.QuoteLineGroupsForQuoteRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/quotes/%v/quote_line_groups", id), params)
}
func Convert(id string, params *quote.ConvertRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/convert", id), params)
}
func UpdateStatus(id string, params *quote.UpdateStatusRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/update_status", id), params)
}
func Delete(id string, params *quote.DeleteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/delete", id), params)
}
func Pdf(id string, params *quote.PdfRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/pdf", id), params)
}
