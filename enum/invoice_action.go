package enum

type InvoiceAction string

const (
	InvoiceActionVoid     InvoiceAction = "void"
	InvoiceActionWriteOff InvoiceAction = "write_off"
)
