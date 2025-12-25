package chargebee

import (
	"fmt"
	"net/url"
)

type PaymentSourceService struct {
	config *ClientConfig
}

func (s *PaymentSourceService) CreateUsingTempToken(req *PaymentSourceCreateUsingTempTokenRequest) (*PaymentSourceCreateUsingTempTokenResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/payment_sources/create_using_temp_token")
	req.isIdempotent = true
	return send[*PaymentSourceCreateUsingTempTokenResponse](req, s.config)
}

func (s *PaymentSourceService) CreateUsingPermanentToken(req *PaymentSourceCreateUsingPermanentTokenRequest) (*PaymentSourceCreateUsingPermanentTokenResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/payment_sources/create_using_permanent_token")
	req.isIdempotent = true
	return send[*PaymentSourceCreateUsingPermanentTokenResponse](req, s.config)
}

func (s *PaymentSourceService) CreateUsingToken(req *PaymentSourceCreateUsingTokenRequest) (*PaymentSourceCreateUsingTokenResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/payment_sources/create_using_token")
	req.isIdempotent = true
	return send[*PaymentSourceCreateUsingTokenResponse](req, s.config)
}

func (s *PaymentSourceService) CreateUsingPaymentIntent(req *PaymentSourceCreateUsingPaymentIntentRequest) (*PaymentSourceCreateUsingPaymentIntentResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/payment_sources/create_using_payment_intent")
	req.isIdempotent = true
	return send[*PaymentSourceCreateUsingPaymentIntentResponse](req, s.config)
}

func (s *PaymentSourceService) CreateVoucherPaymentSource(req *PaymentSourceCreateVoucherPaymentSourceRequest) (*PaymentSourceCreateVoucherPaymentSourceResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/payment_sources/create_voucher_payment_source")
	req.isIdempotent = true
	return send[*PaymentSourceCreateVoucherPaymentSourceResponse](req, s.config)
}

func (s *PaymentSourceService) CreateCard(req *PaymentSourceCreateCardRequest) (*PaymentSourceCreateCardResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/payment_sources/create_card")
	req.isIdempotent = true
	return send[*PaymentSourceCreateCardResponse](req, s.config)
}

func (s *PaymentSourceService) CreateBankAccount(req *PaymentSourceCreateBankAccountRequest) (*PaymentSourceCreateBankAccountResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/payment_sources/create_bank_account")
	req.isIdempotent = true
	return send[*PaymentSourceCreateBankAccountResponse](req, s.config)
}

func (s *PaymentSourceService) UpdateCard(id string, req *PaymentSourceUpdateCardRequest) (*PaymentSourceUpdateCardResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/payment_sources/%v/update_card", url.PathEscape(id))
	req.isIdempotent = true
	return send[*PaymentSourceUpdateCardResponse](req, s.config)
}

func (s *PaymentSourceService) UpdateBankAccount(id string, req *PaymentSourceUpdateBankAccountRequest) (*PaymentSourceUpdateBankAccountResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/payment_sources/%v/update_bank_account", url.PathEscape(id))
	req.isIdempotent = true
	return send[*PaymentSourceUpdateBankAccountResponse](req, s.config)
}

func (s *PaymentSourceService) VerifyBankAccount(id string, req *PaymentSourceVerifyBankAccountRequest) (*PaymentSourceVerifyBankAccountResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/payment_sources/%v/verify_bank_account", url.PathEscape(id))
	req.isIdempotent = true
	return send[*PaymentSourceVerifyBankAccountResponse](req, s.config)
}

func (s *PaymentSourceService) Retrieve(id string) (*PaymentSourceRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/payment_sources/%v", url.PathEscape(id))
	return send[*PaymentSourceRetrieveResponse](req, s.config)
}

func (s *PaymentSourceService) List(req *PaymentSourceListRequest) (*PaymentSourceListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/payment_sources")
	req.isListRequest = true
	return send[*PaymentSourceListResponse](req, s.config)
}

func (s *PaymentSourceService) SwitchGatewayAccount(id string, req *PaymentSourceSwitchGatewayAccountRequest) (*PaymentSourceSwitchGatewayAccountResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/payment_sources/%v/switch_gateway_account", url.PathEscape(id))
	req.isIdempotent = true
	return send[*PaymentSourceSwitchGatewayAccountResponse](req, s.config)
}

func (s *PaymentSourceService) ExportPaymentSource(id string, req *PaymentSourceExportPaymentSourceRequest) (*PaymentSourceExportPaymentSourceResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/payment_sources/%v/export_payment_source", url.PathEscape(id))
	req.isIdempotent = true
	return send[*PaymentSourceExportPaymentSourceResponse](req, s.config)
}

func (s *PaymentSourceService) Delete(id string) (*PaymentSourceDeleteResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/payment_sources/%v/delete", url.PathEscape(id))
	req.isIdempotent = true
	return send[*PaymentSourceDeleteResponse](req, s.config)
}

func (s *PaymentSourceService) DeleteLocal(id string) (*PaymentSourceDeleteLocalResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/payment_sources/%v/delete_local", url.PathEscape(id))
	req.isIdempotent = true
	return send[*PaymentSourceDeleteLocalResponse](req, s.config)
}
