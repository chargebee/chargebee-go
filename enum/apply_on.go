package enum

type ApplyOn string

const (
	ApplyOnInvoiceAmount     ApplyOn = "invoice_amount"
	ApplyOnSpecificItemPrice ApplyOn = "specific_item_price"
)
