package unbilledcharge

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"
	"net/url"
)

func CreateUnbilledCharge(params *unbilledcharge.CreateUnbilledChargeRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/unbilled_charges/create"), params).SetIdempotency(true)
}
func Create(params *unbilledcharge.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/unbilled_charges"), params).SetIdempotency(true)
}
func InvoiceUnbilledCharges(params *unbilledcharge.InvoiceUnbilledChargesRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/unbilled_charges/invoice_unbilled_charges"), params).SetIdempotency(true)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/unbilled_charges/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}
func List(params *unbilledcharge.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/unbilled_charges"), params)
}
func InvoiceNowEstimate(params *unbilledcharge.InvoiceNowEstimateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/unbilled_charges/invoice_now_estimate"), params)
}
