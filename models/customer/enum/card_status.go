package enum

type CardStatus string

const (
	CardStatusNoCard              CardStatus = "no_card"
	CardStatusValid               CardStatus = "valid"
	CardStatusExpiring            CardStatus = "expiring"
	CardStatusExpired             CardStatus = "expired"
	CardStatusPendingVerification CardStatus = "pending_verification"
	CardStatusInvalid             CardStatus = "invalid"
)
