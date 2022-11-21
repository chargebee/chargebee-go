package unbilledcharge

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"
	"net/url"
)

func CreateUnbilledCharge(params *unbilledcharge.CreateUnbilledChargeRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/unbilled_charges/create"), params)
}
func Create(params *unbilledcharge.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/unbilled_charges"), params)
}
func InvoiceUnbilledCharges(params *unbilledcharge.InvoiceUnbilledChargesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/unbilled_charges/invoice_unbilled_charges"), params)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/unbilled_charges/%v/delete", url.PathEscape(id)), nil)
}
func List(params *unbilledcharge.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/unbilled_charges"), params)
}
func InvoiceNowEstimate(params *unbilledcharge.InvoiceNowEstimateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/unbilled_charges/invoice_now_estimate"), params)
}
