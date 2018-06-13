package enum

type Role string

const (
	RolePrimary Role = "primary"
	RoleBackup  Role = "backup"
	RoleNone    Role = "none"
)
