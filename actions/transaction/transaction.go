package transaction

import (
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/transaction"
)

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
