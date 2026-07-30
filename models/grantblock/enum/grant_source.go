package enum

type GrantSource string

const (
	GrantSourceSubscriptionCreated GrantSource = "subscription_created"
	GrantSourceSubscriptionChanged GrantSource = "subscription_changed"
	GrantSourceTopUp               GrantSource = "top_up"
	GrantSourcePromotionalGrants   GrantSource = "promotional_grants"
	GrantSourceRollover            GrantSource = "rollover"
	GrantSourceGrantRenewal        GrantSource = "grant_renewal"
	GrantSourceSubscriptionRenewed GrantSource = "subscription_renewed"
)
