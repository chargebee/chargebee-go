package enum

type CancellationReason string

const (
	CancellationReasonShippingCutOffPassed   CancellationReason = "shipping_cut_off_passed"
	CancellationReasonProductUnsatisfactory  CancellationReason = "product_unsatisfactory"
	CancellationReasonThirdPartyCancellation CancellationReason = "third_party_cancellation"
	CancellationReasonProductNotRequired     CancellationReason = "product_not_required"
	CancellationReasonDeliveryDateMissed     CancellationReason = "delivery_date_missed"
	CancellationReasonAlternativeFound       CancellationReason = "alternative_found"
	CancellationReasonInvoiceWrittenOff      CancellationReason = "invoice_written_off"
	CancellationReasonInvoiceVoided          CancellationReason = "invoice_voided"
	CancellationReasonFraudulentTransaction  CancellationReason = "fraudulent_transaction"
	CancellationReasonPaymentDeclined        CancellationReason = "payment_declined"
	CancellationReasonSubscriptionCancelled  CancellationReason = "subscription_cancelled"
	CancellationReasonProductNotAvailable    CancellationReason = "product_not_available"
	CancellationReasonOthers                 CancellationReason = "others"
	CancellationReasonOrderResent            CancellationReason = "order_resent"
)
