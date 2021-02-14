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
