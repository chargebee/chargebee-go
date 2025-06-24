package paymentvoucher

import (
	"fmt"
	"net/url"

	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/paymentvoucher"
)

func Create(params *paymentvoucher.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/payment_vouchers"), params).SetIdempotency(true)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/payment_vouchers/%v", url.PathEscape(id)), nil)
}
func PaymentVouchersForInvoice(id string, params *paymentvoucher.PaymentVouchersForInvoiceRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/invoices/%v/payment_vouchers", url.PathEscape(id)), params)
}
func PaymentVouchersForCustomer(id string, params *paymentvoucher.PaymentVouchersForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/payment_vouchers", url.PathEscape(id)), params)
}

// Deprecated: This function is deprecated. Please use PaymentVouchersForInvoice instead.
func Payment_vouchersForInvoice(id string, params *paymentvoucher.PaymentVouchersForInvoiceRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/invoices/%v/payment_vouchers", url.PathEscape(id)), params)
}

// Deprecated: This function is deprecated. Please use PaymentVouchersForCustomer instead.
func Payment_vouchersForCustomer(id string, params *paymentvoucher.PaymentVouchersForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/payment_vouchers", url.PathEscape(id)), params)
}
