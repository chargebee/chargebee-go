package enum

type DunningType string

const (
	DunningTypeAutoCollect      DunningType = "auto_collect"
	DunningTypeOffline          DunningType = "offline"
	DunningTypeDirectDebit      DunningType = "direct_debit"
	DunningTypeRealTimePayments DunningType = "real_time_payments"
)
