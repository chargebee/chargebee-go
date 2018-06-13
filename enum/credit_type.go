package enum

type CreditType string

const (
	CreditTypeLoyaltyCredits  CreditType = "loyalty_credits"
	CreditTypeReferralRewards CreditType = "referral_rewards"
	CreditTypeGeneral         CreditType = "general"
)
