package paymentvoucher

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/paymentvoucher"
	"net/url"
)

func Create(params *paymentvoucher.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/payment_vouchers"), params).WithTelemetryResource("paymentVoucher").WithTelemetryOperation("create").SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/payment_vouchers/%v", url.PathEscape(id)), nil).WithTelemetryResource("paymentVoucher").WithTelemetryOperation("retrieve")
}
func PaymentVouchersForInvoice(id string, params *paymentvoucher.PaymentVouchersForInvoiceRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/invoices/%v/payment_vouchers", url.PathEscape(id)), params).WithTelemetryResource("paymentVoucher").WithTelemetryOperation("paymentVouchersForInvoice")
}
func PaymentVouchersForCustomer(id string, params *paymentvoucher.PaymentVouchersForCustomerRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/payment_vouchers", url.PathEscape(id)), params).WithTelemetryResource("paymentVoucher").WithTelemetryOperation("paymentVouchersForCustomer")
}

// Deprecated: This function is deprecated. Please use PaymentVouchersForInvoice instead.
func Payment_vouchersForInvoice(id string, params *paymentvoucher.PaymentVouchersForInvoiceRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/invoices/%v/payment_vouchers", url.PathEscape(id)), params)
}

// Deprecated: This function is deprecated. Please use PaymentVouchersForCustomer instead.
func Payment_vouchersForCustomer(id string, params *paymentvoucher.PaymentVouchersForCustomerRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/payment_vouchers", url.PathEscape(id)), params)
}
