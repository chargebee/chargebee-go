package enum

type ExecuteMode string

const (
	ExecuteModeStopOnFirstTrue  ExecuteMode = "stop_on_first_true"
	ExecuteModeStopOnFirstFalse ExecuteMode = "stop_on_first_false"
	ExecuteModeExecuteAll       ExecuteMode = "execute_all"
	ExecuteModeExecuteAllTrue   ExecuteMode = "execute_all_true"
)
