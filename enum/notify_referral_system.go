package enum

type NotifyReferralSystem string

const (
	NotifyReferralSystemNone                NotifyReferralSystem = "none"
	NotifyReferralSystemFirstPaidConversion NotifyReferralSystem = "first_paid_conversion"
	NotifyReferralSystemAllInvoices         NotifyReferralSystem = "all_invoices"
)
