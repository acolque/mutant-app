package services

type EDna struct {
	Dna      []string
	IsMutant bool
}

type IMutantDb interface {
	Find(EDna) (EDna, error)
	Add(EDna) error
	Update(EDna, EDna) error
	Delete(EDna) error
	GetAll() ([]EDna, error)
}
