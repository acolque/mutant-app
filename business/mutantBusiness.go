package business

type MutantBusiness struct {
}

func (m MutantBusiness) IsMutant(dna []string) bool {
	return true
}
