package wildcare

type Species struct {
	ID   string
	Code string
	Name string
}

type SpeciesDB interface {
	All() []*Species
	Create(*Species) error
	Find(id string) (*Species, error)
	FindByCode(code string) (*Species, error)
}
