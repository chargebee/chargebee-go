package enum

type ApplyOn string

const (
	ApplyOnInvoiceAmount            ApplyOn = "invoice_amount"
	ApplyOnSpecifiedItemsTotal      ApplyOn = "specified_items_total"
	ApplyOnEachSpecifiedItem        ApplyOn = "each_specified_item"
	ApplyOnEachUnitOfSpecifiedItems ApplyOn = "each_unit_of_specified_items"
)
