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

type PatientStorage interface {
	All(PatientStatus) []*Patient
	Create(*Patient) error
	Find(id string) (*Patient, error)
}

type Admission interface {
	New(*Patient, *StaffMember) error
}

type Release interface {
	New(*Patient, *StaffMember) error
}

func (p *Patient) ReferenceID() string {
	return p.ID
}

func (p *Patient) ReferenceName() string {
	return "Patient"
}
