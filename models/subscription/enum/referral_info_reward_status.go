package enum

type ReferralInfoRewardStatus string

const (
	ReferralInfoRewardStatusPending ReferralInfoRewardStatus = "pending"
	ReferralInfoRewardStatusPaid    ReferralInfoRewardStatus = "paid"
	ReferralInfoRewardStatusInvalid ReferralInfoRewardStatus = "invalid"
)
