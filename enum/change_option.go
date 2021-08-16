package enum

type ChangeOption string

const (
	ChangeOptionImmediately  ChangeOption = "immediately"
	ChangeOptionEndOfTerm    ChangeOption = "end_of_term"
	ChangeOptionSpecificDate ChangeOption = "specific_date"
)
