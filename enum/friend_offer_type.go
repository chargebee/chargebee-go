package enum

type FriendOfferType string

const (
	FriendOfferTypeNone       FriendOfferType = "none"
	FriendOfferTypeCoupon     FriendOfferType = "coupon"
	FriendOfferTypeCouponCode FriendOfferType = "coupon_code"
)
