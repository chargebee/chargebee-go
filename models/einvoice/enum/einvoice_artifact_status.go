package enum

type EinvoiceArtifactStatus string

const (
	EinvoiceArtifactStatusScheduled  EinvoiceArtifactStatus = "scheduled"
	EinvoiceArtifactStatusSkipped    EinvoiceArtifactStatus = "skipped"
	EinvoiceArtifactStatusInProgress EinvoiceArtifactStatus = "in_progress"
	EinvoiceArtifactStatusSuccess    EinvoiceArtifactStatus = "success"
	EinvoiceArtifactStatusFailed     EinvoiceArtifactStatus = "failed"
	EinvoiceArtifactStatusRegistered EinvoiceArtifactStatus = "registered"
)
