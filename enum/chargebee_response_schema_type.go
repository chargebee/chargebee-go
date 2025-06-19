package enum

type ChargebeeResponseSchemaType string

const (
	ChargebeeResponseSchemaTypePlansAddons ChargebeeResponseSchemaType = "plans_addons"
	ChargebeeResponseSchemaTypeItems       ChargebeeResponseSchemaType = "items"
	ChargebeeResponseSchemaTypeCompat      ChargebeeResponseSchemaType = "compat"
)
