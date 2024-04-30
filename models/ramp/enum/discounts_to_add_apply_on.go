package enum

type DiscountsToAddApplyOn string

const (
	DiscountsToAddApplyOnInvoiceAmount     DiscountsToAddApplyOn = "invoice_amount"
	DiscountsToAddApplyOnSpecificItemPrice DiscountsToAddApplyOn = "specific_item_price"
)
