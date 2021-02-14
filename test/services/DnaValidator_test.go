package test

import (
	"testing"

	"github.com/mutant-app/services"
	"github.com/stretchr/testify/assert"
)

func TestValidateSizeSuccess(t *testing.T) {
	dna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	v := new(services.DnaValidator)
	result := v.ValidateSize(dna)
	assert.True(t, result, "fallo test TestValidateSize")
}

func TestValidateSizeError(t *testing.T) {
	dna := []string{"ATG", "CAGTGC", "TTATGT", "AGG", "CCCCTA", "TCACTG"}
	v := new(services.DnaValidator)
	result := v.ValidateSize(dna)
	assert.False(t, result, "fallo test TestValidateSizeError")
}

func TestValidateCharsSuccess(t *testing.T) {
	dna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	v := new(services.DnaValidator)
	result := v.ValidateChars(dna)
	assert.True(t, result, "fallo test TestValidateCharsSuccess")
}

func TestValidateCharsError(t *testing.T) {
	dna := []string{"ATGCGA", "XXXYYY", "TTATGT", "AGAAGG", "CCCRRR", "TCACTG"}
	v := new(services.DnaValidator)
	result := v.ValidateChars(dna)
	assert.False(t, result, "fallo test TestValidateCharsError")
}

func TestValidateDNASuccess(t *testing.T) {
	dna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	v := new(services.DnaValidator)
	result := v.ValidateDna(dna)
	assert.True(t, result, "fallo test TestValidateDNASuccess")
}

func TestValidateDNAError(t *testing.T) {
	dna := []string{"ATG", "CAGTGC", "TTATGT", "AGG", "CCCCTA", "TCACTG"}
	v := new(services.DnaValidator)
	result := v.ValidateDna(dna)
	assert.False(t, result, "fallo test TestValidateDNAError")
}

func TestValidateDNADeepError(t *testing.T) {
	dna := "ATG"
	v := new(services.DnaValidator)
	result := v.ValidateDnaDeep(dna)
	assert.False(t, result, "fallo test TestValidateDNADeepError")
}
