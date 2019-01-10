package enum

type OrderLineItemLinkedCreditStatus string

const (
	OrderLineItemLinkedCreditStatusAdjusted  OrderLineItemLinkedCreditStatus = "adjusted"
	OrderLineItemLinkedCreditStatusRefunded  OrderLineItemLinkedCreditStatus = "refunded"
	OrderLineItemLinkedCreditStatusRefundDue OrderLineItemLinkedCreditStatus = "refund_due"
	OrderLineItemLinkedCreditStatusVoided    OrderLineItemLinkedCreditStatus = "voided"
)
