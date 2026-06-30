package transaction

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/transaction"
	"net/url"
)

func CreateAuthorization(params *transaction.CreateAuthorizationRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/transactions/create_authorization"), params).WithTelemetryResource("transaction").WithTelemetryOperation("createAuthorization").SetIdempotency(true)
}
func VoidTransaction(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/transactions/%v/void", url.PathEscape(id)), nil).WithTelemetryResource("transaction").WithTelemetryOperation("voidTransaction").SetIdempotency(true)
}
func RecordRefund(id string, params *transaction.RecordRefundRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/transactions/%v/record_refund", url.PathEscape(id)), params).WithTelemetryResource("transaction").WithTelemetryOperation("recordRefund").SetIdempotency(true)
}
func Reconcile(id string, params *transaction.ReconcileRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/transactions/%v/reconcile", url.PathEscape(id)), params).WithTelemetryResource("transaction").WithTelemetryOperation("reconcile").SetIdempotency(true)
}
func Refund(id string, params *transaction.RefundRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/transactions/%v/refund", url.PathEscape(id)), params).WithTelemetryResource("transaction").WithTelemetryOperation("refund").SetIdempotency(true)
}
func List(params *transaction.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/transactions"), params).WithTelemetryResource("transaction").WithTelemetryOperation("list")
}

// Deprecated: This function is deprecated.
func TransactionsForCustomer(id string, params *transaction.TransactionsForCustomerRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/transactions", url.PathEscape(id)), params).WithTelemetryResource("transaction").WithTelemetryOperation("transactionsForCustomer")
}

// Deprecated: This function is deprecated.
func TransactionsForSubscription(id string, params *transaction.TransactionsForSubscriptionRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/transactions", url.PathEscape(id)), params).WithTelemetryResource("transaction").WithTelemetryOperation("transactionsForSubscription")
}
func PaymentsForInvoice(id string, params *transaction.PaymentsForInvoiceRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/invoices/%v/payments", url.PathEscape(id)), params).WithTelemetryResource("transaction").WithTelemetryOperation("paymentsForInvoice")
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/transactions/%v", url.PathEscape(id)), nil).WithTelemetryResource("transaction").WithTelemetryOperation("retrieve")
}
func DeleteOfflineTransaction(id string, params *transaction.DeleteOfflineTransactionRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/transactions/%v/delete_offline_transaction", url.PathEscape(id)), params).WithTelemetryResource("transaction").WithTelemetryOperation("deleteOfflineTransaction").SetIdempotency(true)
}
