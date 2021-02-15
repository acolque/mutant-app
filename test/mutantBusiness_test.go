package test

import (
	"testing"

	"github.com/mutant-app/business"
	"github.com/stretchr/testify/assert"
)

func TestValidateBusiness(t *testing.T) {
	m := new(business.MutantBusiness)
	dna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

	result := m.IsMutant(dna)

	assert.True(t, result, "fallo test TestValidateBusiness")
}
