package enum

type EndScheduleOn string

const (
	EndScheduleOnAfterNumberOfIntervals EndScheduleOn = "after_number_of_intervals"
	EndScheduleOnSpecificDate           EndScheduleOn = "specific_date"
	EndScheduleOnSubscriptionEnd        EndScheduleOn = "subscription_end"
)
