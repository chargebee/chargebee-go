package enum

type StoreStatus string

const (
	StoreStatusInTrial   StoreStatus = "in_trial"
	StoreStatusActive    StoreStatus = "active"
	StoreStatusCancelled StoreStatus = "cancelled"
	StoreStatusPaused    StoreStatus = "paused"
)
