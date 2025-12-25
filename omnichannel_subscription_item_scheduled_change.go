package chargebee

type OmnichannelSubscriptionItemScheduledChangeChangeType string

const (
	OmnichannelSubscriptionItemScheduledChangeChangeTypeDowngrade OmnichannelSubscriptionItemScheduledChangeChangeType = "downgrade"
	OmnichannelSubscriptionItemScheduledChangeChangeTypePause     OmnichannelSubscriptionItemScheduledChangeChangeType = "pause"
)

// just struct
type OmnichannelSubscriptionItemScheduledChange struct {
	Id                            string                                               `json:"id"`
	OmnichannelSubscriptionItemId string                                               `json:"omnichannel_subscription_item_id"`
	ScheduledAt                   int64                                                `json:"scheduled_at"`
	ChangeType                    OmnichannelSubscriptionItemScheduledChangeChangeType `json:"change_type"`
	CreatedAt                     int64                                                `json:"created_at"`
	ModifiedAt                    int64                                                `json:"modified_at"`
	ResourceVersion               int64                                                `json:"resource_version"`
	CurrentState                  *CurrentState                                        `json:"current_state"`
	ScheduledState                *ScheduledState                                      `json:"scheduled_state"`
	Object                        string                                               `json:"object"`
}

// sub resources
type OmnichannelSubscriptionItemScheduledChangeCurrentState struct {
	ItemIdAtSource string `json:"item_id_at_source"`
	Object         string `json:"object"`
}
type OmnichannelSubscriptionItemScheduledChangeScheduledState struct {
	ItemIdAtSource string `json:"item_id_at_source"`
	Object         string `json:"object"`
}

// operations
// input params
