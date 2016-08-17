package wildcare

type Species struct {
	ID   string
	Name string
}

type SpeciesDB interface {
	All() []*Species
	Create(*Species) error
	Find(id string) (*Species, error)
	FindByName(name string) (*Species, error)
}
