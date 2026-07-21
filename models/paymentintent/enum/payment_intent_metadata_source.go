package enum

type PaymentIntentMetadataSource string

const (
	PaymentIntentMetadataSourceCbJs             PaymentIntentMetadataSource = "cb_js"
	PaymentIntentMetadataSourceComponentsFields PaymentIntentMetadataSource = "components_fields"
	PaymentIntentMetadataSourceCheckoutV3       PaymentIntentMetadataSource = "checkout_v3"
	PaymentIntentMetadataSourcePaynowV3         PaymentIntentMetadataSource = "paynow_v3"
	PaymentIntentMetadataSourcePortalV3         PaymentIntentMetadataSource = "portal_v3"
	PaymentIntentMetadataSourceGiftV3           PaymentIntentMetadataSource = "gift_v3"
	PaymentIntentMetadataSourceCheckoutV4       PaymentIntentMetadataSource = "checkout_v4"
	PaymentIntentMetadataSourcePaymentComponent PaymentIntentMetadataSource = "payment_component"
	PaymentIntentMetadataSourcePcInappV4        PaymentIntentMetadataSource = "pc_inapp_v4"
	PaymentIntentMetadataSourcePcFpcV4          PaymentIntentMetadataSource = "pc_fpc_v4"
)
