package business

import "github.com/mutant-app/services"

type StatsBusiness struct {
	myDb services.IMutantDb
}

type StatResult struct {
	Count_mutant_dna int
	Count_human_dna  int
	Ratio            float64
}

func NewStatsBusiness(db services.IMutantDb) *StatsBusiness {
	b := new(StatsBusiness)
	b.myDb = db
	return b
}

func (s StatsBusiness) GetStats() (StatResult, error) {
	dnas, err := s.myDb.GetAll()
	if _, ok := err.(*services.ConexError); ok {
		return StatResult{}, &services.ConexError{}
	}

	countMutant, countHuman, ratio := 0, 0, 0.0

	for _, dna := range dnas {
		if dna.IsMutant {
			countMutant++
		} else {
			countHuman++
		}
	}

	ratio = float64(countMutant) / float64(countHuman)

	result := StatResult{
		Count_mutant_dna: countMutant,
		Count_human_dna:  countHuman,
		Ratio:            ratio,
	}

	return result, nil
}
