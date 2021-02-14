package services

import (
	"os"
	"strings"
)

type DnaValidator struct {
	chars []rune
}

func (v DnaValidator) ValidateDna(dnaMatrix []string) bool {
	return v.ValidateSize(dnaMatrix) && v.ValidateChars(dnaMatrix)
}

func (v DnaValidator) ValidateDnaDeep(dnaMatrix interface{}) bool {
	_, err := dnaMatrix.([]string)
	return err
}

func (v DnaValidator) ValidateSize(dnaMatrix []string) bool {
	size := len(dnaMatrix[0])

	for _, dna := range dnaMatrix {
		if len(dna) != size {
			return false
		}
		size = len(dna)
	}

	return true
}

func (v DnaValidator) ValidateChars(dnaMatrix []string) bool {
	validChars := []rune{'A', 'T', 'C', 'G'}

	chars := os.Getenv("VALID_CHARS")
	if chars != "" {
		validChars = []rune(strings.ToUpper(chars))
	}

	for _, dna := range dnaMatrix {
		for _, char := range strings.ToUpper(dna) {
			if !Contains(validChars, char) {
				return false
			}
		}
	}

	return true
}

func Contains(list []rune, value rune) bool {
	for _, char := range list {
		if char == value {
			return true
		}
	}

	return false
}
