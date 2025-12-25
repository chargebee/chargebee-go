package chargebee

import (
	"fmt"
	"net/url"
)

type GiftService struct {
	config *ClientConfig
}

func (s *GiftService) Create(req *GiftCreateRequest) (*GiftCreateResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/gifts")
	req.isIdempotent = true
	return send[*GiftCreateResponse](req, s.config)
}

func (s *GiftService) CreateForItems(req *GiftCreateForItemsRequest) (*GiftCreateForItemsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/gifts/create_for_items")
	req.isIdempotent = true
	return send[*GiftCreateForItemsResponse](req, s.config)
}

func (s *GiftService) Retrieve(id string) (*GiftRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/gifts/%v", url.PathEscape(id))
	return send[*GiftRetrieveResponse](req, s.config)
}

func (s *GiftService) List(req *GiftListRequest) (*GiftListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/gifts")
	req.isListRequest = true
	return send[*GiftListResponse](req, s.config)
}

func (s *GiftService) Claim(id string) (*GiftClaimResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/gifts/%v/claim", url.PathEscape(id))
	req.isIdempotent = true
	return send[*GiftClaimResponse](req, s.config)
}

func (s *GiftService) Cancel(id string) (*GiftCancelResponse, error) {
	req := &BlankRequest{}
	req.method = "POST"
	req.path = fmt.Sprintf("/gifts/%v/cancel", url.PathEscape(id))
	req.isIdempotent = true
	return send[*GiftCancelResponse](req, s.config)
}

func (s *GiftService) UpdateGift(id string, req *GiftUpdateGiftRequest) (*GiftUpdateGiftResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/gifts/%v/update_gift", url.PathEscape(id))
	req.isIdempotent = true
	return send[*GiftUpdateGiftResponse](req, s.config)
}
