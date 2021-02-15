package test

import (
	"testing"

	"github.com/mutant-app/services"

	"github.com/stretchr/testify/assert"
)

func TestGenerateWelcomeMsg(t *testing.T) {
	p := new(services.Person)
	result := p.GetName()
	assert.Equal(t, "Arnold", result, "fallo test GenerateWelcomeMsg")
}

// func TestReadENV(t *testing.T) {
// 	var chars string

// 	json.Unmarshal([]byte(`"ATCG"`), &chars)

// 	result := []rune(chars)
// 	expec := []rune{'A', 'T', 'C', 'G'}

// 	assert.Equal(t, expec, result, "fallo test TestReadENV")
// }

// func TestMongodb(t *testing.T) {
// 	db := services.NewMutantMongodb()
// 	item := services.EDna{Dna: []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}, IsMutant: true}
// 	item2 := services.EDna{Dna: []string{"XXXXXX", "CAGTGC", "XXXXXX", "AGACGG", "XXXXXX", "TCACTG"}, IsMutant: true}

// 	result, _ := db.Find(item)
// 	db.Add(item2)
// 	db.Delete(item2)
// 	dnas, _ := db.GetAll()
// 	_ = dnas

// 	assert.NotNil(t, result, "fallo test TestMongodb")
// }
