package enum

type ApplyDiscountOn string

const (
	ApplyDiscountOnPlans             ApplyDiscountOn = "plans"
	ApplyDiscountOnPlansAndAddons    ApplyDiscountOn = "plans_and_addons"
	ApplyDiscountOnPlansWithQuantity ApplyDiscountOn = "plans_with_quantity"
	ApplyDiscountOnNotApplicable     ApplyDiscountOn = "not_applicable"
)
