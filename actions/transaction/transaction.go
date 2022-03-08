package transaction

import (
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/transaction"
	"net/url"
)

func CreateAuthorization(params *transaction.CreateAuthorizationRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/transactions/create_authorization"), params)
}
func VoidTransaction(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/transactions/%v/void", url.PathEscape(id)), nil)
}
func RecordRefund(id string, params *transaction.RecordRefundRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/transactions/%v/record_refund", url.PathEscape(id)), params)
}
func Refund(id string, params *transaction.RefundRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/transactions/%v/refund", url.PathEscape(id)), params)
}
func List(params *transaction.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/transactions"), params)
}
func TransactionsForCustomer(id string, params *transaction.TransactionsForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/transactions", url.PathEscape(id)), params)
}
func TransactionsForSubscription(id string, params *transaction.TransactionsForSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/transactions", url.PathEscape(id)), params)
}
func PaymentsForInvoice(id string, params *transaction.PaymentsForInvoiceRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/invoices/%v/payments", url.PathEscape(id)), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/transactions/%v", url.PathEscape(id)), nil)
}
func DeleteOfflineTransaction(id string, params *transaction.DeleteOfflineTransactionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/transactions/%v/delete_offline_transaction", url.PathEscape(id)), params)
}
