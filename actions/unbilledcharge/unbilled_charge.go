package unbilledcharge

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"
	"net/url"
)

func CreateUnbilledCharge(params *unbilledcharge.CreateUnbilledChargeRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/unbilled_charges/create"), params).WithTelemetryResource("unbilledCharge").WithTelemetryOperation("createUnbilledCharge").SetIdempotency(true)
}
func Create(params *unbilledcharge.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/unbilled_charges"), params).WithTelemetryResource("unbilledCharge").WithTelemetryOperation("create").SetIdempotency(true)
}
func InvoiceUnbilledCharges(params *unbilledcharge.InvoiceUnbilledChargesRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/unbilled_charges/invoice_unbilled_charges"), params).WithTelemetryResource("unbilledCharge").WithTelemetryOperation("invoiceUnbilledCharges").SetIdempotency(true)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/unbilled_charges/%v/delete", url.PathEscape(id)), nil).WithTelemetryResource("unbilledCharge").WithTelemetryOperation("delete").SetIdempotency(true)
}
func List(params *unbilledcharge.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/unbilled_charges"), params).WithTelemetryResource("unbilledCharge").WithTelemetryOperation("list")
}
func InvoiceNowEstimate(params *unbilledcharge.InvoiceNowEstimateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/unbilled_charges/invoice_now_estimate"), params).WithTelemetryResource("unbilledCharge").WithTelemetryOperation("invoiceNowEstimate")
}
