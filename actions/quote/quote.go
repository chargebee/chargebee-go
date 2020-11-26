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
func EditCreateSubForCustomerQuote(id string, params *quote.EditCreateSubForCustomerQuoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/edit_create_subscription_quote", id), params)
}
func CreateSubItemsForCustomerQuote(id string, params *quote.CreateSubItemsForCustomerQuoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/create_subscription_quote_for_items", id), params)
}
func UpdateSubscriptionQuote(params *quote.UpdateSubscriptionQuoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/update_subscription_quote"), params)
}
func EditUpdateSubscriptionQuote(id string, params *quote.EditUpdateSubscriptionQuoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/edit_update_subscription_quote", id), params)
}
func UpdateSubscriptionQuoteForItems(params *quote.UpdateSubscriptionQuoteForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/update_subscription_quote_for_items"), params)
}
func CreateForOnetimeCharges(params *quote.CreateForOnetimeChargesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/create_for_onetime_charges"), params)
}
func CreateForChargeItemsAndCharges(params *quote.CreateForChargeItemsAndChargesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/create_for_charge_items_and_charges"), params)
}
func EditOneTimeQuote(id string, params *quote.EditOneTimeQuoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/edit_one_time_quote", id), params)
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
func ExtendExpiryDate(id string, params *quote.ExtendExpiryDateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/extend_expiry_date", id), params)
}
func Delete(id string, params *quote.DeleteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/delete", id), params)
}
func Pdf(id string, params *quote.PdfRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/pdf", id), params)
}
