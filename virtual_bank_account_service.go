package chargebee

import (
	"fmt"
	"net/url"
)

type VirtualBankAccountService struct {
	config *ClientConfig
}

func (s *VirtualBankAccountService) CreateUsingPermanentToken(req *VirtualBankAccountCreateUsingPermanentTokenRequest) (*VirtualBankAccountCreateUsingPermanentTokenResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/virtual_bank_accounts/create_using_permanent_token")
	req.isIdempotent = true
	return send[*VirtualBankAccountCreateUsingPermanentTokenResponse](req, s.config)
}

func (s *VirtualBankAccountService) Create(req *VirtualBankAccountCreateRequest) (*VirtualBankAccountCreateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/virtual_bank_accounts")
	req.isIdempotent = true
	return send[*VirtualBankAccountCreateResponse](req, s.config)
}

func (s *VirtualBankAccountService) Retrieve(id string) (*VirtualBankAccountRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/virtual_bank_accounts/%v", url.PathEscape(id))
	return send[*VirtualBankAccountRetrieveResponse](req, s.config)
}

func (s *VirtualBankAccountService) List(req *VirtualBankAccountListRequest) (*VirtualBankAccountListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/virtual_bank_accounts")
	req.isListRequest = true
	return send[*VirtualBankAccountListResponse](req, s.config)
}

func (s *VirtualBankAccountService) Delete(id string) (*VirtualBankAccountDeleteResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/virtual_bank_accounts/%v/delete", url.PathEscape(id))
	req.isIdempotent = true
	return send[*VirtualBankAccountDeleteResponse](req, s.config)
}

func (s *VirtualBankAccountService) DeleteLocal(id string) (*VirtualBankAccountDeleteLocalResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/virtual_bank_accounts/%v/delete_local", url.PathEscape(id))
	req.isIdempotent = true
	return send[*VirtualBankAccountDeleteLocalResponse](req, s.config)
}
