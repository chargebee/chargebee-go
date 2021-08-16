package enum

type ChangeOption string

const (
	ChangeOptionEndOfTerm    ChangeOption = "end_of_term"
	ChangeOptionSpecificDate ChangeOption = "specific_date"
	ChangeOptionImmediately  ChangeOption = "immediately"
)
