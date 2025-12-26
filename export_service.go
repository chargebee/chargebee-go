package chargebee

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

type ExportService struct {
	config *ClientConfig
}

func (s *ExportService) Retrieve(id string) (*ExportRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/exports/%v", url.PathEscape(id))
	return send[*ExportRetrieveResponse](req, s.config)
}

func (s *ExportService) RevenueRecognition(req *ExportRevenueRecognitionRequest) (*ExportRevenueRecognitionResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/exports/revenue_recognition")
	req.isIdempotent = true
	return send[*ExportRevenueRecognitionResponse](req, s.config)
}

func (s *ExportService) DeferredRevenue(req *ExportDeferredRevenueRequest) (*ExportDeferredRevenueResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/exports/deferred_revenue")
	req.isIdempotent = true
	return send[*ExportDeferredRevenueResponse](req, s.config)
}

func (s *ExportService) Plans(req *ExportPlansRequest) (*ExportPlansResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/exports/plans")
	req.isIdempotent = true
	return send[*ExportPlansResponse](req, s.config)
}

func (s *ExportService) Addons(req *ExportAddonsRequest) (*ExportAddonsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/exports/addons")
	req.isIdempotent = true
	return send[*ExportAddonsResponse](req, s.config)
}

func (s *ExportService) Coupons(req *ExportCouponsRequest) (*ExportCouponsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/exports/coupons")
	req.isIdempotent = true
	return send[*ExportCouponsResponse](req, s.config)
}

func (s *ExportService) Customers(req *ExportCustomersRequest) (*ExportCustomersResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/exports/customers")
	req.isIdempotent = true
	return send[*ExportCustomersResponse](req, s.config)
}

func (s *ExportService) Subscriptions(req *ExportSubscriptionsRequest) (*ExportSubscriptionsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/exports/subscriptions")
	req.isIdempotent = true
	return send[*ExportSubscriptionsResponse](req, s.config)
}

func (s *ExportService) Invoices(req *ExportInvoicesRequest) (*ExportInvoicesResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/exports/invoices")
	req.isIdempotent = true
	return send[*ExportInvoicesResponse](req, s.config)
}

func (s *ExportService) CreditNotes(req *ExportCreditNotesRequest) (*ExportCreditNotesResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/exports/credit_notes")
	req.isIdempotent = true
	return send[*ExportCreditNotesResponse](req, s.config)
}

func (s *ExportService) Transactions(req *ExportTransactionsRequest) (*ExportTransactionsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/exports/transactions")
	req.isIdempotent = true
	return send[*ExportTransactionsResponse](req, s.config)
}

func (s *ExportService) Orders(req *ExportOrdersRequest) (*ExportOrdersResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/exports/orders")
	req.isIdempotent = true
	return send[*ExportOrdersResponse](req, s.config)
}

func (s *ExportService) ItemFamilies(req *ExportItemFamiliesRequest) (*ExportItemFamiliesResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/exports/item_families")
	req.isIdempotent = true
	return send[*ExportItemFamiliesResponse](req, s.config)
}

func (s *ExportService) Items(req *ExportItemsRequest) (*ExportItemsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/exports/items")
	req.isIdempotent = true
	return send[*ExportItemsResponse](req, s.config)
}

func (s *ExportService) ItemPrices(req *ExportItemPricesRequest) (*ExportItemPricesResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/exports/item_prices")
	req.isIdempotent = true
	return send[*ExportItemPricesResponse](req, s.config)
}

func (s *ExportService) AttachedItems(req *ExportAttachedItemsRequest) (*ExportAttachedItemsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/exports/attached_items")
	req.isIdempotent = true
	return send[*ExportAttachedItemsResponse](req, s.config)
}

func (s *ExportService) DifferentialPrices(req *ExportDifferentialPricesRequest) (*ExportDifferentialPricesResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/exports/differential_prices")
	req.isIdempotent = true
	return send[*ExportDifferentialPricesResponse](req, s.config)
}

func (s *ExportService) PriceVariants(req *ExportPriceVariantsRequest) (*ExportPriceVariantsResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/exports/price_variants")
	req.isIdempotent = true
	return send[*ExportPriceVariantsResponse](req, s.config)
}

func (s *ExportService) WaitForExportCompletion(exp Export) (Export, error) {
	count := 0
	for exp.Status == ExportStatusInProcess {
		if count > 30 {
			return exp, errors.New("Export is taking too long")
		}
		count++
		time.Sleep(ExportWaitInSecs)
		response, err := s.Retrieve(exp.Id)
		if err != nil {
			return exp, err
		}
		exp = *response.Export
	}

	return exp, nil
}
