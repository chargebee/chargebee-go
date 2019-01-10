package enum

type ChargesHandling string

const (
	ChargesHandlingInvoiceImmediately   ChargesHandling = "invoice_immediately"
	ChargesHandlingAddToUnbilledCharges ChargesHandling = "add_to_unbilled_charges"
)
