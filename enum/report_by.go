package enum

type ReportBy string

const (
	ReportByCustomer     ReportBy = "customer"
	ReportByInvoice      ReportBy = "invoice"
	ReportByProduct      ReportBy = "product"
	ReportBySubscription ReportBy = "subscription"
)
