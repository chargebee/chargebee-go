package chargebee

import (
	"fmt"
	"net/url"
)

type CouponCodeService struct {
	config *ClientConfig
}

func (s *CouponCodeService) Create(req *CouponCodeCreateRequest) (*CouponCodeCreateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/coupon_codes")
	req.isIdempotent = true
	return send[*CouponCodeCreateResponse](req, s.config)
}

func (s *CouponCodeService) Retrieve(id string) (*CouponCodeRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/coupon_codes/%v", url.PathEscape(id))
	return send[*CouponCodeRetrieveResponse](req, s.config)
}

func (s *CouponCodeService) List(req *CouponCodeListRequest) (*CouponCodeListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/coupon_codes")
	req.isListRequest = true
	return send[*CouponCodeListResponse](req, s.config)
}

func (s *CouponCodeService) Archive(id string) (*CouponCodeArchiveResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/coupon_codes/%v/archive", url.PathEscape(id))
	req.isIdempotent = true
	return send[*CouponCodeArchiveResponse](req, s.config)
}
