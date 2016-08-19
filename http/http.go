package http

import (
	"net/http"
	"strconv"

	"github.com/larissavoigt/wildcare"
)

type Server struct {
	admission wildcare.Admission
}

func (s *Server) NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/patients/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			r.ParseForm()
			species := r.Form.Get("species_id")
			notes := r.Form.Get("notes")
			status := r.Form.Get("patient_status")
			i64, err := strconv.ParseInt(status, 10, 64)

			if err != nil {

			}

			patient := &wildcare.Patient{
				SpeciesID: species,
				Notes:     notes,
				Status:    wildcare.PatientStatus(i64),
			}

			staff := &wildcare.StaffMember{}

			err = s.admission.New(patient, staff)

			if err != nil {

			}

		}
	})

	return mux
}
