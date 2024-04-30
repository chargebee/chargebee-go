package enum

type ApplyOn string

const (
	ApplyOnInvoiceAmount            ApplyOn = "invoice_amount"
	ApplyOnEachSpecifiedItem        ApplyOn = "each_specified_item"
	ApplyOnSpecifiedItemsTotal      ApplyOn = "specified_items_total"
	ApplyOnEachUnitOfSpecifiedItems ApplyOn = "each_unit_of_specified_items"
)
