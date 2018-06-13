package enum

type PlanConstraint string

const (
	PlanConstraintNone          PlanConstraint = "none"
	PlanConstraintAll           PlanConstraint = "all"
	PlanConstraintSpecific      PlanConstraint = "specific"
	PlanConstraintNotApplicable PlanConstraint = "not_applicable"
)
