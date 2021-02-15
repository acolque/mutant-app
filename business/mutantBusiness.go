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

func (m MutantBusiness) IsMutant(dna []string) (result bool, msg string) {
	result, err := m.myDetector.IsMutant(dna)
	if err != nil {
		msg = "warn: Ocurrio un error en la validacion"
		return
	}

	edna := services.EDna{Dna: dna, IsMutant: result}
	find, err := m.myDb.Find(edna)
	if _, ok := err.(*services.ConexError); ok {
		msg = "warn: Sin conex a Db. No se pudo guardar para estadistica"
		return
	}

	if find.Dna == nil {
		m.myDb.Add(edna)
	}

	return
}
