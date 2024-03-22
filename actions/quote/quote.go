package quote

import (
	"fmt"

	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/quote"
)

func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/quotes/%v", chargebee.IDEscape(id)), nil)
}
func CreateSubForCustomerQuote(id string, params *quote.CreateSubForCustomerQuoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/create_subscription_quote", chargebee.IDEscape(id)), params)
}
func EditCreateSubForCustomerQuote(id string, params *quote.EditCreateSubForCustomerQuoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/edit_create_subscription_quote", chargebee.IDEscape(id)), params)
}
func UpdateSubscriptionQuote(params *quote.UpdateSubscriptionQuoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/update_subscription_quote"), params)
}
func EditUpdateSubscriptionQuote(id string, params *quote.EditUpdateSubscriptionQuoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/edit_update_subscription_quote", chargebee.IDEscape(id)), params)
}
func CreateForOnetimeCharges(params *quote.CreateForOnetimeChargesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/create_for_onetime_charges"), params)
}
func EditOneTimeQuote(id string, params *quote.EditOneTimeQuoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/edit_one_time_quote", chargebee.IDEscape(id)), params)
}
func CreateSubItemsForCustomerQuote(id string, params *quote.CreateSubItemsForCustomerQuoteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/create_subscription_quote_for_items", chargebee.IDEscape(id)), params)
}
func EditCreateSubCustomerQuoteForItems(id string, params *quote.EditCreateSubCustomerQuoteForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/edit_create_subscription_quote_for_items", chargebee.IDEscape(id)), params)
}
func UpdateSubscriptionQuoteForItems(params *quote.UpdateSubscriptionQuoteForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/update_subscription_quote_for_items"), params)
}
func EditUpdateSubscriptionQuoteForItems(id string, params *quote.EditUpdateSubscriptionQuoteForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/edit_update_subscription_quote_for_items", chargebee.IDEscape(id)), params)
}
func CreateForChargeItemsAndCharges(params *quote.CreateForChargeItemsAndChargesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/create_for_charge_items_and_charges"), params)
}
func EditForChargeItemsAndCharges(id string, params *quote.EditForChargeItemsAndChargesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/edit_for_charge_items_and_charges", chargebee.IDEscape(id)), params)
}
func List(params *quote.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/quotes"), params)
}
func QuoteLineGroupsForQuote(id string, params *quote.QuoteLineGroupsForQuoteRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/quotes/%v/quote_line_groups", chargebee.IDEscape(id)), params)
}
func Convert(id string, params *quote.ConvertRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/convert", chargebee.IDEscape(id)), params)
}
func UpdateStatus(id string, params *quote.UpdateStatusRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/update_status", chargebee.IDEscape(id)), params)
}
func ExtendExpiryDate(id string, params *quote.ExtendExpiryDateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/extend_expiry_date", chargebee.IDEscape(id)), params)
}
func Delete(id string, params *quote.DeleteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/delete", chargebee.IDEscape(id)), params)
}
func Pdf(id string, params *quote.PdfRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/pdf", chargebee.IDEscape(id)), params)
}
