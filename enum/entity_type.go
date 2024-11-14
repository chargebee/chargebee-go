package enum

type EntityType string

const (
	EntityTypeCustomer                    EntityType = "customer"
	EntityTypeSubscription                EntityType = "subscription"
	EntityTypeCoupon                      EntityType = "coupon"
	EntityTypePlanItemPrice               EntityType = "plan_item_price"
	EntityTypeAddonItemPrice              EntityType = "addon_item_price"
	EntityTypeChargeItemPrice             EntityType = "charge_item_price"
	EntityTypeInvoice                     EntityType = "invoice"
	EntityTypeQuote                       EntityType = "quote"
	EntityTypeCreditNote                  EntityType = "credit_note"
	EntityTypeTransaction                 EntityType = "transaction"
	EntityTypePlan                        EntityType = "plan"
	EntityTypeAddon                       EntityType = "addon"
	EntityTypeOrder                       EntityType = "order"
	EntityTypeItemFamily                  EntityType = "item_family"
	EntityTypeItem                        EntityType = "item"
	EntityTypeItemPrice                   EntityType = "item_price"
	EntityTypePlanItem                    EntityType = "plan_item"
	EntityTypeAddonItem                   EntityType = "addon_item"
	EntityTypeChargeItem                  EntityType = "charge_item"
	EntityTypePlanPrice                   EntityType = "plan_price"
	EntityTypeAddonPrice                  EntityType = "addon_price"
	EntityTypeChargePrice                 EntityType = "charge_price"
	EntityTypeDifferentialPrice           EntityType = "differential_price"
	EntityTypeAttachedItem                EntityType = "attached_item"
	EntityTypeFeature                     EntityType = "feature"
	EntityTypeSubscriptionEntitlement     EntityType = "subscription_entitlement"
	EntityTypeItemEntitlement             EntityType = "item_entitlement"
	EntityTypeBusinessEntity              EntityType = "business_entity"
	EntityTypePriceVariant                EntityType = "price_variant"
	EntityTypeOmnichannelSubscription     EntityType = "omnichannel_subscription"
	EntityTypeOmnichannelSubscriptionItem EntityType = "omnichannel_subscription_item"
	EntityTypeOmnichannelTransaction      EntityType = "omnichannel_transaction"
)
