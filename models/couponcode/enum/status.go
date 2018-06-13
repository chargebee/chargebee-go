package enum

type Status string

const (
	StatusNotRedeemed Status = "not_redeemed"
	StatusRedeemed    Status = "redeemed"
	StatusArchived    Status = "archived"
)
