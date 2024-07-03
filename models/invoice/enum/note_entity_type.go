package enum

type NoteEntityType string

const (
	NoteEntityTypeCoupon          NoteEntityType = "coupon"
	NoteEntityTypeSubscription    NoteEntityType = "subscription"
	NoteEntityTypeCustomer        NoteEntityType = "customer"
	NoteEntityTypePlanItemPrice   NoteEntityType = "plan_item_price"
	NoteEntityTypeAddonItemPrice  NoteEntityType = "addon_item_price"
	NoteEntityTypeChargeItemPrice NoteEntityType = "charge_item_price"
	NoteEntityTypeTax             NoteEntityType = "tax"
	NoteEntityTypePlan            NoteEntityType = "plan"
	NoteEntityTypeAddon           NoteEntityType = "addon"
)
