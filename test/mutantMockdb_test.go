package test

import (
	"testing"

	"github.com/mutant-app/services"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	mydb := services.NewMutantMockDb()
	item := services.EDna{Dna: []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}, IsMutant: false}

	mydb.Add(item)
	dnas, err := mydb.GetAll()
	_ = err

	assert.Equal(t, 1, len(dnas), "fallo test TestAdd")
}

func TestFind(t *testing.T) {
	mydb := services.NewMutantMockDb()
	item := services.EDna{Dna: []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTT"}, IsMutant: false}
	item2 := services.EDna{Dna: []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}, IsMutant: false}
	item3 := services.EDna{Dna: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}, IsMutant: true}
	mydb.Add(item)
	mydb.Add(item2)
	mydb.Add(item3)

	dna, err := mydb.Find(item2)
	_ = err

	assert.Equal(t, item2.Dna, dna.Dna, "fallo test Find")
}

func TestDelete(t *testing.T) {
	mydb := services.NewMutantMockDb()
	item := services.EDna{Dna: []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTT"}, IsMutant: false}
	item2 := services.EDna{Dna: []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}, IsMutant: false}
	item3 := services.EDna{Dna: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}, IsMutant: true}
	mydb.Add(item)
	mydb.Add(item2)
	mydb.Add(item3)

	mydb.Delete(item3)
	dnas, _ := mydb.GetAll()

	assert.Equal(t, 2, len(dnas), "fallo test Delete")
}

func TestFindNotExist(t *testing.T) {
	mydb := services.NewMutantMockDb()
	item := services.EDna{Dna: []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTT"}, IsMutant: false}
	item2 := services.EDna{Dna: []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}, IsMutant: false}
	item3 := services.EDna{Dna: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}, IsMutant: true}
	mydb.Add(item2)
	mydb.Add(item3)

	_, err := mydb.Find(item)

	assert.NotNil(t, err, "fallo test TestFindNotExist")
}
