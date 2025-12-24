package chargebee

type AddressService struct {
	config *ClientConfig
}

func (s *AddressService) Retrieve(req *AddressRetrieveRequest) (*AddressRetrieveResponse, error) {
	req.method = "GET"
	req.path = "/addresses"
	res, err := send[*AddressRetrieveResponse](req, s.config)
	return *res, err
}

func (s *AddressService) Update(req *AddressUpdateRequest) (*AddressUpdateResponse, error) {
	req.method = "POST"
	req.path = "/addresses"
	req.SetIdempotency(true)
	res, err := send[*AddressUpdateResponse](req, s.config)
	return *res, err
}
