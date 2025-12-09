package chargebee

import (
	"fmt"
	"net/url"
)

type VirtualBankAccountService struct {
	transport *Transport
}

func (s *VirtualBankAccountService) CreateUsingPermanentToken(req *CreateUsingPermanentTokenRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/virtual_bank_accounts/create_using_permanent_token"), req).SetIdempotency(true)
}

func (s *VirtualBankAccountService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/virtual_bank_accounts"), req).SetIdempotency(true)
}

func (s *VirtualBankAccountService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/virtual_bank_accounts/%v", url.PathEscape(id)), nil)
}

func (s *VirtualBankAccountService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/virtual_bank_accounts"), req)
}

func (s *VirtualBankAccountService) Delete(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/virtual_bank_accounts/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *VirtualBankAccountService) DeleteLocal(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/virtual_bank_accounts/%v/delete_local", url.PathEscape(id)), nil).SetIdempotency(true)
}
