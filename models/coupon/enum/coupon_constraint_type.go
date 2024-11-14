package enum

type CouponConstraintType string

const (
	CouponConstraintTypeMaxRedemptions   CouponConstraintType = "max_redemptions"
	CouponConstraintTypeUniqueBy         CouponConstraintType = "unique_by"
	CouponConstraintTypeExistingCustomer CouponConstraintType = "existing_customer"
	CouponConstraintTypeNewCustomer      CouponConstraintType = "new_customer"
)
