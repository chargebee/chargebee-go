package chargebee

import (
	"fmt"
	"net/url"
)

type GiftService struct {
	transport *Transport
}

func (s *GiftService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/gifts"), req).SetIdempotency(true)
}

func (s *GiftService) CreateForItems(req *CreateForItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/gifts/create_for_items"), req).SetIdempotency(true)
}

func (s *GiftService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/gifts/%v", url.PathEscape(id)), nil)
}

func (s *GiftService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/gifts"), req)
}

func (s *GiftService) Claim(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/gifts/%v/claim", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *GiftService) Cancel(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/gifts/%v/cancel", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *GiftService) UpdateGift(id string, req *UpdateGiftRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/gifts/%v/update_gift", url.PathEscape(id)), req).SetIdempotency(true)
}
