package enum

type AccountReceivablesHandling string

const (
	AccountReceivablesHandlingNoAction                  AccountReceivablesHandling = "no_action"
	AccountReceivablesHandlingSchedulePaymentCollection AccountReceivablesHandling = "schedule_payment_collection"
	AccountReceivablesHandlingWriteOff                  AccountReceivablesHandling = "write_off"
)
