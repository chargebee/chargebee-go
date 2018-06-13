package enum

type NoteEntityType string

const (
	NoteEntityTypePlan         NoteEntityType = "plan"
	NoteEntityTypeAddon        NoteEntityType = "addon"
	NoteEntityTypeCoupon       NoteEntityType = "coupon"
	NoteEntityTypeSubscription NoteEntityType = "subscription"
	NoteEntityTypeCustomer     NoteEntityType = "customer"
)
