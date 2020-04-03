package enum

type ActionAtTermEnd string

const (
	ActionAtTermEndRenew     ActionAtTermEnd = "renew"
	ActionAtTermEndEvergreen ActionAtTermEnd = "evergreen"
	ActionAtTermEndCancel    ActionAtTermEnd = "cancel"
	ActionAtTermEndRenewOnce ActionAtTermEnd = "renew_once"
)
