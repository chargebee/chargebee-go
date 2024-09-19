package enum

type ScheduleEntryStatus string

const (
	ScheduleEntryStatusPosted     ScheduleEntryStatus = "posted"
	ScheduleEntryStatusPaymentDue ScheduleEntryStatus = "payment_due"
	ScheduleEntryStatusPaid       ScheduleEntryStatus = "paid"
)
