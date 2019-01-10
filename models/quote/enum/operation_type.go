package enum

type OperationType string

const (
	OperationTypeCreateSubscriptionForCustomer OperationType = "create_subscription_for_customer"
	OperationTypeChangeSubscription            OperationType = "change_subscription"
	OperationTypeOnetimeInvoice                OperationType = "onetime_invoice"
)
