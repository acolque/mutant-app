package handlers

import (
	"net/http"
)

type StatResult struct {
	Count_mutant_dna int
	Count_human_dna  int
	Ratio            float32
}

func Stats(w http.ResponseWriter, r *http.Request) {
	result := StatResult{
		Count_mutant_dna: 40,
		Count_human_dna:  100,
		Ratio:            0.4,
	}

	respondWithJSON(w, http.StatusOK, result)
}
