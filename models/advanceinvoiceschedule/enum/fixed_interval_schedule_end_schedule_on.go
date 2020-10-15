package enum

type FixedIntervalScheduleEndScheduleOn string

const (
	FixedIntervalScheduleEndScheduleOnAfterNumberOfIntervals FixedIntervalScheduleEndScheduleOn = "after_number_of_intervals"
	FixedIntervalScheduleEndScheduleOnSpecificDate           FixedIntervalScheduleEndScheduleOn = "specific_date"
	FixedIntervalScheduleEndScheduleOnSubscriptionEnd        FixedIntervalScheduleEndScheduleOn = "subscription_end"
)
