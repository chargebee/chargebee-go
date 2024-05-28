package enum

type LinkedCreditNoteStatus string

const (
	LinkedCreditNoteStatusAdjusted  LinkedCreditNoteStatus = "adjusted"
	LinkedCreditNoteStatusRefunded  LinkedCreditNoteStatus = "refunded"
	LinkedCreditNoteStatusRefundDue LinkedCreditNoteStatus = "refund_due"
	LinkedCreditNoteStatusVoided    LinkedCreditNoteStatus = "voided"
)
