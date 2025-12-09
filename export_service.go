package chargebee

import (
	"errors"
	"fmt"
	"time"

	"net/url"
)

type ExportService struct {
	transport *Transport
}

func (s *ExportService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/exports/%v", url.PathEscape(id)), nil)
}

func (s *ExportService) RevenueRecognition(req *RevenueRecognitionRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/exports/revenue_recognition"), req).SetIdempotency(true)
}

func (s *ExportService) DeferredRevenue(req *DeferredRevenueRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/exports/deferred_revenue"), req).SetIdempotency(true)
}

func (s *ExportService) Plans(req *PlansRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/exports/plans"), req).SetIdempotency(true)
}

func (s *ExportService) Addons(req *AddonsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/exports/addons"), req).SetIdempotency(true)
}

func (s *ExportService) Coupons(req *CouponsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/exports/coupons"), req).SetIdempotency(true)
}

func (s *ExportService) Customers(req *CustomersRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/exports/customers"), req).SetIdempotency(true)
}

func (s *ExportService) Subscriptions(req *SubscriptionsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/exports/subscriptions"), req).SetIdempotency(true)
}

func (s *ExportService) Invoices(req *InvoicesRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/exports/invoices"), req).SetIdempotency(true)
}

func (s *ExportService) CreditNotes(req *CreditNotesRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/exports/credit_notes"), req).SetIdempotency(true)
}

func (s *ExportService) Transactions(req *TransactionsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/exports/transactions"), req).SetIdempotency(true)
}

func (s *ExportService) Orders(req *OrdersRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/exports/orders"), req).SetIdempotency(true)
}

func (s *ExportService) ItemFamilies(req *ItemFamiliesRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/exports/item_families"), req).SetIdempotency(true)
}

func (s *ExportService) Items(req *ItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/exports/items"), req).SetIdempotency(true)
}

func (s *ExportService) ItemPrices(req *ItemPricesRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/exports/item_prices"), req).SetIdempotency(true)
}

func (s *ExportService) AttachedItems(req *AttachedItemsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/exports/attached_items"), req).SetIdempotency(true)
}

func (s *ExportService) DifferentialPrices(req *DifferentialPricesRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/exports/differential_prices"), req).SetIdempotency(true)
}

func (s *ExportService) PriceVariants(req *PriceVariantsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/exports/price_variants"), req).SetIdempotency(true)
}

func WaitForExportCompletion(exp Export) (Export, error) {
	return WaitForExportCompletionWithEnv(exp, chargebee.DefaultConfig())
}

func WaitForExportCompletionWithEnv(exp Export, env chargebee.Environment) (Export, error) {
	count := 0
	for exp.Status == StatusInProcess {
		if count > 30 {
			return exp, errors.New("'Export is taking too long'")
		}
		count++
		time.Sleep(chargebee.ExportWaitInSecs)
		result, err := Retrieve(exp.Id).RequestWithEnv(env)
		if err != nil {
			return exp, err
		}
		exp = *result.Export
	}
	return exp, nil
}
