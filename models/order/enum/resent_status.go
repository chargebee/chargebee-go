package enum

type ResentStatus string

const (
	ResentStatusFullyResent     ResentStatus = "fully_resent"
	ResentStatusPartiallyResent ResentStatus = "partially_resent"
)
