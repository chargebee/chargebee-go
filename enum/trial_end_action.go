package enum

type TrialEndAction string

const (
	TrialEndActionSiteDefault          TrialEndAction = "site_default"
	TrialEndActionPlanDefault          TrialEndAction = "plan_default"
	TrialEndActionActivateSubscription TrialEndAction = "activate_subscription"
	TrialEndActionCancelSubscription   TrialEndAction = "cancel_subscription"
)
