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
