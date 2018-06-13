package enum

type UnpaidInvoicesHandling string

const (
	UnpaidInvoicesHandlingNoAction                  UnpaidInvoicesHandling = "no_action"
	UnpaidInvoicesHandlingSchedulePaymentCollection UnpaidInvoicesHandling = "schedule_payment_collection"
)
