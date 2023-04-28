package enum

type CouponConstraintType string

const (
	CouponConstraintTypeMaxRedemptions CouponConstraintType = "max_redemptions"
	CouponConstraintTypeUniqueBy       CouponConstraintType = "unique_by"
)
