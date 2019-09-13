package transaction

import (
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/transaction"
)

func CreateAuthorization(params *transaction.CreateAuthorizationRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/transactions/create_authorization"), params)
}
func VoidTransaction(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/transactions/%v/void", id), nil)
}
func RecordRefund(id string, params *transaction.RecordRefundRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/transactions/%v/record_refund", id), params)
}
func List(params *transaction.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/transactions"), params)
}
func TransactionsForCustomer(id string, params *transaction.TransactionsForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/transactions", id), params)
}
func TransactionsForSubscription(id string, params *transaction.TransactionsForSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/transactions", id), params)
}
func PaymentsForInvoice(id string, params *transaction.PaymentsForInvoiceRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/invoices/%v/payments", id), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/transactions/%v", id), nil)
}
func DeleteOfflineTransaction(id string, params *transaction.DeleteOfflineTransactionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/transactions/%v/delete_offline_transaction", id), params)
}
