package wildcare

type StaffMember struct {
	ID    string
	Name  string
	Email string
	Psswd string
	Phone string
	Type  StaffMemberType
}

type StaffMemberType int

type StaffMemberDB interface {
	All() []*StaffMember
	Create(*StaffMember) error
	Find(id string) (*StaffMember, error)
	FindByEmail(email string) (*StaffMember, error)
}
