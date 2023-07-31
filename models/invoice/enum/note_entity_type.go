package enum

type NoteEntityType string

const (
	NoteEntityTypePlan            NoteEntityType = "plan"
	NoteEntityTypeAddon           NoteEntityType = "addon"
	NoteEntityTypeCoupon          NoteEntityType = "coupon"
	NoteEntityTypeSubscription    NoteEntityType = "subscription"
	NoteEntityTypeCustomer        NoteEntityType = "customer"
	NoteEntityTypePlanItemPrice   NoteEntityType = "plan_item_price"
	NoteEntityTypeAddonItemPrice  NoteEntityType = "addon_item_price"
	NoteEntityTypeChargeItemPrice NoteEntityType = "charge_item_price"
	NoteEntityTypeTax             NoteEntityType = "tax"
)
