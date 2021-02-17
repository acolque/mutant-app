package test

import (
	"testing"

	"github.com/mutant-app/business"
	"github.com/mutant-app/services"
	"github.com/stretchr/testify/assert"
)

func TestValidateBusinessIsMutantSucess(t *testing.T) {
	detector := new(services.DnaMutantDetector)
	db := new(services.MutantMockDb)
	myBusiness := business.NewMutantBusiness(detector, db)
	dna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

	result, _ := myBusiness.IsMutant(dna)

	assert.True(t, result, "fallo test TestValidateBusinessIsMutantSucess")
}

func TestValidateBusinessIsMutantError(t *testing.T) {
	detector := new(services.DnaMutantDetector)
	db := new(services.MutantMockDb)
	myBusiness := business.NewMutantBusiness(detector, db)
	dna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAATG", "ACCCTA", "TCACTG"}

	result, _ := myBusiness.IsMutant(dna)

	assert.False(t, result, "fallo test TestValidateBusinessIsMutantError")
}

func TestValidateBusinessIsMutantSuccessWithDBError(t *testing.T) {
	detector := new(services.DnaMutantDetector)
	db := new(services.MutantMongodb)
	myBusiness := business.NewMutantBusiness(detector, db)
	dna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

	_, msg := myBusiness.IsMutant(dna)

	assert.NotEmpty(t, msg, "fallo test TestValidateBusinessIsMutantSuccessWithDBError")
}
