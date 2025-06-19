package enum

type DiscountApplyOnItemType string

const (
	DiscountApplyOnItemTypePlan   DiscountApplyOnItemType = "plan"
	DiscountApplyOnItemTypeAddon  DiscountApplyOnItemType = "addon"
	DiscountApplyOnItemTypeCharge DiscountApplyOnItemType = "charge"
)
