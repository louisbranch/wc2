package wildcare

import "time"

type ActivityType int

type ActivityLog struct {
	Type          ActivityType
	StaffMemberID string
	ReferenceID   string
	Changes       string
	CreatedAt     time.Time
}
