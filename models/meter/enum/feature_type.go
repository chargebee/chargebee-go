package enum

type FeatureType string

const (
	FeatureTypeSwitch   FeatureType = "switch"
	FeatureTypeCustom   FeatureType = "custom"
	FeatureTypeQuantity FeatureType = "quantity"
	FeatureTypeRange    FeatureType = "range"
)
