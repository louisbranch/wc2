package admission

import (
	"time"

	"github.com/larissavoigt/wildcare"
)

type Admission struct {
	patientStorage wildcare.PatientStorage
	logStorage     wildcare.LogStorage
}

func (a *Admission) New(p *wildcare.Patient, m *wildcare.StaffMember) error {
	p.Admission = time.Now()

	err := a.patientStorage.Create(p)

	if err != nil {
		return err
	}

	e, err := wildcare.NewLogEntry(p, m)

	if err != nil {
		return err
	}

	err = a.logStorage.Create(e)

	return err
}
