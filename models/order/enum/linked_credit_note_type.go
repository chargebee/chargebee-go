package enum

type LinkedCreditNoteType string

const (
	LinkedCreditNoteTypeAdjustment LinkedCreditNoteType = "adjustment"
	LinkedCreditNoteTypeRefundable LinkedCreditNoteType = "refundable"
)
