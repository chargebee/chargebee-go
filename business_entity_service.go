package chargebee

import (
	"fmt"
)

type BusinessEntityService struct {
	config *ClientConfig
}

func (s *BusinessEntityService) CreateTransfers(req *BusinessEntityCreateTransfersRequest) (*BusinessEntityCreateTransfersResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/business_entities/transfers")
	req.isIdempotent = true
	return send[*BusinessEntityCreateTransfersResponse](req, s.config)
}

func (s *BusinessEntityService) GetTransfers(req *BusinessEntityGetTransfersRequest) (*BusinessEntityGetTransfersResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/business_entities/transfers")
	req.isListRequest = true
	return send[*BusinessEntityGetTransfersResponse](req, s.config)
}
