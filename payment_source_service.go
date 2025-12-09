package chargebee

import (
	"fmt"
	"net/url"
)

type PaymentSourceService struct {
	transport *Transport
}

func (s *PaymentSourceService) CreateUsingTempToken(req *CreateUsingTempTokenRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/payment_sources/create_using_temp_token"), req).SetIdempotency(true)
}

func (s *PaymentSourceService) CreateUsingPermanentToken(req *CreateUsingPermanentTokenRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/payment_sources/create_using_permanent_token"), req).SetIdempotency(true)
}

func (s *PaymentSourceService) CreateUsingToken(req *CreateUsingTokenRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/payment_sources/create_using_token"), req).SetIdempotency(true)
}

func (s *PaymentSourceService) CreateUsingPaymentIntent(req *CreateUsingPaymentIntentRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/payment_sources/create_using_payment_intent"), req).SetIdempotency(true)
}

func (s *PaymentSourceService) CreateVoucherPaymentSource(req *CreateVoucherPaymentSourceRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/payment_sources/create_voucher_payment_source"), req).SetIdempotency(true)
}

func (s *PaymentSourceService) CreateCard(req *CreateCardRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/payment_sources/create_card"), req).SetIdempotency(true)
}

func (s *PaymentSourceService) CreateBankAccount(req *CreateBankAccountRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/payment_sources/create_bank_account"), req).SetIdempotency(true)
}

func (s *PaymentSourceService) UpdateCard(id string, req *UpdateCardRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/payment_sources/%v/update_card", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *PaymentSourceService) UpdateBankAccount(id string, req *UpdateBankAccountRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/payment_sources/%v/update_bank_account", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *PaymentSourceService) VerifyBankAccount(id string, req *VerifyBankAccountRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/payment_sources/%v/verify_bank_account", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *PaymentSourceService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/payment_sources/%v", url.PathEscape(id)), nil)
}

func (s *PaymentSourceService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/payment_sources"), req)
}

func (s *PaymentSourceService) SwitchGatewayAccount(id string, req *SwitchGatewayAccountRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/payment_sources/%v/switch_gateway_account", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *PaymentSourceService) ExportPaymentSource(id string, req *ExportPaymentSourceRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/payment_sources/%v/export_payment_source", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *PaymentSourceService) Delete(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/payment_sources/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *PaymentSourceService) DeleteLocal(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/payment_sources/%v/delete_local", url.PathEscape(id)), nil).SetIdempotency(true)
}
