package enum

type Type string

const (
	TypeCheckoutNew          Type = "checkout_new"
	TypeCheckoutExisting     Type = "checkout_existing"
	TypeUpdatePaymentMethod  Type = "update_payment_method"
	TypeManagePaymentSources Type = "manage_payment_sources"
	TypeCollectNow           Type = "collect_now"
	TypeExtendSubscription   Type = "extend_subscription"
	TypeCheckoutOneTime      Type = "checkout_one_time"
	TypePreCancel            Type = "pre_cancel"
	TypeViewVoucher          Type = "view_voucher"
	TypeCheckoutGift         Type = "checkout_gift"
	TypeClaimGift            Type = "claim_gift"
	TypeUpdateCard           Type = "update_card"
)
