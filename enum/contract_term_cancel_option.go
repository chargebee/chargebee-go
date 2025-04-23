package enum

type ContractTermCancelOption string

const (
	ContractTermCancelOptionTerminateImmediately         ContractTermCancelOption = "terminate_immediately"
	ContractTermCancelOptionEndOfContractTerm            ContractTermCancelOption = "end_of_contract_term"
	ContractTermCancelOptionSpecificDate                 ContractTermCancelOption = "specific_date"
	ContractTermCancelOptionEndOfSubscriptionBillingTerm ContractTermCancelOption = "end_of_subscription_billing_term"
)
