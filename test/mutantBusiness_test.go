package test

import (
	"testing"

	"github.com/mutant-app/business"
	"github.com/mutant-app/services"
	"github.com/stretchr/testify/assert"
)

func TestValidateBusinessIsMutantSucess(t *testing.T) {
	detector := new(services.DnaMutantDetector)
	myBusiness := business.NewMutantBusiness(detector)
	dna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

	result := myBusiness.IsMutant(dna)

	assert.True(t, result, "fallo test TestValidateBusinessIsMutantSucess")
}

func TestValidateBusinessIsMutantError(t *testing.T) {
	detector := new(services.DnaMutantDetector)
	myBusiness := business.NewMutantBusiness(detector)
	dna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAATG", "ACCCTA", "TCACTG"}

	result := myBusiness.IsMutant(dna)

	assert.False(t, result, "fallo test TestValidateBusinessIsMutantError")
}
