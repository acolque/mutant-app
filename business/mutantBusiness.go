package business

import "github.com/mutant-app/services"

type MutantBusiness struct {
	myDetector services.IDetector
	myDb       services.IMutantDb
}

func NewMutantBusiness(detector services.IDetector, dataBase services.IMutantDb) *MutantBusiness {
	b := new(MutantBusiness)
	b.myDetector = detector
	b.myDb = dataBase
	return b
}

func (m MutantBusiness) IsMutant(dna []string) bool {
	result, _ := m.myDetector.IsMutant(dna)

	edna := services.EDna{Dna: dna, IsMutant: result}

	_, err := m.myDb.Find(edna)
	if err != nil {
		m.myDb.Add(edna)
	}

	return result
}
