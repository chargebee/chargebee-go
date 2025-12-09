package chargebee

import (
	"fmt"
	"net/url"
)

type TransactionService struct {
	transport *Transport
}

func (s *TransactionService) CreateAuthorization(req *CreateAuthorizationRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/transactions/create_authorization"), req).SetIdempotency(true)
}

func (s *TransactionService) VoidTransaction(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/transactions/%v/void", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *TransactionService) RecordRefund(id string, req *RecordRefundRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/transactions/%v/record_refund", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *TransactionService) Reconcile(id string, req *ReconcileRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/transactions/%v/reconcile", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *TransactionService) Refund(id string, req *RefundRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/transactions/%v/refund", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *TransactionService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/transactions"), req)
}

func (s *TransactionService) TransactionsForCustomer(id string, req *TransactionsForCustomerRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/customers/%v/transactions", url.PathEscape(id)), req)
}

func (s *TransactionService) TransactionsForSubscription(id string, req *TransactionsForSubscriptionRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/subscriptions/%v/transactions", url.PathEscape(id)), req)
}

func (s *TransactionService) PaymentsForInvoice(id string, req *PaymentsForInvoiceRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/invoices/%v/payments", url.PathEscape(id)), req)
}

func (s *TransactionService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/transactions/%v", url.PathEscape(id)), nil)
}

func (s *TransactionService) DeleteOfflineTransaction(id string, req *DeleteOfflineTransactionRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/transactions/%v/delete_offline_transaction", url.PathEscape(id)), req).SetIdempotency(true)
}
