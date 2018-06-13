package enum

type Status string

const (
	StatusCreated         Status = "created"
	StatusLoggedIn        Status = "logged_in"
	StatusLoggedOut       Status = "logged_out"
	StatusNotYetActivated Status = "not_yet_activated"
	StatusActivated       Status = "activated"
)
