package enum

type Type string

const (
	TypeAllocation           Type = "allocation"
	TypeCapture              Type = "capture"
	TypeAuthorize            Type = "authorize"
	TypeReleaseAuthorization Type = "release_authorization"
	TypeCaptureAuthorization Type = "capture_authorization"
	TypeExpiry               Type = "expiry"
	TypeVoid                 Type = "void"
	TypeRollover             Type = "rollover"
	TypeAdjustment           Type = "adjustment"
)
