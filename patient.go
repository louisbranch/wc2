package wildcare

import "time"

type PatientStatus int

const (
	PatientInTreatment PatientStatus = iota
	PatientReleased
	PatientDeceased
)

type Patient struct {
	ID        string
	SpeciesID string
	Notes     string
	Status    PatientStatus
	Admission time.Time
	Discharge time.Time
}

type PatientDB interface {
	All(PatientStatus) []*Patient
	Create(*Patient) error
	Find(id string) (*Patient, error)
}
