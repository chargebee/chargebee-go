package enum

type ReferrerRewardType string

const (
	ReferrerRewardTypeNone                      ReferrerRewardType = "none"
	ReferrerRewardTypeReferralDirectReward      ReferrerRewardType = "referral_direct_reward"
	ReferrerRewardTypeCustomPromotionalCredit   ReferrerRewardType = "custom_promotional_credit"
	ReferrerRewardTypeCustomRevenuePercentBased ReferrerRewardType = "custom_revenue_percent_based"
)
