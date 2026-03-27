package enum

type CustomerAcceptanceMethod string

const (
	CustomerAcceptanceMethodEsignAndPay CustomerAcceptanceMethod = "esign_and_pay"
	CustomerAcceptanceMethodEsign       CustomerAcceptanceMethod = "esign"
	CustomerAcceptanceMethodPay         CustomerAcceptanceMethod = "pay"
)
