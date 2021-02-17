package services

import (
	"strings"
)

type DnaMutantDetector struct {
}

func (d DnaMutantDetector) IsMutant(dna []string) (result bool, err error) {

	defer catchError(&err)

	matrix := convertToMatrix(dna)

	wordsToFind := []string{"AAAA", "TTTT", "CCCC", "GGGG"}
	minWordsToFind := 2

	mapWords := make(map[rune]string)
	for _, word := range wordsToFind {
		mapWords[rune(word[0])] = word
	}

	for j, row := range matrix {
		for i := range row {
			firstChar := matrix[j][i]
			if _, ok := mapWords[firstChar]; !ok {
				continue
			}

			for key, wordToFind := range mapWords {
				if findTo(wordToFind, i, j, 1, 0, matrix) ||
					findTo(wordToFind, i, j, 1, 1, matrix) ||
					findTo(wordToFind, i, j, 1, -1, matrix) ||
					findTo(wordToFind, i, j, 0, 1, matrix) {
					delete(mapWords, key)
					minWordsToFind--
					break
				}
			}

			if minWordsToFind == 0 {
				result = true
				return
			}
		}
	}

	result = false
	return
}

func convertToMatrix(arr []string) [][]rune {
	var matrix [][]rune
	for _, row := range arr {
		matrix = append(matrix, []rune(strings.ToUpper(row)))
	}
	return matrix
}

func findTo(wordToFind string, posX, posY, moveR, moveL int, matrix [][]rune) bool {
	lenCrr := 0
	lenFind, lenX, lenY := len(wordToFind), len(matrix[0]), len(matrix)
	var word string

	for lenCrr < lenFind && posX > -1 && posY > -1 && posX < lenX && posY < lenY {
		word += string(matrix[posY][posX])
		posX += moveR
		posY += moveL
		lenCrr++

		if !strings.Contains(wordToFind, word) {
			return false
		}
	}

	return word == wordToFind
}

func catchError(err *error) {
	if r := recover(); r != nil {
		*err = r.(error)
	}
}
