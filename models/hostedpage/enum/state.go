package enum

type State string

const (
	StateCreated      State = "created"
	StateRequested    State = "requested"
	StateSucceeded    State = "succeeded"
	StateCancelled    State = "cancelled"
	StateAcknowledged State = "acknowledged"
	StateFailed       State = "failed"
)
