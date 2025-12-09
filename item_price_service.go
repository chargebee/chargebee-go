package chargebee

import (
	"fmt"
	"net/url"
)

type ItemPriceService struct {
	transport *Transport
}

func (s *ItemPriceService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/item_prices"), req).SetIdempotency(true)
}

func (s *ItemPriceService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/item_prices/%v", url.PathEscape(id)), nil)
}

func (s *ItemPriceService) Update(id string, req *UpdateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/item_prices/%v", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *ItemPriceService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/item_prices"), req)
}

func (s *ItemPriceService) Delete(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/item_prices/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *ItemPriceService) FindApplicableItems(id string, req *FindApplicableItemsRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/item_prices/%v/applicable_items", url.PathEscape(id)), req)
}

func (s *ItemPriceService) FindApplicableItemPrices(id string, req *FindApplicableItemPricesRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/item_prices/%v/applicable_item_prices", url.PathEscape(id)), req)
}

func (s *ItemPriceService) MoveItemPrice(id string, req *MoveItemPriceRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/item_prices/%v/move", url.PathEscape(id)), req).SetIdempotency(true)
}
