package wildcare

import "time"

type PatientStatus int

const (
	InTreatment PatientStatus = iota
	Released
	Deceased
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
	Create(*Patient) error
	Find(id string) (*Patient, error)
	AllAdmitted() []*Patient
}
