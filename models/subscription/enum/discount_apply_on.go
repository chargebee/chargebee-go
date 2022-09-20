package enum

type DiscountApplyOn string

const (
	DiscountApplyOnInvoiceAmount     DiscountApplyOn = "invoice_amount"
	DiscountApplyOnSpecificItemPrice DiscountApplyOn = "specific_item_price"
)
