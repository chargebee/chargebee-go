package enum

type ItemConstraintConstraint string

const (
	ItemConstraintConstraintNone     ItemConstraintConstraint = "none"
	ItemConstraintConstraintAll      ItemConstraintConstraint = "all"
	ItemConstraintConstraintSpecific ItemConstraintConstraint = "specific"
	ItemConstraintConstraintCriteria ItemConstraintConstraint = "criteria"
)
