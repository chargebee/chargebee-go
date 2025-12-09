package chargebee

import (
	"fmt"
	"net/url"
)

type PaymentVoucherService struct {
	transport *Transport
}

func (s *PaymentVoucherService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/payment_vouchers"), req).SetIdempotency(true)
}

func (s *PaymentVoucherService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/payment_vouchers/%v", url.PathEscape(id)), nil)
}

func (s *PaymentVoucherService) PaymentVouchersForInvoice(id string, req *PaymentVouchersForInvoiceRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/invoices/%v/payment_vouchers", url.PathEscape(id)), req)
}

func (s *PaymentVoucherService) PaymentVouchersForCustomer(id string, req *PaymentVouchersForCustomerRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/customers/%v/payment_vouchers", url.PathEscape(id)), req)
}

// Deprecated: This function is deprecated. Please use PaymentVouchersForInvoice instead.
func Payment_vouchersForInvoice(id string, params *paymentvoucher.PaymentVouchersForInvoiceRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/invoices/%v/payment_vouchers", url.PathEscape(id)), params)
}

// Deprecated: This function is deprecated. Please use PaymentVouchersForCustomer instead.
func Payment_vouchersForCustomer(id string, params *paymentvoucher.PaymentVouchersForCustomerRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/payment_vouchers", url.PathEscape(id)), params)
}
