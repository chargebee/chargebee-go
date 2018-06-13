package enum

type State string

const (
	StateCreated      State = "created"
	StateRequested    State = "requested"
	StateSucceeded    State = "succeeded"
	StateCancelled    State = "cancelled"
	StateFailed       State = "failed"
	StateAcknowledged State = "acknowledged"
)
