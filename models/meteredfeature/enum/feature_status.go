package enum

type FeatureStatus string

const (
	FeatureStatusActive   FeatureStatus = "active"
	FeatureStatusArchived FeatureStatus = "archived"
	FeatureStatusDraft    FeatureStatus = "draft"
)
