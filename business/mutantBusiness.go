package business

import "github.com/mutant-app/services"

type MutantBusiness struct {
	myDetector services.IDetector
}

func NewMutantBusiness(detector services.IDetector) *MutantBusiness {
	b := new(MutantBusiness)
	b.myDetector = detector
	return b
}

func (m MutantBusiness) IsMutant(dna []string) bool {
	result, _ := m.myDetector.IsMutant(dna)
	return result
}
