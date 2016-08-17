package wildcare

import "time"

type ActivityType int

const (
	RegisterMember ActivityType = iota
	AdmitPatient
	ReleasePatient
	CreateDiet
	EditDiet
)

type ActivityLog struct {
	Type          ActivityType
	StaffMemberID string
	ReferenceID   string
	Changes       string
	Date          time.Time
}
