package enum

type OrderLineItemLinkedCreditType string

const (
	OrderLineItemLinkedCreditTypeAdjustment OrderLineItemLinkedCreditType = "adjustment"
	OrderLineItemLinkedCreditTypeRefundable OrderLineItemLinkedCreditType = "refundable"
)
