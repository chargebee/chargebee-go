package enum

type UnbilledChargesHandling string

const (
	UnbilledChargesHandlingNoAction UnbilledChargesHandling = "no_action"
	UnbilledChargesHandlingInvoice  UnbilledChargesHandling = "invoice"
)
