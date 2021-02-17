package test

import (
	"testing"

	"github.com/mutant-app/business"
	"github.com/mutant-app/services"
	"github.com/stretchr/testify/assert"
)

func TestValidateBusinessGetStats(t *testing.T) {
	db := new(services.MutantMockDb)
	myBusiness := business.NewStatsBusiness(db)

	result, _ := myBusiness.GetStats()

	assert.NotNil(t, result, "fallo test TestValidateBusinessGetStats")
}

func TestValidateBusinessGetStatsWithOneMutantsTwoHuman(t *testing.T) {
	db := new(services.MutantMockDb)
	detector := new(services.DnaMutantDetector)
	mutantBusiness := business.NewMutantBusiness(detector, db)
	myBusiness := business.NewStatsBusiness(db)
	mutant1 := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	human1 := []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}
	human2 := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAGTG", "TCCCTA", "TCACTG"}

	mutantBusiness.IsMutant(mutant1)
	mutantBusiness.IsMutant(human1)
	mutantBusiness.IsMutant(human2)
	result, _ := myBusiness.GetStats()

	assert.True(t,
		result.Count_mutant_dna == 1 &&
			result.Count_human_dna == 2 &&
			result.Ratio == 0.5,
		"fallo test TestValidateBusinessGetStatsWithOneMutantsTwoHuman")
}
