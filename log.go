package wildcare

import (
	"encoding/json"
	"time"
)

type LogEntry struct {
	ID            string
	Action        string
	StaffMemberID string
	ReferenceID   string
	ReferenceName string
	Changes       string
	Date          time.Time
}

type LogReference interface {
	ReferenceID() string
	ReferenceName() string
}

type LogStorage interface {
	Create(*LogEntry) error
	AllByMember(id string) []*LogEntry
	AllByPatient(id string) []*LogEntry
}

func NewLogEntry(ref LogReference, member *StaffMember) (*LogEntry, error) {
	chg, err := json.Marshal(ref)

	if err != nil {
		return nil, err
	}

	l := &LogEntry{
		StaffMemberID: member.ID,
		ReferenceID:   ref.ReferenceID(),
		ReferenceName: ref.ReferenceName(),
		Changes:       string(chg),
		Date:          time.Now(),
	}

	return l, nil
}
