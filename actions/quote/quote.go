package quote

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/quote"
	"net/url"
)

func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/quotes/%v", url.PathEscape(id)), nil).WithTelemetryResource("quote").WithTelemetryOperation("retrieve")
}
func CreateSubForCustomerQuote(id string, params *quote.CreateSubForCustomerQuoteRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/create_subscription_quote", url.PathEscape(id)), params).WithTelemetryResource("quote").WithTelemetryOperation("createSubForCustomerQuote").SetIdempotency(true)
}
func EditCreateSubForCustomerQuote(id string, params *quote.EditCreateSubForCustomerQuoteRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/edit_create_subscription_quote", url.PathEscape(id)), params).WithTelemetryResource("quote").WithTelemetryOperation("editCreateSubForCustomerQuote").SetIdempotency(true)
}
func UpdateSubscriptionQuote(params *quote.UpdateSubscriptionQuoteRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/update_subscription_quote"), params).WithTelemetryResource("quote").WithTelemetryOperation("updateSubscriptionQuote").SetIdempotency(true)
}
func EditUpdateSubscriptionQuote(id string, params *quote.EditUpdateSubscriptionQuoteRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/edit_update_subscription_quote", url.PathEscape(id)), params).WithTelemetryResource("quote").WithTelemetryOperation("editUpdateSubscriptionQuote").SetIdempotency(true)
}
func CreateForOnetimeCharges(params *quote.CreateForOnetimeChargesRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/create_for_onetime_charges"), params).WithTelemetryResource("quote").WithTelemetryOperation("createForOnetimeCharges").SetIdempotency(true)
}
func EditOneTimeQuote(id string, params *quote.EditOneTimeQuoteRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/edit_one_time_quote", url.PathEscape(id)), params).WithTelemetryResource("quote").WithTelemetryOperation("editOneTimeQuote").SetIdempotency(true)
}
func CreateSubItemsForCustomerQuote(id string, params *quote.CreateSubItemsForCustomerQuoteRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/create_subscription_quote_for_items", url.PathEscape(id)), params).WithTelemetryResource("quote").WithTelemetryOperation("createSubItemsForCustomerQuote").SetIdempotency(true)
}
func EditCreateSubCustomerQuoteForItems(id string, params *quote.EditCreateSubCustomerQuoteForItemsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/edit_create_subscription_quote_for_items", url.PathEscape(id)), params).WithTelemetryResource("quote").WithTelemetryOperation("editCreateSubCustomerQuoteForItems").SetIdempotency(true)
}
func UpdateSubscriptionQuoteForItems(params *quote.UpdateSubscriptionQuoteForItemsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/update_subscription_quote_for_items"), params).WithTelemetryResource("quote").WithTelemetryOperation("updateSubscriptionQuoteForItems").SetIdempotency(true)
}
func EditUpdateSubscriptionQuoteForItems(id string, params *quote.EditUpdateSubscriptionQuoteForItemsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/edit_update_subscription_quote_for_items", url.PathEscape(id)), params).WithTelemetryResource("quote").WithTelemetryOperation("editUpdateSubscriptionQuoteForItems").SetIdempotency(true)
}
func CreateForChargeItemsAndCharges(params *quote.CreateForChargeItemsAndChargesRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/create_for_charge_items_and_charges"), params).WithTelemetryResource("quote").WithTelemetryOperation("createForChargeItemsAndCharges").SetIdempotency(true)
}
func EditForChargeItemsAndCharges(id string, params *quote.EditForChargeItemsAndChargesRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/edit_for_charge_items_and_charges", url.PathEscape(id)), params).WithTelemetryResource("quote").WithTelemetryOperation("editForChargeItemsAndCharges").SetIdempotency(true)
}
func List(params *quote.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/quotes"), params).WithTelemetryResource("quote").WithTelemetryOperation("list")
}
func QuoteLineGroupsForQuote(id string, params *quote.QuoteLineGroupsForQuoteRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/quotes/%v/quote_line_groups", url.PathEscape(id)), params).WithTelemetryResource("quote").WithTelemetryOperation("quoteLineGroupsForQuote")
}
func Convert(id string, params *quote.ConvertRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/convert", url.PathEscape(id)), params).WithTelemetryResource("quote").WithTelemetryOperation("convert").SetIdempotency(true)
}
func UpdateStatus(id string, params *quote.UpdateStatusRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/update_status", url.PathEscape(id)), params).WithTelemetryResource("quote").WithTelemetryOperation("updateStatus").SetIdempotency(true)
}
func ExtendExpiryDate(id string, params *quote.ExtendExpiryDateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/extend_expiry_date", url.PathEscape(id)), params).WithTelemetryResource("quote").WithTelemetryOperation("extendExpiryDate").SetIdempotency(true)
}
func Delete(id string, params *quote.DeleteRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/delete", url.PathEscape(id)), params).WithTelemetryResource("quote").WithTelemetryOperation("delete").SetIdempotency(true)
}
func Pdf(id string, params *quote.PdfRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/pdf", url.PathEscape(id)), params).WithTelemetryResource("quote").WithTelemetryOperation("pdf").SetIdempotency(true)
}

// Deprecated: This function is deprecated.
func RetrieveSignature(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/quotes/%v/retrieve_signature", url.PathEscape(id)), nil).WithTelemetryResource("quote").WithTelemetryOperation("retrieveSignature")
}

// Deprecated: This function is deprecated.
func RetrieveSignedPdf(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/retrieve_signed_pdf", url.PathEscape(id)), nil).WithTelemetryResource("quote").WithTelemetryOperation("retrieveSignedPdf").SetIdempotency(true)
}

// Deprecated: This function is deprecated.
func CreateSignature(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/create_signature", url.PathEscape(id)), nil).WithTelemetryResource("quote").WithTelemetryOperation("createSignature").SetIdempotency(true)
}

// Deprecated: This function is deprecated.
func UpdateSignature(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/update_signature", url.PathEscape(id)), nil).WithTelemetryResource("quote").WithTelemetryOperation("updateSignature").SetIdempotency(true)
}

// Deprecated: This function is deprecated.
func UpdateSignatureStatus(id string, params *quote.UpdateSignatureStatusRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/update_signature_status", url.PathEscape(id)), params).WithTelemetryResource("quote").WithTelemetryOperation("updateSignatureStatus").SetIdempotency(true)
}

// Deprecated: This function is deprecated.
func RefreshSignatureLink(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/quotes/%v/refresh_signature_link", url.PathEscape(id)), nil).WithTelemetryResource("quote").WithTelemetryOperation("refreshSignatureLink").SetIdempotency(true)
}
