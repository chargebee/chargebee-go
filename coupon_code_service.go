package chargebee

import (
	"fmt"
	"net/url"
)

type CouponCodeService struct {
	transport *Transport
}

func (s *CouponCodeService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/coupon_codes"), req).SetIdempotency(true)
}

func (s *CouponCodeService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/coupon_codes/%v", url.PathEscape(id)), nil)
}

func (s *CouponCodeService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/coupon_codes"), req)
}

func (s *CouponCodeService) Archive(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/coupon_codes/%v/archive", url.PathEscape(id)), nil).SetIdempotency(true)
}
