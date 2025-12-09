package chargebee

import (
	"fmt"
	"net/url"
)

type CouponSetService struct {
	transport *Transport
}

func (s *CouponSetService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/coupon_sets"), req).SetIdempotency(true)
}

func (s *CouponSetService) AddCouponCodes(id string, req *AddCouponCodesRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/coupon_sets/%v/add_coupon_codes", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CouponSetService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/coupon_sets"), req)
}

func (s *CouponSetService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/coupon_sets/%v", url.PathEscape(id)), nil)
}

func (s *CouponSetService) Update(id string, req *UpdateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/coupon_sets/%v/update", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CouponSetService) Delete(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/coupon_sets/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *CouponSetService) DeleteUnusedCouponCodes(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/coupon_sets/%v/delete_unused_coupon_codes", url.PathEscape(id)), nil).SetIdempotency(true)
}
