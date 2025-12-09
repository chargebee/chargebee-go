package chargebee

import (
	"fmt"
	"net/url"
)

type CouponService struct {
	transport *Transport
}

func (s *CouponService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/coupons"), req).SetIdempotency(true)
}

func (s *CouponService) CreateForItems(req *CreateForItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/coupons/create_for_items"), req).SetIdempotency(true)
}

func (s *CouponService) UpdateForItems(id string, req *UpdateForItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/coupons/%v/update_for_items", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CouponService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/coupons"), req)
}

func (s *CouponService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/coupons/%v", url.PathEscape(id)), nil)
}

func (s *CouponService) Update(id string, req *UpdateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/coupons/%v", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CouponService) Delete(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/coupons/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *CouponService) Copy(req *CopyRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/coupons/copy"), req).SetIdempotency(true)
}

func (s *CouponService) Unarchive(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/coupons/%v/unarchive", url.PathEscape(id)), nil).SetIdempotency(true)
}
