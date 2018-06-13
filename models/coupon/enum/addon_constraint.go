package enum

type AddonConstraint string

const (
	AddonConstraintNone          AddonConstraint = "none"
	AddonConstraintAll           AddonConstraint = "all"
	AddonConstraintSpecific      AddonConstraint = "specific"
	AddonConstraintNotApplicable AddonConstraint = "not_applicable"
)
